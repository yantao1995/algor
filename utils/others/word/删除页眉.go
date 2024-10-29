package word

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nguyenthenguyen/docx"
)

/*

改库源码


func (d *Docx) DeleteHeader() (err error) {
	return deleteHeaderFooter(d.headers)
}

func deleteHeaderFooter(headerFooter map[string]string) (err error) {
	newString, err := encode(" ")
	if err != nil {
		return err
	}
	for k := range headerFooter {
		headerFooter[k] = newString
	}

	return nil
}
*/

func clear(path string) {
	// 打开现有的 .docx 文件
	doc, err := docx.ReadDocxFile(path)
	if err != nil {
		log.Fatal(err)
	}

	doc.Editable().DeleteHeader()

	if err := doc.Editable().WriteToFile(path + ".docx"); err != nil {
		panic(err)
	}

	doc.Close()
}

// WalkAndProcessFiles 遍历目录并处理所有 Word 文件
func WalkAndProcessFiles(rootDir string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查文件是否是 .docx 文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".docx") {
			fmt.Printf("正在处理文件: %s\n", path)
			clear(path)
		}
		return nil
	})
}

// WalkAndProcessFiles 遍历目录并处理所有 Word 文件
func WalkAndProcessFiles2(rootDir string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查文件是否是 .docx 文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".docx.docx") {
			if err := os.Remove(strings.ReplaceAll(path, ".docx.docx", ".docx")); err != nil {
				panic(err)
			}
			if err := os.Rename(path, strings.ReplaceAll(path, ".docx.docx", ".docx")); err != nil {
				panic(err)
			}
		}
		return nil
	})
}

func ClearDocxHeaders() {
	rootDir := "./" // 目标目录设为当前目录

	if err := WalkAndProcessFiles(rootDir); err != nil {
		fmt.Printf("处理目录失败: %v\n", err)
	} else {
		fmt.Println("所有文件处理完成")
	}

	if err := WalkAndProcessFiles2(rootDir); err != nil {
		fmt.Printf("改名处理目录失败: %v\n", err)
	} else {
		fmt.Println("改名所有文件处理完成")
	}
	fmt.Scanln()
}
