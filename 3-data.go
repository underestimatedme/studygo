package main

import "fmt"

/**
 * 数据类型：包含array,slice,map,struct
 */

func defStruct() {
	// 定义了一个会员，包含name/age/idNo/mobileNo/registerDate字段
	type Member struct {
		name         string
		age          int
		idNo         string
		mobileNo     string
		registerDate string
	}
	zhangsan := Member{"Zhang San", 20, "310xxxxxx", "+86-11111111111", "2022-03-15 10:10"}
	fmt.Println(zhangsan)
	var lisi Member
	lisi.name = "Li Si"
	lisi.age = 21
	lisi.idNo = "333xxxxxx"
	lisi.mobileNo = "+86-1234567890"
	lisi.registerDate = "2022-05-15 00:00"
	fmt.Println(lisi)
	// 定义struct指针
	var zhangsanCopy *Member
	zhangsanCopy = &zhangsan
	fmt.Println(zhangsanCopy)
}

// 切片-slice和python中的类似，string/数组使用[:]这种格式
func defSlice() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))
	nums = append(nums, 0)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3, 4)
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	nums2 := make([]int, len(nums), (cap(nums))*2)
	//复制数组nums熬nums2
	copy(nums2, nums)
	fmt.Println(nums2, len(nums2), cap(nums2))

}

// map和python中的基本一致
func defMap() {
	// 初始化方式1: map初始化为nil
	var memberIdMap map[string]string
	memberIdMap = make(map[string]string)
	memberIdMap["zhangsan"] = "310xxxxxx"
	memberIdMap["lisi"] = "333xxxxxx"
	fmt.Println(memberIdMap)
	// 初始化方式2
	testMap := map[string]string{"a": "1", "b": "2", "c": "3"}
	for key := range testMap {
		fmt.Println(key, testMap[key])
	}
	// 删除元素
	delete(testMap, "a")
}

func main() {
	defStruct()
	defSlice()
	defMap()
}
