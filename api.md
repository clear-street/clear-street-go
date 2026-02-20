# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go/shared#APIError">APIError</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go/shared#BaseResponse">BaseResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go/shared#ResponseMetadata">ResponseMetadata</a>

# Active

## V1

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#SecurityIDSource">SecurityIDSource</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#SecurityType">SecurityType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#SecurityIDSource">SecurityIDSource</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#SecurityType">SecurityType</a>

### Accounts

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#RiskSettingsParam">RiskSettingsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AccountKind">AccountKind</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AccountList">AccountList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AccountSettings">AccountSettings</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AccountStatus">AccountStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AccountSubkind">AccountSubkind</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#OrderList">OrderList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#RiskSettings">RiskSettings</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountGetAccountByIDResponse">ActiveV1AccountGetAccountByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountGetAccountsResponse">ActiveV1AccountGetAccountsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPatchAccountByIDResponse">ActiveV1AccountPatchAccountByIDResponse</a>

Methods:

- <code title="get /active/v1/accounts/{account_id}">client.Active.V1.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountService.GetAccountByID">GetAccountByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountGetAccountByIDResponse">ActiveV1AccountGetAccountByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts">client.Active.V1.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountService.GetAccounts">GetAccounts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountGetAccountsParams">ActiveV1AccountGetAccountsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountGetAccountsResponse">ActiveV1AccountGetAccountsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /active/v1/accounts/{account_id}">client.Active.V1.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountService.PatchAccountByID">PatchAccountByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPatchAccountByIDParams">ActiveV1AccountPatchAccountByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPatchAccountByIDResponse">ActiveV1AccountPatchAccountByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Balances

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AccountBalances">AccountBalances</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#APITimestamp">APITimestamp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MarginType">MarginType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#RegTBalance">RegTBalance</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountBalanceGetAccountBalancesResponse">ActiveV1AccountBalanceGetAccountBalancesResponse</a>

Methods:

- <code title="get /active/v1/accounts/{account_id}/balances">client.Active.V1.Accounts.Balances.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountBalanceService.GetAccountBalances">GetAccountBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountBalanceGetAccountBalancesResponse">ActiveV1AccountBalanceGetAccountBalancesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Locates

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#LocateOrderStatus">LocateOrderStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#LocateOrder">LocateOrder</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#LocateOrderList">LocateOrderList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#LocateOrderStatus">LocateOrderStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateNewLocateRequestResponse">ActiveV1AccountLocateNewLocateRequestResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateGetLocateRequestsResponse">ActiveV1AccountLocateGetLocateRequestsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateUpdateLocateRequestResponse">ActiveV1AccountLocateUpdateLocateRequestResponse</a>

Methods:

- <code title="post /active/v1/accounts/{account_id}/locates">client.Active.V1.Accounts.Locates.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateService.NewLocateRequest">NewLocateRequest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateNewLocateRequestParams">ActiveV1AccountLocateNewLocateRequestParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateNewLocateRequestResponse">ActiveV1AccountLocateNewLocateRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts/{account_id}/locates">client.Active.V1.Accounts.Locates.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateService.GetLocateRequests">GetLocateRequests</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateGetLocateRequestsParams">ActiveV1AccountLocateGetLocateRequestsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateGetLocateRequestsResponse">ActiveV1AccountLocateGetLocateRequestsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /active/v1/accounts/{account_id}/locates">client.Active.V1.Accounts.Locates.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateService.UpdateLocateRequest">UpdateLocateRequest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateUpdateLocateRequestParams">ActiveV1AccountLocateUpdateLocateRequestParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateUpdateLocateRequestResponse">ActiveV1AccountLocateUpdateLocateRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Inventory

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#LocateInventoryItem">LocateInventoryItem</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#LocateInventoryItemList">LocateInventoryItemList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateInventoryGetLocateInventoryResponse">ActiveV1AccountLocateInventoryGetLocateInventoryResponse</a>

Methods:

