func maxArea(heights []int) int {

maxWater:=0
water:=0
left:=0
right:=len(heights)-1
for left<right{
	if heights[left]<heights[right]{
		water=(right-left) *heights[left]
		left++
	}else{
		water=(right-left) *heights[right]
		right--
	}
	if water>maxWater{
		maxWater=water
		}
}
	return maxWater
}
