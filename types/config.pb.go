// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Config struct {
	Title           string      `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Loglevel        string      `protobuf:"bytes,2,opt,name=loglevel" json:"loglevel,omitempty"`
	LogConsoleLevel string      `protobuf:"bytes,3,opt,name=logConsoleLevel" json:"logConsoleLevel,omitempty"`
	LogFile         string      `protobuf:"bytes,4,opt,name=logFile" json:"logFile,omitempty"`
	Store           *Store      `protobuf:"bytes,6,opt,name=store" json:"store,omitempty"`
	LocalStore      *LocalStore `protobuf:"bytes,7,opt,name=localStore" json:"localStore,omitempty"`
	Consensus       *Consensus  `protobuf:"bytes,8,opt,name=consensus" json:"consensus,omitempty"`
	MemPool         *MemPool    `protobuf:"bytes,9,opt,name=memPool" json:"memPool,omitempty"`
	BlockChain      *BlockChain `protobuf:"bytes,10,opt,name=blockChain" json:"blockChain,omitempty"`
	Wallet          *Wallet     `protobuf:"bytes,11,opt,name=wallet" json:"wallet,omitempty"`
	P2P             *P2P        `protobuf:"bytes,12,opt,name=p2p" json:"p2p,omitempty"`
	Rpc             *Rpc        `protobuf:"bytes,13,opt,name=rpc" json:"rpc,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Config) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Config) GetLoglevel() string {
	if m != nil {
		return m.Loglevel
	}
	return ""
}

func (m *Config) GetLogConsoleLevel() string {
	if m != nil {
		return m.LogConsoleLevel
	}
	return ""
}

func (m *Config) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *Config) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *Config) GetLocalStore() *LocalStore {
	if m != nil {
		return m.LocalStore
	}
	return nil
}

func (m *Config) GetConsensus() *Consensus {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *Config) GetMemPool() *MemPool {
	if m != nil {
		return m.MemPool
	}
	return nil
}

func (m *Config) GetBlockChain() *BlockChain {
	if m != nil {
		return m.BlockChain
	}
	return nil
}

func (m *Config) GetWallet() *Wallet {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func (m *Config) GetP2P() *P2P {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *Config) GetRpc() *Rpc {
	if m != nil {
		return m.Rpc
	}
	return nil
}

type MemPool struct {
	PoolCacheSize int64 `protobuf:"varint,1,opt,name=poolCacheSize" json:"poolCacheSize,omitempty"`
	MinTxFee      int64 `protobuf:"varint,2,opt,name=minTxFee" json:"minTxFee,omitempty"`
}

func (m *MemPool) Reset()                    { *m = MemPool{} }
func (m *MemPool) String() string            { return proto.CompactTextString(m) }
func (*MemPool) ProtoMessage()               {}
func (*MemPool) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *MemPool) GetPoolCacheSize() int64 {
	if m != nil {
		return m.PoolCacheSize
	}
	return 0
}

func (m *MemPool) GetMinTxFee() int64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

type Consensus struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Genesis          string `protobuf:"bytes,2,opt,name=genesis" json:"genesis,omitempty"`
	Minerstart       bool   `protobuf:"varint,3,opt,name=minerstart" json:"minerstart,omitempty"`
	NodeId           int64  `protobuf:"varint,4,opt,name=nodeId" json:"nodeId,omitempty"`
	RaftApiPort      int64  `protobuf:"varint,5,opt,name=raftApiPort" json:"raftApiPort,omitempty"`
	IsNewJoinNode    bool   `protobuf:"varint,6,opt,name=isNewJoinNode" json:"isNewJoinNode,omitempty"`
	PeersURL         string `protobuf:"bytes,7,opt,name=peersURL" json:"peersURL,omitempty"`
	ReadOnlyPeersURL string `protobuf:"bytes,8,opt,name=readOnlyPeersURL" json:"readOnlyPeersURL,omitempty"`
	AddPeersURL      string `protobuf:"bytes,9,opt,name=addPeersURL" json:"addPeersURL,omitempty"`
	GenesisBlockTime int64  `protobuf:"varint,10,opt,name=genesisBlockTime" json:"genesisBlockTime,omitempty"`
	HotkeyAddr       string `protobuf:"bytes,11,opt,name=hotkeyAddr" json:"hotkeyAddr,omitempty"`
}

