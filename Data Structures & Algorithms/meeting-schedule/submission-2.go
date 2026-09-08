/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].end
	})
 if len(intervals) == 0 {
        return true
    }

for i := 1; i < len(intervals); i++ {
	prevEnd := intervals[i-1]
	curr:=intervals[i]
    if curr.start < prevEnd.end {
        return false
    }
}
return true
}
