# NFTAuctionV1 合约的 Go CLI

Go 命令行工具，用于与 Sepolia 测试网或主网上的 NFTAuctionV1 合约交互，支持创建拍卖、出价、结束拍卖、提取退款等操作。

## 项目结构

```
nft-go/
├── main.go                  # CLI 入口，命令行参数解析与合约交互逻辑
├── nftauction/
│   └── NFTAuctionV1.go      # abigen 生成的 NFTAuctionV1 合约 Go 绑定
├── mytestnft/
│   └── MyTestNFT.go         # abigen 生成的 MyTestNFT 合约 Go 绑定
├── go.mod
└── go.sum
```

## 前置条件

- Go 1.25+
- 设置环境变量 `ETH_RPC_URL`（Sepolia RPC 节点地址）

```bash
export ETH_RPC_URL=https://sepolia.infura.io/v3/<your-infura-key>
```

## 合约地址（Sepolia）

| 合约 | 地址 |
|------|------|
| NFTAuctionV1 Proxy | `0x0405234edde4f86846c838e7b2bf60de97376459` |
| MyTestNFT | `0x814DACF409d581D4abDF2C38909B82eFF7b7bD1e` |

## 使用方法

### 创建拍卖

```bash
go run main.go --createAuction \
  --privateKey <your-private-key> \
  --nftContract 0x814DACF409d581D4abDF2C38909B82eFF7b7bD1e \
  --tokenId 1 \
  --startPrice 100000000 \
  --startTime 0 \
  --durationHours 24 \
  --allowedTokens 0,1,2
```

流程：
1. 在 NFT 合约上调用 `setApprovalForAll`，授权拍卖合约转移 NFT
2. 调用 `createAuction` 创建拍卖，NFT 从卖家转移到拍卖合约
3. 等待交易上链确认，打印 AuctionId

### 拍卖出价

```bash
go run main.go --bidAuction \
  --privateKey <bidder-private-key> \
  --auctionId 1 \
  --token 0 \
  --amount 200000000
```

- `--token 0` 表示用 ETH 出价
- `--amount` 为出价金额，单位 wei

### 结束拍卖

```bash
go run main.go --endAuction \
  --privateKey <seller-private-key> \
  --auctionId 1
```

拍卖到期后，卖家调用结束拍卖，NFT 转给最高出价者，出价资金结算。

### 提取退款

```bash
go run main.go --refund \
  --privateKey <bidder-private-key> \
  --auctionId 1 \
  --token 0
```

未中标的出价者可提取退款。执行前会查询并打印可退金额。

## 参数说明

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `--privateKey` | string | （必填） | 私钥，hex 格式，不带 0x 前缀 |
| `--nftContract` | string | （createAuction 必填） | NFT 合约地址 |
| `--tokenId` | int64 | `1` | NFT Token ID |
| `--startPrice` | string | `"100000000"` | 起拍价，单位 wei（字符串类型避免 int64 溢出） |
| `--startTime` | int64 | `0` | 距今多少秒后开始，0 表示立即开始 |
| `--durationHours` | int64 | `24` | 拍卖持续时长，单位小时 |
| `--allowedTokens` | string | `"0"` | 允许的支付 token，逗号分隔（0=ETH） |
| `--auctionId` | int64 | `0` | 拍卖 ID |
| `--token` | int64 | `0` | 支付 token（0=ETH） |
| `--amount` | string | `"200000000"` | 出价金额，单位 wei |
| `--proxy` | string | `0x0405...` | 拍卖合约 Proxy 地址 |
| `--testNFTContract` | string | `0x814D...` | MyTestNFT 合约地址 |

## 关键设计

### 交易确认机制

每笔交易发送后，会轮询 `TransactionReceipt` 等待交易上链，并检查 `receipt.Status`：
- `Status = 1`：交易成功
- `Status = 0`：交易 revert，打印 tx hash 供 Etherscan 排查

### Nonce 管理

每笔交易前调用 `getAuth()` 创建新的 `TransactOpts`，由 go-ethereum 自动获取下一个 nonce，避免连续交易时 nonce 冲突。

### Gas 估算

不设置 `GasLimit`，由节点通过 `eth_estimateGas` 自动估算，避免硬编码 gas 不够导致 `out of gas` revert。

### 大数处理

`startPrice` 和 `amount` 使用字符串类型接收，通过 `new(big.Int).SetString()` 解析为 `*big.Int`，避免 int64 溢出（int64 最大约 9.2 ETH）。

## 生成 Go 绑定

```bash
# NFTAuctionV1
abigen --abi NFTAuctionV1.abi.json --pkg nftauction --out nftauction/NFTAuctionV1.go

# MyTestNFT
abigen --abi MyTestNFT.abi.json --pkg mytestnft --out mytestnft/MyTestNFT.go
```