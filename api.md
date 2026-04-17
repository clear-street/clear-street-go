# Shared Response Types

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared#APIError">APIError</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared#BaseResponse">BaseResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go/shared#ResponseMetadata">ResponseMetadata</a>

# Active

## V1

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#APIDecimal64">APIDecimal64</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SecurityIDSource">SecurityIDSource</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SecurityType">SecurityType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#APIDecimal64">APIDecimal64</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SecurityIDSource">SecurityIDSource</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SecurityType">SecurityType</a>

### Accounts

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RiskSettingsParam">RiskSettingsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountKind">AccountKind</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountList">AccountList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountSettings">AccountSettings</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountStatus">AccountStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountSubkind">AccountSubkind</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RiskSettings">RiskSettings</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountGetAccountByIDResponse">ActiveV1AccountGetAccountByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountGetAccountsResponse">ActiveV1AccountGetAccountsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPatchAccountByIDResponse">ActiveV1AccountPatchAccountByIDResponse</a>

Methods:

- <code title="get /active/v1/accounts/{account_id}">client.Active.V1.Accounts.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountService.GetAccountByID">GetAccountByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountGetAccountByIDResponse">ActiveV1AccountGetAccountByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts">client.Active.V1.Accounts.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountService.GetAccounts">GetAccounts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountGetAccountsParams">ActiveV1AccountGetAccountsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountGetAccountsResponse">ActiveV1AccountGetAccountsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /active/v1/accounts/{account_id}">client.Active.V1.Accounts.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountService.PatchAccountByID">PatchAccountByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPatchAccountByIDParams">ActiveV1AccountPatchAccountByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPatchAccountByIDResponse">ActiveV1AccountPatchAccountByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Balances

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountBalances">AccountBalances</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AccountBalancesSod">AccountBalancesSod</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarginDetails">MarginDetails</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarginDetailsUsage">MarginDetailsUsage</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarginTopContributor">MarginTopContributor</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarginType">MarginType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountBalanceGetAccountBalancesResponse">ActiveV1AccountBalanceGetAccountBalancesResponse</a>

Methods:

- <code title="get /active/v1/accounts/{account_id}/balances">client.Active.V1.Accounts.Balances.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountBalanceService.GetAccountBalances">GetAccountBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountBalanceGetAccountBalancesParams">ActiveV1AccountBalanceGetAccountBalancesParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountBalanceGetAccountBalancesResponse">ActiveV1AccountBalanceGetAccountBalancesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Locates

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#LocateOrderStatus">LocateOrderStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#LocateOrder">LocateOrder</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#LocateOrderList">LocateOrderList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#LocateOrderStatus">LocateOrderStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateNewLocateRequestResponse">ActiveV1AccountLocateNewLocateRequestResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateGetLocateRequestsResponse">ActiveV1AccountLocateGetLocateRequestsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateUpdateLocateRequestResponse">ActiveV1AccountLocateUpdateLocateRequestResponse</a>

Methods:

- <code title="post /active/v1/accounts/{account_id}/locates">client.Active.V1.Accounts.Locates.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateService.NewLocateRequest">NewLocateRequest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateNewLocateRequestParams">ActiveV1AccountLocateNewLocateRequestParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateNewLocateRequestResponse">ActiveV1AccountLocateNewLocateRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts/{account_id}/locates">client.Active.V1.Accounts.Locates.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateService.GetLocateRequests">GetLocateRequests</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateGetLocateRequestsParams">ActiveV1AccountLocateGetLocateRequestsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateGetLocateRequestsResponse">ActiveV1AccountLocateGetLocateRequestsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /active/v1/accounts/{account_id}/locates">client.Active.V1.Accounts.Locates.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateService.UpdateLocateRequest">UpdateLocateRequest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateUpdateLocateRequestParams">ActiveV1AccountLocateUpdateLocateRequestParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateUpdateLocateRequestResponse">ActiveV1AccountLocateUpdateLocateRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Inventory

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#LocateInventoryItem">LocateInventoryItem</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#LocateInventoryItemList">LocateInventoryItemList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateInventoryGetLocateInventoryResponse">ActiveV1AccountLocateInventoryGetLocateInventoryResponse</a>

