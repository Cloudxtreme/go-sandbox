// http://he-the-great.livejournal.com/49072.html

package main

import (
	"fmt"
	"sort"
	"time"
)

// Define us a type so we can sort it
type TimeSlice []time.Time

// Forward request for length
func (p TimeSlice) Len() int {
	return len(p)
}

// Define compare
func (p TimeSlice) Less(i, j int) bool {
	return p[i].Before(p[j])
}

// Define swap over an array
func (p TimeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	// Create slice of two times
	events := make([]time.Time)
	// 4:30AM
	events[0], _ = time.Parse("UnixDate", "04:30")
	events[1] = time.Now()
	events[2] = time.Now()
	// 3:12PM
	// events[1], _ = time.Parse("15:04", "03:12")
	// events[1], _ = time.Parse("Mon Jan _2 15:04:05 MST 2006", "46123")
	// Wrap array in type for sorting
	fmt.Println(events)
	// fmt.Println(time.Now().Format("2006/01/02 15:04:05 MST"))
	// fmt.Println(time.Now().Format("1371646123"))
	// fmt.Println(time.Now().Unix())

	sort.Sort(TimeSlice(events))

	fmt.Println(events)

	// fmt.Println(time.Now().Unix())
}
