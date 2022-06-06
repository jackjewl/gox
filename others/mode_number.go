package others

//求众数，占总元素一半以上的数为众数
func ModeNumber(elems []int) (int, bool) {
	candidateIndex := FindModeCandidate(elems)
	mode := elems[candidateIndex]
	if ValidIsModeNumber(elems, candidateIndex) {
		return mode, true
	} else {
		return 0, false
	}
}

func FindModeCandidate(elems []int) (index int) {
	diffCounter := 0
	candidateIndex := 0
	for i := 0; i < len(elems); i++ {
		if diffCounter == 0 {
			candidateIndex = i
			diffCounter = 1
		} else {
			if elems[candidateIndex] == elems[i] {
				diffCounter++
			} else {
				diffCounter--
			}
		}
	}
	return candidateIndex
}

func ValidIsModeNumber(elems []int, candidateIndex int) bool {
	modeCandidate := elems[candidateIndex]
	repeatCounter := 0
	for _, elem := range elems {
		if elem == modeCandidate {
			repeatCounter++
		}
	}
	if repeatCounter > len(elems)>>1 {
		return true
	}
	return false

}
