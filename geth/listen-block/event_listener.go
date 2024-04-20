package listen_block

import (
	"context"
	"math/big"
	"time"

	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractListener interface {
	Listen()
}

func NewContract(client *ethclient.Client, address common.Address, abi abi.ABI) ContractListener {
	return &contract{
		ETHClient: client,
		Address:   address,
		ABI:       abi,
	}
}

type contract struct {
	ETHClient        *ethclient.Client
	Address          common.Address // 合约地址
	ABI              abi.ABI        // 合约对应的abi
	lastScannedBlock uint64         // 上次扫描的最新区块号
}

func (c *contract) Listen() {
	// 获取当前最新的区块号
	latestBlockNum, err := c.ETHClient.BlockNumber(context.Background())
	if err != nil {
		zap.L().Fatal("获取最新区块号失败", zap.Error(err))
	}
	c.lastScannedBlock = latestBlockNum
	// 定时扫描新的区块
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			zap.L().Info("开始扫描区块")
			ctx := context.Background()
			// 获取当前最新的区块号
			latestBlock, err := c.ETHClient.BlockNumber(ctx)
			if err != nil {
				zap.L().Error("获取最新区块号失败", zap.Error(err))
				continue
			}
			// 扫描从上次扫描以来的新区块
			for blockNumber := c.lastScannedBlock + 1; blockNumber <= latestBlock; blockNumber++ {
				// 根据区块号获取区块数据
				block, err := c.ETHClient.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
				if err != nil {
					zap.L().Error("获取区块失败", zap.Error(err), zap.Uint64("blockNumber", blockNumber))
					continue
				}
				// 处理区块中的transaction
				for _, transaction := range block.Transactions() {
					// 获取transaction收据
					receipt, err := c.ETHClient.TransactionReceipt(ctx, transaction.Hash())
					if err != nil {
						zap.L().Error("获取transaction收据失败", zap.Error(err),
							zap.Uint64("blockNumber", blockNumber), zap.Any("transaction", transaction))
						continue
					}
					// 解析log
					for _, log := range receipt.Logs {
						event := struct{}{} // 对应事件的结构体
						if err := c.ABI.UnpackIntoInterface(&event, "ERC20Recovered", log.Data); err != nil {
							e, _ := log.MarshalJSON()
							zap.L().Error("事件解析失败", zap.Error(err), zap.String("e", string(e)))
							continue
						}
					}
				}
			}
			zap.L().Info("等待下一次扫描")
		}
	}
}
