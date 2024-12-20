// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: pkg/proto/nbcontract/v1/httpproxyconfig.proto

package nbcontractv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HTTPProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Proxy endpoint/URL using HTTP through which the traffic should be routed for making HTTP connections outside the cluster
	HttpProxy string `protobuf:"bytes,1,opt,name=http_proxy,json=httpProxy,proto3" json:"http_proxy,omitempty"`
	// Proxy endpoint/URL using HTTPS through which the traffic should be routed for making HTTPS connections outside the cluster
	HttpsProxy string `protobuf:"bytes,2,opt,name=https_proxy,json=httpsProxy,proto3" json:"https_proxy,omitempty"`
	// A list of entries including IPs, domains, or other network CIDRs that should not be accessed through the proxy
	NoProxyEntries []string `protobuf:"bytes,3,rep,name=no_proxy_entries,json=noProxyEntries,proto3" json:"no_proxy_entries,omitempty"`
	// Custom CA certificate that should be trusted by the proxy
	ProxyTrustedCa string `protobuf:"bytes,4,opt,name=proxy_trusted_ca,json=proxyTrustedCa,proto3" json:"proxy_trusted_ca,omitempty"`
}

func (x *HTTPProxyConfig) Reset() {
	*x = HTTPProxyConfig{}
	mi := &file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HTTPProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPProxyConfig) ProtoMessage() {}

func (x *HTTPProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPProxyConfig.ProtoReflect.Descriptor instead.
func (*HTTPProxyConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPProxyConfig) GetHttpProxy() string {
	if x != nil {
		return x.HttpProxy
	}
	return ""
}

func (x *HTTPProxyConfig) GetHttpsProxy() string {
	if x != nil {
		return x.HttpsProxy
	}
	return ""
}

func (x *HTTPProxyConfig) GetNoProxyEntries() []string {
	if x != nil {
		return x.NoProxyEntries
	}
	return nil
}

func (x *HTTPProxyConfig) GetProxyTrustedCa() string {
	if x != nil {
		return x.ProxyTrustedCa
	}
	return ""
}

var File_pkg_proto_nbcontract_v1_httpproxyconfig_proto protoreflect.FileDescriptor

var file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xa5,
	0x01, 0x0a, 0x0f, 0x48, 0x54, 0x54, 0x50, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x65, 0x64, 0x43, 0x61, 0x42, 0xc0, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x48, 0x74,
	0x74, 0x70, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6b, 0x65,
	0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4e,
	0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x4e,
	0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x4e,
	0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDescOnce sync.Once
	file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDescData = file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDesc
)

func file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDescGZIP() []byte {
	file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDescOnce.Do(func() {
		file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDescData)
	})
	return file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDescData
}

var file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_goTypes = []any{
	(*HTTPProxyConfig)(nil), // 0: nbcontract.v1.HTTPProxyConfig
}
var file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_init() }
func file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_init() {
	if File_pkg_proto_nbcontract_v1_httpproxyconfig_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_goTypes,
		DependencyIndexes: file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_depIdxs,
		MessageInfos:      file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_msgTypes,
	}.Build()
	File_pkg_proto_nbcontract_v1_httpproxyconfig_proto = out.File
	file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_rawDesc = nil
	file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_goTypes = nil
	file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_depIdxs = nil
}
