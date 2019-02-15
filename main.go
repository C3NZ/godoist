package main

// We import go-homedir as homedir
import (
    "fmt"
    "os"
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
    must(db.Init(dbPath))

    // Execute the root command (Describe our app)
    must(cmd.RootCmd.Execute());
}

// Error handling function for somewhat useful error outputs
func must(err error) {
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}
