# Shared Response Types

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared#APIError">APIError</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared#BaseResponse">BaseResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared#ResponseMetadata">ResponseMetadata</a>

# V1

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SecurityType">SecurityType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SecurityType">SecurityType</a>

Methods:

- <code title="get /v1/ws">client.V1.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1Service.WebsocketHandler">WebsocketHandler</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Accounts

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RiskSettingsParam">RiskSettingsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountList">AccountList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountSettings">AccountSettings</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountStatus">AccountStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountSubtype">AccountSubtype</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountType">AccountType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RiskSettings">RiskSettings</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountGetAccountByIDResponse">V1AccountGetAccountByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountGetAccountsResponse">V1AccountGetAccountsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPatchAccountByIDResponse">V1AccountPatchAccountByIDResponse</a>

Methods:

- <code title="get /v1/accounts/{account_id}">client.V1.Accounts.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountService.GetAccountByID">GetAccountByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountGetAccountByIDResponse">V1AccountGetAccountByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts">client.V1.Accounts.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountService.GetAccounts">GetAccounts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountGetAccountsParams">V1AccountGetAccountsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountGetAccountsResponse">V1AccountGetAccountsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/accounts/{account_id}">client.V1.Accounts.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountService.PatchAccountByID">PatchAccountByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPatchAccountByIDParams">V1AccountPatchAccountByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPatchAccountByIDResponse">V1AccountPatchAccountByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Balances

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountBalances">AccountBalances</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountBalancesSod">AccountBalancesSod</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarginDetails">MarginDetails</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarginDetailsUsage">MarginDetailsUsage</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarginTopContributor">MarginTopContributor</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarginType">MarginType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountBalanceGetAccountBalancesResponse">V1AccountBalanceGetAccountBalancesResponse</a>

Methods:

- <code title="get /v1/accounts/{account_id}/balances">client.V1.Accounts.Balances.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountBalanceService.GetAccountBalances">GetAccountBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountBalanceGetAccountBalancesParams">V1AccountBalanceGetAccountBalancesParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountBalanceGetAccountBalancesResponse">V1AccountBalanceGetAccountBalancesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewOrderRequestParam">NewOrderRequestParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PositionEffect">PositionEffect</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RequestOrderType">RequestOrderType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RequestTimeInForce">RequestTimeInForce</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Side">Side</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TrailingOffsetType">TrailingOffsetType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CancelOrderRequest">CancelOrderRequest</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewOrderRequest">NewOrderRequest</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderList">OrderList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderStatus">OrderStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderType">OrderType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PositionEffect">PositionEffect</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#QueueState">QueueState</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RequestOrderType">RequestOrderType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RequestTimeInForce">RequestTimeInForce</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Side">Side</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TimeInForce">TimeInForce</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TrailingOffsetType">TrailingOffsetType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderCancelAllOpenOrdersResponse">V1AccountOrderCancelAllOpenOrdersResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderCancelOpenOrderResponse">V1AccountOrderCancelOpenOrderResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderGetOrderByIDResponse">V1AccountOrderGetOrderByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderGetOrdersResponse">V1AccountOrderGetOrdersResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderReplaceOrderResponse">V1AccountOrderReplaceOrderResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderSubmitOrdersResponse">V1AccountOrderSubmitOrdersResponse</a>

Methods:

