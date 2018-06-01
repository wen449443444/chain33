package types

const (
	CoinsX    = "coins"
	TicketX   = "ticket"
	HashlockX = "hashlock"
	RetrieveX = "retrieve"
	NoneX     = "none"
	TokenX    = "token"
	TradeX    = "trade"
	ManageX   = "manage"
	PrivacyX  = "privacy"
)

var (
	AllowDepositExec = []string{"ticket"}
	//TODO:修改这种静态的方式为动态注册
	AllowUserExec          = []string{"coins", "ticket", "hashlock", "retrieve", "none", "token", "trade", "manage", "privacy"}
	GenesisAddr            = "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"
	GenesisBlockTime int64 = 1514533394
	HotkeyAddr             = "12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"
	FundKeyAddr            = "1BQXS6TxaYYG5mADaWij4AxhZZUTpw95a5"
	EmptyValue             = []byte("emptyBVBiCj5jvE15pEiwro8TQRGnJSNsJF") //这字符串表示数据库中的空值
	SuperManager           = []string{"1Bsg9j6gW83sShoee1fZAt9TkUjcrCgA9S", "1Q8hGLfoGe63efeWa8fJ4Pnukhkngt6poK"}
	ConfigPrefix           = "mavl-config-"
	TokenApprs             = []string{
		"1Bsg9j6gW83sShoee1fZAt9TkUjcrCgA9S",
		"1Q8hGLfoGe63efeWa8fJ4Pnukhkngt6poK",
		"1LY8GFia5EiyoTodMLfkB5PHNNpXRqxhyB",
		"1GCzJDS6HbgTQ2emade7mEJGGWFfA15pS9",
		"1JYB8sxi4He5pZWHCd3Zi2nypQ4JMB6AxN",
	}
	//addr:1Cbo5u8V5F3ubWBv9L6qu9wWxKuD3qBVpi,这里只是作为测试用，后面需要修改为系统账户
	ViewPubFee  = "0x0f7b661757fe8471c0b853b09bf526b19537a2f91254494d19874a04119415e8"
	SpendPubFee = "0x64204db5a521771eeeddee59c25aaae6bebe796d564effb6ba11352418002ee3"
	ViewPrivFee  = "0x0f7b661757fe8471c0b853b09bf526b19537a2f91254494d19874a04119415e8"
)

var (
	MinFee             int64 = 1e5
	MinBalanceTransfer int64 = 1e6
)

func SetMinFee(fee int64) {
	if fee < 0 {
		panic("fee less than zero")
	}
	MinFee = fee
	MinBalanceTransfer = fee * 10
}

const (
	//<<<<<<< HEAD
	//Coin                     int64 = 1e8
	//TokenPrecision           int64 = 1e4
	//InputPrecision           float64 = 1e4
	//CoinMultiple             int64 = 1e4
	//MaxCoin                  int64 = 1e17
	//FutureBlockTime          int64 = 16
	//CoinReward               int64 = 18 * Coin //用户回报
	//CoinDevFund              int64 = 12 * Coin //发展基金回报
	//TicketPrice              int64 = 10000 * Coin
	//TicketFrozenTime         int64 = 5        //5s only for test
	//TicketWithdrawTime       int64 = 10       //10s only for test
	//TicketMinerWaitTime      int64 = 2        // 2s only for test
	//MaxTxSize                int64 = 100000   //100K
	//MaxBlockSize             int64 = 10000000 //10M
	//MaxTxNumber              int64 = 1600     //160
	//PowLimitBits             uint32 = uint32(0x1f00ffff)
	//TargetTimespan                   = 144 * 16 * time.Second
	//TargetTimePerBlock               = 16 * time.Second
	//RetargetAdjustmentFactor = 4
	//MaxTxsPerBlock = 100000
	//TokenNameLenLimit = 128
	//TokenSymbolLenLimit              = 16
	//TokenIntroLenLimit = 1024
	//TokenPrecisionLen = 1e6
	//TokenCreatePriceStand = 10000 * Coin
	//InvalidStartTime = 0
	//InvalidStopTime = 0
	//====== =
	Coin                int64   = 1e8
	MaxCoin             int64   = 1e17
	MaxTxSize                   = 100000   //100K
	MaxBlockSize                = 10000000 //10M
	MaxTxsPerBlock              = 100000
	TokenPrecision      int64   = 1e8
	MaxTokenBalance     int64   = 900 * 1e8 * TokenPrecision //900亿
	InputPrecision      float64 = 1e4
	Multiple1E4         int64   = 1e4
	TokenNameLenLimit           = 128
	TokenSymbolLenLimit         = 16
	TokenIntroLenLimit          = 1024
	InvalidStartTime            = 0
	InvalidStopTime             = 0
	BTY                         = "BTY"
	BTYDustThreshold            = 1e8 * 10
	ConfirmedHeight             = 12
	UTXOCacheCount              = 256
)

