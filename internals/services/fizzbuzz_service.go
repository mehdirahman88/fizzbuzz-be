package fizzbuzzService

import "github.com/mehdirahman88/fizzbuzz-be/config"

func CalcFizzBuzz(i int) string {
	fizz := config.Config("FIZZ")
	buzz := config.Config("BUZZ")

	if i%3 == 0 && i%5 == 0 {
		return fizz + buzz
	} else if i%3 == 0 {
		return fizz
	} else if i%5 == 0 {
		return buzz
	}

	return ""
}
