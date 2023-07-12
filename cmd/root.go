package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use: "rabbit",
    Short: "This is the first commad",
    Long: `A longer description 
        for the first command`,
    Run: func(cmd *cobra.Command, args []string){
    },
}

func Execute(){
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}
