package goxSort

import (
	"gox/utils"
	"math/rand"
)

func QuickSort[T any](elements []T, start, end int, compare func(T, T) int) {
	if end-start < 2 {
		//log.Println("return")
		return
	}

	mid := Partition2[T](elements, start, end, compare)
	//log.Println(mid)

	QuickSort[T](elements, start, mid+1, compare)
	QuickSort[T](elements, mid+1, end, compare)

}

//tony hore的轴点构造算法
func Partition1[T any](elements []T, start, end int, compare func(T, T) int) int {
	//左闭右开区间，故此处high需要减1
	end -= 1
	//将首元素随机与low,high之内的元素交换，降低最坏情况出现的概率
	utils.Swap[T](&elements[start], &elements[start+rand.Intn(end-start)])
	pivot := elements[start]

	for start < end {

		for start < end && compare(pivot, elements[end]) <= 0 {
			end--
		}
		elements[start] = elements[end]

		for start < end && compare(pivot, elements[start]) >= 0 {
			start++
		}
		elements[end] = elements[start]

	}
	elements[start] = pivot
	return start
}

//更简明的轴点构造方法
func Partition2[T any](elements []T, start, end int, compare func(T, T) int) int {
	//左闭右开区间，故此处high需要减1
	end -= 1
	//将首元素随机与low,high之内的元素交换，降低最坏情况出现的概率
	utils.Swap[T](&elements[start], &elements[start+rand.Intn(end-start)])
	pivot := elements[start]
	mid := start

	for k := start + 1; k <= end; k++ {
		if compare(elements[k], pivot) < 0 {
			mid++
			utils.Swap[T](&elements[mid], &elements[k])
		}
	}
	utils.Swap[T](&elements[start], &elements[mid])
	return mid
}
