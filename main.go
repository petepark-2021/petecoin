package main

import (
	"github.com/petepark-2021/petecoin/explorer"
	"github.com/petepark-2021/petecoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4004)
}
