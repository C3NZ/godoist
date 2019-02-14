package cmd

import "github.com/spf13/cobra"
import "fmt"
import "strings"

var addCmd = &cobra.Command{
    Use: "add",
    Short: "Add a task to the to do list",
    Run: func(cmd *cobra.Command, args []string) {
        task := strings.Join(args, " ")
        fmt.Printf("Added \"%s\" to your task list \n", task)
    },
}

// Go initializer function called to setup anything before execution
func init() {
    // Register the addCmD as a sub command
    RootCmd.AddCommand(addCmd)
}
