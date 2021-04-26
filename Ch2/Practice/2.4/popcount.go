package popcount

func PopCount(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&1 == 1 {
			n++
		}
		x = x >> 1
	}
	return n
}
