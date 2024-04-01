// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AccountDnsBackResolve implements accountDnsBackResolve operation.
	//
	// Get account's domains.
	//
	// GET /v2/accounts/{account_id}/dns/backresolve
	AccountDnsBackResolve(ctx context.Context, params AccountDnsBackResolveParams) (*DomainNames, error)
	// AddressParse implements addressParse operation.
	//
	// Parse address and display in all formats.
	//
	// GET /v2/address/{account_id}/parse
	AddressParse(ctx context.Context, params AddressParseParams) (*AddressParseOK, error)
	// BlockchainAccountInspect implements blockchainAccountInspect operation.
	//
	// Blockchain account inspect.
	//
	// GET /v2/blockchain/accounts/{account_id}/inspect
	BlockchainAccountInspect(ctx context.Context, params BlockchainAccountInspectParams) (*BlockchainAccountInspect, error)
	// DecodeMessage implements decodeMessage operation.
	//
	// Decode a given message. Only external incoming messages can be decoded currently.
	//
	// POST /v2/message/decode
	DecodeMessage(ctx context.Context, req *DecodeMessageReq) (*DecodedMessage, error)
	// DnsResolve implements dnsResolve operation.
	//
	// DNS resolve for domain name.
	//
	// GET /v2/dns/{domain_name}/resolve
	DnsResolve(ctx context.Context, params DnsResolveParams) (*DnsRecord, error)
	// EmulateMessageToAccountEvent implements emulateMessageToAccountEvent operation.
	//
	// Emulate sending message to blockchain.
	//
	// POST /v2/accounts/{account_id}/events/emulate
	EmulateMessageToAccountEvent(ctx context.Context, req *EmulateMessageToAccountEventReq, params EmulateMessageToAccountEventParams) (*AccountEvent, error)
	// EmulateMessageToEvent implements emulateMessageToEvent operation.
	//
	// Emulate sending message to blockchain.
	//
	// POST /v2/events/emulate
	EmulateMessageToEvent(ctx context.Context, req *EmulateMessageToEventReq, params EmulateMessageToEventParams) (*Event, error)
	// EmulateMessageToTrace implements emulateMessageToTrace operation.
	//
	// Emulate sending message to blockchain.
	//
	// POST /v2/traces/emulate
	EmulateMessageToTrace(ctx context.Context, req *EmulateMessageToTraceReq, params EmulateMessageToTraceParams) (*Trace, error)
	// EmulateMessageToWallet implements emulateMessageToWallet operation.
	//
	// Emulate sending message to blockchain.
	//
	// POST /v2/wallet/emulate
	EmulateMessageToWallet(ctx context.Context, req *EmulateMessageToWalletReq, params EmulateMessageToWalletParams) (*MessageConsequences, error)
	// ExecGetMethodForBlockchainAccount implements execGetMethodForBlockchainAccount operation.
	//
	// Execute get method for account.
	//
	// GET /v2/blockchain/accounts/{account_id}/methods/{method_name}
	ExecGetMethodForBlockchainAccount(ctx context.Context, params ExecGetMethodForBlockchainAccountParams) (*MethodExecutionResult, error)
	// GetAccount implements getAccount operation.
	//
	// Get human-friendly information about an account without low-level details.
	//
	// GET /v2/accounts/{account_id}
	GetAccount(ctx context.Context, params GetAccountParams) (*Account, error)
	// GetAccountDiff implements getAccountDiff operation.
	//
	// Get account's balance change.
	//
	// GET /v2/accounts/{account_id}/diff
	GetAccountDiff(ctx context.Context, params GetAccountDiffParams) (*GetAccountDiffOK, error)
	// GetAccountDnsExpiring implements getAccountDnsExpiring operation.
	//
	// Get expiring account .ton dns.
	//
	// GET /v2/accounts/{account_id}/dns/expiring
	GetAccountDnsExpiring(ctx context.Context, params GetAccountDnsExpiringParams) (*DnsExpiring, error)
	// GetAccountEvent implements getAccountEvent operation.
	//
	// Get event for an account by event_id.
	//
	// GET /v2/accounts/{account_id}/events/{event_id}
	GetAccountEvent(ctx context.Context, params GetAccountEventParams) (*AccountEvent, error)
	// GetAccountEvents implements getAccountEvents operation.
	//
	// Get events for an account. Each event is built on top of a trace which is a series of transactions
	// caused by one inbound message. TonAPI looks for known patterns inside the trace and splits the
	// trace into actions, where a single action represents a meaningful high-level operation like a
	// Jetton Transfer or an NFT Purchase. Actions are expected to be shown to users. It is advised not
	// to build any logic on top of actions because actions can be changed at any time.
	//
	// GET /v2/accounts/{account_id}/events
	GetAccountEvents(ctx context.Context, params GetAccountEventsParams) (*AccountEvents, error)
	// GetAccountInfoByStateInit implements getAccountInfoByStateInit operation.
	//
	// Get account info by state init.
	//
	// POST /v2/tonconnect/stateinit
	GetAccountInfoByStateInit(ctx context.Context, req *GetAccountInfoByStateInitReq) (*AccountInfoByStateInit, error)
	// GetAccountInscriptions implements getAccountInscriptions operation.
	//
	// Get all inscriptions by owner address. It's experimental API and can be dropped in the future.
	//
	// GET /v2/experimental/accounts/{account_id}/inscriptions
	GetAccountInscriptions(ctx context.Context, params GetAccountInscriptionsParams) (*InscriptionBalances, error)
	// GetAccountInscriptionsHistory implements getAccountInscriptionsHistory operation.
	//
	// Get the transfer inscriptions history for account. It's experimental API and can be dropped in the
	// future.
	//
	// GET /v2/experimental/accounts/{account_id}/inscriptions/history
	GetAccountInscriptionsHistory(ctx context.Context, params GetAccountInscriptionsHistoryParams) (*AccountEvents, error)
	// GetAccountInscriptionsHistoryByTicker implements getAccountInscriptionsHistoryByTicker operation.
	//
	// Get the transfer inscriptions history for account. It's experimental API and can be dropped in the
	// future.
	//
	// GET /v2/experimental/accounts/{account_id}/inscriptions/{ticker}/history
	GetAccountInscriptionsHistoryByTicker(ctx context.Context, params GetAccountInscriptionsHistoryByTickerParams) (*AccountEvents, error)
	// GetAccountJettonHistoryByID implements getAccountJettonHistoryByID operation.
	//
	// Get the transfer jetton history for account and jetton.
	//
	// GET /v2/accounts/{account_id}/jettons/{jetton_id}/history
	GetAccountJettonHistoryByID(ctx context.Context, params GetAccountJettonHistoryByIDParams) (*AccountEvents, error)
	// GetAccountJettonsBalances implements getAccountJettonsBalances operation.
	//
	// Get all Jettons balances by owner address.
	//
	// GET /v2/accounts/{account_id}/jettons
	GetAccountJettonsBalances(ctx context.Context, params GetAccountJettonsBalancesParams) (*JettonsBalances, error)
	// GetAccountJettonsHistory implements getAccountJettonsHistory operation.
	//
	// Get the transfer jettons history for account.
	//
	// GET /v2/accounts/{account_id}/jettons/history
	GetAccountJettonsHistory(ctx context.Context, params GetAccountJettonsHistoryParams) (*AccountEvents, error)
	// GetAccountNftHistory implements getAccountNftHistory operation.
	//
	// Get the transfer nft history.
	//
	// GET /v2/accounts/{account_id}/nfts/history
	GetAccountNftHistory(ctx context.Context, params GetAccountNftHistoryParams) (*AccountEvents, error)
	// GetAccountNftItems implements getAccountNftItems operation.
	//
	// Get all NFT items by owner address.
	//
	// GET /v2/accounts/{account_id}/nfts
	GetAccountNftItems(ctx context.Context, params GetAccountNftItemsParams) (*NftItems, error)
	// GetAccountNominatorsPools implements getAccountNominatorsPools operation.
	//
	// All pools where account participates.
	//
	// GET /v2/staking/nominator/{account_id}/pools
	GetAccountNominatorsPools(ctx context.Context, params GetAccountNominatorsPoolsParams) (*AccountStaking, error)
	// GetAccountPublicKey implements getAccountPublicKey operation.
	//
	// Get public key by account id.
	//
	// GET /v2/accounts/{account_id}/publickey
	GetAccountPublicKey(ctx context.Context, params GetAccountPublicKeyParams) (*GetAccountPublicKeyOK, error)
	// GetAccountSeqno implements getAccountSeqno operation.
	//
	// Get account seqno.
	//
	// GET /v2/wallet/{account_id}/seqno
	GetAccountSeqno(ctx context.Context, params GetAccountSeqnoParams) (*Seqno, error)
	// GetAccountSubscriptions implements getAccountSubscriptions operation.
	//
	// Get all subscriptions by wallet address.
	//
	// GET /v2/accounts/{account_id}/subscriptions
	GetAccountSubscriptions(ctx context.Context, params GetAccountSubscriptionsParams) (*Subscriptions, error)
	// GetAccountTraces implements getAccountTraces operation.
	//
	// Get traces for account.
	//
	// GET /v2/accounts/{account_id}/traces
	GetAccountTraces(ctx context.Context, params GetAccountTracesParams) (*TraceIDs, error)
	// GetAccounts implements getAccounts operation.
	//
	// Get human-friendly information about several accounts without low-level details.
	//
	// POST /v2/accounts/_bulk
	GetAccounts(ctx context.Context, req OptGetAccountsReq) (*Accounts, error)
	// GetAllAuctions implements getAllAuctions operation.
	//
	// Get all auctions.
	//
	// GET /v2/dns/auctions
	GetAllAuctions(ctx context.Context, params GetAllAuctionsParams) (*Auctions, error)
	// GetAllRawShardsInfo implements getAllRawShardsInfo operation.
	//
	// Get all raw shards info.
	//
	// GET /v2/liteserver/get_all_shards_info/{block_id}
	GetAllRawShardsInfo(ctx context.Context, params GetAllRawShardsInfoParams) (*GetAllRawShardsInfoOK, error)
	// GetBlockchainAccountTransactions implements getBlockchainAccountTransactions operation.
	//
	// Get account transactions.
	//
	// GET /v2/blockchain/accounts/{account_id}/transactions
	GetBlockchainAccountTransactions(ctx context.Context, params GetBlockchainAccountTransactionsParams) (*Transactions, error)
	// GetBlockchainBlock implements getBlockchainBlock operation.
	//
	// Get blockchain block data.
	//
	// GET /v2/blockchain/blocks/{block_id}
	GetBlockchainBlock(ctx context.Context, params GetBlockchainBlockParams) (*BlockchainBlock, error)
	// GetBlockchainBlockTransactions implements getBlockchainBlockTransactions operation.
	//
	// Get transactions from block.
	//
	// GET /v2/blockchain/blocks/{block_id}/transactions
	GetBlockchainBlockTransactions(ctx context.Context, params GetBlockchainBlockTransactionsParams) (*Transactions, error)
	// GetBlockchainConfig implements getBlockchainConfig operation.
	//
	// Get blockchain config.
	//
	// GET /v2/blockchain/config
	GetBlockchainConfig(ctx context.Context) (*BlockchainConfig, error)
	// GetBlockchainConfigFromBlock implements getBlockchainConfigFromBlock operation.
	//
	// Get blockchain config from a specific block, if present.
	//
	// GET /v2/blockchain/masterchain/{masterchain_seqno}/config
	GetBlockchainConfigFromBlock(ctx context.Context, params GetBlockchainConfigFromBlockParams) (*BlockchainConfig, error)
	// GetBlockchainMasterchainBlocks implements getBlockchainMasterchainBlocks operation.
	//
	// Get all blocks in all shards and workchains between target and previous masterchain block
	// according to shards last blocks snapshot in masterchain.  We don't recommend to build your app
	// around this method because it has problem with scalability and will work very slow in the future.
	//
	// GET /v2/blockchain/masterchain/{masterchain_seqno}/blocks
	GetBlockchainMasterchainBlocks(ctx context.Context, params GetBlockchainMasterchainBlocksParams) (*BlockchainBlocks, error)
	// GetBlockchainMasterchainHead implements getBlockchainMasterchainHead operation.
	//
	// Get last known masterchain block.
	//
	// GET /v2/blockchain/masterchain-head
	GetBlockchainMasterchainHead(ctx context.Context) (*BlockchainBlock, error)
	// GetBlockchainMasterchainShards implements getBlockchainMasterchainShards operation.
	//
	// Get blockchain block shards.
	//
	// GET /v2/blockchain/masterchain/{masterchain_seqno}/shards
	GetBlockchainMasterchainShards(ctx context.Context, params GetBlockchainMasterchainShardsParams) (*BlockchainBlockShards, error)
	// GetBlockchainMasterchainTransactions implements getBlockchainMasterchainTransactions operation.
	//
	// Get all transactions in all shards and workchains between target and previous masterchain block
	// according to shards last blocks snapshot in masterchain. We don't recommend to build your app
	// around this method because it has problem with scalability and will work very slow in the future.
	//
	// GET /v2/blockchain/masterchain/{masterchain_seqno}/transactions
	GetBlockchainMasterchainTransactions(ctx context.Context, params GetBlockchainMasterchainTransactionsParams) (*Transactions, error)
	// GetBlockchainRawAccount implements getBlockchainRawAccount operation.
	//
	// Get low-level information about an account taken directly from the blockchain.
	//
	// GET /v2/blockchain/accounts/{account_id}
	GetBlockchainRawAccount(ctx context.Context, params GetBlockchainRawAccountParams) (*BlockchainRawAccount, error)
	// GetBlockchainTransaction implements getBlockchainTransaction operation.
	//
	// Get transaction data.
	//
	// GET /v2/blockchain/transactions/{transaction_id}
	GetBlockchainTransaction(ctx context.Context, params GetBlockchainTransactionParams) (*Transaction, error)
	// GetBlockchainTransactionByMessageHash implements getBlockchainTransactionByMessageHash operation.
	//
	// Get transaction data by message hash.
	//
	// GET /v2/blockchain/messages/{msg_id}/transaction
	GetBlockchainTransactionByMessageHash(ctx context.Context, params GetBlockchainTransactionByMessageHashParams) (*Transaction, error)
	// GetBlockchainValidators implements getBlockchainValidators operation.
	//
	// Get blockchain validators.
	//
	// GET /v2/blockchain/validators
	GetBlockchainValidators(ctx context.Context) (*Validators, error)
	// GetChartRates implements getChartRates operation.
	//
	// Get chart by token.
	//
	// GET /v2/rates/chart
	GetChartRates(ctx context.Context, params GetChartRatesParams) (*GetChartRatesOK, error)
	// GetDnsInfo implements getDnsInfo operation.
	//
	// Get full information about domain name.
	//
	// GET /v2/dns/{domain_name}
	GetDnsInfo(ctx context.Context, params GetDnsInfoParams) (*DomainInfo, error)
	// GetDomainBids implements getDomainBids operation.
	//
	// Get domain bids.
	//
	// GET /v2/dns/{domain_name}/bids
	GetDomainBids(ctx context.Context, params GetDomainBidsParams) (*DomainBids, error)
	// GetEvent implements getEvent operation.
	//
	// Get an event either by event ID or a hash of any transaction in a trace. An event is built on top
	// of a trace which is a series of transactions caused by one inbound message. TonAPI looks for known
	// patterns inside the trace and splits the trace into actions, where a single action represents a
	// meaningful high-level operation like a Jetton Transfer or an NFT Purchase. Actions are expected to
	// be shown to users. It is advised not to build any logic on top of actions because actions can be
	// changed at any time.
	//
	// GET /v2/events/{event_id}
	GetEvent(ctx context.Context, params GetEventParams) (*Event, error)
	// GetInscriptionOpTemplate implements getInscriptionOpTemplate operation.
	//
	// Return comment for making operation with inscription. please don't use it if you don't know what
	// you are doing.
	//
	// GET /v2/experimental/inscriptions/op-template
	GetInscriptionOpTemplate(ctx context.Context, params GetInscriptionOpTemplateParams) (*GetInscriptionOpTemplateOK, error)
	// GetItemsFromCollection implements getItemsFromCollection operation.
	//
	// Get NFT items from collection by collection address.
	//
	// GET /v2/nfts/collections/{account_id}/items
	GetItemsFromCollection(ctx context.Context, params GetItemsFromCollectionParams) (*NftItems, error)
	// GetJettonHolders implements getJettonHolders operation.
	//
	// Get jetton's holders.
	//
	// GET /v2/jettons/{account_id}/holders
	GetJettonHolders(ctx context.Context, params GetJettonHoldersParams) (*JettonHolders, error)
	// GetJettonInfo implements getJettonInfo operation.
	//
	// Get jetton metadata by jetton master address.
	//
	// GET /v2/jettons/{account_id}
	GetJettonInfo(ctx context.Context, params GetJettonInfoParams) (*JettonInfo, error)
	// GetJettons implements getJettons operation.
	//
	// Get a list of all indexed jetton masters in the blockchain.
	//
	// GET /v2/jettons
	GetJettons(ctx context.Context, params GetJettonsParams) (*Jettons, error)
	// GetJettonsEvents implements getJettonsEvents operation.
	//
	// Get only jetton transfers in the event.
	//
	// GET /v2/events/{event_id}/jettons
	GetJettonsEvents(ctx context.Context, params GetJettonsEventsParams) (*Event, error)
	// GetMarketsRates implements getMarketsRates operation.
	//
	// Get the TON price from markets.
	//
	// GET /v2/rates/markets
	GetMarketsRates(ctx context.Context) (*GetMarketsRatesOK, error)
	// GetNftCollection implements getNftCollection operation.
	//
	// Get NFT collection by collection address.
	//
	// GET /v2/nfts/collections/{account_id}
	GetNftCollection(ctx context.Context, params GetNftCollectionParams) (*NftCollection, error)
	// GetNftCollections implements getNftCollections operation.
	//
	// Get NFT collections.
	//
	// GET /v2/nfts/collections
	GetNftCollections(ctx context.Context, params GetNftCollectionsParams) (*NftCollections, error)
	// GetNftHistoryByID implements getNftHistoryByID operation.
	//
	// Get the transfer nfts history for account.
	//
	// GET /v2/nfts/{account_id}/history
	GetNftHistoryByID(ctx context.Context, params GetNftHistoryByIDParams) (*AccountEvents, error)
	// GetNftItemByAddress implements getNftItemByAddress operation.
	//
	// Get NFT item by its address.
	//
	// GET /v2/nfts/{account_id}
	GetNftItemByAddress(ctx context.Context, params GetNftItemByAddressParams) (*NftItem, error)
	// GetNftItemsByAddresses implements getNftItemsByAddresses operation.
	//
	// Get NFT items by their addresses.
	//
	// POST /v2/nfts/_bulk
	GetNftItemsByAddresses(ctx context.Context, req OptGetNftItemsByAddressesReq) (*NftItems, error)
	// GetRates implements getRates operation.
	//
	// Get the token price to the currency.
	//
	// GET /v2/rates
	GetRates(ctx context.Context, params GetRatesParams) (*GetRatesOK, error)
	// GetRawAccountState implements getRawAccountState operation.
	//
	// Get raw account state.
	//
	// GET /v2/liteserver/get_account_state/{account_id}
	GetRawAccountState(ctx context.Context, params GetRawAccountStateParams) (*GetRawAccountStateOK, error)
	// GetRawBlockProof implements getRawBlockProof operation.
	//
	// Get raw block proof.
	//
	// GET /v2/liteserver/get_block_proof
	GetRawBlockProof(ctx context.Context, params GetRawBlockProofParams) (*GetRawBlockProofOK, error)
	// GetRawBlockchainBlock implements getRawBlockchainBlock operation.
	//
	// Get raw blockchain block.
	//
	// GET /v2/liteserver/get_block/{block_id}
	GetRawBlockchainBlock(ctx context.Context, params GetRawBlockchainBlockParams) (*GetRawBlockchainBlockOK, error)
	// GetRawBlockchainBlockHeader implements getRawBlockchainBlockHeader operation.
	//
	// Get raw blockchain block header.
	//
	// GET /v2/liteserver/get_block_header/{block_id}
	GetRawBlockchainBlockHeader(ctx context.Context, params GetRawBlockchainBlockHeaderParams) (*GetRawBlockchainBlockHeaderOK, error)
	// GetRawBlockchainBlockState implements getRawBlockchainBlockState operation.
	//
	// Get raw blockchain block state.
	//
	// GET /v2/liteserver/get_state/{block_id}
	GetRawBlockchainBlockState(ctx context.Context, params GetRawBlockchainBlockStateParams) (*GetRawBlockchainBlockStateOK, error)
	// GetRawBlockchainConfig implements getRawBlockchainConfig operation.
	//
	// Get raw blockchain config.
	//
	// GET /v2/blockchain/config/raw
	GetRawBlockchainConfig(ctx context.Context) (*RawBlockchainConfig, error)
	// GetRawBlockchainConfigFromBlock implements getRawBlockchainConfigFromBlock operation.
	//
	// Get raw blockchain config from a specific block, if present.
	//
	// GET /v2/blockchain/masterchain/{masterchain_seqno}/config/raw
	GetRawBlockchainConfigFromBlock(ctx context.Context, params GetRawBlockchainConfigFromBlockParams) (*RawBlockchainConfig, error)
	// GetRawConfig implements getRawConfig operation.
	//
	// Get raw config.
	//
	// GET /v2/liteserver/get_config_all/{block_id}
	GetRawConfig(ctx context.Context, params GetRawConfigParams) (*GetRawConfigOK, error)
	// GetRawListBlockTransactions implements getRawListBlockTransactions operation.
	//
	// Get raw list block transactions.
	//
	// GET /v2/liteserver/list_block_transactions/{block_id}
	GetRawListBlockTransactions(ctx context.Context, params GetRawListBlockTransactionsParams) (*GetRawListBlockTransactionsOK, error)
	// GetRawMasterchainInfo implements getRawMasterchainInfo operation.
	//
	// Get raw masterchain info.
	//
	// GET /v2/liteserver/get_masterchain_info
	GetRawMasterchainInfo(ctx context.Context) (*GetRawMasterchainInfoOK, error)
	// GetRawMasterchainInfoExt implements getRawMasterchainInfoExt operation.
	//
	// Get raw masterchain info ext.
	//
	// GET /v2/liteserver/get_masterchain_info_ext
	GetRawMasterchainInfoExt(ctx context.Context, params GetRawMasterchainInfoExtParams) (*GetRawMasterchainInfoExtOK, error)
	// GetRawShardBlockProof implements getRawShardBlockProof operation.
	//
	// Get raw shard block proof.
	//
	// GET /v2/liteserver/get_shard_block_proof/{block_id}
	GetRawShardBlockProof(ctx context.Context, params GetRawShardBlockProofParams) (*GetRawShardBlockProofOK, error)
	// GetRawShardInfo implements getRawShardInfo operation.
	//
	// Get raw shard info.
	//
	// GET /v2/liteserver/get_shard_info/{block_id}
	GetRawShardInfo(ctx context.Context, params GetRawShardInfoParams) (*GetRawShardInfoOK, error)
	// GetRawTime implements getRawTime operation.
	//
	// Get raw time.
	//
	// GET /v2/liteserver/get_time
	GetRawTime(ctx context.Context) (*GetRawTimeOK, error)
	// GetRawTransactions implements getRawTransactions operation.
	//
	// Get raw transactions.
	//
	// GET /v2/liteserver/get_transactions/{account_id}
	GetRawTransactions(ctx context.Context, params GetRawTransactionsParams) (*GetRawTransactionsOK, error)
	// GetStakingPoolHistory implements getStakingPoolHistory operation.
	//
	// Pool history.
	//
	// GET /v2/staking/pool/{account_id}/history
	GetStakingPoolHistory(ctx context.Context, params GetStakingPoolHistoryParams) (*GetStakingPoolHistoryOK, error)
	// GetStakingPoolInfo implements getStakingPoolInfo operation.
	//
	// Stacking pool info.
	//
	// GET /v2/staking/pool/{account_id}
	GetStakingPoolInfo(ctx context.Context, params GetStakingPoolInfoParams) (*GetStakingPoolInfoOK, error)
	// GetStakingPools implements getStakingPools operation.
	//
	// All pools available in network.
	//
	// GET /v2/staking/pools
	GetStakingPools(ctx context.Context, params GetStakingPoolsParams) (*GetStakingPoolsOK, error)
	// GetStorageProviders implements getStorageProviders operation.
	//
	// Get TON storage providers deployed to the blockchain.
	//
	// GET /v2/storage/providers
	GetStorageProviders(ctx context.Context) (*GetStorageProvidersOK, error)
	// GetTonConnectPayload implements getTonConnectPayload operation.
	//
	// Get a payload for further token receipt.
	//
	// GET /v2/tonconnect/payload
	GetTonConnectPayload(ctx context.Context) (*GetTonConnectPayloadOK, error)
	// GetTrace implements getTrace operation.
	//
	// Get the trace by trace ID or hash of any transaction in trace.
	//
	// GET /v2/traces/{trace_id}
	GetTrace(ctx context.Context, params GetTraceParams) (*Trace, error)
	// GetWalletBackup implements getWalletBackup operation.
	//
	// Get backup info.
	//
	// GET /v2/wallet/backup
	GetWalletBackup(ctx context.Context, params GetWalletBackupParams) (*GetWalletBackupOK, error)
	// GetWalletsByPublicKey implements getWalletsByPublicKey operation.
	//
	// Get wallets by public key.
	//
	// GET /v2/pubkeys/{public_key}/wallets
	GetWalletsByPublicKey(ctx context.Context, params GetWalletsByPublicKeyParams) (*Accounts, error)
	// ReduceIndexingLatency implements reduceIndexingLatency operation.
	//
	// Reduce indexing latency.
	//
	// GET /v2/status
	ReduceIndexingLatency(ctx context.Context) (*ServiceStatus, error)
	// ReindexAccount implements reindexAccount operation.
	//
	// Update internal cache for a particular account.
	//
	// POST /v2/accounts/{account_id}/reindex
	ReindexAccount(ctx context.Context, params ReindexAccountParams) error
	// SearchAccounts implements searchAccounts operation.
	//
	// Search by account domain name.
	//
	// GET /v2/accounts/search
	SearchAccounts(ctx context.Context, params SearchAccountsParams) (*FoundAccounts, error)
	// SendBlockchainMessage implements sendBlockchainMessage operation.
	//
	// Send message to blockchain.
	//
	// POST /v2/blockchain/message
	SendBlockchainMessage(ctx context.Context, req *SendBlockchainMessageReq) error
	// SendRawMessage implements sendRawMessage operation.
	//
	// Send raw message to blockchain.
	//
	// POST /v2/liteserver/send_message
	SendRawMessage(ctx context.Context, req *SendRawMessageReq) (*SendRawMessageOK, error)
	// SetWalletBackup implements setWalletBackup operation.
	//
	// Set backup info.
	//
	// PUT /v2/wallet/backup
	SetWalletBackup(ctx context.Context, req SetWalletBackupReq, params SetWalletBackupParams) error
	// TonConnectProof implements tonConnectProof operation.
	//
	// Account verification and token issuance.
	//
	// POST /v2/wallet/auth/proof
	TonConnectProof(ctx context.Context, req *TonConnectProofReq) (*TonConnectProofOK, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