const (
	EventTx                   = 1
	EventGetBlocks            = 2
	EventBlocks               = 3
	EventGetBlockHeight       = 4
	EventReplyBlockHeight     = 5
	EventQueryTx              = 6
	EventTransactionDetail    = 7
	EventReply                = 8
	EventTxBroadcast          = 9
	EventPeerInfo             = 10
	EventTxList               = 11
	EventReplyTxList          = 12
	EventAddBlock             = 13
	EventBlockBroadcast       = 14
	EventFetchBlocks          = 15
	EventAddBlocks            = 16
	EventTxHashList           = 17
	EventTxHashListReply      = 18
	EventGetHeaders           = 19
	EventHeaders              = 20
	EventGetMempoolSize       = 21
	EventMempoolSize          = 22
	EventStoreGet             = 23
	EventStoreSet             = 24
	EventStoreGetReply        = 25
	EventStoreSetReply        = 26
	EventReceipts             = 27
	EventExecTxList           = 28
	EventPeerList             = 29
	EventGetLastHeader        = 30
	EventHeader               = 31
	EventAddBlockDetail       = 32
	EventGetMempool           = 33
	EventGetTransactionByAddr = 34
	EventGetTransactionByHash = 35
	EventReplyTxInfo          = 36
	//wallet event
	EventWalletGetAccountList  = 37
	EventWalletAccountList     = 38
	EventNewAccount            = 39
	EventWalletAccount         = 40
	EventWalletTransactionList = 41
	//EventReplyTxList           = 42
	EventWalletImportprivkey = 43
	EventWalletSendToAddress = 44
	EventWalletSetFee        = 45
	EventWalletSetLabel      = 46
	//EventWalletAccount       = 47
	EventWalletMergeBalance = 48
	EventReplyHashes        = 49
	EventWalletSetPasswd    = 50
	EventWalletLock         = 51
	EventWalletUnLock       = 52
	EventTransactionDetails = 53
	EventBroadcastAddBlock  = 54
	EventGetBlockOverview   = 55
	EventGetAddrOverview    = 56
	EventReplyBlockOverview = 57
	EventReplyAddrOverview  = 58
	EventGetBlockHash       = 59
	EventBlockHash          = 60
	EventGetLastMempool     = 61
	EventWalletGetTickets   = 62
	EventMinerStart         = 63
	EventMinerStop          = 64
	EventWalletTickets      = 65
	EventStoreMemSet        = 66
	EventStoreRollback      = 67
	EventStoreCommit        = 68
	EventCheckBlock         = 69
	//seed
	EventGenSeed      = 70
	EventReplyGenSeed = 71
	EventSaveSeed     = 72
	EventGetSeed      = 73
	EventReplyGetSeed = 74
	EventDelBlock     = 75
	//local store
	EventLocalGet            = 76
	EventLocalReplyValue     = 77
	EventLocalList           = 78
	EventLocalSet            = 79
	EventGetWalletStatus     = 80
	EventCheckTx             = 81
	EventReceiptCheckTx      = 82
	EventQuery               = 83
	EventReplyQuery          = 84
	EventFlushTicket         = 85
	EventFetchBlockHeaders   = 86
	EventAddBlockHeaders     = 87
	EventWalletAutoMiner     = 88
	EventReplyWalletStatus   = 89
	EventGetLastBlock        = 90
	EventBlock               = 91
	EventGetTicketCount      = 92
	EventReplyGetTicketCount = 93
	EventDumpPrivkey         = 94
	EventReplyPrivkey        = 95
	EventIsSync              = 96
	EventReplyIsSync         = 97

	EventCloseTickets        = 98
	EventGetAddrTxs          = 99
	EventReplyAddrTxs        = 100
	EventIsNtpClockSync      = 101
	EventReplyIsNtpClockSync = 102
	EventDelTxList           = 103
	// Token
	EventTokenPreCreate         = 200
	EventReplyTokenPreCreate    = 201
	EventTokenFinishCreate      = 202
	EventReplyTokenFinishCreate = 203
	EventTokenRevokeCreate      = 204
	EventReplyTokenRevokeCreate = 205
	EventSellToken              = 206
	EventReplySellToken         = 207
	EventBuyToken               = 208
	EventReplyBuyToken          = 209
	EventRevokeSellToken        = 210
	EventReplyRevokeSellToken   = 211
	// config
	EventModifyConfig      = 300
	EventReplyModifyConfig = 301

	// privacy
	EventPublic2privacy = iota + 400
	EventReplyPublic2privacy
	EventPrivacy2privacy
	EventReplyPrivacy2privacy
	EventPrivacy2public
	EventReplyPrivacy2public
	EventShowPrivacyPK
	EventReplyShowPrivacyPK
	EventShowPrivacyTransfer
	EventReplyShowPrivacyTransfer
	EventShowPrivacyAccount
	EventReplyShowPrivacyAccount
	EventGetPrivacyTransaction
	EventReplyGetPrivacyTransaction
	EventGetGlobalIndex
	EventReplyGetGlobalIndex
	EventGetUTXOPubKey
	EventReplyGetUTXOPubKey
	EventCreateUTXOs
	EventReplyCreateUTXOs

)

