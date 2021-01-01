package polygon

import "time"

// SymbolsMetadata is the structure that defines symbol
// metadata served through polygon's REST API.
type SymbolsMetadata struct {
	Symbols []struct {
		Symbol  string    `json:"symbol"`
		Name    string    `json:"name"`
		Type    string    `json:"type"`
		Updated time.Time `json:"updated"`
		IsOTC   bool      `json:"isOTC"`
		URL     string    `json:"url"`
	} `json:"symbols"`
}

// HistoricTrades is the structure that defines trade
// data served through polygon's REST API.
type HistoricTrades struct {
	Day string `json:"day"`
	Map struct {
		C1 string `json:"c1"`
		C2 string `json:"c2"`
		C3 string `json:"c3"`
		C4 string `json:"c4"`
		E  string `json:"e"`
		P  string `json:"p"`
		S  string `json:"s"`
		T  string `json:"t"`
	} `json:"map"`
	MsLatency int         `json:"msLatency"`
	Status    string      `json:"status"`
	Symbol    string      `json:"symbol"`
	Ticks     []TradeTick `json:"ticks"`
	Type      string      `json:"type"`
}

// TradeTick is the structure that contains the actual
// tick data included in a HistoricTrades response
type TradeTick struct {
	Timestamp  int64   `json:"t"`
	Price      float64 `json:"p"`
	Size       int     `json:"s"`
	Exchange   string  `json:"e"`
	Condition1 int     `json:"c1"`
	Condition2 int     `json:"c2"`
	Condition3 int     `json:"c3"`
	Condition4 int     `json:"c4"`
}

type MapItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// HistoricTradesV2 is the structure that defines trade
// data served through polygon's REST API.
type HistoricTradesV2 struct {
	ResultsCount int64              `json:"results_count"`
	Ticker       string             `json:"ticker"`
	Results      []TradeTickV2      `json:"results"`
	Map          map[string]MapItem `json:"map"`
}

// TradeTickV2 is the structure that contains the actual
// tick data included in a HistoricTradesV2 response
type TradeTickV2 struct {
	SIPTimestamp         *int64   `json:"t"`
	ParticipantTimestamp *int64   `json:"y"`
	TRFTimestamp         *int64   `json:"f"`
	SequenceNumber       *int     `json:"q"`
	ID                   *string  `json:"i"`
	OrigID               *string  `json:"I"`
	Exchange             *int     `json:"x"`
	TRFID                *int     `json:"r"`
	Size                 *int     `json:"s"`
	Conditions           *[]int   `json:"c"`
	Price                *float64 `json:"p"`
	Tape                 *int     `json:"z"`
	Correction           *int     `json:"e"`
}

// HistoricQuotes is the structure that defines quote
// data served through polygon's REST API.
type HistoricQuotes struct {
	Day string `json:"day"`
	Map struct {
		AE string `json:"aE"`
		AP string `json:"aP"`
		AS string `json:"aS"`
		BE string `json:"bE"`
		BP string `json:"bP"`
		BS string `json:"bS"`
		C  string `json:"c"`
		T  string `json:"t"`
	} `json:"map"`
	MsLatency int         `json:"msLatency"`
	Status    string      `json:"status"`
	Symbol    string      `json:"symbol"`
	Ticks     []QuoteTick `json:"ticks"`
	Type      string      `json:"type"`
}

// QuoteTick is the structure that contains the actual
// tick data included in a HistoricQuotes response
type QuoteTick struct {
	Timestamp   int64   `json:"t"`
	BidExchange string  `json:"bE"`
	AskExchange string  `json:"aE"`
	BidPrice    float64 `json:"bP"`
	AskPrice    float64 `json:"aP"`
	BidSize     int     `json:"bS"`
	AskSize     int     `json:"aS"`
	Condition   int     `json:"c"`
}

// HistoricQuotesV2 is the structure that defines trade
// data served through polygon's REST API.
type HistoricQuotesV2 struct {
	ResultsCount int64              `json:"results_count"`
	Ticker       string             `json:"ticker"`
	Results      []QuoteTickV2      `json:"results"`
	Map          map[string]MapItem `json:"map"`
}

