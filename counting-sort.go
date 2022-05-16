// Challenge
// Given a list of integers, count and return the number of times each value appears as an array of integers.

// Function Description

// Complete the countingSort function in the editor below.

// countingSort has the following parameter(s):

// arr[n]: an array of integers
// Returns

// int[100]: a frequency array

package main

func countingSort(a []int32) []int32 {
	// Write your code here
	var oc = make([]int32, 100)

	for _, v := range a {
		oc[v]++
	}

	return oc
}
