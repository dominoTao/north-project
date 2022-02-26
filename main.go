package main

import (
	"fmt"
	"north-project/routers"
)

func main() {
	r := routers.SetupRouters()

	if err := r.Run("localhost:8080"); err != nil {
		fmt.Errorf("failed start routers")
		return
	}
}
