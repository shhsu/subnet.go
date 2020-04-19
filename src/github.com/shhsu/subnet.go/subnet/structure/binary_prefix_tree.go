package structure

const one PrefixKey = 1
const shift uint64 = PrefixKeyBits - 1

func IsSetAt(key PrefixKey, digit int) bool {
	return ((one << (shift - uint64(digit))) & key) != 0
}

func ChildIndex(key PrefixKey, digit int) int {
	if IsSetAt(key, digit) {
		return 1
	}
	return 0
}

type BinaryPrefixTree interface {
	AddOrReplace(PrefixKey, int, interface{}) interface{}
	Get(PrefixKey) interface{}
}

type BinaryNode struct {
	Data     interface{}
	Children [](*BinaryNode)
}

type BasicBinaryPrefixTree struct {
	root *BinaryNode
}

func NewBasicBinaryPrefixTree() BinaryPrefixTree {
	return &BasicBinaryPrefixTree{
		root: newNode(),
	}
}

func newNode() (t *BinaryNode) {
	return &BinaryNode{}
}

func (t *BasicBinaryPrefixTree) AddOrReplace(key PrefixKey, ranges int, value interface{}) (replaced interface{}) {
	node, _ := t.Locate(key, ranges, true)
	replaced = node.Data
	node.Data = value
	return
}

func (t *BasicBinaryPrefixTree) Get(key PrefixKey) interface{} {
	_, value := t.Locate(key, 64, false)
	return value
}

func (t *BasicBinaryPrefixTree) Locate(key PrefixKey, ranges int, create bool) (cursor *BinaryNode, lastValue interface{}) {
	cursor = t.root
	lastValue = nil
	for i := 0; i < ranges; i++ {
		childIndex := ChildIndex(key, i)
		if cursor.Children[childIndex] == nil {
			if create {
				cursor.Children[childIndex] = newNode()
			} else {
				break
			}
		}
		cursor = cursor.Children[childIndex]
		if cursor != nil && cursor.Data != nil {
			lastValue = cursor.Data
		}
	}
	return
}
