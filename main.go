package main

// We import go-homedir as homedir
import (
    "path/filepath"
    "godoist/cmd"
    "godoist/db"

    homedir "github.com/mitchellh/go-homedir"
)

func main() {
    // Obtain the homedir path for the user currently running this program
    home, _ := homedir.Dir()

    // Join the home directory with our desired file path
    dbPath := filepath.Join(home, "tasks.db")
    
    // Instantiate our db connection/buckets
    err:= db.Init(dbPath)

    // Crash if setting up our database fails
    if err != nil {
        panic(err)
    }
    // Execute the root command (Describe our app)
    cmd.RootCmd.Execute();
}
