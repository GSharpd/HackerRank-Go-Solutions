package main

func main() {
	a := []int32{1, 2, 3, 2, 1}

	lonelyinteger(a)
}

func lonelyinteger(a []int32) int32 {
	// Write your code here
	if len(a) == 1 {
		return a[0]
	}

	intMap := make(map[int32]int32)

	for _, v := range a {
		intMap[v]++
	}
	for _, x := range a {
		if intMap[x] == 1 {
			return x
		}
	}
	return 0
}
