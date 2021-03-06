// Code generated by protoc-gen-go.
// source: dzhyun.zhibiao.proto
// DO NOT EDIT!

package dzhyun

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 指标类型
type ZhiBiao_ZBLeiXing int32

const (
	ZhiBiao_TYPE_EXPLORER  ZhiBiao_ZBLeiXing = 0
	ZhiBiao_TYPE_SYSTEST   ZhiBiao_ZBLeiXing = 1
	ZhiBiao_TYPE_MAIN_PICT ZhiBiao_ZBLeiXing = 2
	ZhiBiao_TYPE_MAIN_ADD  ZhiBiao_ZBLeiXing = 3
	ZhiBiao_TYPE_SUB_PICT  ZhiBiao_ZBLeiXing = 4
	ZhiBiao_TYPE_PAINT_IT  ZhiBiao_ZBLeiXing = 5
	ZhiBiao_TYPE_TEMP_INDI ZhiBiao_ZBLeiXing = 6
	ZhiBiao_TYPE_TECHNIQUE ZhiBiao_ZBLeiXing = 7
	ZhiBiao_TYPE_UNKNOWN   ZhiBiao_ZBLeiXing = 8
)

var ZhiBiao_ZBLeiXing_name = map[int32]string{
	0: "TYPE_EXPLORER",
	1: "TYPE_SYSTEST",
	2: "TYPE_MAIN_PICT",
	3: "TYPE_MAIN_ADD",
	4: "TYPE_SUB_PICT",
	5: "TYPE_PAINT_IT",
	6: "TYPE_TEMP_INDI",
	7: "TYPE_TECHNIQUE",
	8: "TYPE_UNKNOWN",
}
var ZhiBiao_ZBLeiXing_value = map[string]int32{
	"TYPE_EXPLORER":  0,
	"TYPE_SYSTEST":   1,
	"TYPE_MAIN_PICT": 2,
	"TYPE_MAIN_ADD":  3,
	"TYPE_SUB_PICT":  4,
	"TYPE_PAINT_IT":  5,
	"TYPE_TEMP_INDI": 6,
	"TYPE_TECHNIQUE": 7,
	"TYPE_UNKNOWN":   8,
}

func (x ZhiBiao_ZBLeiXing) String() string {
	return proto.EnumName(ZhiBiao_ZBLeiXing_name, int32(x))
}
func (ZhiBiao_ZBLeiXing) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0, 0} }

//   指标文本类型
type ZhiBiao_ZBWenBenLeiXing int32

const (
	ZhiBiao_TEXTTYPE_FORMULA ZhiBiao_ZBWenBenLeiXing = 0
	ZhiBiao_TEXTTYPE_LUA     ZhiBiao_ZBWenBenLeiXing = 1
	ZhiBiao_TEXTTYPE_UNKNOWN ZhiBiao_ZBWenBenLeiXing = 2
)

var ZhiBiao_ZBWenBenLeiXing_name = map[int32]string{
	0: "TEXTTYPE_FORMULA",
	1: "TEXTTYPE_LUA",
	2: "TEXTTYPE_UNKNOWN",
}
var ZhiBiao_ZBWenBenLeiXing_value = map[string]int32{
	"TEXTTYPE_FORMULA": 0,
	"TEXTTYPE_LUA":     1,
	"TEXTTYPE_UNKNOWN": 2,
}

func (x ZhiBiao_ZBWenBenLeiXing) String() string {
	return proto.EnumName(ZhiBiao_ZBWenBenLeiXing_name, int32(x))
}
func (ZhiBiao_ZBWenBenLeiXing) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0, 1} }

