package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"nft-go/mytestnft"
	"nft-go/nftauction"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 1. 创建拍卖 --createAuction --privateKey <your-private-key> --nftContract <nft-contract-address> --tokenId <your-token-id> --startPrice <start-price> --startTime <start-time> --durationHours <duration-hours> --allowedTokens <allowed-tokens>
// 2. 拍卖出价 --bidAuction
func main() {
	rpcUrl := os.Getenv("ETH_RPC_URL")
	if rpcUrl == "" {
		log.Fatal("ETH_RPC_URL is not set")
	}
	fmt.Println("ETH_RPC_URL is set to:", rpcUrl)

	createAuctionFlag := flag.Bool("createAuction", false, "Create a new NFT auction")
	bidAuctionFlag := flag.Bool("bidAuction", false, "Bid on an auction")
	endAuctionFlag := flag.Bool("endAuction", false, "End an auction")
	refundFlag := flag.Bool("refund", false, "Refund from an auction")

	// startPrice、amount 使用字符串类型来接收，而不是数字类型，因为数字类型 Int64 只能最大表示 9.2 ETH , 超过范围会报错，而字符串类型则没有这个bug
	// 同时 cmd 命令行写参数都是用数字，只是接收的时候用字符串来接收
	privateKey := flag.String("privateKey", "", "private key (hex, no 0x)")
	nftContract := flag.String("nftContract", "", "NFT contract address")
	tokenId := flag.Int64("tokenId", 1, "NFT token ID")
	startPrice := flag.String("startPrice", "100000000", "start price in wei")
	startTime := flag.Int64("startTime", 0, "seconds from now to start (0=immediately)")
	durationHours := flag.Int64("durationHours", 24, "duration in hours")
	allowedTokens := flag.String("allowedTokens", "0", "allowed tokens, comma separated (e.g. 0,1,2)")
	auctionId := flag.Int64("auctionId", 0, "auction ID")
	token := flag.Int64("token", 0, "payment token (0=ETH)")
	amount := flag.String("amount", "200000000", "bid amount in wei")
	proxy := flag.String("proxy", "0x0405234edde4f86846c838e7b2bf60de97376459", "auction proxy address")
	testNFTContract := flag.String("testNFTContract", "0x814DACF409d581D4abDF2C38909B82eFF7b7bD1e", "test NFT contract address")

	flag.Parse()

	if !*createAuctionFlag && !*bidAuctionFlag && !*endAuctionFlag && !*refundFlag {
		fmt.Println("Usage:")
		fmt.Println("--createAuction --privateKey <key> --nftContract <addr> --tokenId <id,def:1> --startPrice <wei,def:100000000> --startTime <secondsFromNow,def:0=immediately> --durationHours <hours,def:24> --allowedTokens <csv,def:0> --proxy <addr,def:0x0405...> --testNFTContract <addr,def:0x814DA...>")
		fmt.Println("--bidAuction --privateKey <key> --auctionId <id> --token <0=ETH,def:0> --amount <wei,def:200000000> --proxy <addr,def:0x0405...>")
		fmt.Println("--endAuction --privateKey <key> --auctionId <id> --proxy <addr,def:0x0405...>")
		fmt.Println("--refund --privateKey <key> --auctionId <id> --token <0=ETH,def:0> --proxy <addr,def:0x0405...>")
		os.Exit(1)
	}

	mustSet(*privateKey, "--privateKey")
	privKey := parseKey(*privateKey)
	proxyAddr := common.HexToAddress(*proxy)
	testNFTAddr := common.HexToAddress(*testNFTContract)
	client := dialEth(rpcUrl)
	defer client.Close()
	auction := loadAuction(proxyAddr, client)
	nft := loadTestNFT(testNFTAddr, client)

	switch {
	case *createAuctionFlag:
		mustSet(*nftContract, "--nftContract")
		doCreateAuction(client, nft, auction, privKey, proxyAddr, *nftContract, *tokenId, *startPrice, *startTime, *durationHours, *allowedTokens)
	case *bidAuctionFlag:
		doBidAuction(client, auction, privKey, *auctionId, *token, *amount)
	case *endAuctionFlag:
		doEndAuction(client, auction, privKey, *auctionId)
	case *refundFlag:
		doRefund(client, auction, privKey, *auctionId, *token)
	}
}

