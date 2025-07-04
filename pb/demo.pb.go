// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: demo.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProcessDemoReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Path          string                 `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Player        string                 `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessDemoReq) Reset() {
	*x = ProcessDemoReq{}
	mi := &file_demo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessDemoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDemoReq) ProtoMessage() {}

func (x *ProcessDemoReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDemoReq.ProtoReflect.Descriptor instead.
func (*ProcessDemoReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessDemoReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ProcessDemoReq) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

type KillCounter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int32                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	HeadshotCount int32                  `protobuf:"varint,2,opt,name=headshotCount,proto3" json:"headshotCount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KillCounter) Reset() {
	*x = KillCounter{}
	mi := &file_demo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KillCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillCounter) ProtoMessage() {}

func (x *KillCounter) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillCounter.ProtoReflect.Descriptor instead.
func (*KillCounter) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *KillCounter) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *KillCounter) GetHeadshotCount() int32 {
	if x != nil {
		return x.HeadshotCount
	}
	return 0
}

type Flash struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	EnemiesFlashed int32                  `protobuf:"varint,1,opt,name=enemiesFlashed,proto3" json:"enemiesFlashed,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Flash) Reset() {
	*x = Flash{}
	mi := &file_demo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Flash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flash) ProtoMessage() {}

func (x *Flash) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flash.ProtoReflect.Descriptor instead.
func (*Flash) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{2}
}

func (x *Flash) GetEnemiesFlashed() int32 {
	if x != nil {
		return x.EnemiesFlashed
	}
	return 0
}

type Position struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             float64                `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y             float64                `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Position) Reset() {
	*x = Position{}
	mi := &file_demo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{3}
}

func (x *Position) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Kill struct {
	state             protoimpl.MessageState  `protogen:"open.v1"`
	Overall           *KillCounter            `protobuf:"bytes,1,opt,name=overall,proto3" json:"overall,omitempty"`
	Terrorists        *KillCounter            `protobuf:"bytes,2,opt,name=terrorists,proto3" json:"terrorists,omitempty"`
	CounterTerrorists *KillCounter            `protobuf:"bytes,3,opt,name=counterTerrorists,proto3" json:"counterTerrorists,omitempty"`
	Weapons           map[string]*KillCounter `protobuf:"bytes,4,rep,name=weapons,proto3" json:"weapons,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	KillPositions     []*Position             `protobuf:"bytes,5,rep,name=killPositions,proto3" json:"killPositions,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Kill) Reset() {
	*x = Kill{}
	mi := &file_demo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Kill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kill) ProtoMessage() {}

func (x *Kill) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kill.ProtoReflect.Descriptor instead.
func (*Kill) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{4}
}

func (x *Kill) GetOverall() *KillCounter {
	if x != nil {
		return x.Overall
	}
	return nil
}

func (x *Kill) GetTerrorists() *KillCounter {
	if x != nil {
		return x.Terrorists
	}
	return nil
}

func (x *Kill) GetCounterTerrorists() *KillCounter {
	if x != nil {
		return x.CounterTerrorists
	}
	return nil
}

func (x *Kill) GetWeapons() map[string]*KillCounter {
	if x != nil {
		return x.Weapons
	}
	return nil
}

func (x *Kill) GetKillPositions() []*Position {
	if x != nil {
		return x.KillPositions
	}
	return nil
}

type GameStats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kills         map[string]*Kill       `protobuf:"bytes,1,rep,name=kills,proto3" json:"kills,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Flashes       map[string]*Flash      `protobuf:"bytes,2,rep,name=flashes,proto3" json:"flashes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GameStats) Reset() {
	*x = GameStats{}
	mi := &file_demo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStats) ProtoMessage() {}

func (x *GameStats) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStats.ProtoReflect.Descriptor instead.
func (*GameStats) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{5}
}

func (x *GameStats) GetKills() map[string]*Kill {
	if x != nil {
		return x.Kills
	}
	return nil
}

func (x *GameStats) GetFlashes() map[string]*Flash {
	if x != nil {
		return x.Flashes
	}
	return nil
}

type ProcessDemoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Player        string                 `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Map           string                 `protobuf:"bytes,2,opt,name=map,proto3" json:"map,omitempty"`
	Players       []string               `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
	Stats         *GameStats             `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessDemoResponse) Reset() {
	*x = ProcessDemoResponse{}
	mi := &file_demo_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessDemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDemoResponse) ProtoMessage() {}

func (x *ProcessDemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDemoResponse.ProtoReflect.Descriptor instead.
func (*ProcessDemoResponse) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{6}
}

func (x *ProcessDemoResponse) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *ProcessDemoResponse) GetMap() string {
	if x != nil {
		return x.Map
	}
	return ""
}

func (x *ProcessDemoResponse) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *ProcessDemoResponse) GetStats() *GameStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type Pos struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             float64                `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y             float64                `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pos) Reset() {
	*x = Pos{}
	mi := &file_demo_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pos) ProtoMessage() {}

func (x *Pos) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pos.ProtoReflect.Descriptor instead.
func (*Pos) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{7}
}

func (x *Pos) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Pos) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_demo_proto protoreflect.FileDescriptor

const file_demo_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"demo.proto\x12\x02pb\"<\n" +
	"\x0eProcessDemoReq\x12\x12\n" +
	"\x04path\x18\x01 \x01(\tR\x04path\x12\x16\n" +
	"\x06player\x18\x02 \x01(\tR\x06player\"I\n" +
	"\vKillCounter\x12\x14\n" +
	"\x05count\x18\x01 \x01(\x05R\x05count\x12$\n" +
	"\rheadshotCount\x18\x02 \x01(\x05R\rheadshotCount\"/\n" +
	"\x05Flash\x12&\n" +
	"\x0eenemiesFlashed\x18\x01 \x01(\x05R\x0eenemiesFlashed\"&\n" +
	"\bPosition\x12\f\n" +
	"\x01x\x18\x01 \x01(\x01R\x01x\x12\f\n" +
	"\x01y\x18\x02 \x01(\x01R\x01y\"\xd3\x02\n" +
	"\x04Kill\x12)\n" +
	"\aoverall\x18\x01 \x01(\v2\x0f.pb.KillCounterR\aoverall\x12/\n" +
	"\n" +
	"terrorists\x18\x02 \x01(\v2\x0f.pb.KillCounterR\n" +
	"terrorists\x12=\n" +
	"\x11counterTerrorists\x18\x03 \x01(\v2\x0f.pb.KillCounterR\x11counterTerrorists\x12/\n" +
	"\aweapons\x18\x04 \x03(\v2\x15.pb.Kill.WeaponsEntryR\aweapons\x122\n" +
	"\rkillPositions\x18\x05 \x03(\v2\f.pb.PositionR\rkillPositions\x1aK\n" +
	"\fWeaponsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12%\n" +
	"\x05value\x18\x02 \x01(\v2\x0f.pb.KillCounterR\x05value:\x028\x01\"\xfc\x01\n" +
	"\tGameStats\x12.\n" +
	"\x05kills\x18\x01 \x03(\v2\x18.pb.GameStats.KillsEntryR\x05kills\x124\n" +
	"\aflashes\x18\x02 \x03(\v2\x1a.pb.GameStats.FlashesEntryR\aflashes\x1aB\n" +
	"\n" +
	"KillsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x1e\n" +
	"\x05value\x18\x02 \x01(\v2\b.pb.KillR\x05value:\x028\x01\x1aE\n" +
	"\fFlashesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x1f\n" +
	"\x05value\x18\x02 \x01(\v2\t.pb.FlashR\x05value:\x028\x01\"~\n" +
	"\x13ProcessDemoResponse\x12\x16\n" +
	"\x06player\x18\x01 \x01(\tR\x06player\x12\x10\n" +
	"\x03map\x18\x02 \x01(\tR\x03map\x12\x18\n" +
	"\aplayers\x18\x03 \x03(\tR\aplayers\x12#\n" +
	"\x05stats\x18\x04 \x01(\v2\r.pb.GameStatsR\x05stats\"!\n" +
	"\x03Pos\x12\f\n" +
	"\x01x\x18\x01 \x01(\x01R\x01x\x12\f\n" +
	"\x01y\x18\x02 \x01(\x01R\x01y2t\n" +
	"\x05Stats\x12<\n" +
	"\vProcessDemo\x12\x12.pb.ProcessDemoReq\x1a\x17.pb.ProcessDemoResponse\"\x00\x12-\n" +
	"\n" +
	"ReplayDemo\x12\x12.pb.ProcessDemoReq\x1a\a.pb.Pos\"\x000\x01B&Z$github.com/jakub-galecki/cs_stats/pbb\x06proto3"

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData []byte
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_demo_proto_rawDesc), len(file_demo_proto_rawDesc)))
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_demo_proto_goTypes = []any{
	(*ProcessDemoReq)(nil),      // 0: pb.ProcessDemoReq
	(*KillCounter)(nil),         // 1: pb.KillCounter
	(*Flash)(nil),               // 2: pb.Flash
	(*Position)(nil),            // 3: pb.Position
	(*Kill)(nil),                // 4: pb.Kill
	(*GameStats)(nil),           // 5: pb.GameStats
	(*ProcessDemoResponse)(nil), // 6: pb.ProcessDemoResponse
	(*Pos)(nil),                 // 7: pb.Pos
	nil,                         // 8: pb.Kill.WeaponsEntry
	nil,                         // 9: pb.GameStats.KillsEntry
	nil,                         // 10: pb.GameStats.FlashesEntry
}
var file_demo_proto_depIdxs = []int32{
	1,  // 0: pb.Kill.overall:type_name -> pb.KillCounter
	1,  // 1: pb.Kill.terrorists:type_name -> pb.KillCounter
	1,  // 2: pb.Kill.counterTerrorists:type_name -> pb.KillCounter
	8,  // 3: pb.Kill.weapons:type_name -> pb.Kill.WeaponsEntry
	3,  // 4: pb.Kill.killPositions:type_name -> pb.Position
	9,  // 5: pb.GameStats.kills:type_name -> pb.GameStats.KillsEntry
	10, // 6: pb.GameStats.flashes:type_name -> pb.GameStats.FlashesEntry
	5,  // 7: pb.ProcessDemoResponse.stats:type_name -> pb.GameStats
	1,  // 8: pb.Kill.WeaponsEntry.value:type_name -> pb.KillCounter
	4,  // 9: pb.GameStats.KillsEntry.value:type_name -> pb.Kill
	2,  // 10: pb.GameStats.FlashesEntry.value:type_name -> pb.Flash
	0,  // 11: pb.Stats.ProcessDemo:input_type -> pb.ProcessDemoReq
	0,  // 12: pb.Stats.ReplayDemo:input_type -> pb.ProcessDemoReq
	6,  // 13: pb.Stats.ProcessDemo:output_type -> pb.ProcessDemoResponse
	7,  // 14: pb.Stats.ReplayDemo:output_type -> pb.Pos
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_demo_proto_rawDesc), len(file_demo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}