// 指标
type ZhiBiao struct {
	MingCheng      string                  `protobuf:"bytes,1,opt,name=MingCheng" json:"MingCheng,omitempty"`
	MiaoShu        string                  `protobuf:"bytes,2,opt,name=MiaoShu" json:"MiaoShu,omitempty"`
	YongFa         string                  `protobuf:"bytes,3,opt,name=YongFa" json:"YongFa,omitempty"`
	CanShuJingLing string                  `protobuf:"bytes,4,opt,name=CanShuJingLing" json:"CanShuJingLing,omitempty"`
	JianYiZu       string                  `protobuf:"bytes,5,opt,name=JianYiZu" json:"JianYiZu,omitempty"`
	WenBen         string                  `protobuf:"bytes,6,opt,name=WenBen" json:"WenBen,omitempty"`
	ShiJian        YFloat                   `protobuf:"varint,7,opt,name=ShiJian" json:"ShiJian,omitempty"`
	LeiXing        ZhiBiao_ZBLeiXing       `protobuf:"varint,8,opt,name=LeiXing,enum=dzhyun.ZhiBiao_ZBLeiXing" json:"LeiXing,omitempty"`
	WenBenLeiXing  ZhiBiao_ZBWenBenLeiXing `protobuf:"varint,9,opt,name=WenBenLeiXing,enum=dzhyun.ZhiBiao_ZBWenBenLeiXing" json:"WenBenLeiXing,omitempty"`
	BanBen         YFloat                   `protobuf:"varint,10,opt,name=BanBen" json:"BanBen,omitempty"`
	ShuXing        YFloat                   `protobuf:"varint,11,opt,name=ShuXing" json:"ShuXing,omitempty"`
	MoRenLeiXing   YFloat                   `protobuf:"varint,12,opt,name=MoRenLeiXing" json:"MoRenLeiXing,omitempty"`
	ZiJieMa        string                  `protobuf:"bytes,13,opt,name=ZiJieMa" json:"ZiJieMa,omitempty"`
	ChangYong      bool                    `protobuf:"varint,14,opt,name=ChangYong" json:"ChangYong,omitempty"`
	ZiDingYi       bool                    `protobuf:"varint,15,opt,name=ZiDingYi" json:"ZiDingYi,omitempty"`
	EWaiShuJu      []YFloat                 `protobuf:"varint,16,rep,packed,name=EWaiShuJu" json:"EWaiShuJu,omitempty"`
	CanShu         []*ZhiBiao_ZBCanShu     `protobuf:"bytes,17,rep,name=CanShu" json:"CanShu,omitempty"`
	ShuChu         []*ZhiBiao_ZBShuChu     `protobuf:"bytes,18,rep,name=ShuChu" json:"ShuChu,omitempty"`
}

func (m *ZhiBiao) Reset()                    { *m = ZhiBiao{} }
func (m *ZhiBiao) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiao) ProtoMessage()               {}
func (*ZhiBiao) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *ZhiBiao) GetMingCheng() string {
	if m != nil {
		return m.MingCheng
	}
	return ""
}

func (m *ZhiBiao) GetMiaoShu() string {
	if m != nil {
		return m.MiaoShu
	}
	return ""
}

func (m *ZhiBiao) GetYongFa() string {
	if m != nil {
		return m.YongFa
	}
	return ""
}

func (m *ZhiBiao) GetCanShuJingLing() string {
	if m != nil {
		return m.CanShuJingLing
	}
	return ""
}

func (m *ZhiBiao) GetJianYiZu() string {
	if m != nil {
		return m.JianYiZu
	}
	return ""
}

func (m *ZhiBiao) GetWenBen() string {
	if m != nil {
		return m.WenBen
	}
	return ""
}

func (m *ZhiBiao) GetShiJian() YFloat {
	if m != nil {
		return m.ShiJian
	}
	return 0
}

func (m *ZhiBiao) GetLeiXing() ZhiBiao_ZBLeiXing {
	if m != nil {
		return m.LeiXing
	}
	return ZhiBiao_TYPE_EXPLORER
}

func (m *ZhiBiao) GetWenBenLeiXing() ZhiBiao_ZBWenBenLeiXing {
	if m != nil {
		return m.WenBenLeiXing
	}
	return ZhiBiao_TEXTTYPE_FORMULA
}