- <code title="get /active/v1/accounts/{account_id}/locates/inventory">client.Active.V1.Accounts.Locates.Inventory.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateInventoryService.GetLocateInventory">GetLocateInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateInventoryGetLocateInventoryParams">ActiveV1AccountLocateInventoryGetLocateInventoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountLocateInventoryGetLocateInventoryResponse">ActiveV1AccountLocateInventoryGetLocateInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ApStrategyParam">ApStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#APIDecimal64">APIDecimal64</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#BaseStrategyParams">BaseStrategyParams</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#DarkStrategyParam">DarkStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#DmaStrategyParam">DmaStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#OrderStrategyUnionParam">OrderStrategyUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#OrderType">OrderType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#PovStrategyParam">PovStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Side">Side</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#SorStrategyParam">SorStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#TimeInForce">TimeInForce</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#TrailingOffsetType">TrailingOffsetType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#TwapStrategyParam">TwapStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Urgency">Urgency</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#VwapStrategyParam">VwapStrategyParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ApStrategy">ApStrategy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#APIDecimal64">APIDecimal64</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#BaseStrategyParamsResp">BaseStrategyParamsResp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#DarkStrategy">DarkStrategy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#DmaStrategy">DmaStrategy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#OrderStatus">OrderStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#OrderStrategyUnion">OrderStrategyUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#OrderType">OrderType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#PovStrategy">PovStrategy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Side">Side</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#SorStrategy">SorStrategy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#TimeInForce">TimeInForce</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#TrailingOffsetType">TrailingOffsetType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#TwapStrategy">TwapStrategy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Urgency">Urgency</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#VwapStrategy">VwapStrategy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderCancelAllOrdersResponse">ActiveV1AccountOrderCancelAllOrdersResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderCancelOrderResponse">ActiveV1AccountOrderCancelOrderResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderGetOrderByIDResponse">ActiveV1AccountOrderGetOrderByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderGetOrdersResponse">ActiveV1AccountOrderGetOrdersResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderReplaceOrderResponse">ActiveV1AccountOrderReplaceOrderResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderSubmitOrdersResponse">ActiveV1AccountOrderSubmitOrdersResponse</a>

Methods:

- <code title="delete /active/v1/accounts/{account_id}/orders">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderService.CancelAllOrders">CancelAllOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderCancelAllOrdersParams">ActiveV1AccountOrderCancelAllOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderCancelAllOrdersResponse">ActiveV1AccountOrderCancelAllOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/accounts/{account_id}/orders/{order_id}">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderService.CancelOrder">CancelOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderCancelOrderParams">ActiveV1AccountOrderCancelOrderParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderCancelOrderResponse">ActiveV1AccountOrderCancelOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts/{account_id}/orders/{order_id}">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderService.GetOrderByID">GetOrderByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderGetOrderByIDParams">ActiveV1AccountOrderGetOrderByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderGetOrderByIDResponse">ActiveV1AccountOrderGetOrderByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts/{account_id}/orders">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderService.GetOrders">GetOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderGetOrdersParams">ActiveV1AccountOrderGetOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderGetOrdersResponse">ActiveV1AccountOrderGetOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /active/v1/accounts/{account_id}/orders/{order_id}">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderService.ReplaceOrder">ReplaceOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderReplaceOrderParams">ActiveV1AccountOrderReplaceOrderParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderReplaceOrderResponse">ActiveV1AccountOrderReplaceOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /active/v1/accounts/{account_id}/orders">client.Active.V1.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderService.SubmitOrders">SubmitOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderSubmitOrdersParams">ActiveV1AccountOrderSubmitOrdersParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountOrderSubmitOrdersResponse">ActiveV1AccountOrderSubmitOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### PortfolioHistory

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#PortfolioHistoryResponse">PortfolioHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#PortfolioHistorySegment">PortfolioHistorySegment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse">ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse</a>

Methods:

- <code title="get /active/v1/accounts/{account_id}/portfolio-history">client.Active.V1.Accounts.PortfolioHistory.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPortfolioHistoryService.GetPortfolioHistory">GetPortfolioHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPortfolioHistoryGetPortfolioHistoryParams">ActiveV1AccountPortfolioHistoryGetPortfolioHistoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse">ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Position">Position</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#PositionList">PositionList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionClosePositionResponse">ActiveV1AccountPositionClosePositionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionClosePositionsResponse">ActiveV1AccountPositionClosePositionsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionGetPositionsResponse">ActiveV1AccountPositionGetPositionsResponse</a>

Methods:

