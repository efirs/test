
package model

import "github.com/google/uuid"
    
type SpendingOutpoint struct {
	N int64 `json:"n"`
	TxIndex int64 `json:"tx_index"`
}

type PrevOut struct {
	Addr string `json:"addr"`
	N int64 `json:"n"`
	Script string `json:"script"`
	SpendingOutpoints []SpendingOutpoint `json:"spending_outpoints"`
	Spent bool `json:"spent"`
	TxIndex int64 `json:"tx_index"`
	Type int64 `json:"type"`
	Value int64 `json:"value"`
}

type Input struct {
	Index int64 `json:"index"`
	PrevOut PrevOut `json:"prev_out"`
	Script string `json:"script"`
	Sequence int64 `json:"sequence"`
	Witness string `json:"witness"`
}

type Tx struct {
	DoubleSpend bool `json:"double_spend"`
	Fee int64 `json:"fee"`
	Hash []byte `json:"hash"`
	Inputs []Input `json:"inputs"`
	LockTime int64 `json:"lock_time"`
	Outs []PrevOut `json:"out"`
	Rbf bool `json:"rbf"`
	RelayedBy string `json:"relayed_by"`
	Size int64 `json:"size"`
	Time int64 `json:"time"`
	TxIndex int64 `json:"tx_index"`
	Ver int64 `json:"ver"`
	VinSz int64 `json:"vin_sz"`
	VoutSz int64 `json:"vout_sz"`
	Weight int64 `json:"weight"`
}

type Unconfirmed struct {
	DoubleSpend bool `json:"double_spend"`
	Fee int64 `json:"fee"`
	Hash []byte `json:"hash"`
	Id uuid.UUID `json:"id" tigris:"primaryKey:1,autoGenerate"`
	Inputs []Input `json:"inputs"`
	LockTime int64 `json:"lock_time"`
	Outs []PrevOut `json:"out"`
	Rbf bool `json:"rbf"`
	RelayedBy string `json:"relayed_by"`
	Size int64 `json:"size"`
	Time int64 `json:"time"`
	TxIndex int64 `json:"tx_index"`
	Txs []Tx `json:"txs"`
	Ver int64 `json:"ver"`
	VinSz int64 `json:"vin_sz"`
	VoutSz int64 `json:"vout_sz"`
	Weight int64 `json:"weight"`
}
