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

func (c Crypto) display() {
	fmt.Println("open", c.open)
	fmt.Println("close", c.close)
	fmt.Println("high", c.high)
	fmt.Println("low", c.low)
	fmt.Println("volume", c.volume)
	fmt.Println("created", c.created)
}

func (c Crypto) getVolume() float64 {
	return c.volume
}

func main() {
	BTC := Crypto{34000, 35000, 38000, 32000, 10000, time.Now()}
	BTC.display()
	ETH := new(Crypto)
	ETH = &Crypto{3400, 3500, 3800, 3200, 2000, time.Now()}
	ETH.display()
	fmt.Println(BTC.getVolume())
	fmt.Println(ETH.getVolume())
}
