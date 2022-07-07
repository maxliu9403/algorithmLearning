package minMeetingRooms

import (
	"fmt"
	"testing"
)

func TestMinMeetingRooms(t *testing.T) {
	i1 := [][]int{{0, 30}, {5, 10}, {15, 20}}
	o1 := minMeetingRooms(i1)
	fmt.Println("==", o1)
	
	i2 := [][]int{{7, 10}, {2, 4}}
	o2 := minMeetingRooms(i2)
	fmt.Println("==", o2)

	i3 := [][]int{{1, 10}, {2, 7}, {3, 19}, {8, 12}, {10, 20}, {11, 30}}
	o3 := minMeetingRooms(i3)
	fmt.Println("==", o3)
}
