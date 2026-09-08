func insert(intervals [][]int, newInterval []int) [][]int {
 i:=0
 n:=len(intervals)
 result:=[][]int{}

 for i<n && intervals[i][1]< newInterval[0]{
	result=append(result,intervals[i])
	i++
 }

for i<n && intervals[i][0]<=newInterval[1] && intervals[i][1]>=newInterval[0]{
	newInterval[0]=min(intervals[i][0],newInterval[0])
	newInterval[1]=max(intervals[i][1],newInterval[1])
	i++
}
result=append(result,newInterval)

for i<n{
	result=append(result,intervals[i])
	i++
}
return result

}