// QuoteTickV2 is the structure that contains the actual
// tick data included in a HistoricQuotesV2 response
type QuoteTickV2 struct {
	SIPTimestamp         *int64   `json:"t"`
	ParticipantTimestamp *int64   `json:"y"`
	TRFTimestamp         *int64   `json:"f"`
	SequenceNumber       *int     `json:"q"`
	Indicators           *[]int   `json:"i"`
	BidExchange          *int     `json:"x"`
	AskExchange          *int     `json:"X"`
	TRFID                *int     `json:"r"`
	Size                 *int     `json:"s"`
	Conditions           *[]int   `json:"c"`
	BidPrice             *float64 `json:"p"`
	AskPrice             *float64 `json:"P"`
	BidSize              *int     `json:"s"`
	AskSize              *int     `json:"S"`
	Tape                 *int     `json:"z"`
}

// HistoricAggregates is the structure that defines
// aggregate data served through Polygon's v1 REST API.
type HistoricAggregates struct {
	Symbol        string  `json:"symbol"`
	AggregateType AggType `json:"aggType"`
	Map           struct {
		O string `json:"o"`
		C string `json:"c"`
		H string `json:"h"`
		L string `json:"l"`
		V string `json:"v"`
		D string `json:"d"`
	} `json:"map"`
	Ticks []AggTick `json:"ticks"`
}

// HistoricAggregatesV2 is the structure that defines
// aggregate data served through Polygon's v2 REST API.
type HistoricAggregatesV2 struct {
	Symbol       string    `json:"ticker"`
	Adjusted     bool      `json:"adjusted"`
	QueryCount   int       `json:"queryCount"`
	ResultsCount int       `json:"resultsCount"`
	Ticks        []AggTick `json:"results"`
}

// DailyGroupedV2 is the structure that defines
// grouped daily data served through Polygon's v2 REST API.
type DailyGroupedV2 struct {
	QueryCount   int         `json:"queryCount"`
	ResultsCount int         `json:"resultsCount"`
	Adjusted     bool        `json:"adjusted"`
	Results      []GroupTick `json:"results"`
}

// GroupTick is the structure that defines
// grouped daily data for a single symbol
// served through Polygon's v2 REST API.
type GroupTick struct {
	Symbol            string  `json:"T"`
	Volume            float64 `json:"v"`
	OpenPrice         float64 `json:"o"`
	ClosePrice        float64 `json:"c"`
	HighPrice         float64 `json:"h"`
	LowPrice          float64 `json:"l"`
	EpochMilliseconds int64   `json:"t"`
	VolWeightedAvgPrice float64 `json:"vw"`
}

type GetHistoricTradesParams struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type HistoricTicksV2Params struct {
	Timestamp      int64 `json:"timestamp"`
	TimestampLimit int64 `json:"timestamp_limit"`
	Reverse        bool  `json:"reverse"`
	Limit          int64 `json:"limit"`
}

// AggTick is the structure that contains the actual
// tick data included in a HistoricAggregates response
type AggTick struct {
	Open              float64 `json:"o"`
	High              float64 `json:"h"`
	Low               float64 `json:"l"`
	Close             float64 `json:"c"`
	Volume            float64 `json:"v"`
	EpochMilliseconds int64   `json:"t"`
	Items             int64   `json:"n"` // v2 response only
}

// AggType used in the HistoricAggregates response
type AggType string

const (
	// Minute timeframe aggregates
	Minute AggType = "minute"
	// Day timeframe aggregates
	Day AggType = "day"
)

// polygon stream

// PolygonClientMsg is the standard message sent by clients of the stream interface
type PolygonClientMsg struct {
	Action string `json:"action"`
	Params string `json:"params"`
}

