package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "SuperAdminSeeder" {
		SuperAdminSeeder()
	} else if flag.Arg(0) == "AboutMeSeeder" {
		AboutMeSeeder()
	} else {
		fmt.Println("Unknown command")
	}
}
