package main

import (
	"log"
	"math/big"
)

func bigPrime() {
	p := big.NewInt(2)
	p.Exp(p, big.NewInt(1398269), nil)
	p.Sub(p, big.NewInt(1))
	// Get ready to scroll
	log.Printf("a big prime number is %s", p)
	// Takes a while
	// log.Printf("2^1,398,269-1 is probably prime: %t", p.ProbablyPrime(1))
}

func ScientificNotation(coefficient, exponent int64) *big.Int {
	exp := big.NewInt(10)
	exp = exp.Exp(exp, big.NewInt(exponent), nil)
	coeff := big.NewInt(coefficient)
	return coeff.Mul(coeff, exp)
}

func astrophysics() {
	age := ScientificNotation(43, 16)
	log.Printf("the universe is about %s seconds old", age)
	size := ScientificNotation(88, 25)
	log.Printf("the universe is about %s light years across", size)
	stars := ScientificNotation(5, 22)
	log.Printf("the universe has about %s stars", stars)
	galaxies := ScientificNotation(125, 9)
	log.Printf("the universe has about %s galaxies", galaxies)
}

func primeList() {
	var primesFound int
	two := big.NewInt(2)
	p := big.NewInt(3)
	for {
		if p.ProbablyPrime(1) {
			primesFound++
			log.Printf("%s is a prime number", p)
		}

		if primesFound > 100 {
			break
		}

		p.Add(p, two)
	}
}

func mul() {
	x, _ := new(big.Int).SetString("7612058254738945", 10)
	y, _ := new(big.Int).SetString("9263591128439081", 10)
	z := new(big.Int).Mul(x, y)
	log.Printf("%s x %s = %s", x, y, z)
}

func gcd() {
	a, _ := new(big.Int).SetString("7612058254738945", 10)
	b, _ := new(big.Int).SetString("9263591128439081", 10)
	z := new(big.Int).GCD(nil, nil, a, b)
	log.Printf("the GCD of %s and %s is %s", a, b, z)
}

var one = big.NewRat(1, 1)

func f(i *big.Rat, depth uint64) *big.Rat {
	if depth == 0 {
		return one
	}

	// Doing this is slightly faster
	// than the recursive version.
	c := make(chan *big.Rat, 1)
	go func() {
		n := new(big.Rat).Set(i)
		c <- f(n.Add(n, one), depth-1)
	}()

	num := new(big.Rat).Set(i)
	denom := big.NewRat(2, 1)
	denom = denom.Mul(denom, num)
	denom = denom.Add(denom, one)

	rest := new(big.Rat)
	rest = rest.Mul(num, denom.Inv(denom))
	rest = rest.Mul(rest, <-c)

	ret := big.NewRat(1, 1)
	return ret.Add(ret, rest)
}

func pi() {
	value := f(big.NewRat(1, 1), 500)
	value.Mul(value, big.NewRat(2, 1))
	log.Printf("pi is %s", value.FloatString(100))
}

func main() {
	bigPrime()
	astrophysics()
	primeList()
	mul()
	gcd()
	pi()
}
