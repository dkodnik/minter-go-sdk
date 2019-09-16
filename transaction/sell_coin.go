package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type SellCoinData struct {
	CoinToSell        [10]byte
	ValueToSell       *big.Int
	CoinToBuy         [10]byte
	MinimumValueToBuy *big.Int
}

func NewSellCoinData() *SellCoinData {
	return &SellCoinData{}
}

func (d *SellCoinData) SetCoinToSell(symbol string) *SellCoinData {
	copy(d.CoinToSell[:], symbol)
	return d
}

func (d *SellCoinData) SetCoinToBuy(symbol string) *SellCoinData {
	copy(d.CoinToBuy[:], symbol)
	return d
}

func (d *SellCoinData) SetValueToSell(value *big.Int) *SellCoinData {
	d.ValueToSell = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *SellCoinData) SetMinimumValueToBuy(value *big.Int) *SellCoinData {
	d.MinimumValueToBuy = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *SellCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
