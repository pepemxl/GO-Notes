package main

import (
	"fmt"
	"time"
)

type Crypto struct {
	open    float64
	close   float64
	high    float64
	low     float64
	volume  float64
	created time.Time
}

func main() {
	var BTC Crypto
	BTC.open = 34000
	BTC.close = 35000
	BTC.high = 38000
	BTC.low = 32000
	BTC.created = time.Now()
	ETH := new(Crypto)
	ETH.open = 3400
	ETH.close = 3500
	ETH.high = 3800
	ETH.low = 3200
	//time_sleep := time.Duration(100000)
	//time.Sleep(time_sleep)
	ETH.created = time.Now()
	fmt.Println("BTC", BTC)
	fmt.Println("ETH", ETH)
	fmt.Println("BTC.open", BTC.open)
	fmt.Println("BTC.close", BTC.close)
	fmt.Println("BTC.high", BTC.high)
	fmt.Println("BTC.low", BTC.low)
	fmt.Println("BTC.created", BTC.created)
	fmt.Println("ETH.open", ETH.open)
	fmt.Println("ETH.close", ETH.close)
	fmt.Println("ETH.high", ETH.high)
	fmt.Println("ETH.low", ETH.low)
	fmt.Println("ETH.created", ETH.created)
}
