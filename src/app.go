package src

import (
	"sync"

	"github.com/rudikurniawan99/go-api/src/utils/db"
)

func Init() {
	var loadOnce sync.Once
	loadOnce.Do(func() {
		db.Connect()
	})
}
