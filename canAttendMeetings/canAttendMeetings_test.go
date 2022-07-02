package canAttendMeetings

import (
	"fmt"
	"testing"
)

func TestCanAttendMeetings(t *testing.T) {
	itv1 := [][]int{{0, 30}, {5, 10}, {15, 20}}
	v1 := canAttendMeetings(itv1)
	fmt.Println("intervals1", v1)
	itv2 := [][]int{{7, 10}, {2, 4}}
	v2 := canAttendMeetings(itv2)
	fmt.Println("intervals2", v2)
}
