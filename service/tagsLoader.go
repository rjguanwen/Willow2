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

// 客户ID列表
var CustList []string
// 标签数量
var TagsCount int
// 标签列表
var TagsNameList []string
// 标签存储，key为标签名，value为与CustList相同顺序的该标签的值
var TagsStore = make(map[string] []byte)
// 标签存储，key为标签名，value为与CustList相同顺序的该标签的值的[]byte转换成的int值
var TagsStoreByInt  = make(map[string] int64)



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
			// 标签数量比列数少2，一个是序号一个是客户ID
			TagsCount = columnsCount -2
			// 初始化标签名称列表
			TagsNameList = record[2:]
			log.Printf("标签列表：%v", TagsNameList)
			// 初始化标签存储
			for i := 0; i < len(TagsNameList); i++ {
				// 以标签名为 key，暂存一个空数组
				TagsStore[TagsNameList[i]] = []byte{}
			}
			isFirstLine = false
		} else {
			// 生成客户 ID 列表
			currentCust := record[1]
			CustList = append(CustList, currentCust)
			for j := 0; j < len(TagsNameList); j++ {
				// "j+2" 的原因是 record 多了序号与客户ID两列
				tagValue := record[j+2]
				if tagValue == "True" {
					TagsStore[TagsNameList[j]] = append(TagsStore[TagsNameList[j]], 1)
				} else {
					TagsStore[TagsNameList[j]] = append(TagsStore[TagsNameList[j]], 0)
				}
			}
		}
	}
	//// 下面这部分程序存在错误 -- 待修正
	//// 将标签存储的 []byte 转换为 int32
	//for i := 0; i < len(TagsNameList); i++ {
	//	TagsStoreByInt[TagsNameList[i]] = utils.BytesToInt64(TagsStore[TagsNameList[i]])
	//}
	return err
}

//// 将传入的 [key]value 形式的查询条件，转化为可快速查询的查询标签列表
//func GetQueryList(queryMap map[string] byte) (queryList []byte){
//	queryList = []byte{}
//	for i := 0; i < TagsCount; i++ {
//		// 按顺序过滤所有标签
//		cTag := TagsNameList[i]
//		// 判单当前标签是否包含在传入的查询条件
//		value, isOk := queryMap[cTag]
//		if !isOk {
//			// 如果未包含
//			queryList = append(queryList, 0)
//		} else {
//			queryList = append(queryList, value)
//		}
//	}
//	log.Printf("组成查询标签列表：%v", queryList)
//	return queryList
//}
//
//// 根据查询标签列表，获取客户列表
//func QueryCustByMapTags(queryMap map[string] byte) (custList []string) {
//	queryList := GetQueryList(queryMap)
//	return QueryCustByTags(queryList)
//}

//// 根据查询标签列表，获取客户列表
//func QueryCustByTags(queryTags []string) (custList []string) {
//	var fTagValue []byte
//	noInit := true
//	for i := 0; i < len(queryTags); i++ {
//		tmpTagValue, isOk := TagsStore[queryTags[i]]
//		if isOk {
//			if noInit {
//				fTagValue = tmpTagValue
//				noInit = false
//			} else {
//				//fTagValue = tmpTagValue & fTagValue
//			}
//		}
//	}
//
//}