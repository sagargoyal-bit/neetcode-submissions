func findDuplicate(nums []int) int {
    // using flodd cycle detection
slow:=nums[0]
fast:=nums[0]

//intersection point 
for {
	slow =nums[slow]
	fast= nums[nums[fast]]
	if fast==slow{
		break
	}
}
//find starting point
slow =nums[0]
 for slow !=fast{
	fast=nums[fast]
	slow=nums[slow]
 }
 return slow

}