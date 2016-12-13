// Code generated by protoc-gen-go.
// source: dzhyun.jianpanbao.proto
// DO NOT EDIT!

package dzhyun

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 业务类型
type JPBLeiXing int32

const (
	JPBLeiXing_TYPE_OBJ        JPBLeiXing = 0
	JPBLeiXing_TYPE_INDI       JPBLeiXing = 1
	JPBLeiXing_TYPE_TOPIC      JPBLeiXing = 2
	JPBLeiXing_TYPE_LHB        JPBLeiXing = 3
	JPBLeiXing_TYPE_BLKOBJ     JPBLeiXing = 4
	JPBLeiXing_TYPE_TOPICANALY JPBLeiXing = 5
	JPBLeiXing_TYPE_OBJPHONE   JPBLeiXing = 6
)

var JPBLeiXing_name = map[int32]string{
	0: "TYPE_OBJ",
	1: "TYPE_INDI",
	2: "TYPE_TOPIC",
	3: "TYPE_LHB",
	4: "TYPE_BLKOBJ",
	5: "TYPE_TOPICANALY",
	6: "TYPE_OBJPHONE",
}
var JPBLeiXing_value = map[string]int32{
	"TYPE_OBJ":        0,
	"TYPE_INDI":       1,
	"TYPE_TOPIC":      2,
	"TYPE_LHB":        3,
	"TYPE_BLKOBJ":     4,
	"TYPE_TOPICANALY": 5,
	"TYPE_OBJPHONE":   6,
}

func (x JPBLeiXing) String() string {
	return proto.EnumName(JPBLeiXing_name, int32(x))
}
func (JPBLeiXing) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

type JPBShuJu struct {
	DaiMa     string `protobuf:"bytes,1,opt,name=DaiMa" json:"DaiMa,omitempty"`
	MingCheng string `protobuf:"bytes,2,opt,name=MingCheng" json:"MingCheng,omitempty"`
	ShuXing   string `protobuf:"bytes,3,opt,name=ShuXing" json:"ShuXing,omitempty"`
	KuoZhan   string `protobuf:"bytes,4,opt,name=KuoZhan" json:"KuoZhan,omitempty"`
}

func (m *JPBShuJu) Reset()                    { *m = JPBShuJu{} }
func (m *JPBShuJu) String() string            { return proto.CompactTextString(m) }
func (*JPBShuJu) ProtoMessage()               {}
func (*JPBShuJu) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *JPBShuJu) GetDaiMa() string {
	if m != nil {
		return m.DaiMa
	}
	return ""
}

func (m *JPBShuJu) GetMingCheng() string {
	if m != nil {
		return m.MingCheng
	}
	return ""
}

func (m *JPBShuJu) GetShuXing() string {
	if m != nil {
		return m.ShuXing
	}
	return ""
}

func (m *JPBShuJu) GetKuoZhan() string {
	if m != nil {
		return m.KuoZhan
	}
	return ""
}

type JPBShuChu struct {
	LeiXing JPBLeiXing  `protobuf:"varint,1,opt,name=LeiXing,enum=dzhyun.JPBLeiXing" json:"LeiXing,omitempty"`
	ShuJu   []*JPBShuJu `protobuf:"bytes,2,rep,name=ShuJu" json:"ShuJu,omitempty"`
}

func (m *JPBShuChu) Reset()                    { *m = JPBShuChu{} }
func (m *JPBShuChu) String() string            { return proto.CompactTextString(m) }
func (*JPBShuChu) ProtoMessage()               {}
func (*JPBShuChu) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *JPBShuChu) GetLeiXing() JPBLeiXing {
	if m != nil {
		return m.LeiXing
	}
	return JPBLeiXing_TYPE_OBJ
}

func (m *JPBShuChu) GetShuJu() []*JPBShuJu {
	if m != nil {
		return m.ShuJu
	}
	return nil
}

type JianPanBaoShuChu struct {
	GuanJianZi string       `protobuf:"bytes,1,opt,name=GuanJianZi" json:"GuanJianZi,omitempty"`
	JieGuo     []*JPBShuChu `protobuf:"bytes,2,rep,name=JieGuo" json:"JieGuo,omitempty"`
}

func (m *JianPanBaoShuChu) Reset()                    { *m = JianPanBaoShuChu{} }
func (m *JianPanBaoShuChu) String() string            { return proto.CompactTextString(m) }
func (*JianPanBaoShuChu) ProtoMessage()               {}
func (*JianPanBaoShuChu) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *JianPanBaoShuChu) GetGuanJianZi() string {
	if m != nil {
		return m.GuanJianZi
	}
	return ""
}

func (m *JianPanBaoShuChu) GetJieGuo() []*JPBShuChu {
	if m != nil {
		return m.JieGuo
	}
	return nil
}

