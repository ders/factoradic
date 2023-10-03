# Factoradic

[![Go Reference](https://pkg.go.dev/badge/github.com/ders/factoradic.svg)](https://pkg.go.dev/github.com/ders/factoradic)

Factoradic provides support for [factorial numbers](https://xkcd.com/2835/).

## Quickstart

```
var n Number = 23
fmt.Printf("Your number is %s.", n) // "Your number is 321."
```

```
n, _ := ParseNumber("321")
fmt.Printf("Your number in base 10 is %d.", n) // "Your number in base 10 is 23."
```
