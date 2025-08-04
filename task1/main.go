


func twoSum(nums []int, target int) []int {
    idx_map := make(map[int]int)
    for idx, val := range nums {
        complement := target - val 
        if targ_idx, exists := idx_map[complement]; exists{
            return []int{idx, targ_idx}
        }
        idx_map[val] = idx 
    }
    return nil
}

/*
赋值  :=
map初始化    make   ；  中括号 [int]int
返回临时数组   []int{...}
*/




func merge(intervals [][]int) [][]int {
     if len(intervals) == 0{
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
    
	res := [][] int{}
	mb := intervals[0]

	for _, val := range intervals[1:] {
		if val[0] <= mb[1]{
			mb[1] = val[1]
		}else{
			res = append(res, mb)
			mb = val 
		}
	}
    res = append(res, mb)
    return res
}


/**
sort.Slice 排序
二维数组 [][]int
res = append(res, mb)
*/



func removeDuplicates(nums []int) int {
    idx := 0
    for _, val := range nums{
        if val != nums[idx]{
            idx += 1
            nums[idx] = val
        }
    }
    return idx+1
}




func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] < 9{
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    return append([]int{1}, digits...)
}

/*
	数组倒序的方法 
	1、从后往前便利；
	for i := len(digits) - 1; i >= 0; i-- {}
	2、翻转后输出
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
	3、递归倒序
	func p(arr []int, len){
	    if len >= 0{
		    print(arr[len])
		    p(arr, len - 1)
		}
	}

	往数组前加一个元素
	1. append + 切片  append([]int{1}, slice...)       ------ 创建临时切片
	2. 开辟一块新内存，拷贝       					    ------ 更高效
		news := make([]int, len(slice)+1)
		copy(news[1:], slice)
*/





func longestCommonPrefix(strs []string) string {
    prefix := strs[0]
    for k := 0; k < len(strs); k++ {
        for j := 0; j < len(prefix); j++ {
            if j >= len(strs[k]) || strs[k][j] != prefix[j] {
                prefix = prefix[:j]
                break
            }
        }
        if prefix == "" {
            return ""
        }
    }
    return prefix
}


/*
字符串 切分： strings.Split("a,b,c", ",")      // ["a", "b", "c"]
字符串 拼接： strings.Join([]string{"a", "b", "c"}, "-") // "a-b-c"
字符串 子串： strings.Contains(s, "word")   返回bool
替换   strings.Replace(s, "xx", "replace_word", 1)  // 还有 ReplaceAll
string.ToUpper("xxxx")  ToLower("xx")

strconv :  Atoi 转 int   Itoa 转字符串

[]byte("str") <---> string(bytes)
strings.TrimSpace(s)
*/




func isPalindrome(x int) bool {
    if (x < 0)  || (x % 10 == 0 ) {
        return false
    }

    rev := 0
    for rev < x / 10{
        rev =  rev * 10 + x % 10
        x /= 10
    }

    return  rev == x || rev == x /10
}




func singleNumber(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    sta := nums[0]
    for _, item:= range nums[1:]{       // _, item 
        sta ^= item
    }
    return sta
}