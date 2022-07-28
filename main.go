package main

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/rudikurniawan99/go-api/src"
)

func main() {
	var loadOnce sync.Once
	godotenv.Load()

	loadOnce.Do(func() {
		src.Init()
	})
}