func (m *Consensus) Reset()                    { *m = Consensus{} }
func (m *Consensus) String() string            { return proto.CompactTextString(m) }
func (*Consensus) ProtoMessage()               {}
func (*Consensus) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Consensus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consensus) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *Consensus) GetMinerstart() bool {
	if m != nil {
		return m.Minerstart
	}
	return false
}

func (m *Consensus) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *Consensus) GetRaftApiPort() int64 {
	if m != nil {
		return m.RaftApiPort
	}
	return 0
}

func (m *Consensus) GetIsNewJoinNode() bool {
	if m != nil {
		return m.IsNewJoinNode
	}
	return false
}

func (m *Consensus) GetPeersURL() string {
	if m != nil {
		return m.PeersURL
	}
	return ""
}

func (m *Consensus) GetReadOnlyPeersURL() string {
	if m != nil {
		return m.ReadOnlyPeersURL
	}
	return ""
}

func (m *Consensus) GetAddPeersURL() string {
	if m != nil {
		return m.AddPeersURL
	}
	return ""
}

func (m *Consensus) GetGenesisBlockTime() int64 {
	if m != nil {
		return m.GenesisBlockTime
	}
	return 0
}

func (m *Consensus) GetHotkeyAddr() string {
	if m != nil {
		return m.HotkeyAddr
	}
	return ""
}

type Wallet struct {
	MinFee   int64  `protobuf:"varint,1,opt,name=minFee" json:"minFee,omitempty"`
	Driver   string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath   string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	SignType string `protobuf:"bytes,4,opt,name=signType" json:"signType,omitempty"`
}

func (m *Wallet) Reset()                    { *m = Wallet{} }
func (m *Wallet) String() string            { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()               {}
func (*Wallet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Wallet) GetMinFee() int64 {
	if m != nil {
		return m.MinFee
	}
	return 0
}

func (m *Wallet) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Wallet) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *Wallet) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

type Store struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Driver string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *Store) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Store) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Store) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

type LocalStore struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	DbPath string `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *LocalStore) Reset()                    { *m = LocalStore{} }
func (m *LocalStore) String() string            { return proto.CompactTextString(m) }
func (*LocalStore) ProtoMessage()               {}
func (*LocalStore) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *LocalStore) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *LocalStore) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

type BlockChain struct {
	DefCacheSize        int64  `protobuf:"varint,1,opt,name=defCacheSize" json:"defCacheSize,omitempty"`
	MaxFetchBlockNum    int64  `protobuf:"varint,2,opt,name=maxFetchBlockNum" json:"maxFetchBlockNum,omitempty"`
	TimeoutSeconds      int64  `protobuf:"varint,3,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	BatchBlockNum       int64  `protobuf:"varint,4,opt,name=batchBlockNum" json:"batchBlockNum,omitempty"`
	Driver              string `protobuf:"bytes,5,opt,name=driver" json:"driver,omitempty"`
	DbPath              string `protobuf:"bytes,6,opt,name=dbPath" json:"dbPath,omitempty"`
	IsStrongConsistency bool   `protobuf:"varint,7,opt,name=isStrongConsistency" json:"isStrongConsistency,omitempty"`
	SingleMode          bool   `protobuf:"varint,8,opt,name=singleMode" json:"singleMode,omitempty"`
}

func (m *BlockChain) Reset()                    { *m = BlockChain{} }
func (m *BlockChain) String() string            { return proto.CompactTextString(m) }
func (*BlockChain) ProtoMessage()               {}
func (*BlockChain) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *BlockChain) GetDefCacheSize() int64 {
	if m != nil {
		return m.DefCacheSize
	}
	return 0
}