func (m *ZhiBiao) GetBanBen() YFloat {
	if m != nil {
		return m.BanBen
	}
	return 0
}

func (m *ZhiBiao) GetShuXing() YFloat {
	if m != nil {
		return m.ShuXing
	}
	return 0
}

func (m *ZhiBiao) GetMoRenLeiXing() YFloat {
	if m != nil {
		return m.MoRenLeiXing
	}
	return 0
}

func (m *ZhiBiao) GetZiJieMa() string {
	if m != nil {
		return m.ZiJieMa
	}
	return ""
}

func (m *ZhiBiao) GetChangYong() bool {
	if m != nil {
		return m.ChangYong
	}
	return false
}

func (m *ZhiBiao) GetZiDingYi() bool {
	if m != nil {
		return m.ZiDingYi
	}
	return false
}

func (m *ZhiBiao) GetEWaiShuJu() []YFloat {
	if m != nil {
		return m.EWaiShuJu
	}
	return nil
}

func (m *ZhiBiao) GetCanShu() []*ZhiBiao_ZBCanShu {
	if m != nil {
		return m.CanShu
	}
	return nil
}

func (m *ZhiBiao) GetShuChu() []*ZhiBiao_ZBShuChu {
	if m != nil {
		return m.ShuChu
	}
	return nil
}

// 指标输出
type ZhiBiao_ZBShuChu struct {
	MingCheng       string                            `protobuf:"bytes,1,opt,name=MingCheng" json:"MingCheng,omitempty"`
	LeiXing         ZhiBiaoShuChu_ZBShuXing_SXLeiXing `protobuf:"varint,2,opt,name=LeiXing,enum=dzhyun.ZhiBiaoShuChu_ZBShuXing_SXLeiXing" json:"LeiXing,omitempty"`
	YiDong          YFloat                             `protobuf:"varint,3,opt,name=YiDong" json:"YiDong,omitempty"`
	ShuXing         YFloat                             `protobuf:"varint,4,opt,name=ShuXing" json:"ShuXing,omitempty"`
	YanSe           YFloat                             `protobuf:"varint,5,opt,name=YanSe" json:"YanSe,omitempty"`
	BianLiangWeiZhi YFloat                             `protobuf:"varint,6,opt,name=BianLiangWeiZhi" json:"BianLiangWeiZhi,omitempty"`
	KuoZhanShuXing  YFloat                             `protobuf:"varint,7,opt,name=KuoZhanShuXing" json:"KuoZhanShuXing,omitempty"`
}

func (m *ZhiBiao_ZBShuChu) Reset()                    { *m = ZhiBiao_ZBShuChu{} }
func (m *ZhiBiao_ZBShuChu) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiao_ZBShuChu) ProtoMessage()               {}
func (*ZhiBiao_ZBShuChu) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0, 0} }

func (m *ZhiBiao_ZBShuChu) GetMingCheng() string {
	if m != nil {
		return m.MingCheng
	}
	return ""
}

func (m *ZhiBiao_ZBShuChu) GetLeiXing() ZhiBiaoShuChu_ZBShuXing_SXLeiXing {
	if m != nil {
		return m.LeiXing
	}
	return ZhiBiaoShuChu_ZBShuXing_TYPE_TEMP_EXPRESION
}

func (m *ZhiBiao_ZBShuChu) GetYiDong() YFloat {
	if m != nil {
		return m.YiDong
	}
	return 0
}

func (m *ZhiBiao_ZBShuChu) GetShuXing() YFloat {
	if m != nil {
		return m.ShuXing
	}
	return 0
}

func (m *ZhiBiao_ZBShuChu) GetYanSe() YFloat {
	if m != nil {
		return m.YanSe
	}
	return 0
}

func (m *ZhiBiao_ZBShuChu) GetBianLiangWeiZhi() YFloat {
	if m != nil {
		return m.BianLiangWeiZhi
	}
	return 0
}

func (m *ZhiBiao_ZBShuChu) GetKuoZhanShuXing() YFloat {
	if m != nil {
		return m.KuoZhanShuXing
	}
	return 0
}

