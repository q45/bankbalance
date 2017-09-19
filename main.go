package main

import (
	"fmt"
	"time"

	"github.com/q/bankbalance/bank"
)

func main() {

	// Alice:
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
	}()

	// Bob:
	go bank.Deposit(100)

	time.Sleep(time.Second * 5)
	fmt.Println(bank.Balance())

}
