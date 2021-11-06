package tools

// 数组排序
func SortSliceInt(data []int64) []int64 {
	for i:=0; i<len(data); i++ {
		for j:=i+1; j<len(data); j++ {
			if data[i] > data[j] {
				data[j],data[i] = data[i],data[j]
			}
		}
	}
	return data
}

// 有序数组 去重
func RemoveDup(data []int64) []int64 {
	if len(data) <= 1 {
		return data
	}else {
		i := 0
		for j:= 1; j<len(data);j++ {
			if data[i] != data[j] {
				i++
				data[i] = data[j]
			}
		}
		return data[:i+1]
	}
}

// int 类型  不排序去重
func RemoveIntByMap(data []int64) []int64 {
	//存放返回的不重复切片
	var result []int64
	// 存放不重复主键
	tempMap := map[int64]byte{}
	for _, e := range data {
		l := len(tempMap)
		//当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		tempMap[e] = 0
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

// string 去重

//slice去重
func RemoveStringByMap(slc []string) []string {
	var result []string          //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}