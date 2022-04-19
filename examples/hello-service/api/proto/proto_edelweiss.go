package proto

import (
	pd9 "bytes"
	pd8 "context"
	pd10 "errors"
	pd2 "fmt"
	pd7 "io"
	pd4 "net/http"
	pd12 "net/url"

	pd5 "github.com/ipfs/go-log"
	pd13 "github.com/ipld/edelweiss/services"
	pd1 "github.com/ipld/edelweiss/values"
	pd6 "github.com/ipld/go-ipld-prime"
	pd11 "github.com/ipld/go-ipld-prime/codec/dagjson"
	pd3 "github.com/ipld/go-ipld-prime/datamodel"
)

// -- protocol type Hello_IdentifyArg --

type Hello_IdentifyArg struct {
}

func (x Hello_IdentifyArg) Node() pd3.Node {
	return x
}

func (x *Hello_IdentifyArg) Parse(n pd3.Node) error {
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd2.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd3.Null); err != nil {
			return err
		}
	}
	return nil
}

type Hello_IdentifyArg_MapIterator struct {
	i int64
	s *Hello_IdentifyArg
}

func (x *Hello_IdentifyArg_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd1.ErrNA
}

func (x *Hello_IdentifyArg_MapIterator) Done() bool {
	return x.i+1 >= 0
}

func (x Hello_IdentifyArg) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x Hello_IdentifyArg) LookupByString(key string) (pd3.Node, error) {
	switch key {

	}
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyArg) LookupByNode(key pd3.Node) (pd3.Node, error) {
	switch key.Kind() {
	case pd3.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd3.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyArg) LookupByIndex(idx int64) (pd3.Node, error) {
	switch idx {

	}
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyArg) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {

	}
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyArg) MapIterator() pd3.MapIterator {
	return &Hello_IdentifyArg_MapIterator{-1, &x}
}

func (x Hello_IdentifyArg) ListIterator() pd3.ListIterator {
	return nil
}

func (x Hello_IdentifyArg) Length() int64 {
	return 0
}

func (x Hello_IdentifyArg) IsAbsent() bool {
	return false
}

func (x Hello_IdentifyArg) IsNull() bool {
	return false
}

func (x Hello_IdentifyArg) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x Hello_IdentifyArg) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Hello_IdentifyArg) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Hello_IdentifyArg) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Hello_IdentifyArg) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyArg) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyArg) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type AnonList1 --

type AnonList1 []pd1.String

func (v AnonList1) Node() pd3.Node {
	return v
}

func (v *AnonList1) Parse(n pd3.Node) error {
	if n.Kind() == pd3.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd3.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList1, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList1) Kind() pd3.Kind {
	return pd3.Kind_List
}

func (AnonList1) LookupByString(string) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) LookupByNode(key pd3.Node) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList1) LookupByIndex(i int64) (pd3.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList1) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList1) MapIterator() pd3.MapIterator {
	return nil
}

func (v AnonList1) ListIterator() pd3.ListIterator {
	return &AnonList1_ListIterator{v, 0}
}

func (v AnonList1) Length() int64 {
	return int64(len(v))
}

func (AnonList1) IsAbsent() bool {
	return false
}

func (AnonList1) IsNull() bool {
	return false
}

func (v AnonList1) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (AnonList1) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList1) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList1) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList1) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) Prototype() pd3.NodePrototype {
	return nil // not needed
}

type AnonList1_ListIterator struct {
	list AnonList1
	at   int64
}

func (iter *AnonList1_ListIterator) Next() (int64, pd3.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList1_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}

// -- protocol type Hello_IdentifyResult --

type Hello_IdentifyResult struct {
	Methods AnonList1
}

func (x Hello_IdentifyResult) Node() pd3.Node {
	return x
}

func (x *Hello_IdentifyResult) Parse(n pd3.Node) error {
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
		"Methods": x.Methods.Parse,
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd2.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
				case "Methods":
					if _, notParsed := fieldMap["Methods"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "Methods")
					}
					if err := x.Methods.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Methods")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd3.Null); err != nil {
			return err
		}
	}
	return nil
}

type Hello_IdentifyResult_MapIterator struct {
	i int64
	s *Hello_IdentifyResult
}

func (x *Hello_IdentifyResult_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Methods"), x.s.Methods.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *Hello_IdentifyResult_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x Hello_IdentifyResult) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x Hello_IdentifyResult) LookupByString(key string) (pd3.Node, error) {
	switch key {
	case "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyResult) LookupByNode(key pd3.Node) (pd3.Node, error) {
	switch key.Kind() {
	case pd3.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd3.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyResult) LookupByIndex(idx int64) (pd3.Node, error) {
	switch idx {
	case 0:
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyResult) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {
	case "0", "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyResult) MapIterator() pd3.MapIterator {
	return &Hello_IdentifyResult_MapIterator{-1, &x}
}

func (x Hello_IdentifyResult) ListIterator() pd3.ListIterator {
	return nil
}

func (x Hello_IdentifyResult) Length() int64 {
	return 1
}

func (x Hello_IdentifyResult) IsAbsent() bool {
	return false
}

func (x Hello_IdentifyResult) IsNull() bool {
	return false
}

func (x Hello_IdentifyResult) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x Hello_IdentifyResult) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Hello_IdentifyResult) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Hello_IdentifyResult) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Hello_IdentifyResult) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyResult) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x Hello_IdentifyResult) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type Hello_Error --

