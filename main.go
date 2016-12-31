package main

import (
	"fmt"
	"github.com/rtravitz/coparkfinder/seed"
)

func main() {
	parks := seed.UnmarshalCSV()
	for _, park := range parks {
		fmt.Println(park)
	}
}
