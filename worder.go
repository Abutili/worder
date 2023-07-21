package main

import (
	"log"
	"os"

	"abutili.com/worder/cmd/worder"
	//w "abutili.com/worder/pkg/worder"
	//"github.com/spf13/cobra"
	//"github.com/spf13/pflag"
)

func main() {
	// set all logging to stdout for now
	log.SetOutput(os.Stdout)
	log.SetOutput(os.Stderr)

	// call into root Cobra command
	worder.Execute()
}
