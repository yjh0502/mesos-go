// Code generated by protoc-gen-gogo.
// source: state.proto
// DO NOT EDIT!

package mesosproto

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io2 "io"
import code_google_com_p_gogoprotobuf_proto4 "code.google.com/p/gogoprotobuf/proto"

import fmt6 "fmt"
import strings4 "strings"
import reflect4 "reflect"

import fmt7 "fmt"
import strings5 "strings"
import code_google_com_p_gogoprotobuf_proto5 "code.google.com/p/gogoprotobuf/proto"
import sort2 "sort"
import strconv2 "strconv"
import reflect5 "reflect"

import fmt8 "fmt"
import bytes2 "bytes"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Operation_Type int32

const (
	Operation_SNAPSHOT Operation_Type = 1
	Operation_EXPUNGE  Operation_Type = 2
)

var Operation_Type_name = map[int32]string{
	1: "SNAPSHOT",
	2: "EXPUNGE",
}
var Operation_Type_value = map[string]int32{
	"SNAPSHOT": 1,
	"EXPUNGE":  2,
}

func (x Operation_Type) Enum() *Operation_Type {
	p := new(Operation_Type)
	*p = x
	return p
}
func (x Operation_Type) String() string {
	return proto.EnumName(Operation_Type_name, int32(x))
}
func (x *Operation_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Operation_Type_value, data, "Operation_Type")
	if err != nil {
		return err
	}
	*x = Operation_Type(value)
	return nil
}

// Describes a state entry, a versioned (via a UUID) key/value pair.
type Entry struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Uuid             []byte  `protobuf:"bytes,2,req,name=uuid" json:"uuid,omitempty"`
	Value            []byte  `protobuf:"bytes,3,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Entry) Reset()      { *m = Entry{} }
func (*Entry) ProtoMessage() {}

func (m *Entry) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Entry) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *Entry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Describes an operation used in the log storage implementation.
type Operation struct {
	Type             *Operation_Type     `protobuf:"varint,1,req,name=type,enum=mesosproto.Operation_Type" json:"type,omitempty"`
	Snapshot         *Operation_Snapshot `protobuf:"bytes,2,opt,name=snapshot" json:"snapshot,omitempty"`
	Expunge          *Operation_Expunge  `protobuf:"bytes,3,opt,name=expunge" json:"expunge,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Operation) Reset()      { *m = Operation{} }
func (*Operation) ProtoMessage() {}

func (m *Operation) GetType() Operation_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Operation_SNAPSHOT
}

func (m *Operation) GetSnapshot() *Operation_Snapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func (m *Operation) GetExpunge() *Operation_Expunge {
	if m != nil {
		return m.Expunge
	}
	return nil
}

