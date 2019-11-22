package array

// This function is used to search second large off array.
func SecondLargeOfArray(arr []int) int {
	large := 0
	slarge := 0
	for _, v := range arr {
		if v > large {
			slarge = large
			large = v
		}
		if v > slarge && v < large {
			slarge = v
		}
	}
	return slarge
}

func SumWithReturnOnValue(val []int) int {
	resp := SumOfArray(val)
	if resp > 9 {
		resp1 := []int{}
		for _, s := range IntToArray(resp) {
			resp1 = append(resp1, int(s))
		}
		return SumWithReturnOnValue(resp1)
	} else {
		return resp
	}
	return resp
}

func IntToArray(val int) []byte {
	var rightMost, tempIntVar int
	var byteArray []byte

	tempIntVar = val
	// need to reverse the order
	for {
		rightMost = tempIntVar % 10
		byteArray = append(byteArray, byte(rightMost))

		tempIntVar /= 10

		if tempIntVar == 0 {
			break
		}
	}
	fixByteArray := []byte{}
	for i := range byteArray {
		n := byteArray[len(byteArray)-1-i]
		fixByteArray = append(fixByteArray, n)
	}

	return fixByteArray
}

func SumOfArray(val []int) int {
	resp := 0
	for _, a := range val {
		resp = resp + a
	}
	return resp
}
