package vfs

type GetPortfolioResponse struct {
	StatusCode int     `json:"statusCode"`
	ErrorCode  *string `json:"errorCode"`
	Message    *string `json:"message"`
	ErrorField *string `json:"errorField"`
	Data       []struct {
		TransDate            int         `json:"transDate"`
		SubAccoCd            int         `json:"subAccoCd"`
		SubAccountNo         string      `json:"subAccountNo"`
		CustName             string      `json:"custName"`
		SecCd                string      `json:"secCd"`
		RemainQty            int         `json:"remainQty"`
		SellAvailQty         interface{} `json:"sellAvailQty"`
		RightQty             int         `json:"rightQty"`
		RightAmt             int         `json:"rightAmt"`
		InvestmentAmt        int         `json:"investmentAmt"`
		AvgPrice             float64     `json:"avgPrice"`
		CurrentPrice         float64     `json:"currentPrice"`
		InvestmentAmtInday   float64     `json:"investmentAmtInday"`
		ProfitAmtInday       interface{} `json:"profitAmtInday"`
		ProfitAmtAcm         float64     `json:"profitAmtAcm"`
		FloorPrice           float64     `json:"floorPrice"`
		BasicPrice           float64     `json:"basicPrice"`
		CeilingPrice         float64     `json:"ceilingPrice"`
		LastPrice            interface{} `json:"lastPrice"`
		ChangePoint          float64     `json:"changePoint"`
		ChangePercent        float64     `json:"changePercent"`
		InvestmentValueInday interface{} `json:"investmentValueInday"`
		ProfitValueInday     interface{} `json:"profitValueInday"`
		ProfitValueAcm       interface{} `json:"profitValueAcm"`
		ProfitPercentAcm     float64     `json:"profitPercentAcm"`
		SecClassId           int         `json:"secClassId"`
		SecClassName         string      `json:"secClassName"`
		MarginInterest       int         `json:"marginInterest"`
		AvailQty             int         `json:"availQty"`
	} `json:"data"`
}

type GetTodayTransactionResponse struct {
	StatusCode int         `json:"statusCode"`
	ErrorCode  interface{} `json:"errorCode"`
	Message    interface{} `json:"message"`
	ErrorField interface{} `json:"errorField"`
	Data       []struct {
		AlloDate           int         `json:"alloDate"`
		RefNo              int         `json:"refNo"`
		DelCd              int         `json:"delCd"`
		SubAccoCd          int         `json:"subAccoCd"`
		SecCd              string      `json:"secCd"`
		Qty                int         `json:"qty"`
		TradeType          int         `json:"tradeType"`
		Price              float64     `json:"price"`
		Amount             int         `json:"amount"`
		GroupCategory      string      `json:"groupCategory"`
		TaskCd             string      `json:"taskCd"`
		RemainQtyAcm       int         `json:"remainQtyAcm"`
		InvestmentValueAcm interface{} `json:"investmentValueAcm"`
		ProfitValue        interface{} `json:"profitValue"`
		ProfitAmt          float64     `json:"profitAmt"`
		FeeAmt             int         `json:"feeAmt"`
		TaxAmt             int         `json:"taxAmt"`
		TotalfeeAmt        interface{} `json:"totalfeeAmt"`
		TotaltaxAmt        interface{} `json:"totaltaxAmt"`
		OrderNo            string      `json:"orderNo"`
		ConfirmNo          interface{} `json:"confirmNo"`
		MarketCd           int         `json:"marketCd"`
		RightexecDate      interface{} `json:"rightexecDate"`
		RightexecNo        interface{} `json:"rightexecNo"`
		TransactionCd      interface{} `json:"transactionCd"`
		Remarks            string      `json:"remarks"`
		RegDateTime        int64       `json:"regDateTime"`
		RegUserId          string      `json:"regUserId"`
		UpdDateTime        int64       `json:"updDateTime"`
		UpdUserId          string      `json:"updUserId"`
		FromDate           interface{} `json:"fromDate"`
		ToDate             interface{} `json:"toDate"`
		SubAccNo           string      `json:"subAccNo"`
		CustType           interface{} `json:"custType"`
		CustFamilyName     string      `json:"custFamilyName"`
		CustGiveName       string      `json:"custGiveName"`
	} `json:"data"`
}