- <code title="delete /v1/accounts/{account_id}/orders">client.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderService.CancelAllOpenOrders">CancelAllOpenOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderCancelAllOpenOrdersParams">V1AccountOrderCancelAllOpenOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderCancelAllOpenOrdersResponse">V1AccountOrderCancelAllOpenOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/accounts/{account_id}/orders/{order_id}">client.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderService.CancelOpenOrder">CancelOpenOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderCancelOpenOrderParams">V1AccountOrderCancelOpenOrderParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderCancelOpenOrderResponse">V1AccountOrderCancelOpenOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/{account_id}/orders/{order_id}">client.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderService.GetOrderByID">GetOrderByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderGetOrderByIDParams">V1AccountOrderGetOrderByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderGetOrderByIDResponse">V1AccountOrderGetOrderByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/{account_id}/orders">client.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderService.GetOrders">GetOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderGetOrdersParams">V1AccountOrderGetOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderGetOrdersResponse">V1AccountOrderGetOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/accounts/{account_id}/orders/{order_id}">client.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderService.ReplaceOrder">ReplaceOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderReplaceOrderParams">V1AccountOrderReplaceOrderParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderReplaceOrderResponse">V1AccountOrderReplaceOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/accounts/{account_id}/orders">client.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderService.SubmitOrders">SubmitOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderSubmitOrdersParams">V1AccountOrderSubmitOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountOrderSubmitOrdersResponse">V1AccountOrderSubmitOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### PortfolioHistory

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PortfolioHistoryResponse">PortfolioHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PortfolioHistorySegment">PortfolioHistorySegment</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPortfolioHistoryGetPortfolioHistoryResponse">V1AccountPortfolioHistoryGetPortfolioHistoryResponse</a>

Methods:

- <code title="get /v1/accounts/{account_id}/portfolio-history">client.V1.Accounts.PortfolioHistory.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPortfolioHistoryService.GetPortfolioHistory">GetPortfolioHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPortfolioHistoryGetPortfolioHistoryParams">V1AccountPortfolioHistoryGetPortfolioHistoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPortfolioHistoryGetPortfolioHistoryResponse">V1AccountPortfolioHistoryGetPortfolioHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Position">Position</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PositionList">PositionList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PositionType">PositionType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionClosePositionResponse">V1AccountPositionClosePositionResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionClosePositionsResponse">V1AccountPositionClosePositionsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionGetPositionsResponse">V1AccountPositionGetPositionsResponse</a>

Methods:

- <code title="delete /v1/accounts/{account_id}/positions/{instrument_id}">client.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionService.ClosePosition">ClosePosition</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, instrumentID <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>, params <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionClosePositionParams">V1AccountPositionClosePositionParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionClosePositionResponse">V1AccountPositionClosePositionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/accounts/{account_id}/positions">client.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionService.ClosePositions">ClosePositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionClosePositionsParams">V1AccountPositionClosePositionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionClosePositionsResponse">V1AccountPositionClosePositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/{account_id}/positions">client.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionService.GetPositions">GetPositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionGetPositionsParams">V1AccountPositionGetPositionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1AccountPositionGetPositionsResponse">V1AccountPositionGetPositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Calendars

### MarketHours

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketType">MarketType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DayType">DayType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketHoursDetail">MarketHoursDetail</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketHoursDetailList">MarketHoursDetailList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketSessionType">MarketSessionType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketStatus">MarketStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketType">MarketType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SessionSchedule">SessionSchedule</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TradingSessions">TradingSessions</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1CalendarMarketHourGetMarketHoursCalendarResponse">V1CalendarMarketHourGetMarketHoursCalendarResponse</a>

Methods:

- <code title="get /v1/calendars/market-hours">client.V1.Calendars.MarketHours.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1CalendarMarketHourService.GetMarketHoursCalendar">GetMarketHoursCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1CalendarMarketHourGetMarketHoursCalendarParams">V1CalendarMarketHourGetMarketHoursCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1CalendarMarketHourGetMarketHoursCalendarResponse">V1CalendarMarketHourGetMarketHoursCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Clock

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ClockDetail">ClockDetail</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ClockGetClockResponse">V1ClockGetClockResponse</a>

Methods:

- <code title="get /v1/clock">client.V1.Clock.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ClockService.GetClock">GetClock</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ClockGetClockResponse">V1ClockGetClockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Instruments

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContractType">ContractType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AnalystRating">AnalystRating</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContractType">ContractType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ExerciseStyle">ExerciseStyle</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FiscalPeriodType">FiscalPeriodType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Instrument">Instrument</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentCore">InstrumentCore</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentCoreList">InstrumentCoreList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentEarnings">InstrumentEarnings</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ListingType">ListingType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OptionsContract">OptionsContract</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OptionsContractList">OptionsContractList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentGetInstrumentByIDResponse">V1InstrumentGetInstrumentByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentGetInstrumentsResponse">V1InstrumentGetInstrumentsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentSearchInstrumentsResponse">V1InstrumentSearchInstrumentsResponse</a>

