package fletcher

// Fletcher16 calculates Fletcher16 checksum of
// data slice
func Fletcher16(data []byte) uint16 {
	var sum1 uint16 = 0xFF
	var sum2 uint16 = 0xFF

	var i = 0
	var len = len(data)

	for len > 0 {
		tlen := len
		if len > 20 {
			tlen = 20
		}
		len -= tlen

		sum1 += uint16(data[i])
		sum2 += sum1
		i++
		tlen--
		for tlen > 0 {
			sum1 += uint16(data[i])
			sum2 += sum1
			i++
			tlen--
		}
		sum1 = (sum1 & 0xff) + (sum1 >> 8)
		sum2 = (sum2 & 0xff) + (sum2 >> 8)
	}
	/* Second reduction step to reduce sums to 8 bits */
	sum1 = (sum1 & 0xff) + (sum1 >> 8)
	sum2 = (sum2 & 0xff) + (sum2 >> 8)
	return sum2<<8 | sum1
}
