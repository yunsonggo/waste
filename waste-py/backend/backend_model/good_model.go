package backend_model

type Good struct {
	Id          int64
	Name        string  `xorm:"varchar(30)" json:"name"`        // 商品名称
	Description string  `xorm:"varchar(60)" json:"description"` // 商品描述
	Icon        string  `xorm:"varchar(255)" json:"icon"`       // 商品图标
	Price       float32 `xorm:"float" json:"price"`             // 现售单价
	OldPrice    float32 `xorm:"float" json:"old_price"`         // 原价
	GoodsUnit   string  `xorm:"varchar(30)" json:"goods_unit"`  // 商品单位
	IsHot       int     `xorm:"int" json:"is_hot"`              // 热卖 true 1 false 0
	SellCount   int64   `xorm:"int" json:"sell_count"`          // 已售数量
	Station     int     `xorm:"int" json:"station"`             // 排序位置
	Tage        string  `xorm:"varchar(10)" json:"tage"`        // 图标
	Centent     string  `xorm:"text" json:"content"`
}
