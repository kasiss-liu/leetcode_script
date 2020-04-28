package main

//一种指针控制的原地算法 在本地运行失败
import (
	"testing"
	"unsafe"
)

// 两个强制转换函数，具体方法请翻源码思考，这里不多做解释
func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func reverseWord(s string) string {
	sArr := str2bytes(s)
	sSize := len(sArr)
	mid := sSize / 2
	curLeft := 0  // 维护的一个实际有用长度的指针
	find := false // 表示是否找到单词
	for i := 0; i < mid; i++ {
		// 先将左边元素无条件复制到右边
		tmp := sArr[sSize-i-1]
		sArr[sSize-i-1] = sArr[i]

		// 如果发现了单词
		if find {
			// 将当下元素复制到实际的左指针
			sArr[curLeft] = tmp
			curLeft++
			if tmp == 32 {
				find = false // 发现空格后 要将发现单词的状态置否
			}
		} else {
			if tmp != 32 {
				// 否则发现单词，开始复制
				sArr[curLeft] = tmp
				curLeft++
				find = true
			}
		}
	}
	for i := mid; i < sSize; i++ {
		// 到这个循环，字符串的右半边都是反转过的左半边
		// 所以从右半边遍历，按刚刚的方法一个个将元素复制到实际的左指针
		if find {
			sArr[curLeft] = sArr[i]
			curLeft++
			if sArr[i] == 32 {
				find = false
			}
		} else {
			if sArr[i] != 32 {
				sArr[curLeft] = sArr[i]
				curLeft++
				find = true
			}
		}
	}
	// 判断有用长度是否为0
	if curLeft == 0 {
		return ""
	}
	// 判断最后一个元素是不是空格
	if sArr[curLeft-1] == 32 {
		curLeft--
	}
	find = false
	// 双指针，记录单词的开始下标和结束下表
	left := 0
	right := 0
	for i := 0; i < curLeft; i++ {
		if find {
			if sArr[i] == 32 {
				find = false
				right = i - 1
				// 当发现单词，对left到right执行一次翻转
				for left < right {
					tmp := sArr[left]
					sArr[left] = sArr[right]
					sArr[right] = tmp
					left++
					right--
				}
			}
		} else {
			if sArr[i] != 32 {
				find = true
				left = i
			}
		}
	}
	// 最后一个单词被忽略了，这里补回
	right = curLeft - 1
	for left < right {
		tmp := sArr[left]
		sArr[left] = sArr[right]
		sArr[right] = tmp
		left++
		right--
	}
	// 最后将有用长度的字符串返回为答案即可
	return bytes2str(sArr[0:curLeft])
}

func TestReversWordStay(t *testing.T) {
	t.Error(reverseWord("ni hao"))
}
