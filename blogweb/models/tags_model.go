package models

import (
	"strings"

)

func HandleTagsListData(tags []string)map[string]int  {
	var tagsMap = make(map[string]int)

	for _,tag:=range tags{   //把通过Tasg取到的数据 遍历
		s1:=strings.Split(tag,"&")  //分割
		for _,value:=range s1{  //再遍历分割后的数据
			tagsMap[value]++  //查找字符串数组里有多少重复项
			//相同的value会++
		}
	}

return tagsMap
}