type Hello_Error struct {
	Code pd1.String
}

func (x Hello_Error) Node() pd3.Node {
	return x
}

func (x *Hello_Error) Parse(n pd3.Node) error {
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
		"Code": x.Code.Parse,
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd2.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
				case "Code":
					if _, notParsed := fieldMap["Code"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "Code")
					}
					if err := x.Code.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Code")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd3.Null); err != nil {
			return err
		}
	}
	return nil
}

type Hello_Error_MapIterator struct {
	i int64
	s *Hello_Error
}

func (x *Hello_Error_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Code"), x.s.Code.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *Hello_Error_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x Hello_Error) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x Hello_Error) LookupByString(key string) (pd3.Node, error) {
	switch key {
	case "Code":
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Hello_Error) LookupByNode(key pd3.Node) (pd3.Node, error) {
	switch key.Kind() {
	case pd3.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd3.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x Hello_Error) LookupByIndex(idx int64) (pd3.Node, error) {
	switch idx {
	case 0:
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Hello_Error) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {
	case "0", "Code":
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Hello_Error) MapIterator() pd3.MapIterator {
	return &Hello_Error_MapIterator{-1, &x}
}

func (x Hello_Error) ListIterator() pd3.ListIterator {
	return nil
}

func (x Hello_Error) Length() int64 {
	return 1
}

func (x Hello_Error) IsAbsent() bool {
	return false
}

func (x Hello_Error) IsNull() bool {
	return false
}

func (x Hello_Error) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x Hello_Error) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Hello_Error) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Hello_Error) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Hello_Error) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Hello_Error) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x Hello_Error) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type AnonInductive4 --

type AnonInductive4 struct {
	Identify *Hello_IdentifyArg
	Hello    *HelloRequest
}

func (x *AnonInductive4) Parse(n pd3.Node) error {
	*x = AnonInductive4{}
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd2.Errorf("inductive map key is not a string")
	}
	switch k {
	case "IdentifyRequest":
		var y Hello_IdentifyArg
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Identify = &y
		return nil
	case "HelloRequest":
		var y HelloRequest
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Hello = &y
		return nil

	}

	return pd2.Errorf("inductive map has no applicable keys")

}

type AnonInductive4_MapIterator struct {
	done bool
	s    *AnonInductive4
}

func (x *AnonInductive4_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
		case x.s.Identify != nil:
			return pd1.String("IdentifyRequest"), x.s.Identify.Node(), nil
		case x.s.Hello != nil:
			return pd1.String("HelloRequest"), x.s.Hello.Node(), nil

		default:
			return nil, nil, pd2.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive4_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive4) Node() pd3.Node {
	return x
}

func (x AnonInductive4) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x AnonInductive4) LookupByString(key string) (pd3.Node, error) {
	switch {
	case x.Identify != nil && key == "IdentifyRequest":
		return x.Identify.Node(), nil
	case x.Hello != nil && key == "HelloRequest":
		return x.Hello.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive4) LookupByNode(key pd3.Node) (pd3.Node, error) {
	if key.Kind() != pd3.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive4) LookupByIndex(idx int64) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {
	case "IdentifyRequest":
		return x.Identify.Node(), nil
	case "HelloRequest":
		return x.Hello.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive4) MapIterator() pd3.MapIterator {
	return &AnonInductive4_MapIterator{false, &x}
}

func (x AnonInductive4) ListIterator() pd3.ListIterator {
	return nil
}

func (x AnonInductive4) Length() int64 {
	return 1
}

func (x AnonInductive4) IsAbsent() bool {
	return false
}

func (x AnonInductive4) IsNull() bool {
	return false
}

func (x AnonInductive4) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x AnonInductive4) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive4) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive4) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x AnonInductive4) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type AnonInductive5 --

type AnonInductive5 struct {
	Identify *Hello_IdentifyResult
	Hello    *HelloResponse
	Error    *Hello_Error
}

func (x *AnonInductive5) Parse(n pd3.Node) error {
	*x = AnonInductive5{}
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd2.Errorf("inductive map key is not a string")
	}
	switch k {
	case "IdentifyResponse":
		var y Hello_IdentifyResult
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Identify = &y
		return nil
	case "HelloResponse":
		var y HelloResponse
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Hello = &y
		return nil
	case "Error":
		var y Hello_Error
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Error = &y
		return nil

	}

	return pd2.Errorf("inductive map has no applicable keys")

}

type AnonInductive5_MapIterator struct {
	done bool
	s    *AnonInductive5
}

func (x *AnonInductive5_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
		case x.s.Identify != nil:
			return pd1.String("IdentifyResponse"), x.s.Identify.Node(), nil
		case x.s.Hello != nil:
			return pd1.String("HelloResponse"), x.s.Hello.Node(), nil
		case x.s.Error != nil:
			return pd1.String("Error"), x.s.Error.Node(), nil

		default:
			return nil, nil, pd2.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive5_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive5) Node() pd3.Node {
	return x
}

