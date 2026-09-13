/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
 if len(intervals)==0{
	return 0
 }

 startTime:=make([]int,len(intervals))
 endTime:=make([]int,len(intervals))
//sperate start and end time 

 for i:=0;i<len(intervals);i++{
	startTime[i]=intervals[i].start
	endTime[i]=intervals[i].end
 }


//  sort both the arrays

sort.Ints(startTime)
sort.Ints(endTime)

i, j := 0, 0
rooms:=0
maxrooms:=0

for i<len(startTime){
	if startTime[i]<endTime[j]{
		rooms++
		i++
		if rooms>maxrooms{
			maxrooms=rooms
		}
	}else{
		rooms--
		j++
	}
}
return maxrooms
}
