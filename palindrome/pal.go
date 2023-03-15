package palindrome

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	aux := x
	reversed := 0
	for {
		reversed = reversed*10 + aux%10
		aux = aux / 10
		if aux == 0 {
			break
		}
	}
	return x == reversed
}
