package string

func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	a1, a2, res := []byte(num1), []byte(num2), make([]int, len(num1)+len(num2))
	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a2); j++ {
			res[i+j+1] += int(a1[i]-'0') * int(a2[j]-'0')
		}
	}
	for i := len(res) - 1; i > 0; i-- {
		res[i-1] += res[i] / 10
		res[i] %= 10

	}
	if res[0] == 0 {
		res = res[1:]
	}
	ans := make([]byte, len(res))
	for i := 0; i < len(res); i++ {
		ans[i] = byte(res[i]) + '0'
	}
	return string(ans)
}