var eventName = map[int]string{
	1:  "EventTx",
	2:  "EventGetBlocks",
	3:  "EventBlocks",
	4:  "EventGetBlockHeight",
	5:  "EventReplyBlockHeight",
	6:  "EventQueryTx",
	7:  "EventTransactionDetail",
	8:  "EventReply",
	9:  "EventTxBroadcast",
	10: "EventPeerInfo",
	11: "EventTxList",
	12: "EventReplyTxList",
	13: "EventAddBlock",
	14: "EventBlockBroadcast",
	15: "EventFetchBlocks",
	16: "EventAddBlocks",
	17: "EventTxHashList",
	18: "EventTxHashListReply",
	19: "EventGetHeaders",
	20: "EventHeaders",
	21: "EventGetMempoolSize",
	22: "EventMempoolSize",
	23: "EventStoreGet",
	24: "EventStoreSet",
	25: "EventStoreGetReply",
	26: "EventStoreSetReply",
	27: "EventReceipts",
	28: "EventExecTxList",
	29: "EventPeerList",
	30: "EventGetLastHeader",
	31: "EventHeader",
	32: "EventAddBlockDetail",
	33: "EventGetMempool",
	34: "EventGetTransactionByAddr",
	35: "EventGetTransactionByHash",
	36: "EventReplyTxInfo",
	37: "EventWalletGetAccountList",
	38: "EventWalletAccountList",
	39: "EventNewAccount",
	40: "EventWalletAccount",
	41: "EventWalletTransactionList",
	//42: "EventReplyTxList",
	43: "EventWalletImportPrivKey",
	44: "EventWalletSendToAddress",
	45: "EventWalletSetFee",
	46: "EventWalletSetLabel",
	//47: "EventWalletAccount",
	48:  "EventWalletMergeBalance",
	49:  "EventReplyHashes",
	50:  "EventWalletSetPasswd",
	51:  "EventWalletLock",
	52:  "EventWalletUnLock",
	53:  "EventTransactionDetails",
	54:  "EventBroadcastAddBlock",
	55:  "EventGetBlockOverview",
	56:  "EventGetAddrOverview",
	57:  "EventReplyBlockOverview",
	58:  "EventReplyAddrOverview",
	59:  "EventGetBlockHash",
	60:  "EventBlockHash",
	61:  "EventGetLastMempool",
	62:  "EventWalletGetTickets",
	63:  "EventMinerStart",
	64:  "EventMinerStop",
	65:  "EventWalletTickets",
	66:  "EventStoreMemSet",
	67:  "EventStoreRollback",
	68:  "EventStoreCommit",
	69:  "EventCheckBlock",
	70:  "EventGenSeed",
	71:  "EventReplyGenSeed",
	72:  "EventSaveSeed",
	73:  "EventGetSeed",
	74:  "EventReplyGetSeed",
	75:  "EventDelBlock",
	76:  "EventLocalGet",
	77:  "EventLocalReplyValue",
	78:  "EventLocalList",
	79:  "EventLocalSet",
	80:  "EventGetWalletStatus",
	81:  "EventCheckTx",
	82:  "EventReceiptCheckTx",
	83:  "EventQuery",
	84:  "EventReplyQuery",
	85:  "EventFlushTicket",
	86:  "EventFetchBlockHeaders",
	87:  "EventAddBlockHeaders",
	88:  "EventWalletAutoMiner",
	89:  "EventReplyWalletStatus",
	90:  "EventGetLastBlock",
	91:  "EventBlock",
	92:  "EventGetTicketCount",
	93:  "EventReplyGetTicketCount",
	94:  "EventDumpPrivkey",
	95:  "EventReplyPrivkey",
	96:  "EventIsSync",
	97:  "EventReplyIsSync",
	98:  "EventCloseTickets",
	99:  "EventGetAddrTxs",
	100: "EventReplyAddrTxs",
	101: "EventIsNtpClockSync",
	102: "EventReplyIsNtpClockSync",
	103: "EventDelTxList",
	// Token
	EventTokenPreCreate:         "EventTokenPreCreate",
	EventReplyTokenPreCreate:    "EventReplyTokenPreCreate",
	EventTokenFinishCreate:      "EventTokenFinishCreate",
	EventReplyTokenFinishCreate: "EventReplyTokenFinishCreate",
	EventTokenRevokeCreate:      "EventTokenRevokeCreate",
	EventReplyTokenRevokeCreate: "EventReplyTokenRevokeCreate",
	EventSellToken:              "EventSellToken",
	EventReplySellToken:         "EventReplySellToken",
	EventBuyToken:               "EventBuyToken",
	EventReplyBuyToken:          "EventReplyBuyToken",
	EventRevokeSellToken:        "EventRevokeSellToken",
	EventReplyRevokeSellToken:   "EventReplyRevokeSellToken",
	// config
	EventModifyConfig:      "EventModifyConfig",
	EventReplyModifyConfig: "EventReplyModifyConfig",

	//privacy
	EventPublic2privacy:             "EventPublic2privacy",
	EventReplyPublic2privacy:        "EventReplyPublic2privacy",
	EventPrivacy2privacy:            "EventPrivacy2privacy",
	EventReplyPrivacy2privacy:       "EventReplyPrivacy2privacy",
	EventPrivacy2public:             "EventPrivacy2public",
	EventReplyPrivacy2public:        "EventReplyPrivacy2public",
	EventShowPrivacyPK:              "EventShowPrivacyPK",
	EventReplyShowPrivacyPK:         "EventReplyShowPrivacyPK",
	EventShowPrivacyTransfer:        "EventShowPrivacyTransfer",
	EventReplyShowPrivacyTransfer:   "EventReplyShowPrivacyTransfer",
	EventShowPrivacyAccount:         "EventShowPrivacyAccount",
	EventReplyShowPrivacyAccount:    "EventReplyShowPrivacyAccount",
	EventGetPrivacyTransaction:      "EventGetPrivacyTransaction",
	EventReplyGetPrivacyTransaction: "EventReplyGetPrivacyTransaction",
	EventGetGlobalIndex:             "EventGetGlobalIndex",
	EventReplyGetGlobalIndex:        "EventReplyGetGlobalIndex",
	EventGetUTXOPubKey:              "EventGetUTXOPubKey",
	EventReplyGetUTXOPubKey:         "EventReplyGetUTXOPubKey",

}

