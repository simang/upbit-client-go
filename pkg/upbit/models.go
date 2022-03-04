package upbit

type Account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

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
			MinTotal  string      `json:"min_total"`
		} `json:"bid"`
		Ask struct {
			Currency  string      `json:"currency"`
			PriceUnit interface{} `json:"price_unit"`
			MinTotal  string      `json:"min_total"`
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
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	Price           string `json:"price"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	Volume          string `json:"volume"`
	RemainingVolume string `json:"remaining_volume"`
	ReservedFee     string `json:"reserved_fee"`
	RemainingFee    string `json:"remaining_fee"`
	PaidFee         string `json:"paid_fee"`
	Locked          string `json:"locked"`
	ExecutedVolume  string `json:"executed_volume"`
	TradesCount     int    `json:"trades_count"`
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
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	Price           string `json:"price"`
	AvgPrice        string `json:"avg_price"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	Volume          string `json:"volume"`
	RemainingVolume string `json:"remaining_volume"`
	ReservedFee     string `json:"reserved_fee"`
	RemainingFee    string `json:"remaining_fee"`
	PaidFee         string `json:"paid_fee"`
	Locked          string `json:"locked"`
	ExecutedVolume  string `json:"executed_volume"`
	TradesCount     int    `json:"trades_count"`
}

type Market struct {
	Market        string `json:"market"`
	KoreanName    string `json:"korean_name"`
	EnglishName   string `json:"english_name"`
	MarketWarning string `json:"market_warning"`
}

type Ticker struct {
	Market             string  `json:"market"`
	TradeDate          string  `json:"trade_date"`
	TradeTime          string  `json:"trade_time"`
	TradeDateKst       string  `json:"trade_date_kst"`
	TradeTimeKst       string  `json:"trade_time_kst"`
	TradeTimestamp     int64   `json:"trade_timestamp"`
	OpeningPrice       float64 `json:"opening_price"`
	HighPrice          float64 `json:"high_price"`
	LowPrice           float64 `json:"low_price"`
	TradePrice         float64 `json:"trade_price"`
	PrevClosingPrice   float64 `json:"prev_closing_price"`
	Change             string  `json:"change"`
	ChangePrice        float64 `json:"change_price"`
	ChangeRate         float64 `json:"change_rate"`
	SignedChangePrice  float64 `json:"signed_change_price"`
	SignedChangeRate   float64 `json:"signed_change_rate"`
	TradeVolume        float64 `json:"trade_volume"`
	AccTradePrice      float64 `json:"acc_trade_price"`
	AccTradePrice24H   float64 `json:"acc_trade_price_24h"`
	AccTradeVolume     float64 `json:"acc_trade_volume"`
	AccTradeVolume24H  float64 `json:"acc_trade_volume_24h"`
	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"`
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	Timestamp          int64   `json:"timestamp"`
}

type Candle struct {
	Market               string  `json:"market"`
	CandleDateTimeUtc    string  `json:"candle_date_time_utc"`
	CandleDateTimeKst    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TradePrice           float64 `json:"trade_price"`
	Timestamp            int64   `json:"timestamp"`
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
	Unit                 int     `json:"unit"`
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
	Identifier string `url:"identifier,omitempty" json:"identifier,omitempty"`
	Market     string `url:"market" json:"market"`
	OrdType    string `url:"ord_type" json:"ord_type"`
	Price      string `url:"price,omitempty" json:"price,omitempty"`
	Side       string `url:"side" json:"side"`
	Volume     string `url:"volume,omitempty" json:"volume,omitempty"`
}

type AllMarketParam struct {
	IsDetails bool `url:"isDetails,omitempty"`
}

type TickerParam struct {
	Markets string `url:"markets"`
}

type CandleParam struct {
	Market string `url:"market"`
	Count  int32  `url:"count"`
	To     string `url:"to"` // yyyy-MM-dd'T'HH:mm:ss'Z' or yyyy-MM-dd HH:mm:ss
}

type ErrorMessage struct {
	Error struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	} `json:"error"`
}
