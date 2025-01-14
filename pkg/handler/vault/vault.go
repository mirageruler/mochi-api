package vault

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	vaulttxquery "github.com/defipod/mochi/pkg/repo/vault_transaction"
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

func (h *Handler) CreateVault(c *gin.Context) {
	var req request.CreateVaultRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateVault] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !validateVaultName(req.Name) {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(nil, "[validation.validateVaultName] - vault name exceed 24 characters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vault name exceed 24 characters"})
		return
	}

	vault, err := h.entities.CreateVault(&req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateVault] - failed to create vault")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Vault name is already exist for this server"})
			return
		}
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateVault] - failed to create vault")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vault, nil, nil, nil))
}

// GetVaults     godoc
// @Summary     Get vaults
// @Description Get vaults
// @Tags        Vault
// @Accept      json
// @Produce     json
// @Param       req   query  request.GetVaultsRequest true  "get vaults request"
// @Success     200 {object} response.GetVaultsResponse
// @Router      /vault [get]
func (h *Handler) GetVaults(c *gin.Context) {
	var req request.GetVaultsRequest
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetVaults] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}

	vault, err := h.entities.GetVaults(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetVaults] entity.GetVaults() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vault, nil, nil, nil))
}

// GetVault     godoc
// @Summary     Get vault
// @Description Get vault
// @Tags        Vault
// @Accept      json
// @Produce     json
// @Param       req   query  request.GetVaultRequest true  "get vault request"
// @Success     200 {object} response.GetVaultResponse
// @Router      /vault [get]
func (h *Handler) GetVault(c *gin.Context) {
	var req request.GetVaultRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.GetVault] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}

	vault, err := h.entities.GetVault(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetVault] entity.GetVault() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vault, nil, nil, nil))
}

func (h *Handler) GetVaultConfigChannel(c *gin.Context) {
	guildId := c.Query("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetVault] - guildId is empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}

	vaultInfo, err := h.entities.GetVaultConfigChannel(guildId)
	if err != nil {
		h.log.Error(err, "[handler.GetVaultConfigChannel] - failed to get vault config channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vaultInfo, nil, nil, nil))
}

func (h *Handler) CreateConfigChannel(c *gin.Context) {
	var req request.CreateVaultConfigChannelRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "channelID": req.ChannelId}).Error(err, "[handler.CreateConfigChannel] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.CreateVaultConfigChannel(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "channelID": req.ChannelId}).Error(err, "[handler.CreateVaultConfigChannel] - failed to create vault config channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) CreateConfigThreshold(c *gin.Context) {
	var req request.CreateConfigThresholdRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateConfigThreshold] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	vaultConfigChannel, err := h.entities.CreateConfigThreshold(&req)
	if err != nil {
		if err.Error() == "vault not found" {
			h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateConfigThreshold] - vault not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "This vault is not existed yet"})
			return
		}
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateConfigThreshold] - failed to create vault config channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vaultConfigChannel, nil, nil, nil))
}

func (h *Handler) CreateTreasurerRequest(c *gin.Context) {
	var req request.CreateTreasurerRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "userProfileId": req.UserProfileId, "vaultName": req.VaultName, "message": req.Message}).Error(err, "[handler.CreateAddTreasurerRequest] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	treasurerReq, err := h.entities.CreateTreasurerRequest(&req)
	if err != nil {
		if strings.Contains(err.Error(), "vault not exist") {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateAddTreasurerRequest] - user not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This vault is not existed yet"})
			return
		}

		if strings.Contains(err.Error(), "balance not enough") {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateAddTreasurerRequest] - balance not enough")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
			return
		}

		if strings.Contains(err.Error(), "user not in list treasurers") {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateAddTreasurerRequest] - user not in list treasurers")
			c.JSON(http.StatusBadRequest, gin.H{"error": "User is not member of this vault"})
			return
		}

		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateAddTreasurerRequest] - failed to create add treasurer req")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](treasurerReq, nil, nil, nil))
}

