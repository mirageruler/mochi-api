package response

import (
	"time"

	friendtechtypes "github.com/consolelabs/mochi-typeset/common/transaction/friendtech/typeset"
	"github.com/defipod/mochi/pkg/model"
	"github.com/shopspring/decimal"
)

type FriendTechKeysResponse struct {
	Data []FriendTechKey `json:"data"`
}

// FriendTechKey camelCase, to handle response from FriendTech API
type FriendTechKey struct {
	ID                    int64     `json:"id"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	Address               string    `json:"address"`
	TwitterUsername       string    `json:"twitterUsername"`
	TwitterPfpUrl         string    `json:"twitterPfpUrl"`
	ProfileChecked        bool      `json:"profileChecked"`
	Price                 float64   `json:"price"`
	Supply                int       `json:"supply"`
	Holders               int       `json:"holders"`
	PriceChangePercentage float64   `json:"priceChangePercentage"`
}

// Also FriendTechKey but snake_case, to return to Mochi client
type TrackedFriendTechKey struct {
	ID                    int64     `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Address               string    `json:"address"`
	TwitterUsername       string    `json:"twitter_username"`
	TwitterPfpUrl         string    `json:"twitter_pfp_url"`
	ProfileChecked        bool      `json:"profile_checked"`
	Price                 float64   `json:"price"`
	Supply                int       `json:"supply"`
	Holders               int       `json:"holders"`
	PriceChangePercentage float64   `json:"price_change_percentage"`
}

func CamelCaseToSnakeCaseFriendTechKey(f *FriendTechKey) *TrackedFriendTechKey {
	return &TrackedFriendTechKey{
		ID:                    f.ID,
		CreatedAt:             f.CreatedAt,
		UpdatedAt:             f.UpdatedAt,
		Address:               f.Address,
		TwitterUsername:       f.TwitterUsername,
		TwitterPfpUrl:         f.TwitterPfpUrl,
		ProfileChecked:        f.ProfileChecked,
		Price:                 f.Price,
		Supply:                f.Supply,
		Holders:               f.Holders,
		PriceChangePercentage: f.PriceChangePercentage,
	}
}

type GetUserFriendTechKeyWatchlistResponse struct {
	Data []FriendTechKeyWatchlistItemResponse `json:"data"`
}

type TrackFriendTechKeyResponse struct {
	Data FriendTechKeyWatchlistItemResponse `json:"data"`
}

type FriendTechKeyWatchlistItemResponse struct {
	Id              int                   `json:"id"`
	KeyAddress      string                `json:"key_address"`
	ProfileId       string                `json:"profile_id"`
	IncreaseAlertAt int                   `json:"increase_alert_at"`
	DecreaseAlertAt int                   `json:"decrease_alert_at"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
	Metadata        *TrackedFriendTechKey `json:"metadata"`
}

func TrackingFriendTechKeyModelToResponse(m model.FriendTechKeyWatchlistItem, metadata *FriendTechKey) *FriendTechKeyWatchlistItemResponse {
	return &FriendTechKeyWatchlistItemResponse{
		Id:              m.Id,
		KeyAddress:      m.KeyAddress,
		ProfileId:       m.ProfileId,
		IncreaseAlertAt: m.IncreaseAlertAt,
		DecreaseAlertAt: m.DecreaseAlertAt,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		Metadata:        CamelCaseToSnakeCaseFriendTechKey(metadata),
	}
}

type FriendTechKeyPriceHistoryResponse struct {
	Data []FriendTechKeyPrice `json:"data"`
}

type FriendTechKeyPrice struct {
	Time   time.Time       `json:"time"`
	Supply int             `json:"supply"`
	Holder int             `json:"holder"`
	Price  decimal.Decimal `json:"price"`
}

type FriendTechKeyTransactionsResponse struct {
	Data []friendtechtypes.FriendTechKeyTransaction `json:"data"`
}
