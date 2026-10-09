func characterReplacement(s string, k int) int {
left:=0
maxlength:=0
maxrep:=0
seen:=make(map[byte]int)
	for right:=0;right<len(s);right++ {
		ch:=s[right]
		seen[ch]=seen[ch]+1
		if seen[ch]>maxrep{
			maxrep=seen[ch]
		}

		for (right-left+1)-maxrep>k{
			seen[s[left]]--
			left++
		}
		//calculate the length
				length:=(right-left)+1
				if length>maxlength{
					maxlength=length
				}
		
		}

	return maxlength
}
