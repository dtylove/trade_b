package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"mathcing/matching"
)

func main()  {
	a, _ := decimal.NewFromString("5.5")
	b, _ := decimal.NewFromString("1.11")

	c := a.Div(b)

	fmt.Print(c)
	matching.InitEngine(4)
}