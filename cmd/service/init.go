package main

import (
	"github.com/dsc-sgu/atcc/internal/config"
	"github.com/dsc-sgu/atcc/internal/log"
)

func init() {
	config.C = config.Load("config/atcc.yaml")

	log.Init()
}
