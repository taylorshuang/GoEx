package builder

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	. "github.com/taylorshuang/GoEx"
	"github.com/taylorshuang/GoEx/bigone"
	"github.com/taylorshuang/GoEx/binance"
	"github.com/taylorshuang/GoEx/bitfinex"
	"github.com/taylorshuang/GoEx/bithumb"
	"github.com/taylorshuang/GoEx/bitstamp"
	"github.com/taylorshuang/GoEx/bittrex"
	"github.com/taylorshuang/GoEx/coin58"
	"github.com/taylorshuang/GoEx/coinex"
	"github.com/taylorshuang/GoEx/fcoin"
	"github.com/taylorshuang/GoEx/gateio"
	"github.com/taylorshuang/GoEx/gdax"
	"github.com/taylorshuang/GoEx/hitbtc"
	"github.com/taylorshuang/GoEx/huobi"
	"github.com/taylorshuang/GoEx/kraken"
	"github.com/taylorshuang/GoEx/lbank"
	"github.com/taylorshuang/GoEx/okcoin"
	"github.com/taylorshuang/GoEx/okex"
	"github.com/taylorshuang/GoEx/poloniex"
	"github.com/taylorshuang/GoEx/zb"
)

type APIBuilder struct {
	client        *http.Client
	httpTimeout   time.Duration
	apiKey        string
	secretkey     string
	clientId      string
	apiPassphrase string
}

func NewAPIBuilder() (builder *APIBuilder) {
	_client := http.DefaultClient
	transport := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 4 * time.Second,
	}
	_client.Transport = transport
	return &APIBuilder{client: _client}
}

func NewCustomAPIBuilder(client *http.Client) (builder *APIBuilder) {
	return &APIBuilder{client: client}
}

func (builder *APIBuilder) APIKey(key string) (_builder *APIBuilder) {
	builder.apiKey = key
	return builder
}

func (builder *APIBuilder) APISecretkey(key string) (_builder *APIBuilder) {
	builder.secretkey = key
	return builder
}

func (builder *APIBuilder) HttpProxy(proxyUrl string) (_builder *APIBuilder) {
	if proxyUrl == "" {
		return builder
	}
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return builder
	}
	transport := builder.client.Transport.(*http.Transport)
	transport.Proxy = http.ProxyURL(proxy)
	return builder
}

func (builder *APIBuilder) ClientID(id string) (_builder *APIBuilder) {
	builder.clientId = id
	return builder
}

func (builder *APIBuilder) HttpTimeout(timeout time.Duration) (_builder *APIBuilder) {
	builder.httpTimeout = timeout
	builder.client.Timeout = timeout
	transport := builder.client.Transport.(*http.Transport)
	if transport != nil {
		transport.ResponseHeaderTimeout = timeout
		transport.TLSHandshakeTimeout = timeout
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		}
	}
	return builder
}

func (builder *APIBuilder) ApiPassphrase(apiPassphrase string) (_builder *APIBuilder) {
	builder.apiPassphrase = apiPassphrase
	return builder
}

func (builder *APIBuilder) Build(exName string) (api API) {
	var _api API
	switch exName {
	//case OKCOIN_CN:
	//	_api = okcoin.New(builder.client, builder.apiKey, builder.secretkey)
	case POLONIEX:
		_api = poloniex.New(builder.client, builder.apiKey, builder.secretkey)
	case OKCOIN_COM:
		_api = okcoin.NewCOM(builder.client, builder.apiKey, builder.secretkey)
	case BITSTAMP:
		_api = bitstamp.NewBitstamp(builder.client, builder.apiKey, builder.secretkey, builder.clientId)
	case HUOBI_PRO:
		_api = huobi.NewHuoBiProSpot(builder.client, builder.apiKey, builder.secretkey)
	case OKEX:
		_api = okcoin.NewOKExSpot(builder.client, builder.apiKey, builder.secretkey)
	case BITFINEX:
		_api = bitfinex.New(builder.client, builder.apiKey, builder.secretkey)
	case KRAKEN:
		_api = kraken.New(builder.client, builder.apiKey, builder.secretkey)
	case BINANCE:
		_api = binance.New(builder.client, builder.apiKey, builder.secretkey)
	case BITTREX:
		_api = bittrex.New(builder.client, builder.apiKey, builder.secretkey)
	case BITHUMB:
		_api = bithumb.New(builder.client, builder.apiKey, builder.secretkey)
	case GDAX:
		_api = gdax.New(builder.client, builder.apiKey, builder.secretkey)
	case GATEIO:
		_api = gateio.New(builder.client, builder.apiKey, builder.secretkey)
	case ZB:
		_api = zb.New(builder.client, builder.apiKey, builder.secretkey)
	case COINEX:
		_api = coinex.New(builder.client, builder.apiKey, builder.secretkey)
	case FCOIN:
		_api = fcoin.NewFCoin(builder.client, builder.apiKey, builder.secretkey)
	case COIN58:
		_api = coin58.New58Coin(builder.client, builder.apiKey, builder.secretkey)
	case BIGONE:
		_api = bigone.New(builder.client, builder.apiKey, builder.secretkey)
	case HITBTC:
		_api = hitbtc.New(builder.client, builder.apiKey, builder.secretkey)
	case LBANK:
		_api = lbank.New(builder.client, builder.apiKey, builder.secretkey)
	default:
		println("exchange name error [" + exName + "].")

	}
	return _api
}

func (builder *APIBuilder) BuildFuture(exName string) (api FutureRestAPI) {
	switch exName {
	case OKEX_FUTURE:
		return okcoin.NewOKEx(builder.client, builder.apiKey, builder.secretkey)
	case HBDM:
		return huobi.NewHbdm(&APIConfig{HttpClient: builder.client, ApiKey: builder.apiKey, ApiSecretKey: builder.secretkey})
	case OKEX_SWAP:
		return okex.NewOKExSwap(&APIConfig{HttpClient: builder.client, Endpoint: "https://www.okex.com", ApiKey: builder.apiKey, ApiSecretKey: builder.secretkey, ApiPassphrase: builder.apiPassphrase})
	default:
		println(fmt.Sprintf("%s not support", exName))
		return nil
	}
}
