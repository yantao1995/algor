package excel

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

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
	xlsxFile, err := xlsx.OpenFile(name + ".xlsx")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	if len(xlsxFile.Sheets) < 1 {
		fmt.Println("未找到sheet1")
		return
	}
	sheet1 := xlsxFile.Sheets[0]
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
	err = xlsxFile.Save(copyName)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCarNumberGenerator(t *testing.T) {
	fmt.Printf("生成的个数:")
	count := 0
	fmt.Scanln(&count)
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	xls := xlsx.NewFile()
	source := "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sheet1, err := xls.AddSheet("car number")
	if err != nil {
		panic(err)
	}
	ls := int64(len(source))
	getRandStr := func() string {
		str := "苏U"
		for i := 0; i < 5; i++ {
			str += string(source[rd.Int63n(ls)])
		}
		return str
	}
	for i := 0; i < count; i++ {
		sheet1.AddRow().AddCell().SetString(getRandStr())
	}
	xls.Save("./车牌号生成.xlsx")
}

func TestModifyCell(t *testing.T) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}
	for k := range files {
		if strings.Contains(files[k].Name(), ".xlsx") {
			xlsxFile, err := xlsx.OpenFile(files[k].Name())
			if err != nil {
				panic("打开文件失败:" + err.Error())
			}
			sheet0 := xlsxFile.Sheets[0]
			sheet0.Rows[2].Cells[1].SetString("男")
			xlsxFile.Save("./" + files[k].Name())
		}
	}
}

func TestModifyFileName(t *testing.T) {
	fmt.Printf("输入文件名中需要被替换的文字:")
	source := ""
	fmt.Scanln(&source)
	fmt.Printf("输入文件名中替换的文字:")
	target := ""
	fmt.Scanln(&target)
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}
	for k := range files {
		if strings.Contains(files[k].Name(), source) {
			targetName := strings.ReplaceAll(files[k].Name(), source, target)
			os.Rename("./"+files[k].Name(), targetName)
		}
	}
}
