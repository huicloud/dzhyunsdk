// Code generated by protoc-gen-go.
// source: dzhyun.zhibiaojisuan.proto
// DO NOT EDIT!

package dzhyun

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 输出属性类型
type ZhiBiaoShuChu_ZBShuXing_SXLeiXing int32

const (
	ZhiBiaoShuChu_ZBShuXing_TYPE_TEMP_EXPRESION  ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 0
	ZhiBiaoShuChu_ZBShuXing_TYPE_CURV_LINE       ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 1
	ZhiBiaoShuChu_ZBShuXing_TYPE_STICK_LINE      ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 2
	ZhiBiaoShuChu_ZBShuXing_TYPE_COLORSTICK_LINE ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 3
	ZhiBiaoShuChu_ZBShuXing_TYPE_VOLSTICK_LINE   ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 4
	ZhiBiaoShuChu_ZBShuXing_TYPE_LINESTICK_LINE  ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 5
	ZhiBiaoShuChu_ZBShuXing_TYPE_CROSS_DOT       ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 6
	ZhiBiaoShuChu_ZBShuXing_TYPE_CIRCLE_DOT      ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 7
	ZhiBiaoShuChu_ZBShuXing_TYPE_POINT_DOT       ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 8
	ZhiBiaoShuChu_ZBShuXing_TYPE_STICK3D_LINE    ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 9
	ZhiBiaoShuChu_ZBShuXing_TYPE_COLOR3D_LINE    ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 10
	ZhiBiaoShuChu_ZBShuXing_TYPE_DOT_DOT         ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 11
	ZhiBiaoShuChu_ZBShuXing_TYPE_DASH_DOT        ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 12
	ZhiBiaoShuChu_ZBShuXing_TYPE_PERCENT_BAR     ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 13
	ZhiBiaoShuChu_ZBShuXing_TYPE_ENTER_LONG      ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 100
	ZhiBiaoShuChu_ZBShuXing_TYPE_EXIT_LONG       ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 101
	ZhiBiaoShuChu_ZBShuXing_TYPE_ENTER_SHORT     ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 102
	ZhiBiaoShuChu_ZBShuXing_TYPE_EXIT_SHORT      ZhiBiaoShuChu_ZBShuXing_SXLeiXing = 103
)

var ZhiBiaoShuChu_ZBShuXing_SXLeiXing_name = map[int32]string{
	0:   "TYPE_TEMP_EXPRESION",
	1:   "TYPE_CURV_LINE",
	2:   "TYPE_STICK_LINE",
	3:   "TYPE_COLORSTICK_LINE",
	4:   "TYPE_VOLSTICK_LINE",
	5:   "TYPE_LINESTICK_LINE",
	6:   "TYPE_CROSS_DOT",
	7:   "TYPE_CIRCLE_DOT",
	8:   "TYPE_POINT_DOT",
	9:   "TYPE_STICK3D_LINE",
	10:  "TYPE_COLOR3D_LINE",
	11:  "TYPE_DOT_DOT",
	12:  "TYPE_DASH_DOT",
	13:  "TYPE_PERCENT_BAR",
	100: "TYPE_ENTER_LONG",
	101: "TYPE_EXIT_LONG",
	102: "TYPE_ENTER_SHORT",
	103: "TYPE_EXIT_SHORT",
}
var ZhiBiaoShuChu_ZBShuXing_SXLeiXing_value = map[string]int32{
	"TYPE_TEMP_EXPRESION":  0,
	"TYPE_CURV_LINE":       1,
	"TYPE_STICK_LINE":      2,
	"TYPE_COLORSTICK_LINE": 3,
	"TYPE_VOLSTICK_LINE":   4,
	"TYPE_LINESTICK_LINE":  5,
	"TYPE_CROSS_DOT":       6,
	"TYPE_CIRCLE_DOT":      7,
	"TYPE_POINT_DOT":       8,
	"TYPE_STICK3D_LINE":    9,
	"TYPE_COLOR3D_LINE":    10,
	"TYPE_DOT_DOT":         11,
	"TYPE_DASH_DOT":        12,
	"TYPE_PERCENT_BAR":     13,
	"TYPE_ENTER_LONG":      100,
	"TYPE_EXIT_LONG":       101,
	"TYPE_ENTER_SHORT":     102,
	"TYPE_EXIT_SHORT":      103,
}

func (x ZhiBiaoShuChu_ZBShuXing_SXLeiXing) String() string {
	return proto.EnumName(ZhiBiaoShuChu_ZBShuXing_SXLeiXing_name, int32(x))
}
func (ZhiBiaoShuChu_ZBShuXing_SXLeiXing) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor14, []int{0, 1, 0}
}