func (x AnonInductive5) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x AnonInductive5) LookupByString(key string) (pd3.Node, error) {
	switch {
	case x.Identify != nil && key == "IdentifyResponse":
		return x.Identify.Node(), nil
	case x.Hello != nil && key == "HelloResponse":
		return x.Hello.Node(), nil
	case x.Error != nil && key == "Error":
		return x.Error.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive5) LookupByNode(key pd3.Node) (pd3.Node, error) {
	if key.Kind() != pd3.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive5) LookupByIndex(idx int64) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {
	case "IdentifyResponse":
		return x.Identify.Node(), nil
	case "HelloResponse":
		return x.Hello.Node(), nil
	case "Error":
		return x.Error.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive5) MapIterator() pd3.MapIterator {
	return &AnonInductive5_MapIterator{false, &x}
}

func (x AnonInductive5) ListIterator() pd3.ListIterator {
	return nil
}

func (x AnonInductive5) Length() int64 {
	return 1
}

func (x AnonInductive5) IsAbsent() bool {
	return false
}

func (x AnonInductive5) IsNull() bool {
	return false
}

func (x AnonInductive5) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x AnonInductive5) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive5) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive5) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x AnonInductive5) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) Prototype() pd3.NodePrototype {
	return nil
}

var logger_client_Hello = pd5.Logger("service/client/Hello")

type Hello_Client interface {
	Identify(ctx pd8.Context, req *Hello_IdentifyArg) ([]*Hello_IdentifyResult, error)

	Hello(ctx pd8.Context, req *HelloRequest) ([]*HelloResponse, error)

	Identify_Async(ctx pd8.Context, req *Hello_IdentifyArg) (<-chan Hello_Identify_AsyncResult, error)

	Hello_Async(ctx pd8.Context, req *HelloRequest) (<-chan Hello_Hello_AsyncResult, error)
}

type Hello_Identify_AsyncResult struct {
	Resp *Hello_IdentifyResult
	Err  error
}

type Hello_Hello_AsyncResult struct {
	Resp *HelloResponse
	Err  error
}

type Hello_ClientOption func(*client_Hello) error

type client_Hello struct {
	httpClient *pd4.Client
	endpoint   *pd12.URL
}

func Hello_Client_WithHTTPClient(hc *pd4.Client) Hello_ClientOption {
	return func(c *client_Hello) error {
		c.httpClient = hc
		return nil
	}
}

