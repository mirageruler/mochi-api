package usernftwatchlistitem

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q UserNftWatchlistQuery) ([]model.UserNftWatchlistItem, int64, error) {
	var items []model.UserNftWatchlistItem
	var total int64
	db := pg.db.Table("user_nft_watchlist_items")
	if q.UserID != "" {
		db = db.Where("user_id = ?", q.UserID)
	}
	if q.Symbol != "" {
		db = db.Where("symbol ILIKE ?", q.Symbol)
	}
	if q.CollectionAddress != "" {
		db = db.Where("collection_address = ?", q.CollectionAddress)
	}
	if q.ChainID != "" {
		db = db.Where("chain_id = ?", q.ChainID)
	}

	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return items, total, db.Find(&items).Error
}

func (pg *pg) Create(item *model.UserNftWatchlistItem) error {
	return pg.db.Create(item).Error
}

func (pg *pg) Delete(userID, symbol string) (int64, error) {
	tx := pg.db.Where("user_id = ? AND symbol ILIKE ?", userID, symbol).Delete(&model.UserNftWatchlistItem{})
	return tx.RowsAffected, tx.Error
}