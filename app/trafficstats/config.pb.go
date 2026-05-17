// Code generated manually; source: app/trafficstats/config.proto

package trafficstats

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
	Listen        string                 `protobuf:"bytes,1,opt,name=listen,proto3" json:"listen,omitempty"`
	Secret        string                 `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_trafficstats_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_trafficstats_config_proto_msgTypes[0]
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
	return file_app_trafficstats_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetListen() string {
	if x != nil {
		return x.Listen
	}
	return ""
}

func (x *Config) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

var File_app_trafficstats_config_proto protoreflect.FileDescriptor

const file_app_trafficstats_config_proto_rawDesc = "" +
	"\x0a\x1d\x61\x70\x70\x2f\x74\x72\x61\x66\x66\x69\x63\x73\x74\x61" +
	"\x74\x73\x2f\x63\x6f\x6e\x66\x69\x67\x2e\x70\x72\x6f\x74\x6f\x12" +
	"\x15\x78\x72\x61\x79\x2e\x61\x70\x70\x2e\x74\x72\x61\x66\x66\x69" +
	"\x63\x73\x74\x61\x74\x73\x22\x28\x0a\x06\x43\x6f\x6e\x66\x69" +
	"\x67\x12\x0e\x0a\x06\x6c\x69\x73\x74\x65\x6e\x18\x01\x20\x01" +
	"\x28\x09\x12\x0e\x0a\x06\x73\x65\x63\x72\x65\x74\x18\x02\x20\x01" +
	"\x28\x09\x42\x2c\x5a\x2a\x67\x69\x74\x68\x75\x62\x2e\x63\x6f" +
	"\x6d\x2f\x78\x74\x6c\x73\x2f\x78\x72\x61\x79\x2d\x63\x6f\x72" +
	"\x65\x2f\x61\x70\x70\x2f\x74\x72\x61\x66\x66\x69\x63\x73\x74" +
	"\x61\x74\x73\x62\x06\x70\x72\x6f\x74\x6f\x33"

var (
	file_app_trafficstats_config_proto_rawDescOnce sync.Once
	file_app_trafficstats_config_proto_rawDescData []byte
)

func file_app_trafficstats_config_proto_rawDescGZIP() []byte {
	file_app_trafficstats_config_proto_rawDescOnce.Do(func() {
		file_app_trafficstats_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_trafficstats_config_proto_rawDesc), len(file_app_trafficstats_config_proto_rawDesc)))
	})
	return file_app_trafficstats_config_proto_rawDescData
}

var file_app_trafficstats_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_trafficstats_config_proto_goTypes = []any{
	(*Config)(nil), // 0: xray.app.trafficstats.Config
}
var file_app_trafficstats_config_proto_depIdxs = []int32{}

func init() { file_app_trafficstats_config_proto_init() }
func file_app_trafficstats_config_proto_init() {
	if File_app_trafficstats_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_trafficstats_config_proto_rawDesc), len(file_app_trafficstats_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_trafficstats_config_proto_goTypes,
		DependencyIndexes: file_app_trafficstats_config_proto_depIdxs,
		MessageInfos:      file_app_trafficstats_config_proto_msgTypes,
	}.Build()
	File_app_trafficstats_config_proto = out.File
	file_app_trafficstats_config_proto_goTypes = nil
	file_app_trafficstats_config_proto_depIdxs = nil
}