Methods:

- <code title="get /active/v1/accounts/{account_id}/locates/inventory">client.Active.V1.Accounts.Locates.Inventory.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateInventoryService.GetLocateInventory">GetLocateInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateInventoryGetLocateInventoryParams">ActiveV1AccountLocateInventoryGetLocateInventoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountLocateInventoryGetLocateInventoryResponse">ActiveV1AccountLocateInventoryGetLocateInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ApStrategyParam">ApStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#BaseStrategyParams">BaseStrategyParams</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DarkStrategyParam">DarkStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DmaStrategyParam">DmaStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderStrategyUnionParam">OrderStrategyUnionParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderType">OrderType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PovStrategyParam">PovStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Side">Side</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TimeInForce">TimeInForce</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TrailingOffsetType">TrailingOffsetType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TwapStrategyParam">TwapStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Urgency">Urgency</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#VwapStrategyParam">VwapStrategyParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ApStrategy">ApStrategy</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#BaseStrategyParamsResp">BaseStrategyParamsResp</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DarkStrategy">DarkStrategy</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DmaStrategy">DmaStrategy</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderList">OrderList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderStatus">OrderStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderStrategyUnion">OrderStrategyUnion</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderType">OrderType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PovStrategy">PovStrategy</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Side">Side</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TimeInForce">TimeInForce</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TrailingOffsetType">TrailingOffsetType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#TwapStrategy">TwapStrategy</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Urgency">Urgency</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#VwapStrategy">VwapStrategy</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderCancelAllOrdersResponse">ActiveV1AccountOrderCancelAllOrdersResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderCancelOrderResponse">ActiveV1AccountOrderCancelOrderResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderGetOrderByIDResponse">ActiveV1AccountOrderGetOrderByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderGetOrdersResponse">ActiveV1AccountOrderGetOrdersResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderReplaceOrderResponse">ActiveV1AccountOrderReplaceOrderResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderSubmitOrdersResponse">ActiveV1AccountOrderSubmitOrdersResponse</a>

Methods:

- <code title="delete /active/v1/accounts/{account_id}/orders">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderService.CancelAllOrders">CancelAllOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderCancelAllOrdersParams">ActiveV1AccountOrderCancelAllOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderCancelAllOrdersResponse">ActiveV1AccountOrderCancelAllOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/accounts/{account_id}/orders/{order_id}">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderService.CancelOrder">CancelOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderCancelOrderParams">ActiveV1AccountOrderCancelOrderParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderCancelOrderResponse">ActiveV1AccountOrderCancelOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts/{account_id}/orders/{order_id}">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderService.GetOrderByID">GetOrderByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderGetOrderByIDParams">ActiveV1AccountOrderGetOrderByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderGetOrderByIDResponse">ActiveV1AccountOrderGetOrderByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts/{account_id}/orders">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderService.GetOrders">GetOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderGetOrdersParams">ActiveV1AccountOrderGetOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderGetOrdersResponse">ActiveV1AccountOrderGetOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /active/v1/accounts/{account_id}/orders/{order_id}">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderService.ReplaceOrder">ReplaceOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderReplaceOrderParams">ActiveV1AccountOrderReplaceOrderParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderReplaceOrderResponse">ActiveV1AccountOrderReplaceOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /active/v1/accounts/{account_id}/orders">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderService.SubmitOrders">SubmitOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderSubmitOrdersParams">ActiveV1AccountOrderSubmitOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountOrderSubmitOrdersResponse">ActiveV1AccountOrderSubmitOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### PortfolioHistory

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PortfolioHistoryResponse">PortfolioHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PortfolioHistorySegment">PortfolioHistorySegment</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse">ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse</a>

