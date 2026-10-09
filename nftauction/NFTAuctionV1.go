// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nftauction

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = time.Tick
	_ = context.Background
)

// NFTAuctionV1AuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type NFTAuctionV1AuctionInfo struct {
	AuctionId              *big.Int
	TokenId                *big.Int
	StartPrice             *big.Int
	StartTime              *big.Int
	DurationHours          *big.Int
	EndTime                *big.Int
	CurrHighestPrice       *big.Int
	HighestBidder          common.Address
	NftContract            common.Address
	Seller                 common.Address
	IsCreated              bool
	IsEnded                bool
	IsToken                bool
	CurrHighestDecimals    uint8
	HighestBidToken        *big.Int
	CurrHighestTokenAmount *big.Int
	AllowedTokens          []*big.Int
}

// NFTAuctionV1CreateAuctionParams is an auto generated low-level Go binding around an user-defined struct.
type NFTAuctionV1CreateAuctionParams struct {
	NftContract   common.Address
	TokenId       *big.Int
	StartPrice    *big.Int
	StartTime     *big.Int
	DurationHours *big.Int
	AllowedTokens []*big.Int
}

// NFTAuctionV1FeedResult is an auto generated low-level Go binding around an user-defined struct.
type NFTAuctionV1FeedResult struct {
	TokenDecimals uint8
	RawPrice      *big.Int
	UpdateTime    *big.Int
	FeedDecimals  uint8
	Decimals      uint8
	Usd18Value    *big.Int
}

// NFTAuctionV1TokenInitConfig is an auto generated low-level Go binding around an user-defined struct.
type NFTAuctionV1TokenInitConfig struct {
	Token     *big.Int
	TokenAddr common.Address
	FeedAddr  common.Address
}

// NFTAuctionV1MetaData contains all meta data concerning the NFTAuctionV1 contract.
var NFTAuctionV1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addTokenCfg\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchAddTokenCfg\",\"inputs\":[{\"name\":\"tokenInitList\",\"type\":\"tuple[]\",\"internalType\":\"structNFTAuctionV1.TokenInitConfig[]\",\"components\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bidAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAuction\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structNFTAuctionV1.CreateAuctionParams\",\"components\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"durationHours\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowedTokens\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delTokenCfg\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAuctionCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAuctionInfo\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNFTAuctionV1.AuctionInfo\",\"components\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"durationHours\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currHighestPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isCreated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isEnded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isToken\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"currHighestDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"highestBidToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currHighestTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowedTokens\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBidPriceReturns\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeedAddr\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeedAddrIsExists\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNtfToken2AuctionId\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAddr\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAddrIsExists\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenInitCfg\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNFTAuctionV1.TokenInitConfig\",\"components\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenIsExists\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUSDByToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"feedRt\",\"type\":\"tuple\",\"internalType\":\"structNFTAuctionV1.FeedResult\",\"components\":[{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rawPrice\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"updateTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"usd18Value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nftSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenInitList\",\"type\":\"tuple[]\",\"internalType\":\"structNFTAuctionV1.TokenInitConfig[]\",\"components\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refund\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updCfgFeedAddr\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updCfgTokenAddr\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AddTokenCfg\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feedAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operateTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionCancel\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionCreated\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isToken\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"allowedTokens\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"startDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"startTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionEnd\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidInfo\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bidTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bidPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bidDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelTokenCfg\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feedAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operateTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GetUSDByToken\",\"inputs\":[{\"name\":\"currTime\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"updateTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"rawPrice\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"feedDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"usd18Value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Refund\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"withdrawee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdCfgFeedAddr\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldFeedAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newFeedAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operateTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdCfgTokenAddr\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operateTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AllowedTokenSizeOver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionAlreadyExists\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AuctionAlreadyStarted\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AuctionEnded\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AuctionExpired\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AuctionNotExists\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AuctionNotExpired\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AuctionNotStarted\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"DurationHoursOutOfRange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EndAuctionTransferFailed\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeedAddrExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeedAddrSameBefore\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidBidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBidToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCheckType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeedAddr\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNftContractAddr\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRawPrice\",\"inputs\":[{\"name\":\"rawPrice\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenAddr\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NftNotApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuctionSeller\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotEthTokenCfg\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotHighestBidderOrSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotNftOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotOverCurrHighestPrice\",\"inputs\":[{\"name\":\"currHighestPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotRefund\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RefundTransferFailed\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SellerCannotBid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartPriceMustGtZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimeMustGtCurrTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimeOverMaxValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAddrExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAddrSameBefore\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenCfgExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenCfgNotExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"refundAfterBid\",\"inputs\":[]}]",
}

// NFTAuctionV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTAuctionV1MetaData.ABI instead.
var NFTAuctionV1ABI = NFTAuctionV1MetaData.ABI

// NFTAuctionV1 is an auto generated Go binding around an Ethereum contract.
type NFTAuctionV1 struct {
	NFTAuctionV1Caller     // Read-only binding to the contract
	NFTAuctionV1Transactor // Write-only binding to the contract
	NFTAuctionV1Filterer   // Log filterer for contract events
}

