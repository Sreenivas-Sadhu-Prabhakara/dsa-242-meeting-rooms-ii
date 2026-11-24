package main

import (
	"fmt"
	"sort"
)

// MinRooms returns the fewest rooms so all meetings run without sharing.
// A meeting ending exactly when another starts frees its room in time.
// Sort starts and ends separately and sweep: peak simultaneous meetings.
func MinRooms(intervals [][2]int) int {
	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, iv := range intervals {
		starts[i] = iv[0]
		ends[i] = iv[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)
	rooms, peak := 0, 0
	i, j := 0, 0
	for i < n {
		if starts[i] < ends[j] { // strictly before an end: truly concurrent
			rooms++
			if rooms > peak {
				peak = rooms
			}
			i++
		} else {
			rooms--
			j++
		}
	}
	return peak
}

func main() {
	fmt.Println(MinRooms([][2]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(MinRooms([][2]int{{1, 5}, {5, 8}}))
}
