// 此文件的作用是将 code examples 下的代码传入到数据库中
package main

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/initialization"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("get files")
	f := readFiles()
	fmt.Println("convert type")
	codeExamples := convertType(f)
	fmt.Printf("save to db %d records\n", len(codeExamples))
	saveToDB(codeExamples)
}

type ChainRpcCodeExamples struct {
	Chain      string `json:"chain"`
	Cli        string `json:"cli"`
	Javascript string `json:"javascript"`
	Python     string `json:"python"`
	Go         string `json:"go"`
}

func convertType(files []FileInfo) []ChainRpcCodeExamples {
	m := make(map[string]ChainRpcCodeExamples)
	for _, file := range files {
		chain := strings.Split(file.RelativePath, "/")[0]
		languageType := strings.Split(file.RelativePath, "/")[1]
		switch languageType {
		case "cli.txt":
			if _, ok := m[chain]; !ok {
				m[chain] = ChainRpcCodeExamples{}
			}
			c := m[chain]
			c.Chain = chain
			c.Cli = file.Content
			m[chain] = c
		case "js.txt":
			if _, ok := m[chain]; !ok {
				m[chain] = ChainRpcCodeExamples{}
			}
			c := m[chain]
			c.Chain = chain
			c.Javascript = file.Content
			m[chain] = c
		case "python.txt":
			if _, ok := m[chain]; !ok {
				m[chain] = ChainRpcCodeExamples{}
			}
			c := m[chain]
			c.Chain = chain
			c.Python = file.Content
			m[chain] = c
		case "go.txt":
			if _, ok := m[chain]; !ok {
				m[chain] = ChainRpcCodeExamples{}
			}
			c := m[chain]
			c.Chain = chain
			c.Go = file.Content
			m[chain] = c
		}
	}
	// map 转 slice
	var codeExamples []ChainRpcCodeExamples
	for _, v := range m {
		codeExamples = append(codeExamples, v)
	}
	return codeExamples
}

func saveToDB(codeExamples []ChainRpcCodeExamples) {
	db, _ := application.GetBean[*gorm.DB]("db")
	// 清空 code examples table
	err := db.Exec("truncate table t_cl_rpc_code_examples").Error
	if err != nil {
		panic(err)
	}
	err = db.Table("t_cl_rpc_code_examples").Create(&codeExamples).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("save to db success")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	initialization.InitDB()
}

func readFiles() []FileInfo {
	txtFiles, err := readTxtFilesFromDir(".")
	if err != nil {
		panic(err)
	}
	return txtFiles
}

type FileInfo struct {
	Content      string
	RelativePath string
}

func readTxtFilesFromDir(dirPath string) ([]FileInfo, error) {
	var txtFiles []FileInfo

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			relativePath, err := filepath.Rel(dirPath, path)
			if err != nil {
				return err
			}

			txtFiles = append(txtFiles, FileInfo{
				Content:      string(content),
				RelativePath: relativePath,
			})
		}
		return nil
	})

	return txtFiles, err
}