func New_Hello_Client(endpoint string, opts ...Hello_ClientOption) (*client_Hello, error) {
	u, err := pd12.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &client_Hello{endpoint: u, httpClient: pd4.DefaultClient}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *client_Hello) Identify(ctx pd8.Context, req *Hello_IdentifyArg) ([]*Hello_IdentifyResult, error) {
	ctx, cancel := pd8.WithCancel(ctx)
	defer cancel()
	ch, err := c.Identify_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*Hello_IdentifyResult
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					logger_client_Hello.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_Hello) Identify_Async(ctx pd8.Context, req *Hello_IdentifyArg) (<-chan Hello_Identify_AsyncResult, error) {
	envelope := &AnonInductive4{
		Identify: req,
	}

	buf, err := pd6.Encode(envelope, pd11.Encode)
	if err != nil {
		return nil, pd2.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd12.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd4.NewRequestWithContext(ctx, "POST", u.String(), pd9.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			"application/vnd.ipfs.rpc+dag-json; version=1",
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd2.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan Hello_Identify_AsyncResult, 1)
	go process_Hello_Identify_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_Hello_Identify_AsyncResult(ctx pd8.Context, ch chan<- Hello_Identify_AsyncResult, r pd7.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- Hello_Identify_AsyncResult{Err: pd13.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd6.DecodeStreaming(r, pd11.Decode)
		if pd10.Is(err, pd7.EOF) || pd10.Is(err, pd7.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- Hello_Identify_AsyncResult{Err: pd13.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- Hello_Identify_AsyncResult{Err: pd13.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- Hello_Identify_AsyncResult{Err: pd13.ErrService{Cause: pd10.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.Identify == nil {
			continue
		}
		ch <- Hello_Identify_AsyncResult{Resp: env.Identify}
	}
}

func (c *client_Hello) Hello(ctx pd8.Context, req *HelloRequest) ([]*HelloResponse, error) {
	ctx, cancel := pd8.WithCancel(ctx)
	defer cancel()
	ch, err := c.Hello_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*HelloResponse
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					logger_client_Hello.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_Hello) Hello_Async(ctx pd8.Context, req *HelloRequest) (<-chan Hello_Hello_AsyncResult, error) {
	envelope := &AnonInductive4{
		Hello: req,
	}

	buf, err := pd6.Encode(envelope, pd11.Encode)
	if err != nil {
		return nil, pd2.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd12.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd4.NewRequestWithContext(ctx, "POST", u.String(), pd9.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			"application/vnd.ipfs.rpc+dag-json; version=1",
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd2.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan Hello_Hello_AsyncResult, 1)
	go process_Hello_Hello_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_Hello_Hello_AsyncResult(ctx pd8.Context, ch chan<- Hello_Hello_AsyncResult, r pd7.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- Hello_Hello_AsyncResult{Err: pd13.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd6.DecodeStreaming(r, pd11.Decode)
		if pd10.Is(err, pd7.EOF) || pd10.Is(err, pd7.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- Hello_Hello_AsyncResult{Err: pd13.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- Hello_Hello_AsyncResult{Err: pd13.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- Hello_Hello_AsyncResult{Err: pd13.ErrService{Cause: pd10.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.Hello == nil {
			continue
		}
		ch <- Hello_Hello_AsyncResult{Resp: env.Hello}
	}
}

var logger_server_Hello = pd5.Logger("service/server/Hello")

type Hello_Server interface {
	Hello(ctx pd8.Context, req *HelloRequest, respCh chan<- *Hello_Hello_AsyncResult) error
}

func Hello_AsyncHandler(s Hello_Server) pd4.HandlerFunc {
	return func(writer pd4.ResponseWriter, request *pd4.Request) {
		// parse request
		msg := request.URL.Query().Get("q")
		n, err := pd6.Decode([]byte(msg), pd11.Decode)
		if err != nil {
			logger_server_Hello.Errorf("received request not decodeable (%v)", err)
			writer.WriteHeader(400)
			return
		}
		env := &AnonInductive4{}
		if err = env.Parse(n); err != nil {
			logger_server_Hello.Errorf("parsing call envelope (%v)", err)
			writer.WriteHeader(400)
			return
		}

		writer.Header()["Content-Type"] = []string{
			"application/vnd.ipfs.rpc+dag-json; version=1",
		}

		// demultiplex request
		switch {

		case env.Hello != nil:
			ch := make(chan *Hello_Hello_AsyncResult)
			if err = s.Hello(pd8.Background(), env.Hello, ch); err != nil {
				logger_server_Hello.Errorf("get p2p provider rejected request (%v)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				var env *AnonInductive5
				if resp.Err != nil {
					env = &AnonInductive5{Error: &Hello_Error{Code: pd1.String(resp.Err.Error())}}
				} else {
					env = &AnonInductive5{Hello: resp.Resp}
				}
				var buf pd9.Buffer
				if err = pd6.EncodeStreaming(&buf, env, pd11.Encode); err != nil {
					logger_server_Hello.Errorf("cannot encode response (%v)", err)
					continue
				}
				buf.WriteByte("\n"[0])
				writer.Write(buf.Bytes())
			}

		case env.Identify != nil:
			var env *AnonInductive5
			env = &AnonInductive5{
				Identify: &Hello_IdentifyResult{
					Methods: []pd1.String{
						"Hello",
					},
				},
			}
			var buf pd9.Buffer
			if err = pd6.EncodeStreaming(&buf, env, pd11.Encode); err != nil {
				logger_server_Hello.Errorf("cannot encode identify response (%v)", err)
				writer.WriteHeader(500)
				return
			}
			buf.WriteByte("\n"[0])
			writer.Write(buf.Bytes())

		default:
			logger_server_Hello.Errorf("missing or unknown request")
			writer.WriteHeader(404)
		}
	}
}

// -- protocol type AnonList7 --

type AnonList7 []pd1.String

func (v AnonList7) Node() pd3.Node {
	return v
}

func (v *AnonList7) Parse(n pd3.Node) error {
	if n.Kind() == pd3.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd3.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList7, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList7) Kind() pd3.Kind {
	return pd3.Kind_List
}

func (AnonList7) LookupByString(string) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList7) LookupByNode(key pd3.Node) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList7) LookupByIndex(i int64) (pd3.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList7) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList7) MapIterator() pd3.MapIterator {
	return nil
}

func (v AnonList7) ListIterator() pd3.ListIterator {
	return &AnonList7_ListIterator{v, 0}
}

func (v AnonList7) Length() int64 {
	return int64(len(v))
}

func (AnonList7) IsAbsent() bool {
	return false
}

func (AnonList7) IsNull() bool {
	return false
}

func (v AnonList7) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (AnonList7) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList7) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList7) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList7) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList7) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList7) Prototype() pd3.NodePrototype {
	return nil // not needed
}

type AnonList7_ListIterator struct {
	list AnonList7
	at   int64
}

func (iter *AnonList7_ListIterator) Next() (int64, pd3.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList7_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}

// -- protocol type AnonInductive8 --

type AnonInductive8 struct {
	US *USAddress
	SK *SKAddress

	Other   string
	Address *AnonList7
}

func (x *AnonInductive8) Parse(n pd3.Node) error {
	*x = AnonInductive8{}
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd2.Errorf("inductive map key is not a string")
	}
	switch k {
	case "US":
		var y USAddress
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.US = &y
		return nil
	case "SouthKorea":
		var y SKAddress
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.SK = &y
		return nil

	default:
		var y AnonList7
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Other = k
		x.Address = &y
		return nil

	}

}

type AnonInductive8_MapIterator struct {
	done bool
	s    *AnonInductive8
}

func (x *AnonInductive8_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
		case x.s.US != nil:
			return pd1.String("US"), x.s.US.Node(), nil
		case x.s.SK != nil:
			return pd1.String("SouthKorea"), x.s.SK.Node(), nil

		case x.s.Address != nil:
			return pd1.String(x.s.Other), x.s.Address.Node(), nil

		default:
			return nil, nil, pd2.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive8_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive8) Node() pd3.Node {
	return x
}

func (x AnonInductive8) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x AnonInductive8) LookupByString(key string) (pd3.Node, error) {
	switch {
	case x.US != nil && key == "US":
		return x.US.Node(), nil
	case x.SK != nil && key == "SouthKorea":
		return x.SK.Node(), nil

	case x.Address != nil && key == x.Other:
		return x.Address.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive8) LookupByNode(key pd3.Node) (pd3.Node, error) {
	if key.Kind() != pd3.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive8) LookupByIndex(idx int64) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive8) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {
	case "US":
		return x.US.Node(), nil
	case "SouthKorea":
		return x.SK.Node(), nil

	case x.Other:
		return x.Address.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive8) MapIterator() pd3.MapIterator {
	return &AnonInductive8_MapIterator{false, &x}
}

func (x AnonInductive8) ListIterator() pd3.ListIterator {
	return nil
}

func (x AnonInductive8) Length() int64 {
	return 1
}

func (x AnonInductive8) IsAbsent() bool {
	return false
}

func (x AnonInductive8) IsNull() bool {
	return false
}

func (x AnonInductive8) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x AnonInductive8) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive8) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive8) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x AnonInductive8) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive8) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive8) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type HelloRequest --