// 绘图类型
type ZhiBiaoShuChu_ZBHuiTu_HTLeiXing int32

const (
	ZhiBiaoShuChu_ZBHuiTu_TYPE_NOLINE       ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 0
	ZhiBiaoShuChu_ZBHuiTu_TYPE_POLYLINE     ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 1
	ZhiBiaoShuChu_ZBHuiTu_TYPE_LINE         ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 2
	ZhiBiaoShuChu_ZBHuiTu_TYPE_STICKLINE    ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 3
	ZhiBiaoShuChu_ZBHuiTu_TYPE_TEXT         ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 4
	ZhiBiaoShuChu_ZBHuiTu_TYPE_ICON         ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 5
	ZhiBiaoShuChu_ZBHuiTu_TYPE_TIP_TEXT     ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 6
	ZhiBiaoShuChu_ZBHuiTu_TYPE_BACK_GRD     ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 7
	ZhiBiaoShuChu_ZBHuiTu_TYPE_BACK_GRDLAST ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 8
	ZhiBiaoShuChu_ZBHuiTu_TYPE_DRAWBMP      ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 9
	ZhiBiaoShuChu_ZBHuiTu_TYPE_VERTLINE     ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 10
	ZhiBiaoShuChu_ZBHuiTu_TYPE_TEXTABS      ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 11
	ZhiBiaoShuChu_ZBHuiTu_TYPE_TEXTREL      ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 12
	ZhiBiaoShuChu_ZBHuiTu_TYPE_RECTABS      ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 13
	ZhiBiaoShuChu_ZBHuiTu_TYPE_RECTREL      ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 14
	ZhiBiaoShuChu_ZBHuiTu_TYPE_FLAGTEXT     ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 15
	ZhiBiaoShuChu_ZBHuiTu_TYPE_MOVETEXT     ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 16
	ZhiBiaoShuChu_ZBHuiTu_TYPE_HORILINE     ZhiBiaoShuChu_ZBHuiTu_HTLeiXing = 17
)

var ZhiBiaoShuChu_ZBHuiTu_HTLeiXing_name = map[int32]string{
	0:  "TYPE_NOLINE",
	1:  "TYPE_POLYLINE",
	2:  "TYPE_LINE",
	3:  "TYPE_STICKLINE",
	4:  "TYPE_TEXT",
	5:  "TYPE_ICON",
	6:  "TYPE_TIP_TEXT",
	7:  "TYPE_BACK_GRD",
	8:  "TYPE_BACK_GRDLAST",
	9:  "TYPE_DRAWBMP",
	10: "TYPE_VERTLINE",
	11: "TYPE_TEXTABS",
	12: "TYPE_TEXTREL",
	13: "TYPE_RECTABS",
	14: "TYPE_RECTREL",
	15: "TYPE_FLAGTEXT",
	16: "TYPE_MOVETEXT",
	17: "TYPE_HORILINE",
}
var ZhiBiaoShuChu_ZBHuiTu_HTLeiXing_value = map[string]int32{
	"TYPE_NOLINE":       0,
	"TYPE_POLYLINE":     1,
	"TYPE_LINE":         2,
	"TYPE_STICKLINE":    3,
	"TYPE_TEXT":         4,
	"TYPE_ICON":         5,
	"TYPE_TIP_TEXT":     6,
	"TYPE_BACK_GRD":     7,
	"TYPE_BACK_GRDLAST": 8,
	"TYPE_DRAWBMP":      9,
	"TYPE_VERTLINE":     10,
	"TYPE_TEXTABS":      11,
	"TYPE_TEXTREL":      12,
	"TYPE_RECTABS":      13,
	"TYPE_RECTREL":      14,
	"TYPE_FLAGTEXT":     15,
	"TYPE_MOVETEXT":     16,
	"TYPE_HORILINE":     17,
}

func (x ZhiBiaoShuChu_ZBHuiTu_HTLeiXing) String() string {
	return proto.EnumName(ZhiBiaoShuChu_ZBHuiTu_HTLeiXing_name, int32(x))
}
func (ZhiBiaoShuChu_ZBHuiTu_HTLeiXing) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor14, []int{0, 2, 0}
}