func doCreateAuction(client *ethclient.Client, nft *mytestnft.MyTestNFT, auction *nftauction.NFTAuctionV1, privKey *ecdsa.PrivateKey, proxyAddr common.Address, nftAddr string, tokenId int64, startPriceStr string, startTime int64, durationHours int64, allowedTokensStr string) {
	callerAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	fmt.Printf("Caller: %s\n", callerAddr.Hex())

	auth := getAuth(client, privKey)
	fmt.Println("Approving auction contract to transfer NFT...")
	tx, err := nft.SetApprovalForAll(auth, proxyAddr, true)
	checkErr(err)
	fmt.Printf("Approve tx: %s\n", tx.Hash().Hex())

	var receipt *types.Receipt
	receipt = waitReceipt(client, tx.Hash())
	if receipt.Status != 1 {
		log.Fatalf("SetApprovalForAll failed! tx: %s", tx.Hash().Hex())
	}
	fmt.Println("Approve confirmed")

	price, ok := new(big.Int).SetString(startPriceStr, 10)
	if !ok {
		log.Fatal("invalid startPrice")
	}

	// 每笔交易创建新的 auth，nonce 自动递增
	auth = getAuth(client, privKey)
	params := nftauction.NFTAuctionV1CreateAuctionParams{
		NftContract:   common.HexToAddress(nftAddr),
		TokenId:       big.NewInt(tokenId),
		StartPrice:    price,
		StartTime:     calcStartTime(startTime),
		DurationHours: big.NewInt(durationHours),
		AllowedTokens: parseTokenList(allowedTokensStr),
	}

	fmt.Println("Creating auction...")
	tx, err = auction.CreateAuction(auth, params) // 仅发送交易，不等待确认
	checkErr(err)
	fmt.Printf("CreateAuction tx: %s\n", tx.Hash().Hex())

	receipt = waitReceipt(client, tx.Hash())
	if receipt.Status != 1 {
		log.Fatalf("CreateAuction failed! receipt status: %d, tx: %s", receipt.Status, tx.Hash().Hex())
	}
	fmt.Printf("CreateAuction confirmed in block %d\n", receipt.BlockNumber.Uint64())

	auctionCount, err := auction.GetAuctionCount(&bind.CallOpts{Context: context.Background()})
	checkErr(err)
	fmt.Printf("Auction created! AuctionId: %s\n", auctionCount.String())
}

func doBidAuction(client *ethclient.Client, auction *nftauction.NFTAuctionV1, privKey *ecdsa.PrivateKey, auctionId int64, token int64, amountStr string) {
	bidAmount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		log.Fatal("invalid amount")
	}

	auth := getAuth(client, privKey)
	auth.Value = bidAmount

	fmt.Printf("Bidding on auction %d with %s wei...\n", auctionId, bidAmount.String())
	tx, err := auction.BidAuction(auth, big.NewInt(auctionId), big.NewInt(token), bidAmount)
	checkErr(err)
	fmt.Printf("BidAuction tx: %s\n", tx.Hash().Hex())

	receipt := waitReceipt(client, tx.Hash())
	if receipt.Status != 1 {
		log.Fatalf("BidAuction failed! tx: %s", tx.Hash().Hex())
	}
	fmt.Printf("BidAuction confirmed in block %d\n", receipt.BlockNumber.Uint64())

	info, err := auction.GetAuctionInfo(&bind.CallOpts{Context: context.Background()}, big.NewInt(auctionId))
	checkErr(err)
	fmt.Printf("Highest bidder: %s, Highest price: %s\n", info.HighestBidder.Hex(), info.CurrHighestPrice.String())
}

