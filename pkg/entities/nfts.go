package entities

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/contracts/erc721"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	usernftwatchlistitem "github.com/defipod/mochi/pkg/repo/user_nft_watchlist_items"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

var (
	mapChainChainId = map[string]string{
		"eth":    "1",
		"heco":   "128",
		"bsc":    "56",
		"matic":  "137",
		"op":     "10",
		"btt":    "199",
		"okt":    "66",
		"movr":   "1285",
		"celo":   "42220",
		"metis":  "1088",
		"cro":    "25",
		"xdai":   "0x64",
		"boba":   "288",
		"ftm":    "250",
		"avax":   "0xa86a",
		"arb":    "42161",
		"aurora": "1313161554",
	}
)

func (e *Entity) GetNFTDetail(symbol, tokenID, guildID string) (*response.IndexerGetNFTTokenDetailResponseWithSuggestions, error) {
	suggest := []response.CollectionSuggestions{}
	// handle query by address
	if len(symbol) > 1 && symbol[:2] == "0x" {
		data, err := e.GetNFTDetailByAddress(symbol, tokenID)
		if err != nil {
			e.log.Errorf(err, "[e.GetNFTDetailByAddress] failed to get nft collection by address")
			return nil, err
		}
		return data, nil
	}

	// get collection
	collections, err := e.repo.NFTCollection.GetBySymbolorName(symbol)
	// cannot find collection => return suggested collections
	if err != nil || len(collections) == 0 {
		suggest, err = e.GetNFTSuggestion(symbol, tokenID)
		if err != nil {
			e.log.Errorf(err, "[repo.NFTCollection.GetBySymbolorName] failed to get nft collection by symbol %s", symbol)
			return nil, err
		}
		return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
			Suggestions: suggest,
		}, nil
	}

	// found multiple symbols => only suggest those
	if len(collections) != 1 {
		var defaultSymbol *response.CollectionSuggestions
		// check default symbol
		symbols, _ := e.GetDefaultCollectionSymbol(guildID)
		for _, col := range collections {
			// if found default symbol
			if len(symbols) != 0 && checkIsDefaultSymbol(symbols, &col) {
				def := response.CollectionSuggestions{
					Address: col.Address,
					Chain:   util.ConvertChainIDToChain(col.ChainID),
					Name:    col.Name,
					Symbol:  col.Symbol,
				}
				defaultSymbol = &def
			}
			suggest = append(suggest, response.CollectionSuggestions{
				Name:    col.Name,
				Symbol:  col.Symbol,
				Address: col.Address,
				Chain:   util.ConvertChainIDToChain(col.ChainID),
			})
		}
		return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
			Suggestions:   suggest,
			DefaultSymbol: defaultSymbol,
		}, nil
	}

	// db returned 1 collection
	collection := collections[0]
	data, err := e.getTokenDetailFromIndexer(collection.Address, tokenID)
	if err != nil {
		e.log.Errorf(err, "[e.getTokenDetailFromIndexer] failed to get nft indexer detail")
		return nil, err
	}

	finalData := make([]response.NftListingMarketplace, 0)
	if len(data.Data.Marketplace) > 0 {
		for _, marketplace := range data.Data.Marketplace {
			marketplace.ItemUrl = util.GetTokenMarketplaceUrl(collection.Address, collection.Symbol, tokenID, marketplace.PlatformName)
			finalData = append(finalData, marketplace)
		}
		data.Data.Marketplace = finalData
	}

	// empty response
	if data == nil {
		e.log.Infof("[indexer.GetNFTDetail] no nft data from indexer")
		err := fmt.Errorf("no nft data from indexer")
		return nil, err
	}

	data.Data.Image = util.StandardizeUri(data.Data.Image)
	return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
		Data:        data.Data,
		Suggestions: suggest,
	}, nil
}

func (e *Entity) GetNFTActivity(collectionAddress, tokenID, query string) (*response.GetNFTActivityResponse, error) {
	res, err := e.getNFTActivityFromIndexer(collectionAddress, tokenID, query)
	if err != nil {
		e.log.Errorf(err, "[e.getNFTActivityFromIndexer] failed to get nft indexer activity")
		return nil, err
	}

	// empty response
	if res == nil {
		e.log.Infof("[indexer.GetNFTActivity] no nft data from indexer")
		err := fmt.Errorf("no nft activity data from indexer")
		return nil, err
	}

	return &response.GetNFTActivityResponse{
		Data: response.GetNFTActivityData{
			Data: res.Data,
			Metadata: util.Pagination{
				Page:  res.Page,
				Size:  res.Size,
				Total: res.Total,
			},
		},
	}, nil
}

func (e *Entity) GetNFTTokenTransactionHistory(collectionAddress, tokenID string) (*response.IndexerGetNFTTokenTxHistoryResponse, error) {
	res, err := e.getNftTokenTransactionHistory(collectionAddress, tokenID)
	if err != nil {
		e.log.Errorf(err, "[e.GetNFTTokenTransactionHistory] failed to get nft indexer activity")
		return nil, err
	}

	return res, nil
}