func (h *Handler) AddTreasurerToVault(c *gin.Context) {
	var req request.AddTreasurerToVaultRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.AddTreasurerToVault] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	treasurer, err := h.entities.AddTreasurerToVault(&req)
	if err != nil {
		if strings.Contains(err.Error(), "request not exist") {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.AddTreasurerToVault] - request not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This request is not exist"})
			return
		}
		if strings.Contains(err.Error(), "duplicate key value") {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.AddTreasurerToVault] - user already in vault")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This user is already added to this vault"})
			return
		}

		if strings.Contains(err.Error(), "vault not exist") {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.AddTreasurerToVault] - user not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This vault is not existed yet"})
			return
		}

		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.AddTreasurerToVault] - failed to add treasurer to vault")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](treasurer, nil, nil, nil))
}

func (h *Handler) RemoveTreasurerFromVault(c *gin.Context) {
	var req request.RemoveTreasurerToVaultRequest

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.RemoveTreasurerFromVault] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	treasurer, err := h.entities.RemoveTreasurerFromVault(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.RemoveTreasurerFromVault] - failed to remove treasurer from vault")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](treasurer, nil, nil, nil))
}

func (h *Handler) CreateTreasurerSubmission(c *gin.Context) {
	var req request.CreateTreasurerSubmission

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateTreasurerSubmission] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// TODO: check if submission is finish

	treasurerSubmission, err := h.entities.CreateTreasurerSubmission(&req)
	if err != nil {
		if strings.Contains(err.Error(), "submission already processed") {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateTreasurerSubmission] - user not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This submission is already processed"})
			return
		}

		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateTreasurerSubmission] - failed to create treasurer submission")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](treasurerSubmission, nil, nil, nil))
}

func (h *Handler) GetVaultDetail(c *gin.Context) {
	vaultName := c.Query("vaultName")
	guildId := c.Query("guildId")

	vaultDetail, err := h.entities.GetVaultDetail(vaultName, guildId)
	if err != nil {
		if strings.Contains(err.Error(), "vault not exist") {
			h.log.Fields(logger.Fields{"vaultName": vaultName, "guildId": guildId}).Error(err, "[handler.AddTreasurerToVault] - user not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This vault is not existed yet"})
			return
		}
		h.log.Fields(logger.Fields{"vaultName": vaultName, "guildId": guildId}).Error(err, "[handler.GetVaultDetail] - failed to get vault detail")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](vaultDetail, nil, nil, nil))
}

func (h *Handler) TransferVaultToken(c *gin.Context) {
	var req request.TransferVaultTokenRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.TransferVaultToken] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.TransferVaultToken(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.TransferVaultToken] - failed to transfer vault token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](gin.H{"message": "ok"}, nil, nil, nil))
}

func (h *Handler) GetTreasurerRequest(c *gin.Context) {
	requestId := c.Param("request_id")
	treasurerRequest, err := h.entities.GetTreasurerRequest(requestId)
	if err != nil {
		h.log.Fields(logger.Fields{"requestId": requestId}).Error(err, "[handler.GetTreasurerRequest] - failed to get treasurer request")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](treasurerRequest, nil, nil, nil))
}

func (h *Handler) GetVaultTransactions(c *gin.Context) {
	vaultId := c.Param("vault_id")
	vaultIdInt, err := strconv.Atoi(vaultId)
	if err != nil {
		h.log.Fields(logger.Fields{"vaultId": vaultId}).Error(err, "[handler.GetVaultTransactions] - failed to convert vault id to int")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	query := vaulttxquery.VaultTransactionQuery{
		StartTime: c.Query("start_time"),
		EndTime:   c.Query("end_time"),
		VaultId:   int64(vaultIdInt),
	}

	vaultTransactions, err := h.entities.GetVaultTransactions(query)
	if err != nil {
		h.log.Fields(logger.Fields{"vaultId": vaultId}).Error(err, "[handler.GetVaultTransactions] - failed to get vault transactions")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](vaultTransactions, nil, nil, nil))
}