- <code title="delete /active/v1/accounts/{account_id}/positions/{security_id_source}/{security_id}">client.Active.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionService.ClosePosition">ClosePosition</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionClosePositionParams">ActiveV1AccountPositionClosePositionParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionClosePositionResponse">ActiveV1AccountPositionClosePositionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/accounts/{account_id}/positions">client.Active.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionService.ClosePositions">ClosePositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionClosePositionsParams">ActiveV1AccountPositionClosePositionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionClosePositionsResponse">ActiveV1AccountPositionClosePositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/accounts/{account_id}/positions">client.Active.V1.Accounts.Positions.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionService.GetPositions">GetPositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionGetPositionsParams">ActiveV1AccountPositionGetPositionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AccountPositionGetPositionsResponse">ActiveV1AccountPositionGetPositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#APIKeyListEntry">APIKeyListEntry</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#APIKeyListEntryList">APIKeyListEntryList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Revocation">Revocation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#RevocationList">RevocationList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyNewResponse">ActiveV1APIKeyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyListResponse">ActiveV1APIKeyListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyRevokeResponse">ActiveV1APIKeyRevokeResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyRevokeAllResponse">ActiveV1APIKeyRevokeAllResponse</a>

Methods:

- <code title="post /active/v1/api_keys">client.Active.V1.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyNewParams">ActiveV1APIKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyNewResponse">ActiveV1APIKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/api_keys">client.Active.V1.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyListResponse">ActiveV1APIKeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/api_keys/{id}">client.Active.V1.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyRevokeResponse">ActiveV1APIKeyRevokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/api_keys">client.Active.V1.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyService.RevokeAll">RevokeAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1APIKeyRevokeAllResponse">ActiveV1APIKeyRevokeAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Assistant

#### Prompts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#PromptResult">PromptResult</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#PromptStatus">PromptStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#RunPromptResponse">RunPromptResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AssistantPromptGetPromptResultResponse">ActiveV1AssistantPromptGetPromptResultResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AssistantPromptRunPromptResponse">ActiveV1AssistantPromptRunPromptResponse</a>

Methods:

- <code title="get /active/v1/assistant/prompts/{id}">client.Active.V1.Assistant.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AssistantPromptService.GetPromptResult">GetPromptResult</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AssistantPromptGetPromptResultParams">ActiveV1AssistantPromptGetPromptResultParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AssistantPromptGetPromptResultResponse">ActiveV1AssistantPromptGetPromptResultResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /active/v1/assistant/prompts">client.Active.V1.Assistant.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AssistantPromptService.RunPrompt">RunPrompt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AssistantPromptRunPromptParams">ActiveV1AssistantPromptRunPromptParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1AssistantPromptRunPromptResponse">ActiveV1AssistantPromptRunPromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Calendars

#### Dividends

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#DividendCalendarEvent">DividendCalendarEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#DividendCalendarEventList">DividendCalendarEventList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarDividendGetDividendsCalendarResponse">ActiveV1CalendarDividendGetDividendsCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/dividends">client.Active.V1.Calendars.Dividends.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarDividendService.GetDividendsCalendar">GetDividendsCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarDividendGetDividendsCalendarParams">ActiveV1CalendarDividendGetDividendsCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarDividendGetDividendsCalendarResponse">ActiveV1CalendarDividendGetDividendsCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Earnings

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#EarningsCalendarEvent">EarningsCalendarEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#EarningsCalendarEventList">EarningsCalendarEventList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarEarningGetEarningsCalendarResponse">ActiveV1CalendarEarningGetEarningsCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/earnings">client.Active.V1.Calendars.Earnings.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarEarningService.GetEarningsCalendar">GetEarningsCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarEarningGetEarningsCalendarParams">ActiveV1CalendarEarningGetEarningsCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarEarningGetEarningsCalendarResponse">ActiveV1CalendarEarningGetEarningsCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Economic

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#EconomicCalendarEvent">EconomicCalendarEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#EconomicCalendarEventList">EconomicCalendarEventList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarEconomicGetEconomicCalendarResponse">ActiveV1CalendarEconomicGetEconomicCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/economic">client.Active.V1.Calendars.Economic.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarEconomicService.GetEconomicCalendar">GetEconomicCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarEconomicGetEconomicCalendarParams">ActiveV1CalendarEconomicGetEconomicCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarEconomicGetEconomicCalendarResponse">ActiveV1CalendarEconomicGetEconomicCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### MarketHours

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MarketHoursDetail">MarketHoursDetail</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MarketHoursDetailList">MarketHoursDetailList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse">ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/market-hours">client.Active.V1.Calendars.MarketHours.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarMarketHourService.GetMarketHoursCalendar">GetMarketHoursCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarMarketHourGetMarketHoursCalendarParams">ActiveV1CalendarMarketHourGetMarketHoursCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse">ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### MergersAcquisitions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MergersAcquisitionsEvent">MergersAcquisitionsEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MergersAcquisitionsEventList">MergersAcquisitionsEventList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse">ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/mergers-acquisitions">client.Active.V1.Calendars.MergersAcquisitions.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarMergersAcquisitionService.GetMergersAndAcquisitionsCalendar">GetMergersAndAcquisitionsCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams">ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse">ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Splits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#StockSplitEvent">StockSplitEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#StockSplitEventList">StockSplitEventList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarSplitGetSplitsCalendarResponse">ActiveV1CalendarSplitGetSplitsCalendarResponse</a>

