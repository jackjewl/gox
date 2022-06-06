package others

func Power2BF(n int) int {
	pow := 1
	for n > 0 {
		pow <<= 1
		n--
	}
	return pow
}

func Power2(n int) int {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		return Power2(n/2) * Power2(n/2)
	} else {
		return Power2(n/2) * Power2(n/2) * 2
	}
}
