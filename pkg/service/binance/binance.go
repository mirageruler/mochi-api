package binance

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Binance struct {
	getExchangeInfoURL string
	getSymbolKlinesURL string
}

func NewService() Service {
	return &Binance{
		getExchangeInfoURL: "https://api.binance.com/api/v3/exchangeInfo",
		getSymbolKlinesURL: "https://api.binance.com/api/v3/uiKlines?symbol=%s&interval=1h&limit=168", // 168h = 7d
	}
}

func (b *Binance) GetExchangeInfo(symbol string) (*response.GetExchangeInfoResponse, error, int) {
	symbol = strings.Replace(symbol, "/", "", 1)
	res := &response.GetExchangeInfoResponse{}
	url := b.getExchangeInfoURL
	if symbol != "" {
		url = fmt.Sprintf("%s?symbol=%s", url, symbol)
	}
	statusCode, err := util.FetchData(url, res)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("binance.GetExchangeInfo() failed: %v", err), statusCode
	}
	return res, nil, http.StatusOK
}

func (b *Binance) GetKlinesBySymbol(symbol string) ([]response.GetKlinesDataResponse, error, int) {
	symbol = strings.Replace(symbol, "/", "", 1)
	data := make([][]interface{}, 0)
	statusCode, err := util.FetchData(fmt.Sprintf(b.getSymbolKlinesURL, symbol), &data)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("binance.GetKlinesBySymbol() failed: %v", err), statusCode
	}
	res := make([]response.GetKlinesDataResponse, 0, len(data))
	for _, item := range data {
		res = append(res, response.GetKlinesDataResponse{
			OPrice: item[1].(string),
			HPrice: item[2].(string),
			LPrice: item[3].(string),
			CPrice: item[4].(string),
		})
	}
	return res, nil, http.StatusOK
}