Methods:

- <code title="get /active/v1/calendars/splits">client.Active.V1.Calendars.Splits.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarSplitService.GetSplitsCalendar">GetSplitsCalendar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarSplitGetSplitsCalendarParams">ActiveV1CalendarSplitGetSplitsCalendarParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarSplitGetSplitsCalendarResponse">ActiveV1CalendarSplitGetSplitsCalendarResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Summary

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#CalendarDateSummary">CalendarDateSummary</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#CalendarDateSummaryList">CalendarDateSummaryList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarSummaryGetCalendarSummaryResponse">ActiveV1CalendarSummaryGetCalendarSummaryResponse</a>

Methods:

- <code title="get /active/v1/calendars/summary">client.Active.V1.Calendars.Summary.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarSummaryService.GetCalendarSummary">GetCalendarSummary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarSummaryGetCalendarSummaryParams">ActiveV1CalendarSummaryGetCalendarSummaryParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1CalendarSummaryGetCalendarSummaryResponse">ActiveV1CalendarSummaryGetCalendarSummaryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Instruments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Instrument">Instrument</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentCore">InstrumentCore</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentCoreList">InstrumentCoreList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentEarnings">InstrumentEarnings</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentQuote">InstrumentQuote</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentSecurityID">InstrumentSecurityID</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentGetInstrumentByIDResponse">ActiveV1InstrumentGetInstrumentByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentGetInstrumentsResponse">ActiveV1InstrumentGetInstrumentsResponse</a>

Methods:

- <code title="get /active/v1/instruments/{security_id_source}/{security_id}">client.Active.V1.Instruments.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentService.GetInstrumentByID">GetInstrumentByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentGetInstrumentByIDParams">ActiveV1InstrumentGetInstrumentByIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentGetInstrumentByIDResponse">ActiveV1InstrumentGetInstrumentByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/instruments">client.Active.V1.Instruments.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentService.GetInstruments">GetInstruments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentGetInstrumentsParams">ActiveV1InstrumentGetInstrumentsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentGetInstrumentsResponse">ActiveV1InstrumentGetInstrumentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### AnalystReporting

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AnalystDistribution">AnalystDistribution</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AnalystRating">AnalystRating</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentAnalystConsensus">InstrumentAnalystConsensus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#PriceTarget">PriceTarget</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse">ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse</a>

Methods:

- <code title="get /active/v1/instruments/{security_id_source}/{security_id}/analyst-reporting">client.Active.V1.Instruments.AnalystReporting.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentAnalystReportingService.GetInstrumentAnalystConsensus">GetInstrumentAnalystConsensus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams">ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse">ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Events

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentAllEventsData">InstrumentAllEventsData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentDividendEvent">InstrumentDividendEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentEventsByDate">InstrumentEventsByDate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentEventsData">InstrumentEventsData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#InstrumentSplitEvent">InstrumentSplitEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentEventGetAllInstrumentEventsResponse">ActiveV1InstrumentEventGetAllInstrumentEventsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentEventGetInstrumentEventsResponse">ActiveV1InstrumentEventGetInstrumentEventsResponse</a>

Methods:

- <code title="get /active/v1/instruments/events">client.Active.V1.Instruments.Events.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentEventService.GetAllInstrumentEvents">GetAllInstrumentEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentEventGetAllInstrumentEventsParams">ActiveV1InstrumentEventGetAllInstrumentEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentEventGetAllInstrumentEventsResponse">ActiveV1InstrumentEventGetAllInstrumentEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/instruments/{security_id_source}/{security_id}/events">client.Active.V1.Instruments.Events.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentEventService.GetInstrumentEvents">GetInstrumentEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentEventGetInstrumentEventsParams">ActiveV1InstrumentEventGetInstrumentEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentEventGetInstrumentEventsResponse">ActiveV1InstrumentEventGetInstrumentEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Reporting

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentReportingGetInstrumentReportingResponse">ActiveV1InstrumentReportingGetInstrumentReportingResponse</a>

