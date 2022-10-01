package offchain_tip_bot_chain

import (
	"strings"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetAll(f Filter) ([]model.OffchainTipBotChain, error) {
	var rs []model.OffchainTipBotChain

	db := pg.db.
		Group("offchain_tip_bot_chains.id,offchain_tip_bot_chains.chain_id").
		Order("offchain_tip_bot_chains.chain_name").
		Preload("Tokens").
		Preload("Contracts").
		Joins(`
	JOIN offchain_tip_bot_tokens_chains tc ON tc.chain_id = offchain_tip_bot_chains.id 
	JOIN offchain_tip_bot_tokens t ON tc.token_id = t.id
	JOIN offchain_tip_bot_contracts c ON c.chain_id = offchain_tip_bot_chains.id`)

	switch {
	case f.TokenID != "":
		db = db.Where("t.token_id = ?", f.TokenID)
	case f.TokenSymbol != "":
		db = db.Where("lower(t.token_symbol) = ?", strings.ToLower(f.TokenSymbol))
	}

	if f.IsContractAvailable {
		db = db.Where("c.assign_status = ?", consts.OffchainTipBotContractAssignStatusAvailable)
	}

	return rs, db.Find(&rs).Error
}

func (pg *pg) GetByID(id string) (model.OffchainTipBotChain, error) {
	var rs model.OffchainTipBotChain
	return rs, pg.db.Where("chain_id = ?", id).First(&rs).Error
}