// 指标参数
type ZhiBiao_ZBCanShu struct {
	MingCheng  string `protobuf:"bytes,1,opt,name=MingCheng" json:"MingCheng,omitempty"`
	MoRenZhi   YFloat  `protobuf:"varint,2,opt,name=MoRenZhi" json:"MoRenZhi,omitempty"`
	ZuiDaZhi   YFloat  `protobuf:"varint,3,opt,name=ZuiDaZhi" json:"ZuiDaZhi,omitempty"`
	ZuiXiaoZhi YFloat  `protobuf:"varint,4,opt,name=ZuiXiaoZhi" json:"ZuiXiaoZhi,omitempty"`
	BuChang    YFloat  `protobuf:"varint,5,opt,name=BuChang" json:"BuChang,omitempty"`
}

func (m *ZhiBiao_ZBCanShu) Reset()                    { *m = ZhiBiao_ZBCanShu{} }
func (m *ZhiBiao_ZBCanShu) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiao_ZBCanShu) ProtoMessage()               {}
func (*ZhiBiao_ZBCanShu) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0, 1} }

func (m *ZhiBiao_ZBCanShu) GetMingCheng() string {
	if m != nil {
		return m.MingCheng
	}
	return ""
}

func (m *ZhiBiao_ZBCanShu) GetMoRenZhi() YFloat {
	if m != nil {
		return m.MoRenZhi
	}
	return 0
}

func (m *ZhiBiao_ZBCanShu) GetZuiDaZhi() YFloat {
	if m != nil {
		return m.ZuiDaZhi
	}
	return 0
}

func (m *ZhiBiao_ZBCanShu) GetZuiXiaoZhi() YFloat {
	if m != nil {
		return m.ZuiXiaoZhi
	}
	return 0
}

func (m *ZhiBiao_ZBCanShu) GetBuChang() YFloat {
	if m != nil {
		return m.BuChang
	}
	return 0
}

func init() {
	proto.RegisterType((*ZhiBiao)(nil), "dzhyun.ZhiBiao")
	proto.RegisterType((*ZhiBiao_ZBShuChu)(nil), "dzhyun.ZhiBiao.ZBShuChu")
	proto.RegisterType((*ZhiBiao_ZBCanShu)(nil), "dzhyun.ZhiBiao.ZBCanShu")
	proto.RegisterEnum("dzhyun.ZhiBiao_ZBLeiXing", ZhiBiao_ZBLeiXing_name, ZhiBiao_ZBLeiXing_value)
	proto.RegisterEnum("dzhyun.ZhiBiao_ZBWenBenLeiXing", ZhiBiao_ZBWenBenLeiXing_name, ZhiBiao_ZBWenBenLeiXing_value)
}