Methods:

- <code title="get /active/v1/instruments/{security_id_source}/{security_id}/reporting">client.Active.V1.Instruments.Reporting.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentReportingService.GetInstrumentReporting">GetInstrumentReporting</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, securityID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentReportingGetInstrumentReportingParams">ActiveV1InstrumentReportingGetInstrumentReportingParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentReportingGetInstrumentReportingResponse">ActiveV1InstrumentReportingGetInstrumentReportingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Venues

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Venue">Venue</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#VenueList">VenueList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentVenueGetVenuesResponse">ActiveV1InstrumentVenueGetVenuesResponse</a>

Methods:

- <code title="get /active/v1/instruments/venues">client.Active.V1.Instruments.Venues.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentVenueService.GetVenues">GetVenues</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1InstrumentVenueGetVenuesResponse">ActiveV1InstrumentVenueGetVenuesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Iris

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Thread">Thread</a>

#### Feedback

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#CreateFeedbackResponse">CreateFeedbackResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisFeedbackNewFeedbackResponse">ActiveV1IrisFeedbackNewFeedbackResponse</a>

Methods:

- <code title="post /active/v1/iris/feedback">client.Active.V1.Iris.Feedback.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisFeedbackService.NewFeedback">NewFeedback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisFeedbackNewFeedbackParams">ActiveV1IrisFeedbackNewFeedbackParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisFeedbackNewFeedbackResponse">ActiveV1IrisFeedbackNewFeedbackResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Runs

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Capability">Capability</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#CancelRunResponse">CancelRunResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Capability">Capability</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ContentPartUnion">ContentPartUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#GetRunResponse">GetRunResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MessageContent">MessageContent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MessageRole">MessageRole</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Run">Run</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#RunStatus">RunStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#StartRunResponse">StartRunResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunCancelRunResponse">ActiveV1IrisRunCancelRunResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunGetRunResponse">ActiveV1IrisRunGetRunResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunStartRunResponse">ActiveV1IrisRunStartRunResponse</a>

Methods:

- <code title="delete /active/v1/iris/runs/{run_id}">client.Active.V1.Iris.Runs.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunService.CancelRun">CancelRun</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunCancelRunParams">ActiveV1IrisRunCancelRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunCancelRunResponse">ActiveV1IrisRunCancelRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/iris/runs/{run_id}">client.Active.V1.Iris.Runs.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunService.GetRun">GetRun</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunGetRunParams">ActiveV1IrisRunGetRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunGetRunResponse">ActiveV1IrisRunGetRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /active/v1/iris/runs">client.Active.V1.Iris.Runs.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunService.StartRun">StartRun</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunStartRunParams">ActiveV1IrisRunStartRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisRunStartRunResponse">ActiveV1IrisRunStartRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Threads

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#GetThreadResponse">GetThreadResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ListThreadsResponse">ListThreadsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadGetThreadResponse">ActiveV1IrisThreadGetThreadResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadListThreadsResponse">ActiveV1IrisThreadListThreadsResponse</a>

Methods:

- <code title="get /active/v1/iris/threads/{thread_id}">client.Active.V1.Iris.Threads.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadService.GetThread">GetThread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadGetThreadParams">ActiveV1IrisThreadGetThreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadGetThreadResponse">ActiveV1IrisThreadGetThreadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/iris/threads">client.Active.V1.Iris.Threads.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadService.ListThreads">ListThreads</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadListThreadsParams">ActiveV1IrisThreadListThreadsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadListThreadsResponse">ActiveV1IrisThreadListThreadsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ListMessagesResponse">ListMessagesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadMessageListMessagesResponse">ActiveV1IrisThreadMessageListMessagesResponse</a>

Methods:

- <code title="get /active/v1/iris/threads/{thread_id}/messages">client.Active.V1.Iris.Threads.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadMessageService.ListMessages">ListMessages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadMessageListMessagesParams">ActiveV1IrisThreadMessageListMessagesParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1IrisThreadMessageListMessagesResponse">ActiveV1IrisThreadMessageListMessagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### MarketData

#### Snapshot

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MarketDataSnapshot">MarketDataSnapshot</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#MarketDataSnapshotList">MarketDataSnapshotList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1MarketDataSnapshotGetSnapshotsResponse">ActiveV1MarketDataSnapshotGetSnapshotsResponse</a>

