func eraseOverlapIntervals(intervals [][]int) int {
	// sort the intervals
sort.Slice(intervals,func(i,j int)bool{
	return intervals[i][0]<intervals[j][0]
})
// [[1,2],[1,4][2,4]]
count:=0
prev:=intervals[0]
for i:=1;i<len(intervals);i++{
	curr:=intervals[i]
	//overlapping
	if prev[1]>curr[0]{
		count++
		//keep the interval of the smaller value for no future overlapping
		if curr[1]<prev[1]{
			prev=curr
		}
	}else{
		prev=curr
	}
}
return count
}