// Describes a "snapshot" operation.
type Operation_Snapshot struct {
	Entry            *Entry `protobuf:"bytes,1,req,name=entry" json:"entry,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Operation_Snapshot) Reset()      { *m = Operation_Snapshot{} }
func (*Operation_Snapshot) ProtoMessage() {}

func (m *Operation_Snapshot) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

// Describes an "expunge" operation.
type Operation_Expunge struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Operation_Expunge) Reset()      { *m = Operation_Expunge{} }
func (*Operation_Expunge) ProtoMessage() {}

func (m *Operation_Expunge) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("mesosproto.Operation_Type", Operation_Type_name, Operation_Type_value)
}
func (m *Entry) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io2.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto4.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io2.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io2.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Name = &s
			index = postIndex
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto4.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io2.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io2.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid, data[index:postIndex]...)
			index = postIndex
		case 3:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto4.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io2.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io2.ErrUnexpectedEOF
			}
			m.Value = append(m.Value, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto4.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io2.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Operation) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io2.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto4.ErrWrongType
			}
			var v Operation_Type
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io2.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (Operation_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Type = &v
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto4.ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io2.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io2.ErrUnexpectedEOF
			}
			if m.Snapshot == nil {
				m.Snapshot = &Operation_Snapshot{}
			}
			if err := m.Snapshot.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 3:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto4.ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io2.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io2.ErrUnexpectedEOF
			}
			if m.Expunge == nil {
				m.Expunge = &Operation_Expunge{}
			}
			if err := m.Expunge.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto4.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io2.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Operation_Snapshot) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io2.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto4.ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io2.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io2.ErrUnexpectedEOF
			}
			if m.Entry == nil {
				m.Entry = &Entry{}
			}
			if err := m.Entry.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto4.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io2.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Operation_Expunge) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io2.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto4.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io2.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io2.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Name = &s
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto4.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io2.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *Entry) String() string {
	if this == nil {
		return "nil"
	}
	s := strings4.Join([]string{`&Entry{`,
		`Name:` + valueToStringState(this.Name) + `,`,
		`Uuid:` + valueToStringState(this.Uuid) + `,`,
		`Value:` + valueToStringState(this.Value) + `,`,
		`XXX_unrecognized:` + fmt6.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Operation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings4.Join([]string{`&Operation{`,
		`Type:` + valueToStringState(this.Type) + `,`,
		`Snapshot:` + strings4.Replace(fmt6.Sprintf("%v", this.Snapshot), "Operation_Snapshot", "Operation_Snapshot", 1) + `,`,
		`Expunge:` + strings4.Replace(fmt6.Sprintf("%v", this.Expunge), "Operation_Expunge", "Operation_Expunge", 1) + `,`,
		`XXX_unrecognized:` + fmt6.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Operation_Snapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings4.Join([]string{`&Operation_Snapshot{`,
		`Entry:` + strings4.Replace(fmt6.Sprintf("%v", this.Entry), "Entry", "Entry", 1) + `,`,
		`XXX_unrecognized:` + fmt6.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Operation_Expunge) String() string {
	if this == nil {
		return "nil"
	}
	s := strings4.Join([]string{`&Operation_Expunge{`,
		`Name:` + valueToStringState(this.Name) + `,`,
		`XXX_unrecognized:` + fmt6.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringState(v interface{}) string {
	rv := reflect4.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect4.Indirect(rv).Interface()
	return fmt6.Sprintf("*%v", pv)
}
func (m *Entry) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovState(uint64(l))
	}
	if m.Uuid != nil {
		l = len(m.Uuid)
		n += 1 + l + sovState(uint64(l))
	}
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovState(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Operation) Size() (n int) {
	var l int
	_ = l
	if m.Type != nil {
		n += 1 + sovState(uint64(*m.Type))
	}
	if m.Snapshot != nil {
		l = m.Snapshot.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Expunge != nil {
		l = m.Expunge.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Operation_Snapshot) Size() (n int) {
	var l int
	_ = l
	if m.Entry != nil {
		l = m.Entry.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Operation_Expunge) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovState(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovState(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedEntry(r randyState, easy bool) *Entry {
	this := &Entry{}
	v1 := randStringState(r)
	this.Name = &v1
	v2 := r.Intn(100)
	this.Uuid = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Uuid[i] = byte(r.Intn(256))
	}
	v3 := r.Intn(100)
	this.Value = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.Value[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedState(r, 4)
	}
	return this
}

func NewPopulatedOperation(r randyState, easy bool) *Operation {
	this := &Operation{}
	v4 := Operation_Type([]int32{1, 2}[r.Intn(2)])
	this.Type = &v4
	if r.Intn(10) != 0 {
		this.Snapshot = NewPopulatedOperation_Snapshot(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Expunge = NewPopulatedOperation_Expunge(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedState(r, 4)
	}
	return this
}

func NewPopulatedOperation_Snapshot(r randyState, easy bool) *Operation_Snapshot {
	this := &Operation_Snapshot{}
	this.Entry = NewPopulatedEntry(r, easy)
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedState(r, 2)
	}
	return this
}

func NewPopulatedOperation_Expunge(r randyState, easy bool) *Operation_Expunge {
	this := &Operation_Expunge{}
	v5 := randStringState(r)
	this.Name = &v5
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedState(r, 2)
	}
	return this
}

type randyState interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneState(r randyState) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringState(r randyState) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneState(r)
	}
	return string(tmps)
}
func randUnrecognizedState(r randyState, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldState(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldState(data []byte, r randyState, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateState(data, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		data = encodeVarintPopulateState(data, uint64(v7))
	case 1:
		data = encodeVarintPopulateState(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateState(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateState(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateState(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateState(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *Entry) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Entry) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintState(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.Uuid != nil {
		data[i] = 0x12
		i++
		i = encodeVarintState(data, i, uint64(len(m.Uuid)))
		i += copy(data[i:], m.Uuid)
	}
	if m.Value != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintState(data, i, uint64(len(m.Value)))
		i += copy(data[i:], m.Value)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Operation) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Operation) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != nil {
		data[i] = 0x8
		i++
		i = encodeVarintState(data, i, uint64(*m.Type))
	}
	if m.Snapshot != nil {
		data[i] = 0x12
		i++
		i = encodeVarintState(data, i, uint64(m.Snapshot.Size()))
		n1, err := m.Snapshot.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Expunge != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintState(data, i, uint64(m.Expunge.Size()))
		n2, err := m.Expunge.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Operation_Snapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Operation_Snapshot) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Entry != nil {
		data[i] = 0xa
		i++
		i = encodeVarintState(data, i, uint64(m.Entry.Size()))
		n3, err := m.Entry.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Operation_Expunge) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Operation_Expunge) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintState(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64State(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32State(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintState(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *Entry) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings5.Join([]string{`&mesosproto.Entry{` + `Name:` + valueToGoStringState(this.Name, "string"), `Uuid:` + valueToGoStringState(this.Uuid, "byte"), `Value:` + valueToGoStringState(this.Value, "byte"), `XXX_unrecognized:` + fmt7.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Operation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings5.Join([]string{`&mesosproto.Operation{` + `Type:` + valueToGoStringState(this.Type, "mesosproto.Operation_Type"), `Snapshot:` + fmt7.Sprintf("%#v", this.Snapshot), `Expunge:` + fmt7.Sprintf("%#v", this.Expunge), `XXX_unrecognized:` + fmt7.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Operation_Snapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings5.Join([]string{`&mesosproto.Operation_Snapshot{` + `Entry:` + fmt7.Sprintf("%#v", this.Entry), `XXX_unrecognized:` + fmt7.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Operation_Expunge) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings5.Join([]string{`&mesosproto.Operation_Expunge{` + `Name:` + valueToGoStringState(this.Name, "string"), `XXX_unrecognized:` + fmt7.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringState(v interface{}, typ string) string {
	rv := reflect5.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect5.Indirect(rv).Interface()
	return fmt7.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringState(e map[int32]code_google_com_p_gogoprotobuf_proto5.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort2.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv2.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings5.Join(ss, ",") + "}"
	return s
}
func (this *Entry) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt8.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Entry)
	if !ok {
		return fmt8.Errorf("that is not of type *Entry")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt8.Errorf("that is type *Entry but is nil && this != nil")
	} else if this == nil {
		return fmt8.Errorf("that is type *Entrybut is not nil && this == nil")
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return fmt8.Errorf("Name this(%v) Not Equal that(%v)", *this.Name, *that1.Name)
		}
	} else if this.Name != nil {
		return fmt8.Errorf("this.Name == nil && that.Name != nil")
	} else if that1.Name != nil {
		return fmt8.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if !bytes2.Equal(this.Uuid, that1.Uuid) {
		return fmt8.Errorf("Uuid this(%v) Not Equal that(%v)", this.Uuid, that1.Uuid)
	}
	if !bytes2.Equal(this.Value, that1.Value) {
		return fmt8.Errorf("Value this(%v) Not Equal that(%v)", this.Value, that1.Value)
	}
	if !bytes2.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt8.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Entry) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Entry)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return false
		}
	} else if this.Name != nil {
		return false
	} else if that1.Name != nil {
		return false
	}
	if !bytes2.Equal(this.Uuid, that1.Uuid) {
		return false
	}
	if !bytes2.Equal(this.Value, that1.Value) {
		return false
	}
	if !bytes2.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Operation) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt8.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Operation)
	if !ok {
		return fmt8.Errorf("that is not of type *Operation")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt8.Errorf("that is type *Operation but is nil && this != nil")
	} else if this == nil {
		return fmt8.Errorf("that is type *Operationbut is not nil && this == nil")
	}
	if this.Type != nil && that1.Type != nil {
		if *this.Type != *that1.Type {
			return fmt8.Errorf("Type this(%v) Not Equal that(%v)", *this.Type, *that1.Type)
		}
	} else if this.Type != nil {
		return fmt8.Errorf("this.Type == nil && that.Type != nil")
	} else if that1.Type != nil {
		return fmt8.Errorf("Type this(%v) Not Equal that(%v)", this.Type, that1.Type)
	}
	if !this.Snapshot.Equal(that1.Snapshot) {
		return fmt8.Errorf("Snapshot this(%v) Not Equal that(%v)", this.Snapshot, that1.Snapshot)
	}
	if !this.Expunge.Equal(that1.Expunge) {
		return fmt8.Errorf("Expunge this(%v) Not Equal that(%v)", this.Expunge, that1.Expunge)
	}
	if !bytes2.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt8.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Operation) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Operation)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Type != nil && that1.Type != nil {
		if *this.Type != *that1.Type {
			return false
		}
	} else if this.Type != nil {
		return false
	} else if that1.Type != nil {
		return false
	}
	if !this.Snapshot.Equal(that1.Snapshot) {
		return false
	}
	if !this.Expunge.Equal(that1.Expunge) {
		return false
	}
	if !bytes2.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Operation_Snapshot) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt8.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Operation_Snapshot)
	if !ok {
		return fmt8.Errorf("that is not of type *Operation_Snapshot")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt8.Errorf("that is type *Operation_Snapshot but is nil && this != nil")
	} else if this == nil {
		return fmt8.Errorf("that is type *Operation_Snapshotbut is not nil && this == nil")
	}
	if !this.Entry.Equal(that1.Entry) {
		return fmt8.Errorf("Entry this(%v) Not Equal that(%v)", this.Entry, that1.Entry)
	}
	if !bytes2.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt8.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Operation_Snapshot) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Operation_Snapshot)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Entry.Equal(that1.Entry) {
		return false
	}
	if !bytes2.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Operation_Expunge) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt8.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Operation_Expunge)
	if !ok {
		return fmt8.Errorf("that is not of type *Operation_Expunge")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt8.Errorf("that is type *Operation_Expunge but is nil && this != nil")
	} else if this == nil {
		return fmt8.Errorf("that is type *Operation_Expungebut is not nil && this == nil")
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return fmt8.Errorf("Name this(%v) Not Equal that(%v)", *this.Name, *that1.Name)
		}
	} else if this.Name != nil {
		return fmt8.Errorf("this.Name == nil && that.Name != nil")
	} else if that1.Name != nil {
		return fmt8.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if !bytes2.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt8.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Operation_Expunge) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Operation_Expunge)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return false
		}
	} else if this.Name != nil {
		return false
	} else if that1.Name != nil {
		return false
	}
	if !bytes2.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
