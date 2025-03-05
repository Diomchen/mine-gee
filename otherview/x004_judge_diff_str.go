package otherview

func JudgeDiffStr(str1 string) bool {
	var mask1, mask2, mask3, mask4 uint64
	var mask *uint64

	for _, s := range str1 {
		n := uint64(s)
		if n < 64 {
			mask = &mask1
		} else if n < 128 {
			mask = &mask2
			n -= 64
		} else if n < 192 {
			mask = &mask3
			n -= 128
		} else {
			mask = &mask4
			n -= 192
		}

		if (*mask)&(1<<n) != 0 {
			return false
		}
		*mask = (*mask) | uint64(1<<n)
	}
	return true
}