Methods:

- <code title="get /active/v1/accounts/{account_id}/portfolio-history">client.Active.V1.Accounts.PortfolioHistory.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPortfolioHistoryService.GetPortfolioHistory">GetPortfolioHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPortfolioHistoryGetPortfolioHistoryParams">ActiveV1AccountPortfolioHistoryGetPortfolioHistoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse">ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Position">Position</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PositionList">PositionList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PositionType">PositionType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionClosePositionResponse">ActiveV1AccountPositionClosePositionResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionClosePositionsResponse">ActiveV1AccountPositionClosePositionsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionGetPositionsResponse">ActiveV1AccountPositionGetPositionsResponse</a>

Methods:

- <code title="delete /active/v1/accounts/{account_id}/positions/{security_id_source}/{security_id}">client.Active.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionService.ClosePosition">ClosePosition</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionClosePositionParams">ActiveV1AccountPositionClosePositionParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionClosePositionResponse">ActiveV1AccountPositionClosePositionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/accounts/{account_id}/positions">client.Active.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionService.ClosePositions">ClosePositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionClosePositionsParams">ActiveV1AccountPositionClosePositionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionClosePositionsResponse">ActiveV1AccountPositionClosePositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts/{account_id}/positions">client.Active.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionService.GetPositions">GetPositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionGetPositionsParams">ActiveV1AccountPositionGetPositionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1AccountPositionGetPositionsResponse">ActiveV1AccountPositionGetPositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#APIKeyListEntry">APIKeyListEntry</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#APIKeyListEntryList">APIKeyListEntryList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Revocation">Revocation</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#RevocationList">RevocationList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyNewResponse">ActiveV1APIKeyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyListResponse">ActiveV1APIKeyListResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyRevokeResponse">ActiveV1APIKeyRevokeResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyRevokeAllResponse">ActiveV1APIKeyRevokeAllResponse</a>

Methods:

- <code title="post /active/v1/api_keys">client.Active.V1.APIKeys.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyNewParams">ActiveV1APIKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyNewResponse">ActiveV1APIKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/api_keys">client.Active.V1.APIKeys.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyListResponse">ActiveV1APIKeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/api_keys/{id}">client.Active.V1.APIKeys.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyRevokeResponse">ActiveV1APIKeyRevokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/api_keys">client.Active.V1.APIKeys.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyService.RevokeAll">RevokeAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1APIKeyRevokeAllResponse">ActiveV1APIKeyRevokeAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Calendars

