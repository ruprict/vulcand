package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ruprict/vulcand/plugin/registry"
	"github.com/ruprict/vulcand/service"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := service.Run(registry.GetRegistry()); err != nil {
		fmt.Printf("Service exited with error: %s\n", err)
		os.Exit(255)
	} else {
		fmt.Println("Service exited gracefully")
	}
}
