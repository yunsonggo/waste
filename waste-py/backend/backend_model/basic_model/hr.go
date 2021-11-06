package basic_model

type Hr2020 struct {
	Id          int64
	BearDate    string  `xorm:"varchar(20)" json:"bear_date"` // 载明日期
	UserName    string  `xorm:"varchar(20)" json:"user_name"` // 姓名
	UserId      string  `xorm:"varchar(20)" json:"user_id"`   // 身份证
	UserStation string `xorm:"varchar(30)" json:"user_station"` // 部门
	UserZhiwei string `xorm:"varchar(30)" json:"user_zhiwei"` // 职位
	UserDate string `xorm:"varchar(30)" json:"user_date"` // 入职时间
	UserType string `xorm:"varchar(10)" json:"user_type"` // 农民
	BasicYanglao float64 `xorm:"float" json:"basic_yanglao"` // 养老基数
	BasicShiye float64 `xorm:"float" json:"basic_shiye"` // 失业基数
	BasicGongshang float64 `xorm:"float" json:"basic_gongshang"` // 工伤基数
	BasicYiliao	float64 `xorm:"float" json:"basic_yiliao"` // 医疗基数
	BasicShengyu float64 `xorm:"float" json:"basic_shengyu"`
	BasicGjj float64 `xorm:"float" json:"basic_gjj"`
	PayYanglaoCom float64 `xorm:"float" json:"pay_yanglao_com"`
	PayYanglaoPer float64 `xorm:"float" json:"pay_yanglao_per"`
	PayShiyeCom float64 `xorm:"float" json:"pay_shiye_com"`
	PayShiyePer float64 `xorm:"float" json:"pay_shiye_per"`
	PayGongshangCom float64 `xorm:"float" json:"pay_gongshang_com"`
	PayYiliaoCom float64 `xorm:"float" json:"pay_yiliao_com"`
	PayYiliaoPer float64 `xorm:"float" json:"pay_yiliao_per"`
	PayShengyuCom float64 `xorm:"float" json:"pay_shengyu_com"`
	PayGjjCom float64 `xorm:"float" json:"pay_gjj_com"`
	PayGjjPer float64 `xorm:"float" json:"pay_gjj_per"`
	PayTotalCom float64 `xorm:"float" json:"pay_total_com"`
	PayTotalPer float64 `xorm:"float" json:"pay_total_per"`
	PayTotal float64 `xorm:"float" json:"pay_total"`
}