// NFTAuctionV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type NFTAuctionV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTAuctionV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTAuctionV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTAuctionV1Session struct {
	Contract     *NFTAuctionV1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTAuctionV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTAuctionV1CallerSession struct {
	Contract *NFTAuctionV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NFTAuctionV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTAuctionV1TransactorSession struct {
	Contract     *NFTAuctionV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NFTAuctionV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type NFTAuctionV1Raw struct {
	Contract *NFTAuctionV1 // Generic contract binding to access the raw methods on
}

// NFTAuctionV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTAuctionV1CallerRaw struct {
	Contract *NFTAuctionV1Caller // Generic read-only contract binding to access the raw methods on
}

// NFTAuctionV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTAuctionV1TransactorRaw struct {
	Contract *NFTAuctionV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTAuctionV1 creates a new instance of NFTAuctionV1, bound to a specific deployed contract.
func NewNFTAuctionV1(address common.Address, backend bind.ContractBackend) (*NFTAuctionV1, error) {
	contract, err := bindNFTAuctionV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1{NFTAuctionV1Caller: NFTAuctionV1Caller{contract: contract}, NFTAuctionV1Transactor: NFTAuctionV1Transactor{contract: contract}, NFTAuctionV1Filterer: NFTAuctionV1Filterer{contract: contract}}, nil
}

// NewNFTAuctionV1Caller creates a new read-only instance of NFTAuctionV1, bound to a specific deployed contract.
func NewNFTAuctionV1Caller(address common.Address, caller bind.ContractCaller) (*NFTAuctionV1Caller, error) {
	contract, err := bindNFTAuctionV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1Caller{contract: contract}, nil
}

// NewNFTAuctionV1Transactor creates a new write-only instance of NFTAuctionV1, bound to a specific deployed contract.
func NewNFTAuctionV1Transactor(address common.Address, transactor bind.ContractTransactor) (*NFTAuctionV1Transactor, error) {
	contract, err := bindNFTAuctionV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1Transactor{contract: contract}, nil
}

// NewNFTAuctionV1Filterer creates a new log filterer instance of NFTAuctionV1, bound to a specific deployed contract.
func NewNFTAuctionV1Filterer(address common.Address, filterer bind.ContractFilterer) (*NFTAuctionV1Filterer, error) {
	contract, err := bindNFTAuctionV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1Filterer{contract: contract}, nil
}

// bindNFTAuctionV1 binds a generic wrapper to an already deployed contract.
func bindNFTAuctionV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTAuctionV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTAuctionV1 *NFTAuctionV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTAuctionV1.Contract.NFTAuctionV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTAuctionV1 *NFTAuctionV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.NFTAuctionV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTAuctionV1 *NFTAuctionV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.NFTAuctionV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTAuctionV1 *NFTAuctionV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTAuctionV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTAuctionV1 *NFTAuctionV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTAuctionV1 *NFTAuctionV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _NFTAuctionV1.Contract.UPGRADEINTERFACEVERSION(&_NFTAuctionV1.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NFTAuctionV1.Contract.UPGRADEINTERFACEVERSION(&_NFTAuctionV1.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFTAuctionV1.Contract.BalanceOf(&_NFTAuctionV1.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFTAuctionV1.Contract.BalanceOf(&_NFTAuctionV1.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFTAuctionV1.Contract.GetApproved(&_NFTAuctionV1.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFTAuctionV1.Contract.GetApproved(&_NFTAuctionV1.CallOpts, tokenId)
}

// GetAuctionCount is a free data retrieval call binding the contract method 0xc44e6640.
//
// Solidity: function getAuctionCount() view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetAuctionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getAuctionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAuctionCount is a free data retrieval call binding the contract method 0xc44e6640.
//
// Solidity: function getAuctionCount() view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetAuctionCount() (*big.Int, error) {
	return _NFTAuctionV1.Contract.GetAuctionCount(&_NFTAuctionV1.CallOpts)
}

// GetAuctionCount is a free data retrieval call binding the contract method 0xc44e6640.
//
// Solidity: function getAuctionCount() view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetAuctionCount() (*big.Int, error) {
	return _NFTAuctionV1.Contract.GetAuctionCount(&_NFTAuctionV1.CallOpts)
}

// GetAuctionInfo is a free data retrieval call binding the contract method 0xfc3fc4ed.
//
// Solidity: function getAuctionInfo(uint256 auctionId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,bool,bool,bool,uint8,uint256,uint256,uint256[]))
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetAuctionInfo(opts *bind.CallOpts, auctionId *big.Int) (NFTAuctionV1AuctionInfo, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getAuctionInfo", auctionId)

	if err != nil {
		return *new(NFTAuctionV1AuctionInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(NFTAuctionV1AuctionInfo)).(*NFTAuctionV1AuctionInfo)

	return out0, err

}

// GetAuctionInfo is a free data retrieval call binding the contract method 0xfc3fc4ed.
//
// Solidity: function getAuctionInfo(uint256 auctionId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,bool,bool,bool,uint8,uint256,uint256,uint256[]))
func (_NFTAuctionV1 *NFTAuctionV1Session) GetAuctionInfo(auctionId *big.Int) (NFTAuctionV1AuctionInfo, error) {
	return _NFTAuctionV1.Contract.GetAuctionInfo(&_NFTAuctionV1.CallOpts, auctionId)
}

// GetAuctionInfo is a free data retrieval call binding the contract method 0xfc3fc4ed.
//
// Solidity: function getAuctionInfo(uint256 auctionId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,bool,bool,bool,uint8,uint256,uint256,uint256[]))
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetAuctionInfo(auctionId *big.Int) (NFTAuctionV1AuctionInfo, error) {
	return _NFTAuctionV1.Contract.GetAuctionInfo(&_NFTAuctionV1.CallOpts, auctionId)
}

// GetBidPriceReturns is a free data retrieval call binding the contract method 0xd483fe1a.
//
// Solidity: function getBidPriceReturns(uint256 auctionId, address bidder, uint256 token) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetBidPriceReturns(opts *bind.CallOpts, auctionId *big.Int, bidder common.Address, token *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getBidPriceReturns", auctionId, bidder, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidPriceReturns is a free data retrieval call binding the contract method 0xd483fe1a.
//
// Solidity: function getBidPriceReturns(uint256 auctionId, address bidder, uint256 token) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetBidPriceReturns(auctionId *big.Int, bidder common.Address, token *big.Int) (*big.Int, error) {
	return _NFTAuctionV1.Contract.GetBidPriceReturns(&_NFTAuctionV1.CallOpts, auctionId, bidder, token)
}

// GetBidPriceReturns is a free data retrieval call binding the contract method 0xd483fe1a.
//
// Solidity: function getBidPriceReturns(uint256 auctionId, address bidder, uint256 token) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetBidPriceReturns(auctionId *big.Int, bidder common.Address, token *big.Int) (*big.Int, error) {
	return _NFTAuctionV1.Contract.GetBidPriceReturns(&_NFTAuctionV1.CallOpts, auctionId, bidder, token)
}

// GetFeedAddr is a free data retrieval call binding the contract method 0x9bc49cb1.
//
// Solidity: function getFeedAddr(uint256 token) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetFeedAddr(opts *bind.CallOpts, token *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getFeedAddr", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeedAddr is a free data retrieval call binding the contract method 0x9bc49cb1.
//
// Solidity: function getFeedAddr(uint256 token) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetFeedAddr(token *big.Int) (common.Address, error) {
	return _NFTAuctionV1.Contract.GetFeedAddr(&_NFTAuctionV1.CallOpts, token)
}

// GetFeedAddr is a free data retrieval call binding the contract method 0x9bc49cb1.
//
// Solidity: function getFeedAddr(uint256 token) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetFeedAddr(token *big.Int) (common.Address, error) {
	return _NFTAuctionV1.Contract.GetFeedAddr(&_NFTAuctionV1.CallOpts, token)
}

// GetFeedAddrIsExists is a free data retrieval call binding the contract method 0xdfbf7bfb.
//
// Solidity: function getFeedAddrIsExists(address addr) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetFeedAddrIsExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getFeedAddrIsExists", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeedAddrIsExists is a free data retrieval call binding the contract method 0xdfbf7bfb.
//
// Solidity: function getFeedAddrIsExists(address addr) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetFeedAddrIsExists(addr common.Address) (bool, error) {
	return _NFTAuctionV1.Contract.GetFeedAddrIsExists(&_NFTAuctionV1.CallOpts, addr)
}

// GetFeedAddrIsExists is a free data retrieval call binding the contract method 0xdfbf7bfb.
//
// Solidity: function getFeedAddrIsExists(address addr) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetFeedAddrIsExists(addr common.Address) (bool, error) {
	return _NFTAuctionV1.Contract.GetFeedAddrIsExists(&_NFTAuctionV1.CallOpts, addr)
}

// GetNtfToken2AuctionId is a free data retrieval call binding the contract method 0xa5878cd2.
//
// Solidity: function getNtfToken2AuctionId(address nftContract, uint256 tokenId) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetNtfToken2AuctionId(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getNtfToken2AuctionId", nftContract, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNtfToken2AuctionId is a free data retrieval call binding the contract method 0xa5878cd2.
//
// Solidity: function getNtfToken2AuctionId(address nftContract, uint256 tokenId) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetNtfToken2AuctionId(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _NFTAuctionV1.Contract.GetNtfToken2AuctionId(&_NFTAuctionV1.CallOpts, nftContract, tokenId)
}

// GetNtfToken2AuctionId is a free data retrieval call binding the contract method 0xa5878cd2.
//
// Solidity: function getNtfToken2AuctionId(address nftContract, uint256 tokenId) view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetNtfToken2AuctionId(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _NFTAuctionV1.Contract.GetNtfToken2AuctionId(&_NFTAuctionV1.CallOpts, nftContract, tokenId)
}

// GetTokenAddr is a free data retrieval call binding the contract method 0x04569f5f.
//
// Solidity: function getTokenAddr(uint256 token) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetTokenAddr(opts *bind.CallOpts, token *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getTokenAddr", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddr is a free data retrieval call binding the contract method 0x04569f5f.
//
// Solidity: function getTokenAddr(uint256 token) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetTokenAddr(token *big.Int) (common.Address, error) {
	return _NFTAuctionV1.Contract.GetTokenAddr(&_NFTAuctionV1.CallOpts, token)
}

