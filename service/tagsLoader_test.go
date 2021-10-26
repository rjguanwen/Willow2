// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2021/10/23 13:53
// @Update  姓名（需要改）  2021/10/23 13:53
package service

import (
	"fmt"
	"testing"
)

func TestLoadTagsData(t *testing.T) {
	LoadTagsData("../data/tags-all-sample.csv")
	fmt.Println("-CustList---")
	fmt.Println(CustList)
	fmt.Println("-TagsCount---")
	fmt.Println(TagsCount)
	fmt.Println("-TagsNameList---")
	fmt.Println(TagsNameList)
	fmt.Println("-TagsStore---")
	fmt.Println(TagsStore)
	fmt.Println("-len(TagsStore)---")
	fmt.Println(len(TagsStore))
	fmt.Println("-TagsStoreByInt---")
	fmt.Println(TagsStoreByInt)

	fmt.Println("---------------------------")
	//queryMap := make(map[string]byte)
	//queryMap["A"] = 1
	//queryMap["Z"] = 1
	//queryMap["AA"] = 1
	//queryMap["AB"] = 1
	//queryMap["AC"] = 1
	//queryMap["BB"] = 1
	//queryList := GetQueryList(queryMap)
	//fmt.Println("-queryList---")
	//fmt.Println(queryList)
}