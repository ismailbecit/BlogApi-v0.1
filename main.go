package main

import (
	"blogapi/api/config"
	"blogapi/api/router"
	"blogapi/repository"
)

func main() {
	config.Init()
	repository.Set()
	router.Router()
}
