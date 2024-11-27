package main

import (
	"gin-gorm-oj/router"
)

func main() {
	r := router.Router()

	r.Run("10.0.4.16:10003") // listen and serve on 121.4.174.222:10003
}
