package main

import "testing"

func TestMinRooms(t *testing.T) {
	cases := []struct {
		in   [][2]int
		want int
	}{
		{[][2]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[][2]int{{7, 10}, {2, 4}}, 1},
		{[][2]int{{1, 5}, {5, 8}}, 1}, // back-to-back shares one room
		{[][2]int{{1, 5}, {2, 6}, {3, 7}}, 3},
		{[][2]int{}, 0},
		{[][2]int{{4, 9}}, 1},
		{[][2]int{{1, 2}, {1, 2}, {1, 2}}, 3},
	}
	for _, c := range cases {
		if got := MinRooms(c.in); got != c.want {
			t.Errorf("MinRooms(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
