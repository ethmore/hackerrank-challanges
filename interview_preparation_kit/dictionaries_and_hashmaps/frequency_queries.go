package dictionariesandhashmaps

func freqQuery(queries [][]int32) []int32 {
	hMap := make(map[int32]int32)
	freqMap := make(map[int32]int32)
	var occur []int32

	for _, elem := range queries {
		if elem[0] == 1 {
			hMap[elem[1]]++
			freqMap[hMap[elem[1]]]++

			if freqMap[hMap[elem[1]]-1] > 0 {
				freqMap[hMap[elem[1]]-1]--
			}

		} else if elem[0] == 2 {
			if hMap[elem[1]] > 0 {
				freqMap[hMap[elem[1]]]--
				freqMap[hMap[elem[1]]-1]++
				hMap[elem[1]]--
			}

		} else if elem[0] == 3 {
			if freqMap[elem[1]] > 0 {
				occur = append(occur, 1)
			} else {
				occur = append(occur, 0)

			}

		}
	}

	return occur
}
