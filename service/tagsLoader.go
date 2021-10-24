// @Title  tagsLoader
// @Description  标签加载
// @Author  zhengbin  2021/10/22 16:33
// @Update  姓名（需要改）  2021/10/22 16:33
package service

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// 读取消费者标签数据文件
func LoadTagsData(dataFilePath string) (err error){
	dataFile, err := os.OpenFile(dataFilePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("文件打开错误：%s", dataFilePath)
		return err
	}
	defer dataFile.Close()

	fileReader := csv.NewReader(dataFile)

	isFirstLine := true
	for {
		// 循环按行读取文件
		record, err := fileReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("文件读取错误：", err)
			return err
		}
		if isFirstLine {
			columnsCount := len(record)
			log.Printf("列数：%d", columnsCount)
			tagsList := record[2:]
			log.Printf("标签列表：%v", tagsList)
			isFirstLine = false
			continue
		}
		firstCust := record[1]
		firstCustTags := record[2:]
		log.Printf("消费者ID：%s", firstCust)
		log.Printf("消费者标签：%v", firstCustTags)
		break
	}
	return err
}