#### Dividends

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DividendCalendarEvent">DividendCalendarEvent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DividendCalendarEventList">DividendCalendarEventList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DividendFrequency">DividendFrequency</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarDividendGetDividendsCalendarResponse">ActiveV1CalendarDividendGetDividendsCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/dividends">client.Active.V1.Calendars.Dividends.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarDividendService.GetDividendsCalendar">GetDividendsCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarDividendGetDividendsCalendarParams">ActiveV1CalendarDividendGetDividendsCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarDividendGetDividendsCalendarResponse">ActiveV1CalendarDividendGetDividendsCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Earnings

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EarningsCalendarEvent">EarningsCalendarEvent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EarningsCalendarEventList">EarningsCalendarEventList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarEarningGetEarningsCalendarResponse">ActiveV1CalendarEarningGetEarningsCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/earnings">client.Active.V1.Calendars.Earnings.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarEarningService.GetEarningsCalendar">GetEarningsCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarEarningGetEarningsCalendarParams">ActiveV1CalendarEarningGetEarningsCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarEarningGetEarningsCalendarResponse">ActiveV1CalendarEarningGetEarningsCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Economic

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EconomicCalendarEvent">EconomicCalendarEvent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EconomicCalendarEventList">EconomicCalendarEventList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#EconomicEventImpact">EconomicEventImpact</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarEconomicGetEconomicCalendarResponse">ActiveV1CalendarEconomicGetEconomicCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/economic">client.Active.V1.Calendars.Economic.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarEconomicService.GetEconomicCalendar">GetEconomicCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarEconomicGetEconomicCalendarParams">ActiveV1CalendarEconomicGetEconomicCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarEconomicGetEconomicCalendarResponse">ActiveV1CalendarEconomicGetEconomicCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### MarketHours

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
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse">ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/market-hours">client.Active.V1.Calendars.MarketHours.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarMarketHourService.GetMarketHoursCalendar">GetMarketHoursCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarMarketHourGetMarketHoursCalendarParams">ActiveV1CalendarMarketHourGetMarketHoursCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse">ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### MergersAcquisitions

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MergersAcquisitionsEvent">MergersAcquisitionsEvent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MergersAcquisitionsEventList">MergersAcquisitionsEventList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse">ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/mergers-acquisitions">client.Active.V1.Calendars.MergersAcquisitions.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarMergersAcquisitionService.GetMergersAndAcquisitionsCalendar">GetMergersAndAcquisitionsCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams">ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse">ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Splits

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#StockSplitEvent">StockSplitEvent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#StockSplitEventList">StockSplitEventList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarSplitGetSplitsCalendarResponse">ActiveV1CalendarSplitGetSplitsCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/splits">client.Active.V1.Calendars.Splits.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarSplitService.GetSplitsCalendar">GetSplitsCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarSplitGetSplitsCalendarParams">ActiveV1CalendarSplitGetSplitsCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarSplitGetSplitsCalendarResponse">ActiveV1CalendarSplitGetSplitsCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Summary

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CalendarDateSummary">CalendarDateSummary</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CalendarDateSummaryList">CalendarDateSummaryList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarSummaryGetCalendarSummaryResponse">ActiveV1CalendarSummaryGetCalendarSummaryResponse</a>

Methods:

- <code title="get /active/v1/calendars/summary">client.Active.V1.Calendars.Summary.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarSummaryService.GetCalendarSummary">GetCalendarSummary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarSummaryGetCalendarSummaryParams">ActiveV1CalendarSummaryGetCalendarSummaryParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1CalendarSummaryGetCalendarSummaryResponse">ActiveV1CalendarSummaryGetCalendarSummaryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Clock

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ClockDetail">ClockDetail</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ClockGetClockResponse">ActiveV1ClockGetClockResponse</a>

Methods:

- <code title="get /active/v1/clock">client.Active.V1.Clock.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ClockService.GetClock">GetClock</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ClockGetClockResponse">ActiveV1ClockGetClockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Instruments

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContractType">ContractType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AnalystRating">AnalystRating</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContractType">ContractType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ExerciseStyle">ExerciseStyle</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Instrument">Instrument</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentCore">InstrumentCore</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentCoreList">InstrumentCoreList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentEarnings">InstrumentEarnings</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentQuote">InstrumentQuote</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentSecurityID">InstrumentSecurityID</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ListingType">ListingType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OptionsContract">OptionsContract</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OptionsContractList">OptionsContractList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentGetInstrumentByIDResponse">ActiveV1InstrumentGetInstrumentByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentGetInstrumentsResponse">ActiveV1InstrumentGetInstrumentsResponse</a>

Methods:

- <code title="get /active/v1/instruments/{security_id_source}/{security_id}">client.Active.V1.Instruments.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentService.GetInstrumentByID">GetInstrumentByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentGetInstrumentByIDParams">ActiveV1InstrumentGetInstrumentByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentGetInstrumentByIDResponse">ActiveV1InstrumentGetInstrumentByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/instruments">client.Active.V1.Instruments.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentService.GetInstruments">GetInstruments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentGetInstrumentsParams">ActiveV1InstrumentGetInstrumentsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentGetInstrumentsResponse">ActiveV1InstrumentGetInstrumentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### AnalystReporting

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AnalystDistribution">AnalystDistribution</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#InstrumentAnalystConsensus">InstrumentAnalystConsensus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PriceTarget">PriceTarget</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse">ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse</a>

Methods:

