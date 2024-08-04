package main

import (
	"github.com/dsc-sgu/atcc/internal/config"
	"github.com/dsc-sgu/atcc/internal/log"
	"github.com/dsc-sgu/atcc/internal/random"
)

func init() {
	config.Load("config/atcc.yaml")
	log.Init()
	random.Init()
}
