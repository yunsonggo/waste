package basic_server

import (
	"waste-py/backend/backend_model/basic_model"
	"waste-py/dal"
)

type HrServer interface {
	// 插入表格多条数据
	InsertHr2020(hrArr []basic_model.Hr2020) (err error)
	// 从数据库读取所有数据
	QueryHrDataAll() (hrArr []basic_model.Hr2020,err error)
}

type HrService struct {}

func NewHrService() HrServer {
	return &HrService{}
}

// 插入表格多条数据
func(hrs *HrService) InsertHr2020(hrArr []basic_model.Hr2020) (err error) {
	_,err = dal.DB.Insert(hrArr)
	return
}
// 从数据库读取所有数据
func(hrs *HrService) QueryHrDataAll() (hrArr []basic_model.Hr2020,err error) {
	err = dal.DB.Find(&hrArr)
	return
}