package model

import (
	"time"
)

//订单表。每下一个单 单独一个订单。订单状态跟桌子状态相关。
type OrderInfoTbl struct {
	ID          int       `gorm:"primary_key" json:"-"` //
	OrderID     string    `json:"order_id"`             //	订单号
	OrderItemID string    `json:"order_item_id"`        //	子订单号
	TableID     string    `json:"table_id"`             //	桌子号
	GoodsID     string    `json:"goods_id"`             //	菜品ID
	GoodsName   string    `json:"goods_name"`           //	菜品名
	Count       int       `json:"count"`                //	购买数量
	UnitPrice   int       `json:"unit_price"`           //	单价（金额分）
	Status      int       `json:"status"`               //	商品状态（-1:已取消, 1:已创建，2:已出单，3:已支付）
	Did         string    `json:"did"`                  //	下单设备id
	CreatTime   time.Time `json:"creat_time"`           //	创建时间
}

//支付列表
type PayTbl struct {
	ID         int       `gorm:"primary_key" json:"-"`                      //
	TableID    string    `gorm:"unique_index:table_id" json:"table_id"`     //	桌子号
	PayOrderID string    `gorm:"unique_index:table_id" json:"pay_order_id"` //	订单号
	Amount     int       `json:"amount"`                                    //	消费金额
	Tax        int       `json:"tax"`                                       //	交税金额
	Tip        int       `json:"tip"`                                       //	小费
	Method     int       `json:"method"`                                    //	支付方式(1:现金，2:信用卡，3:其他)
	CreatTime  time.Time `json:"creat_time"`                                //	支付时间
}

//统计分析类型
type AnalyticsTbl struct {
	ID     int    `gorm:"primary_key" json:"-"`         //
	Topic  string `gorm:"index:idx_name" json:"topic"`  //	主题
	Bundle string `gorm:"index:idx_name" json:"bundle"` //	包
	Pid    string `json:"pid"`                          //	类别
	Data   int64  `json:"data"`                         //	数据
	Attach int64  `json:"attach"`                       //	附加数据
	Hour0  int64  `gorm:"index:idx_name" json:"hour0"`  //	整点时间
	Day0   int64  `json:"day0"`                         //	当天的0点
	Week0  int64  `json:"week0"`                        //	当周周天零点
	Month0 int64  `json:"month0"`                       //	当月1号零点
}

//操作日志表
type OpLogTbl struct {
	ID         int       `gorm:"primary_key" json:"-"` //
	Operator   string    `json:"operator"`             //	操作人
	Receive    string    `json:"receive"`              //	接收方
	CreateTime time.Time `json:"create_time"`          //	操作时间
	Topic      string    `json:"topic"`                //	操作主题
	Bundle     string    `json:"bundle"`               //	子主体
	Pid        string    `json:"pid"`                  //	细分主体
	IPAddr     string    `json:"ip_addr"`              //	操作IP地址
	Num0       int       `json:"num0"`                 //	附加数字1
	Num1       int       `json:"num1"`                 //	附加数字2
	Num2       int       `json:"num2"`                 //	附加数字3
	Attach0    string    `json:"attach0"`              //	附加字符串1
	Attach1    string    `json:"attach1"`              //	附加字符串2
	Attach2    string    `json:"attach2"`              //	附加字符串3
}

//商品类别排序表
type GtypeRankTbl struct {
	ID       int    `gorm:"primary_key" json:"-"` //
	RankInfo string `json:"rank_info"`            //
}

type PayGoodsTbl struct {
	ID         int       `gorm:"primary_key" json:"-"`      //
	PayOrderID string    `gorm:"index" json:"pay_order_id"` //	支付单号
	OrderID    string    `gorm:"index" json:"order_id"`     //	订单号
	TableID    string    `json:"table_id"`                  //	桌子号
	GoodsID    string    `json:"goods_id"`                  //	菜品ID
	GoodsName  string    `json:"goods_name"`                //	菜品名
	Nume       int       `json:"nume"`                      //	分子
	Deno       int       `json:"deno"`                      //	分母
	UnitPrice  int       `json:"unit_price"`                //	单价（金额分）
	Status     int       `json:"status"`                    //	商品状态（-1:已取消, 1:已创建，2:已出单，3:已支付）
	CreatTime  time.Time `json:"creat_time"`                //	创建时间
}

//翻译kv数据表
type TranslTbl struct {
	ID    int    `gorm:"primary_key" json:"-"` //
	Key   string `gorm:"unique" json:"key"`    //	key(中文)
	Value string `json:"value"`                //	值(英文)
}

//商品信息表
type GoodsInfoTbl struct {
	ID         int       `gorm:"primary_key" json:"-"`                    //
	GoodsID    string    `gorm:"unique" json:"goods_id"`                  //	商品编号
	GoodsName  string    `gorm:"index:goods_id" json:"goods_name"`        //	商品名称
	TypeName   string    `gorm:"index:goods_id_4;index" json:"type_name"` //	商品类别
	BasePrice  int       `json:"base_price"`                              //	价格（单位：分）
	IsVaild    int       `json:"is_vaild"`                                //	是否有效（1：有效/正常  2：无效/下架）
	CreateTime time.Time `json:"create_time"`                             //	创建时间
	Stock      int       `json:"stock"`                                   //	库存（-1：无限库存 ）
	PicURL     string    `json:"pic_url"`                                 //	轮播图url
}

//商品类别表
//111
type GoodsTypeTbl struct {
	ID   int    `gorm:"primary_key" json:"-"` //
	Name string `gorm:"unique" json:"name"`   //	类别123
}
