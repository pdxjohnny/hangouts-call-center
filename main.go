package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/hangouts-call-center/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "hangouts-call-center"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
