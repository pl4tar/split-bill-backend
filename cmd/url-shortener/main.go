package main

import (
	"fmt"
	"split-bill-backend/internal/config"
)

func main() {
	cfg := config.MustLoad()

	// TODO DELETE THIS!!!
	fmt.Println(cfg)

}
