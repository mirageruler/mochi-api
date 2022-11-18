package nghenhan

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Nghenhan struct {
	baseUrl string
}

func NewService() Service {
	return &Nghenhan{
		baseUrl: "https://cex.console.so/api/v1",
	}
}

func (n *Nghenhan) GetFiatHistoricalChart(base, target, interval string, limit int) (*response.NghenhanFiatHistoricalChartResponse, error) {
	url := n.baseUrl + fmt.Sprintf("/rate?base=%s&target=%s&interval=%s&limit=%v", base, target, interval, limit)
	data := response.NghenhanFiatHistoricalChartResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   map[string]string{"Content-Type": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[nghenhan.GetFiatHistoricalChart] util.SendRequest() failed: %v", err)
	}
	return &data, nil
}