type HelloRequest struct {
	Name    pd1.String
	Address AnonInductive8
}

func (x HelloRequest) Node() pd3.Node {
	return x
}

func (x *HelloRequest) Parse(n pd3.Node) error {
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
		"Name":    x.Name.Parse,
		"Address": x.Address.Parse,
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd2.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
				case "Name":
					if _, notParsed := fieldMap["Name"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "Name")
					}
					if err := x.Name.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Name")
				case "Address":
					if _, notParsed := fieldMap["Address"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "Address")
					}
					if err := x.Address.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Address")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd3.Null); err != nil {
			return err
		}
	}
	return nil
}

type HelloRequest_MapIterator struct {
	i int64
	s *HelloRequest
}

func (x *HelloRequest_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Name"), x.s.Name.Node(), nil
	case 1:
		return pd1.String("Address"), x.s.Address.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *HelloRequest_MapIterator) Done() bool {
	return x.i+1 >= 2
}

func (x HelloRequest) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x HelloRequest) LookupByString(key string) (pd3.Node, error) {
	switch key {
	case "Name":
		return x.Name.Node(), nil
	case "Address":
		return x.Address.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x HelloRequest) LookupByNode(key pd3.Node) (pd3.Node, error) {
	switch key.Kind() {
	case pd3.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd3.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x HelloRequest) LookupByIndex(idx int64) (pd3.Node, error) {
	switch idx {
	case 0:
		return x.Name.Node(), nil
	case 1:
		return x.Address.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x HelloRequest) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {
	case "0", "Name":
		return x.Name.Node(), nil
	case "1", "Address":
		return x.Address.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x HelloRequest) MapIterator() pd3.MapIterator {
	return &HelloRequest_MapIterator{-1, &x}
}

func (x HelloRequest) ListIterator() pd3.ListIterator {
	return nil
}

func (x HelloRequest) Length() int64 {
	return 2
}

func (x HelloRequest) IsAbsent() bool {
	return false
}

func (x HelloRequest) IsNull() bool {
	return false
}

func (x HelloRequest) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x HelloRequest) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x HelloRequest) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x HelloRequest) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x HelloRequest) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x HelloRequest) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x HelloRequest) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type USAddress --

type USAddress struct {
	Street pd1.String
	City   pd1.String
	State  State
	ZIP    pd1.Int
}

func (x USAddress) Node() pd3.Node {
	return x
}

func (x *USAddress) Parse(n pd3.Node) error {
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
		"street": x.Street.Parse,
		"city":   x.City.Parse,
		"state":  x.State.Parse,
		"zip":    x.ZIP.Parse,
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd2.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
				case "street":
					if _, notParsed := fieldMap["street"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "street")
					}
					if err := x.Street.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "street")
				case "city":
					if _, notParsed := fieldMap["city"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "city")
					}
					if err := x.City.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "city")
				case "state":
					if _, notParsed := fieldMap["state"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "state")
					}
					if err := x.State.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "state")
				case "zip":
					if _, notParsed := fieldMap["zip"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "zip")
					}
					if err := x.ZIP.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "zip")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd3.Null); err != nil {
			return err
		}
	}
	return nil
}

type USAddress_MapIterator struct {
	i int64
	s *USAddress
}

func (x *USAddress_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("street"), x.s.Street.Node(), nil
	case 1:
		return pd1.String("city"), x.s.City.Node(), nil
	case 2:
		return pd1.String("state"), x.s.State.Node(), nil
	case 3:
		return pd1.String("zip"), x.s.ZIP.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *USAddress_MapIterator) Done() bool {
	return x.i+1 >= 4
}

func (x USAddress) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x USAddress) LookupByString(key string) (pd3.Node, error) {
	switch key {
	case "street":
		return x.Street.Node(), nil
	case "city":
		return x.City.Node(), nil
	case "state":
		return x.State.Node(), nil
	case "zip":
		return x.ZIP.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x USAddress) LookupByNode(key pd3.Node) (pd3.Node, error) {
	switch key.Kind() {
	case pd3.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd3.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x USAddress) LookupByIndex(idx int64) (pd3.Node, error) {
	switch idx {
	case 0:
		return x.Street.Node(), nil
	case 1:
		return x.City.Node(), nil
	case 2:
		return x.State.Node(), nil
	case 3:
		return x.ZIP.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x USAddress) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {
	case "0", "street":
		return x.Street.Node(), nil
	case "1", "city":
		return x.City.Node(), nil
	case "2", "state":
		return x.State.Node(), nil
	case "3", "zip":
		return x.ZIP.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x USAddress) MapIterator() pd3.MapIterator {
	return &USAddress_MapIterator{-1, &x}
}

func (x USAddress) ListIterator() pd3.ListIterator {
	return nil
}

func (x USAddress) Length() int64 {
	return 4
}

func (x USAddress) IsAbsent() bool {
	return false
}

func (x USAddress) IsNull() bool {
	return false
}

func (x USAddress) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x USAddress) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x USAddress) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x USAddress) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x USAddress) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x USAddress) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x USAddress) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type AnonSingletonString11 --

