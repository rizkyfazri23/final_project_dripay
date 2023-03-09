package main

import (
	"github.com/rizkyfazri23/dripay/delivery"
	_ "github.com/lib/pq"
)

func main() {
	// Run the server
	delivery.Server().Run()
}
