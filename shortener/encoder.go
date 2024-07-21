package shortener

const base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var counter uint64 =  56800235584


//increases the counter and returns its equivalent base62 encoding
func base62() (base62 string) {
	
	base62 = ""
	counter++
	temp := counter
	
	for temp > 0 {
		remainder := temp % 62
		base62 = string(base62Digits[remainder]) + base62
		temp /= 62

	}

	return
}