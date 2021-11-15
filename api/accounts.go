package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-agent/message/npool"
	"github.com/NpoolPlatform/sphinx-agent/pkg/app"
)

func (s *Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (resp *npool.AccountBalance, err error) {
	resp, err = app.GetBalance(in.CoinId, in.Address, in.TimestampUtc)
	return
}
