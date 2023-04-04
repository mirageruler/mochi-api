package treasurer

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Create(treasurer *model.Treasurer) (*model.Treasurer, error) {
	return treasurer, pg.db.Create(treasurer).Error
}

func (pg *pg) GetByVaultId(vaultId int64) (treasurers []model.Treasurer, err error) {
	return treasurers, pg.db.Where("vault_id = ?", vaultId).Find(&treasurers).Error
}
