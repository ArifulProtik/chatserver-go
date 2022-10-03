package main

import (
	"fmt"

	"github.com/ArifulProtik/chatserver-go/config"
)

func main() {
	cfg, err := config.New("./", "config")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.AppInfo.Name)
}
