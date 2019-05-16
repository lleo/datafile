package datafile

type LinearBptree interface {
	BlockI
	Order() int
	NumberOfEntries() int
	Get(int) ([]byte, bool)
	Put(int, []byte) (LinearBptree, error)
	Del(int) (LinearBptree, bool)
}

type linBptNodeI interface {
	BlockI
	equals(linBptNodeI) bool
	isToBig() bool
	isLeaf() bool
	findLeftMost() linBptNodeI
	size() int
	order() int
	halfFullSize() int
	depth() int
}

//type linBptIntrNodeI interface {
//	linBptNodeI
//	get(int) (linBptNodeI, bool)
//	put(int, linBptNodeI) (linBptIntrNodeI, error)
//	del(int) (linBptIntrNodeI, error)
//}
//
//type linBptLeafNodeI interface {
//	linBptNodeI
//	get(int) ([]byte, bool)
//	put(int, []byte) (linBptLeafNodeI, error)
//	del(int) (linBptLeafNodeI, error)
//}

type linearBptreeS struct {
	btid     BlockTypeId
	root     linBptNodeI
	pageSize uint32
	order    int
	maxEnt   int
	numEnts  int
	depth    int
}

type lbptIntrNodeS struct {
	btid   BlockTypeId
	vals   []linBptNodeI
	depth_ int
	order_ int
}

type lbptLeafNodeS struct {
	btid   BlockTypeId
	vals   [][]byte
	depth_ int
	order_ int
}

func newLinearBptree(btid BlockTypeId, pagesize uint32) LinearBptree {
	var t = &linearBptreeS{
		btid:     btid,
		root:     &lbptLeafNodeS{},
		pageSize: pagesize,
		order:    calcLbptOrder(pagesize),
		maxEnt:   0,
		numEnts:  0,
		depth:    0,
	}

	return t
}

func calcLbptOrder(pagesize uint32) int {
	//FIXME: NOT IMPLEMENTED
	return int(pagesize)
}

// linearBptreeS as BlockI
//
func (t *linearBptreeS) TypeAndId() BlockTypeId {
	return t.btid
}

func (t *linearBptreeS) Type() BlockType {
	return btid2Type(t.btid)
}

func (t *linearBptreeS) Id() BlockId {
	return BlockId(uint64(t.btid) & BlockIdMask)
}

func (t *linearBptreeS) Marshal() ([]byte, error) {
	return nil, NotImplemented("*linearBptreeS->Marshal")
}

func (t *linearBptreeS) Write(p []byte) (int, error) {
	return -1, NotImplemented("*linearBptreeS->Write")
}

// linearBptreeS as LinearBptree
//
func (t *linearBptreeS) Order() int {
	return t.order
}

func (t *linearBptreeS) NumberOfEntries() int {
	return t.numEnts
}

func (t *linearBptreeS) Get(i int) ([]byte, bool) {
	return nil, false
}

func (t *linearBptreeS) Put(i int, d []byte) (LinearBptree, error) {
	return nil, NotImplemented("linearBptreeS->Put")
}

func (t *linearBptreeS) Del(i int) (LinearBptree, bool) {
	return nil, false
}

// lbptIntrNodeS as BlockI
//
func (n *lbptIntrNodeS) TypeAndId() BlockTypeId {
	return n.btid
}

func (n *lbptIntrNodeS) Type() BlockType {
	return btid2Type(n.btid)
}

func (n *lbptIntrNodeS) Id() BlockId {
	return BlockId(uint64(n.btid) & BlockIdMask)
}

func (n *lbptIntrNodeS) Marshal() ([]byte, error) {
	return nil, NotImplemented("*lbptIntrNodeS->Marshal")
}

func (n *lbptIntrNodeS) Write(p []byte) (int, error) {
	return -1, NotImplemented("*lbptIntrNodeS->Write")
}

// lbptIntrNodeS as linBptNodeI
//
//FIXME: I don't know if it is necessary to have an eq() or equals() function
func (n *lbptIntrNodeS) eq(o *lbptIntrNodeS) bool {
	if n == o {
		return true
	}
	if n == nil || o == nil {
		return false
	}
	if n.depth_ != o.depth_ || n.order_ != o.order_ {
		return false
	}
	if len(n.vals) != len(o.vals) {
		return false
	}
	for i := 0; i < len(n.vals); i++ {
		if !n.vals[i].equals(o.vals[i]) {
			return false
		}
	}
	return true
}

func (n *lbptIntrNodeS) equals(o linBptNodeI) bool {
	other, ok := o.(*lbptIntrNodeS)
	if !ok {
		return false
	}

	if !n.eq(other) {
		return false
	}
	return true
}

func (n *lbptIntrNodeS) isToBig() bool {
	if len(n.vals) > n.order_ {
		return true
	}
	return false
}

func (n *lbptIntrNodeS) isLeaf() bool {
	return false
}

func (n *lbptIntrNodeS) findLeftMost() linBptNodeI {
	for i := 0; i < len(n.vals); i++ {
		if n.vals[i] != nil {
			return n.vals[i].findLeftMost()
		}
	}
	return nil
}

func (n *lbptIntrNodeS) size() int {
	return len(n.vals)
}

func (n *lbptIntrNodeS) order() int {
	return n.order_
}

func (n *lbptIntrNodeS) halfFullSize() int {
	return n.order_ / 2
}

func (n *lbptIntrNodeS) depth() int {
	return n.depth_
}

// lbptLeafNodeS as BlockI
//
func (t *lbptLeafNodeS) TypeAndId() BlockTypeId {
	return t.btid
}

func (t *lbptLeafNodeS) Type() BlockType {
	return btid2Type(t.btid)
}

func (t *lbptLeafNodeS) Id() BlockId {
	return BlockId(uint64(t.btid) & BlockIdMask)
}

func (t *lbptLeafNodeS) Marshal() ([]byte, error) {
	return nil, NotImplemented("*lbptLeafNodeS->Marshal")
}

func (t *lbptLeafNodeS) Write(p []byte) (int, error) {
	return -1, NotImplemented("*lbptLeafNodeS->Write")
}

// lbptLeafNodeS as libBptNodeI
//
//FIXME: I don't know if it is necessary to have an eq() or equals() function
func (n *lbptLeafNodeS) eq(o *lbptLeafNodeS) bool {
	if n == o {
		return true
	}
	if n == nil || o == nil {
		return false
	}
	if n.depth_ != o.depth_ || n.order_ != o.order_ {
		return false
	}
	if len(n.vals) != len(o.vals) {
		return false
	}
	for i := 0; i < len(n.vals); i++ {
		//this is weird; converting two generic []byte values to a string to compare them.
		if string(n.vals[i]) != string(o.vals[i]) {
			return false
		}
	}
	return true
}

func (n *lbptLeafNodeS) equals(o linBptNodeI) bool {
	other, ok := o.(*lbptLeafNodeS)
	if !ok {
		return false
	}
	if !n.eq(other) {
		return false
	}
	return true
}

func (n *lbptLeafNodeS) isToBig() bool {
	if len(n.vals) > n.order_ {
		return true
	}
	return false
}

func (n *lbptLeafNodeS) isLeaf() bool {
	return false
}

func (n *lbptLeafNodeS) findLeftMost() linBptNodeI {
	return n
}

func (n *lbptLeafNodeS) size() int {
	return len(n.vals)
}

func (n *lbptLeafNodeS) order() int {
	return n.order_
}

func (n *lbptLeafNodeS) halfFullSize() int {
	return n.order_ / 2
}

func (n *lbptLeafNodeS) depth() int {
	return n.depth_
}