// GetTokenAddr is a free data retrieval call binding the contract method 0x04569f5f.
//
// Solidity: function getTokenAddr(uint256 token) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetTokenAddr(token *big.Int) (common.Address, error) {
	return _NFTAuctionV1.Contract.GetTokenAddr(&_NFTAuctionV1.CallOpts, token)
}

// GetTokenAddrIsExists is a free data retrieval call binding the contract method 0x88f2b2bd.
//
// Solidity: function getTokenAddrIsExists(address addr) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetTokenAddrIsExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getTokenAddrIsExists", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTokenAddrIsExists is a free data retrieval call binding the contract method 0x88f2b2bd.
//
// Solidity: function getTokenAddrIsExists(address addr) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetTokenAddrIsExists(addr common.Address) (bool, error) {
	return _NFTAuctionV1.Contract.GetTokenAddrIsExists(&_NFTAuctionV1.CallOpts, addr)
}

// GetTokenAddrIsExists is a free data retrieval call binding the contract method 0x88f2b2bd.
//
// Solidity: function getTokenAddrIsExists(address addr) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetTokenAddrIsExists(addr common.Address) (bool, error) {
	return _NFTAuctionV1.Contract.GetTokenAddrIsExists(&_NFTAuctionV1.CallOpts, addr)
}

// GetTokenCount is a free data retrieval call binding the contract method 0x78a89567.
//
// Solidity: function getTokenCount() view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCount is a free data retrieval call binding the contract method 0x78a89567.
//
// Solidity: function getTokenCount() view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetTokenCount() (*big.Int, error) {
	return _NFTAuctionV1.Contract.GetTokenCount(&_NFTAuctionV1.CallOpts)
}

// GetTokenCount is a free data retrieval call binding the contract method 0x78a89567.
//
// Solidity: function getTokenCount() view returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetTokenCount() (*big.Int, error) {
	return _NFTAuctionV1.Contract.GetTokenCount(&_NFTAuctionV1.CallOpts)
}

// GetTokenInitCfg is a free data retrieval call binding the contract method 0x844f433c.
//
// Solidity: function getTokenInitCfg(uint256 auctionId, uint256 token) view returns((uint256,address,address))
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetTokenInitCfg(opts *bind.CallOpts, auctionId *big.Int, token *big.Int) (NFTAuctionV1TokenInitConfig, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getTokenInitCfg", auctionId, token)

	if err != nil {
		return *new(NFTAuctionV1TokenInitConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(NFTAuctionV1TokenInitConfig)).(*NFTAuctionV1TokenInitConfig)

	return out0, err

}

// GetTokenInitCfg is a free data retrieval call binding the contract method 0x844f433c.
//
// Solidity: function getTokenInitCfg(uint256 auctionId, uint256 token) view returns((uint256,address,address))
func (_NFTAuctionV1 *NFTAuctionV1Session) GetTokenInitCfg(auctionId *big.Int, token *big.Int) (NFTAuctionV1TokenInitConfig, error) {
	return _NFTAuctionV1.Contract.GetTokenInitCfg(&_NFTAuctionV1.CallOpts, auctionId, token)
}

// GetTokenInitCfg is a free data retrieval call binding the contract method 0x844f433c.
//
// Solidity: function getTokenInitCfg(uint256 auctionId, uint256 token) view returns((uint256,address,address))
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetTokenInitCfg(auctionId *big.Int, token *big.Int) (NFTAuctionV1TokenInitConfig, error) {
	return _NFTAuctionV1.Contract.GetTokenInitCfg(&_NFTAuctionV1.CallOpts, auctionId, token)
}

// GetTokenIsExists is a free data retrieval call binding the contract method 0x94719489.
//
// Solidity: function getTokenIsExists(uint256 token) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Caller) GetTokenIsExists(opts *bind.CallOpts, token *big.Int) (bool, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "getTokenIsExists", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTokenIsExists is a free data retrieval call binding the contract method 0x94719489.
//
// Solidity: function getTokenIsExists(uint256 token) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetTokenIsExists(token *big.Int) (bool, error) {
	return _NFTAuctionV1.Contract.GetTokenIsExists(&_NFTAuctionV1.CallOpts, token)
}

