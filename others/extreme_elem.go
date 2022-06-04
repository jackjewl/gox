
package others

func NotExtremeElem(elems []int)int{
	min,max:=0,0
	for i:=0;i<3;i++{
		if elems[i]>elems[max]{
			max=i
		}
		if elems[i]<elems[min]{
			min=i
		}
	}
	
	for i:=0;i<3;i++{
		if min!=i&&max!=i{
			return elems[i]
		}
	}

	return 0
}

	

