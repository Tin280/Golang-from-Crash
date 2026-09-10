package sprint

func IsPrime(n int) bool {
	if n<=0 {
		return false
	}
	if n = 1 {
		return true
	}
	for i:=2; i<n ; i++ {
		if n%i==0 {
			return false
		}
	} 
	return true
}