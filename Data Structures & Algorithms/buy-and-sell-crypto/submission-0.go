func maxProfit(prices []int) int {
	l:=0
	r:=1
	maxPrice:=0
	
	for r<len(prices){
		if prices[r]> prices[l]{
			price:=prices[r]-prices[l]
			if price>maxPrice{
				maxPrice=price
			}
		}else{
			l=r
		}
		r++
	}
	return maxPrice
}