Methods:

- <code title="get /active/v1/market-data/snapshot">client.Active.V1.MarketData.Snapshot.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1MarketDataSnapshotService.GetSnapshots">GetSnapshots</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1MarketDataSnapshotGetSnapshotsParams">ActiveV1MarketDataSnapshotGetSnapshotsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1MarketDataSnapshotGetSnapshotsResponse">ActiveV1MarketDataSnapshotGetSnapshotsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### News

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#NewsItem">NewsItem</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#NewsItemList">NewsItemList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1NewsGetNewsResponse">ActiveV1NewsGetNewsResponse</a>

Methods:

- <code title="get /active/v1/news">client.Active.V1.News.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1NewsService.GetNews">GetNews</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1NewsGetNewsParams">ActiveV1NewsGetNewsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1NewsGetNewsResponse">ActiveV1NewsGetNewsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Screener

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ScreenerItem">ScreenerItem</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ScreenerItemList">ScreenerItemList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1ScreenerGetScreenerResponse">ActiveV1ScreenerGetScreenerResponse</a>

Methods:

- <code title="get /active/v1/screener">client.Active.V1.Screener.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1ScreenerService.GetScreener">GetScreener</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1ScreenerGetScreenerParams">ActiveV1ScreenerGetScreenerParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1ScreenerGetScreenerResponse">ActiveV1ScreenerGetScreenerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Version

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#Version">Version</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1VersionGetVersionResponse">ActiveV1VersionGetVersionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1VersionUpdateVersionResponse">ActiveV1VersionUpdateVersionResponse</a>

Methods:

- <code title="get /active/v1/version">client.Active.V1.Version.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1VersionService.GetVersion">GetVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1VersionGetVersionResponse">ActiveV1VersionGetVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /active/v1/version">client.Active.V1.Version.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1VersionService.UpdateVersion">UpdateVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1VersionUpdateVersionResponse">ActiveV1VersionUpdateVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Watchlists

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#WatchlistDetail">WatchlistDetail</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#WatchlistEntry">WatchlistEntry</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#WatchlistEntryList">WatchlistEntryList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#WatchlistItemEntry">WatchlistItemEntry</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistNewWatchlistResponse">ActiveV1WatchlistNewWatchlistResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistGetWatchlistByIDResponse">ActiveV1WatchlistGetWatchlistByIDResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistGetWatchlistsResponse">ActiveV1WatchlistGetWatchlistsResponse</a>

Methods:

- <code title="post /active/v1/watchlists">client.Active.V1.Watchlists.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistService.NewWatchlist">NewWatchlist</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistNewWatchlistParams">ActiveV1WatchlistNewWatchlistParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistNewWatchlistResponse">ActiveV1WatchlistNewWatchlistResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/watchlists/{watchlist_id}">client.Active.V1.Watchlists.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistService.DeleteWatchlist">DeleteWatchlist</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /active/v1/watchlists/{watchlist_id}">client.Active.V1.Watchlists.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistService.GetWatchlistByID">GetWatchlistByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistGetWatchlistByIDResponse">ActiveV1WatchlistGetWatchlistByIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /active/v1/watchlists">client.Active.V1.Watchlists.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistService.GetWatchlists">GetWatchlists</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistGetWatchlistsResponse">ActiveV1WatchlistGetWatchlistsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Items

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#AddWatchlistItemData">AddWatchlistItemData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistItemAddWatchlistItemResponse">ActiveV1WatchlistItemAddWatchlistItemResponse</a>

Methods:

- <code title="post /active/v1/watchlists/{watchlist_id}/items">client.Active.V1.Watchlists.Items.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistItemService.AddWatchlistItem">AddWatchlistItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, watchlistID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistItemAddWatchlistItemParams">ActiveV1WatchlistItemAddWatchlistItemParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistItemAddWatchlistItemResponse">ActiveV1WatchlistItemAddWatchlistItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /active/v1/watchlists/{watchlist_id}/items/{item_id}">client.Active.V1.Watchlists.Items.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistItemService.DeleteWatchlistItem">DeleteWatchlistItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go">clearstreet</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WatchlistItemDeleteWatchlistItemParams">ActiveV1WatchlistItemDeleteWatchlistItemParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Ws

Methods:

- <code title="get /active/v1/ws">client.Active.V1.Ws.<a href="https://pkg.go.dev/github.com/stainless-sdks/clear-street-go#ActiveV1WService.WebsocketHandler">WebsocketHandler</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