//ty = 1 -> secp256k1
//ty = 2 -> ed25519
//ty = 3 -> sm2
const (
	Invalid           = 0
	SECP256K1         = 1
	ED25519           = 2
	SM2               = 3
	OnetimeED25519    = 4
	RingBaseonED25519 = 5
)

//const (
//	SignTypeInvalid        = 0
//	SignTypeSecp256k1      = 1
//	SignTypeED25519        = 2
//	SignTypeSM2            = 3
//	SignTypeOnetimeED25519 = 4
//	SignTypeRing           = 5
//)

const (
	SignNameSecp256k1      = "secp256k1"
	SignNameED25519        = "ed25519"
	SignNameSM2            = "sm2"
	SignNameOnetimeED25519 = "onetimeed25519"
	SignNameRing           = "RingSignatue"
)

var MapSignType2name = map[int]string{
	SECP256K1:         SignNameSecp256k1,
	ED25519:           SignNameED25519,
	SM2:               SignNameSM2,
	OnetimeED25519:    SignNameOnetimeED25519,
	RingBaseonED25519: SignNameRing,
}

var MapSignName2Type = map[string]int{
	SignNameSecp256k1:      SECP256K1,
	SignNameED25519:        ED25519,
	SignNameSM2:            SM2,
	SignNameOnetimeED25519: OnetimeED25519,
	SignNameRing:           RingBaseonED25519,
}

