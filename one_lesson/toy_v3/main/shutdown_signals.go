package main

import (
	"os"
	"syscall"
)

var ShutDownSignals = []os.Signal{
	os.Kill,os.Interrupt,syscall.SIGKILL,
}
