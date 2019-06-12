package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taylorshuang/GoEx"
)

var builder = NewAPIBuilder()

func TestAPIBuilder_Build(t *testing.T) {
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(goex.OKCOIN_COM).GetExchangeName(), goex.OKCOIN_COM)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(goex.HUOBI_PRO).GetExchangeName(), goex.HUOBI_PRO)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(goex.ZB).GetExchangeName(), goex.ZB)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(goex.BIGONE).GetExchangeName(), goex.BIGONE)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(goex.OKEX).GetExchangeName(), goex.OKEX)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(goex.POLONIEX).GetExchangeName(), goex.POLONIEX)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(goex.KRAKEN).GetExchangeName(), goex.KRAKEN)
	assert.Equal(t, builder.APIKey("").APISecretkey("").BuildFuture(goex.HBDM).GetExchangeName(), goex.HBDM)
	assert.Equal(t, builder.APIKey("").APISecretkey("").BuildFuture(goex.LBANK).GetExchangeName(), goex.LBANK)
}
