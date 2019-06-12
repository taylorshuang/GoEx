package lbank

import (
	"fmt"
	"net/http"

	. "github.com/taylorshuang/GoEx"
)

const (
	EXCHANGE_NAME = "lbank.com"

	BASE_URL  = "https://api.lbkex.com/v1"
	DEPTH_URI = "/market/depth"
)

// Lbank represent a Lbank client.
type Lbank struct {
	accessKey,
	secretKey string
	httpClient *http.Client
}

func New(client *http.Client, accessKey, secretKey string) *Lbank {
	return &Lbank{accessKey, secretKey, client}

}

func (l *Lbank) GetExchangeName() string {
	return EXCHANGE_NAME
}

//
func (lbank *Lbank) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
	merge := 1
	apiUrl := BASE_URL + DEPTH_URI + fmt.Sprintf("?symbol=%s&size=%d&merge=%d", currency.ToLower(), size, merge)

	resp, err := HttpGet(lbank.httpClient, apiUrl)
	if err != nil {
		return nil, err
	}
	println("resp:", resp)

	depth := new(Depth)

	// for _, bid := range bids {
	// 	_bid := bid.(map[string]interface{})
	// 	amount := ToFloat64(_bid["amount"])
	// 	price := ToFloat64(_bid["price"])
	// 	dr := DepthRecord{Amount: amount, Price: price}
	// 	depth.BidList = append(depth.BidList, dr)
	// }

	// for _, ask := range asks {
	// 	_ask := ask.(map[string]interface{})
	// 	amount := ToFloat64(_ask["amount"])
	// 	price := ToFloat64(_ask["price"])
	// 	dr := DepthRecord{Amount: amount, Price: price}
	// 	depth.AskList = append(depth.AskList, dr)
	// }

	return depth, nil
}
