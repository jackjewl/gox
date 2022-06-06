package others

import (
	"gox/sort/slice"
	"gox/utils"
)

//找出第k大的数

func LinearSelectKth(elems []int, k int) (index, value int) {
	groupLen := 5

	return SelectKth(elems, k, groupLen)

}

func SelectKth(elems []int, k int, groupLen int) (index, value int) {

	//1当问题规模elems的长度足够小的时候，用任何一种选取算法来返回，这里选用排序之后取第k大的元素
	//2否则，将elems分为若干组，每一组大小为 groupLen
	//3对每一个小分组，进行排序，求得每个小分组的中位数
	//4 对所有的中位数，求所有中位数的中位数，调用递归来取得中位数，假设为m
	//5根据m来分类元素，partition划分
	//(6)减除无用的区间，继续递归找

	//(1)
	if len(elems) < groupLen {
		return SelectKthMin(elems, k)
	}

	//(2),(3)
	middleIndex, _ := DevideIntoGroupAndFindMiddleNumber(elems, k, groupLen)

	//(4),(5)
	utils.Swap[int](&elems[0], &elems[middleIndex])
	pivot := Partition[int](elems, 0, len(elems), func(a, b int) int {
		return a - b
	})

	//(6)
	if pivot == k {
		return k, elems[k]
	} else if k < pivot {
		return LinearSelectKth(elems[:pivot], k)
	} else {
		return LinearSelectKth(elems[pivot+1:], k-pivot)
	}
}

func SelectKthMin(elems []int, k int) (index, value int) {
	//快排
	slice.QuickSort[int](elems, 0, len(elems), func(a, b int) int { return a - b })
	return k - 1, elems[k-1]
}

//
func DevideIntoGroupAndFindMiddleNumber(elems []int, k int, groupLen int) (index, value int) {

	middleLocation := 0
	i := 0
	for ; i < len(elems)-groupLen; i += groupLen {
		part := elems[i : i+groupLen]
		SelectKthMin(elems, len(part)>>1)
		utils.Swap[int](&elems[middleLocation], &part[len(part)>>1])
		middleLocation++
	}

	if i-groupLen <= len(elems)-1 {
		part := elems[i-groupLen : len(elems)]
		SelectKthMin(elems, len(part)>>1)
		utils.Swap[int](&elems[middleLocation], &part[len(part)>>1])
		middleLocation++
	}

	return LinearSelectKth(elems[:middleLocation], middleLocation>>1)

}

//更简明的轴点构造方法
func Partition[T any](elems []T, low, high int, compare func(T, T) int) int {
	//左闭右开区间，故此处high需要减1
	high -= 1
	//将首元素随机与low,high之内的元素交换，降低最坏情况出现的概率
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
