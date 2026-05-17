// Code generated manually; source: app/auth/config.proto

package auth

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuthUrl       string                 `protobuf:"bytes,1,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	NodeToken     string                 `protobuf:"bytes,2,opt,name=node_token,json=nodeToken,proto3" json:"node_token,omitempty"`
	NodeId        string                 `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Protocol      string                 `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Timeout       int64                  `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_auth_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_auth_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Config) Descriptor() ([]byte, []int) {
	return file_app_auth_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *Config) GetNodeToken() string {
	if x != nil {
		return x.NodeToken
	}
	return ""
}

func (x *Config) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Config) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Config) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

var File_app_auth_config_proto protoreflect.FileDescriptor

const file_app_auth_config_proto_rawDesc = "" +
	"\x0a\x15\x61\x70\x70\x2f\x61\x75\x74\x68\x2f\x63\x6f\x6e\x66\x69" +
	"\x67\x2e\x70\x72\x6f\x74\x6f\x12\x0d\x78\x72\x61\x79\x2e\x61\x70" +
	"\x70\x2e\x61\x75\x74\x68\x22\x7e\x0a\x06\x43\x6f\x6e\x66\x69" +
	"\x67\x12\x19\x0a\x08\x61\x75\x74\x68\x5f\x75\x72\x6c\x18\x01" +
	"\x20\x01\x28\x09\x52\x07\x61\x75\x74\x68\x55\x72\x6c\x12\x1d" +
	"\x0a\x0a\x6e\x6f\x64\x65\x5f\x74\x6f\x6b\x65\x6e\x18\x02\x20\x01" +
	"\x28\x09\x52\x09\x6e\x6f\x64\x65\x54\x6f\x6b\x65\x6e\x12\x17" +
	"\x0a\x07\x6e\x6f\x64\x65\x5f\x69\x64\x18\x03\x20\x01\x28\x09" +
	"\x52\x06\x6e\x6f\x64\x65\x49\x64\x12\x10\x0a\x08\x70\x72\x6f" +
	"\x74\x6f\x63\x6f\x6c\x18\x04\x20\x01\x28\x09\x12\x0f\x0a\x07" +
	"\x74\x69\x6d\x65\x6f\x75\x74\x18\x05\x20\x01\x28\x03\x42\x24" +
	"\x5a\x22\x67\x69\x74\x68\x75\x62\x2e\x63\x6f\x6d\x2f\x78\x74" +
	"\x6c\x73\x2f\x78\x72\x61\x79\x2d\x63\x6f\x72\x65\x2f\x61\x70" +
	"\x70\x2f\x61\x75\x74\x68\x62\x06\x70\x72\x6f\x74\x6f\x33"

var (
	file_app_auth_config_proto_rawDescOnce sync.Once
	file_app_auth_config_proto_rawDescData []byte
)

func file_app_auth_config_proto_rawDescGZIP() []byte {
	file_app_auth_config_proto_rawDescOnce.Do(func() {
		file_app_auth_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_auth_config_proto_rawDesc), len(file_app_auth_config_proto_rawDesc)))
	})
	return file_app_auth_config_proto_rawDescData
}

var file_app_auth_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_auth_config_proto_goTypes = []any{
	(*Config)(nil), // 0: xray.app.auth.Config
}
var file_app_auth_config_proto_depIdxs = []int32{}

func init() { file_app_auth_config_proto_init() }
func file_app_auth_config_proto_init() {
	if File_app_auth_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_auth_config_proto_rawDesc), len(file_app_auth_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_auth_config_proto_goTypes,
		DependencyIndexes: file_app_auth_config_proto_depIdxs,
		MessageInfos:      file_app_auth_config_proto_msgTypes,
	}.Build()
	File_app_auth_config_proto = out.File
	file_app_auth_config_proto_goTypes = nil
	file_app_auth_config_proto_depIdxs = nil
}
