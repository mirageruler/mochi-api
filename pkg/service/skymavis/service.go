package skymavis

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetAddressFarming(address string) (*response.WalletFarmingResponse, error)
}