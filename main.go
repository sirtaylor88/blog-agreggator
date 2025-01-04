package main

import (
	"encoding/json"
	"fmt"

	"github.com/sirtaylor88/go-blog-agreggator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	if err = config.SetUser("sirtaylor88", cfg); err != nil {
		fmt.Println(err)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	data, _ := json.MarshalIndent(cfg, "", "\t")
	fmt.Println(string(data))
}
