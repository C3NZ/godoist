package cmd

import "github.com/spf13/cobra"
import "fmt"

var addCmd = &cobra.Command{
    Use: "add",
    Short: "Add a task to the to do list",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Add called")  
    },
}

func init() {
    // Register the addCmD as a sub command
    RootCmd.AddCommand(addCmd)
}