type AnonSingletonString11 struct{}

func (AnonSingletonString11) Parse(n pd3.Node) error {
	if n.Kind() != pd3.Kind_String {
		return pd1.ErrNA
	}
	v, err := n.AsString()
	if err != nil {
		return err
	}
	if v != "CA" {
		return pd1.ErrNA
	}
	return nil
}

func (v AnonSingletonString11) Node() pd3.Node {
	return v
}

func (AnonSingletonString11) Kind() pd3.Kind {
	return pd3.Kind_String
}

func (AnonSingletonString11) LookupByString(string) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString11) LookupByNode(key pd3.Node) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString11) LookupByIndex(idx int64) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString11) LookupBySegment(_ pd3.PathSegment) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString11) MapIterator() pd3.MapIterator {
	return nil
}

func (AnonSingletonString11) ListIterator() pd3.ListIterator {
	return nil
}

func (AnonSingletonString11) Length() int64 {
	return -1
}

func (AnonSingletonString11) IsAbsent() bool {
	return false
}

func (AnonSingletonString11) IsNull() bool {
	return false
}

func (v AnonSingletonString11) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (AnonSingletonString11) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonSingletonString11) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonSingletonString11) AsString() (string, error) {
	return "CA", nil
}

func (AnonSingletonString11) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString11) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString11) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type AnonSingletonString12 --

type AnonSingletonString12 struct{}

func (AnonSingletonString12) Parse(n pd3.Node) error {
	if n.Kind() != pd3.Kind_String {
		return pd1.ErrNA
	}
	v, err := n.AsString()
	if err != nil {
		return err
	}
	if v != "NY" {
		return pd1.ErrNA
	}
	return nil
}

func (v AnonSingletonString12) Node() pd3.Node {
	return v
}

func (AnonSingletonString12) Kind() pd3.Kind {
	return pd3.Kind_String
}

func (AnonSingletonString12) LookupByString(string) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString12) LookupByNode(key pd3.Node) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString12) LookupByIndex(idx int64) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString12) LookupBySegment(_ pd3.PathSegment) (pd3.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString12) MapIterator() pd3.MapIterator {
	return nil
}

func (AnonSingletonString12) ListIterator() pd3.ListIterator {
	return nil
}

func (AnonSingletonString12) Length() int64 {
	return -1
}

func (AnonSingletonString12) IsAbsent() bool {
	return false
}

func (AnonSingletonString12) IsNull() bool {
	return false
}

func (v AnonSingletonString12) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (AnonSingletonString12) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonSingletonString12) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonSingletonString12) AsString() (string, error) {
	return "NY", nil
}

func (AnonSingletonString12) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString12) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonSingletonString12) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type State --

type State struct {
	CA    *AnonSingletonString11
	NY    *AnonSingletonString12
	Other *pd1.String
}

func (x *State) Parse(n pd3.Node) error {
	*x = State{}

	var CA AnonSingletonString11
	if err := CA.Parse(n); err == nil {
		x.CA = &CA
		return nil
	}

	var NY AnonSingletonString12
	if err := NY.Parse(n); err == nil {
		x.NY = &NY
		return nil
	}

	var Other pd1.String
	if err := Other.Parse(n); err == nil {
		x.Other = &Other
		return nil
	}

	return pd2.Errorf("no union cases parses")
}

func (x State) Node() pd3.Node {
	if x.CA != nil {
		return x.CA
	}
	if x.NY != nil {
		return x.NY
	}
	if x.Other != nil {
		return x.Other
	}

	return nil
}

// proxy Node methods to active case

func (x State) Kind() pd3.Kind {
	if x.CA != nil {
		return x.CA.Kind()
	}
	if x.NY != nil {
		return x.NY.Kind()
	}
	if x.Other != nil {
		return x.Other.Kind()
	}

	return pd3.Kind_Invalid
}

