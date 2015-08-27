package lib

func Primesieve(limit int) []bool {
	primes := make([]bool, limit)

	for i := 2; i < limit; i++ {
		primes[i] = true
	}

	for i := 2; i < limit; i++ {
		if primes[i] {
			j := i + i
			for j < limit {
				primes[j] = false
				j += i
			}
		}
	}
	return primes
}