- <code title="get /active/v1/instruments/{security_id_source}/{security_id}/analyst-reporting">client.Active.V1.Instruments.AnalystReporting.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentAnalystReportingService.GetInstrumentAnalystConsensus">GetInstrumentAnalystConsensus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams">ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse">ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Events

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
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentEventGetAllInstrumentEventsResponse">ActiveV1InstrumentEventGetAllInstrumentEventsResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentEventGetInstrumentEventsResponse">ActiveV1InstrumentEventGetInstrumentEventsResponse</a>

Methods:

- <code title="get /active/v1/instruments/events">client.Active.V1.Instruments.Events.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentEventService.GetAllInstrumentEvents">GetAllInstrumentEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentEventGetAllInstrumentEventsParams">ActiveV1InstrumentEventGetAllInstrumentEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentEventGetAllInstrumentEventsResponse">ActiveV1InstrumentEventGetAllInstrumentEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/instruments/{security_id_source}/{security_id}/events">client.Active.V1.Instruments.Events.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentEventService.GetInstrumentEvents">GetInstrumentEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentEventGetInstrumentEventsParams">ActiveV1InstrumentEventGetInstrumentEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentEventGetInstrumentEventsResponse">ActiveV1InstrumentEventGetInstrumentEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Options

##### Contracts

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentOptionContractGetOptionContractsResponse">ActiveV1InstrumentOptionContractGetOptionContractsResponse</a>

Methods:

- <code title="get /active/v1/instruments/options/contracts">client.Active.V1.Instruments.Options.Contracts.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentOptionContractService.GetOptionContracts">GetOptionContracts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentOptionContractGetOptionContractsParams">ActiveV1InstrumentOptionContractGetOptionContractsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentOptionContractGetOptionContractsResponse">ActiveV1InstrumentOptionContractGetOptionContractsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Reporting

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentReportingGetInstrumentReportingResponse">ActiveV1InstrumentReportingGetInstrumentReportingResponse</a>

Methods:

- <code title="get /active/v1/instruments/{security_id_source}/{security_id}/reporting">client.Active.V1.Instruments.Reporting.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentReportingService.GetInstrumentReporting">GetInstrumentReporting</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentReportingGetInstrumentReportingParams">ActiveV1InstrumentReportingGetInstrumentReportingParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentReportingGetInstrumentReportingResponse">ActiveV1InstrumentReportingGetInstrumentReportingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Venues

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#DisplayType">DisplayType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#GtdAccepts">GtdAccepts</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Venue">Venue</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#VenueList">VenueList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#VenueSession">VenueSession</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentVenueGetVenuesResponse">ActiveV1InstrumentVenueGetVenuesResponse</a>

Methods:

- <code title="get /active/v1/instruments/venues">client.Active.V1.Instruments.Venues.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentVenueService.GetVenues">GetVenues</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1InstrumentVenueGetVenuesResponse">ActiveV1InstrumentVenueGetVenuesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### MarketData

#### Snapshot

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketDataSnapshot">MarketDataSnapshot</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MarketDataSnapshotList">MarketDataSnapshotList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SnapshotLastTrade">SnapshotLastTrade</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SnapshotQuote">SnapshotQuote</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SnapshotSession">SnapshotSession</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1MarketDataSnapshotGetSnapshotsResponse">ActiveV1MarketDataSnapshotGetSnapshotsResponse</a>

Methods:

- <code title="get /active/v1/market-data/snapshot">client.Active.V1.MarketData.Snapshot.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1MarketDataSnapshotService.GetSnapshots">GetSnapshots</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1MarketDataSnapshotGetSnapshotsParams">ActiveV1MarketDataSnapshotGetSnapshotsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1MarketDataSnapshotGetSnapshotsResponse">ActiveV1MarketDataSnapshotGetSnapshotsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### News

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewsInstrument">NewsInstrument</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewsItem">NewsItem</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewsItemList">NewsItemList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#NewsType">NewsType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1NewsGetNewsResponse">ActiveV1NewsGetNewsResponse</a>

