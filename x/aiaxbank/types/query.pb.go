// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aiaxbank/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryDenomRepresentationRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryDenomRepresentationRequest) Reset()         { *m = QueryDenomRepresentationRequest{} }
func (m *QueryDenomRepresentationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomRepresentationRequest) ProtoMessage()    {}
func (*QueryDenomRepresentationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa428b2ca4ab705f, []int{0}
}
func (m *QueryDenomRepresentationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomRepresentationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomRepresentationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomRepresentationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomRepresentationRequest.Merge(m, src)
}
func (m *QueryDenomRepresentationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomRepresentationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomRepresentationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomRepresentationRequest proto.InternalMessageInfo

type QueryDenomRepresentationResponse struct {
	ExternalAddress string `protobuf:"bytes,1,opt,name=external_address,json=externalAddress,proto3" json:"external_address,omitempty"`
	InternalAddress string `protobuf:"bytes,2,opt,name=internal_address,json=internalAddress,proto3" json:"internal_address,omitempty"`
}

func (m *QueryDenomRepresentationResponse) Reset()         { *m = QueryDenomRepresentationResponse{} }
func (m *QueryDenomRepresentationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomRepresentationResponse) ProtoMessage()    {}
func (*QueryDenomRepresentationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa428b2ca4ab705f, []int{1}
}
func (m *QueryDenomRepresentationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomRepresentationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomRepresentationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomRepresentationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomRepresentationResponse.Merge(m, src)
}
func (m *QueryDenomRepresentationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomRepresentationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomRepresentationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomRepresentationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryDenomRepresentationRequest)(nil), "aiaxbank.v1.QueryDenomRepresentationRequest")
	proto.RegisterType((*QueryDenomRepresentationResponse)(nil), "aiaxbank.v1.QueryDenomRepresentationResponse")
}

func init() { proto.RegisterFile("aiaxbank/v1/query.proto", fileDescriptor_aa428b2ca4ab705f) }

var fileDescriptor_aa428b2ca4ab705f = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcc, 0x4c, 0xac,
	0x48, 0x4a, 0xcc, 0xcb, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x49, 0xe8, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7,
	0x83, 0xc5, 0xf5, 0x41, 0x2c, 0x88, 0x12, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd,
	0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88,
	0xac, 0x92, 0x23, 0x97, 0x7c, 0x20, 0xc8, 0x3c, 0x97, 0xd4, 0xbc, 0xfc, 0xdc, 0xa0, 0xd4, 0x82,
	0xa2, 0xd4, 0xe2, 0xd4, 0x3c, 0x88, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11,
	0x2e, 0xd6, 0x14, 0x90, 0xac, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x63, 0xc5, 0xd1,
	0xb1, 0x40, 0x9e, 0xe1, 0xc5, 0x02, 0x79, 0x06, 0xa5, 0x26, 0x46, 0x2e, 0x05, 0xdc, 0x66, 0x14,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x69, 0x72, 0x09, 0xa4, 0x56, 0x94, 0xa4, 0x16, 0xe5, 0x25,
	0xe6, 0xc4, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x43, 0xcd, 0xe3, 0x87, 0x89, 0x3b, 0x42,
	0x84, 0x41, 0x4a, 0x33, 0xf3, 0xd0, 0x94, 0x32, 0x41, 0x94, 0xc2, 0xc4, 0xa1, 0x4a, 0x11, 0x8e,
	0x30, 0x5a, 0xca, 0xc8, 0xc5, 0x0a, 0x76, 0x84, 0xd0, 0x6c, 0x46, 0x2e, 0x61, 0x2c, 0x2e, 0x11,
	0xd2, 0xd1, 0x43, 0x0a, 0x2b, 0x3d, 0x02, 0x9e, 0x96, 0xd2, 0x25, 0x52, 0x35, 0xc4, 0x7b, 0x4a,
	0xea, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x52, 0x14, 0x92, 0xd7, 0x47, 0x8e, 0x29, 0x70, 0x48, 0xc5,
	0x17, 0xa5, 0x16, 0x14, 0xe9, 0x57, 0x83, 0xd9, 0xb5, 0x4e, 0x3e, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f,
	0x0b, 0x36, 0x44, 0x37, 0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x1b, 0xca, 0xc9, 0x4f, 0x49, 0xd5,
	0xaf, 0x40, 0x98, 0x5e, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x8e, 0x44, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc1, 0xdf, 0x8f, 0xc0, 0x20, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	DenomRepresentation(ctx context.Context, in *QueryDenomRepresentationRequest, opts ...grpc.CallOption) (*QueryDenomRepresentationResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DenomRepresentation(ctx context.Context, in *QueryDenomRepresentationRequest, opts ...grpc.CallOption) (*QueryDenomRepresentationResponse, error) {
	out := new(QueryDenomRepresentationResponse)
	err := c.cc.Invoke(ctx, "/aiaxbank.v1.Query/DenomRepresentation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	DenomRepresentation(context.Context, *QueryDenomRepresentationRequest) (*QueryDenomRepresentationResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DenomRepresentation(ctx context.Context, req *QueryDenomRepresentationRequest) (*QueryDenomRepresentationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomRepresentation not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DenomRepresentation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomRepresentationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomRepresentation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aiaxbank.v1.Query/DenomRepresentation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomRepresentation(ctx, req.(*QueryDenomRepresentationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aiaxbank.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DenomRepresentation",
			Handler:    _Query_DenomRepresentation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aiaxbank/v1/query.proto",
}

func (m *QueryDenomRepresentationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomRepresentationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomRepresentationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomRepresentationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomRepresentationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomRepresentationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InternalAddress) > 0 {
		i -= len(m.InternalAddress)
		copy(dAtA[i:], m.InternalAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.InternalAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ExternalAddress) > 0 {
		i -= len(m.ExternalAddress)
		copy(dAtA[i:], m.ExternalAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ExternalAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryDenomRepresentationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomRepresentationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExternalAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.InternalAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDenomRepresentationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomRepresentationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomRepresentationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDenomRepresentationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomRepresentationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomRepresentationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InternalAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