func checkIsDefaultSymbol(defaults []model.GuildConfigDefaultCollection, symbol *model.NFTCollection) bool {
	for _, def := range defaults {
		if def.Address == symbol.Address && def.ChainID == symbol.ChainID {
			return true
		}
	}
	return false
}

type NFTCollectionData struct {
	TokenAddress string    `json:"token_address"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	ContractType string    `json:"contract_type"`
	SyncedAt     time.Time `json:"synced_at"`
}

type MoralisMessageFail struct {
	Message string `json:"message"`
}

func (e *Entity) GetNFTSuggestion(symbol string, tokenID string) ([]response.CollectionSuggestions, error) {
	// get collections that are correct 50%
	matches, err := e.repo.NFTCollection.GetSuggestionsBySymbolorName(symbol, len(symbol)/2)
	if err != nil {
		if err.Error() == "record not found" {
			e.log.Info("[repo.NFTCollection.GetSuggestionsBySymbolorName] found no suggestions")
			return nil, fmt.Errorf("found no suggestions")
		} else {
			e.log.Errorf(err, "[repo.NFTCollection.GetSuggestionsBySymbolorName] failed to get nft suggestions for symbol %s", symbol)
			return nil, fmt.Errorf("[repo.NFTCollection.GetSuggestionsBySymbolorName] failed to get nft suggestions: %s", err)
		}
	}

	res := []response.CollectionSuggestions{}
	for _, col := range matches {
		res = append(res, response.CollectionSuggestions{
			Name:    col.Name,
			Symbol:  col.Symbol,
			Address: col.Address,
			Chain:   util.ConvertChainIDToChain(col.ChainID),
		})
	}

	return res, nil
}

func (e *Entity) getTokenDetailFromIndexer(address string, tokenID string) (*response.IndexerGetNFTTokenDetailResponse, error) {
	data, err := e.indexer.GetNFTDetail(address, tokenID)
	// cannot find collection in indexer
	if err != nil {
		if err.Error() == "record not found" {
			e.log.Errorf(err, "[indexer.GetNFTDetail] indexer: record nft not found")
			err = fmt.Errorf("Token not found")
		} else {
			e.log.Errorf(err, "[indexer.GetNFTDetail] failed to get nft from indexer")
			err = fmt.Errorf("[e.GetNFTDetail] failed to get nft from indexer: %v", err)
		}
		return nil, err
	}
	return data, nil
}

func (e *Entity) getNFTActivityFromIndexer(collectionAddress, tokenID, query string) (*response.IndexerGetNFTActivityResponse, error) {
	data, err := e.indexer.GetNFTActivity(collectionAddress, tokenID, query)
	if err != nil {
		if err.Error() == "record not found" {
			e.log.Errorf(err, "[indexer.GetNFTActivity] indexer: record nft activity not found")
			err = fmt.Errorf("token not found")
		} else {
			e.log.Errorf(err, "[indexer.GetNFTActivity] failed to get nft activity from indexer")
			err = fmt.Errorf("[e.GetNFTActivity] failed to get nft activity from indexer: %v", err)
		}
		return nil, err
	}
	return data, nil
}

func (e *Entity) getNftTokenTransactionHistory(collectionAddress, tokenID string) (*response.IndexerGetNFTTokenTxHistoryResponse, error) {
	data, err := e.indexer.GetNFTTokenTxHistory(collectionAddress, tokenID)
	if err != nil {
		e.log.Fields(logger.Fields{"collectionAddress": collectionAddress, "tokenID": tokenID}).Errorf(err, "[indexer.GetNFTTokenTxHistory] failed to get nft token tx history from indexer")
		return nil, err
	}

	return data, nil
}

func (e *Entity) GetNFTDetailByAddress(address string, tokenID string) (*response.IndexerGetNFTTokenDetailResponseWithSuggestions, error) {
	exist, err := e.CheckExistNftCollection(address)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] failed to get nft collection by address %s", address)
		return nil, err
	}

	if !exist {
		return nil, fmt.Errorf("record not found")
	}

	data, err := e.getTokenDetailFromIndexer(address, tokenID)
	if err != nil {
		e.log.Errorf(err, "[e.getTokenDetailFromIndexer] failed to get nft indexer detail")
		return nil, err
	}

	collection, err := e.repo.NFTCollection.GetByAddress(address)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] failed to get nft collection")
		return nil, err
	}

	finalData := make([]response.NftListingMarketplace, 0)
	if len(data.Data.Marketplace) > 0 {
		for _, marketplace := range data.Data.Marketplace {
			marketplace.ItemUrl = util.GetTokenMarketplaceUrl(collection.Address, collection.Symbol, tokenID, marketplace.PlatformName)
			finalData = append(finalData, marketplace)
		}
		data.Data.Marketplace = finalData
	}
	return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
		Data: data.Data,
	}, nil

}

func (e *Entity) CheckExistNftCollection(address string) (bool, error) {
	_, err := e.repo.NFTCollection.GetByAddress(address)
	// cannot find collection in db
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		} else {
			e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] failed to get nft collection by address %s", address)
			err = errors.New("failed to get nft collection")
			return false, err
		}
	}
	return true, nil
}

func (e *Entity) CheckIsSync(address string) (bool, error) {
	indexerContract, err := e.indexer.GetNFTContract(address)
	if err != nil {
		e.log.Errorf(err, "[indexer.GetNFTContract] failed to get nft contract by address %s", address)
		return false, err
	}

	return indexerContract.IsSynced, nil
}

func GetNFTCollectionFromMoralis(address, chain string, cfg config.Config) (*NFTCollectionData, error) {
	colData := &NFTCollectionData{}
	moralisApi := "https://deep-index.moralis.io/api/v2/nft/%s/metadata?chain=%s"
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(moralisApi, address, chain), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Key", cfg.MoralisXApiKey)
	q := req.URL.Query()

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		mesErr := &MoralisMessageFail{}
		mes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(mes, &mesErr)
		if err != nil {
			return nil, err
		}
		err = fmt.Errorf("%v", mesErr.Message)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &colData)
	if err != nil {
		return nil, err
	}

	return colData, nil
}

func (e *Entity) CreateSolanaNFTCollection(req request.CreateNFTCollectionRequest) (nftCollection *model.NFTCollection, err error) {
	checkExistNFT, err := e.CheckExistNftCollection(req.Address)
	if err != nil {
		e.log.Errorf(err, "[e.CheckExistNftCollection] failed to check if nft exist: %v", err)
		return nil, err
	}

	if checkExistNFT {
		return e.handleExistCollection(req)
	}

	// get solana metadata collection from blockchain api
	solanaCollection, err := e.svc.BlockchainApi.GetSolanaCollection(req.Address)
	if err != nil {
		e.log.Errorf(err, "[e.svc.BlockchainApi.GetSolanaCollection] failed to get solana collection: %v", err)
		return nil, err
	}

	err = e.indexer.CreateERC721Contract(indexer.CreateERC721ContractRequest{
		Address: req.Address,
		ChainID: 0,
	})
	if err != nil {
		e.log.Errorf(err, "[CreateERC721Contract] failed to create erc721 contract: %v", err)
		return nil, fmt.Errorf("Failed to create erc721 contract: %v", err)
	}

	nftCollection, err = e.repo.NFTCollection.Create(model.NFTCollection{
		Address:    req.Address,
		Symbol:     solanaCollection.Data.Symbol,
		Name:       solanaCollection.Data.Name,
		ChainID:    "0",
		ERCFormat:  "ERC721",
		IsVerified: true,
		Author:     req.Author,
		Image:      solanaCollection.OffChainData.Image,
	})
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.Create] cannot add collection: %v", err)
		return nil, fmt.Errorf("Cannot add collection: %v", err)
	}

	err = e.svc.Discord.NotifyAddNewCollection(req.GuildID, solanaCollection.Data.Name, solanaCollection.Data.Symbol, util.ConvertChainIDToChain("sol"), solanaCollection.OffChainData.Image)
	if err != nil {
		e.log.Errorf(err, "[e.svc.Discord.NotifyAddNewCollection] cannot send embed message: %v", err)
		return nil, fmt.Errorf("Cannot send embed message: %v", err)
	}
	return
}

func (e *Entity) CreateEVMNFTCollection(req request.CreateNFTCollectionRequest) (nftCollection *model.NFTCollection, err error) {
	address := e.HandleMarketplaceLink(req.Address, req.ChainID)
	if address == "collection does not have an address" {
		e.log.Infof("[e.HandleMarketplaceLink] collection %s does not have address", req.Address)
		return nil, fmt.Errorf("Collection does not have an address")
	}

	checksumAddress, err := util.ConvertToChecksumAddr(address)
	if err != nil {
		e.log.Errorf(err, "[util.ConvertToChecksumAddr] failed to convert checksum address: %v", err)
		return nil, fmt.Errorf("Failed to validate address: %v", err)
	}

	checkExistNFT, err := e.CheckExistNftCollection(checksumAddress)
	if err != nil {
		e.log.Errorf(err, "[e.CheckExistNftCollection] failed to check if nft exist: %v", err)
		return nil, err
	}

	req.Address = checksumAddress
	if checkExistNFT {
		return e.handleExistCollection(req)
	}

	convertedChainId := util.ConvertChainToChainId(req.ChainID)
	chainID, err := strconv.Atoi(convertedChainId)
	if err != nil {
		e.log.Errorf(err, "[util.ConvertChainToChainId] failed to convert chain to chainId: %v", err)
		return nil, fmt.Errorf("Failed to convert chain to chainId: %v", err)
	}

	image, err := e.getImageFromMarketPlace(chainID, req.Address)
	if err != nil {
		e.log.Errorf(err, "[e.getImageFromMarketPlace] failed to get image from market place: %v", err)
		return nil, err
	}

	// query name and symbol from contract
	name, symbol, err := e.abi.GetNameAndSymbol(req.Address, int64(chainID))
	if err != nil {
		e.log.Errorf(err, "[GetNameAndSymbol] cannot get name and symbol of contract: %s | chainId %d", req.Address, chainID)
		return nil, fmt.Errorf("Cannot get name and symbol of contract: %v", err)
	}

	// host image to cloud if necessary
	image, err = e.svc.Cloud.HostImageToGCS(image, strings.ReplaceAll(name, " ", ""))
	if err != nil {
		e.log.Errorf(err, "[cloud.HostImageToGCS] failed to host image to GCS: %v", err)
		return nil, err
	}

	err = e.indexer.CreateERC721Contract(indexer.CreateERC721ContractRequest{
		Address: req.Address,
		ChainID: chainID,
	})
	if err != nil {
		e.log.Errorf(err, "[CreateERC721Contract] failed to create erc721 contract: %v", err)
		return nil, fmt.Errorf("Failed to create erc721 contract: %v", err)
	}

	nftCollection, err = e.repo.NFTCollection.Create(model.NFTCollection{
		Address:    req.Address,
		Symbol:     symbol,
		Name:       name,
		ChainID:    convertedChainId,
		ERCFormat:  "ERC721",
		IsVerified: true,
		Author:     req.Author,
		Image:      image,
	})
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.Create] cannot add collection: %v", err)
		return nil, fmt.Errorf("Cannot add collection: %v", err)
	}

	err = e.svc.Discord.NotifyAddNewCollection(req.GuildID, name, symbol, util.ConvertChainIDToChain(convertedChainId), image)
	if err != nil {
		e.log.Errorf(err, "[e.svc.Discord.NotifyAddNewCollection] cannot send embed message: %v", err)
		return nil, fmt.Errorf("Cannot send embed message: %v", err)
	}

	//Add collection to podtown
	go CreatePodtownNFTCollection(model.NFTCollection{
		Address:   req.Address,
		Symbol:    symbol,
		Name:      name,
		ChainID:   convertedChainId,
		ERCFormat: "ERC721",
	}, e.cfg)

	return
}

func (e *Entity) handleExistCollection(req request.CreateNFTCollectionRequest) (*model.NFTCollection, error) {
	isSync, err := e.CheckIsSync(req.Address)
	if err != nil {
		e.log.Errorf(err, "[e.CheckIsSync] failed to check if nft is synced: %v", err)
		return nil, err
	}

	if !isSync {
		e.log.Infof("[e.CheckIsSync] Already added. Nft is in sync progress")
		return nil, fmt.Errorf("Already added. Nft is in sync progress")
	} else {
		e.log.Infof("[e.CheckIsSync] Already added. Nft is done with sync")
		return nil, fmt.Errorf("Already added. Nft is done with sync")
	}
}

func (e *Entity) getImageFromMarketPlace(chainID int, address string) (string, error) {
	if chainID == 1 {
		collection, err := e.marketplace.GetOpenseaAssetContract(address)
		if err != nil {
			e.log.Errorf(err, "[GetOpenseaAssetContract] cannot get contract: %s | chainId %d", address, chainID)
			return "", fmt.Errorf("Cannot get contract: %v", err)
		}
		return collection.Collection.Image, nil
	}
	if chainID == 250 {
		collection, err := e.marketplace.GetCollectionFromPaintswap(address)
		if err != nil {
			e.log.Errorf(err, "[GetCollectionFromPaintswap] cannot get collection: %s | chainId %d", address, chainID)
			return "", fmt.Errorf("Cannot get collection: %v", err)
		}
		return collection.Collection.Image, nil
	}
	if chainID == 10 {
		collection, err := e.marketplace.GetCollectionFromQuixotic(address)
		if err != nil {
			e.log.Errorf(err, "[GetCollectionFromQuixotic] cannot get collection: %s | chainId %d", address, chainID)
			return "", fmt.Errorf("Cannot get collection: %v", err)
		}
		return collection.ImageUrl, nil
	}

	return "", nil
}

func CreatePodtownNFTCollection(collection model.NFTCollection, cfg config.Config) (err error) {
	body, err := json.Marshal(collection)
	if err != nil {
		return err
	}
	jsonBody := bytes.NewBuffer(body)
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/nft/collection", cfg.PodtownServerHost), jsonBody)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return nil
}

func (e *Entity) ListAllNFTCollections() ([]model.NFTCollection, error) {
	return e.repo.NFTCollection.ListAll()
}

func (e *Entity) ListAllNFTCollectionConfigs() ([]model.NFTCollectionConfig, error) {
	return e.repo.NFTCollection.ListAllNFTCollectionConfigs()
}

func (e *Entity) GetNFTBalanceFunc(config model.NFTCollectionConfig) (func(address string) (int, error), error) {

	chainID, err := strconv.Atoi(config.ChainID)
	if err != nil {
		e.log.Errorf(err, "[strconv.Atoi] failed to convert chain id %s to int", config.ChainID)
		return nil, fmt.Errorf("failed to convert chain id %s to int: %v", config.ChainID, err)
	}

	chain, err := e.repo.Chain.GetByID(chainID)
	if err != nil {
		e.log.Errorf(err, "[repo.Chain.GetByID] failed to get chain by id %s", config.ChainID)
		return nil, fmt.Errorf("failed to get chain by id %s: %v", config.ChainID, err)
	}

	client, err := ethclient.Dial(chain.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to chain client: %v", err.Error())
	}

	var balanceOf func(string) (int, error)
	switch strings.ToLower(config.ERCFormat) {
	case "721", "erc721":
		contract721, err := erc721.NewErc721(common.HexToAddress(config.Address), client)
		if err != nil {
			e.log.Errorf(err, "[erc721.NewErc721] failed to init erc721 contract")
			return nil, fmt.Errorf("failed to init erc721 contract: %v", err.Error())
		}

		balanceOf = func(address string) (int, error) {
			b, err := contract721.BalanceOf(nil, common.HexToAddress(address))
			if err != nil {
				e.log.Errorf(err, "[contract721.BalanceOf] failed to get balance of %s in chain %s", address, config.ChainID)
				return 0, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, config.ChainID, err.Error())
			}
			return int(b.Int64()), nil
		}

	// case "1155", "erc1155":
	// 	contract1155, err := erc1155.NewErc1155(common.HexToAddress(config.Address), client)
	// 	if err != nil {
	// 		e.log.Errorf(err, "[erc1155.NewErc1155] failed to init erc1155 contract")
	// 		return nil, fmt.Errorf("failed to init erc1155 contract: %v", err.Error())
	// 	}

	// 	tokenID, err := strconv.ParseInt(config.TokenID, 10, 64)
	// 	if err != nil {
	// 		e.log.Errorf(err, "[strconv.ParseInt] token id is not valid")
	// 		return nil, fmt.Errorf("token id is not valid")
	// 	}

	// 	balanceOf = func(address string) (int, error) {
	// 		b, err := contract1155.BalanceOf(nil, common.HexToAddress(address), big.NewInt(tokenID))
	// 		if err != nil {
	// 			e.log.Errorf(err, "[contract1155.BalanceOf] failed to get balance of %s in chain %s", address, config.ChainID)
	// 			return 0, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, config.ChainID, err.Error())
	// 		}
	// 		return int(b.Int64()), nil
	// 	}

	default:
		e.log.Errorf(err, "[GetNFTBalanceFunc] erc format %s not supported", config.ERCFormat)
		return nil, fmt.Errorf("erc format %s not supported", config.ERCFormat)
	}

	return balanceOf, nil
}

func (e *Entity) NewUserNFTBalance(balance model.UserNFTBalance) error {
	err := e.repo.UserNFTBalance.Upsert(balance)
	if err != nil {
		e.log.Errorf(err, "[repo.UserNFTBalance.Upsert] failed to upsert user nft balance")
		return fmt.Errorf("failed to upsert user nft balance: %v", err.Error())
	}
	return nil
}

func (e *Entity) GetNFTCollectionTickers(symbol, rawQuery string) (*response.IndexerNFTCollectionTickersResponse, error) {
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Infof("Nft colletion ticker not found, symbol: %s", symbol)
			return &response.IndexerNFTCollectionTickersResponse{Data: nil}, nil
		}
		e.log.Errorf(err, "[repo.NFTCollection.GetBySymbol] failed to get nft collection by symbol %s", symbol)
		return nil, err
	}

	res, err := e.indexer.GetNFTCollectionTickers(collection.Address, rawQuery)
	if err != nil {
		e.log.Errorf(err, "[indexer.GetNFTCollectionTickers] failed to get nft collection tickers by %s and %s", collection.Address, rawQuery)
		return nil, err
	}

	for _, ts := range res.Data.Tickers.Timestamps {
		time := time.UnixMilli(ts)
		res.Data.Tickers.Times = append(res.Data.Tickers.Times, time.Format("01-02"))
	}
	return res, nil
}

func (e *Entity) GetNFTCollections(p string, s string) (*response.NFTCollectionsResponse, error) {
	page, _ := strconv.Atoi(p)
	size, _ := strconv.Atoi(s)
	data, total, err := e.repo.NFTCollection.ListAllWithPaging(page, size)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.ListAllWithPaging] failed to list all nft collection with paging")
		return nil, err
	}

	for i, _ := range data {
		data[i].Image = util.StandardizeUri(data[i].Image)
	}

	return &response.NFTCollectionsResponse{
		Data: response.NFTCollectionsData{
			Metadata: util.Pagination{
				Page:  int64(page),
				Size:  int64(size),
				Total: total,
			},
			Data: data,
		},
	}, err
}

func (e *Entity) GetNFTTokens(symbol, query string) (*response.IndexerGetNFTTokensResponse, error) {
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetBySymbol] failed to get nft collection by symbol %s", symbol)
		return nil, err
	}
	if collection.Address == "" {
		e.log.Errorf(err, "[GetNFTTokens] invalid address - collection %s", collection.ID.UUID)
		return nil, fmt.Errorf("invalid address - collection %s", collection.ID.UUID)
	}
	data, err := e.svc.Indexer.GetNFTTokens(collection.Address, query)
	if err != nil {
		e.log.Errorf(err, "[svc.Indexer.GetNFTTokens] failed to get nft tokens by %s and  %s", collection.Address, query)
		return nil, err
	}
	return data, nil
}

func (e *Entity) CreateNFTSalesTracker(addr, platform, guildID string) error {
	checksum, err := util.ConvertToChecksumAddr(addr)
	if err != nil {
		e.log.Errorf(err, "[util.ConvertToChecksumAddr] cannot convert to checksum")
		return fmt.Errorf("invalid contract address")
	}
	config, err := e.GetSalesTrackerConfig(guildID)
	if err != nil {
		e.log.Errorf(err, "[e.GetSalesTrackerConfig] fail to get sale track config by guildID %d", guildID)
		return err
	}

	return e.repo.NFTSalesTracker.FirstOrCreate(&model.InsertNFTSalesTracker{
		ContractAddress: checksum,
		Platform:        platform,
		SalesConfigID:   config.ID.UUID.String(),
	})
}

func (e *Entity) GetDetailNftCollection(symbol string) (*model.NFTCollectionDetail, error) {
	if symbol[:2] == "0x" {
		data, err := e.repo.NFTCollection.GetByAddress(symbol)
		if err != nil {
			e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] failed to get nft collection by address")
			return nil, err
		}
		return e.GetCollectionWithChainDetail(data), nil
	}

	collection, err := e.repo.NFTCollection.GetBySymbolorName(symbol)
	if err != nil || len(collection) == 0 {
		e.log.Errorf(err, "[repo.NFTCollection.GetBySymbolorName] failed to get nft collection by %s", symbol)
		return nil, err
	}

	return e.GetCollectionWithChainDetail(&collection[0]), nil
}

func (e *Entity) GetCollectionWithChainDetail(collection *model.NFTCollection) *model.NFTCollectionDetail {
	chainId, _ := strconv.Atoi(collection.ChainID)
	var chainPtr *model.Chain

	// if chain not exist return chain=nil
	chain, err := e.repo.Chain.GetByID(chainId)
	chainPtr = &chain
	if err != nil {
		e.log.Infof("[e.repo.Chain.GetByID] failed to get chain: %s", collection.ChainID)
		chainPtr = nil
	}
	return &model.NFTCollectionDetail{
		ID:         collection.ID,
		Address:    collection.Address,
		Name:       collection.Name,
		Symbol:     collection.Symbol,
		ChainID:    collection.ChainID,
		Chain:      chainPtr,
		ERCFormat:  collection.ERCFormat,
		IsVerified: collection.IsVerified,
		CreatedAt:  collection.CreatedAt,
		Image:      util.StandardizeUri(collection.Image),
		Author:     collection.Author,
	}
}

func (e *Entity) GetAllNFTSalesTracker() ([]response.NFTSalesTrackerResponse, error) {
	resp := []response.NFTSalesTrackerResponse{}
	data, err := e.repo.NFTSalesTracker.GetAll()
	if err != nil {
		e.log.Errorf(err, "[repo.NFTSalesTracker.GetAll] failed to get all nft sales trackers")
		return nil, err
	}
	for _, item := range data {
		resp = append(resp, response.NFTSalesTrackerResponse{
			ContractAddress: item.ContractAddress,
			Platform:        item.Platform,
			GuildID:         item.GuildConfigSalesTracker.GuildID,
			ChannelID:       item.GuildConfigSalesTracker.ChannelID,
		})
	}
	return resp, nil
}

func (e *Entity) DeleteNFTSalesTracker(guildID, contractAddress string) error {
	return e.repo.NFTSalesTracker.DeleteNFTSalesTrackerByContractAddress(contractAddress)
}

func (e *Entity) GetNFTSaleSTrackerByGuildID(guildID string) (*response.NFTSalesTrackerGuildResponse, error) {
	data, err := e.repo.NFTSalesTracker.GetSalesTrackerByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entity.GetNFTSaleSTrackerByGuildID] failed to get nft sales trackers")
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	return &response.NFTSalesTrackerGuildResponse{
		ID:         data[0].SalesConfigID,
		GuildID:    guildID,
		ChannelID:  data[0].GuildConfigSalesTracker.ChannelID,
		Collection: data,
	}, nil
}

func (e *Entity) GetNewListedNFTCollection(interval string, page string, size string) (*response.NFTNewListedResponse, error) {
	itv, _ := strconv.Atoi(interval)
	pg, _ := strconv.Atoi(page)
	lim, _ := strconv.Atoi(size)
	data, total, err := e.repo.NFTCollection.GetNewListed(itv, pg, lim)
	for i, ele := range data {
		chainId, _ := strconv.Atoi(ele.ChainID)
		chain, err := e.repo.Chain.GetByID(chainId)
		if err != nil {
			e.log.Errorf(err, "[repo.Chain.GetByID] failed to get chain %d", chainId)
			return nil, err
		}
		data[i].Chain = chain.Name
	}
	return &response.NFTNewListedResponse{
		Pagination: util.Pagination{
			Page:  int64(pg),
			Size:  int64(lim),
			Total: total,
		},
		Data: data,
	}, err
}

func (e *Entity) GetNftMetadataAttrIcon() (*response.NftMetadataAttrIconResponse, error) {
	data, err := e.indexer.GetNftMetadataAttrIcon()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *Entity) GetCollectionCount() (*response.NFTCollectionCount, error) {
	_, nr_of_eth, err := e.repo.NFTCollection.GetByChain(1)
	if err != nil {
		e.log.Errorf(err, "[e.GetCollectionCount] cannot count number of ETH collections")
		return nil, err
	}
	_, nr_of_ftm, err := e.repo.NFTCollection.GetByChain(250)
	if err != nil {
		e.log.Errorf(err, "[e.GetCollectionCount] cannot count number of FTM collections")
		return nil, err
	}
	_, nr_of_op, err := e.repo.NFTCollection.GetByChain(10)
	if err != nil {
		e.log.Errorf(err, "[e.GetCollectionCount] cannot count number of OP collections")
		return nil, err
	}
	return &response.NFTCollectionCount{
		Total:    nr_of_eth + nr_of_ftm + nr_of_op,
		ETHCount: nr_of_eth,
		FTMCount: nr_of_ftm,
		OPCount:  nr_of_op,
	}, nil
}

func (e *Entity) GetNFTCollectionByAddressChain(address, chainId string) (*model.NFTCollection, error) {
	collection, err := e.repo.NFTCollection.GetByAddressChainId(address, chainId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Infof("Collection is not exist, address: %s", address)
			return nil, nil
		}
		e.log.Errorf(err, "[repo.NFTCollection.GetNFTCollectionByAddress] failed to get nft collection by address %s", address)
		return nil, err
	}
	collection.Image = util.StandardizeUri(collection.Image)

	return collection, nil
}

func (e *Entity) UpdateNFTCollection(address string) error {
	collection, err := e.repo.NFTCollection.GetByAddress(address)
	if err != nil {
		e.log.Errorf(err, "[e.UpdateNFTCollection] cannot get address")
		return err
	}
	// if image already valid, function return same string
	image, err := e.svc.Cloud.HostImageToGCS(collection.Image, strings.ReplaceAll(collection.Name, " ", ""))
	if err != nil {
		e.log.Errorf(err, "[e.UpdateNFTCollection] cannot host image")
		return err
	}
	if image != collection.Image {
		err := e.repo.NFTCollection.UpdateImage(address, image)
		if err != nil {
			e.log.Errorf(err, "[e.UpdateNFTCollection] cannot update image")
			return err
		}
	}
	return nil
}

func (e *Entity) AddNftWatchlist(req request.AddNftWatchlistRequest) (*response.NftWatchlistSuggestResponse, error) {
	if req.CollectionAddress == "" && req.Chain == "" {
		suggestNftCollection, collection, err := e.SuggestCollection(req)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[e.SuggestCollection] failed to get nft collection suggestion")
			return nil, err
		}

		// case symbol has suggestion
		if suggestNftCollection != nil && collection == nil {
			return &response.NftWatchlistSuggestResponse{Data: suggestNftCollection}, nil
		}

		if collection != nil && suggestNftCollection == nil {
			req.CollectionAddress = collection.Address
			req.Chain = util.ConvertChainIDToChain(collection.ChainID)
		}
	}

	// case has collection address + chain id / not have but suggest return 1 result
	collection, err := e.repo.NFTCollection.GetByAddressChainId(req.CollectionAddress, util.ConvertChainToChainId(req.Chain))
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[e.AddNftWatchlist] failed to get nft collection by address and chain id")
		return nil, err
	}

	chainID, _ := strconv.Atoi(collection.ChainID)
	err = e.repo.UserNftWatchlistItem.Create(&model.UserNftWatchlistItem{
		UserID:            req.UserID,
		Symbol:            collection.Symbol,
		CollectionAddress: collection.Address,
		ChainId:           int64(chainID),
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[UserNftWatchlistItem.Create] failed to create user nft wl")
		return nil, err
	}

	return &response.NftWatchlistSuggestResponse{Data: nil}, nil
}

func (e *Entity) SuggestCollection(req request.AddNftWatchlistRequest) (*response.NftWatchlistSuggest, *model.NFTCollection, error) {
	// get collection
	suggest := []response.CollectionSuggestions{}
	collections, err := e.repo.NFTCollection.GetBySymbolorName(req.CollectionSymbol)
	// cannot find collection => return suggested collections
	if err != nil || len(collections) == 0 {
		suggest, err = e.GetNFTSuggestion(req.CollectionSymbol, "")
		if err != nil {
			e.log.Errorf(err, "[repo.NFTCollection.GetBySymbolorName] failed to get nft collection by symbol %s", req.CollectionSymbol)
			return nil, nil, err
		}
		return &response.NftWatchlistSuggest{
			Suggestions: suggest,
		}, nil, nil
	}

	// found multiple symbols => only suggest those
	if len(collections) > 1 {
		var defaultSymbol *response.CollectionSuggestions
		// check default symbol
		symbols, _ := e.GetDefaultCollectionSymbol(req.GuildID)
		for _, col := range collections {
			// if found default symbol
			if len(symbols) != 0 && checkIsDefaultSymbol(symbols, &col) {
				def := response.CollectionSuggestions{
					Address: col.Address,
					Chain:   util.ConvertChainIDToChain(col.ChainID),
					Name:    col.Name,
					Symbol:  col.Symbol,
				}
				defaultSymbol = &def
			}
			suggest = append(suggest, response.CollectionSuggestions{
				Name:    col.Name,
				Symbol:  col.Symbol,
				Address: col.Address,
				Chain:   util.ConvertChainIDToChain(col.ChainID),
			})
		}
		return &response.NftWatchlistSuggest{
			Suggestions:   suggest,
			DefaultSymbol: defaultSymbol,
		}, nil, nil
	}
	return nil, &collections[0], nil
}

func (e *Entity) DeleteNftWatchlist(req request.DeleteNftWatchlistRequest) error {
	rows, err := e.repo.UserNftWatchlistItem.Delete(req.UserID, req.Symbol)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.DeleteNftWatchlist] repo.UserNftWatchlistItem.Delete() failed")
	}
	if rows == 0 {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.DeleteNftWatchlist] item not found")
		return baseerrs.ErrRecordNotFound
	}
	return err
}

func (e *Entity) GetNftWatchlist(req *request.GetNftWatchlistRequest) (*response.GetNftWatchlistResponse, error) {
	q := usernftwatchlistitem.UserNftWatchlistQuery{
		UserID: req.UserID,
		Offset: req.Page * req.Size,
		Limit:  req.Size,
	}
	list, _, err := e.repo.UserNftWatchlistItem.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.GetUserWatchlist] repo.UserWatchlistItem.List() failed")
		return nil, err
	}

	res := make([]response.GetNftWatchlist, 0)
	for _, itm := range list {
		data, err := e.indexer.GetNFTCollectionTickersForWl(itm.CollectionAddress)
		if err != nil {
			e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.GetUserWatchlist] indexer.GetNFTCollectionTickersForWl failed")
			return nil, err
		}

		collection, err := e.repo.NFTCollection.GetByAddressChainId(itm.CollectionAddress, strconv.Itoa(int(itm.ChainId)))
		if err != nil {
			e.log.Fields(logger.Fields{"CollectionAddress": itm.CollectionAddress, "ChainId": itm.ChainId}).Error(err, "[entity.GetUserWatchlist] repo.NFTCollection.GetByAddressChainId failed")
			return nil, err
		}

		if data == nil {
			res = append(res, response.GetNftWatchlist{
				Symbol:                   itm.Symbol,
				Image:                    collection.Image,
				IsPair:                   false,
				Name:                     collection.Name,
				PriceChangePercentage24h: 0,
				SparkLineIn7d: response.SparkLineIn7d{
					Price: []float64{},
				},
				FloorPrice:                        0,
				PriceChangePercentage7dInCurrency: 0,
			})
			continue
		}

		price := make([]float64, 0)
		for _, ticker := range data.Data.Tickers.Prices {
			bigFloatFloorPrice7d := util.StringWeiToEther(ticker.Amount, int(ticker.Token.Decimals))
			floatFloorPrice7d, _ := bigFloatFloorPrice7d.Float64()
			price = append(price, floatFloorPrice7d)
		}
		floatFloorPrice, _ := util.StringWeiToEther(data.Data.FloorPrice.Amount, int(data.Data.FloorPrice.Token.Decimals)).Float64()

		itmRes := response.GetNftWatchlist{
			Symbol:                   itm.Symbol,
			Image:                    collection.Image,
			IsPair:                   false,
			Name:                     collection.Name,
			PriceChangePercentage24h: 0,
			SparkLineIn7d: response.SparkLineIn7d{
				Price: price,
			},
			FloorPrice:                        floatFloorPrice,
			PriceChangePercentage7dInCurrency: (price[len(price)-1] - price[0]) / price[0],
			Token:                             data.Data.FloorPrice.Token,
		}
		res = append(res, itmRes)
	}
	return &response.GetNftWatchlistResponse{Data: res}, nil
}