type PolygonAuthMsg struct {
	Event   string `json:"ev"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// PolygonServerMsg contains the field that is present in all responses to identify their type
type PolgyonServerMsg struct {
	Event string `json:"ev"`
}

// StreamTrade is the structure that defines a trade that
// polygon transmits via websocket protocol.
type StreamTrade struct {
	Symbol     string  `json:"sym"`
	Exchange   int     `json:"x"`
	TradeID    string  `json:"i"`
	Price      float64 `json:"p"`
	Size       int64   `json:"s"`
	Timestamp  int64   `json:"t"`
	Conditions []int   `json:"c"`
}

// StreamQuote is the structure that defines a quote that
// polygon transmits via websocket protocol.
type StreamQuote struct {
	Symbol      string  `json:"sym"`
	Condition   int     `json:"c"`
	BidExchange int     `json:"bx"`
	AskExchange int     `json:"ax"`
	BidPrice    float64 `json:"bp"`
	AskPrice    float64 `json:"ap"`
	BidSize     int64   `json:"bs"`
	AskSize     int64   `json:"as"`
	Timestamp   int64   `json:"t"`
}

// StreamAggregate is the structure that defines an aggregate that
// polygon transmits via websocket protocol.
type StreamAggregate struct {
	Event             string  `json:"ev"`
	Symbol            string  `json:"sym"`
	Volume            int     `json:"v"`
	AccumulatedVolume int     `json:"av"`
	OpeningPrice      float64 `json:"op"`
	VWAP              float64 `json:"vw"`
	OpenPrice         float64 `json:"o"`
	ClosePrice        float64 `json:"c"`
	HighPrice         float64 `json:"h"`
	LowPrice          float64 `json:"l"`
	Average           float64 `json:"a"`
	StartTimestamp    int64   `json:"s"`
	EndTimestamp      int64   `json:"e"`
}

// Exchange defines the Stocks / Equities "Exchange" endpoint response
type StockExchange struct {
	Id     int64  `json:"id"`
	Type   string `json:"type"`
	Market string `json:"market"`
	Mic    string `json:"mic"`
	Name   string `json:"name"`
	Tape   string `json:"tape"`
}

// Ticker API

type Ticker struct {
	Ticker      string `json:"ticker"`
	Name        string `json:"name"`
	Market      string `json:"market"`
	Locale      string `json:"locale"`
	Currency    string `json:"currency"`
	Active      bool   `json:"active"`
	PrimaryExch string `json:"primaryExch"`
	Type        string `json:"type"`
	Updated     string `json:"updated"`
	URL         string `json:"url"`
}

type TickerList struct {
	Count          int64    `json:"count"`
	Page           int64    `json:"page"`
	ResultsPerPage int64    `jso		n:"perPage"`
	Tickers        []Ticker `jso		n:"tickers"`
}		

// Symbol Details		

type SymbolDetails struct {		
	ListDate			string		`json:"listdate"`
	Exchange			string		`json:"exchange"`
	ExchangeSymbol 		string		`json:"exchangeSymbol"`
	Bloomberg			string		`json:"bloomberg"`
	Type				string		`json:"type"`
	Symbol				string		`json:"symbol"`
	Name				string		`json:"name"`
	CIK					string		`json:"cik"`
	Figi				string		`json:"figi"`
	LEI					string		`json:"lei"`
	SIC					int64		`json:"sic"`
	Country				string		`json:"country"`
	Industry			string		`json:"industry"`
	Sector				string		`json:"sector"`
	MarketCap			int64		`json:"marketcap"`
	Employees			int64		`json:"employees"`
	Phone				string		`json:"phone"`
	CEO					string		`json:"ceo"`
	URL					string		`json:"url"`
	Description			string		`json:"description"`
	HQAddress			string		`json:"hq_address"`
	HQState				string		`json:"hq_state"`
	HQCountry			string		`json:"hq_country"`
	Similar				[]string	`json:"similar"`	
	Tags				[]string	`json:"tags"`	
	Updated				string		`json:"updated"`
	Active				bool		`json:"active"`
	HasDetail			bool

}

type FinancialsV2 struct {
	Status 		string
	Count  		int64
	Results 	[]FinancialResultsV2 
}

type FinancialResultsV2 struct {
	Symbol																string		`json:"ticker"`
	Period																string 		`json:"period"`																																				
	CalendarDate														string																	`json:"calendarDate"`			
	ReportPeriod														string																	`json:"reportPeriod"`			
	Updated																string																		`json:"updated"`		
	AccumulatedOtherComprehensiveIncome									float64										`json:"accumulatedOtherComprehensiveIncome"`									
	Assets																float64																	`json:"assets"`		
	AssetsAverage														float64															`json:"assetsAverage"`				
	AssetsCurrent														float64															`json:"assetsCurrent"`				
	AssetTurnover														float64															`json:"assetTurnover"`				
	AssetsNonCurrent													float64															`json:"assetsNonCurrent"`				
	BookValuePerShare													float64															`json:"bookValuePerShare"`				
	CapitalExpenditure													float64														`json:"capitalExpenditure"`					
	CashAndEquivalents													float64														`json:"cashAndEquivalents"`					
	CashAndEquivalentsUSD												float64													`json:"cashAndEquivalentsUSD"`						
	CostOfRevenue														float64															`json:"costOfRevenue"`				
	ConsolidatedIncome													float64														`json:"consolidatedIncome"`					
	CurrentRatio														float64																`json:"currentRatio"`			
	DebtToEquityRatio													float64															`json:"debtToEquityRatio"`				
	Debt																float64																		`json:"debt"`	
	DebtCurrent															float64																`json:"debtCurrent"`			
	DebtNonCurrent														float64															`json:"debtNonCurrent"`				
	DebtUSD																float64																	`json:"debtUSD"`		
	DeferredRevenue														float64															`json:"deferredRevenue"`				
	DepreciationAmortizationAndAccretion								float64										`json:"depreciationAmortizationAndAccretion"`									
	Deposits															float64																	`json:"deposits"`		
	DividendYield														float64															`json:"dividendYield"`				
	DividendsPerBasicCommonShare										float64												`json:"dividendsPerBasicCommonShare"`							
	EarningBeforeInterestTaxes											float64												`json:"earningBeforeInterestTaxes"`							
	EarningsBeforeInterestTaxesDepreciationAmortization					float64						`json:"earningsBeforeInterestTaxesDepreciationAmortization"`													
	EBITDAMargin														float64																`json:"EBITDAMargin"`			
	EarningsBeforeInterestTaxesDepreciationAmortizationUSD				float64					`json:"earningsBeforeInterestTaxesDepreciationAmortizationUSD"`														
	EarningBeforeInterestTaxesUSD										float64											`json:"earningBeforeInterestTaxesUSD"`								
	EarningsBeforeTax													float64														`json:"earningsBeforeTax"`					
	EarningsPerBasicShare												float64														`json:"earningsPerBasicShare"`					
	EarningsPerDilutedShare												float64													`json:"earningsPerDilutedShare"`						
	EarningsPerBasicShareUSD											float64													`json:"earningsPerBasicShareUSD"`						
	ShareholdersEquity													float64														`json:"shareholdersEquity"`					
	AverageEquity														float64															`json:"averageEquity"`				
	ShareholdersEquityUSD												float64													`json:"shareholdersEquityUSD"`						
	EnterpriseValue														float64															`json:"enterpriseValue"`				
	EnterpriseValueOverEBIT												float64													`json:"enterpriseValueOverEBIT"`						
	EnterpriseValueOverEBITDA											float64													`json:"enterpriseValueOverEBITDA"`						
	FreeCashFlow														float64																`json:"freeCashFlow"`			
	FreeCashFlowPerShare												float64														`json:"freeCashFlowPerShare"`					
	ForeignCurrencyUSDExchangeRate										float64											`json:"foreignCurrencyUSDExchangeRate"`								
	GrossProfit															float64																`json:"grossProfit"`			
	GrossMargin															float64																`json:"grossMargin"`			
	GoodwillAndIntangibleAssets											float64												`json:"goodwillAndIntangibleAssets"`							
	InterestExpense														float64															`json:"interestExpense"`				
	InvestedCapital														float64															`json:"investedCapital"`				
	InvestedCapitalAverage												float64													`json:"investedCapitalAverage"`						
	Inventory															float64																`json:"inventory"`			
	Investments															float64																`json:"investments"`			
	InvestmentsCurrent													float64														`json:"investmentsCurrent"`					
	InvestmentsNonCurrent												float64													`json:"investmentsNonCurrent"`						
	TotalLiabilities													float64															`json:"totalLiabilities"`				
	CurrentLiabilities													float64														`json:"currentLiabilities"`					
	LiabilitiesNonCurrent												float64													`json:"liabilitiesNonCurrent"`						
	MarketCapitalization												float64														`json:"marketCapitalization"`					
	NetCashFlow															float64																`json:"netCashFlow"`			
	NetCashFlowBusinessAcquisitionsDisposals							float64									`json:"netCashFlowBusinessAcquisitionsDisposals"`										
	IssuanceEquityShares												float64														`json:"issuanceEquityShares"`					
	IssuanceDebtSecurities												float64													`json:"issuanceDebtSecurities"`						
	PaymentDividendsOtherCashDistributions								float64									`json:"paymentDividendsOtherCashDistributions"`										
	NetCashFlowFromFinancing											float64													`json:"netCashFlowFromFinancing"`						
	NetCashFlowFromInvesting											float64													`json:"netCashFlowFromInvesting"`						
	NetCashFlowInvestmentAcquisitionsDisposals							float64								`json:"netCashFlowInvestmentAcquisitionsDisposals"`											
	NetCashFlowFromOperations											float64												`json:"netCashFlowFromOperations"`							
	EffectOfExchangeRateChangesOnCash									float64										`json:"effectOfExchangeRateChangesOnCash"`									
	NetIncome															float64																`json:"netIncome"`			
	NetIncomeCommonStock												float64														`json:"netIncomeCommonStock"`					
	NetIncomeCommonStockUSD												float64													`json:"netIncomeCommonStockUSD"`						
	NetLossIncomeFromDiscontinuedOperations								float64									`json:"netLossIncomeFromDiscontinuedOperations"`										
	NetIncomeToNonControllingInterests									float64										`json:"netIncomeToNonControllingInterests"`									
	ProfitMargin														float64																`json:"profitMargin"`			
	OperatingExpenses													float64														`json:"operatingExpenses"`					
	OperatingIncome														float64															`json:"operatingIncome"`				
	TradeAndNonTradePayables											float64													`json:"tradeAndNonTradePayables"`						
	PayoutRatio															float64																`json:"payoutRatio"`			
	PriceToBookValue													float64															`json:"priceToBookValue"`				
	PriceEarnings														float64																`json:"priceEarnings"`			
	PriceToEarningsRatio												float64														`json:"priceToEarningsRatio"`					
	PropertyPlantEquipmentNet											float64												`json:"propertyPlantEquipmentNet"`							
	PreferredDividendsIncomeStatementImpact								float64									`json:"preferredDividendsIncomeStatementImpact"`										
	SharePriceAdjustedClose												float64													`json:"sharePriceAdjustedClose"`						
	PriceSales															float64																`json:"priceSales"`			
	PriceToSalesRatio													float64															`json:"priceToSalesRatio"`				
	TradeAndNonTradeReceivables											float64												`json:"tradeAndNonTradeReceivables"`							
	AccumulatedRetainedEarningsDeficit									float64										`json:"accumulatedRetainedEarningsDeficit"`									
	Revenues															float64																	`json:"revenues"`		
	RevenuesUSD															float64																`json:"revenuesUSD"`			
	ResearchAndDevelopmentExpense										float64											`json:"researchAndDevelopmentExpense"`								
	ReturnOnAverageAssets												float64													`json:"returnOnAverageAssets"`						
	ReturnOnAverageEquity												float64													`json:"returnOnAverageEquity"`						
	ReturnOnInvestedCapital												float64													`json:"returnOnInvestedCapital"`						
	ReturnOnSales														float64																`json:"returnOnSales"`			
	ShareBasedCompensation												float64													`json:"shareBasedCompensation"`						
	SellingGeneralAndAdministrativeExpense								float64									`json:"sellingGeneralAndAdministrativeExpense"`										
	ShareFactor															float64																`json:"shareFactor"`			
	Shares																float64																	`json:"shares"`		
	WeightedAverageShares												float64													`json:"weightedAverageShares"`						
	WeightedAverageSharesDiluted										float64												`json:"weightedAverageSharesDiluted"`							
	SalesPerShare														float64																`json:"salesPerShare"`			
	TangibleAssetValue													float64														`json:"tangibleAssetValue"`					
	TaxAssets															float64																`json:"taxAssets"`			
	IncomeTaxExpense													float64															`json:"incomeTaxExpense"`				
	TaxLiabilities														float64															`json:"taxLiabilities"`				
	TangibleAssetsBookValuePerShare										float64											`json:"tangibleAssetsBookValuePerShare"`								
	WorkingCapital														float64															`json:"workingCapital"`
}				