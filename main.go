package main

import (
	_ "github.com/lib/pq"
	"github.com/rizkyfazri23/dripay/delivery"
)

func main() {
	delivery.Server().Run()
}
