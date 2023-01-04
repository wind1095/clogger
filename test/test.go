package main

import (
	"github.com/wind1095/clogger"
	"time"
)

func main() {
	clogger.Init("./log", "cdr", 45*24*time.Hour)

	clogger.Info("test")
	clogger.Error("test2")
}