Methods:

- <code title="get /v1/instruments/{instrument_id}">client.V1.Instruments.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentService.GetInstrumentByID">GetInstrumentByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, instrumentID <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentGetInstrumentByIDParams">V1InstrumentGetInstrumentByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentGetInstrumentByIDResponse">V1InstrumentGetInstrumentByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/instruments">client.V1.Instruments.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentService.GetInstruments">GetInstruments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentGetInstrumentsParams">V1InstrumentGetInstrumentsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentGetInstrumentsResponse">V1InstrumentGetInstrumentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/instruments/search">client.V1.Instruments.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentService.SearchInstruments">SearchInstruments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentSearchInstrumentsParams">V1InstrumentSearchInstrumentsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentSearchInstrumentsResponse">V1InstrumentSearchInstrumentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### AnalystReporting

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AnalystDistribution">AnalystDistribution</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentAnalystConsensus">InstrumentAnalystConsensus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PriceTarget">PriceTarget</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse">V1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse</a>

Methods:

- <code title="get /v1/instruments/{instrument_id}/analyst-reporting">client.V1.Instruments.AnalystReporting.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentAnalystReportingService.GetInstrumentAnalystConsensus">GetInstrumentAnalystConsensus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, instrumentID <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams">V1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse">V1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### BalanceSheets

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentBalanceSheetStatement">InstrumentBalanceSheetStatement</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentBalanceSheetStatementList">InstrumentBalanceSheetStatementList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsResponse">V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsResponse</a>

Methods:

- <code title="get /v1/instruments/{instrument_id}/balance-sheets">client.V1.Instruments.BalanceSheets.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentBalanceSheetService.GetInstrumentBalanceSheetStatements">GetInstrumentBalanceSheetStatements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, instrumentID <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsParams">V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsResponse">V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### CashFlowStatements

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentCashFlowStatement">InstrumentCashFlowStatement</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentCashFlowStatementList">InstrumentCashFlowStatementList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsResponse">V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsResponse</a>

Methods:

- <code title="get /v1/instruments/{instrument_id}/cash-flow-statements">client.V1.Instruments.CashFlowStatements.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentCashFlowStatementService.GetInstrumentCashFlowStatements">GetInstrumentCashFlowStatements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, instrumentID <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsParams">V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsResponse">V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Events

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AllEventsEventType">AllEventsEventType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AllEventsEventType">AllEventsEventType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentAllEventsData">InstrumentAllEventsData</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentDividendEvent">InstrumentDividendEvent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentEventEnvelope">InstrumentEventEnvelope</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentEventIpoItem">InstrumentEventIpoItem</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentEventsByDate">InstrumentEventsByDate</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentEventsData">InstrumentEventsData</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentSplitEvent">InstrumentSplitEvent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentEventGetAllInstrumentEventsResponse">V1InstrumentEventGetAllInstrumentEventsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentEventGetInstrumentEventsResponse">V1InstrumentEventGetInstrumentEventsResponse</a>

Methods:

- <code title="get /v1/instruments/events">client.V1.Instruments.Events.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentEventService.GetAllInstrumentEvents">GetAllInstrumentEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentEventGetAllInstrumentEventsParams">V1InstrumentEventGetAllInstrumentEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentEventGetAllInstrumentEventsResponse">V1InstrumentEventGetAllInstrumentEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/instruments/{instrument_id}/events">client.V1.Instruments.Events.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentEventService.GetInstrumentEvents">GetInstrumentEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, instrumentID <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentEventGetInstrumentEventsParams">V1InstrumentEventGetInstrumentEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentEventGetInstrumentEventsResponse">V1InstrumentEventGetInstrumentEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Fundamentals

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentFundamentals">InstrumentFundamentals</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentFundamentalGetInstrumentFundamentalsResponse">V1InstrumentFundamentalGetInstrumentFundamentalsResponse</a>

Methods:

