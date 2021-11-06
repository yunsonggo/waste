package model

type Cities struct {
	Id       int64
	CityId   int64  `xorm:"int" json:"city_id"`
	CityName string `xorm:"varchar(20)" json:"city_name"`
}

type Area struct {
	Id       int64
	CitiesId int64  `xorm:"int" json:"cities_id"`
	AreaName string `xorm:"varchar(255)" json:"area_name"`
}

type Position struct {
	Id              int64
	AreaId          int64  `xorm:"int" json:"area_id"`
	PositionName    string `xorm:"varchar(255)" json:"position_name"`
	PositionManager string `xorm:"varchar(30)" json:"position_manager"`
	PositionMobile  string `xorm:"varchar(11)" json:"position_mobile"`
}