// GetTokenIsExists is a free data retrieval call binding the contract method 0x94719489.
//
// Solidity: function getTokenIsExists(uint256 token) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) GetTokenIsExists(token *big.Int) (bool, error) {
	return _NFTAuctionV1.Contract.GetTokenIsExists(&_NFTAuctionV1.CallOpts, token)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFTAuctionV1.Contract.IsApprovedForAll(&_NFTAuctionV1.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFTAuctionV1.Contract.IsApprovedForAll(&_NFTAuctionV1.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1Session) Name() (string, error) {
	return _NFTAuctionV1.Contract.Name(&_NFTAuctionV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) Name() (string, error) {
	return _NFTAuctionV1.Contract.Name(&_NFTAuctionV1.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_NFTAuctionV1 *NFTAuctionV1Caller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_NFTAuctionV1 *NFTAuctionV1Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _NFTAuctionV1.Contract.OnERC721Received(&_NFTAuctionV1.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _NFTAuctionV1.Contract.OnERC721Received(&_NFTAuctionV1.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Session) Owner() (common.Address, error) {
	return _NFTAuctionV1.Contract.Owner(&_NFTAuctionV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) Owner() (common.Address, error) {
	return _NFTAuctionV1.Contract.Owner(&_NFTAuctionV1.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFTAuctionV1.Contract.OwnerOf(&_NFTAuctionV1.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFTAuctionV1.Contract.OwnerOf(&_NFTAuctionV1.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuctionV1 *NFTAuctionV1Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuctionV1 *NFTAuctionV1Session) ProxiableUUID() ([32]byte, error) {
	return _NFTAuctionV1.Contract.ProxiableUUID(&_NFTAuctionV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) ProxiableUUID() ([32]byte, error) {
	return _NFTAuctionV1.Contract.ProxiableUUID(&_NFTAuctionV1.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFTAuctionV1.Contract.SupportsInterface(&_NFTAuctionV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFTAuctionV1.Contract.SupportsInterface(&_NFTAuctionV1.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1Session) Symbol() (string, error) {
	return _NFTAuctionV1.Contract.Symbol(&_NFTAuctionV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) Symbol() (string, error) {
	return _NFTAuctionV1.Contract.Symbol(&_NFTAuctionV1.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NFTAuctionV1.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1Session) TokenURI(tokenId *big.Int) (string, error) {
	return _NFTAuctionV1.Contract.TokenURI(&_NFTAuctionV1.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFTAuctionV1 *NFTAuctionV1CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NFTAuctionV1.Contract.TokenURI(&_NFTAuctionV1.CallOpts, tokenId)
}

// AddTokenCfg is a paid mutator transaction binding the contract method 0xd12adadb.
//
// Solidity: function addTokenCfg(uint256 token, address tokenAddr, address feedAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) AddTokenCfg(opts *bind.TransactOpts, token *big.Int, tokenAddr common.Address, feedAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "addTokenCfg", token, tokenAddr, feedAddr)
}

// AddTokenCfg is a paid mutator transaction binding the contract method 0xd12adadb.
//
// Solidity: function addTokenCfg(uint256 token, address tokenAddr, address feedAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) AddTokenCfg(token *big.Int, tokenAddr common.Address, feedAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.AddTokenCfg(&_NFTAuctionV1.TransactOpts, token, tokenAddr, feedAddr)
}

// AddTokenCfg is a paid mutator transaction binding the contract method 0xd12adadb.
//
// Solidity: function addTokenCfg(uint256 token, address tokenAddr, address feedAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) AddTokenCfg(token *big.Int, tokenAddr common.Address, feedAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.AddTokenCfg(&_NFTAuctionV1.TransactOpts, token, tokenAddr, feedAddr)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.Approve(&_NFTAuctionV1.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.Approve(&_NFTAuctionV1.TransactOpts, to, tokenId)
}

// BatchAddTokenCfg is a paid mutator transaction binding the contract method 0x40ecd3da.
//
// Solidity: function batchAddTokenCfg((uint256,address,address)[] tokenInitList) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) BatchAddTokenCfg(opts *bind.TransactOpts, tokenInitList []NFTAuctionV1TokenInitConfig) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "batchAddTokenCfg", tokenInitList)
}

// BatchAddTokenCfg is a paid mutator transaction binding the contract method 0x40ecd3da.
//
// Solidity: function batchAddTokenCfg((uint256,address,address)[] tokenInitList) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) BatchAddTokenCfg(tokenInitList []NFTAuctionV1TokenInitConfig) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.BatchAddTokenCfg(&_NFTAuctionV1.TransactOpts, tokenInitList)
}

// BatchAddTokenCfg is a paid mutator transaction binding the contract method 0x40ecd3da.
//
// Solidity: function batchAddTokenCfg((uint256,address,address)[] tokenInitList) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) BatchAddTokenCfg(tokenInitList []NFTAuctionV1TokenInitConfig) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.BatchAddTokenCfg(&_NFTAuctionV1.TransactOpts, tokenInitList)
}

// BidAuction is a paid mutator transaction binding the contract method 0x6874a3fe.
//
// Solidity: function bidAuction(uint256 auctionId, uint256 token, uint256 amount) payable returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) BidAuction(opts *bind.TransactOpts, auctionId *big.Int, token *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "bidAuction", auctionId, token, amount)
}

// BidAuction is a paid mutator transaction binding the contract method 0x6874a3fe.
//
// Solidity: function bidAuction(uint256 auctionId, uint256 token, uint256 amount) payable returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) BidAuction(auctionId *big.Int, token *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.BidAuction(&_NFTAuctionV1.TransactOpts, auctionId, token, amount)
}

// BidAuction is a paid mutator transaction binding the contract method 0x6874a3fe.
//
// Solidity: function bidAuction(uint256 auctionId, uint256 token, uint256 amount) payable returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) BidAuction(auctionId *big.Int, token *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.BidAuction(&_NFTAuctionV1.TransactOpts, auctionId, token, amount)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) CancelAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "cancelAuction", auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) CancelAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.CancelAuction(&_NFTAuctionV1.TransactOpts, auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) CancelAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.CancelAuction(&_NFTAuctionV1.TransactOpts, auctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xa305fac5.
//
// Solidity: function createAuction((address,uint256,uint256,uint256,uint256,uint256[]) params) returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Transactor) CreateAuction(opts *bind.TransactOpts, params NFTAuctionV1CreateAuctionParams) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "createAuction", params)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xa305fac5.
//
// Solidity: function createAuction((address,uint256,uint256,uint256,uint256,uint256[]) params) returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1Session) CreateAuction(params NFTAuctionV1CreateAuctionParams) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.CreateAuction(&_NFTAuctionV1.TransactOpts, params)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xa305fac5.
//
// Solidity: function createAuction((address,uint256,uint256,uint256,uint256,uint256[]) params) returns(uint256)
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) CreateAuction(params NFTAuctionV1CreateAuctionParams) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.CreateAuction(&_NFTAuctionV1.TransactOpts, params)
}

// DelTokenCfg is a paid mutator transaction binding the contract method 0x0ae0de87.
//
// Solidity: function delTokenCfg(uint256 token) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) DelTokenCfg(opts *bind.TransactOpts, token *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "delTokenCfg", token)
}

// DelTokenCfg is a paid mutator transaction binding the contract method 0x0ae0de87.
//
// Solidity: function delTokenCfg(uint256 token) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) DelTokenCfg(token *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.DelTokenCfg(&_NFTAuctionV1.TransactOpts, token)
}

// DelTokenCfg is a paid mutator transaction binding the contract method 0x0ae0de87.
//
// Solidity: function delTokenCfg(uint256 token) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) DelTokenCfg(token *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.DelTokenCfg(&_NFTAuctionV1.TransactOpts, token)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) EndAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "endAuction", auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.EndAuction(&_NFTAuctionV1.TransactOpts, auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.EndAuction(&_NFTAuctionV1.TransactOpts, auctionId)
}

// GetUSDByToken is a paid mutator transaction binding the contract method 0x66e72eec.
//
// Solidity: function getUSDByToken(uint256 token, uint256 tokenAmount) returns((uint8,int256,uint256,uint8,uint8,uint256) feedRt)
func (_NFTAuctionV1 *NFTAuctionV1Transactor) GetUSDByToken(opts *bind.TransactOpts, token *big.Int, tokenAmount *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "getUSDByToken", token, tokenAmount)
}

// GetUSDByToken is a paid mutator transaction binding the contract method 0x66e72eec.
//
// Solidity: function getUSDByToken(uint256 token, uint256 tokenAmount) returns((uint8,int256,uint256,uint8,uint8,uint256) feedRt)
func (_NFTAuctionV1 *NFTAuctionV1Session) GetUSDByToken(token *big.Int, tokenAmount *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.GetUSDByToken(&_NFTAuctionV1.TransactOpts, token, tokenAmount)
}

// GetUSDByToken is a paid mutator transaction binding the contract method 0x66e72eec.
//
// Solidity: function getUSDByToken(uint256 token, uint256 tokenAmount) returns((uint8,int256,uint256,uint8,uint8,uint256) feedRt)
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) GetUSDByToken(token *big.Int, tokenAmount *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.GetUSDByToken(&_NFTAuctionV1.TransactOpts, token, tokenAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xe05c7a15.
//
// Solidity: function initialize(address initialOwner, string nftName, string nftSymbol, (uint256,address,address)[] tokenInitList) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, nftName string, nftSymbol string, tokenInitList []NFTAuctionV1TokenInitConfig) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "initialize", initialOwner, nftName, nftSymbol, tokenInitList)
}

// Initialize is a paid mutator transaction binding the contract method 0xe05c7a15.
//
// Solidity: function initialize(address initialOwner, string nftName, string nftSymbol, (uint256,address,address)[] tokenInitList) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) Initialize(initialOwner common.Address, nftName string, nftSymbol string, tokenInitList []NFTAuctionV1TokenInitConfig) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.Initialize(&_NFTAuctionV1.TransactOpts, initialOwner, nftName, nftSymbol, tokenInitList)
}

// Initialize is a paid mutator transaction binding the contract method 0xe05c7a15.
//
// Solidity: function initialize(address initialOwner, string nftName, string nftSymbol, (uint256,address,address)[] tokenInitList) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) Initialize(initialOwner common.Address, nftName string, nftSymbol string, tokenInitList []NFTAuctionV1TokenInitConfig) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.Initialize(&_NFTAuctionV1.TransactOpts, initialOwner, nftName, nftSymbol, tokenInitList)
}

// Refund is a paid mutator transaction binding the contract method 0x5af36e3e.
//
// Solidity: function refund(uint256 auctionId, uint256 token) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) Refund(opts *bind.TransactOpts, auctionId *big.Int, token *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "refund", auctionId, token)
}