func init() {
	proto.RegisterType((*JPBShuJu)(nil), "dzhyun.JPBShuJu")
	proto.RegisterType((*JPBShuChu)(nil), "dzhyun.JPBShuChu")
	proto.RegisterType((*JianPanBaoShuChu)(nil), "dzhyun.JianPanBaoShuChu")
	proto.RegisterEnum("dzhyun.JPBLeiXing", JPBLeiXing_name, JPBLeiXing_value)
}

func init() { proto.RegisterFile("dzhyun.jianpanbao.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x4f, 0xc2, 0x40,
	0x14, 0xc4, 0x2d, 0x7f, 0x0a, 0x7d, 0x08, 0x2c, 0x4f, 0x13, 0x7b, 0x30, 0x86, 0x70, 0x30, 0x68,
	0x0c, 0x07, 0xfc, 0x04, 0xb4, 0x10, 0xa0, 0x14, 0x68, 0x0a, 0x07, 0x21, 0x31, 0x66, 0x51, 0xc2,
	0xae, 0x89, 0xbb, 0x84, 0xb8, 0x07, 0x3d, 0xf9, 0xd1, 0x4d, 0xb7, 0x5b, 0xc1, 0x53, 0x33, 0xf3,
	0x9b, 0xbe, 0x99, 0xa6, 0x70, 0xf5, 0xf6, 0xcd, 0xbe, 0x94, 0xe8, 0xbc, 0x73, 0x2a, 0xf6, 0x54,
	0x6c, 0xa8, 0xec, 0xec, 0x0f, 0xf2, 0x53, 0xa2, 0x9d, 0x82, 0xd6, 0x01, 0xca, 0x41, 0xe4, 0x2d,
	0x98, 0x0a, 0x14, 0x5e, 0x42, 0xb1, 0x4f, 0xf9, 0x94, 0xba, 0x56, 0xd3, 0x6a, 0x3b, 0x71, 0x2a,
	0xf0, 0x1a, 0x9c, 0x29, 0x17, 0x3b, 0x9f, 0x6d, 0xc5, 0xce, 0xcd, 0x69, 0x72, 0x34, 0xd0, 0x85,
	0xd2, 0x82, 0xa9, 0x27, 0x2e, 0x76, 0x6e, 0x5e, 0xb3, 0x4c, 0x26, 0x64, 0xa2, 0xe4, 0x9a, 0x51,
	0xe1, 0x16, 0x52, 0x62, 0x64, 0x8b, 0x82, 0x93, 0x76, 0xfa, 0x4c, 0xe1, 0x03, 0x94, 0xc2, 0x2d,
	0xd7, 0x07, 0x92, 0xda, 0x5a, 0x17, 0x3b, 0x66, 0x73, 0x10, 0x79, 0x86, 0xc4, 0x59, 0x04, 0x6f,
	0xa1, 0xa8, 0xb7, 0xba, 0xb9, 0x66, 0xbe, 0x5d, 0xe9, 0x92, 0x93, 0xac, 0xf6, 0xe3, 0x14, 0xb7,
	0x9e, 0x81, 0x04, 0x9c, 0x8a, 0x88, 0x0a, 0x8f, 0x4a, 0xd3, 0x74, 0x03, 0x30, 0x54, 0x54, 0x24,
	0xfe, 0x9a, 0x9b, 0x6f, 0x3c, 0x71, 0xf0, 0x0e, 0xec, 0x80, 0x6f, 0x87, 0x4a, 0x9a, 0xe3, 0x8d,
	0xff, 0xc7, 0x7d, 0xa6, 0x62, 0x13, 0xb8, 0xff, 0xb1, 0x00, 0x8e, 0xf3, 0xf0, 0x1c, 0xca, 0xcb,
	0x55, 0x34, 0x78, 0x99, 0x7b, 0x01, 0x39, 0xc3, 0x2a, 0x38, 0x5a, 0x8d, 0x67, 0xfd, 0x31, 0xb1,
	0xb0, 0x06, 0xa0, 0xe5, 0x72, 0x1e, 0x8d, 0x7d, 0x92, 0xfb, 0x0b, 0x87, 0x23, 0x8f, 0xe4, 0xb1,
	0x0e, 0x15, 0xad, 0xbc, 0x70, 0x92, 0xbc, 0x5d, 0xc0, 0x0b, 0xa8, 0x1f, 0xe3, 0xbd, 0x59, 0x2f,
	0x5c, 0x91, 0x22, 0x36, 0xa0, 0x9a, 0x15, 0x44, 0xa3, 0xf9, 0x6c, 0x40, 0x6c, 0x0f, 0x81, 0xbc,
	0xca, 0x8f, 0x6c, 0xa2, 0xfe, 0xa9, 0x1b, 0x5b, 0x3f, 0x1e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x84, 0x45, 0x6f, 0x7e, 0xf6, 0x01, 0x00, 0x00,
}