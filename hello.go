package main

import "fmt"
//Traigo el package quote de https://pkg.go.dev/
//Una vez que hago el import, tengo que hacer go mod tidy
// para que lo localice y lo traiga 
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Opt())
}