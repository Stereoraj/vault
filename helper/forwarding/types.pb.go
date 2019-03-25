// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helper/forwarding/types.proto

package forwarding

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	// Not used right now but reserving in case it turns out that streaming
	// makes things more economical on the gRPC side
	//uint64 id = 1;
	Method               string                  `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Url                  *URL                    `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	HeaderEntries        map[string]*HeaderEntry `protobuf:"bytes,4,rep,name=header_entries,json=headerEntries,proto3" json:"header_entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte                  `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Host                 string                  `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	RemoteAddr           string                  `protobuf:"bytes,7,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	PeerCertificates     [][]byte                `protobuf:"bytes,8,rep,name=peer_certificates,json=peerCertificates,proto3" json:"peer_certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38697de88a2f47c, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetUrl() *URL {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Request) GetHeaderEntries() map[string]*HeaderEntry {
	if m != nil {
		return m.HeaderEntries
	}
	return nil
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Request) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *Request) GetPeerCertificates() [][]byte {
	if m != nil {
		return m.PeerCertificates
	}
	return nil
}

type URL struct {
	Scheme string `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Opaque string `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// This isn't needed now but might be in the future, so we'll skip the
	// number to keep the ordering in net/url
	//UserInfo user = 3;
	Host    string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Path    string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	RawPath string `protobuf:"bytes,6,opt,name=raw_path,json=rawPath,proto3" json:"raw_path,omitempty"`
	// This also isn't needed right now, but we'll reserve the number
	//bool force_query = 7;
	RawQuery             string   `protobuf:"bytes,8,opt,name=raw_query,json=rawQuery,proto3" json:"raw_query,omitempty"`
	Fragment             string   `protobuf:"bytes,9,opt,name=fragment,proto3" json:"fragment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38697de88a2f47c, []int{1}
}

func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (m *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(m, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *URL) GetOpaque() string {
	if m != nil {
		return m.Opaque
	}
	return ""
}

func (m *URL) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *URL) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *URL) GetRawPath() string {
	if m != nil {
		return m.RawPath
	}
	return ""
}

func (m *URL) GetRawQuery() string {
	if m != nil {
		return m.RawQuery
	}
	return ""
}

func (m *URL) GetFragment() string {
	if m != nil {
		return m.Fragment
	}
	return ""
}

type HeaderEntry struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderEntry) Reset()         { *m = HeaderEntry{} }
func (m *HeaderEntry) String() string { return proto.CompactTextString(m) }
func (*HeaderEntry) ProtoMessage()    {}
func (*HeaderEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38697de88a2f47c, []int{2}
}

func (m *HeaderEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderEntry.Unmarshal(m, b)
}
func (m *HeaderEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderEntry.Marshal(b, m, deterministic)
}
func (m *HeaderEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderEntry.Merge(m, src)
}
func (m *HeaderEntry) XXX_Size() int {
	return xxx_messageInfo_HeaderEntry.Size(m)
}
func (m *HeaderEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderEntry.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderEntry proto.InternalMessageInfo

func (m *HeaderEntry) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Response struct {
	// Not used right now but reserving in case it turns out that streaming
	// makes things more economical on the gRPC side
	//uint64 id = 1;
	StatusCode uint32 `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Body       []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// Added in 0.6.2 to ensure that the content-type is set appropriately, as
	// well as any other information
	HeaderEntries        map[string]*HeaderEntry `protobuf:"bytes,4,rep,name=header_entries,json=headerEntries,proto3" json:"header_entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastRemoteWal        uint64                  `protobuf:"varint,5,opt,name=last_remote_wal,json=lastRemoteWal,proto3" json:"last_remote_wal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38697de88a2f47c, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() uint32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Response) GetHeaderEntries() map[string]*HeaderEntry {
	if m != nil {
		return m.HeaderEntries
	}
	return nil
}

func (m *Response) GetLastRemoteWal() uint64 {
	if m != nil {
		return m.LastRemoteWal
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "forwarding.Request")
	proto.RegisterMapType((map[string]*HeaderEntry)(nil), "forwarding.Request.HeaderEntriesEntry")
	proto.RegisterType((*URL)(nil), "forwarding.URL")
	proto.RegisterType((*HeaderEntry)(nil), "forwarding.HeaderEntry")
	proto.RegisterType((*Response)(nil), "forwarding.Response")
	proto.RegisterMapType((map[string]*HeaderEntry)(nil), "forwarding.Response.HeaderEntriesEntry")
}

func init() { proto.RegisterFile("helper/forwarding/types.proto", fileDescriptor_e38697de88a2f47c) }

