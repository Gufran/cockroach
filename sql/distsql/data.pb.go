// Code generated by protoc-gen-gogo.
// source: cockroach/sql/distsql/data.proto
// DO NOT EDIT!

package distsql

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/roachpb"
import cockroach_sql_sqlbase1 "github.com/cockroachdb/cockroach/sql/sqlbase"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Expression struct {
	// TODO(radu): TBD how this will be used
	Version string `protobuf:"bytes,1,opt,name=version" json:"version"`
	// SQL expressions are passed as a string, with ValArgs ($0, $1, ..) used for
	// "input" variables.
	Expr string `protobuf:"bytes,2,opt,name=expr" json:"expr"`
}

func (m *Expression) Reset()                    { *m = Expression{} }
func (m *Expression) String() string            { return proto.CompactTextString(m) }
func (*Expression) ProtoMessage()               {}
func (*Expression) Descriptor() ([]byte, []int) { return fileDescriptorData, []int{0} }

type TableReaderSpan struct {
	// TODO(radu): the dist_sql APIs should be agnostic to how we map tables to
	// KVs. The span should be described as starting and ending lists of values
	// for a prefix of the index columns, along with inclusive/exclusive flags.
	Span cockroach_roachpb1.Span `protobuf:"bytes,1,opt,name=span" json:"span"`
}

func (m *TableReaderSpan) Reset()                    { *m = TableReaderSpan{} }
func (m *TableReaderSpan) String() string            { return proto.CompactTextString(m) }
func (*TableReaderSpan) ProtoMessage()               {}
func (*TableReaderSpan) Descriptor() ([]byte, []int) { return fileDescriptorData, []int{1} }

// TableReaderSpec is the specification for a table reader. A table reader
// performs KV operations to retrieve rows for a table and outputs the desired
// columns of the rows that pass a filter expression.
type TableReaderSpec struct {
	Table cockroach_sql_sqlbase1.TableDescriptor `protobuf:"bytes,1,opt,name=table" json:"table"`
	// If 0, we use the primary index. If non-zero, we use the index_idx-th index,
	// i.e. table.indexes[index_idx-1]
	IndexIdx uint32            `protobuf:"varint,2,opt,name=index_idx,json=indexIdx" json:"index_idx"`
	Reverse  bool              `protobuf:"varint,3,opt,name=reverse" json:"reverse"`
	Spans    []TableReaderSpan `protobuf:"bytes,4,rep,name=spans" json:"spans"`
	// The filter expression references the columns in the table (table.columns)
	// via $0, $1, etc. If a secondary index is used, the columns that are not
	// available as part of the index cannot be referenced.
	Filter Expression `protobuf:"bytes,5,opt,name=filter" json:"filter"`
	// The table reader will only produce values for these columns, referenced by
	// their indices in table.columns.
	OutputColumns []uint32 `protobuf:"varint,6,rep,name=output_columns,json=outputColumns" json:"output_columns,omitempty"`
}

func (m *TableReaderSpec) Reset()                    { *m = TableReaderSpec{} }
func (m *TableReaderSpec) String() string            { return proto.CompactTextString(m) }
func (*TableReaderSpec) ProtoMessage()               {}
func (*TableReaderSpec) Descriptor() ([]byte, []int) { return fileDescriptorData, []int{2} }

// FlowSpec describes a "flow" which is a subgraph of a distributed SQL
// computation consisting of processors and streams.
type FlowSpec struct {
	// TODO(radu): for now a flow is just a table reader.
	Reader *TableReaderSpec `protobuf:"bytes,2,opt,name=reader" json:"reader,omitempty"`
}

func (m *FlowSpec) Reset()                    { *m = FlowSpec{} }
func (m *FlowSpec) String() string            { return proto.CompactTextString(m) }
func (*FlowSpec) ProtoMessage()               {}
func (*FlowSpec) Descriptor() ([]byte, []int) { return fileDescriptorData, []int{3} }

func init() {
	proto.RegisterType((*Expression)(nil), "cockroach.sql.distsql.Expression")
	proto.RegisterType((*TableReaderSpan)(nil), "cockroach.sql.distsql.TableReaderSpan")
	proto.RegisterType((*TableReaderSpec)(nil), "cockroach.sql.distsql.TableReaderSpec")
	proto.RegisterType((*FlowSpec)(nil), "cockroach.sql.distsql.FlowSpec")
}
func (m *Expression) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Expression) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintData(data, i, uint64(len(m.Version)))
	i += copy(data[i:], m.Version)
	data[i] = 0x12
	i++
	i = encodeVarintData(data, i, uint64(len(m.Expr)))
	i += copy(data[i:], m.Expr)
	return i, nil
}

