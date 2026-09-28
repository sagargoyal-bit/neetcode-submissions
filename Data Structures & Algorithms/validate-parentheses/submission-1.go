func isValid(s string) bool {
   stack:=[]byte{}
	for i:=0;i<len(s);i++{
		ch:=s[i]
		
		if ch =='('||ch=='{'||ch=='['{
			stack=append(stack,ch)
			continue
		}
		//if no opening brackets
		if len(stack)==0{
			return false
		}
		//check for closing brackets

		pop :=stack[len(stack)-1]
		if ch ==')' && pop !='('{
			return false
		}
		if ch =='}' && pop !='{'{
			return false
		}
		if ch ==']' && pop !='['{
			return false
		}

		stack=stack[:len(stack)-1]
	}
	return len(stack)==0
}
