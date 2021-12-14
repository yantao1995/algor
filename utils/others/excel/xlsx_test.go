package excel

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/tealeg/xlsx"
)

func TestChangeColor(t *testing.T) {
	name := ""
	fmt.Print("输入文件名:")
	fmt.Scanln(&name)
	excel(name)
	fmt.Print("执行完成，按回车键退出...")
	fmt.Scanln(&name)
}
func excel(name string) {
	copyName := name + "_修改后的.xlsx"
	xlsxfile, err := xlsx.OpenFile(name + ".xlsx")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	if len(xlsxfile.Sheets) < 1 {
		fmt.Println("未找到sheet1")
		return
	}
	sheet1 := xlsxfile.Sheets[0]
	style1 := xlsx.NewStyle()
	style2 := xlsx.NewStyle()
	style3 := xlsx.NewStyle()
	style4 := xlsx.NewStyle()
	style5 := xlsx.NewStyle()
	style6 := xlsx.NewStyle()
	style7 := xlsx.NewStyle()
	style1.Font.Color = "ff88c2"
	style2.Font.Color = "ff3333"
	style3.Font.Color = "ffff33"
	style4.Font.Color = "99ff99"
	style5.Font.Color = "227700"
	style6.Font.Color = "00bbff"
	style7.Font.Color = "ccccff"
	for k := range sheet1.Rows {
		if k > 0 {
			val, err := strconv.Atoi(sheet1.Rows[k].Cells[1].Value)
			if err != nil {
				continue
			}
			if val < 50 {
				sheet1.Rows[k].Cells[1].SetStyle(style1)
			} else if val >= 90 && val < 95 {
				sheet1.Rows[k].Cells[1].SetStyle(style2)
			} else if val >= 85 && val < 90 {
				sheet1.Rows[k].Cells[1].SetStyle(style3)
			} else if val >= 80 && val < 85 {
				sheet1.Rows[k].Cells[1].SetStyle(style4)
			} else if val >= 70 && val < 80 {
				sheet1.Rows[k].Cells[1].SetStyle(style5)
			} else if val >= 60 && val < 70 {
				sheet1.Rows[k].Cells[1].SetStyle(style6)
			} else if val >= 50 && val < 60 {
				sheet1.Rows[k].Cells[1].SetStyle(style7)
			}
		}
	}
	err = xlsxfile.Save(copyName)
	if err != nil {
		fmt.Println(err)
	}
}
