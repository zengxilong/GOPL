package popcount

func PopCountClean(x uint64) int {
	n := 0
	for x != 0 {
		n++
		x = x & (x - 1)
	}
	return n
}
