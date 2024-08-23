package main

import (
	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/log"
	"github.com/dsc-sgu/shawty/internal/random"
)

func init() {
	config.Load("config/shawty.yaml")
	log.Init()
	random.Init()
}
