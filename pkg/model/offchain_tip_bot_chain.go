package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotChain struct {
	ID          uuid.UUID                 `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	ChainID     string                    `json:"chain_id"`
	ChainName   string                    `json:"chain_name"`
	Currency    string                    `json:"currency"`
	RPCURL      string                    `json:"rpc_url"`
	ExplorerURL string                    `json:"explorer_url"`
	Status      int                       `json:"status"`
	Tokens      []*OffchainTipBotToken    `json:"tokens" gorm:"many2many:offchain_tip_bot_tokens_chains;foreignKey:ID;joinForeignKey:ChainID;References:ID;joinReferences:TokenID"`
	Contracts   []*OffchainTipBotContract `json:"contracts" gorm:"foreignKey:ChainID"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
	DeletedAt   *time.Time                `json:"-"`
}

func (OffchainTipBotChain) TableName() string {
	return "offchain_tip_bot_chains"
}
