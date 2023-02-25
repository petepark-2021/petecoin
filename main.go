package main

import (
	"github.com/petepark-2021/petecoin/cli"
	"github.com/petepark-2021/petecoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