// 指标计算输出
type ZhiBiaoShuChu struct {
	Obj     string                     `protobuf:"bytes,1,opt,name=Obj" json:"Obj,omitempty"`
	ShuJu   []*ZhiBiaoShuChu_ZBShuJu   `protobuf:"bytes,2,rep,name=ShuJu" json:"ShuJu,omitempty"`
	ShuXing []*ZhiBiaoShuChu_ZBShuXing `protobuf:"bytes,3,rep,name=ShuXing" json:"ShuXing,omitempty"`
	HuiTu   []*ZhiBiaoShuChu_ZBHuiTu   `protobuf:"bytes,4,rep,name=HuiTu" json:"HuiTu,omitempty"`
}

func (m *ZhiBiaoShuChu) Reset()                    { *m = ZhiBiaoShuChu{} }
func (m *ZhiBiaoShuChu) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiaoShuChu) ProtoMessage()               {}
func (*ZhiBiaoShuChu) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *ZhiBiaoShuChu) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *ZhiBiaoShuChu) GetShuJu() []*ZhiBiaoShuChu_ZBShuJu {
	if m != nil {
		return m.ShuJu
	}
	return nil
}

func (m *ZhiBiaoShuChu) GetShuXing() []*ZhiBiaoShuChu_ZBShuXing {
	if m != nil {
		return m.ShuXing
	}
	return nil
}

func (m *ZhiBiaoShuChu) GetHuiTu() []*ZhiBiaoShuChu_ZBHuiTu {
	if m != nil {
		return m.HuiTu
	}
	return nil
}

// 指标输出数据
type ZhiBiaoShuChu_ZBShuJu struct {
	ShiJian YFloat   `protobuf:"varint,1,opt,name=ShiJian" json:"ShiJian,omitempty"`
	JieGuo  []YFloat `protobuf:"varint,2,rep,packed,name=JieGuo" json:"JieGuo,omitempty"`
}

func (m *ZhiBiaoShuChu_ZBShuJu) Reset()                    { *m = ZhiBiaoShuChu_ZBShuJu{} }
func (m *ZhiBiaoShuChu_ZBShuJu) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiaoShuChu_ZBShuJu) ProtoMessage()               {}
func (*ZhiBiaoShuChu_ZBShuJu) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0, 0} }

func (m *ZhiBiaoShuChu_ZBShuJu) GetShiJian() YFloat {
	if m != nil {
		return m.ShiJian
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuJu) GetJieGuo() []YFloat {
	if m != nil {
		return m.JieGuo
	}
	return nil
}

// 指标输出属性
type ZhiBiaoShuChu_ZBShuXing struct {
	MingCheng       string                            `protobuf:"bytes,1,opt,name=MingCheng" json:"MingCheng,omitempty"`
	YanSe           YFloat                             `protobuf:"varint,2,opt,name=YanSe" json:"YanSe,omitempty"`
	LeiXing         ZhiBiaoShuChu_ZBShuXing_SXLeiXing `protobuf:"varint,3,opt,name=LeiXing,enum=dzhyun.ZhiBiaoShuChu_ZBShuXing_SXLeiXing" json:"LeiXing,omitempty"`
	KuanDu          YFloat                             `protobuf:"varint,4,opt,name=KuanDu" json:"KuanDu,omitempty"`
	JingDu          YFloat                             `protobuf:"varint,5,opt,name=JingDu" json:"JingDu,omitempty"`
	DuiQi           YFloat                             `protobuf:"varint,6,opt,name=DuiQi" json:"DuiQi,omitempty"`
	ShuXing         YFloat                             `protobuf:"varint,7,opt,name=ShuXing" json:"ShuXing,omitempty"`
	YiDong          YFloat                             `protobuf:"varint,8,opt,name=YiDong" json:"YiDong,omitempty"`
	CengCi          YFloat                             `protobuf:"varint,9,opt,name=CengCi" json:"CengCi,omitempty"`
	BianLiangWeiZhi YFloat                             `protobuf:"varint,10,opt,name=BianLiangWeiZhi" json:"BianLiangWeiZhi,omitempty"`
	KuoZhanShuXing  YFloat                             `protobuf:"varint,11,opt,name=KuoZhanShuXing" json:"KuoZhanShuXing,omitempty"`
	YouXiaoWeiZhi   YFloat                             `protobuf:"varint,12,opt,name=YouXiaoWeiZhi" json:"YouXiaoWeiZhi,omitempty"`
}

func (m *ZhiBiaoShuChu_ZBShuXing) Reset()                    { *m = ZhiBiaoShuChu_ZBShuXing{} }
func (m *ZhiBiaoShuChu_ZBShuXing) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiaoShuChu_ZBShuXing) ProtoMessage()               {}
func (*ZhiBiaoShuChu_ZBShuXing) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0, 1} }