func (m *BlockChain) GetMaxFetchBlockNum() int64 {
	if m != nil {
		return m.MaxFetchBlockNum
	}
	return 0
}

func (m *BlockChain) GetTimeoutSeconds() int64 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *BlockChain) GetBatchBlockNum() int64 {
	if m != nil {
		return m.BatchBlockNum
	}
	return 0
}

func (m *BlockChain) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *BlockChain) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *BlockChain) GetIsStrongConsistency() bool {
	if m != nil {
		return m.IsStrongConsistency
	}
	return false
}

func (m *BlockChain) GetSingleMode() bool {
	if m != nil {
		return m.SingleMode
	}
	return false
}

type P2P struct {
	SeedPort     int32    `protobuf:"varint,1,opt,name=seedPort" json:"seedPort,omitempty"`
	DbPath       string   `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
	GrpcLogFile  string   `protobuf:"bytes,3,opt,name=grpcLogFile" json:"grpcLogFile,omitempty"`
	IsSeed       bool     `protobuf:"varint,4,opt,name=isSeed" json:"isSeed,omitempty"`
	ServerStart  bool     `protobuf:"varint,5,opt,name=serverStart" json:"serverStart,omitempty"`
	Seeds        []string `protobuf:"bytes,6,rep,name=seeds" json:"seeds,omitempty"`
	Enable       bool     `protobuf:"varint,7,opt,name=enable" json:"enable,omitempty"`
	MsgCacheSize int32    `protobuf:"varint,8,opt,name=msgCacheSize" json:"msgCacheSize,omitempty"`
	Version      int32    `protobuf:"varint,9,opt,name=version" json:"version,omitempty"`
	VerMix       int32    `protobuf:"varint,10,opt,name=verMix" json:"verMix,omitempty"`
	VerMax       int32    `protobuf:"varint,11,opt,name=verMax" json:"verMax,omitempty"`
}

func (m *P2P) Reset()                    { *m = P2P{} }
func (m *P2P) String() string            { return proto.CompactTextString(m) }
func (*P2P) ProtoMessage()               {}
func (*P2P) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *P2P) GetSeedPort() int32 {
	if m != nil {
		return m.SeedPort
	}
	return 0
}

func (m *P2P) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *P2P) GetGrpcLogFile() string {
	if m != nil {
		return m.GrpcLogFile
	}
	return ""
}

func (m *P2P) GetIsSeed() bool {
	if m != nil {
		return m.IsSeed
	}
	return false
}

func (m *P2P) GetServerStart() bool {
	if m != nil {
		return m.ServerStart
	}
	return false
}

func (m *P2P) GetSeeds() []string {
	if m != nil {
		return m.Seeds
	}
	return nil
}

func (m *P2P) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *P2P) GetMsgCacheSize() int32 {
	if m != nil {
		return m.MsgCacheSize
	}
	return 0
}

func (m *P2P) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *P2P) GetVerMix() int32 {
	if m != nil {
		return m.VerMix
	}
	return 0
}

func (m *P2P) GetVerMax() int32 {
	if m != nil {
		return m.VerMax
	}
	return 0
}

type Rpc struct {
	JrpcBindAddr string   `protobuf:"bytes,1,opt,name=jrpcBindAddr" json:"jrpcBindAddr,omitempty"`
	GrpcBindAddr string   `protobuf:"bytes,2,opt,name=grpcBindAddr" json:"grpcBindAddr,omitempty"`
	Whitlist     []string `protobuf:"bytes,3,rep,name=whitlist" json:"whitlist,omitempty"`
}

func (m *Rpc) Reset()                    { *m = Rpc{} }
func (m *Rpc) String() string            { return proto.CompactTextString(m) }
func (*Rpc) ProtoMessage()               {}
func (*Rpc) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *Rpc) GetJrpcBindAddr() string {
	if m != nil {
		return m.JrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetGrpcBindAddr() string {
	if m != nil {
		return m.GrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetWhitlist() []string {
	if m != nil {
		return m.Whitlist
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "types.Config")
	proto.RegisterType((*MemPool)(nil), "types.MemPool")
	proto.RegisterType((*Consensus)(nil), "types.Consensus")
	proto.RegisterType((*Wallet)(nil), "types.Wallet")
	proto.RegisterType((*Store)(nil), "types.Store")
	proto.RegisterType((*LocalStore)(nil), "types.LocalStore")
	proto.RegisterType((*BlockChain)(nil), "types.BlockChain")
	proto.RegisterType((*P2P)(nil), "types.P2P")
	proto.RegisterType((*Rpc)(nil), "types.Rpc")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xe4, 0x44,
	0x10, 0x56, 0xc6, 0xeb, 0x89, 0xa7, 0x93, 0x2c, 0x4b, 0x83, 0x90, 0x85, 0x10, 0x8a, 0x2c, 0x40,
	0x11, 0x87, 0x08, 0xc2, 0x95, 0xcb, 0xee, 0x48, 0x91, 0x60, 0x93, 0x30, 0xea, 0x09, 0xe2, 0xec,
	0xb1, 0x2b, 0x9e, 0x66, 0xdb, 0xdd, 0x56, 0x77, 0x67, 0x32, 0xc3, 0x33, 0x71, 0xe1, 0x21, 0x78,
	0x1e, 0x5e, 0x01, 0x55, 0xb9, 0xed, 0xb1, 0xf3, 0x73, 0xd8, 0xdb, 0x7c, 0x5f, 0x7d, 0xd5, 0xd3,
	0x55, 0xf5, 0x55, 0x9b, 0x1d, 0x17, 0x46, 0xdf, 0xc9, 0xea, 0xbc, 0xb1, 0xc6, 0x1b, 0x1e, 0xfb,
	0x5d, 0x03, 0x2e, 0xfb, 0x37, 0x62, 0xd3, 0x39, 0xf1, 0xfc, 0x73, 0x16, 0x7b, 0xe9, 0x15, 0xa4,
	0x07, 0xa7, 0x07, 0x67, 0x33, 0xd1, 0x02, 0xfe, 0x25, 0x4b, 0x94, 0xa9, 0x14, 0x6c, 0x40, 0xa5,
	0x13, 0x0a, 0xf4, 0x98, 0x9f, 0xb1, 0x4f, 0x94, 0xa9, 0xe6, 0x46, 0x3b, 0xa3, 0xe0, 0x8a, 0x24,
	0x11, 0x49, 0x1e, 0xd3, 0x3c, 0x65, 0x87, 0xca, 0x54, 0x97, 0x52, 0x41, 0xfa, 0x8a, 0x14, 0x1d,
	0xe4, 0x19, 0x8b, 0x9d, 0x37, 0x16, 0xd2, 0xe9, 0xe9, 0xc1, 0xd9, 0xd1, 0xc5, 0xf1, 0x39, 0xdd,
	0xeb, 0x7c, 0x89, 0x9c, 0x68, 0x43, 0xfc, 0x47, 0xc6, 0x94, 0x29, 0x72, 0x45, 0x64, 0x7a, 0x48,
	0xc2, 0x4f, 0x83, 0xf0, 0xaa, 0x0f, 0x88, 0x81, 0x88, 0x9f, 0xb3, 0x59, 0x61, 0xb4, 0x03, 0xed,
	0xee, 0x5d, 0x9a, 0x50, 0xc6, 0x9b, 0x90, 0x31, 0xef, 0x78, 0xb1, 0x97, 0xf0, 0x33, 0x76, 0x58,
	0x43, 0xbd, 0x30, 0x46, 0xa5, 0x33, 0x52, 0xbf, 0x0e, 0xea, 0xeb, 0x96, 0x15, 0x5d, 0x18, 0x2f,
	0xb3, 0x52, 0xa6, 0xf8, 0x30, 0x5f, 0xe7, 0x52, 0xa7, 0x6c, 0x74, 0x99, 0x77, 0x7d, 0x40, 0x0c,
	0x44, 0xfc, 0x5b, 0x36, 0x7d, 0xc8, 0x95, 0x02, 0x9f, 0x1e, 0x91, 0xfc, 0x24, 0xc8, 0xff, 0x20,
	0x52, 0x84, 0x20, 0xff, 0x8a, 0x45, 0xcd, 0x45, 0x93, 0x1e, 0x93, 0x86, 0x05, 0xcd, 0xe2, 0x62,
	0x21, 0x90, 0xc6, 0xa8, 0x6d, 0x8a, 0xf4, 0x64, 0x14, 0x15, 0x4d, 0x21, 0x90, 0xce, 0xde, 0xb3,
	0xc3, 0x70, 0x53, 0xfe, 0x0d, 0x3b, 0x69, 0x8c, 0x51, 0xf3, 0xbc, 0x58, 0xc3, 0x52, 0xfe, 0xd5,
	0xce, 0x33, 0x12, 0x63, 0x12, 0xe7, 0x5a, 0x4b, 0x7d, 0xbb, 0xbd, 0x04, 0xa0, 0xb9, 0x46, 0xa2,
	0xc7, 0xd9, 0x7f, 0x13, 0x36, 0xeb, 0xbb, 0xc4, 0x39, 0x7b, 0xa5, 0xf3, 0xba, 0xb3, 0x05, 0xfd,
	0xc6, 0x79, 0x56, 0xa0, 0xc1, 0x49, 0x17, 0x4c, 0xd1, 0x41, 0xfe, 0x35, 0x63, 0xb5, 0xd4, 0x60,
	0x9d, 0xcf, 0xad, 0x27, 0x3b, 0x24, 0x62, 0xc0, 0xf0, 0x2f, 0xd8, 0x54, 0x9b, 0x12, 0x7e, 0x29,
	0xc9, 0x08, 0x91, 0x08, 0x88, 0x9f, 0xb2, 0x23, 0x9b, 0xdf, 0xf9, 0xb7, 0x8d, 0x5c, 0x18, 0xeb,
	0xd3, 0x98, 0x82, 0x43, 0x0a, 0xeb, 0x92, 0xee, 0x06, 0x1e, 0x7e, 0x35, 0x52, 0xdf, 0x98, 0xb2,
	0x75, 0x4c, 0x22, 0xc6, 0x24, 0xd6, 0xd5, 0x00, 0x58, 0xf7, 0xbb, 0xb8, 0x22, 0xa7, 0xcc, 0x44,
	0x8f, 0xf9, 0xf7, 0xec, 0x8d, 0x85, 0xbc, 0xfc, 0x4d, 0xab, 0xdd, 0xa2, 0xd3, 0x24, 0xa4, 0x79,
	0xc2, 0xe3, 0x7d, 0xf2, 0xb2, 0xec, 0x65, 0x33, 0x92, 0x0d, 0x29, 0x3c, 0x2d, 0x14, 0x4d, 0x63,
	0xbf, 0x95, 0x35, 0x90, 0x1d, 0x22, 0xf1, 0x84, 0xc7, 0xae, 0xac, 0x8d, 0xff, 0x00, 0xbb, 0xb7,
	0x65, 0x69, 0xc9, 0x05, 0x33, 0x31, 0x60, 0x32, 0xc5, 0xa6, 0xad, 0x19, 0xb0, 0x3f, 0xb5, 0xd4,
	0x38, 0x95, 0x76, 0x6c, 0x01, 0x21, 0x5f, 0x5a, 0xb9, 0x01, 0x1b, 0x1a, 0x1e, 0x10, 0xf1, 0xab,
	0x45, 0xee, 0xd7, 0x61, 0xf5, 0x02, 0xc2, 0x3e, 0x38, 0x59, 0xe9, 0xdb, 0x5d, 0xd3, 0xad, 0x5c,
	0x8f, 0xb3, 0xf7, 0x2c, 0x6e, 0xb7, 0xe4, 0xb9, 0xd1, 0x7e, 0xe4, 0x1f, 0x65, 0x3f, 0x33, 0xb6,
	0xdf, 0xc1, 0x41, 0xf6, 0xc1, 0x0b, 0xd9, 0x93, 0x51, 0xf6, 0xdf, 0x13, 0xc6, 0xf6, 0x5b, 0xc3,
	0x33, 0x76, 0x5c, 0xc2, 0xdd, 0x63, 0xeb, 0x8e, 0x38, 0xec, 0x7b, 0x9d, 0x6f, 0x2f, 0xc1, 0x17,
	0x6b, 0xca, 0xbc, 0xb9, 0xaf, 0x83, 0x83, 0x9f, 0xf0, 0xfc, 0x3b, 0xf6, 0xda, 0xcb, 0x1a, 0xcc,
	0xbd, 0x5f, 0x42, 0x61, 0x74, 0xe9, 0xe8, 0xf2, 0x91, 0x78, 0xc4, 0xa2, 0xb7, 0x56, 0xf9, 0xf0,
	0xc0, 0xd6, 0x9c, 0x63, 0x72, 0x50, 0x5c, 0xfc, 0x42, 0x71, 0xd3, 0xd1, 0x0c, 0x7e, 0x60, 0x9f,
	0x49, 0xb7, 0xf4, 0xd6, 0x68, 0x7a, 0x0d, 0xa5, 0xf3, 0xa0, 0x8b, 0x1d, 0xd9, 0x32, 0x11, 0xcf,
	0x85, 0xd0, 0x27, 0x4e, 0xea, 0x4a, 0xc1, 0x35, 0x1a, 0x3c, 0x69, 0xb7, 0x67, 0xcf, 0x64, 0xff,
	0x4c, 0x58, 0xb4, 0xb8, 0x58, 0xd0, 0x74, 0x01, 0x4a, 0x5a, 0x15, 0xec, 0x51, 0x2c, 0x7a, 0xfc,
	0x52, 0xab, 0xd1, 0xd1, 0x95, 0x6d, 0x8a, 0xab, 0xf0, 0x0e, 0xb7, 0x53, 0x1c, 0x52, 0x98, 0x29,
	0xdd, 0x12, 0xa0, 0xdd, 0xcd, 0x44, 0x04, 0x84, 0x99, 0x0e, 0xec, 0x06, 0xec, 0x92, 0x96, 0x3a,
	0xa6, 0xe0, 0x90, 0xc2, 0x6f, 0x07, 0xfe, 0xbf, 0x4b, 0xa7, 0xa7, 0x11, 0x7e, 0x3b, 0x08, 0xe0,
	0x79, 0xa0, 0xf3, 0x95, 0x82, 0x50, 0x72, 0x40, 0x38, 0xe5, 0xda, 0x55, 0xfb, 0x29, 0x27, 0x54,
	0xc1, 0x88, 0xc3, 0x17, 0x66, 0x03, 0xd6, 0x49, 0xa3, 0x69, 0xf7, 0x62, 0xd1, 0x41, 0x3c, 0x75,
	0x03, 0xf6, 0x5a, 0x6e, 0x69, 0xdb, 0x62, 0x11, 0x50, 0xc7, 0xe7, 0x5b, 0xda, 0xaf, 0xc0, 0xe7,
	0xdb, 0x4c, 0xb2, 0x48, 0x34, 0x05, 0xfe, 0xe9, 0x9f, 0xb6, 0x29, 0xde, 0x49, 0x5d, 0xd2, 0x12,
	0xb6, 0xfe, 0x1c, 0x71, 0xa8, 0xa9, 0x86, 0x9a, 0xb6, 0x81, 0x23, 0x0e, 0x5b, 0xff, 0xb0, 0x96,
	0x5e, 0x49, 0x87, 0xcf, 0x1b, 0x56, 0xdb, 0xe3, 0xd5, 0x94, 0xbe, 0xad, 0x3f, 0xfd, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0xb1, 0xef, 0xb6, 0x19, 0x6b, 0x07, 0x00, 0x00,
}
