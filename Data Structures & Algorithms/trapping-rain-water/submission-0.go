func trap(height []int) int {
	left:=0
	right:=len(height)-1

	maxLeft:=0
	maxRight:=0
	maxWater:=0
	for left<right{
	if height[left]<height[right]{
		if height[left]>maxLeft{
			maxLeft= height[left]
		}else{
			maxWater=maxWater+(maxLeft -height[left])
		}
		left++
	}else{
		if height[right]>maxRight{
			maxRight= height[right]
		}else{
			maxWater=maxWater+(maxRight -height[right])
		}
		right--
	}
	}
	return maxWater

}
