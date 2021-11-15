package grpcws

import (
	"context"

	plugin "github.com/NpoolPlatform/sphinx-plugin/message/npool"
	sign "github.com/NpoolPlatform/sphinx-sign/message/npool"
)

type NodePlugin interface {
	// 获取预签名信息
	GetSignInfo(ctx context.Context, in *plugin.GetSignInfoRequest) (resp *plugin.SignInfo, err error)
	// 余额查询
	GetBalance(ctx context.Context, in *plugin.GetBalanceRequest) (resp *plugin.AccountBalance, err error)
	// 广播交易
	BroadcastScript(ctx context.Context, in *plugin.BroadcastScriptRequest) (resp *plugin.BroadcastScriptResponse, err error)
	// 交易状态查询
	GetTxStatus(ctx context.Context, in *plugin.GetTxStatusRequest) (resp *plugin.GetTxStatusResponse, err error)
	// TODO 账户交易查询
	GetTxJSON(ctx context.Context, in *plugin.GetTxJSONRequest) (resp *plugin.AccountTxJSON, err error)
}

type NodeSign interface {
	CreateAccount(ctx context.Context, in *sign.CreateAccountRequest) (resp *sign.CreateAccountResponse, err error)
	SignScript(ctx context.Context, in *sign.SignScriptRequest) (resp *sign.SignScriptResponse, err error)
}
