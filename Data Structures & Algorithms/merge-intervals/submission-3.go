func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//crete and blank list
	lists := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		prev := lists[len(lists)-1]
		current := intervals[i]

		if current[0] <= prev[1] {
			prev[0] = min(prev[0], current[0])
			prev[1] = max(prev[1], current[1])
		} else {
			lists = append(lists, intervals[i])
		}
	}
	return lists
}
