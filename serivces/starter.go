package main

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
)

func main() {
	wsAllMarketsStatHandler := func(event binance.WsAllMarketsStatEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsAllMarketsStatServe(wsAllMarketsStatHandler,errHandler)
	if err != nil {
		panic(err)
	}
	<-doneC
}
