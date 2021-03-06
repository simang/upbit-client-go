package upbit

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
)

// Upbit API client
type Upbit struct {
	apiUrl     string
	accessKey  string
	secretKey  string
	httpClient *http.Client
	logger     *logrus.Logger
}

// NewUpbit
func NewUpbit(
	apiUrl string, accessKey string, secretKey string, httpClient *http.Client, logger *logrus.Logger) *Upbit {
	if apiUrl == "" {
		apiUrl = "https://api.upbit.com"
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Upbit{
		apiUrl:     apiUrl,
		accessKey:  accessKey,
		secretKey:  secretKey,
		httpClient: httpClient,
		logger:     logger,
	}
}

func (u *Upbit) Accounts() ([]Account, error) {
	path := "/v1/accounts"
	var model []Account
	err := u.callApi(http.MethodGet, path, nil, &model, true)
	if err != nil {
		return nil, u.errorf("Failed to call api %s: %s", path, err.Error())
	}
	return model, nil
}

func (u *Upbit) Chance(param *ChanceParam) (*Chance, error) {
	path := "/v1/orders/chance"
	var model Chance
	err := u.callApi(http.MethodGet, path, param, &model, true)
	if err != nil {
		return nil, u.errorf("Failed to call api %s: %s", path, err.Error())
	}
	return &model, nil
}

func (u *Upbit) Order(param *OrderParam) (*OrderDetail, error) {
	path := "/v1/order"
	var model OrderDetail
	err := u.callApi(http.MethodGet, path, param, &model, true)
	if err != nil {
		return nil, u.errorf("Failed to call api %s: %s", path, err.Error())
	}
	return &model, nil
}

func (u *Upbit) CancelOrder(param *OrderParam) (*OrderDetail, error) {
	path := "/v1/order"
	var model OrderDetail
	err := u.callApi(http.MethodDelete, path, param, &model, true)
	if err != nil {
		return nil, u.errorf("Failed to call api %s: %s", path, err.Error())
	}
	return &model, nil
}

func (u *Upbit) Orders(param *OrdersParam) ([]Order, error) {
	path := "/v1/orders"
	var model []Order
	err := u.callApi(http.MethodGet, path, param, &model, true)
	if err != nil {
		return nil, u.errorf("Failed to call api %s: %s", path, err.Error())
	}
	return model, nil
}

func (u *Upbit) CreateOrders(param *PostOrdersParam) (*Order, error) {
	path := "/v1/orders"
	var model Order
	err := u.callApi(http.MethodPost, path, param, &model, true)
	if err != nil {
		return nil, u.errorf("Failed to call api %s: %s", path, err.Error())
	}
	return &model, nil
}

// QUOTATION
func (u *Upbit) AllMarket(param *AllMarketParam) ([]Market, error) {
	path := "/v1/market/all"
	var model []Market
	err := u.callApi(http.MethodGet, path, param, &model, false)
	if err != nil {
		return nil, u.errorf("Failed to call api %s: %s", path, err.Error())
	}
	return model, nil
}

func (u *Upbit) Ticker(param *TickerParam) ([]Ticker, error) {
	path := "/v1/ticker"
	var model []Ticker
	err := u.callApi(http.MethodGet, path, param, &model, false)
	if err != nil {
		return nil, u.errorf("Failed to call api %s: %s", path, err.Error())
	}
	return model, nil
}

func (u *Upbit) callApi(method string, url string, params interface{}, model interface{}, withAuth bool) error {
	nonce := uuid.New().String()
	claims := jwt.MapClaims{}
	claims["access_key"] = u.accessKey
	claims["nonce"] = nonce

	var encodedQuery string
	if params != nil {
		v, err := query.Values(params)
		if err != nil {
			return err
		}
		encodedQuery = v.Encode()
		u.logger.Debugf("encoded query: %+v", encodedQuery)
		if withAuth {
			h := sha512.New()
			_, err = h.Write([]byte(encodedQuery))
			if err != nil {
				return err
			}
			claims["query_hash"] = hex.EncodeToString(h.Sum(nil))
			claims["query_hash_alg"] = "SHA512"
		}
	}

	u.logger.Debugf("claims: %+v", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.secretKey))
	if err != nil {
		return u.errorf("Failed to get signed token : %s", err.Error())
	}

	var request *http.Request
	if method == http.MethodPost || method == http.MethodPut {
		if params != nil {
			body, err := json.Marshal(params)
			if err != nil {
				return err
			}
			u.logger.Debugf("body: %s", string(body))
			request, err = http.NewRequest(method, fmt.Sprintf("%s%s", u.apiUrl, url), bytes.NewBuffer(body))
		} else {
			request, err = http.NewRequest(method, fmt.Sprintf("%s%s", u.apiUrl, url), nil)
		}
	} else {
		if encodedQuery != "" {
			encodedQuery = fmt.Sprintf("?%s", encodedQuery)
		}
		request, err = http.NewRequest(method, fmt.Sprintf("%s%s%s", u.apiUrl, url, encodedQuery), nil)
	}
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))

	response, err := u.httpClient.Do(request)
	if err != nil {
		return err
	}
	body := response.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	u.logger.Debugf("Response from api status %d: %s", response.StatusCode, string(data))
	if response.StatusCode >= 300 || response.StatusCode < 200 {
		var msg ErrorMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return u.errorf("Failed to unmarshal response error: %s", err.Error())
		}
		return u.errorf("HTTP status %d: %+v", response.StatusCode, msg)
	}

	if err := json.Unmarshal(data, model); err != nil {
		return u.errorf("Failed to unmarshal response error: %s", err.Error())
	}
	return nil
}

func (u *Upbit) errorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	u.logger.Errorf(msg)
	return errors.New(msg)
}
