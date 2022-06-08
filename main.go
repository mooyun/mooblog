package main

import (
	"mooblog/model"
	"mooblog/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
}