- <code title="get /v1/instruments/{instrument_id}/fundamentals">client.V1.Instruments.Fundamentals.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentFundamentalService.GetInstrumentFundamentals">GetInstrumentFundamentals</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, instrumentID <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentFundamentalGetInstrumentFundamentalsResponse">V1InstrumentFundamentalGetInstrumentFundamentalsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### IncomeStatements

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIncomeStatement">InstrumentIncomeStatement</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIncomeStatementList">InstrumentIncomeStatementList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentIncomeStatementGetInstrumentIncomeStatementsResponse">V1InstrumentIncomeStatementGetInstrumentIncomeStatementsResponse</a>

Methods:

- <code title="get /v1/instruments/{instrument_id}/income-statements">client.V1.Instruments.IncomeStatements.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentIncomeStatementService.GetInstrumentIncomeStatements">GetInstrumentIncomeStatements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, instrumentID <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentIDOrSymbol">InstrumentIDOrSymbol</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentIncomeStatementGetInstrumentIncomeStatementsParams">V1InstrumentIncomeStatementGetInstrumentIncomeStatementsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentIncomeStatementGetInstrumentIncomeStatementsResponse">V1InstrumentIncomeStatementGetInstrumentIncomeStatementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Options

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentOptionGetOptionContractsResponse">V1InstrumentOptionGetOptionContractsResponse</a>

Methods:

- <code title="get /v1/instruments/options/contracts">client.V1.Instruments.Options.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentOptionService.GetOptionContracts">GetOptionContracts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentOptionGetOptionContractsParams">V1InstrumentOptionGetOptionContractsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1InstrumentOptionGetOptionContractsResponse">V1InstrumentOptionGetOptionContractsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## MarketData

### DailySummary

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DailySummary">DailySummary</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DailySummaryList">DailySummaryList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1MarketDataDailySummaryGetDailySummariesResponse">V1MarketDataDailySummaryGetDailySummariesResponse</a>

Methods:

- <code title="get /v1/market-data/daily-summary">client.V1.MarketData.DailySummary.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1MarketDataDailySummaryService.GetDailySummaries">GetDailySummaries</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1MarketDataDailySummaryGetDailySummariesParams">V1MarketDataDailySummaryGetDailySummariesParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1MarketDataDailySummaryGetDailySummariesResponse">V1MarketDataDailySummaryGetDailySummariesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Snapshot

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketDataSnapshot">MarketDataSnapshot</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketDataSnapshotList">MarketDataSnapshotList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SnapshotLastTrade">SnapshotLastTrade</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SnapshotQuote">SnapshotQuote</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SnapshotSession">SnapshotSession</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1MarketDataSnapshotGetSnapshotsResponse">V1MarketDataSnapshotGetSnapshotsResponse</a>

Methods:

- <code title="get /v1/market-data/snapshot">client.V1.MarketData.Snapshot.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1MarketDataSnapshotService.GetSnapshots">GetSnapshots</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1MarketDataSnapshotGetSnapshotsParams">V1MarketDataSnapshotGetSnapshotsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1MarketDataSnapshotGetSnapshotsResponse">V1MarketDataSnapshotGetSnapshotsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## News

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewsInstrument">NewsInstrument</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewsItem">NewsItem</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewsItemList">NewsItemList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewsType">NewsType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1NewsGetNewsResponse">V1NewsGetNewsResponse</a>

Methods:

- <code title="get /v1/news">client.V1.News.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1NewsService.GetNews">GetNews</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1NewsGetNewsParams">V1NewsGetNewsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1NewsGetNewsResponse">V1NewsGetNewsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OmniAI

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EntitlementCode">EntitlementCode</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActionButton">ActionButton</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CancelResponsePayload">CancelResponsePayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ChartPayload">ChartPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ChartPoint">ChartPoint</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ChartSeries">ChartSeries</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartChartPayload">ContentPartChartPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartCustomPayload">ContentPartCustomPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartStructuredActionPayload">ContentPartStructuredActionPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartSuggestedActionsPayload">ContentPartSuggestedActionsPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartTextPayload">ContentPartTextPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartThinkingPayload">ContentPartThinkingPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CreateFeedbackResponse">CreateFeedbackResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CreateMessageResponse">CreateMessageResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CreateThreadResponse">CreateThreadResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DataChart">DataChart</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EntitlementAgreementKey">EntitlementAgreementKey</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EntitlementCode">EntitlementCode</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ErrorStatus">ErrorStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageContent">MessageContent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageContentPartUnion">MessageContentPartUnion</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageList">MessageList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageOutcome">MessageOutcome</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageRole">MessageRole</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OpenChartAction">OpenChartAction</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OpenEntitlementConsentAction">OpenEntitlementConsentAction</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OpenScreenerAction">OpenScreenerAction</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PrefillOrderActionUnion">PrefillOrderActionUnion</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PromptButtonAction">PromptButtonAction</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Response">Response</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ResponseContent">ResponseContent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ResponseContentPartUnion">ResponseContentPartUnion</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ResponseStatus">ResponseStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#StructuredActionUnion">StructuredActionUnion</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#StructuredActionButtonAction">StructuredActionButtonAction</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SuggestedActionsPayload">SuggestedActionsPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SymbolChart">SymbolChart</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Thread">Thread</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ThreadList">ThreadList</a>

