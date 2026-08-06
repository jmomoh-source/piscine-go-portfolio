package piscine

func FindPrevPrime(nb int) int {
    // Your code here
	if nb < 2 {
		return 0
	} 

	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	if n <= 1 {
		return false 
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false 
		}
	}
	return true
}


// func FindPrevPrime(nb int) int {
// 	for i := nb; i >= 2; i-- {
// 		prime := true
// 		for j := 2; j < i; j++ {
// 			if i%j == 0 {
// 				prime = false
// 				break
// 			}
// 		}
// 		if prime {
// 			return i
// 		}
// 	}
// 	return 0
// }