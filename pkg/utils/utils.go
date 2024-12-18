package utils

func ContainPos(posCollection [][2]int, pos [2]int) bool {
	for _, val := range posCollection {
		if val == pos {
			return true
		}
	}
	return false
}

func UpdSlice(sliceToUpdate [][2]int, givenSlice [][2]int) {
	for i := range sliceToUpdate {
		sliceToUpdate[i] = givenSlice[i]
	}
}