func doEndAuction(client *ethclient.Client, auction *nftauction.NFTAuctionV1, privKey *ecdsa.PrivateKey, auctionId int64) {
	auth := getAuth(client, privKey)
	fmt.Printf("Ending auction %d...\n", auctionId)
	tx, err := auction.EndAuction(auth, big.NewInt(auctionId))
	checkErr(err)
	fmt.Printf("EndAuction tx: %s\n", tx.Hash().Hex())

	receipt := waitReceipt(client, tx.Hash())
	if receipt.Status != 1 {
		log.Fatalf("EndAuction failed! tx: %s", tx.Hash().Hex())
	}
	fmt.Printf("EndAuction confirmed in block %d\n", receipt.BlockNumber.Uint64())

	info, err := auction.GetAuctionInfo(&bind.CallOpts{Context: context.Background()}, big.NewInt(auctionId))
	checkErr(err)
	fmt.Printf("Auction ended: isEnded=%v\n", info.IsEnded)
}

func doRefund(client *ethclient.Client, auction *nftauction.NFTAuctionV1, privKey *ecdsa.PrivateKey, auctionId int64, token int64) {
	callerAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	bidPrice, err := auction.GetBidPriceReturns(&bind.CallOpts{Context: context.Background()}, big.NewInt(auctionId), callerAddr, big.NewInt(token))
	checkErr(err)
	fmt.Printf("Refund amount: %s wei\n", bidPrice.String())

	auth := getAuth(client, privKey)
	fmt.Printf("Refunding from auction %d for token %d...\n", auctionId, token)
	tx, err := auction.Refund(auth, big.NewInt(auctionId), big.NewInt(token))
	checkErr(err)
	fmt.Printf("Refund tx: %s\n", tx.Hash().Hex())

	receipt := waitReceipt(client, tx.Hash())
	if receipt.Status != 1 {
		log.Fatalf("Refund failed! tx: %s", tx.Hash().Hex())
	}
	fmt.Printf("Refund confirmed in block %d, refunded %s wei\n", receipt.BlockNumber.Uint64(), bidPrice.String())
}

func waitReceipt(client *ethclient.Client, txHash common.Hash) *types.Receipt {
	fmt.Printf("Waiting for tx %s to be on-chain...\n", txHash.Hex())
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt
		}
		time.Sleep(2 * time.Second)
	}
}

func dialEth(rpcUrl string) *ethclient.Client {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client:", err)
	}
	return client
}

func loadAuction(addr common.Address, client *ethclient.Client) *nftauction.NFTAuctionV1 {
	auction, err := nftauction.NewNFTAuctionV1(addr, client)
	if err != nil {
		log.Fatal("Failed to load auction contract:", err)
	}
	return auction
}

func loadTestNFT(addr common.Address, client *ethclient.Client) *mytestnft.MyTestNFT {
	nft, err := mytestnft.NewMyTestNFT(addr, client)
	if err != nil {
		log.Fatal("Failed to load test NFT contract:", err)
	}
	return nft
}

func parseKey(hex string) *ecdsa.PrivateKey {
	key, err := crypto.HexToECDSA(hex)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}
	return key
}

func getAuth(client *ethclient.Client, privKey *ecdsa.PrivateKey) *bind.TransactOpts {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to get chain ID:", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatal("Failed to create auth:", err)
	}
	return auth
}

func calcStartTime(offsetSeconds int64) *big.Int {
	if offsetSeconds == 0 {
		return big.NewInt(0)
	}
	return big.NewInt(time.Now().Unix() + offsetSeconds)
}

func parseTokenList(s string) []*big.Int {
	parts := strings.Split(s, ",")
	result := make([]*big.Int, len(parts))
	for i, p := range parts {
		val, ok := new(big.Int).SetString(strings.TrimSpace(p), 10)
		if !ok {
			log.Fatalf("invalid token value: %s", p)
		}
		result[i] = val
	}
	return result
}

func mustSet(val string, name string) {
	if val == "" {
		log.Fatalf("%s is required", name)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