Methods:

- <code title="get /active/v1/news">client.Active.V1.News.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1NewsService.GetNews">GetNews</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1NewsGetNewsParams">ActiveV1NewsGetNewsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1NewsGetNewsResponse">ActiveV1NewsGetNewsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OmniAI

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CancelResponsePayload">CancelResponsePayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartChartPayload">ContentPartChartPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartCustomPayload">ContentPartCustomPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartStructuredActionPayload">ContentPartStructuredActionPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartSuggestedActionsPayload">ContentPartSuggestedActionsPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartTextPayload">ContentPartTextPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ContentPartThinkingPayload">ContentPartThinkingPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CreateFeedbackResponse">CreateFeedbackResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CreateMessageResponse">CreateMessageResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#CreateThreadResponse">CreateThreadResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ErrorStatus">ErrorStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageContent">MessageContent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageContentPartUnion">MessageContentPartUnion</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageList">MessageList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageOutcome">MessageOutcome</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#MessageRole">MessageRole</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OpenChartAction">OpenChartAction</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OpenScreenerAction">OpenScreenerAction</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderPayload">OrderPayload</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#OrderStrategyType">OrderStrategyType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#PrefillOrderAction">PrefillOrderAction</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Response">Response</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ResponseContent">ResponseContent</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ResponseContentPartUnion">ResponseContentPartUnion</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ResponseStatus">ResponseStatus</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#StructuredActionUnion">StructuredActionUnion</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Thread">Thread</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ThreadList">ThreadList</a>

#### Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIMessageGetMessageResponse">ActiveV1OmniAIMessageGetMessageResponse</a>

Methods:

- <code title="get /active/v1/omni-ai/messages/{message_id}">client.Active.V1.OmniAI.Messages.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIMessageService.GetMessage">GetMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIMessageGetMessageParams">ActiveV1OmniAIMessageGetMessageParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIMessageGetMessageResponse">ActiveV1OmniAIMessageGetMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Feedback

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIMessageFeedbackNewFeedbackResponse">ActiveV1OmniAIMessageFeedbackNewFeedbackResponse</a>

Methods:

- <code title="post /active/v1/omni-ai/messages/{message_id}/feedback">client.Active.V1.OmniAI.Messages.Feedback.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIMessageFeedbackService.NewFeedback">NewFeedback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIMessageFeedbackNewFeedbackParams">ActiveV1OmniAIMessageFeedbackNewFeedbackParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIMessageFeedbackNewFeedbackResponse">ActiveV1OmniAIMessageFeedbackNewFeedbackResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Responses

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIResponseCancelResponseResponse">ActiveV1OmniAIResponseCancelResponseResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIResponseGetResponseResponse">ActiveV1OmniAIResponseGetResponseResponse</a>

Methods:

- <code title="delete /active/v1/omni-ai/responses/{response_id}">client.Active.V1.OmniAI.Responses.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIResponseService.CancelResponse">CancelResponse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIResponseCancelResponseParams">ActiveV1OmniAIResponseCancelResponseParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIResponseCancelResponseResponse">ActiveV1OmniAIResponseCancelResponseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/omni-ai/responses/{response_id}">client.Active.V1.OmniAI.Responses.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIResponseService.GetResponse">GetResponse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIResponseGetResponseParams">ActiveV1OmniAIResponseGetResponseParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIResponseGetResponseResponse">ActiveV1OmniAIResponseGetResponseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Threads

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadNewThreadResponse">ActiveV1OmniAIThreadNewThreadResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadGetThreadResponse">ActiveV1OmniAIThreadGetThreadResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadListThreadsResponse">ActiveV1OmniAIThreadListThreadsResponse</a>

Methods:

