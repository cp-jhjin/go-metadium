// Copyright 2018-2022 The go-metadium / go-metadium Authors

package miner

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

var (
	ErrNotInitialized = errors.New("not initialized")

	IsMinerFunc                 func() bool
	AmPartnerFunc               func() bool
	IsPartnerFunc               func(string) bool
	AmHubFunc                   func(string) int
	LogBlockFunc                func(int64, common.Hash)
	CalculateRewardsFunc        func(*big.Int, *big.Int, *big.Int, func(common.Address, *big.Int)) (*common.Address, []byte, error)
	VerifyRewardsFunc           func(*big.Int, string) error
	GetCoinbaseFunc             func(height *big.Int) (coinbase common.Address, err error)
	SignBlockFunc               func(height *big.Int, hash common.Hash, isPangyo bool) (coinbase common.Address, nodeId, sig []byte, err error)
	VerifyBlockSigFunc          func(height *big.Int, coinbase common.Address, nodeId []byte, hash common.Hash, sig []byte, isPangyo bool) bool
	RequirePendingTxsFunc       func() bool
	VerifyBlockRewardsFunc      func(height *big.Int) interface{}
	SuggestGasPriceFunc         func() *big.Int
	GetBlockBuildParametersFunc func(height *big.Int) (blockInterval int64, maxBaseFee, gasLimit *big.Int, baseFeeMaxChangeRate, gasTargetPercentage int64, err error)
	AcquireMiningTokenFunc      func(height *big.Int, parentHash common.Hash) (bool, error)
	ReleaseMiningTokenFunc      func(height *big.Int, hash, parentHash common.Hash) error
	HasMiningTokenFunc          func() bool
)

func IsMiner() bool {
	if IsMinerFunc == nil {
		return false
	} else {
		return IsMinerFunc()
	}
}

func IsPartner(id string) bool {
	if IsPartnerFunc == nil {
		return false
	} else {
		return IsPartnerFunc(id)
	}
}

func AmPartner() bool {
	if AmPartnerFunc == nil {
		return false
	} else {
		return AmPartnerFunc()
	}
}

func AmHub(id string) int {
	if AmHubFunc == nil {
		return -1
	} else {
		return AmHubFunc(id)
	}
}

func LogBlock(height int64, hash common.Hash) {
	if LogBlockFunc != nil {
		LogBlockFunc(height, hash)
	}
}

func AcquireMiningToken(height *big.Int, parentHash common.Hash) (bool, error) {
	if AcquireMiningTokenFunc == nil {
		return false, ErrNotInitialized
	}
	return AcquireMiningTokenFunc(height, parentHash)
}

func ReleaseMiningToken(height *big.Int, hash, parentHash common.Hash) error {
	if ReleaseMiningTokenFunc == nil {
		return ErrNotInitialized
	}
	return ReleaseMiningTokenFunc(height, hash, parentHash)
}

func HasMiningToken() bool {
	if HasMiningTokenFunc == nil {
		return false
	}
	return HasMiningTokenFunc()
}

func IsPoW() bool {
	return params.ConsensusMethod == params.ConsensusPoW
}

func CalculateRewards(num, blockReward, fees *big.Int, addBalance func(common.Address, *big.Int)) (*common.Address, []byte, error) {
	if CalculateRewardsFunc == nil {
		return nil, nil, ErrNotInitialized
	} else {
		return CalculateRewardsFunc(num, blockReward, fees, addBalance)
	}
}

func VerifyRewards(num *big.Int, rewards string) error {
	if VerifyRewardsFunc == nil {
		return ErrNotInitialized
	} else {
		return VerifyRewardsFunc(num, rewards)
	}
}

func GetCoinbase(height *big.Int) (coinbase common.Address, err error) {
	if GetCoinbaseFunc == nil {
		err = ErrNotInitialized
	} else {
		coinbase, err = GetCoinbaseFunc(height)
	}
	return
}

func SignBlock(height *big.Int, hash common.Hash, isPangyo bool) (coinbase common.Address, nodeId, sig []byte, err error) {
	if SignBlockFunc == nil {
		err = ErrNotInitialized
	} else {
		coinbase, nodeId, sig, err = SignBlockFunc(height, hash, isPangyo)
	}
	return
}

func VerifyBlockSig(height *big.Int, coinbase common.Address, nodeId []byte, hash common.Hash, sig []byte, isPangyo bool) bool {
	if VerifyBlockSigFunc == nil {
		return false
	} else {
		return VerifyBlockSigFunc(height, coinbase, nodeId, hash, sig, isPangyo)
	}
}

func RequirePendingTxs() bool {
	if RequirePendingTxsFunc == nil {
		return false
	} else {
		return RequirePendingTxsFunc()
	}
}

func VerifyBlockRewards(height *big.Int) interface{} {
	if VerifyBlockRewardsFunc == nil {
		return false
	} else {
		return VerifyBlockRewardsFunc(height)
	}
}

func SuggestGasPrice() *big.Int {
	if SuggestGasPriceFunc == nil {
		return big.NewInt(100 * params.GWei)
	} else {
		return SuggestGasPriceFunc()
	}
}

func GetBlockBuildParameters(height *big.Int) (blockInterval int64, maxBaseFee, gasLimit *big.Int, baseFeeMaxChangeRate, gasTargetPercentage int64, err error) {
	if GetBlockBuildParametersFunc == nil {
		// default values
		return 15, big.NewInt(0), big.NewInt(0), 0, 100, ErrNotInitialized
	} else {
		return GetBlockBuildParametersFunc(height)
	}
}

// EOF