func (m *ZhiBiaoShuChu_ZBShuXing) GetMingCheng() string {
	if m != nil {
		return m.MingCheng
	}
	return ""
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetYanSe() YFloat {
	if m != nil {
		return m.YanSe
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetLeiXing() ZhiBiaoShuChu_ZBShuXing_SXLeiXing {
	if m != nil {
		return m.LeiXing
	}
	return ZhiBiaoShuChu_ZBShuXing_TYPE_TEMP_EXPRESION
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetKuanDu() YFloat {
	if m != nil {
		return m.KuanDu
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetJingDu() YFloat {
	if m != nil {
		return m.JingDu
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetDuiQi() YFloat {
	if m != nil {
		return m.DuiQi
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetShuXing() YFloat {
	if m != nil {
		return m.ShuXing
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetYiDong() YFloat {
	if m != nil {
		return m.YiDong
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetCengCi() YFloat {
	if m != nil {
		return m.CengCi
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetBianLiangWeiZhi() YFloat {
	if m != nil {
		return m.BianLiangWeiZhi
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetKuoZhanShuXing() YFloat {
	if m != nil {
		return m.KuoZhanShuXing
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBShuXing) GetYouXiaoWeiZhi() YFloat {
	if m != nil {
		return m.YouXiaoWeiZhi
	}
	return 0
}

// 指标绘图指令输出
type ZhiBiaoShuChu_ZBHuiTu struct {
	LeiXing              ZhiBiaoShuChu_ZBHuiTu_HTLeiXing   `protobuf:"varint,1,opt,name=LeiXing,enum=dzhyun.ZhiBiaoShuChu_ZBHuiTu_HTLeiXing" json:"LeiXing,omitempty"`
	KuanDu               YFloat                             `protobuf:"varint,2,opt,name=KuanDu" json:"KuanDu,omitempty"`
	ShuXing              YFloat                             `protobuf:"varint,3,opt,name=ShuXing" json:"ShuXing,omitempty"`
	ShangCiJiSuan        YFloat                             `protobuf:"varint,4,opt,name=ShangCiJiSuan" json:"ShangCiJiSuan,omitempty"`
	YanSe                YFloat                             `protobuf:"varint,5,opt,name=YanSe" json:"YanSe,omitempty"`
	ShuChuLeiXing        ZhiBiaoShuChu_ZBShuXing_SXLeiXing `protobuf:"varint,6,opt,name=ShuChuLeiXing,enum=dzhyun.ZhiBiaoShuChu_ZBShuXing_SXLeiXing" json:"ShuChuLeiXing,omitempty"`
	ShuChuShuXing        YFloat                             `protobuf:"varint,7,opt,name=ShuChuShuXing" json:"ShuChuShuXing,omitempty"`
	ShuChuKuoZhanShuXing YFloat                             `protobuf:"varint,8,opt,name=ShuChuKuoZhanShuXing" json:"ShuChuKuoZhanShuXing,omitempty"`
	WenBen               []string                          `protobuf:"bytes,9,rep,name=WenBen" json:"WenBen,omitempty"`
	ShuJu                []*ZhiBiaoShuChu_ZBHuiTu_HTShuJu  `protobuf:"bytes,10,rep,name=ShuJu" json:"ShuJu,omitempty"`
}

func (m *ZhiBiaoShuChu_ZBHuiTu) Reset()                    { *m = ZhiBiaoShuChu_ZBHuiTu{} }
func (m *ZhiBiaoShuChu_ZBHuiTu) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiaoShuChu_ZBHuiTu) ProtoMessage()               {}
func (*ZhiBiaoShuChu_ZBHuiTu) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0, 2} }

func (m *ZhiBiaoShuChu_ZBHuiTu) GetLeiXing() ZhiBiaoShuChu_ZBHuiTu_HTLeiXing {
	if m != nil {
		return m.LeiXing
	}
	return ZhiBiaoShuChu_ZBHuiTu_TYPE_NOLINE
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetKuanDu() YFloat {
	if m != nil {
		return m.KuanDu
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetShuXing() YFloat {
	if m != nil {
		return m.ShuXing
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetShangCiJiSuan() YFloat {
	if m != nil {
		return m.ShangCiJiSuan
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetYanSe() YFloat {
	if m != nil {
		return m.YanSe
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetShuChuLeiXing() ZhiBiaoShuChu_ZBShuXing_SXLeiXing {
	if m != nil {
		return m.ShuChuLeiXing
	}
	return ZhiBiaoShuChu_ZBShuXing_TYPE_TEMP_EXPRESION
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetShuChuShuXing() YFloat {
	if m != nil {
		return m.ShuChuShuXing
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetShuChuKuoZhanShuXing() YFloat {
	if m != nil {
		return m.ShuChuKuoZhanShuXing
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetWenBen() []string {
	if m != nil {
		return m.WenBen
	}
	return nil
}

func (m *ZhiBiaoShuChu_ZBHuiTu) GetShuJu() []*ZhiBiaoShuChu_ZBHuiTu_HTShuJu {
	if m != nil {
		return m.ShuJu
	}
	return nil
}

// 绘图指令数据
type ZhiBiaoShuChu_ZBHuiTu_HTShuJu struct {
	WeiZhi YFloat `protobuf:"varint,1,opt,name=WeiZhi" json:"WeiZhi,omitempty"`
	JiaGe  YFloat `protobuf:"varint,2,opt,name=JiaGe" json:"JiaGe,omitempty"`
	CanShu YFloat `protobuf:"varint,3,opt,name=CanShu" json:"CanShu,omitempty"`
}

func (m *ZhiBiaoShuChu_ZBHuiTu_HTShuJu) Reset()         { *m = ZhiBiaoShuChu_ZBHuiTu_HTShuJu{} }
func (m *ZhiBiaoShuChu_ZBHuiTu_HTShuJu) String() string { return proto.CompactTextString(m) }
func (*ZhiBiaoShuChu_ZBHuiTu_HTShuJu) ProtoMessage()    {}
func (*ZhiBiaoShuChu_ZBHuiTu_HTShuJu) Descriptor() ([]byte, []int) {
	return fileDescriptor14, []int{0, 2, 0}
}

func (m *ZhiBiaoShuChu_ZBHuiTu_HTShuJu) GetWeiZhi() YFloat {
	if m != nil {
		return m.WeiZhi
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBHuiTu_HTShuJu) GetJiaGe() YFloat {
	if m != nil {
		return m.JiaGe
	}
	return 0
}

func (m *ZhiBiaoShuChu_ZBHuiTu_HTShuJu) GetCanShu() YFloat {
	if m != nil {
		return m.CanShu
	}
	return 0
}

// 指标属性输出
type ZhiBiaoShuXingShuChu struct {
	ShuChu []*ZhiBiaoShuChu_ZBShuXing `protobuf:"bytes,1,rep,name=ShuChu" json:"ShuChu,omitempty"`
}

func (m *ZhiBiaoShuXingShuChu) Reset()                    { *m = ZhiBiaoShuXingShuChu{} }
func (m *ZhiBiaoShuXingShuChu) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiaoShuXingShuChu) ProtoMessage()               {}
func (*ZhiBiaoShuXingShuChu) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *ZhiBiaoShuXingShuChu) GetShuChu() []*ZhiBiaoShuChu_ZBShuXing {
	if m != nil {
		return m.ShuChu
	}
	return nil
}

// 指标绘图指令输出
type ZhiBiaoHuiTuShuChu struct {
	ShuChu []*ZhiBiaoShuChu_ZBHuiTu `protobuf:"bytes,1,rep,name=ShuChu" json:"ShuChu,omitempty"`
}

func (m *ZhiBiaoHuiTuShuChu) Reset()                    { *m = ZhiBiaoHuiTuShuChu{} }
func (m *ZhiBiaoHuiTuShuChu) String() string            { return proto.CompactTextString(m) }
func (*ZhiBiaoHuiTuShuChu) ProtoMessage()               {}
func (*ZhiBiaoHuiTuShuChu) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *ZhiBiaoHuiTuShuChu) GetShuChu() []*ZhiBiaoShuChu_ZBHuiTu {
	if m != nil {
		return m.ShuChu
	}
	return nil
}

func init() {
	proto.RegisterType((*ZhiBiaoShuChu)(nil), "dzhyun.ZhiBiaoShuChu")
	proto.RegisterType((*ZhiBiaoShuChu_ZBShuJu)(nil), "dzhyun.ZhiBiaoShuChu.ZBShuJu")
	proto.RegisterType((*ZhiBiaoShuChu_ZBShuXing)(nil), "dzhyun.ZhiBiaoShuChu.ZBShuXing")
	proto.RegisterType((*ZhiBiaoShuChu_ZBHuiTu)(nil), "dzhyun.ZhiBiaoShuChu.ZBHuiTu")
	proto.RegisterType((*ZhiBiaoShuChu_ZBHuiTu_HTShuJu)(nil), "dzhyun.ZhiBiaoShuChu.ZBHuiTu.HTShuJu")
	proto.RegisterType((*ZhiBiaoShuXingShuChu)(nil), "dzhyun.ZhiBiaoShuXingShuChu")
	proto.RegisterType((*ZhiBiaoHuiTuShuChu)(nil), "dzhyun.ZhiBiaoHuiTuShuChu")
	proto.RegisterEnum("dzhyun.ZhiBiaoShuChu_ZBShuXing_SXLeiXing", ZhiBiaoShuChu_ZBShuXing_SXLeiXing_name, ZhiBiaoShuChu_ZBShuXing_SXLeiXing_value)
	proto.RegisterEnum("dzhyun.ZhiBiaoShuChu_ZBHuiTu_HTLeiXing", ZhiBiaoShuChu_ZBHuiTu_HTLeiXing_name, ZhiBiaoShuChu_ZBHuiTu_HTLeiXing_value)
}

func init() { proto.RegisterFile("dzhyun.zhibiaojisuan.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xd1, 0x92, 0xe2, 0x44,
	0x14, 0x5d, 0x60, 0x09, 0x9b, 0xcb, 0x30, 0xf4, 0xf4, 0xe2, 0x9a, 0xa2, 0xb4, 0x9c, 0xa2, 0x56,
	0xc5, 0x17, 0x1e, 0x66, 0xca, 0xb2, 0xac, 0x7d, 0x82, 0x10, 0x21, 0x90, 0xa1, 0xb1, 0x13, 0x67,
	0x60, 0x5e, 0xa8, 0x8c, 0x1b, 0x49, 0x6f, 0x95, 0x89, 0xa5, 0xe6, 0xc1, 0xfd, 0x05, 0xdf, 0xfc,
	0x03, 0xbf, 0xc9, 0x2f, 0xf0, 0x4f, 0xac, 0xbe, 0xdd, 0x69, 0x08, 0x65, 0xcd, 0xe8, 0x13, 0xb9,
	0xa7, 0xcf, 0xbd, 0xf7, 0x74, 0xe7, 0x9e, 0x0e, 0xd0, 0x7f, 0xfb, 0x3e, 0xfd, 0xad, 0xc8, 0x46,
	0xef, 0x53, 0xf1, 0x20, 0xe2, 0xfc, 0x9d, 0xf8, 0xa5, 0x88, 0xb3, 0xd1, 0x4f, 0x3f, 0xe7, 0xbf,
	0xe6, 0xd4, 0x52, 0x6b, 0x83, 0xdf, 0xbb, 0xd0, 0xb9, 0x4f, 0xc5, 0x44, 0xc4, 0x79, 0x98, 0x16,
	0x6e, 0x5a, 0x50, 0x02, 0x0d, 0xf6, 0xf0, 0xce, 0xa9, 0x5d, 0xd6, 0x86, 0x36, 0x97, 0x8f, 0xf4,
	0x1a, 0x9a, 0x61, 0x5a, 0x2c, 0x0a, 0xa7, 0x7e, 0xd9, 0x18, 0xb6, 0xaf, 0x3e, 0x1e, 0xe9, 0xba,
	0x95, 0xbc, 0xd1, 0xfd, 0x04, 0x49, 0x5c, 0x71, 0xe9, 0xd7, 0xd0, 0x0a, 0xd3, 0x62, 0x23, 0xb2,
	0xbd, 0xd3, 0xc0, 0xb4, 0x4f, 0x1e, 0x49, 0x93, 0x34, 0x5e, 0xf2, 0x65, 0xbf, 0x79, 0x21, 0xa2,
	0xc2, 0x79, 0xfe, 0x78, 0x3f, 0x24, 0x71, 0xc5, 0xed, 0xbf, 0x81, 0x96, 0x56, 0x40, 0x1d, 0xd9,
	0x5a, 0x2c, 0x44, 0x9c, 0xe1, 0x2e, 0x1a, 0xbc, 0x0c, 0xe9, 0x2b, 0xb0, 0x16, 0x22, 0x99, 0x15,
	0x39, 0x6e, 0xa5, 0xc1, 0x75, 0xd4, 0xff, 0xd3, 0x02, 0xdb, 0x08, 0xa1, 0x1f, 0x81, 0x7d, 0x23,
	0xb2, 0xbd, 0x9b, 0x26, 0xd9, 0x5e, 0x9f, 0xc3, 0x01, 0xa0, 0x3d, 0x68, 0x6e, 0xe3, 0x2c, 0x4c,
	0x9c, 0x3a, 0xd6, 0x56, 0x01, 0x75, 0xa1, 0x15, 0x24, 0x42, 0x6f, 0xb7, 0x36, 0x3c, 0xbf, 0xfa,
	0xe2, 0x89, 0xed, 0x8e, 0xc2, 0x8d, 0x4e, 0xe0, 0x65, 0xa6, 0x94, 0xb7, 0x2c, 0xe2, 0x6c, 0x2a,
	0x77, 0x2e, 0x6b, 0xeb, 0x48, 0xc9, 0xce, 0xf6, 0xd3, 0xc2, 0x69, 0x2a, 0x5c, 0x45, 0x52, 0xca,
	0xb4, 0x10, 0xdf, 0x0a, 0xc7, 0x52, 0x52, 0x30, 0x50, 0xdb, 0x57, 0x27, 0xdf, 0x2a, 0xb7, 0x5f,
	0x94, 0xf5, 0xb7, 0x62, 0x9a, 0x67, 0x7b, 0xe7, 0x85, 0xaa, 0xa3, 0x22, 0x89, 0xbb, 0x49, 0xb6,
	0x77, 0x85, 0x63, 0x2b, 0x5c, 0x45, 0x74, 0x08, 0xdd, 0x89, 0x88, 0xb3, 0x40, 0xc4, 0xd9, 0xfe,
	0x2e, 0x11, 0xf7, 0xa9, 0x70, 0x00, 0x09, 0xa7, 0x30, 0xfd, 0x0c, 0xce, 0x97, 0x45, 0x7e, 0x9f,
	0xc6, 0x59, 0xd9, 0xba, 0x8d, 0xc4, 0x13, 0x94, 0xbe, 0x86, 0xce, 0x36, 0x2f, 0x36, 0x22, 0xce,
	0x75, 0xbd, 0x33, 0xa4, 0x55, 0xc1, 0xc1, 0x1f, 0x0d, 0xb0, 0xcd, 0xf1, 0xd0, 0x0f, 0xe1, 0x65,
	0xb4, 0x5d, 0x7b, 0xbb, 0xc8, 0xbb, 0x59, 0xef, 0xbc, 0xcd, 0x9a, 0x7b, 0xa1, 0xcf, 0x56, 0xe4,
	0x19, 0xa5, 0x70, 0x8e, 0x0b, 0xee, 0x77, 0xfc, 0x76, 0x17, 0xf8, 0x2b, 0x8f, 0xd4, 0xe8, 0x4b,
	0xe8, 0x22, 0x16, 0x46, 0xbe, 0xbb, 0x54, 0x60, 0x9d, 0x3a, 0xd0, 0x53, 0x44, 0x16, 0x30, 0x7e,
	0xb4, 0xd2, 0xa0, 0xaf, 0x80, 0xe2, 0xca, 0x2d, 0x0b, 0x8e, 0xf0, 0xe7, 0xa6, 0xa7, 0x0c, 0x8f,
	0x16, 0x9a, 0x87, 0x9e, 0x9c, 0x85, 0xe1, 0x6e, 0xca, 0x22, 0x62, 0x99, 0x9e, 0xae, 0xcf, 0xdd,
	0xc0, 0x43, 0xb0, 0x65, 0x88, 0x6b, 0xe6, 0xaf, 0x22, 0xc4, 0x5e, 0xd0, 0x0f, 0xe0, 0xe2, 0x20,
	0xee, 0x7a, 0xaa, 0x6a, 0xda, 0x06, 0x46, 0x79, 0x25, 0x0c, 0x94, 0xc0, 0x19, 0xc2, 0x53, 0xa6,
	0xf2, 0xdb, 0xf4, 0x02, 0x3a, 0x0a, 0x19, 0x87, 0x73, 0x84, 0xce, 0x68, 0x0f, 0x88, 0x6a, 0xe3,
	0x71, 0xd7, 0x5b, 0x45, 0xbb, 0xc9, 0x98, 0x93, 0x8e, 0x51, 0xe4, 0xad, 0x22, 0x8f, 0xef, 0x02,
	0xb6, 0x9a, 0x91, 0xb7, 0x46, 0x91, 0xb7, 0xf1, 0x23, 0x85, 0x25, 0x26, 0x5d, 0x11, 0xc3, 0x39,
	0xe3, 0x11, 0xf9, 0xe1, 0x90, 0x2e, 0x99, 0x0a, 0xdc, 0xf7, 0xff, 0xb6, 0xa4, 0xc3, 0xd0, 0x6c,
	0x74, 0x7c, 0x98, 0xf6, 0x1a, 0x4e, 0xfb, 0xe7, 0x8f, 0x7a, 0x74, 0x34, 0x8f, 0x1e, 0x99, 0xf5,
	0x7a, 0x65, 0xd6, 0x9d, 0xe3, 0x7b, 0xa3, 0x32, 0xbd, 0xaf, 0xa1, 0x13, 0xa6, 0xb1, 0x1c, 0xcc,
	0x85, 0x08, 0x8b, 0x38, 0xd3, 0x26, 0xa9, 0x82, 0x07, 0x7b, 0x36, 0x8f, 0xed, 0xc9, 0x64, 0xae,
	0x94, 0x54, 0xca, 0xb6, 0xfe, 0xaf, 0x49, 0xab, 0xf9, 0x4a, 0x8c, 0x04, 0xaa, 0x56, 0xab, 0x82,
	0xf4, 0x0a, 0x7a, 0x0a, 0x38, 0x31, 0x87, 0xb2, 0xdf, 0xbf, 0xae, 0xc9, 0x83, 0xb9, 0x4b, 0xb2,
	0x49, 0x92, 0x39, 0xf6, 0x65, 0x63, 0x68, 0x73, 0x1d, 0xd1, 0x37, 0xe5, 0x2d, 0x0c, 0x78, 0x2b,
	0x7e, 0xfa, 0xd4, 0x89, 0x1f, 0xdf, 0xc6, 0x7d, 0x06, 0x2d, 0x8d, 0xa8, 0xfa, 0xe8, 0x3d, 0x75,
	0x39, 0xea, 0x48, 0x1e, 0xdc, 0x42, 0xc4, 0x33, 0x73, 0xaf, 0x61, 0x80, 0x57, 0x03, 0x6a, 0xd3,
	0x6f, 0x43, 0x47, 0x83, 0xbf, 0xea, 0x60, 0x9b, 0xb7, 0x4a, 0xbb, 0xd0, 0xc6, 0x81, 0x59, 0x31,
	0x9c, 0xdd, 0x67, 0x66, 0x52, 0xd7, 0x2c, 0xd8, 0x6a, 0x67, 0x76, 0xc0, 0x36, 0x96, 0x22, 0x75,
	0x33, 0x8d, 0xe8, 0x05, 0xed, 0xc6, 0x92, 0x12, 0x79, 0x9b, 0x88, 0x3c, 0x37, 0xa1, 0xef, 0xb2,
	0x15, 0x69, 0x9a, 0x9a, 0x91, 0xbf, 0x56, 0x0c, 0xcb, 0x40, 0x93, 0xb1, 0xbb, 0xdc, 0xcd, 0xf8,
	0x94, 0xb4, 0x8c, 0x99, 0x4a, 0x28, 0x18, 0x87, 0xd2, 0x7a, 0xc6, 0x4c, 0x7c, 0x7c, 0x37, 0xb9,
	0x59, 0x13, 0xdb, 0xe4, 0xde, 0x7a, 0x3c, 0x3a, 0x71, 0x9c, 0xac, 0x3e, 0x9e, 0x84, 0xa4, 0x5d,
	0x41, 0xb8, 0x17, 0x90, 0x33, 0x83, 0x70, 0xcf, 0x45, 0x4e, 0xa7, 0x82, 0x48, 0xce, 0xb9, 0x29,
	0xfd, 0x4d, 0x30, 0x9e, 0xa1, 0xd2, 0xae, 0x81, 0x6e, 0xd8, 0xad, 0x87, 0x10, 0x31, 0xd0, 0x9c,
	0x71, 0x1f, 0x05, 0x5c, 0x0c, 0x18, 0xf4, 0x0e, 0xaf, 0x53, 0x9e, 0xac, 0xfe, 0x26, 0x7f, 0x05,
	0x96, 0x7a, 0x72, 0x6a, 0xff, 0xed, 0x5b, 0xaa, 0xe9, 0x83, 0x25, 0x50, 0x4d, 0xc1, 0xb1, 0xd0,
	0xe5, 0xbe, 0x3c, 0x29, 0xf7, 0xc4, 0x17, 0x56, 0x93, 0x27, 0x14, 0xc8, 0xf7, 0xf9, 0x8f, 0x25,
	0x17, 0xff, 0x47, 0x3c, 0x58, 0xf8, 0x73, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x3d,
	0x01, 0x36, 0x6c, 0x08, 0x00, 0x00,
}
