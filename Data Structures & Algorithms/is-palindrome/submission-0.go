func isPalindrome(s string) bool {

 sample:=[]byte{}
s = strings.ToLower(s)
 for i:=0;i<len(s);i++{
	if (s[i]>='A' && s[i]<='Z') ||(s[i]>='a' && s[i]<='z')  ||(s[i]>='0' && s[i]<='9'){
		sample=append(sample,s[i])
	}
 }
 left:=0
 right:=len(sample)-1
 for left<right{
	if sample[left]==sample[right]{
		left++
		right--
	}else{
		return false
	}	
 }
 return true
}
