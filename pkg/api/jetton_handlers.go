package api

import (
	"context"
	"errors"
	"math/big"

	"github.com/tonkeeper/opentonapi/pkg/bath"
	"github.com/tonkeeper/tongo"

	"github.com/tonkeeper/opentonapi/pkg/core"
	"github.com/tonkeeper/opentonapi/pkg/oas"
)

func (h Handler) GetJettonsBalances(ctx context.Context, params oas.GetJettonsBalancesParams) (oas.GetJettonsBalancesRes, error) {
	accountID, err := tongo.ParseAccountID(params.AccountID)
	if err != nil {
		return &oas.BadRequest{Error: err.Error()}, nil
	}
	wallets, err := h.storage.GetJettonWalletsByOwnerAddress(ctx, accountID)
	if err != nil {
		if errors.Is(err, core.ErrEntityNotFound) {
			return &oas.NotFound{Error: err.Error()}, nil
		}
		return &oas.InternalError{Error: err.Error()}, nil
	}
	var balances = oas.JettonsBalances{
		Balances: make([]oas.JettonBalance, 0, len(wallets)),
	}
	for _, wallet := range wallets {
		jettonBalance := oas.JettonBalance{
			Balance:       wallet.Balance.String(),
			WalletAddress: convertAccountAddress(wallet.Address, h.addressBook),
		}
		meta, err := h.storage.GetJettonMasterMetadata(ctx, wallet.JettonAddress)
		if err != nil && !errors.Is(err, core.ErrEntityNotFound) {
			return &oas.InternalError{Error: err.Error()}, nil
		}
		jettonBalance.Jetton = jettonPreview(h.addressBook, wallet.JettonAddress, meta, h.previewGenerator)
		balances.Balances = append(balances.Balances, jettonBalance)
	}

	return &balances, nil
}

func (h Handler) GetJettonInfo(ctx context.Context, params oas.GetJettonInfoParams) (r oas.GetJettonInfoRes, err error) {
	verification := oas.JettonVerificationTypeNone
	account, err := tongo.ParseAccountID(params.AccountID)
	if err != nil {
		return &oas.BadRequest{Error: err.Error()}, nil
	}
	meta, err := h.storage.GetJettonMasterMetadata(ctx, account)
	metadata := oas.JettonMetadata{Address: account.ToRaw()}
	info, ok := h.addressBook.GetJettonInfoByAddress(account)
	if ok {
		meta.Name = rewriteIfNotEmpty(meta.Name, info.Name)
		meta.Description = rewriteIfNotEmpty(meta.Description, info.Description)
		meta.Image = rewriteIfNotEmpty(meta.Image, info.Image)
		meta.Symbol = rewriteIfNotEmpty(meta.Symbol, info.Symbol)
		verification = oas.JettonVerificationTypeWhitelist
	}
	metadata.Name = meta.Name
	metadata.Symbol = meta.Symbol
	metadata.Decimals = meta.Decimals
	metadata.Social = info.Social
	metadata.Websites = info.Websites

	if meta.Description != "" {
		metadata.Description.SetTo(meta.Description)
	}
	if meta.Image != "" {
		metadata.Image.SetTo(meta.Image)
	}
	data, err := h.storage.GetJettonMasterData(ctx, account)
	if err != nil {
		return &oas.InternalError{Error: err.Error()}, nil
	}
	supply := big.Int(data.TotalSupply)
	return &oas.JettonInfo{
		Mintable:     data.Mintable != 0,
		TotalSupply:  supply.String(),
		Metadata:     metadata,
		Verification: verification,
	}, nil
}

func (h Handler) GetJettonsHistory(ctx context.Context, params oas.GetJettonsHistoryParams) (res oas.GetJettonsHistoryRes, err error) {
	account, err := tongo.ParseAccountID(params.AccountID)
	if err != nil {
		return &oas.BadRequest{Error: err.Error()}, nil
	}
	traceIDs, err := h.storage.GetAccountJettonsHistory(ctx, account, params.Limit, optIntToPointer(params.BeforeLt), optIntToPointer(params.StartDate), optIntToPointer(params.EndDate))
	if err != nil {
		return &oas.InternalError{Error: err.Error()}, nil
	}
	events := make([]oas.AccountEvent, len(traceIDs))
	var lastLT uint64
	for idx, traceID := range traceIDs {
		trace, err := h.storage.GetTrace(ctx, traceID)
		if err != nil {
			return &oas.InternalError{Error: err.Error()}, nil
		}
		bubble := bath.FromTrace(trace)
		actions, fees := bath.CollectActions(bubble, &account)
		event := oas.AccountEvent{
			EventID:    trace.Hash.Hex(),
			Account:    convertAccountAddress(account, h.addressBook),
			Timestamp:  trace.Utime,
			Fee:        oas.Fee{Account: convertAccountAddress(account, h.addressBook)},
			IsScam:     false,
			Lt:         int64(trace.Lt),
			InProgress: trace.InProgress(),
		}
		for _, fee := range fees {
			if fee.WhoPay == account {
				event.Fee = convertFees(fee, h.addressBook)
				break
			}
		}
		for _, action := range actions {
			convertedAction, spamDetected := h.convertAction(ctx, action)
			if !event.IsScam && spamDetected {
				event.IsScam = true
			}
			event.Actions = append(event.Actions, convertedAction)
		}
		if len(event.Actions) == 0 {
			event.Actions = []oas.Action{{
				Type:   oas.ActionTypeUnknown,
				Status: oas.ActionStatusOk,
			}}
		}
		events[idx] = event
		lastLT = trace.Lt
	}

	return &oas.AccountEvents{Events: events, NextFrom: int64(lastLT)}, nil
}
