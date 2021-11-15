package app

import (
	"github.com/NpoolPlatform/sphinx-agent/message/npool"
)

func GetTxStatus(transactionIDChain string) (resp *npool.GetTxStatusResponse, err error) {
	resp = &npool.GetTxStatusResponse{
		AmountFloat64:      0,
		AmountUint64:       0,
		AddressFrom:        "",
		AddressTo:          "",
		TransactionIdChain: transactionIDChain,
		CreatetimeUtc:      0,
		UpdatetimeUtc:      0,
		IsSuccess:          false,
		IsFailed:           false,
		NumBlocksConfirmed: 0,
	}
	return
}

func GetBalance(coinID int32, address string, timestampUTC int64) (resp *npool.AccountBalance, err error) {
	resp = &npool.AccountBalance{
		CoinId:        coinID,
		Address:       address,
		TimestampUtc:  timestampUTC,
		AmountFloat64: 0,
		AmountUint64:  0,
	}
	return
}
