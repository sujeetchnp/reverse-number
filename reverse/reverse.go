package reverse

func ReverseNumber(num int64) int64 {
	var reversed int64 = 0
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10 // num = num / 10
	}
	return reversed
}
