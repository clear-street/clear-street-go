// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/stainless-sdks/clear-street-go/option"
)

// ActiveV1Service contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1Service] method instead.
type ActiveV1Service struct {
	Options     []option.RequestOption
	Accounts    ActiveV1AccountService
	APIKeys     ActiveV1APIKeyService
	Assistant   ActiveV1AssistantService
	Calendars   ActiveV1CalendarService
	Instruments ActiveV1InstrumentService
	Iris        ActiveV1IrisService
	MarketData  ActiveV1MarketDataService
	News        ActiveV1NewsService
	Screener    ActiveV1ScreenerService
	Version     ActiveV1VersionService
	Watchlists  ActiveV1WatchlistService
	Ws          ActiveV1WService
}

// NewActiveV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1Service(opts ...option.RequestOption) (r ActiveV1Service) {
	r = ActiveV1Service{}
	r.Options = opts
	r.Accounts = NewActiveV1AccountService(opts...)
	r.APIKeys = NewActiveV1APIKeyService(opts...)
	r.Assistant = NewActiveV1AssistantService(opts...)
	r.Calendars = NewActiveV1CalendarService(opts...)
	r.Instruments = NewActiveV1InstrumentService(opts...)
	r.Iris = NewActiveV1IrisService(opts...)
	r.MarketData = NewActiveV1MarketDataService(opts...)
	r.News = NewActiveV1NewsService(opts...)
	r.Screener = NewActiveV1ScreenerService(opts...)
	r.Version = NewActiveV1VersionService(opts...)
	r.Watchlists = NewActiveV1WatchlistService(opts...)
	r.Ws = NewActiveV1WService(opts...)
	return
}

// Security identifier source
type SecurityIDSource string

const (
	SecurityIDSourceCms                           SecurityIDSource = "CMS"
	SecurityIDSourceClst                          SecurityIDSource = "CLST"
	SecurityIDSourceOpra                          SecurityIDSource = "OPRA"
	SecurityIDSourceFigi                          SecurityIDSource = "FIGI"
	SecurityIDSourceCusip                         SecurityIDSource = "CUSIP"
	SecurityIDSourceCurrency                      SecurityIDSource = "CURRENCY"
	SecurityIDSourceFmp                           SecurityIDSource = "FMP"
	SecurityIDSourceOems                          SecurityIDSource = "OEMS"
	SecurityIDSourceSedol                         SecurityIDSource = "SEDOL"
	SecurityIDSourceQuik                          SecurityIDSource = "QUIK"
	SecurityIDSourceIsin                          SecurityIDSource = "ISIN"
	SecurityIDSourceRic                           SecurityIDSource = "RIC"
	SecurityIDSourceCountry                       SecurityIDSource = "COUNTRY"
	SecurityIDSourceExchange                      SecurityIDSource = "EXCHANGE"
	SecurityIDSourceCta                           SecurityIDSource = "CTA"
	SecurityIDSourceBloomberg                     SecurityIDSource = "BLOOMBERG"
	SecurityIDSourceWertpapier                    SecurityIDSource = "WERTPAPIER"
	SecurityIDSourceDutch                         SecurityIDSource = "DUTCH"
	SecurityIDSourceValoren                       SecurityIDSource = "VALOREN"
	SecurityIDSourceSicovam                       SecurityIDSource = "SICOVAM"
	SecurityIDSourceBelgian                       SecurityIDSource = "BELGIAN"
	SecurityIDSourceCommon                        SecurityIDSource = "COMMON"
	SecurityIDSourceClearingHouse                 SecurityIDSource = "CLEARING_HOUSE"
	SecurityIDSourceIsdaFpmlSpecification         SecurityIDSource = "ISDA_FPML_SPECIFICATION"
	SecurityIDSourceIsdaFpmlURL                   SecurityIDSource = "ISDA_FPML_URL"
	SecurityIDSourceLetterOfCredit                SecurityIDSource = "LETTER_OF_CREDIT"
	SecurityIDSourceMarketplaceAssignedIdentifier SecurityIDSource = "MARKETPLACE_ASSIGNED_IDENTIFIER"
	SecurityIDSourceMarkitRedEntityClip           SecurityIDSource = "MARKIT_RED_ENTITY_CLIP"
	SecurityIDSourceMarkitRedPairClip             SecurityIDSource = "MARKIT_RED_PAIR_CLIP"
	SecurityIDSourceCftc                          SecurityIDSource = "CFTC"
	SecurityIDSourceIsdaCommodityReferencePrice   SecurityIDSource = "ISDA_COMMODITY_REFERENCE_PRICE"
	SecurityIDSourceLegalEntityIdentifier         SecurityIDSource = "LEGAL_ENTITY_IDENTIFIER"
	SecurityIDSourceSynthetic                     SecurityIDSource = "SYNTHETIC"
	SecurityIDSourceFidessaInstrumentMnemonic     SecurityIDSource = "FIDESSA_INSTRUMENT_MNEMONIC"
	SecurityIDSourceIndexName                     SecurityIDSource = "INDEX_NAME"
	SecurityIDSourceUniformSymbol                 SecurityIDSource = "UNIFORM_SYMBOL"
	SecurityIDSourceDigitalTokenIdentifier        SecurityIDSource = "DIGITAL_TOKEN_IDENTIFIER"
	SecurityIDSourceMassive                       SecurityIDSource = "MASSIVE"
	SecurityIDSourceOther                         SecurityIDSource = "OTHER"
)

// Security type
type SecurityType string

const (
	SecurityTypeCommonStock    SecurityType = "COMMON_STOCK"
	SecurityTypePreferredStock SecurityType = "PREFERRED_STOCK"
	SecurityTypeCorporateBond  SecurityType = "CORPORATE_BOND"
	SecurityTypeOption         SecurityType = "OPTION"
	SecurityTypeFuture         SecurityType = "FUTURE"
	SecurityTypeWarrant        SecurityType = "WARRANT"
	SecurityTypeCash           SecurityType = "CASH"
	SecurityTypeOther          SecurityType = "OTHER"
)
