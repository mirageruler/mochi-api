package defi

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Handler struct {
	entities *entities.Entity
	log      logger.Logger
}

func New(entities *entities.Entity, logger logger.Logger) IHandler {
	return &Handler{
		entities: entities,
		log:      logger,
	}
}

// GetHistoricalMarketChart     godoc
// @Summary     Get historical market chart
// @Description Get historical market chart
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       coin_id   query  string true  "Coin ID"
// @Param       day   query  int true  "Day"
// @Param       currency   query  string false  "Currency" default(usd)
// @Success     200 {object} response.GetHistoricalMarketChartResponse
// @Router      /defi/market-chart [get]
func (h *Handler) GetHistoricalMarketChart(c *gin.Context) {
	var req request.GetMarketChartRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetHistoricalMarketChart] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err, statusCode := h.entities.GetHistoricalMarketChart(&req)
	if err != nil {
		h.log.Error(err, "[handler.GetHistoricalMarketChart] - failed to get historical market chart")
		c.JSON(statusCode, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetSupportedToken     godoc
// @Summary     Get supported token by address and chain id
// @Description Get supported token by address and chain id
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       address   query  string true  "token address"
// @Param       chain   query  string true  "token chain"
// @Success     200 {object} response.GetSupportedTokenResponse
// @Router      /defi/token [get]
func (h *Handler) GetSupportedToken(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		h.log.Info("[handler.GetSupportedToken] - address is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "address is required"})
		return
	}
	chain := c.Query("chain")
	if chain == "" {
		h.log.Info("[handler.GetSupportedToken] - chain is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "chain is required"})
		return
	}
	token, err := h.entities.GetSupportedToken(address, chain)
	if err != nil {
		h.log.Error(err, "[handler.GetSupportedToken] - failed to get supported token")
		c.JSON(baseerrs.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(token, nil, nil, nil))
}

// GetSupportedTokens     godoc
// @Summary     Get supported tokens
// @Description Get supported tokens
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetSupportedTokensResponse
// @Router      /defi/tokens [get]
func (h *Handler) GetSupportedTokens(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	if page == "" {
		page = "0"
	}
	if size == "" {
		size = "15"
	}
	tokens, pagination, err := h.entities.GetSupportedTokens(page, size)
	if err != nil {
		h.log.Error(err, "[handler.GetSupportedTokens] - failed to get supported tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tokens, pagination, nil, nil))
}

// GetCoin     godoc
// @Summary     Get coin
// @Description Get coin
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       id   path  string true  "Coin ID"
// @Success     200 {object} response.GetCoinResponseWrapper
// @Router      /defi/coins/{id} [get]
func (h *Handler) GetCoin(c *gin.Context) {
	coinID := c.Param("id")
	if coinID == "" {
		h.log.Info("[handler.GetCoin] - coin id missing")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}

	isDominanceChart := strings.EqualFold(c.Query("is_dominance_chart"), "true")

	data, err, statusCode := h.entities.GetCoinData(coinID, isDominanceChart)
	if err != nil {
		h.log.Error(err, "[handler.GetCoin] - failed to get coin data")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// SearchCoins     godoc
// @Summary     Search coin
// @Description Search coin
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       query   query  string true  "coin query"
// @Success     200 {object} response.SearchCoinResponse
// @Router      /defi/coins [get]
func (h *Handler) SearchCoins(c *gin.Context) {
	req := request.SearchCoinRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.SearchCoins] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	tokens, err := h.entities.SearchCoins(req.Query, req.GuildId, req.NoDefault)
	if err != nil {
		h.log.Error(err, "[handler.SearchCoins] entities.SearchCoins() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tokens, nil, nil, nil))
}

// GetTokenInfo		 godoc
// @Summary     Get token info
// @Description Get token info
// @Tags        Defi
// @Tags        Public
// @Accept      json
// @Produce     json
// @Param       query   query  string true  "token query"
// @Success     200 {object} response.TokenInfoResponse
// @Router      /defi/tokens/info/{id} [get]
func (h *Handler) GetTokenInfo(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		h.log.Info("[handler.GetTokenInfo] query is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}

	token, err := h.entities.GetTokenInfo(id)
	if err != nil {
		h.log.Error(err, "[handler.GetTokenInfo] entities.GetTokenInfo() failed")
		c.JSON(baseerrs.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(token, nil, nil, nil))
}

// CompareToken     godoc
// @Summary     Compare token
// @Description Compare token
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       base   query  string true  "base token"
// @Param       target   query  string true  "target token"
// @Param       interval   query  string true  "compare interval"
// @Param       guild_id   query  string false  "Guild ID"
// @Success     200 {object} response.CompareTokenResponse
// @Router      /defi/coins/compare [get]
func (h *Handler) CompareToken(c *gin.Context) {
	base := c.Query("base")
	target := c.Query("target")
	interval := c.Query("interval")
	guildID := c.Query("guild_id")

	if base == "" {
		h.log.Info("[handler.CompareToken] base is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "source symbol is required"})
		return
	}

	if target == "" {
		h.log.Info("[handler.CompareToken] target is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "target symbol is required"})
		return
	}
	if interval == "" {
		h.log.Info("[handler.CompareToken] interval empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "interval is required"})
		return
	}

	res, err := h.entities.CompareToken(base, target, interval, guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"base": base, "target": target}).Error(err, "[handler.CompareToken] entity.CompareToken failed")
		c.JSON(baseerrs.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// GetFiatHistoricalExchangeRates     godoc
// @Summary     Get historical market chart
// @Description Remove from user's watchlist
// @Tags        Fiat
// @Accept      json
// @Produce     json
// @Param       req query request.GetFiatHistoricalExchangeRatesRequest true "request"
// @Success     200 {object} response.GetFiatHistoricalExchangeRatesResponse
// @Router      /fiat/historical-exchange-rates [get]
func (h *Handler) GetFiatHistoricalExchangeRates(c *gin.Context) {
	var req request.GetFiatHistoricalExchangeRatesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetFiatHistoricalExchangeRates] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.entities.GetFiatHistoricalExchangeRates(req)
	if err != nil {
		h.log.Error(err, "[handler.GetFiatHistoricalExchangeRates] entity.GetFiatHistoricalExchangeRates() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// AddContract   godoc
// @Summary     List All Chain
// @Description List All Chain
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetListAllChainsResponse
// @Router      /defi/chains [get]
func (h *Handler) ListAllChain(c *gin.Context) {
	returnChain, err := h.entities.ListAllChain()
	if err != nil {
		h.log.Error(err, "[handler.ListAllChain] - failed to list all chains")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(returnChain, nil, nil, nil))
}

// AddToWatchlist     godoc
// @Summary     Add to user's price alert
// @Description Add to user's price alert
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       req body request.AddTokenPriceAlertRequest true "request"
// @Success     200 {object} response.AddTokenPriceAlertResponse
// @Router      /defi/price-alert [post]
func (h *Handler) AddTokenPriceAlert(c *gin.Context) {
	var req request.AddTokenPriceAlertRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.AddTokenPriceAlert] Bind() failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.entities.AddTokenPriceAlert(req)
	if err != nil {
		h.log.Error(err, "[handler.AddTokenPriceAlert] entity.AddTokenPriceAlert() failed")
		c.JSON(baseerrs.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetUserListPriceAlert     godoc
// @Summary     Get user's price alerts
// @Description Get user's price alerts
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       req query request.GetUserListPriceAlertRequest true "request"
// @Success     200 {object} response.ListTokenPriceAlertResponse
// @Router      /defi/price-alert [get]
func (h *Handler) GetUserListPriceAlert(c *gin.Context) {
	var req request.GetUserListPriceAlertRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetUserListPriceAlert] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.entities.GetUserListPriceAlert(req)
	if err != nil {
		h.log.Error(err, "[handler.GetUserListPriceAlert] entity.GetUserWatchlist() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// RemoveTokenPriceAlert     godoc
// @Summary     Remove from user's price alerts
// @Description Remove from user's price alerts
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       id query string true "id"
// @Success     200 {object} object
// @Router      /defi/price-alert [delete]
func (h *Handler) RemoveTokenPriceAlert(c *gin.Context) {
	alertID := c.Query("id")
	if alertID == "" {
		h.log.Info("[handler.RemoveTokenPriceAlert] - id is required")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}
	err := h.entities.RemoveTokenPriceAlert(alertID)
	if err != nil {
		h.log.Error(err, "[handler.RemoveTokenPriceAlert] entity.RemoveTokenPriceAlert() failed")
		code := http.StatusInternalServerError
		if err == baseerrs.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, nil, nil))
}

// GetCoin     godoc
// @Summary     Get coin data from Binance Exchange
// @Description Get coin data from Binance Exchange
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Coin ID"
// @Success     200 {object} response.GetCoinResponseWrapper
// @Router      /defi/coins/binance/{symbol} [get]
func (h *Handler) GetBinanceCoinData(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		h.log.Info("[handler.GetBinanceCoinData] - symbol missing")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}

	data, err, statusCode := h.entities.GetBinanceCoinPrice(symbol)
	if err != nil {
		h.log.Error(err, "[handler.GetBinanceCoinData] - failed to get coin data")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetUserRequestTokens     godoc
// @Summary     Get tokens requested by user
// @Description Get tokens requested by user
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetSupportedTokensResponse
// @Router      /defi/token-support [get]
func (h *Handler) GetUserRequestTokens(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	status := c.Query("status")
	if page == "" {
		page = "0"
	}
	if size == "" {
		size = "15"
	}
	tokens, pagination, err := h.entities.GetUserRequestTokens(request.GetUserSupportTokenRequest{Page: page, Size: size, Status: status})
	if err != nil {
		h.log.Error(err, "[handler.GetSupportedTokens] - failed to get supported tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tokens, pagination, nil, nil))
}

// CreateUserTokenSupportRequest     godoc
// @Summary     Request support token
// @Description Request support token
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       Request body  request.CreateUserTokenSupportRequest true  "Create user token support request"
// @Success     200 {object} response.CreateUserTokenSupportRequest
// @Router      /defi/token-support [post]
func (h *Handler) CreateUserTokenSupportRequest(c *gin.Context) {
	req := &request.CreateUserTokenSupportRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateUserTokenSupportRequest] - c.ShouldBindJSON failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res, err := h.entities.CreateUserTokenSupportRequest(*req)
	if err != nil {
		h.log.Error(err, "[handler.CreateUserTokenSupportRequest] - entities.CreateUserTokenSupportRequest failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// FindTokenByContractAddress godoc
// @Summary     Find token by contract address
// @Description Find token by contract address
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       chain_id query string true  "Chain ID"
// @Param       address query string true  "Contract address"
// @Success     200 {object} response.FindTokenByContractAddressResponse
// @Router      /defi/custom-tokens [get]
func (h *Handler) FindTokenByContractAddress(c *gin.Context) {
	req := &request.FindTokenByContractAddressRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.FindTokenByContractAddress] - c.ShouldBindQuery failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.FindTokenByContractAddress(req.ChainId, req.Address)
	if err != nil {
		h.log.Error(err, "[handler.FindTokenByContractAddress] - entities.FindTokenByContractAddress failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// ApproveUserTokenSupportRequest     godoc
// @Summary     Approve support token request
// @Description Approve support token request
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       id path int true  "Support Token Request ID"
// @Success     200 {object} response.CreateUserTokenSupportRequest
// @Router      /defi/token-support/{id}/approve [put]
func (h *Handler) ApproveUserTokenSupportRequest(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.ApproveUserTokenSupportRequest] - invalid id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf("invalid request id"), nil))
		return
	}
	res, err := h.entities.ApproveTokenSupportRequest(id)
	if err != nil {
		h.log.Error(err, "[handler.ApproveUserTokenSupportRequest] - entities.ApproveTokenSupportRequest() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// RejectUserTokenSupportRequest     godoc
// @Summary     Reject support token request
// @Description Reject support token request
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       id path  int true  "Support Token Request ID"
// @Success     200 {object} response.CreateUserTokenSupportRequest
// @Router      /defi/token-support/{id}/reject [put]
func (h *Handler) RejectUserTokenSupportRequest(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.RejectTokenSupportRequest] - invalid id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf("invalid request id"), nil))
		return
	}
	res, err := h.entities.RejectTokenSupportRequest(id)
	if err != nil {
		h.log.Error(err, "[handler.RejectTokenSupportRequest] - entities.RejectTokenSupportRequest() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// GetGasTracker     godoc
// @Summary     Get gas tracker of all chain
// @Description Get gas tracker of all chain
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GasTrackerResponseData
// @Router      /defi/gas-tracker [get]
func (h *Handler) GetGasTracker(c *gin.Context) {
	gasTracker, err := h.entities.GetGasTracker()
	if err != nil {
		h.log.Error(err, "[handler.GetGasTracker] - entities.GetGasTracker() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(gasTracker, nil, nil, nil))
}

// GetChainGasTracker     godoc
// @Summary     Get gas tracker of one chain
// @Description Get gas tracker of one chain
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       chain   path  string true  "chain"
// @Success     200 {object} response.ChainGasTrackerResponseData
// @Router      /defi/gas-tracker/{chain} [get]
func (h *Handler) GetChainGasTracker(c *gin.Context) {
	chain := c.Param("chain")
	if chain == "" {
		h.log.Error(fmt.Errorf("chain is empty"), "[handler.GetChainGasTracker] - chain is empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf("chain is required"), nil))
		return
	}

	gasTracker, err := h.entities.GetChainGasTracker(chain)
	if err != nil {
		h.log.Error(err, "[handler.GetGasTracker] - entities.GetGasTracker() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(gasTracker, nil, nil, nil))
}

// GetCoinsMarketData     godoc
// @Summary     Get coins market data of top coins
// @Description Get coins market data of top coins
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       page   query  string false  "page"
// @Param       page_size   query  string false  "page_size"
// @Param       order   query  string false  "accepted values: price_change_percentage_7d_asc, price_change_percentage_7d_desc, price_change_percentage_1h_asc, price_change_percentage_1h_desc, price_change_percentage_24h_asc, price_change_percentage_24h_desc"
// @Success     200 {object} response.GetCoinsMarketDataResponse
// @Router      /defi/market-data [get]
func (h *Handler) GetCoinsMarketData(c *gin.Context) {
	req := request.GetMarketDataRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "c.ShouldBindQuery() - cannot parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf(err.Error()), nil))
		return
	}

	data, err := h.entities.GetCoinsMarketData(req)
	if err != nil {
		h.log.Error(err, "[handler.GetCoinsMarketData] entity.GetCoinsMarketData() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetTrendingSearch     godoc
// @Summary     Get trending search of coins
// @Description Get trending search of coins
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetTrendingSearch
// @Router      /defi/trending [get]
func (h *Handler) GetTrendingSearch(c *gin.Context) {
	data, err := h.entities.GetTrendingSearch()
	if err != nil {
		h.log.Error(err, "[handler.GetTrendingCoins] entity.GetTrendingCoins() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// TopGainerLoser     godoc
// @Summary     Get top 300 gainer and loser
// @Description Get top 300 gainer and loser
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       duration   query  string false  "default: 24h, accepted value: 1h, 24h, 7d, 14d, 30d, 60d, 1y"
// @Success     200 {object} response.GetTopGainerLoser
// @Router      /defi/top-gainer-loser [get]
func (h *Handler) TopGainerLoser(c *gin.Context) {
	req := request.TopGainerLoserRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "c.ShouldBindQuery() - cannot parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf(err.Error()), nil))
		return
	}

	data, err := h.entities.GetTopLoserGainer(req)
	if err != nil {
		h.log.Error(err, "[handler.TopGainerLoser] entity.GetTopLoserGainer() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// SearchKeys     godoc
// @Summary     Search coin
// @Description Search coin
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       query   query  string true  "coin query"
// @Success     200 {object} response.FriendTechKeysResponse
// @Router      /defi/keys [get]
func (h *Handler) SearchKeys(c *gin.Context) {
	req := request.SearchFriendTechKeysRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.SearchKeys] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	tokens, err := h.entities.SearchFriendTechKeys(request.SearchFriendTechKeysRequest{Query: req.Query, Limit: req.Limit})
	if err != nil {
		h.log.Error(err, "[handler.SearchKeys] entities.SearchFriendTechKeys() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tokens.Data, nil, nil, nil))
}

// TrackFriendTechKey     godoc
// @Summary     Track a specific friend tech key by adding to user's watchlist
// @Description Track a specific friend tech key by adding to user's watchlist
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       req body request.TrackFriendTechKeyRequest true "request"
// @Success     200 {object} response.TrackFriendTechKeyResponse
// @Router      /defi/tracking-keys [post]
func (h *Handler) TrackFriendTechKey(c *gin.Context) {
	req := request.TrackFriendTechKeyRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.TrackFriendTechKey] - c.ShouldBindQuery() - cannot parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf(err.Error()), nil))
		return
	}

	data, err := h.entities.TrackFriendTechKey(req.ProfileId, req.KeyAddress, req.IncreaseAlertAt, req.DecreaseAlertAt)
	if err != nil {
		h.log.Error(err, "[handler.TrackFriendTechKey] entity.TrackFriendTechKey() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	searchKeyResult, err := h.entities.SearchFriendTechKeys(request.SearchFriendTechKeysRequest{
		Query: data.KeyAddress,
		Limit: 1,
	})
	if err != nil {
		h.log.Error(err, "[handler.TrackFriendTechKey] entities.SearchFriendTechKeys() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var keyMetadata *response.FriendTechKey
	if len(searchKeyResult.Data) > 0 && strings.EqualFold(searchKeyResult.Data[0].Address, data.KeyAddress) {
		keyMetadata = &searchKeyResult.Data[0]
	}

	if keyMetadata != nil {
		keyMetadata.PriceChangePercentage, err = h.entities.CalculateFriendTechKeyPriceChangePercentage(data.KeyAddress)
		if err != nil {
			h.log.Error(err, "[handler.TrackFriendTechKey] entities.CalculateFriendTechKeyPriceChangePercentage() failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.TrackingFriendTechKeyModelToResponse(
		*data,
		keyMetadata,
	), nil, nil, nil))
}

// UntrackFriendTechKey     godoc
// @Summary     Untrack a specific friend tech key by removing from user's watchlist
// @Description Untrack a specific friend tech key by removing from user's watchlist
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       id path int true "id"
// @Success     200 {string} string "ok"
// @Router      /defi/tracking-keys/{id} [delete]
func (h *Handler) UntrackFriendTechKey(c *gin.Context) {
	// id from path
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.UntrackFriendTechKey] - invalid id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf("invalid request id"), nil))
		return
	}

	if err = h.entities.UnTrackFriendTechKey(id); err != nil {
		h.log.Error(err, "[handler.UntrackFriendTechKey] entity.UnTrackFriendTechKey() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse("ok", nil, nil, nil))
}

// UpdateFriendTechKeyTrack     godoc
// @Summary     Update friend tech key track config
// @Description Update friend tech key track config
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       id path int true "id"
// @Param       req body request.UpdateFriendTechKeyTrackRequest true "request"
// @Success     200 {object} response.TrackFriendTechKeyResponse
// @Router      /defi/tracking-keys/{id} [put]
func (h *Handler) UpdateFriendTechKeyTrack(c *gin.Context) {
	// id from path
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.UpdateFriendTechKeyTrack] - invalid id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf("invalid request id"), nil))
		return
	}

	req := request.UpdateFriendTechKeyTrackRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UpdateFriendTechKeyTrack] c.ShouldBindJSON() - cannot parse query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf(err.Error()), nil))
		return
	}

	updatedTrack, err := h.entities.UpdateFriendTechKeyTrack(id, req.IncreaseAlertAt, req.DecreaseAlertAt)
	if err != nil {
		h.log.Error(err, "[handler.UpdateFriendTechKeyTrack] entity.UpdateFriendTechKeyTrack() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	searchKeyResult, err := h.entities.SearchFriendTechKeys(request.SearchFriendTechKeysRequest{
		Query: updatedTrack.KeyAddress,
		Limit: 1,
	})
	if err != nil {
		h.log.Error(err, "[handler.TrackFriendTechKey] entities.SearchFriendTechKeys() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var keyMetadata *response.FriendTechKey
	if len(searchKeyResult.Data) > 0 && strings.EqualFold(searchKeyResult.Data[0].Address, updatedTrack.KeyAddress) {
		keyMetadata = &searchKeyResult.Data[0]
	}

	if keyMetadata != nil {
		keyMetadata.PriceChangePercentage, err = h.entities.CalculateFriendTechKeyPriceChangePercentage(updatedTrack.KeyAddress)
		if err != nil {
			h.log.Error(err, "[handler.UpdateFriendTechKeyTrack] entities.CalculateFriendTechKeyPriceChangePercentage() failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.TrackingFriendTechKeyModelToResponse(
		*updatedTrack,
		keyMetadata,
	), nil, nil, nil))
}

// GetUserFriendTechKeyWatchlist     godoc
// @Summary     Get user's friend tech key watchlist
// @Description Get user's friend tech key watchlist
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       profile_id query string true "profile_id"
// @Success     200 {object} response.TrackFriendTechKeyResponse
// @Router      /defi/tracking-keys [get]
func (h *Handler) GetUserFriendTechKeyWatchlist(c *gin.Context) {
	// user profile id from query param
	profileIdStr := c.Query("profile_id")
	if profileIdStr == "" {
		h.log.Fields(logger.Fields{"profile_id": profileIdStr}).Error(fmt.Errorf("profile_id is empty"), "[handler.GetUserFriendTechKeyWatchlist] - invalid profile_id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, fmt.Errorf("invalid profile_id"), nil))
		return
	}

	watchlist, err := h.entities.GetUserFriendTechKeyWatchlist(profileIdStr)
	if err != nil {
		h.log.Error(err, "[handler.GetUserFriendTechKeyWatchlist] entity.GetUserFriendTechKeyWatchlist() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	resp := make([]response.FriendTechKeyWatchlistItemResponse, 0)
	for _, trackingKey := range watchlist {
		searchKeyResult, err := h.entities.SearchFriendTechKeys(request.SearchFriendTechKeysRequest{
			Query: trackingKey.KeyAddress,
			Limit: 1,
		})
		if err != nil {
			h.log.Error(err, "[handler.TrackFriendTechKey] entities.SearchFriendTechKeys() failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var keyMetadata *response.FriendTechKey
		if len(searchKeyResult.Data) > 0 && strings.EqualFold(searchKeyResult.Data[0].Address, trackingKey.KeyAddress) {
			keyMetadata = &searchKeyResult.Data[0]
		}

		if keyMetadata != nil {
			keyMetadata.PriceChangePercentage, err = h.entities.CalculateFriendTechKeyPriceChangePercentage(trackingKey.KeyAddress)
			if err != nil {
				h.log.Error(err, "[handler.TrackFriendTechKey] entities.CalculateFriendTechKeyPriceChangePercentage() failed")
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		resp = append(resp, *response.TrackingFriendTechKeyModelToResponse(trackingKey, keyMetadata))
	}

	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}