var fileDescriptor_e38697de88a2f47c = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0xb4, 0x49, 0x26, 0x0d, 0x2d, 0x7b, 0x80, 0xa5, 0x08, 0x61, 0x22, 0x51, 0x22,
	0x01, 0x8e, 0x14, 0x2e, 0x88, 0x1b, 0x54, 0x48, 0x1c, 0x0a, 0x82, 0x95, 0x2a, 0x04, 0x17, 0x6b,
	0xe3, 0x9d, 0x64, 0x2d, 0xec, 0xac, 0xb3, 0xbb, 0x6e, 0xe4, 0xdf, 0xe0, 0x4f, 0xf8, 0x27, 0x3e,
	0x04, 0xed, 0xda, 0x34, 0x46, 0x15, 0x12, 0x17, 0x4e, 0x99, 0xf7, 0xde, 0x64, 0x3c, 0x6f, 0x66,
	0x16, 0x1e, 0x48, 0xcc, 0x4b, 0xd4, 0xf3, 0x95, 0xd2, 0x3b, 0xae, 0x45, 0xb6, 0x59, 0xcf, 0x6d,
	0x5d, 0xa2, 0x89, 0x4b, 0xad, 0xac, 0x22, 0xb0, 0xe7, 0xa7, 0x3f, 0x7b, 0x30, 0x60, 0xb8, 0xad,
	0xd0, 0x58, 0x72, 0x07, 0x0e, 0x0b, 0xb4, 0x52, 0x09, 0xda, 0x8b, 0x82, 0xd9, 0x88, 0xb5, 0x88,
	0x3c, 0x82, 0xb0, 0xd2, 0x39, 0x0d, 0xa3, 0x60, 0x36, 0x5e, 0x1c, 0xc7, 0xfb, 0x7f, 0xc7, 0x97,
	0xec, 0x82, 0x39, 0x8d, 0xbc, 0x87, 0x5b, 0x12, 0xb9, 0x40, 0x9d, 0xe0, 0xc6, 0xea, 0x0c, 0x0d,
	0xed, 0x47, 0xe1, 0x6c, 0xbc, 0x38, 0xeb, 0x66, 0xb7, 0xdf, 0x89, 0xdf, 0xf9, 0xcc, 0xb7, 0x4d,
	0xa2, 0xfb, 0xa9, 0xd9, 0x44, 0x76, 0x39, 0x42, 0xa0, 0xbf, 0x54, 0xa2, 0xa6, 0x07, 0x51, 0x30,
	0x3b, 0x62, 0x3e, 0x76, 0x9c, 0x54, 0xc6, 0xd2, 0x43, 0xdf, 0x9b, 0x8f, 0xc9, 0x43, 0x18, 0x6b,
	0x2c, 0x94, 0xc5, 0x84, 0x0b, 0xa1, 0xe9, 0xc0, 0x4b, 0xd0, 0x50, 0xaf, 0x85, 0xd0, 0xe4, 0x29,
	0xdc, 0x2e, 0x11, 0x75, 0x92, 0xa2, 0xb6, 0xd9, 0x2a, 0x4b, 0xb9, 0x45, 0x43, 0x87, 0x51, 0x38,
	0x3b, 0x62, 0x27, 0x4e, 0x38, 0xef, 0xf0, 0xa7, 0x5f, 0x80, 0xdc, 0x6c, 0x8d, 0x9c, 0x40, 0xf8,
	0x0d, 0x6b, 0x1a, 0xf8, 0xda, 0x2e, 0x24, 0xcf, 0xe1, 0xe0, 0x8a, 0xe7, 0x15, 0xfa, 0x31, 0x8d,
	0x17, 0x77, 0xbb, 0x1e, 0xf7, 0x05, 0x6a, 0xd6, 0x64, 0xbd, 0xea, 0xbd, 0x0c, 0xa6, 0x3f, 0x02,
	0x08, 0x2f, 0xd9, 0x85, 0x1b, 0xb1, 0x49, 0x25, 0x16, 0xd8, 0xd6, 0x6b, 0x91, 0xe3, 0x55, 0xc9,
	0xb7, 0x6d, 0xcd, 0x11, 0x6b, 0xd1, 0xb5, 0xe9, 0x7e, 0xc7, 0x34, 0x81, 0x7e, 0xc9, 0xad, 0xf4,
	0xc3, 0x19, 0x31, 0x1f, 0x93, 0x7b, 0x30, 0xd4, 0x7c, 0x97, 0x78, 0xbe, 0x19, 0xd0, 0x40, 0xf3,
	0xdd, 0x47, 0x27, 0xdd, 0x87, 0x91, 0x93, 0xb6, 0x15, 0xea, 0x9a, 0x0e, 0xbd, 0xe6, 0x72, 0x3f,
	0x39, 0x4c, 0x4e, 0x61, 0xb8, 0xd2, 0x7c, 0x5d, 0xe0, 0xc6, 0xd2, 0x51, 0xa3, 0xfd, 0xc6, 0xd3,
	0xc7, 0x30, 0xee, 0xb8, 0x71, 0x2d, 0x7a, 0x3f, 0x86, 0x06, 0x51, 0xe8, 0x5a, 0x6c, 0xd0, 0xf4,
	0x7b, 0x0f, 0x86, 0x0c, 0x4d, 0xa9, 0x36, 0x06, 0xdd, 0x42, 0x8c, 0xe5, 0xb6, 0x32, 0x49, 0xaa,
	0x44, 0x63, 0x66, 0xc2, 0xa0, 0xa1, 0xce, 0x95, 0xc0, 0xeb, 0xcd, 0x86, 0x9d, 0xcd, 0x7e, 0xf8,
	0xcb, 0xf1, 0x3c, 0xf9, 0xf3, 0x78, 0x9a, 0x4f, 0xfc, 0xc3, 0xf5, 0x9c, 0xc1, 0x71, 0xce, 0x8d,
	0x4d, 0xda, 0xd3, 0xd8, 0xf1, 0xdc, 0xcf, 0xaa, 0xcf, 0x26, 0x8e, 0x66, 0x9e, 0xfd, 0xcc, 0xf3,
	0xff, 0xb8, 0xef, 0x37, 0xf1, 0xd7, 0x67, 0xeb, 0xcc, 0xca, 0x6a, 0x19, 0xa7, 0xaa, 0x98, 0x4b,
	0x6e, 0x64, 0x96, 0x2a, 0x5d, 0xce, 0xaf, 0x78, 0x95, 0xdb, 0xf9, 0x8d, 0xe7, 0xb9, 0x3c, 0xf4,
	0x2f, 0xf3, 0xc5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xbd, 0xb1, 0xfc, 0xba, 0x03, 0x00,
	0x00,
}