//log type
const (
	TyLogReserved = 0
	TyLogErr      = 1
	TyLogFee      = 2
	//coins
	TyLogTransfer        = 3
	TyLogGenesis         = 4
	TyLogDeposit         = 5
	TyLogExecTransfer    = 6
	TyLogExecWithdraw    = 7
	TyLogExecDeposit     = 8
	TyLogExecFrozen      = 9
	TyLogExecActive      = 10
	TyLogGenesisTransfer = 11
	TyLogGenesisDeposit  = 12

	//log for ticket
	TyLogNewTicket   = 111
	TyLogCloseTicket = 112
	TyLogMinerTicket = 113
	TyLogTicketBind  = 114

	//log for token create
	TyLogPreCreateToken    = 211
	TyLogFinishCreateToken = 212
	TyLogRevokeCreateToken = 213

	//log for trade
	TyLogTradeSell            = 310
	TyLogTradeBuy             = 311
	TyLogTradeRevoke          = 312
	TyLogTokenTransfer        = 313
	TyLogTokenGenesis         = 314
	TyLogTokenDeposit         = 315
	TyLogTokenExecTransfer    = 316
	TyLogTokenExecWithdraw    = 317
	TyLogTokenExecDeposit     = 318
	TyLogTokenExecFrozen      = 319
	TyLogTokenExecActive      = 320
	TyLogTokenGenesisTransfer = 321
	TyLogTokenGenesisDeposit  = 322

	// log for config
	TyLogModifyConfig = 410

	// log for privacy
	TyLogPrivacyFee = iota + 500
	TyLogPrivacyFeeUTXO
	TyLogPrivacyInput
	TyLogPrivacyOutput
)

//exec type
const (
	ExecErr  = 0
	ExecPack = 1
	ExecOk   = 2
)

const (
	PrivacyUTXONormal = iota
	PrivacyUTXOChange
	PrivacyUTXOFee
)

const (
	InvalidAction = iota
	CoinsActionTransfer
	CoinsActionGenesis
	CoinsActionWithdraw

	//action for token
	ActionTransfer
	ActionGenesis
	ActionWithdraw
	TokenActionPreCreate
	TokenActionFinishCreate
	TokenActionRevokeCreate
	//action type for privacy
	ActionPublic2Privacy = iota + 100
	ActionPrivacy2Privacy
	ActionPrivacy2Public
)

//ticket
const (
	TicketActionGenesis = 11
	TicketActionOpen    = 12
	TicketActionClose   = 13
	TicketActionList    = 14 //读的接口不直接经过transaction
	TicketActionInfos   = 15 //读的接口不直接经过transaction
	TicketActionMiner   = 16
	TicketActionBind    = 17
)

//hashlock const
const (
	HashlockActionLock   = 1
	HashlockActionSend   = 2
	HashlockActionUnlock = 3
)

//retrieve
const (
	RetrievePre    = 1
	RetrievePerf   = 2
	RetrieveBackup = 3
	RetrieveCancel = 4
)

const (
	TokenStatusPreCreated = iota
	TokenStatusCreated
	TokenStatusCreateRevoked
)

const (
	TradeSell = iota
	TradeBuy
	TradeRevokeSell
)

//0->not start, 1->on sale, 2->sold out, 3->revoke, 4->expired
const (
	NotStart = iota
	OnSale
	SoldOut
	Revoked
	Expired
)

var SellOrderStatus = map[int32]string{
	NotStart: "NotStart",
	OnSale:   "OnSale",
	SoldOut:  "SoldOut",
	Revoked:  "Revoked",
	Expired:  "Expired",
}

const (
	ManageActionModifyConfig = iota
)

const (
	ConfigItemArrayConfig = iota
	ConfigItemIntConfig
	ConfigItemStringConfig
)

var MapSellOrderStatusStr2Int = map[string]int32{
	"onsale":  OnSale,
	"soldout": SoldOut,
	"revoked": Revoked,
}

//hard fork block height
const (
	ForkV1             = 75260
	ForkV2_add_token   = 100899
	ForkV3             = 110000
	ForkV4_add_manage  = 120000
	ForkV5_add_privacy = 200000
)
