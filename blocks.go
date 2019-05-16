package datafile

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

var _ = fmt.Print

type BlockType uint8

const (
	rootOfRootsT BlockType = iota
	linIntrBptNodeT
	linLeafBptNodeT
	kvIntrBptNodeT
	kvLeafBptNodeT
	blobT
	stringT
	numBlockTypes
)

var BlockTypes = []string{
	"rootOfRootsT",
	"linIntrBptNodeT",
	"linLeafBptNodeT",
	"kvIntrBptNodeT",
	"kvLeafBptNodeT",
	"blobT",
	"stringT",
}

var BlockTypeMap map[string]BlockType

func init() {
	BlockTypeMap = make(map[string]BlockType)
	for typ, name := range BlockTypes {
		BlockTypeMap[name] = BlockType(uint8(typ))
	}
}

type BlockId uint64

type BlockTypeId uint64

const BlockIdMask uint64 = 0x0000ffffffffffff
const BlockIdShift uint64 = 48
const BlockTypeMask uint64 = 0xff00000000000000 // 0xff<<56
const BlockTypeShift uint64 = 56
const BlockSubmapMask uint64 = 0x00ff000000000000 // 0xff<<48
const BlockSubmapShift uint64 = 48

func BlockTypeFromBtid(btid BlockTypeId) BlockType {
	return BlockType(uint64(btid) >> BlockTypeShift)
	//return BlockType((uint64(btid) & BlockTypeMask) >> BlockTypeShift)
	//return BlockType((uint64(btid) & (0xff << 56)) >> 56)
}

func BlockIdFromBtid(btid BlockTypeId) BlockId {
	return BlockId(uint64(btid) & BlockIdMask)
	//return BlockId(uint64(btid) & 0x00ffffffffffffff)
}

func BlockSubmapFromBtid(btid BlockTypeId) (submapType byte, mapmask byte) {
	submapinfo := BlockSubmapMask & uint64(btid)
	submapType = byte(submapinfo & 0xffff0000)
	mapmask = byte(submapinfo & 0x0000ffff)
	return submapType, mapmask
}

var DatafileSig *uuid.UUID
var DatafileHeaderSig *uuid.UUID
var RootOfRootsSig *uuid.UUID
var LinBptIntrNodeSig *uuid.UUID
var LinBptLeafNodeSig *uuid.UUID
var KvBptIntrNodeSig *uuid.UUID
var KvBptLeafNodeSig *uuid.UUID
var BlobBlockSig *uuid.UUID
var StringBlockSig *uuid.UUID

func init() {
	//DatafileSig = &uuid.NewV5(uuid.NamespaceURL, "github.com/lleo/datafile")
	//DatafileHeaderSig = &uuid.NewV5(DatafileSig, "DatafileHeader")
	//RootOfRootsSig = &uuid.NewV5(DatafileSig, "RootOfRoots")
	//LinBptIntrNodeSig = &uuid.NewV5(DatafileSig, "LinBptIntrNode")
	//LinBptLeafNodeSig = &uuid.NewV5(DatafileSig, "LinBptLeafNode")
	//KvBptIntrNodeSig = &uuid.NewV5(DatafileSig, "KvBptIntrNode")
	//KvBptLeafNodeSig = &uuid.NewV5(DatafileSig, "KvBptLeafNode")
	//BlobBlockSig = &uuid.NewV5(DatafileSig, "BlobBlock")
	//StringBlockSig = &uuid.NewV5(DatafileSig, "StringBlock")

	var datafileSig = uuid.NewV5(uuid.NamespaceURL, "github.com/lleo/datafile")
	DatafileSig = &datafileSig

	var datafileHeaderSig = uuid.NewV5(uuid.NamespaceURL, "github.com/lleo/datafile")
	DatafileHeaderSig = &datafileHeaderSig

	var rootOfRootsSig = uuid.NewV5(*DatafileSig, "RootOfRoots")
	RootOfRootsSig = &rootOfRootsSig

	var linBptIntrNodeSig = uuid.NewV5(*DatafileSig, "LinBptIntrNode")
	LinBptIntrNodeSig = &linBptIntrNodeSig

	var linBptLeafNodeSig = uuid.NewV5(*DatafileSig, "LinBptLeafNode")
	LinBptLeafNodeSig = &linBptLeafNodeSig

	var kvBptIntrNodeSig = uuid.NewV5(*DatafileSig, "KvBptIntrNode")
	LinBptIntrNodeSig = &kvBptIntrNodeSig

	var kvBptLeafNodeSig = uuid.NewV5(*DatafileSig, "KvBptLeafNode")
	LinBptLeafNodeSig = &kvBptLeafNodeSig

	var blobBlockSig = uuid.NewV5(*DatafileSig, "BlobBlock")
	BlobBlockSig = &blobBlockSig

	var stringBlockSig = uuid.NewV5(*DatafileSig, "StringBlock")
	StringBlockSig = &stringBlockSig
}

var BlockType2Sig = []*uuid.UUID{
	RootOfRootsSig,
	LinBptIntrNodeSig,
	LinBptLeafNodeSig,
	KvBptIntrNodeSig,
	KvBptLeafNodeSig,
	BlobBlockSig,
	StringBlockSig,
}

type BlockI interface {
	TypeAndId() BlockTypeId
	Type() BlockType
	Id() BlockId
	Marshal() ([]byte, error)
	Write(p []byte) (int, error)
	//isBlock/isBptNode
}

type BptNodeI interface {
	BlockI
	//Generic B+Tree interface (not get/put/del)
}

type LinearBptNodeI interface {
	BptNodeI
	//...
}

type KeyvalBptNodeI interface {
	BptNodeI
	//...
}

type DataBlockI interface {
	BlockI
	Order2() byte
	GetData() []byte
}

type blockBaseS struct {
	sig *uuid.UUID
	typ BlockType
	id  BlockId
	//rawBlk []byte
}

type linIntrNodeS struct {
	blockBaseS
	used  byte
	btids []BlockTypeId
	nodes []BptNodeI
}

type linLeafNodeS struct {
	blockBaseS
	used  byte
	btids []BlockTypeId
	dBlks []DataBlockI
}

type blobBlkS struct {
	blockBaseS
	//next BlockTypeId //continuation default=0
	order2 int
	used   uint32
	bytes  []byte
}

func btid2Type(btid BlockTypeId) BlockType {
	return BlockType((uint64(btid) & BlockTypeMask) >> BlockTypeShift)
}

func newRootOfRoots(pageSize uint32) *linIntrNodeS {
	var x linIntrNodeS
	x.blockBaseS = blockBaseS{sig: RootOfRootsSig, typ: rootOfRootsT}
	//x.blockBaseS.sig = RootOfRootsSig
	//x.blockBaseS.typ = rootOfRootsT

	x = linIntrNodeS{
		blockBaseS: blockBaseS{
			sig: RootOfRootsSig,
			typ: rootOfRootsT,
		},
		used: 0,
	}
	return &x
}

func (b *linIntrNodeS) BlockTypeId() BlockTypeId {
	return BlockTypeId((uint64(b.typ) << 56) & uint64(b.id))
}