func init() { proto.RegisterFile("dzhyun.zhibiao.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x8e, 0xdb, 0xfc, 0x6c, 0xdb, 0xd4, 0x5d, 0x55, 0x47, 0x7b, 0xa2, 0x23, 0x88, 0x72,
	0x81, 0xcc, 0x4d, 0x84, 0xda, 0x27, 0x88, 0x1d, 0x57, 0xb8, 0x8d, 0x9d, 0xb0, 0x76, 0x94, 0xc4,
	0x37, 0xd1, 0x16, 0x2c, 0x7b, 0x90, 0x58, 0x23, 0xc0, 0x17, 0xf4, 0x05, 0x78, 0x07, 0xde, 0x80,
	0x27, 0xe0, 0xf5, 0xd0, 0xce, 0xda, 0x4e, 0x93, 0x0a, 0xb8, 0x8a, 0xbe, 0x6f, 0xbe, 0x99, 0x9d,
	0x99, 0x6f, 0x1c, 0x72, 0xf9, 0xee, 0x21, 0xff, 0x5a, 0xca, 0xf1, 0x43, 0x0e, 0xf7, 0x20, 0x8a,
	0xf1, 0xc7, 0x4f, 0xc5, 0x97, 0x82, 0xb6, 0x35, 0x3b, 0x18, 0xec, 0x47, 0xdf, 0xc3, 0xe7, 0x52,
	0x48, 0xad, 0x19, 0xfd, 0x20, 0xa4, 0x93, 0xe4, 0xe0, 0x80, 0x28, 0xe8, 0xff, 0xa4, 0x17, 0x80,
	0xcc, 0xdc, 0x3c, 0x95, 0x19, 0x33, 0x86, 0x86, 0xdd, 0xe3, 0x3b, 0x82, 0x32, 0xd2, 0x09, 0x40,
	0x14, 0x51, 0x5e, 0xb2, 0x16, 0xc6, 0x6a, 0x48, 0xff, 0x25, 0xed, 0x4d, 0x21, 0xb3, 0x1b, 0xc1,
	0x4c, 0x0c, 0x54, 0x88, 0xbe, 0x20, 0x7d, 0x57, 0xc8, 0x28, 0x2f, 0x6f, 0x41, 0x66, 0x33, 0x90,
	0x19, 0x3b, 0xc2, 0xf8, 0x01, 0x4b, 0x07, 0xa4, 0x7b, 0x0b, 0x42, 0x6e, 0x20, 0x29, 0xd9, 0x31,
	0x2a, 0x1a, 0xac, 0x6a, 0xaf, 0x52, 0xe9, 0xa4, 0x92, 0xb5, 0x75, 0x6d, 0x8d, 0x54, 0x37, 0x51,
	0x0e, 0x4a, 0xc6, 0x3a, 0x43, 0xc3, 0x36, 0x79, 0x0d, 0xe9, 0x35, 0xe9, 0xcc, 0x52, 0x58, 0xab,
	0xe7, 0xba, 0x43, 0xc3, 0xee, 0x5f, 0xfd, 0x37, 0xae, 0xe6, 0xaf, 0xe6, 0x1c, 0x27, 0x4e, 0x25,
	0xe0, 0xb5, 0x92, 0x7a, 0xe4, 0x4c, 0x17, 0xae, 0x53, 0x7b, 0x98, 0xfa, 0xfc, 0x69, 0xea, 0x9e,
	0x8c, 0xef, 0x67, 0xa9, 0x6e, 0x1d, 0x81, 0xdd, 0x12, 0x6c, 0xaa, 0x42, 0xba, 0xdb, 0x12, 0x0b,
	0x9f, 0xd4, 0xdd, 0x22, 0xa4, 0x23, 0x72, 0x1a, 0x14, 0x7c, 0xf7, 0xee, 0x29, 0x86, 0xf7, 0x38,
	0x95, 0x9d, 0xc0, 0x2d, 0xa4, 0x81, 0x60, 0x67, 0x7a, 0xf3, 0x15, 0x54, 0x8e, 0xb9, 0xb9, 0x90,
	0x99, 0x5a, 0x38, 0xeb, 0x0f, 0x0d, 0xbb, 0xcb, 0x77, 0x84, 0xda, 0x6b, 0x02, 0x53, 0x90, 0xd9,
	0x06, 0xd8, 0x39, 0x06, 0x1b, 0xac, 0x32, 0xbd, 0x95, 0x00, 0x65, 0x43, 0xc9, 0xac, 0xa1, 0x69,
	0x9b, 0x7c, 0x47, 0xd0, 0x57, 0xa4, 0xad, 0x3d, 0x62, 0x17, 0x43, 0xd3, 0x3e, 0xb9, 0x62, 0x4f,
	0xf7, 0xa0, 0xe3, 0xbc, 0xd2, 0xa9, 0x8c, 0x28, 0x2f, 0xdd, 0xbc, 0x64, 0xf4, 0x77, 0x19, 0x3a,
	0xce, 0x2b, 0xdd, 0xe0, 0x5b, 0x8b, 0x74, 0x6b, 0xf2, 0x2f, 0xa7, 0xe7, 0xee, 0x2c, 0x6d, 0xa1,
	0x2f, 0x2f, 0x0f, 0xaa, 0xeb, 0x2a, 0xfa, 0x0d, 0x25, 0x1b, 0x47, 0xeb, 0x27, 0x16, 0xab, 0x2b,
	0x85, 0xa9, 0x5a, 0x94, 0xa9, 0xbd, 0xd1, 0xe8, 0xb1, 0x37, 0x47, 0xfb, 0xde, 0x5c, 0x92, 0xe3,
	0x8d, 0x90, 0x51, 0x8a, 0x47, 0x69, 0x72, 0x0d, 0xa8, 0x4d, 0xce, 0x1d, 0x10, 0x72, 0x06, 0x42,
	0x66, 0xab, 0x14, 0x92, 0x1c, 0xf0, 0x34, 0x4d, 0x7e, 0x48, 0xab, 0xfb, 0xbf, 0x2b, 0x8b, 0x24,
	0xc7, 0x0d, 0xe1, 0x03, 0xfa, 0x54, 0x0f, 0xd8, 0xc1, 0x77, 0x43, 0x6d, 0xa2, 0x5a, 0xe4, 0x9f,
	0x37, 0x31, 0x20, 0x5d, 0x3c, 0x0d, 0xf5, 0x6a, 0x0b, 0x8b, 0x35, 0x18, 0xed, 0x2e, 0x61, 0x2a,
	0x54, 0x4c, 0x8f, 0xd8, 0x60, 0xfa, 0x8c, 0x90, 0xa4, 0x84, 0x35, 0x88, 0x42, 0x45, 0xf5, 0x9c,
	0x8f, 0x18, 0xb5, 0x04, 0xa7, 0xc4, 0xcb, 0xa9, 0x86, 0xad, 0xe1, 0xe8, 0xa7, 0x41, 0x7a, 0xcd,
	0x07, 0x43, 0x2f, 0xc8, 0x59, 0xbc, 0x59, 0x78, 0x5b, 0x6f, 0xbd, 0x98, 0xcd, 0xb9, 0xc7, 0xad,
	0x7f, 0xa8, 0x45, 0x4e, 0x91, 0x8a, 0x36, 0x51, 0xec, 0x45, 0xb1, 0x65, 0x50, 0x4a, 0xfa, 0xc8,
	0x04, 0x13, 0x3f, 0xdc, 0x2e, 0x7c, 0x37, 0xb6, 0x5a, 0x4d, 0x22, 0x72, 0x93, 0xe9, 0xd4, 0x32,
	0x1b, 0x2a, 0x5a, 0x3a, 0x5a, 0x75, 0xd4, 0x50, 0x8b, 0x89, 0x1f, 0xc6, 0x5b, 0x3f, 0xb6, 0x8e,
	0x9b, 0x62, 0xb1, 0x17, 0x2c, 0xb6, 0x7e, 0x38, 0xf5, 0xad, 0xf6, 0x23, 0xce, 0x7d, 0x1d, 0xfa,
	0x6f, 0x96, 0x9e, 0xd5, 0x69, 0xda, 0x58, 0x86, 0x77, 0xe1, 0x7c, 0x15, 0x5a, 0xdd, 0xd1, 0x9c,
	0x9c, 0x1f, 0x7c, 0xae, 0xf4, 0x92, 0x58, 0xb1, 0xb7, 0x8e, 0x51, 0x78, 0x33, 0xe7, 0xc1, 0x72,
	0x36, 0xa9, 0x26, 0xa8, 0xd9, 0xd9, 0x72, 0x62, 0x19, 0x7b, 0xba, 0xba, 0x60, 0xcb, 0xa1, 0xc4,
	0x7a, 0x5b, 0x7c, 0xa8, 0x4f, 0x0f, 0xff, 0x3f, 0xef, 0xdb, 0xf8, 0x73, 0xfd, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x05, 0x1b, 0x5a, 0xba, 0x82, 0x05, 0x00, 0x00,
}