func (x State) LookupByString(key string) (pd3.Node, error) {
	if x.CA != nil {
		return x.CA.LookupByString(key)
	}
	if x.NY != nil {
		return x.NY.LookupByString(key)
	}
	if x.Other != nil {
		return x.Other.LookupByString(key)
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x State) LookupByNode(key pd3.Node) (pd3.Node, error) {
	if x.CA != nil {
		return x.CA.LookupByNode(key)
	}
	if x.NY != nil {
		return x.NY.LookupByNode(key)
	}
	if x.Other != nil {
		return x.Other.LookupByNode(key)
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x State) LookupByIndex(idx int64) (pd3.Node, error) {
	if x.CA != nil {
		return x.CA.LookupByIndex(idx)
	}
	if x.NY != nil {
		return x.NY.LookupByIndex(idx)
	}
	if x.Other != nil {
		return x.Other.LookupByIndex(idx)
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x State) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	if x.CA != nil {
		return x.CA.LookupBySegment(seg)
	}
	if x.NY != nil {
		return x.NY.LookupBySegment(seg)
	}
	if x.Other != nil {
		return x.Other.LookupBySegment(seg)
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x State) MapIterator() pd3.MapIterator {
	if x.CA != nil {
		return x.CA.MapIterator()
	}
	if x.NY != nil {
		return x.NY.MapIterator()
	}
	if x.Other != nil {
		return x.Other.MapIterator()
	}

	return nil
}

func (x State) ListIterator() pd3.ListIterator {
	if x.CA != nil {
		return x.CA.ListIterator()
	}
	if x.NY != nil {
		return x.NY.ListIterator()
	}
	if x.Other != nil {
		return x.Other.ListIterator()
	}

	return nil
}

func (x State) Length() int64 {
	if x.CA != nil {
		return x.CA.Length()
	}
	if x.NY != nil {
		return x.NY.Length()
	}
	if x.Other != nil {
		return x.Other.Length()
	}

	return -1
}

func (x State) IsAbsent() bool {
	if x.CA != nil {
		return x.CA.IsAbsent()
	}
	if x.NY != nil {
		return x.NY.IsAbsent()
	}
	if x.Other != nil {
		return x.Other.IsAbsent()
	}

	return false
}

func (x State) IsNull() bool {
	if x.CA != nil {
		return x.CA.IsNull()
	}
	if x.NY != nil {
		return x.NY.IsNull()
	}
	if x.Other != nil {
		return x.Other.IsNull()
	}

	return false
}

func (x State) AsBool() (bool, error) {
	if x.CA != nil {
		return x.CA.AsBool()
	}
	if x.NY != nil {
		return x.NY.AsBool()
	}
	if x.Other != nil {
		return x.Other.AsBool()
	}

	return false, pd2.Errorf("no active union case found")
}

func (x State) AsInt() (int64, error) {
	if x.CA != nil {
		return x.CA.AsInt()
	}
	if x.NY != nil {
		return x.NY.AsInt()
	}
	if x.Other != nil {
		return x.Other.AsInt()
	}

	return 0, pd2.Errorf("no active union case found")
}

func (x State) AsFloat() (float64, error) {
	if x.CA != nil {
		return x.CA.AsFloat()
	}
	if x.NY != nil {
		return x.NY.AsFloat()
	}
	if x.Other != nil {
		return x.Other.AsFloat()
	}

	return 0.0, pd2.Errorf("no active union case found")
}

func (x State) AsString() (string, error) {
	if x.CA != nil {
		return x.CA.AsString()
	}
	if x.NY != nil {
		return x.NY.AsString()
	}
	if x.Other != nil {
		return x.Other.AsString()
	}

	return "", pd2.Errorf("no active union case found")
}

func (x State) AsBytes() ([]byte, error) {
	if x.CA != nil {
		return x.CA.AsBytes()
	}
	if x.NY != nil {
		return x.NY.AsBytes()
	}
	if x.Other != nil {
		return x.Other.AsBytes()
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x State) AsLink() (pd3.Link, error) {
	if x.CA != nil {
		return x.CA.AsLink()
	}
	if x.NY != nil {
		return x.NY.AsLink()
	}
	if x.Other != nil {
		return x.Other.AsLink()
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x State) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type SKAddress --

type SKAddress struct {
	Street     pd1.String
	City       pd1.String
	Province   pd1.String
	PostalCode pd1.Int
}

func (x SKAddress) Node() pd3.Node {
	return x
}

func (x *SKAddress) Parse(n pd3.Node) error {
	if n.Kind() != pd3.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
		"street":      x.Street.Parse,
		"city":        x.City.Parse,
		"province":    x.Province.Parse,
		"postal_code": x.PostalCode.Parse,
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd2.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
				case "street":
					if _, notParsed := fieldMap["street"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "street")
					}
					if err := x.Street.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "street")
				case "city":
					if _, notParsed := fieldMap["city"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "city")
					}
					if err := x.City.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "city")
				case "province":
					if _, notParsed := fieldMap["province"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "province")
					}
					if err := x.Province.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "province")
				case "postal_code":
					if _, notParsed := fieldMap["postal_code"]; !notParsed {
						return pd2.Errorf("field %s already parsed", "postal_code")
					}
					if err := x.PostalCode.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "postal_code")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd3.Null); err != nil {
			return err
		}
	}
	return nil
}

type SKAddress_MapIterator struct {
	i int64
	s *SKAddress
}

func (x *SKAddress_MapIterator) Next() (key pd3.Node, value pd3.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("street"), x.s.Street.Node(), nil
	case 1:
		return pd1.String("city"), x.s.City.Node(), nil
	case 2:
		return pd1.String("province"), x.s.Province.Node(), nil
	case 3:
		return pd1.String("postal_code"), x.s.PostalCode.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *SKAddress_MapIterator) Done() bool {
	return x.i+1 >= 4
}

func (x SKAddress) Kind() pd3.Kind {
	return pd3.Kind_Map
}