- <code title="post /active/v1/omni-ai/threads">client.Active.V1.OmniAI.Threads.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadService.NewThread">NewThread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadNewThreadParams">ActiveV1OmniAIThreadNewThreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadNewThreadResponse">ActiveV1OmniAIThreadNewThreadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/omni-ai/threads/{thread_id}">client.Active.V1.OmniAI.Threads.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadService.GetThread">GetThread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadGetThreadParams">ActiveV1OmniAIThreadGetThreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadGetThreadResponse">ActiveV1OmniAIThreadGetThreadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/omni-ai/threads">client.Active.V1.OmniAI.Threads.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadService.ListThreads">ListThreads</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadListThreadsParams">ActiveV1OmniAIThreadListThreadsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadListThreadsResponse">ActiveV1OmniAIThreadListThreadsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadMessageNewMessageResponse">ActiveV1OmniAIThreadMessageNewMessageResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadMessageListMessagesResponse">ActiveV1OmniAIThreadMessageListMessagesResponse</a>

Methods:

- <code title="post /active/v1/omni-ai/threads/{thread_id}/messages">client.Active.V1.OmniAI.Threads.Messages.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadMessageService.NewMessage">NewMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadMessageNewMessageParams">ActiveV1OmniAIThreadMessageNewMessageParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadMessageNewMessageResponse">ActiveV1OmniAIThreadMessageNewMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/omni-ai/threads/{thread_id}/messages">client.Active.V1.OmniAI.Threads.Messages.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadMessageService.ListMessages">ListMessages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadMessageListMessagesParams">ActiveV1OmniAIThreadMessageListMessagesParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadMessageListMessagesResponse">ActiveV1OmniAIThreadMessageListMessagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Response

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadResponseGetThreadResponseResponse">ActiveV1OmniAIThreadResponseGetThreadResponseResponse</a>

Methods:

- <code title="get /active/v1/omni-ai/threads/{thread_id}/response">client.Active.V1.OmniAI.Threads.Response.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadResponseService.GetThreadResponse">GetThreadResponse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadResponseGetThreadResponseParams">ActiveV1OmniAIThreadResponseGetThreadResponseParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1OmniAIThreadResponseGetThreadResponseResponse">ActiveV1OmniAIThreadResponseGetThreadResponseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### SavedScreeners

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SavedScreenerFilterParam">SavedScreenerFilterParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#SavedScreenerFilter">SavedScreenerFilter</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerEntry">ScreenerEntry</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerEntryList">ScreenerEntryList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerNewScreenerResponse">ActiveV1SavedScreenerNewScreenerResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerGetScreenerByIDResponse">ActiveV1SavedScreenerGetScreenerByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerListScreenersResponse">ActiveV1SavedScreenerListScreenersResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerUpdateScreenerResponse">ActiveV1SavedScreenerUpdateScreenerResponse</a>

Methods:

- <code title="post /active/v1/saved-screeners">client.Active.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerService.NewScreener">NewScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerNewScreenerParams">ActiveV1SavedScreenerNewScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerNewScreenerResponse">ActiveV1SavedScreenerNewScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/saved-screeners/{screener_id}">client.Active.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerService.DeleteScreener">DeleteScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, screenerID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /active/v1/saved-screeners/{screener_id}">client.Active.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerService.GetScreenerByID">GetScreenerByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, screenerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerGetScreenerByIDResponse">ActiveV1SavedScreenerGetScreenerByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/saved-screeners">client.Active.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerService.ListScreeners">ListScreeners</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerListScreenersResponse">ActiveV1SavedScreenerListScreenersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /active/v1/saved-screeners/{screener_id}">client.Active.V1.SavedScreeners.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerService.UpdateScreener">UpdateScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, screenerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerUpdateScreenerParams">ActiveV1SavedScreenerUpdateScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1SavedScreenerUpdateScreenerResponse">ActiveV1SavedScreenerUpdateScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Screener

Params Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldLookback">FieldLookback</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldPeriod">FieldPeriod</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldRefParam">FieldRefParam</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldType">FieldType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldLookback">FieldLookback</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldPeriod">FieldPeriod</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldRef">FieldRef</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#FieldType">FieldType</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerColumn">ScreenerColumn</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerFilter">ScreenerFilter</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerItem">ScreenerItem</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerItemList">ScreenerItemList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerRow">ScreenerRow</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ScreenerRowList">ScreenerRowList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ScreenerGetScreenerResponse">ActiveV1ScreenerGetScreenerResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ScreenerSearchScreenerResponse">ActiveV1ScreenerSearchScreenerResponse</a>

