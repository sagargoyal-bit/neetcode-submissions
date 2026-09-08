func eraseOverlapIntervals(intervals [][]int) int {
//sort inertval
sort.Slice(intervals,func(i,j int)bool{
	return intervals[i][1]<intervals[j][1]
})
//check overlapping
// [[1,2],[1,4],[2,4]]
count:=0
current:=intervals[0]
for i:=1;i<len(intervals);i++{
	next:=intervals[i]
	if current[1]<=next[0]{
	current=next
	}else{
	count++
	}
}
return count
}