func (x SKAddress) LookupByString(key string) (pd3.Node, error) {
	switch key {
	case "street":
		return x.Street.Node(), nil
	case "city":
		return x.City.Node(), nil
	case "province":
		return x.Province.Node(), nil
	case "postal_code":
		return x.PostalCode.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x SKAddress) LookupByNode(key pd3.Node) (pd3.Node, error) {
	switch key.Kind() {
	case pd3.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd3.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x SKAddress) LookupByIndex(idx int64) (pd3.Node, error) {
	switch idx {
	case 0:
		return x.Street.Node(), nil
	case 1:
		return x.City.Node(), nil
	case 2:
		return x.Province.Node(), nil
	case 3:
		return x.PostalCode.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x SKAddress) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	switch seg.String() {
	case "0", "street":
		return x.Street.Node(), nil
	case "1", "city":
		return x.City.Node(), nil
	case "2", "province":
		return x.Province.Node(), nil
	case "3", "postal_code":
		return x.PostalCode.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x SKAddress) MapIterator() pd3.MapIterator {
	return &SKAddress_MapIterator{-1, &x}
}

func (x SKAddress) ListIterator() pd3.ListIterator {
	return nil
}

func (x SKAddress) Length() int64 {
	return 4
}

func (x SKAddress) IsAbsent() bool {
	return false
}

func (x SKAddress) IsNull() bool {
	return false
}

func (x SKAddress) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x SKAddress) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x SKAddress) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x SKAddress) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x SKAddress) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x SKAddress) AsLink() (pd3.Link, error) {
	return nil, pd1.ErrNA
}

func (x SKAddress) Prototype() pd3.NodePrototype {
	return nil
}

// -- protocol type HelloResponse --

type HelloResponse struct {
	English *pd1.String
	Korean  *pd1.String
}

func (x *HelloResponse) Parse(n pd3.Node) error {
	*x = HelloResponse{}

	var English pd1.String
	if err := English.Parse(n); err == nil {
		x.English = &English
		return nil
	}

	var Korean pd1.String
	if err := Korean.Parse(n); err == nil {
		x.Korean = &Korean
		return nil
	}

	return pd2.Errorf("no union cases parses")
}

func (x HelloResponse) Node() pd3.Node {
	if x.English != nil {
		return x.English
	}
	if x.Korean != nil {
		return x.Korean
	}

	return nil
}

// proxy Node methods to active case

func (x HelloResponse) Kind() pd3.Kind {
	if x.English != nil {
		return x.English.Kind()
	}
	if x.Korean != nil {
		return x.Korean.Kind()
	}

	return pd3.Kind_Invalid
}

func (x HelloResponse) LookupByString(key string) (pd3.Node, error) {
	if x.English != nil {
		return x.English.LookupByString(key)
	}
	if x.Korean != nil {
		return x.Korean.LookupByString(key)
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x HelloResponse) LookupByNode(key pd3.Node) (pd3.Node, error) {
	if x.English != nil {
		return x.English.LookupByNode(key)
	}
	if x.Korean != nil {
		return x.Korean.LookupByNode(key)
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x HelloResponse) LookupByIndex(idx int64) (pd3.Node, error) {
	if x.English != nil {
		return x.English.LookupByIndex(idx)
	}
	if x.Korean != nil {
		return x.Korean.LookupByIndex(idx)
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x HelloResponse) LookupBySegment(seg pd3.PathSegment) (pd3.Node, error) {
	if x.English != nil {
		return x.English.LookupBySegment(seg)
	}
	if x.Korean != nil {
		return x.Korean.LookupBySegment(seg)
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x HelloResponse) MapIterator() pd3.MapIterator {
	if x.English != nil {
		return x.English.MapIterator()
	}
	if x.Korean != nil {
		return x.Korean.MapIterator()
	}

	return nil
}

func (x HelloResponse) ListIterator() pd3.ListIterator {
	if x.English != nil {
		return x.English.ListIterator()
	}
	if x.Korean != nil {
		return x.Korean.ListIterator()
	}

	return nil
}

func (x HelloResponse) Length() int64 {
	if x.English != nil {
		return x.English.Length()
	}
	if x.Korean != nil {
		return x.Korean.Length()
	}

	return -1
}

func (x HelloResponse) IsAbsent() bool {
	if x.English != nil {
		return x.English.IsAbsent()
	}
	if x.Korean != nil {
		return x.Korean.IsAbsent()
	}

	return false
}

func (x HelloResponse) IsNull() bool {
	if x.English != nil {
		return x.English.IsNull()
	}
	if x.Korean != nil {
		return x.Korean.IsNull()
	}

	return false
}

func (x HelloResponse) AsBool() (bool, error) {
	if x.English != nil {
		return x.English.AsBool()
	}
	if x.Korean != nil {
		return x.Korean.AsBool()
	}

	return false, pd2.Errorf("no active union case found")
}

func (x HelloResponse) AsInt() (int64, error) {
	if x.English != nil {
		return x.English.AsInt()
	}
	if x.Korean != nil {
		return x.Korean.AsInt()
	}

	return 0, pd2.Errorf("no active union case found")
}

func (x HelloResponse) AsFloat() (float64, error) {
	if x.English != nil {
		return x.English.AsFloat()
	}
	if x.Korean != nil {
		return x.Korean.AsFloat()
	}

	return 0.0, pd2.Errorf("no active union case found")
}

func (x HelloResponse) AsString() (string, error) {
	if x.English != nil {
		return x.English.AsString()
	}
	if x.Korean != nil {
		return x.Korean.AsString()
	}

	return "", pd2.Errorf("no active union case found")
}

func (x HelloResponse) AsBytes() ([]byte, error) {
	if x.English != nil {
		return x.English.AsBytes()
	}
	if x.Korean != nil {
		return x.Korean.AsBytes()
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x HelloResponse) AsLink() (pd3.Link, error) {
	if x.English != nil {
		return x.English.AsLink()
	}
	if x.Korean != nil {
		return x.Korean.AsLink()
	}

	return nil, pd2.Errorf("no active union case found")
}

func (x HelloResponse) Prototype() pd3.NodePrototype {
	return nil
}
