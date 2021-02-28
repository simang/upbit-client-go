package upbit

import "time"

type Account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

type Accounts []Account

type Chance struct {
	BidFee string `json:"bid_fee"`
	AskFee string `json:"ask_fee"`
	Market struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		OrderTypes []string `json:"order_types"`
		OrderSides []string `json:"order_sides"`
		Bid        struct {
			Currency  string      `json:"currency"`
			PriceUnit interface{} `json:"price_unit"`
			MinTotal  int         `json:"min_total"`
		} `json:"bid"`
		Ask struct {
			Currency  string      `json:"currency"`
			PriceUnit interface{} `json:"price_unit"`
			MinTotal  int         `json:"min_total"`
		} `json:"ask"`
		MaxTotal string `json:"max_total"`
		State    string `json:"state"`
	} `json:"market"`
	BidAccount struct {
		Currency            string `json:"currency"`
		Balance             string `json:"balance"`
		Locked              string `json:"locked"`
		AvgBuyPrice         string `json:"avg_buy_price"`
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
		UnitCurrency        string `json:"unit_currency"`
	} `json:"bid_account"`
	AskAccount struct {
		Currency            string `json:"currency"`
		Balance             string `json:"balance"`
		Locked              string `json:"locked"`
		AvgBuyPrice         string `json:"avg_buy_price"`
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
		UnitCurrency        string `json:"unit_currency"`
	} `json:"ask_account"`
}

type OrderDetail struct {
	UUID            string    `json:"uuid"`
	Side            string    `json:"side"`
	OrdType         string    `json:"ord_type"`
	Price           string    `json:"price"`
	State           string    `json:"state"`
	Market          string    `json:"market"`
	CreatedAt       time.Time `json:"created_at"`
	Volume          string    `json:"volume"`
	RemainingVolume string    `json:"remaining_volume"`
	ReservedFee     string    `json:"reserved_fee"`
	RemainingFee    string    `json:"remaining_fee"`
	PaidFee         string    `json:"paid_fee"`
	Locked          string    `json:"locked"`
	ExecutedVolume  string    `json:"executed_volume"`
	TradesCount     int       `json:"trades_count"`
	Trades          []struct {
		Market string `json:"market"`
		UUID   string `json:"uuid"`
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Funds  string `json:"funds"`
		Side   string `json:"side"`
	} `json:"trades"`
}

type Order struct {
	UUID            string    `json:"uuid"`
	Side            string    `json:"side"`
	OrdType         string    `json:"ord_type"`
	Price           string    `json:"price"`
	AvgPrice        string    `json:"avg_price"`
	State           string    `json:"state"`
	Market          string    `json:"market"`
	CreatedAt       time.Time `json:"created_at"`
	Volume          string    `json:"volume"`
	RemainingVolume string    `json:"remaining_volume"`
	ReservedFee     string    `json:"reserved_fee"`
	RemainingFee    string    `json:"remaining_fee"`
	PaidFee         string    `json:"paid_fee"`
	Locked          string    `json:"locked"`
	ExecutedVolume  string    `json:"executed_volume"`
	TradesCount     int       `json:"trades_count"`
}

type Orders []Order

type Market struct {
	Market        string `json:"market"`
	KoreanName    string `json:"korean_name"`
	EnglishName   string `json:"english_name"`
	MarketWarning string `json:"market_warning"`
}

type ChanceParam struct {
	Market string `url:"market"`
}

type OrderParam struct {
	Uuid       string `url:"uuid,omitempty"`
	Identifier string `url:"identifier,omitempty"`
}

type OrdersParam struct {
	Market      string   `url:"market,omitempty"`
	State       string   `url:"state,omitempty"`
	States      []string `url:"states,omitempty"`
	Uuids       []string `url:"uuids,omitempty"`
	Identifiers []string `url:"identifiers,omitempty"`
	Page        int32    `url:"page,omitempty"`
	Limit       int32    `url:"limit,omitempty"`
	OrderBy     string   `url:"order_by,omitempty"`
}

type PostOrdersParam struct {
	Market     string `url:"market"`
	Side       string `url:"side"`
	Volume     string `url:"volume,omitempty"`
	Price      string `url:"price,omitempty"`
	OrdType    string `url:"ord_type"`
	Identifier string `url:"identifier,omitempty"`
}

type AllMarketParam struct {
	IsDetails bool `url:"isDetails,omitempty"`
}

type ErrorMessage struct {
	Error struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	} `json:"error"`
}