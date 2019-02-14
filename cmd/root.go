package cmd

import "github.com/spf13/cobra"

// Defining a cobra command
// Use - the command use
// Short - A short description of your command
// Long - A long description of your command
// Run - What is used to run your command(?)
var RootCmd = &cobra.Command{
    Use:   "godoist",
    Short: "godoist is a CLI task manager",
}
