// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.16.0
// source: travel.proto

package travel

import (
	v1 "github.com/webmin7761/go-school/homework/final/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TravelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//起飞机场
	OrgAirport string `protobuf:"bytes,1,opt,name=org_airport,json=orgAirport,proto3" json:"org_airport,omitempty"`
	//到达机场
	ArrAirport string `protobuf:"bytes,2,opt,name=arr_airport,json=arrAirport,proto3" json:"arr_airport,omitempty"`
	//航班日期
	FlightDatetime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=flight_datetime,json=flightDatetime,proto3" json:"flight_datetime,omitempty"`
}

func (x *TravelRequest) Reset() {
	*x = TravelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_travel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelRequest) ProtoMessage() {}

func (x *TravelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_travel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelRequest.ProtoReflect.Descriptor instead.
func (*TravelRequest) Descriptor() ([]byte, []int) {
	return file_travel_proto_rawDescGZIP(), []int{0}
}

func (x *TravelRequest) GetOrgAirport() string {
	if x != nil {
		return x.OrgAirport
	}
	return ""
}

func (x *TravelRequest) GetArrAirport() string {
	if x != nil {
		return x.ArrAirport
	}
	return ""
}

func (x *TravelRequest) GetFlightDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.FlightDatetime
	}
	return nil
}

type TravelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//处理结果
	Result *v1.Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	//起飞机场
	OrgAirport string `protobuf:"bytes,2,opt,name=org_airport,json=orgAirport,proto3" json:"org_airport,omitempty"`
	//到达机场
	ArrAirport string `protobuf:"bytes,3,opt,name=arr_airport,json=arrAirport,proto3" json:"arr_airport,omitempty"`
	//航班日期
	FlightDatetime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=flight_datetime,json=flightDatetime,proto3" json:"flight_datetime,omitempty"`
	//推荐信息
	TravelMessage string `protobuf:"bytes,5,opt,name=travel_message,json=travelMessage,proto3" json:"travel_message,omitempty"`
}

func (x *TravelResponse) Reset() {
	*x = TravelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_travel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelResponse) ProtoMessage() {}

func (x *TravelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_travel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelResponse.ProtoReflect.Descriptor instead.
func (*TravelResponse) Descriptor() ([]byte, []int) {
	return file_travel_proto_rawDescGZIP(), []int{1}
}

func (x *TravelResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *TravelResponse) GetOrgAirport() string {
	if x != nil {
		return x.OrgAirport
	}
	return ""
}

func (x *TravelResponse) GetArrAirport() string {
	if x != nil {
		return x.ArrAirport
	}
	return ""
}

func (x *TravelResponse) GetFlightDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.FlightDatetime
	}
	return nil
}

func (x *TravelResponse) GetTravelMessage() string {
	if x != nil {
		return x.TravelMessage
	}
	return ""
}

var File_travel_proto protoreflect.FileDescriptor

var file_travel_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01,
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x67, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x67, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x72, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x72, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x43, 0x0a, 0x0f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xe6, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x67, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x67, 0x41, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x72, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x72, 0x41, 0x69, 0x72, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x43, 0x0a, 0x0f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x64, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x6d, 0x69, 0x6e, 0x37, 0x37, 0x36, 0x31, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_travel_proto_rawDescOnce sync.Once
	file_travel_proto_rawDescData = file_travel_proto_rawDesc
)

func file_travel_proto_rawDescGZIP() []byte {
	file_travel_proto_rawDescOnce.Do(func() {
		file_travel_proto_rawDescData = protoimpl.X.CompressGZIP(file_travel_proto_rawDescData)
	})
	return file_travel_proto_rawDescData
}

var file_travel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_travel_proto_goTypes = []interface{}{
	(*TravelRequest)(nil),         // 0: travel.TravelRequest
	(*TravelResponse)(nil),        // 1: travel.TravelResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*v1.Result)(nil),             // 3: common.Result
}
var file_travel_proto_depIdxs = []int32{
	2, // 0: travel.TravelRequest.flight_datetime:type_name -> google.protobuf.Timestamp
	3, // 1: travel.TravelResponse.result:type_name -> common.Result
	2, // 2: travel.TravelResponse.flight_datetime:type_name -> google.protobuf.Timestamp
	0, // 3: travel.TravelService.Query:input_type -> travel.TravelRequest
	1, // 4: travel.TravelService.Query:output_type -> travel.TravelResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_travel_proto_init() }
func file_travel_proto_init() {
	if File_travel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_travel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_travel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_travel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_travel_proto_goTypes,
		DependencyIndexes: file_travel_proto_depIdxs,
		MessageInfos:      file_travel_proto_msgTypes,
	}.Build()
	File_travel_proto = out.File
	file_travel_proto_rawDesc = nil
	file_travel_proto_goTypes = nil
	file_travel_proto_depIdxs = nil
}