func (m *TableReaderSpan) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TableReaderSpan) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintData(data, i, uint64(m.Span.Size()))
	n1, err := m.Span.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *TableReaderSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TableReaderSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintData(data, i, uint64(m.Table.Size()))
	n2, err := m.Table.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x10
	i++
	i = encodeVarintData(data, i, uint64(m.IndexIdx))
	data[i] = 0x18
	i++
	if m.Reverse {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	if len(m.Spans) > 0 {
		for _, msg := range m.Spans {
			data[i] = 0x22
			i++
			i = encodeVarintData(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x2a
	i++
	i = encodeVarintData(data, i, uint64(m.Filter.Size()))
	n3, err := m.Filter.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.OutputColumns) > 0 {
		for _, num := range m.OutputColumns {
			data[i] = 0x30
			i++
			i = encodeVarintData(data, i, uint64(num))
		}
	}
	return i, nil
}

func (m *FlowSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *FlowSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Reader != nil {
		data[i] = 0x12
		i++
		i = encodeVarintData(data, i, uint64(m.Reader.Size()))
		n4, err := m.Reader.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeFixed64Data(data []byte, offset int, v uint64) int {
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
func encodeFixed32Data(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintData(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Expression) Size() (n int) {
	var l int
	_ = l
	l = len(m.Version)
	n += 1 + l + sovData(uint64(l))
	l = len(m.Expr)
	n += 1 + l + sovData(uint64(l))
	return n
}

func (m *TableReaderSpan) Size() (n int) {
	var l int
	_ = l
	l = m.Span.Size()
	n += 1 + l + sovData(uint64(l))
	return n
}

func (m *TableReaderSpec) Size() (n int) {
	var l int
	_ = l
	l = m.Table.Size()
	n += 1 + l + sovData(uint64(l))
	n += 1 + sovData(uint64(m.IndexIdx))
	n += 2
	if len(m.Spans) > 0 {
		for _, e := range m.Spans {
			l = e.Size()
			n += 1 + l + sovData(uint64(l))
		}
	}
	l = m.Filter.Size()
	n += 1 + l + sovData(uint64(l))
	if len(m.OutputColumns) > 0 {
		for _, e := range m.OutputColumns {
			n += 1 + sovData(uint64(e))
		}
	}
	return n
}

func (m *FlowSpec) Size() (n int) {
	var l int
	_ = l
	if m.Reader != nil {
		l = m.Reader.Size()
		n += 1 + l + sovData(uint64(l))
	}
	return n
}

func sovData(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozData(x uint64) (n int) {
	return sovData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Expression) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Expression: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Expression: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TableReaderSpan) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TableReaderSpan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableReaderSpan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Span", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Span.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TableReaderSpec) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TableReaderSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableReaderSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Table", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Table.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexIdx", wireType)
			}
			m.IndexIdx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.IndexIdx |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reverse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Reverse = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spans = append(m.Spans, TableReaderSpan{})
			if err := m.Spans[len(m.Spans)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Filter.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputColumns", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OutputColumns = append(m.OutputColumns, v)
		default:
			iNdEx = preIndex
			skippy, err := skipData(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FlowSpec) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FlowSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlowSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Reader == nil {
				m.Reader = &TableReaderSpec{}
			}
			if err := m.Reader.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipData(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowData
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowData
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowData
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthData
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowData
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipData(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthData = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowData   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorData = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4e, 0xe3, 0x30,
	0x18, 0xec, 0x4f, 0xfa, 0xe7, 0x55, 0x77, 0x25, 0x6b, 0x57, 0x1b, 0x55, 0x28, 0xb4, 0x91, 0xa8,
	0x38, 0x39, 0x82, 0x07, 0x00, 0xa9, 0x94, 0x4a, 0x70, 0x0c, 0x1c, 0x10, 0x97, 0x2a, 0x75, 0x4c,
	0x89, 0x08, 0x75, 0xb0, 0x1d, 0xc8, 0x63, 0xf0, 0x50, 0x1c, 0x7a, 0xe4, 0xc8, 0x09, 0xf1, 0xf3,
	0x22, 0xd8, 0x8e, 0x43, 0x7f, 0x04, 0x12, 0x87, 0x2f, 0xb1, 0xc6, 0x33, 0xa3, 0x99, 0xcf, 0xa0,
	0x8b, 0x29, 0xbe, 0x62, 0x34, 0xc0, 0x97, 0x1e, 0xbf, 0x89, 0xbd, 0x30, 0xe2, 0x42, 0xff, 0x03,
	0x11, 0xa0, 0x84, 0x51, 0x41, 0xe1, 0xbf, 0x4f, 0x06, 0x92, 0x37, 0xc8, 0x30, 0x3a, 0x1b, 0x0b,
	0xa1, 0xfe, 0x26, 0x93, 0x25, 0x51, 0xa7, 0xbf, 0x6a, 0x2b, 0x67, 0x12, 0x70, 0xe2, 0x71, 0xc1,
	0x52, 0x2c, 0x52, 0x46, 0x42, 0xc3, 0xfb, 0x3b, 0xa5, 0x53, 0xaa, 0x8f, 0x9e, 0x3a, 0xe5, 0xa8,
	0x3b, 0x02, 0xe0, 0x30, 0x4b, 0x18, 0xe1, 0x3c, 0xa2, 0x33, 0xe8, 0x80, 0xc6, 0x2d, 0x61, 0xea,
	0x68, 0x97, 0xbb, 0xe5, 0xed, 0xd6, 0xc0, 0x9a, 0x3f, 0x6f, 0x96, 0xfc, 0x02, 0x84, 0x36, 0xb0,
	0x88, 0x64, 0xdb, 0x95, 0xa5, 0x4b, 0x8d, 0xb8, 0x43, 0xf0, 0xe7, 0x34, 0x98, 0xc4, 0xc4, 0x27,
	0x41, 0x48, 0xd8, 0x49, 0x12, 0xcc, 0xe0, 0x0e, 0xb0, 0xb8, 0xfc, 0x6b, 0xa7, 0x5f, 0xbb, 0xff,
	0xd1, 0xa2, 0x9c, 0x69, 0x81, 0x14, 0xad, 0x70, 0x51, 0x54, 0xf7, 0xa1, 0xb2, 0x66, 0x43, 0x30,
	0x1c, 0x80, 0x9a, 0x50, 0x90, 0xf1, 0xe9, 0xa3, 0xd5, 0x25, 0x99, 0xbe, 0x48, 0xcb, 0x86, 0x84,
	0x63, 0x16, 0x25, 0x82, 0x32, 0x63, 0x9b, 0x4b, 0x61, 0x0f, 0xb4, 0xa2, 0x59, 0x48, 0xb2, 0x71,
	0x14, 0x66, 0x3a, 0x7c, 0xdb, 0xdc, 0x37, 0x35, 0x7c, 0x14, 0x66, 0xaa, 0x3a, 0x23, 0xaa, 0x27,
	0xb1, 0xab, 0x92, 0xd0, 0x2c, 0xaa, 0x1b, 0x50, 0xc5, 0x50, 0x11, 0xb9, 0x6d, 0x75, 0xab, 0x5f,
	0xc4, 0x30, 0x6f, 0x85, 0xd6, 0x96, 0x50, 0xc4, 0xd0, 0x52, 0xb8, 0x0f, 0xea, 0x17, 0x51, 0x2c,
	0x08, 0xb3, 0x6b, 0xba, 0x4b, 0xef, 0x1b, 0x93, 0xc5, 0x8b, 0x18, 0xbd, 0x91, 0xc1, 0x2d, 0xf0,
	0x9b, 0xa6, 0x22, 0x49, 0xc5, 0x18, 0xd3, 0x38, 0xbd, 0x96, 0x69, 0xea, 0x32, 0x4d, 0xdb, 0x6f,
	0xe7, 0xe8, 0x41, 0x0e, 0xba, 0xc7, 0xa0, 0x39, 0x8a, 0xe9, 0x9d, 0x5e, 0xdf, 0x1e, 0xa8, 0x33,
	0x1d, 0x47, 0xf7, 0xfe, 0x61, 0x70, 0x82, 0x7d, 0xa3, 0x1a, 0xf4, 0xe6, 0xaf, 0x4e, 0x69, 0xfe,
	0xe6, 0x94, 0x1f, 0xe5, 0x3c, 0xc9, 0x79, 0x91, 0x73, 0xff, 0xee, 0x94, 0xce, 0x1b, 0x46, 0x7a,
	0x56, 0xf9, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x4e, 0x67, 0x51, 0xdb, 0x02, 0x00, 0x00,
}