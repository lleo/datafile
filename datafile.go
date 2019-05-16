package datafile

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

var _ = fmt.Println
var _ = log.Println

const PageSize = 512

type headerInfo struct {
	sig          [16]byte
	pageSize     uint32
	rootOfRootId BlockTypeId
}

//rootOfRoots indexes
const (
	linearStringsRoot = iota
	linearBlobsRoot
	namedKeyValRoots
	numRoots
	//additionalIndexes //???
)

var Roots = []string{
	"linearStringsRoot",
	"linearBlobRoot",
	"namedKeyValRoots",
}

type DatafileS struct {
	prev         *DatafileS
	dataFile     *os.File
	fileName     string
	openTxn      *txn
	tableOrder   uint32
	leafCapacity uint32
	pageSize     uint32
	root         *linIntrNodeS //rootOfRoots
	//index        map[string]*Index
}

func Create(filename string, pagesize uint32) (*DatafileS, error) {
	var df = initialize(filename, pagesize)
	var err error
	df.dataFile, err = os.Create(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "failed os.Create(%q)", filename)
	}
	df.writeHeaders()
	return df, nil
}

func initialize(filename string, pagesize uint32) *DatafileS {
	var df = &DatafileS{
		prev:         nil,
		dataFile:     nil,
		fileName:     filename,
		openTxn:      nil,
		tableOrder:   calcLinBptOrder(pagesize),
		leafCapacity: calcLeafCapacity(pagesize),
		//index:        make(map[string]*Index),
		pageSize: pagesize,
		root: &linIntrNodeS{
			blockBaseS{typ: rootOfRootsT, id: 0},
			0,
			[]BlockTypeId{},
			[]BptNodeI{},
		},
	}
	return df
}

func calcLinBptOrder(pagesize uint32) uint32 {
	//base libBptNodeS == UUID+NodeType+tableLen
	return (pagesize - 16 - 1 - 1) / 8
	// pagesize LinBptOrder KvBptOrder
	//       32           1 0
	//       64           5 2
	//      128          13 6
	//      256          29 14
	//      512          61 30
	//     1024         125 62
	//     2048         253 126
	//     4096         509 254
}

func calcKvBptOrder(pagesize uint32) uint32 {
	return
}

func calcLeafCapacity(pagesize uint32) uint32 {
	return pagesize
}

func Open(filename string) (*DatafileS, error) {
	return nil, NotImplemented("Open")
}

func (df *DatafileS) StartTxn() error {
	if df.prev != nil {
		return PrevExists("(df *DatafileS) StartTxn")
	}
	return NotImplemented("StartTxn")
}

func (df *DatafileS) writeHeaders() error {
	return NotImplemented("writeHeaders")
}

func (df *DatafileS) AllocBlock(size uint64, bt BlockType) (
	BlockTypeId,
	error,
) {
	return 0, NotImplemented("AllocBlob")
}

func (df *DatafileS) FreeBlock(btid BlockTypeId) error {
	return NotImplemented("FreeBlock")
}

//func (df *DatafileS) LookupByIndex(
//	idxName string,
//	v []interface{},
//) ([]interface{}, error) {
//	t := df.StartTxn()
//	found, err := t.LookupByIndex(indexName, v)
//	if err != nil {
//		t.Rollback()
//		return err
//	}
//	err = t.Commit()
//	return found, err
//}
//
//func (df *DatafileS) CreateIndex(
//	indexName string,
//	fieldNames []string,
//	typ IndexType,
//) (bool, error) {
//	t := df.StartTxn()
//	created, err := t.CreateIndex(indexName)
//	if err != nil {
//		t.Rollback()
//		return created, err
//	}
//	err = t.Commit()
//	return created, err
//}
//
//func (df *DatafileS) DropIndex(idxName string) (bool, error) {
//	t := df.StartTxn()
//	dropped, err := t.DropIndex(indexName)
//	if err != nil {
//		t.Rollback()
//		return dropped, err
//	}
//	err = t.Commit()
//	return dropped, err
//}