### EntitlementAgreements

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EntitlementAgreementResource">EntitlementAgreementResource</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EntitlementAgreementResourceList">EntitlementAgreementResourceList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementAgreementGetEntitlementAgreementsResponse">V1OmniAIEntitlementAgreementGetEntitlementAgreementsResponse</a>

Methods:

- <code title="get /v1/omni-ai/entitlement-agreements">client.V1.OmniAI.EntitlementAgreements.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementAgreementService.GetEntitlementAgreements">GetEntitlementAgreements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementAgreementGetEntitlementAgreementsResponse">V1OmniAIEntitlementAgreementGetEntitlementAgreementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Entitlements

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DeleteEntitlementResponse">DeleteEntitlementResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EntitlementResource">EntitlementResource</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EntitlementResourceList">EntitlementResourceList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementNewEntitlementsResponse">V1OmniAIEntitlementNewEntitlementsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementDeleteEntitlementResponse">V1OmniAIEntitlementDeleteEntitlementResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementGetEntitlementsResponse">V1OmniAIEntitlementGetEntitlementsResponse</a>

Methods:

- <code title="post /v1/omni-ai/entitlements">client.V1.OmniAI.Entitlements.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementService.NewEntitlements">NewEntitlements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementNewEntitlementsParams">V1OmniAIEntitlementNewEntitlementsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementNewEntitlementsResponse">V1OmniAIEntitlementNewEntitlementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/omni-ai/entitlements/{entitlement_id}">client.V1.OmniAI.Entitlements.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementService.DeleteEntitlement">DeleteEntitlement</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entitlementID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementDeleteEntitlementResponse">V1OmniAIEntitlementDeleteEntitlementResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/omni-ai/entitlements">client.V1.OmniAI.Entitlements.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementService.GetEntitlements">GetEntitlements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementGetEntitlementsParams">V1OmniAIEntitlementGetEntitlementsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIEntitlementGetEntitlementsResponse">V1OmniAIEntitlementGetEntitlementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIMessageGetMessageByIDResponse">V1OmniAIMessageGetMessageByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIMessageSubmitFeedbackResponse">V1OmniAIMessageSubmitFeedbackResponse</a>

Methods:

- <code title="get /v1/omni-ai/messages/{message_id}">client.V1.OmniAI.Messages.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIMessageService.GetMessageByID">GetMessageByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIMessageGetMessageByIDParams">V1OmniAIMessageGetMessageByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIMessageGetMessageByIDResponse">V1OmniAIMessageGetMessageByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/omni-ai/messages/{message_id}/feedback">client.V1.OmniAI.Messages.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIMessageService.SubmitFeedback">SubmitFeedback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIMessageSubmitFeedbackParams">V1OmniAIMessageSubmitFeedbackParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIMessageSubmitFeedbackResponse">V1OmniAIMessageSubmitFeedbackResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Responses

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIResponseCancelResponseResponse">V1OmniAIResponseCancelResponseResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIResponseGetResponseByIDResponse">V1OmniAIResponseGetResponseByIDResponse</a>

Methods:

