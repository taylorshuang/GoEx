package lbank

import (
	"net/http"
	"testing"

	"github.com/taylorshuang/GoEx"
)

var lbank = New(http.DefaultClient, "", "")

func TestGetDepth(t *testing.T) {
	dep, _ := lbank.GetDepth(2, goex.ETH_BTC)
	t.Log(dep.AskList)
	t.Log(dep.BidList)
}