// Refund is a paid mutator transaction binding the contract method 0x5af36e3e.
//
// Solidity: function refund(uint256 auctionId, uint256 token) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) Refund(auctionId *big.Int, token *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.Refund(&_NFTAuctionV1.TransactOpts, auctionId, token)
}

// Refund is a paid mutator transaction binding the contract method 0x5af36e3e.
//
// Solidity: function refund(uint256 auctionId, uint256 token) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) Refund(auctionId *big.Int, token *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.Refund(&_NFTAuctionV1.TransactOpts, auctionId, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.RenounceOwnership(&_NFTAuctionV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.RenounceOwnership(&_NFTAuctionV1.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.SafeTransferFrom(&_NFTAuctionV1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.SafeTransferFrom(&_NFTAuctionV1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.SafeTransferFrom0(&_NFTAuctionV1.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.SafeTransferFrom0(&_NFTAuctionV1.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.SetApprovalForAll(&_NFTAuctionV1.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.SetApprovalForAll(&_NFTAuctionV1.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.TransferFrom(&_NFTAuctionV1.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.TransferFrom(&_NFTAuctionV1.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.TransferOwnership(&_NFTAuctionV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.TransferOwnership(&_NFTAuctionV1.TransactOpts, newOwner)
}

// UpdCfgFeedAddr is a paid mutator transaction binding the contract method 0xd16da9ba.
//
// Solidity: function updCfgFeedAddr(uint256 token, address feedAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) UpdCfgFeedAddr(opts *bind.TransactOpts, token *big.Int, feedAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "updCfgFeedAddr", token, feedAddr)
}

// UpdCfgFeedAddr is a paid mutator transaction binding the contract method 0xd16da9ba.
//
// Solidity: function updCfgFeedAddr(uint256 token, address feedAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) UpdCfgFeedAddr(token *big.Int, feedAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.UpdCfgFeedAddr(&_NFTAuctionV1.TransactOpts, token, feedAddr)
}

// UpdCfgFeedAddr is a paid mutator transaction binding the contract method 0xd16da9ba.
//
// Solidity: function updCfgFeedAddr(uint256 token, address feedAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) UpdCfgFeedAddr(token *big.Int, feedAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.UpdCfgFeedAddr(&_NFTAuctionV1.TransactOpts, token, feedAddr)
}

// UpdCfgTokenAddr is a paid mutator transaction binding the contract method 0xbff50f5b.
//
// Solidity: function updCfgTokenAddr(uint256 token, address tokenAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) UpdCfgTokenAddr(opts *bind.TransactOpts, token *big.Int, tokenAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "updCfgTokenAddr", token, tokenAddr)
}

// UpdCfgTokenAddr is a paid mutator transaction binding the contract method 0xbff50f5b.
//
// Solidity: function updCfgTokenAddr(uint256 token, address tokenAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) UpdCfgTokenAddr(token *big.Int, tokenAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.UpdCfgTokenAddr(&_NFTAuctionV1.TransactOpts, token, tokenAddr)
}

// UpdCfgTokenAddr is a paid mutator transaction binding the contract method 0xbff50f5b.
//
// Solidity: function updCfgTokenAddr(uint256 token, address tokenAddr) returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) UpdCfgTokenAddr(token *big.Int, tokenAddr common.Address) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.UpdCfgTokenAddr(&_NFTAuctionV1.TransactOpts, token, tokenAddr)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuctionV1 *NFTAuctionV1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuctionV1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuctionV1 *NFTAuctionV1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.UpgradeToAndCall(&_NFTAuctionV1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuctionV1 *NFTAuctionV1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuctionV1.Contract.UpgradeToAndCall(&_NFTAuctionV1.TransactOpts, newImplementation, data)
}