- <code title="delete /v1/omni-ai/responses/{response_id}">client.V1.OmniAI.Responses.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIResponseService.CancelResponse">CancelResponse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIResponseCancelResponseParams">V1OmniAIResponseCancelResponseParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIResponseCancelResponseResponse">V1OmniAIResponseCancelResponseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/omni-ai/responses/{response_id}">client.V1.OmniAI.Responses.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIResponseService.GetResponseByID">GetResponseByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIResponseGetResponseByIDParams">V1OmniAIResponseGetResponseByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIResponseGetResponseByIDResponse">V1OmniAIResponseGetResponseByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Threads

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadNewThreadResponse">V1OmniAIThreadNewThreadResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadByIDResponse">V1OmniAIThreadGetThreadByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadResponseResponse">V1OmniAIThreadGetThreadResponseResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadsResponse">V1OmniAIThreadGetThreadsResponse</a>

Methods:

- <code title="post /v1/omni-ai/threads">client.V1.OmniAI.Threads.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadService.NewThread">NewThread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadNewThreadParams">V1OmniAIThreadNewThreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadNewThreadResponse">V1OmniAIThreadNewThreadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/omni-ai/threads/{thread_id}">client.V1.OmniAI.Threads.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadService.GetThreadByID">GetThreadByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadByIDParams">V1OmniAIThreadGetThreadByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadByIDResponse">V1OmniAIThreadGetThreadByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/omni-ai/threads/{thread_id}/response">client.V1.OmniAI.Threads.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadService.GetThreadResponse">GetThreadResponse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadResponseParams">V1OmniAIThreadGetThreadResponseParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadResponseResponse">V1OmniAIThreadGetThreadResponseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/omni-ai/threads">client.V1.OmniAI.Threads.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadService.GetThreads">GetThreads</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadsParams">V1OmniAIThreadGetThreadsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadGetThreadsResponse">V1OmniAIThreadGetThreadsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadMessageNewMessageResponse">V1OmniAIThreadMessageNewMessageResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadMessageGetMessagesResponse">V1OmniAIThreadMessageGetMessagesResponse</a>

Methods:

- <code title="post /v1/omni-ai/threads/{thread_id}/messages">client.V1.OmniAI.Threads.Messages.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadMessageService.NewMessage">NewMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadMessageNewMessageParams">V1OmniAIThreadMessageNewMessageParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadMessageNewMessageResponse">V1OmniAIThreadMessageNewMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/omni-ai/threads/{thread_id}/messages">client.V1.OmniAI.Threads.Messages.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadMessageService.GetMessages">GetMessages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadMessageGetMessagesParams">V1OmniAIThreadMessageGetMessagesParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1OmniAIThreadMessageGetMessagesResponse">V1OmniAIThreadMessageGetMessagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SavedScreeners

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerEntry">ScreenerEntry</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerEntryList">ScreenerEntryList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerNewScreenerResponse">V1SavedScreenerNewScreenerResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerGetScreenerByIDResponse">V1SavedScreenerGetScreenerByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerGetScreenersResponse">V1SavedScreenerGetScreenersResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerReplaceScreenerResponse">V1SavedScreenerReplaceScreenerResponse</a>

Methods:

- <code title="post /v1/saved-screeners">client.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerService.NewScreener">NewScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerNewScreenerParams">V1SavedScreenerNewScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerNewScreenerResponse">V1SavedScreenerNewScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/saved-screeners/{screener_id}">client.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerService.DeleteScreener">DeleteScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, screenerID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/saved-screeners/{screener_id}">client.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerService.GetScreenerByID">GetScreenerByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, screenerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerGetScreenerByIDResponse">V1SavedScreenerGetScreenerByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/saved-screeners">client.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerService.GetScreeners">GetScreeners</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerGetScreenersResponse">V1SavedScreenerGetScreenersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/saved-screeners/{screener_id}">client.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerService.ReplaceScreener">ReplaceScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, screenerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerReplaceScreenerParams">V1SavedScreenerReplaceScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1SavedScreenerReplaceScreenerResponse">V1SavedScreenerReplaceScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Screener

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldLookback">FieldLookback</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldPeriod">FieldPeriod</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldRefParam">FieldRefParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldType">FieldType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FilterOpSpecParam">FilterOpSpecParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FilterOperator">FilterOperator</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FilterValueParam">FilterValueParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ModifierParam">ModifierParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ModifierOp">ModifierOp</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OperatorArg">OperatorArg</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SearchFilterParam">SearchFilterParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#VariableParam">VariableParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldLookback">FieldLookback</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldPeriod">FieldPeriod</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldRef">FieldRef</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldType">FieldType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FilterOpSpec">FilterOpSpec</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FilterOperator">FilterOperator</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FilterValue">FilterValue</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Modifier">Modifier</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ModifierOp">ModifierOp</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OperatorArg">OperatorArg</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerColumn">ScreenerColumn</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerFilter">ScreenerFilter</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerItem">ScreenerItem</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerItemList">ScreenerItemList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerRow">ScreenerRow</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerRowList">ScreenerRowList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SearchFilter">SearchFilter</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Variable">Variable</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ScreenerGetScreenerResponse">V1ScreenerGetScreenerResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ScreenerSearchScreenerResponse">V1ScreenerSearchScreenerResponse</a>

