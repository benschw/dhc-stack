package main

import (
    "flag"
    "runtime"
)

/**
 * Implementation of `Sieve of Eratosthenes` algorithm
 * starting with first prime (2)...
 * - eliminate its multiples
 * - next un-eliminated number is the next prime
 * - (repeat)
 *
 * http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
 */
func main() {
    nCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(nCPU)

    var primes int

    flag.IntVar(&primes, "primes", 10, "prime numbers to output")
    flag.Parse()

    ch := make(chan int)
    defer close(ch)

    go Generate(ch)
    for i := 0; i < primes; i++ {
        prime := <-ch
        print(prime, "\n")
        ch1 := make(chan int)
        go Filter(ch, ch1, prime)
        ch = ch1
    }
}

func Generate(ch chan<- int) {
    for i := 2; ; i++ {
        ch <- i
    }
}

func Filter(in <-chan int, out chan<- int, prime int) {
    for {
        i := <-in
        if i%prime != 0 {
            out <- i
        }
    }
}


