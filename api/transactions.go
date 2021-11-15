package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-agent/message/npool"
	"github.com/NpoolPlatform/sphinx-agent/pkg/app"
)

func (s *Server) GetTxStatus(ctx context.Context, in *npool.GetTxStatusRequest) (resp *npool.GetTxStatusResponse, err error) {
	resp, err = app.GetTxStatus(in.TransactionIdChain)
	return
}