Methods:

- <code title="get /v1/screener">client.V1.Screener.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ScreenerService.GetScreener">GetScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ScreenerGetScreenerParams">V1ScreenerGetScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ScreenerGetScreenerResponse">V1ScreenerGetScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/screener">client.V1.Screener.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ScreenerService.SearchScreener">SearchScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ScreenerSearchScreenerParams">V1ScreenerSearchScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1ScreenerSearchScreenerResponse">V1ScreenerSearchScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Version

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Version">Version</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1VersionGetVersionResponse">V1VersionGetVersionResponse</a>

Methods:

- <code title="get /v1/version">client.V1.Version.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1VersionService.GetVersion">GetVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1VersionGetVersionResponse">V1VersionGetVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Watchlists

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#WatchlistDetail">WatchlistDetail</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#WatchlistEntry">WatchlistEntry</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#WatchlistEntryList">WatchlistEntryList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#WatchlistItemEntry">WatchlistItemEntry</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistNewWatchlistResponse">V1WatchlistNewWatchlistResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistDeleteWatchlistResponse">V1WatchlistDeleteWatchlistResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistGetWatchlistByIDResponse">V1WatchlistGetWatchlistByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistGetWatchlistsResponse">V1WatchlistGetWatchlistsResponse</a>

Methods:

- <code title="post /v1/watchlists">client.V1.Watchlists.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistService.NewWatchlist">NewWatchlist</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistNewWatchlistParams">V1WatchlistNewWatchlistParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistNewWatchlistResponse">V1WatchlistNewWatchlistResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/watchlists/{watchlist_id}">client.V1.Watchlists.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistService.DeleteWatchlist">DeleteWatchlist</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistDeleteWatchlistResponse">V1WatchlistDeleteWatchlistResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/watchlists/{watchlist_id}">client.V1.Watchlists.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistService.GetWatchlistByID">GetWatchlistByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistGetWatchlistByIDResponse">V1WatchlistGetWatchlistByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/watchlists">client.V1.Watchlists.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistService.GetWatchlists">GetWatchlists</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistGetWatchlistsParams">V1WatchlistGetWatchlistsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistGetWatchlistsResponse">V1WatchlistGetWatchlistsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Items

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AddWatchlistItemData">AddWatchlistItemData</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistItemAddWatchlistItemResponse">V1WatchlistItemAddWatchlistItemResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistItemDeleteWatchlistItemResponse">V1WatchlistItemDeleteWatchlistItemResponse</a>

Methods:

- <code title="post /v1/watchlists/{watchlist_id}/items">client.V1.Watchlists.Items.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistItemService.AddWatchlistItem">AddWatchlistItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistItemAddWatchlistItemParams">V1WatchlistItemAddWatchlistItemParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistItemAddWatchlistItemResponse">V1WatchlistItemAddWatchlistItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/watchlists/{watchlist_id}/items/{item_id}">client.V1.Watchlists.Items.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistItemService.DeleteWatchlistItem">DeleteWatchlistItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistItemDeleteWatchlistItemParams">V1WatchlistItemDeleteWatchlistItemParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#V1WatchlistItemDeleteWatchlistItemResponse">V1WatchlistItemDeleteWatchlistItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
