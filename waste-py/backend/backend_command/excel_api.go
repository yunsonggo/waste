package backend_command

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"waste-py/backend/backend_model/basic_model"
	"waste-py/backend/backend_server/basic_server"
	"waste-py/common"
	"waste-py/tools"
)

var hrs = basic_server.NewHrService()
var (
	// 1 月 | 1 农 0 非
	yanglaoCom10 float64 = 0.16
	yanglaoPer10 float64 = 0.08
	shiyeCom10 float64 = 0.007
	shiyePer10 float64 = 0.003
	gongshangCom10 float64 = 0.002
	yiliaoCom10 float64 = 0.07
	yiliaoPer10 float64 = 0.02
	shengyuCom10 float64 = 0.0
	gjjCom10 float64 = 0.06
	gjjPer10 float64 = 0.06

	yanglaoCom11 float64 = 0.16
	yanglaoPer11 float64 = 0.08
	shiyeCom11 float64 = 0.007
	shiyePer11 float64 = 0.0
	gongshangCom11 float64 = 0.002
	yiliaoCom11 float64 = 0.02
	yiliaoPer11 float64 = 0.0
	shengyuCom11 float64 = 0.0
	gjjCom11 float64 = 0.06
	gjjPer11 float64 = 0.06

	yanglaoCom20 float64 = 0.16
	yanglaoPer20 float64 = 0.08
	shiyeCom20 float64 = 0.007
	shiyePer20 float64 = 0.003
	gongshangCom20 float64 = 0.002
	yiliaoCom20 float64 = 0.07
	yiliaoPer20 float64 = 0.02
	shengyuCom20 float64 = 0.0
	gjjCom20 float64 = 0.06
	gjjPer20 float64 = 0.06

	yanglaoCom21 float64 = 0.16
	yanglaoPer21 float64 = 0.08
	shiyeCom21 float64 = 0.007
	shiyePer21 float64 = 0.0
	gongshangCom21 float64 = 0.002
	yiliaoCom21 float64 = 0.02
	yiliaoPer21 float64 = 0.0
	shengyuCom21 float64 = 0.0
	gjjCom21 float64 = 0.06
	gjjPer21 float64 = 0.06

	yanglaoCom30 float64 = 0.16
	yanglaoPer30 float64 = 0.08
	shiyeCom30 float64 = 0.007
	shiyePer30 float64 = 0.003
	gongshangCom30 float64 = 0.002
	yiliaoCom30 float64 = 0.07
	yiliaoPer30 float64 = 0.02
	shengyuCom30 float64 = 0.0
	gjjCom30 float64 = 0.06
	gjjPer30 float64 = 0.06

	yanglaoCom31 float64 = 0.16
	yanglaoPer31 float64 = 0.08
	shiyeCom31 float64 = 0.007
	shiyePer31 float64 = 0.0
	gongshangCom31 float64 = 0.002
	yiliaoCom31 float64 = 0.02
	yiliaoPer31 float64 = 0.0
	shengyuCom31 float64 = 0.0
	gjjCom31 float64 = 0.06
	gjjPer31 float64 = 0.06

	yanglaoCom40 float64 = 0.16
	yanglaoPer40 float64 = 0.08
	shiyeCom40 float64 = 0.007
	shiyePer40 float64 = 0.003
	gongshangCom40 float64 = 0.002
	yiliaoCom40 float64 = 0.07
	yiliaoPer40 float64 = 0.02
	shengyuCom40 float64 = 0.0
	gjjCom40 float64 = 0.06
	gjjPer40 float64 = 0.06

	yanglaoCom41 float64 = 0.16
	yanglaoPer41 float64 = 0.08
	shiyeCom41 float64 = 0.007
	shiyePer41 float64 = 0.0
	gongshangCom41 float64 = 0.002
	yiliaoCom41 float64 = 0.02
	yiliaoPer41 float64 = 0.0
	shengyuCom41 float64 = 0.0
	gjjCom41 float64 = 0.06
	gjjPer41 float64 = 0.06

	yanglaoCom50 float64 = 0.16
	yanglaoPer50 float64 = 0.08
	shiyeCom50 float64 = 0.007
	shiyePer50 float64 = 0.003
	gongshangCom50 float64 = 0.002
	yiliaoCom50 float64 = 0.07
	yiliaoPer50 float64 = 0.02
	shengyuCom50 float64 = 0.0
	gjjCom50 float64 = 0.06
	gjjPer50 float64 = 0.06

	yanglaoCom51 float64 = 0.16
	yanglaoPer51 float64 = 0.08
	shiyeCom51 float64 = 0.007
	shiyePer51 float64 = 0.0
	gongshangCom51 float64 = 0.002
	yiliaoCom51 float64 = 0.02
	yiliaoPer51 float64 = 0.0
	shengyuCom51 float64 = 0.0
	gjjCom51 float64 = 0.06
	gjjPer51 float64 = 0.06

	yanglaoCom60 float64 = 0.16
	yanglaoPer60 float64 = 0.08
	shiyeCom60 float64 = 0.007
	shiyePer60 float64 = 0.003
	gongshangCom60 float64 = 0.002
	yiliaoCom60 float64 = 0.07
	yiliaoPer60 float64 = 0.02
	shengyuCom60 float64 = 0.0
	gjjCom60 float64 = 0.06
	gjjPer60 float64 = 0.06

	yanglaoCom61 float64 = 0.16
	yanglaoPer61 float64 = 0.08
	shiyeCom61 float64 = 0.007
	shiyePer61 float64 = 0.0
	gongshangCom61 float64 = 0.002
	yiliaoCom61 float64 = 0.02
	yiliaoPer61 float64 = 0.0
	shengyuCom61 float64 = 0.0
	gjjCom61 float64 = 0.06
	gjjPer61 float64 = 0.06

	yanglaoCom70 float64 = 0.16
	yanglaoPer70 float64 = 0.08
	shiyeCom70 float64 = 0.007
	shiyePer70 float64 = 0.003
	gongshangCom70 float64 = 0.002
	yiliaoCom70 float64 = 0.07
	yiliaoPer70 float64 = 0.02
	shengyuCom70 float64 = 0.0
	gjjCom70 float64 = 0.06
	gjjPer70 float64 = 0.06

	yanglaoCom71 float64 = 0.16
	yanglaoPer71 float64 = 0.08
	shiyeCom71 float64 = 0.007
	shiyePer71 float64 = 0.0
	gongshangCom71 float64 = 0.002
	yiliaoCom71 float64 = 0.02
	yiliaoPer71 float64 = 0.0
	shengyuCom71 float64 = 0.0
	gjjCom71 float64 = 0.06
	gjjPer71 float64 = 0.06

	yanglaoCom80 float64 = 0.16
	yanglaoPer80 float64 = 0.08
	shiyeCom80 float64 = 0.007
	shiyePer80 float64 = 0.003
	gongshangCom80 float64 = 0.002
	yiliaoCom80 float64 = 0.07
	yiliaoPer80 float64 = 0.02
	shengyuCom80 float64 = 0.0
	gjjCom80 float64 = 0.06
	gjjPer80 float64 = 0.06

	yanglaoCom81 float64 = 0.16
	yanglaoPer81 float64 = 0.08
	shiyeCom81 float64 = 0.007
	shiyePer81 float64 = 0.0
	gongshangCom81 float64 = 0.002
	yiliaoCom81 float64 = 0.02
	yiliaoPer81 float64 = 0.0
	shengyuCom81 float64 = 0.0
	gjjCom81 float64 = 0.06
	gjjPer81 float64 = 0.06

	yanglaoCom90 float64 = 0.16
	yanglaoPer90 float64 = 0.08
	shiyeCom90 float64 = 0.007
	shiyePer90 float64 = 0.003
	gongshangCom90 float64 = 0.002
	yiliaoCom90 float64 = 0.07
	yiliaoPer90 float64 = 0.02
	shengyuCom90 float64 = 0.0
	gjjCom90 float64 = 0.06
	gjjPer90 float64 = 0.06

	yanglaoCom91 float64 = 0.16
	yanglaoPer91 float64 = 0.08
	shiyeCom91 float64 = 0.007
	shiyePer91 float64 = 0.0
	gongshangCom91 float64 = 0.002
	yiliaoCom91 float64 = 0.02
	yiliaoPer91 float64 = 0.0
	shengyuCom91 float64 = 0.0
	gjjCom91 float64 = 0.06
	gjjPer91 float64 = 0.06

	yanglaoCom100 float64 = 0.16
	yanglaoPer100 float64 = 0.08
	shiyeCom100 float64 = 0.007
	shiyePer100 float64 = 0.003
	gongshangCom100 float64 = 0.002
	yiliaoCom100 float64 = 0.07
	yiliaoPer100 float64 = 0.02
	shengyuCom100 float64 = 0.0
	gjjCom100 float64 = 0.06
	gjjPer100 float64 = 0.06

	yanglaoCom101 float64 = 0.16
	yanglaoPer101 float64 = 0.08
	shiyeCom101 float64 = 0.007
	shiyePer101 float64 = 0.0
	gongshangCom101 float64 = 0.002
	yiliaoCom101 float64 = 0.02
	yiliaoPer101 float64 = 0.0
	shengyuCom101 float64 = 0.0
	gjjCom101 float64 = 0.06
	gjjPer101 float64 = 0.06

	yanglaoCom110 float64 = 0.16
	yanglaoPer110 float64 = 0.08
	shiyeCom110 float64 = 0.007
	shiyePer110 float64 = 0.003
	gongshangCom110 float64 = 0.002
	yiliaoCom110 float64 = 0.07
	yiliaoPer110 float64 = 0.02
	shengyuCom110 float64 = 0.0
	gjjCom110 float64 = 0.06
	gjjPer110 float64 = 0.06

	yanglaoCom111 float64 = 0.16
	yanglaoPer111 float64 = 0.08
	shiyeCom111 float64 = 0.007
	shiyePer111 float64 = 0.0
	gongshangCom111 float64 = 0.002
	yiliaoCom111 float64 = 0.02
	yiliaoPer111 float64 = 0.0
	shengyuCom111 float64 = 0.0
	gjjCom111 float64 = 0.06
	gjjPer111 float64 = 0.06

	yanglaoCom120 float64 = 0.16
	yanglaoPer120 float64 = 0.08
	shiyeCom120 float64 = 0.007
	shiyePer120 float64 = 0.003
	gongshangCom120 float64 = 0.002
	yiliaoCom120 float64 = 0.07
	yiliaoPer120 float64 = 0.02
	shengyuCom120 float64 = 0.0
	gjjCom120 float64 = 0.06
	gjjPer120 float64 = 0.06

	yanglaoCom121 float64 = 0.16
	yanglaoPer121 float64 = 0.08
	shiyeCom121 float64 = 0.007
	shiyePer121 float64 = 0.0
	gongshangCom121 float64 = 0.002
	yiliaoCom121 float64 = 0.02
	yiliaoPer121 float64 = 0.0
	shengyuCom121 float64 = 0.0
	gjjCom121 float64 = 0.06
	gjjPer121 float64 = 0.06
)
func UploadExcel(ctx *gin.Context) {
	// 获取数据文件
	_, excelFileHeader, err := ctx.Request.FormFile("excel_file")
	if excelFileHeader != nil {
		if err != nil {
			common.Failed(ctx, "获取文件参数失败")
			return
		}
		fileName := excelFileHeader.Filename
		filePath := "public/upload/excel/" + fileName
		// 文件已存在则删除旧文件
		if tools.IsFIleExist(filePath) {
			_ = os.Remove(filePath)
		}
		// 文件保存
		saveFileErr := ctx.SaveUploadedFile(excelFileHeader, filePath)
		if saveFileErr != nil {
			common.Failed(ctx, "保存文件失败")
			return
		}
		// 打开文件
		excelFile, openExcelErr := excelize.OpenFile(filePath)
		if openExcelErr != nil {
			common.Failed(ctx, "打开服务器EXCEL文件失败")
			return
		}
		// 处理文件数据
		rows, _ := excelFile.GetRows("Sheet" + "1")
		var hrArr []basic_model.Hr2020
		for key, row := range rows {
			if key > 0 {
				hr := new(basic_model.Hr2020)
				hr.BearDate = row[0]
				hr.UserName = row[1]
				hr.UserId = row[2]
				hr.UserStation = row[3]
				hr.UserZhiwei = row[4]
				hr.UserDate = row[5]
				hr.UserType = row[6]
				BasicYanglao,_ := strconv.ParseFloat(row[7],64)
				hr.BasicYanglao = BasicYanglao
				BasicShiye,_ := strconv.ParseFloat(row[8],64)
				hr.BasicShiye = BasicShiye
				BasicGongshang,_ := strconv.ParseFloat(row[9],64)
				hr.BasicGongshang = BasicGongshang
				BasicYiliao,_ := strconv.ParseFloat(row[10],64)
				hr.BasicYiliao = BasicYiliao
				BasicShengyu,_ := strconv.ParseFloat(row[11],64)
				hr.BasicShengyu = BasicShengyu
				BasicGjj,_ := strconv.ParseFloat(row[12],64)
				hr.BasicGjj = BasicGjj
				// 计算
				switch hr.BearDate {
				case "202001":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom11
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer11
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom11
						hr.PayShiyePer = hr.BasicShiye * shiyeCom11
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom11
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom11
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer11
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom11
						hr.PayGjjCom = hr.BasicGjj * gjjCom11
						hr.PayGjjPer = hr.BasicGjj * gjjPer11
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom10
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer10
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom10
						hr.PayShiyePer = hr.BasicShiye * shiyeCom10
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom10
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom10
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer10
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom10
						hr.PayGjjCom = hr.BasicGjj * gjjCom10
						hr.PayGjjPer = hr.BasicGjj * gjjPer10
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202002":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom21
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer21
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom21
						hr.PayShiyePer = hr.BasicShiye * shiyeCom21
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom21
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom21
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer21
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom21
						hr.PayGjjCom = hr.BasicGjj * gjjCom21
						hr.PayGjjPer = hr.BasicGjj * gjjPer21
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom20
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer20
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom20
						hr.PayShiyePer = hr.BasicShiye * shiyeCom20
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom20
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom20
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer20
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom20
						hr.PayGjjCom = hr.BasicGjj * gjjCom20
						hr.PayGjjPer = hr.BasicGjj * gjjPer20
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202003":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom31
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer31
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom31
						hr.PayShiyePer = hr.BasicShiye * shiyeCom31
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom31
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom31
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer31
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom31
						hr.PayGjjCom = hr.BasicGjj * gjjCom31
						hr.PayGjjPer = hr.BasicGjj * gjjPer31
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom30
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer30
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom30
						hr.PayShiyePer = hr.BasicShiye * shiyeCom30
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom30
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom30
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer30
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom30
						hr.PayGjjCom = hr.BasicGjj * gjjCom30
						hr.PayGjjPer = hr.BasicGjj * gjjPer30
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202004":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom41
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer41
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom41
						hr.PayShiyePer = hr.BasicShiye * shiyeCom41
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom41
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom41
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer41
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom41
						hr.PayGjjCom = hr.BasicGjj * gjjCom41
						hr.PayGjjPer = hr.BasicGjj * gjjPer41
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom40
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer40
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom40
						hr.PayShiyePer = hr.BasicShiye * shiyeCom40
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom40
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom40
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer40
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom40
						hr.PayGjjCom = hr.BasicGjj * gjjCom40
						hr.PayGjjPer = hr.BasicGjj * gjjPer40
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202005":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom51
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer51
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom51
						hr.PayShiyePer = hr.BasicShiye * shiyeCom51
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom51
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom51
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer51
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom51
						hr.PayGjjCom = hr.BasicGjj * gjjCom51
						hr.PayGjjPer = hr.BasicGjj * gjjPer51
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom50
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer50
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom50
						hr.PayShiyePer = hr.BasicShiye * shiyeCom50
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom50
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom50
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer50
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom50
						hr.PayGjjCom = hr.BasicGjj * gjjCom50
						hr.PayGjjPer = hr.BasicGjj * gjjPer50
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202006":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom61
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer61
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom61
						hr.PayShiyePer = hr.BasicShiye * shiyeCom61
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom61
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom61
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer61
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom61
						hr.PayGjjCom = hr.BasicGjj * gjjCom61
						hr.PayGjjPer = hr.BasicGjj * gjjPer61
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom60
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer60
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom60
						hr.PayShiyePer = hr.BasicShiye * shiyeCom60
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom60
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom60
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer60
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom60
						hr.PayGjjCom = hr.BasicGjj * gjjCom60
						hr.PayGjjPer = hr.BasicGjj * gjjPer60
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202007":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom71
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer71
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom71
						hr.PayShiyePer = hr.BasicShiye * shiyeCom71
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom71
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom71
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer71
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom71
						hr.PayGjjCom = hr.BasicGjj * gjjCom71
						hr.PayGjjPer = hr.BasicGjj * gjjPer71
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom70
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer70
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom70
						hr.PayShiyePer = hr.BasicShiye * shiyeCom70
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom70
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom70
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer70
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom70
						hr.PayGjjCom = hr.BasicGjj * gjjCom70
						hr.PayGjjPer = hr.BasicGjj * gjjPer70
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202008":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom81
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer81
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom81
						hr.PayShiyePer = hr.BasicShiye * shiyeCom81
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom81
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom81
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer81
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom81
						hr.PayGjjCom = hr.BasicGjj * gjjCom81
						hr.PayGjjPer = hr.BasicGjj * gjjPer81
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom80
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer80
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom80
						hr.PayShiyePer = hr.BasicShiye * shiyeCom80
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom80
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom80
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer80
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom80
						hr.PayGjjCom = hr.BasicGjj * gjjCom80
						hr.PayGjjPer = hr.BasicGjj * gjjPer80
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202009":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom91
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer91
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom91
						hr.PayShiyePer = hr.BasicShiye * shiyeCom91
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom91
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom91
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer91
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom91
						hr.PayGjjCom = hr.BasicGjj * gjjCom91
						hr.PayGjjPer = hr.BasicGjj * gjjPer91
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom90
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer90
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom90
						hr.PayShiyePer = hr.BasicShiye * shiyeCom90
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom90
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom90
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer90
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom90
						hr.PayGjjCom = hr.BasicGjj * gjjCom90
						hr.PayGjjPer = hr.BasicGjj * gjjPer90
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202010":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom101
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer101
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom101
						hr.PayShiyePer = hr.BasicShiye * shiyeCom101
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom101
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom101
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer101
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom101
						hr.PayGjjCom = hr.BasicGjj * gjjCom101
						hr.PayGjjPer = hr.BasicGjj * gjjPer101
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom100
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer100
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom100
						hr.PayShiyePer = hr.BasicShiye * shiyeCom100
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom100
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom100
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer100
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom100
						hr.PayGjjCom = hr.BasicGjj * gjjCom100
						hr.PayGjjPer = hr.BasicGjj * gjjPer100
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202011":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom111
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer111
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom111
						hr.PayShiyePer = hr.BasicShiye * shiyeCom111
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom111
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom111
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer111
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom111
						hr.PayGjjCom = hr.BasicGjj * gjjCom111
						hr.PayGjjPer = hr.BasicGjj * gjjPer111
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom110
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer110
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom110
						hr.PayShiyePer = hr.BasicShiye * shiyeCom110
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom110
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom110
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer110
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom110
						hr.PayGjjCom = hr.BasicGjj * gjjCom110
						hr.PayGjjPer = hr.BasicGjj * gjjPer110
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				case "202012":
					if hr.UserType == "农" {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom121
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer121
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom121
						hr.PayShiyePer = hr.BasicShiye * shiyeCom121
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom121
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom121
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer121
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom121
						hr.PayGjjCom = hr.BasicGjj * gjjCom121
						hr.PayGjjPer = hr.BasicGjj * gjjPer121
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					} else {
						hr.PayYanglaoCom = hr.BasicYanglao * yanglaoCom120
						hr.PayYanglaoPer = hr.BasicYanglao * yanglaoPer120
						hr.PayShiyeCom = hr.BasicShiye * shiyeCom120
						hr.PayShiyePer = hr.BasicShiye * shiyeCom120
						hr.PayGongshangCom = hr.BasicGongshang * gongshangCom120
						hr.PayYiliaoCom = hr.BasicYiliao * yiliaoCom120
						hr.PayYiliaoPer = hr.BasicYiliao * yiliaoPer120
						hr.PayShengyuCom = hr.BasicShengyu * shengyuCom120
						hr.PayGjjCom = hr.BasicGjj * gjjCom120
						hr.PayGjjPer = hr.BasicGjj * gjjPer120
						hr.PayTotalCom = hr.PayYanglaoCom + hr.PayShiyeCom + hr.PayGongshangCom + hr.PayYiliaoCom + hr.PayShengyuCom + hr.PayGjjCom
						hr.PayTotalPer = hr.PayYanglaoPer + hr.PayShiyePer + hr.PayYiliaoPer + hr.PayGjjPer
						hr.PayTotal = hr.PayTotalCom + hr.PayTotalPer
					}
				}

				hrArr = append(hrArr, *hr)
			}
		}
		// 数据入库
		insertErr := hrs.InsertHr2020(hrArr)
		if insertErr != nil {
			common.Failed(ctx,"插入数据入库失败")
			return
		}
		// 导出新excel文件
		newFile := excelize.NewFile()
		//index := newFile.NewSheet("Sheet2")
		// 表头
		headerTitle := []string{"Id","BearDate","UserName","UserId","UserStation","UserZhiwei","UserDate","UserType",
			"BasicYanglao","BasicShiye","BasicGongshang","BasicYiliao","BasicShengyu","BasicGjj","PayYanglaoCom",
		"PayYanglaoPer","PayShiyeCom","PayShiyePer","PayGongshangCom","PayYiliaoCom","PayYiliaoPer","PayShengyuCom",
		"PayGjjCom","PayGjjPer","PayTotalCom","PayTotalPer","PayTotal"}
		_ = newFile.SetSheetRow("Sheet1","A1",&headerTitle)
		// 读取数据库数据
		resAll,resAllErr := hrs.QueryHrDataAll()
		if resAllErr != nil {
			common.Failed(ctx,"获取所有失败")
			return
		}
		// 数据写入文件
		for index,rowData := range resAll {
			data := []interface{}{
				rowData.Id,
				rowData.BearDate,
				rowData.UserName,
				rowData.UserId,
				rowData.UserStation,

				rowData.UserZhiwei,
				rowData.UserDate,
				rowData.UserType,
				rowData.BasicYanglao,
				rowData.BasicShiye,

				rowData.BasicGongshang,
				rowData.BasicYiliao,
				rowData.BasicShengyu,
				rowData.BasicGjj,
				rowData.PayYanglaoCom,

				rowData.PayYanglaoPer,
				rowData.PayShiyeCom,
				rowData.PayShiyePer,
				rowData.PayGongshangCom,
				rowData.PayYiliaoCom,

				rowData.PayYiliaoPer,
				rowData.PayShengyuCom,
				rowData.PayGjjCom,
				rowData.PayGjjPer,
				rowData.PayTotalCom,

				rowData.PayTotalPer,
				rowData.PayTotal,
			}

			writeRowErr := newFile.SetSheetRow("Sheet1","A" + strconv.Itoa(index + 2),&data)
			if writeRowErr != nil {
				common.Failed(ctx,writeRowErr.Error())
				return
			}
		}
		newFilePath := "public/upload/excel/" + "res" + fileName
		saveNewErr := newFile.SaveAs(newFilePath)
		if saveNewErr != nil {
			common.Failed(ctx,"保存新文件错误")
			return
		}
		common.Success(ctx,resAll)
		return
	}
	common.Failed(ctx, "文件不存在")
	return
}