Methods:

- <code title="get /active/v1/screener">client.Active.V1.Screener.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ScreenerService.GetScreener">GetScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ScreenerGetScreenerParams">ActiveV1ScreenerGetScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ScreenerGetScreenerResponse">ActiveV1ScreenerGetScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /active/v1/screener">client.Active.V1.Screener.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ScreenerService.SearchScreener">SearchScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ScreenerSearchScreenerParams">ActiveV1ScreenerSearchScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1ScreenerSearchScreenerResponse">ActiveV1ScreenerSearchScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Version

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#Version">Version</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1VersionGetVersionResponse">ActiveV1VersionGetVersionResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1VersionUpdateVersionResponse">ActiveV1VersionUpdateVersionResponse</a>

Methods:

- <code title="get /active/v1/version">client.Active.V1.Version.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1VersionService.GetVersion">GetVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1VersionGetVersionResponse">ActiveV1VersionGetVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /active/v1/version">client.Active.V1.Version.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1VersionService.UpdateVersion">UpdateVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1VersionUpdateVersionResponse">ActiveV1VersionUpdateVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Watchlists

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#WatchlistDetail">WatchlistDetail</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#WatchlistEntry">WatchlistEntry</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#WatchlistEntryList">WatchlistEntryList</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#WatchlistItemEntry">WatchlistItemEntry</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistNewWatchlistResponse">ActiveV1WatchlistNewWatchlistResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistGetWatchlistByIDResponse">ActiveV1WatchlistGetWatchlistByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistGetWatchlistsResponse">ActiveV1WatchlistGetWatchlistsResponse</a>

Methods:

- <code title="post /active/v1/watchlists">client.Active.V1.Watchlists.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistService.NewWatchlist">NewWatchlist</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistNewWatchlistParams">ActiveV1WatchlistNewWatchlistParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistNewWatchlistResponse">ActiveV1WatchlistNewWatchlistResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/watchlists/{watchlist_id}">client.Active.V1.Watchlists.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistService.DeleteWatchlist">DeleteWatchlist</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /active/v1/watchlists/{watchlist_id}">client.Active.V1.Watchlists.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistService.GetWatchlistByID">GetWatchlistByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistGetWatchlistByIDResponse">ActiveV1WatchlistGetWatchlistByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/watchlists">client.Active.V1.Watchlists.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistService.GetWatchlists">GetWatchlists</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistGetWatchlistsResponse">ActiveV1WatchlistGetWatchlistsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Items

Response Types:

- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#AddWatchlistItemData">AddWatchlistItemData</a>
- <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistItemAddWatchlistItemResponse">ActiveV1WatchlistItemAddWatchlistItemResponse</a>

Methods:

- <code title="post /active/v1/watchlists/{watchlist_id}/items">client.Active.V1.Watchlists.Items.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistItemService.AddWatchlistItem">AddWatchlistItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistItemAddWatchlistItemParams">ActiveV1WatchlistItemAddWatchlistItemParams</a>) (\*<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistItemAddWatchlistItemResponse">ActiveV1WatchlistItemAddWatchlistItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/watchlists/{watchlist_id}/items/{item_id}">client.Active.V1.Watchlists.Items.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistItemService.DeleteWatchlistItem">DeleteWatchlistItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clear-street/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WatchlistItemDeleteWatchlistItemParams">ActiveV1WatchlistItemDeleteWatchlistItemParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Ws

Methods:

- <code title="get /active/v1/ws">client.Active.V1.Ws.<a href="https://pkg.go.dev/github.com/clear-street/clear-street-go#ActiveV1WService.WebsocketHandler">WebsocketHandler</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
