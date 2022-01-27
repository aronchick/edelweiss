package iplddata

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Byte byte

func (Byte) Def() def.Type {
	return def.Byte{}
}

func (v *Byte) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_Int {
		return ErrNA
	} else {
		if i, _ := n.AsInt(); int64(byte(i)) != i {
			return ErrNA
		} else {
			*(*byte)(v) = byte(i)
			return nil
		}
	}
}

func (v Byte) Node() datamodel.Node {
	return v
}

// datamodel.Node implementation

func (Byte) Kind() datamodel.Kind {
	return datamodel.Kind_Int
}

func (Byte) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Byte) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Byte) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Byte) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Byte) MapIterator() datamodel.MapIterator {
	return nil
}

func (Byte) ListIterator() datamodel.ListIterator {
	return nil
}

func (Byte) Length() int64 {
	return -1
}

func (Byte) IsAbsent() bool {
	return false
}

func (Byte) IsNull() bool {
	return false
}

func (Byte) AsBool() (bool, error) {
	return false, ErrNA
}

func (v Byte) AsInt() (int64, error) {
	return int64(v), nil
}

func (Byte) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (Byte) AsString() (string, error) {
	return "", ErrNA
}

func (Byte) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Byte) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Byte) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

func TryParseByte(n datamodel.Node) (Byte, error) {
	var x Byte
	return x, x.Parse(n)
}
