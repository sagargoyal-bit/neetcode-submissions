func lengthOfLongestSubstring(s string) int {
	left:=0
	maxLength:=0
	hasseen:=make(map[byte]int)

	for right:=0;right<len(s);right++{
			ch:=s[right]
			index,exists:=hasseen[ch];
			if exists{
				left=max(left,index+1)
			}
		
			hasseen[ch]=right

			currentLength:=(right-left)+1
			if currentLength >maxLength{
				maxLength=currentLength
			}

	}
	return maxLength
}
