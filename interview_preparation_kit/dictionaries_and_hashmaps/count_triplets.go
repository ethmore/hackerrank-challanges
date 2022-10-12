package dictionariesandhashmaps

func countTriplets(arr []int64, r int64) int64 {
	triplets := int64(0)
	occur := make(map[int64]int64)
	occurPairs := make(map[int64]int64)

	for i := int64(len(arr)) - 1; i >= 0; i-- {

		val, inOcurrs := occurPairs[arr[i]*r]
		if inOcurrs {
			triplets += val
		}
		val2, inSeq := occur[arr[i]*r]
		if inSeq {
			occurPairs[arr[i]] = occurPairs[arr[i]] + val2
		}
		occur[arr[i]] = occur[arr[i]] + 1
	}

	return triplets
}