// NFTAuctionV1AddTokenCfgIterator is returned from FilterAddTokenCfg and is used to iterate over the raw logs and unpacked data for AddTokenCfg events raised by the NFTAuctionV1 contract.
type NFTAuctionV1AddTokenCfgIterator struct {
	Event *NFTAuctionV1AddTokenCfg // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1AddTokenCfgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1AddTokenCfg)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1AddTokenCfg)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1AddTokenCfgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1AddTokenCfgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1AddTokenCfg represents a AddTokenCfg event raised by the NFTAuctionV1 contract.
type NFTAuctionV1AddTokenCfg struct {
	Caller      common.Address
	Token       *big.Int
	TokenAddr   common.Address
	FeedAddr    common.Address
	OperateTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddTokenCfg is a free log retrieval operation binding the contract event 0x181031e73abce0c4566a5b09802ce0c1d6d43f6dfde462b27347948e7bc4d7c2.
//
// Solidity: event AddTokenCfg(address indexed caller, uint256 indexed token, address tokenAddr, address feedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterAddTokenCfg(opts *bind.FilterOpts, caller []common.Address, token []*big.Int) (*NFTAuctionV1AddTokenCfgIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "AddTokenCfg", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1AddTokenCfgIterator{contract: _NFTAuctionV1.contract, event: "AddTokenCfg", logs: logs, sub: sub}, nil
}

// WatchAddTokenCfg is a free log subscription operation binding the contract event 0x181031e73abce0c4566a5b09802ce0c1d6d43f6dfde462b27347948e7bc4d7c2.
//
// Solidity: event AddTokenCfg(address indexed caller, uint256 indexed token, address tokenAddr, address feedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchAddTokenCfg(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1AddTokenCfg, caller []common.Address, token []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "AddTokenCfg", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1AddTokenCfg)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "AddTokenCfg", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddTokenCfg is a log parse operation binding the contract event 0x181031e73abce0c4566a5b09802ce0c1d6d43f6dfde462b27347948e7bc4d7c2.
//
// Solidity: event AddTokenCfg(address indexed caller, uint256 indexed token, address tokenAddr, address feedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseAddTokenCfg(log types.Log) (*NFTAuctionV1AddTokenCfg, error) {
	event := new(NFTAuctionV1AddTokenCfg)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "AddTokenCfg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NFTAuctionV1 contract.
type NFTAuctionV1ApprovalIterator struct {
	Event *NFTAuctionV1Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1Approval represents a Approval event raised by the NFTAuctionV1 contract.
type NFTAuctionV1Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NFTAuctionV1ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1ApprovalIterator{contract: _NFTAuctionV1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1Approval)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "Approval", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseApproval(log types.Log) (*NFTAuctionV1Approval, error) {
	event := new(NFTAuctionV1Approval)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NFTAuctionV1 contract.
type NFTAuctionV1ApprovalForAllIterator struct {
	Event *NFTAuctionV1ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1ApprovalForAll represents a ApprovalForAll event raised by the NFTAuctionV1 contract.
type NFTAuctionV1ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NFTAuctionV1ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1ApprovalForAllIterator{contract: _NFTAuctionV1.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1ApprovalForAll)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseApprovalForAll(log types.Log) (*NFTAuctionV1ApprovalForAll, error) {
	event := new(NFTAuctionV1ApprovalForAll)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1AuctionCancelIterator is returned from FilterAuctionCancel and is used to iterate over the raw logs and unpacked data for AuctionCancel events raised by the NFTAuctionV1 contract.
type NFTAuctionV1AuctionCancelIterator struct {
	Event *NFTAuctionV1AuctionCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1AuctionCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1AuctionCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1AuctionCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1AuctionCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1AuctionCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1AuctionCancel represents a AuctionCancel event raised by the NFTAuctionV1 contract.
type NFTAuctionV1AuctionCancel struct {
	AuctionId *big.Int
	Operator  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancel is a free log retrieval operation binding the contract event 0x8b384800a88ae5cdb7c656124e10d94998a89ad2813894368cf6aa20bc5869bc.
//
// Solidity: event AuctionCancel(uint256 indexed auctionId, address indexed operator)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterAuctionCancel(opts *bind.FilterOpts, auctionId []*big.Int, operator []common.Address) (*NFTAuctionV1AuctionCancelIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "AuctionCancel", auctionIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1AuctionCancelIterator{contract: _NFTAuctionV1.contract, event: "AuctionCancel", logs: logs, sub: sub}, nil
}

// WatchAuctionCancel is a free log subscription operation binding the contract event 0x8b384800a88ae5cdb7c656124e10d94998a89ad2813894368cf6aa20bc5869bc.
//
// Solidity: event AuctionCancel(uint256 indexed auctionId, address indexed operator)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchAuctionCancel(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1AuctionCancel, auctionId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "AuctionCancel", auctionIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1AuctionCancel)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "AuctionCancel", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionCancel is a log parse operation binding the contract event 0x8b384800a88ae5cdb7c656124e10d94998a89ad2813894368cf6aa20bc5869bc.
//
// Solidity: event AuctionCancel(uint256 indexed auctionId, address indexed operator)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseAuctionCancel(log types.Log) (*NFTAuctionV1AuctionCancel, error) {
	event := new(NFTAuctionV1AuctionCancel)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "AuctionCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1AuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the NFTAuctionV1 contract.
type NFTAuctionV1AuctionCreatedIterator struct {
	Event *NFTAuctionV1AuctionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1AuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1AuctionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1AuctionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1AuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1AuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1AuctionCreated represents a AuctionCreated event raised by the NFTAuctionV1 contract.
type NFTAuctionV1AuctionCreated struct {
	AuctionId        *big.Int
	Seller           common.Address
	NftContract      common.Address
	TokenId          *big.Int
	StartPrice       *big.Int
	StartTime        *big.Int
	EndTime          *big.Int
	IsToken          bool
	AllowedTokens    []*big.Int
	StartDecimals    uint8
	StartTokenAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xceccea861d7a4eb82880a4b362d1bb992cde178297f05bf37ec8bbc83c026630.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime, bool isToken, uint256[] allowedTokens, uint8 startDecimals, uint256 startTokenAmount)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterAuctionCreated(opts *bind.FilterOpts, auctionId []*big.Int, seller []common.Address, nftContract []common.Address) (*NFTAuctionV1AuctionCreatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1AuctionCreatedIterator{contract: _NFTAuctionV1.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xceccea861d7a4eb82880a4b362d1bb992cde178297f05bf37ec8bbc83c026630.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime, bool isToken, uint256[] allowedTokens, uint8 startDecimals, uint256 startTokenAmount)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1AuctionCreated, auctionId []*big.Int, seller []common.Address, nftContract []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1AuctionCreated)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionCreated is a log parse operation binding the contract event 0xceccea861d7a4eb82880a4b362d1bb992cde178297f05bf37ec8bbc83c026630.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime, bool isToken, uint256[] allowedTokens, uint8 startDecimals, uint256 startTokenAmount)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseAuctionCreated(log types.Log) (*NFTAuctionV1AuctionCreated, error) {
	event := new(NFTAuctionV1AuctionCreated)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1AuctionEndIterator is returned from FilterAuctionEnd and is used to iterate over the raw logs and unpacked data for AuctionEnd events raised by the NFTAuctionV1 contract.
type NFTAuctionV1AuctionEndIterator struct {
	Event *NFTAuctionV1AuctionEnd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1AuctionEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1AuctionEnd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1AuctionEnd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1AuctionEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1AuctionEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1AuctionEnd represents a AuctionEnd event raised by the NFTAuctionV1 contract.
type NFTAuctionV1AuctionEnd struct {
	AuctionId   *big.Int
	Operator    common.Address
	Winner      common.Address
	Price       *big.Int
	EndTime     *big.Int
	Token       *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnd is a free log retrieval operation binding the contract event 0xebc8ae6394094f18cd0d0d2fd804c09ff165f7938fff11dbd938ae93842bed4c.
//
// Solidity: event AuctionEnd(uint256 indexed auctionId, address indexed operator, address indexed winner, uint256 price, uint256 endTime, uint256 token, uint256 tokenAmount)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterAuctionEnd(opts *bind.FilterOpts, auctionId []*big.Int, operator []common.Address, winner []common.Address) (*NFTAuctionV1AuctionEndIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "AuctionEnd", auctionIdRule, operatorRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1AuctionEndIterator{contract: _NFTAuctionV1.contract, event: "AuctionEnd", logs: logs, sub: sub}, nil
}

// WatchAuctionEnd is a free log subscription operation binding the contract event 0xebc8ae6394094f18cd0d0d2fd804c09ff165f7938fff11dbd938ae93842bed4c.
//
// Solidity: event AuctionEnd(uint256 indexed auctionId, address indexed operator, address indexed winner, uint256 price, uint256 endTime, uint256 token, uint256 tokenAmount)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchAuctionEnd(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1AuctionEnd, auctionId []*big.Int, operator []common.Address, winner []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "AuctionEnd", auctionIdRule, operatorRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1AuctionEnd)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "AuctionEnd", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionEnd is a log parse operation binding the contract event 0xebc8ae6394094f18cd0d0d2fd804c09ff165f7938fff11dbd938ae93842bed4c.
//
// Solidity: event AuctionEnd(uint256 indexed auctionId, address indexed operator, address indexed winner, uint256 price, uint256 endTime, uint256 token, uint256 tokenAmount)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseAuctionEnd(log types.Log) (*NFTAuctionV1AuctionEnd, error) {
	event := new(NFTAuctionV1AuctionEnd)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "AuctionEnd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1BidInfoIterator is returned from FilterBidInfo and is used to iterate over the raw logs and unpacked data for BidInfo events raised by the NFTAuctionV1 contract.
type NFTAuctionV1BidInfoIterator struct {
	Event *NFTAuctionV1BidInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1BidInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1BidInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1BidInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1BidInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1BidInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1BidInfo represents a BidInfo event raised by the NFTAuctionV1 contract.
type NFTAuctionV1BidInfo struct {
	AuctionId   *big.Int
	Bidder      common.Address
	BidTime     *big.Int
	BidPrice    *big.Int
	Token       *big.Int
	TokenAmount *big.Int
	BidDecimals uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBidInfo is a free log retrieval operation binding the contract event 0x6b66738e7d64789fd6830f6e3b1147ac5354fd8d379d27a45215cb88be4d4077.
//
// Solidity: event BidInfo(uint256 indexed auctionId, address indexed bidder, uint256 bidTime, uint256 bidPrice, uint256 token, uint256 tokenAmount, uint8 bidDecimals)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterBidInfo(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address) (*NFTAuctionV1BidInfoIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "BidInfo", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1BidInfoIterator{contract: _NFTAuctionV1.contract, event: "BidInfo", logs: logs, sub: sub}, nil
}

// WatchBidInfo is a free log subscription operation binding the contract event 0x6b66738e7d64789fd6830f6e3b1147ac5354fd8d379d27a45215cb88be4d4077.
//
// Solidity: event BidInfo(uint256 indexed auctionId, address indexed bidder, uint256 bidTime, uint256 bidPrice, uint256 token, uint256 tokenAmount, uint8 bidDecimals)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchBidInfo(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1BidInfo, auctionId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "BidInfo", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1BidInfo)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "BidInfo", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidInfo is a log parse operation binding the contract event 0x6b66738e7d64789fd6830f6e3b1147ac5354fd8d379d27a45215cb88be4d4077.
//
// Solidity: event BidInfo(uint256 indexed auctionId, address indexed bidder, uint256 bidTime, uint256 bidPrice, uint256 token, uint256 tokenAmount, uint8 bidDecimals)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseBidInfo(log types.Log) (*NFTAuctionV1BidInfo, error) {
	event := new(NFTAuctionV1BidInfo)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "BidInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1DelTokenCfgIterator is returned from FilterDelTokenCfg and is used to iterate over the raw logs and unpacked data for DelTokenCfg events raised by the NFTAuctionV1 contract.
type NFTAuctionV1DelTokenCfgIterator struct {
	Event *NFTAuctionV1DelTokenCfg // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1DelTokenCfgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1DelTokenCfg)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1DelTokenCfg)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1DelTokenCfgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1DelTokenCfgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1DelTokenCfg represents a DelTokenCfg event raised by the NFTAuctionV1 contract.
type NFTAuctionV1DelTokenCfg struct {
	Caller      common.Address
	Token       *big.Int
	TokenAddr   common.Address
	FeedAddr    common.Address
	OperateTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelTokenCfg is a free log retrieval operation binding the contract event 0x2379824be8ff5ad32a33166e0e3cfb9544f26b7016ab1b092f4bd84548f64a79.
//
// Solidity: event DelTokenCfg(address indexed caller, uint256 indexed token, address tokenAddr, address feedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterDelTokenCfg(opts *bind.FilterOpts, caller []common.Address, token []*big.Int) (*NFTAuctionV1DelTokenCfgIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "DelTokenCfg", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1DelTokenCfgIterator{contract: _NFTAuctionV1.contract, event: "DelTokenCfg", logs: logs, sub: sub}, nil
}

// WatchDelTokenCfg is a free log subscription operation binding the contract event 0x2379824be8ff5ad32a33166e0e3cfb9544f26b7016ab1b092f4bd84548f64a79.
//
// Solidity: event DelTokenCfg(address indexed caller, uint256 indexed token, address tokenAddr, address feedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchDelTokenCfg(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1DelTokenCfg, caller []common.Address, token []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "DelTokenCfg", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1DelTokenCfg)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "DelTokenCfg", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelTokenCfg is a log parse operation binding the contract event 0x2379824be8ff5ad32a33166e0e3cfb9544f26b7016ab1b092f4bd84548f64a79.
//
// Solidity: event DelTokenCfg(address indexed caller, uint256 indexed token, address tokenAddr, address feedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseDelTokenCfg(log types.Log) (*NFTAuctionV1DelTokenCfg, error) {
	event := new(NFTAuctionV1DelTokenCfg)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "DelTokenCfg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1GetUSDByTokenIterator is returned from FilterGetUSDByToken and is used to iterate over the raw logs and unpacked data for GetUSDByToken events raised by the NFTAuctionV1 contract.
type NFTAuctionV1GetUSDByTokenIterator struct {
	Event *NFTAuctionV1GetUSDByToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1GetUSDByTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1GetUSDByToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1GetUSDByToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1GetUSDByTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1GetUSDByTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1GetUSDByToken represents a GetUSDByToken event raised by the NFTAuctionV1 contract.
type NFTAuctionV1GetUSDByToken struct {
	CurrTime      *big.Int
	Token         *big.Int
	TokenDecimals uint8
	UpdateTime    *big.Int
	RawPrice      *big.Int
	FeedDecimals  uint8
	Decimals      uint8
	Usd18Value    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGetUSDByToken is a free log retrieval operation binding the contract event 0x05da27f16f090f0b0180e84627a9a10d91cc8d8b5fa8307c89a3ddf6f92d4db8.
//
// Solidity: event GetUSDByToken(uint256 indexed currTime, uint256 indexed token, uint8 tokenDecimals, uint256 updateTime, int256 rawPrice, uint8 feedDecimals, uint8 decimals, uint256 usd18Value)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterGetUSDByToken(opts *bind.FilterOpts, currTime []*big.Int, token []*big.Int) (*NFTAuctionV1GetUSDByTokenIterator, error) {

	var currTimeRule []interface{}
	for _, currTimeItem := range currTime {
		currTimeRule = append(currTimeRule, currTimeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "GetUSDByToken", currTimeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1GetUSDByTokenIterator{contract: _NFTAuctionV1.contract, event: "GetUSDByToken", logs: logs, sub: sub}, nil
}

// WatchGetUSDByToken is a free log subscription operation binding the contract event 0x05da27f16f090f0b0180e84627a9a10d91cc8d8b5fa8307c89a3ddf6f92d4db8.
//
// Solidity: event GetUSDByToken(uint256 indexed currTime, uint256 indexed token, uint8 tokenDecimals, uint256 updateTime, int256 rawPrice, uint8 feedDecimals, uint8 decimals, uint256 usd18Value)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchGetUSDByToken(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1GetUSDByToken, currTime []*big.Int, token []*big.Int) (event.Subscription, error) {

	var currTimeRule []interface{}
	for _, currTimeItem := range currTime {
		currTimeRule = append(currTimeRule, currTimeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "GetUSDByToken", currTimeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1GetUSDByToken)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "GetUSDByToken", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGetUSDByToken is a log parse operation binding the contract event 0x05da27f16f090f0b0180e84627a9a10d91cc8d8b5fa8307c89a3ddf6f92d4db8.
//
// Solidity: event GetUSDByToken(uint256 indexed currTime, uint256 indexed token, uint8 tokenDecimals, uint256 updateTime, int256 rawPrice, uint8 feedDecimals, uint8 decimals, uint256 usd18Value)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseGetUSDByToken(log types.Log) (*NFTAuctionV1GetUSDByToken, error) {
	event := new(NFTAuctionV1GetUSDByToken)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "GetUSDByToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NFTAuctionV1 contract.
type NFTAuctionV1InitializedIterator struct {
	Event *NFTAuctionV1Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1Initialized represents a Initialized event raised by the NFTAuctionV1 contract.
type NFTAuctionV1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*NFTAuctionV1InitializedIterator, error) {

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1InitializedIterator{contract: _NFTAuctionV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1Initialized) (event.Subscription, error) {

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1Initialized)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "Initialized", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseInitialized(log types.Log) (*NFTAuctionV1Initialized, error) {
	event := new(NFTAuctionV1Initialized)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTAuctionV1 contract.
type NFTAuctionV1OwnershipTransferredIterator struct {
	Event *NFTAuctionV1OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1OwnershipTransferred represents a OwnershipTransferred event raised by the NFTAuctionV1 contract.
type NFTAuctionV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTAuctionV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1OwnershipTransferredIterator{contract: _NFTAuctionV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1OwnershipTransferred)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseOwnershipTransferred(log types.Log) (*NFTAuctionV1OwnershipTransferred, error) {
	event := new(NFTAuctionV1OwnershipTransferred)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1RefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the NFTAuctionV1 contract.
type NFTAuctionV1RefundIterator struct {
	Event *NFTAuctionV1Refund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1RefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1Refund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1Refund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1RefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1RefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1Refund represents a Refund event raised by the NFTAuctionV1 contract.
type NFTAuctionV1Refund struct {
	AuctionId  *big.Int
	Withdrawee common.Address
	Token      *big.Int
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xaba4210cad11f1112f31ba136a61cdc0f6ab4c2b478da8c93eec1d9769e97a89.
//
// Solidity: event Refund(uint256 indexed auctionId, address indexed withdrawee, uint256 indexed token, uint256 price)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterRefund(opts *bind.FilterOpts, auctionId []*big.Int, withdrawee []common.Address, token []*big.Int) (*NFTAuctionV1RefundIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var withdraweeRule []interface{}
	for _, withdraweeItem := range withdrawee {
		withdraweeRule = append(withdraweeRule, withdraweeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "Refund", auctionIdRule, withdraweeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1RefundIterator{contract: _NFTAuctionV1.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xaba4210cad11f1112f31ba136a61cdc0f6ab4c2b478da8c93eec1d9769e97a89.
//
// Solidity: event Refund(uint256 indexed auctionId, address indexed withdrawee, uint256 indexed token, uint256 price)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1Refund, auctionId []*big.Int, withdrawee []common.Address, token []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var withdraweeRule []interface{}
	for _, withdraweeItem := range withdrawee {
		withdraweeRule = append(withdraweeRule, withdraweeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "Refund", auctionIdRule, withdraweeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1Refund)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "Refund", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefund is a log parse operation binding the contract event 0xaba4210cad11f1112f31ba136a61cdc0f6ab4c2b478da8c93eec1d9769e97a89.
//
// Solidity: event Refund(uint256 indexed auctionId, address indexed withdrawee, uint256 indexed token, uint256 price)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseRefund(log types.Log) (*NFTAuctionV1Refund, error) {
	event := new(NFTAuctionV1Refund)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NFTAuctionV1 contract.
type NFTAuctionV1TransferIterator struct {
	Event *NFTAuctionV1Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1Transfer represents a Transfer event raised by the NFTAuctionV1 contract.
type NFTAuctionV1Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NFTAuctionV1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1TransferIterator{contract: _NFTAuctionV1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1Transfer)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "Transfer", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseTransfer(log types.Log) (*NFTAuctionV1Transfer, error) {
	event := new(NFTAuctionV1Transfer)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1UpdCfgFeedAddrIterator is returned from FilterUpdCfgFeedAddr and is used to iterate over the raw logs and unpacked data for UpdCfgFeedAddr events raised by the NFTAuctionV1 contract.
type NFTAuctionV1UpdCfgFeedAddrIterator struct {
	Event *NFTAuctionV1UpdCfgFeedAddr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1UpdCfgFeedAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1UpdCfgFeedAddr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1UpdCfgFeedAddr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1UpdCfgFeedAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1UpdCfgFeedAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1UpdCfgFeedAddr represents a UpdCfgFeedAddr event raised by the NFTAuctionV1 contract.
type NFTAuctionV1UpdCfgFeedAddr struct {
	Caller      common.Address
	Token       *big.Int
	OldFeedAddr common.Address
	NewFeedAddr common.Address
	OperateTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdCfgFeedAddr is a free log retrieval operation binding the contract event 0x11749f4114942dc57cfd11d5d845989eac701c4067d2e6d03ce98e0b161bf492.
//
// Solidity: event UpdCfgFeedAddr(address indexed caller, uint256 indexed token, address oldFeedAddr, address newFeedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterUpdCfgFeedAddr(opts *bind.FilterOpts, caller []common.Address, token []*big.Int) (*NFTAuctionV1UpdCfgFeedAddrIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "UpdCfgFeedAddr", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1UpdCfgFeedAddrIterator{contract: _NFTAuctionV1.contract, event: "UpdCfgFeedAddr", logs: logs, sub: sub}, nil
}

// WatchUpdCfgFeedAddr is a free log subscription operation binding the contract event 0x11749f4114942dc57cfd11d5d845989eac701c4067d2e6d03ce98e0b161bf492.
//
// Solidity: event UpdCfgFeedAddr(address indexed caller, uint256 indexed token, address oldFeedAddr, address newFeedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchUpdCfgFeedAddr(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1UpdCfgFeedAddr, caller []common.Address, token []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "UpdCfgFeedAddr", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1UpdCfgFeedAddr)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "UpdCfgFeedAddr", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdCfgFeedAddr is a log parse operation binding the contract event 0x11749f4114942dc57cfd11d5d845989eac701c4067d2e6d03ce98e0b161bf492.
//
// Solidity: event UpdCfgFeedAddr(address indexed caller, uint256 indexed token, address oldFeedAddr, address newFeedAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseUpdCfgFeedAddr(log types.Log) (*NFTAuctionV1UpdCfgFeedAddr, error) {
	event := new(NFTAuctionV1UpdCfgFeedAddr)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "UpdCfgFeedAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1UpdCfgTokenAddrIterator is returned from FilterUpdCfgTokenAddr and is used to iterate over the raw logs and unpacked data for UpdCfgTokenAddr events raised by the NFTAuctionV1 contract.
type NFTAuctionV1UpdCfgTokenAddrIterator struct {
	Event *NFTAuctionV1UpdCfgTokenAddr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1UpdCfgTokenAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1UpdCfgTokenAddr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1UpdCfgTokenAddr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1UpdCfgTokenAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1UpdCfgTokenAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1UpdCfgTokenAddr represents a UpdCfgTokenAddr event raised by the NFTAuctionV1 contract.
type NFTAuctionV1UpdCfgTokenAddr struct {
	Caller       common.Address
	Token        *big.Int
	OldTokenAddr common.Address
	NewTokenAddr common.Address
	OperateTime  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdCfgTokenAddr is a free log retrieval operation binding the contract event 0xb215a9d1c56860b9b3973ef8282aa762bce38cac65d0c74bf7db7afec4217296.
//
// Solidity: event UpdCfgTokenAddr(address indexed caller, uint256 indexed token, address oldTokenAddr, address newTokenAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterUpdCfgTokenAddr(opts *bind.FilterOpts, caller []common.Address, token []*big.Int) (*NFTAuctionV1UpdCfgTokenAddrIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "UpdCfgTokenAddr", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1UpdCfgTokenAddrIterator{contract: _NFTAuctionV1.contract, event: "UpdCfgTokenAddr", logs: logs, sub: sub}, nil
}

// WatchUpdCfgTokenAddr is a free log subscription operation binding the contract event 0xb215a9d1c56860b9b3973ef8282aa762bce38cac65d0c74bf7db7afec4217296.
//
// Solidity: event UpdCfgTokenAddr(address indexed caller, uint256 indexed token, address oldTokenAddr, address newTokenAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchUpdCfgTokenAddr(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1UpdCfgTokenAddr, caller []common.Address, token []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "UpdCfgTokenAddr", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1UpdCfgTokenAddr)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "UpdCfgTokenAddr", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdCfgTokenAddr is a log parse operation binding the contract event 0xb215a9d1c56860b9b3973ef8282aa762bce38cac65d0c74bf7db7afec4217296.
//
// Solidity: event UpdCfgTokenAddr(address indexed caller, uint256 indexed token, address oldTokenAddr, address newTokenAddr, uint256 operateTime)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseUpdCfgTokenAddr(log types.Log) (*NFTAuctionV1UpdCfgTokenAddr, error) {
	event := new(NFTAuctionV1UpdCfgTokenAddr)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "UpdCfgTokenAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionV1UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NFTAuctionV1 contract.
type NFTAuctionV1UpgradedIterator struct {
	Event *NFTAuctionV1Upgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionV1UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionV1Upgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionV1Upgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionV1UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionV1UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionV1Upgraded represents a Upgraded event raised by the NFTAuctionV1 contract.
type NFTAuctionV1Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NFTAuctionV1UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionV1UpgradedIterator{contract: _NFTAuctionV1.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NFTAuctionV1Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NFTAuctionV1.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionV1Upgraded)
				if err := _NFTAuctionV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuctionV1 *NFTAuctionV1Filterer) ParseUpgraded(log types.Log) (*NFTAuctionV1Upgraded, error) {
	event := new(NFTAuctionV1Upgraded)
	if err := _NFTAuctionV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
