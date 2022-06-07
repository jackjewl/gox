package slice

import (
	"gox/utils"
	"math/rand"
)

func QuickSort[T any](elems []T, low, high int, compare func(T, T) int) {
	if high-low < 2 {
		//log.Println("return")
		return
	}

	mid := Partition2[T](elems, low, high, compare)
	//log.Println(mid)

	QuickSort[T](elems, low, mid+1, compare)
	QuickSort[T](elems, mid+1, high, compare)

}

//tony hore的轴点构造算法
func Partition1[T any](elems []T, low, high int, compare func(T, T) int) int {
	//左闭右开区间，故此处high需要减1
	high -= 1
	//将首元素随机与low,high之内的元素交换，降低最坏情况出现的概率
	utils.Swap[T](&elems[low], &elems[low+rand.Intn(high-low)])
	pivot := elems[low]

	for low < high {

		for low < high && compare(pivot, elems[high]) <= 0 {
			high--
		}
		elems[low] = elems[high]

		for low < high && compare(pivot, elems[low]) >= 0 {
			low++
		}
		elems[high] = elems[low]

	}
	elems[low] = pivot
	return low
}

//更简明的轴点构造方法
func Partition2[T any](elems []T, low, high int, compare func(T, T) int) int {
	//左闭右开区间，故此处high需要减1
	high -= 1
	//将首元素随机与low,high之内的元素交换，降低最坏情况出现的概率
	utils.Swap[T](&elems[low], &elems[low+rand.Intn(high-low)])
	pivot := elems[low]
	mid := low

	for k := low + 1; k <= high; k++ {
		if compare(elems[k], pivot) < 0 {
			mid++
			utils.Swap[T](&elems[mid], &elems[k])
		}
	}
	utils.Swap[T](&elems[low], &elems[mid])
	return mid
}
