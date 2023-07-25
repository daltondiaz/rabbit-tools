package cmd

import (
	"github.com/daltondiaz/rabbit-tools/internal"
	"github.com/spf13/cobra"
)

var connectionCmd = &cobra.Command {
    Use: "conn",
    Short: "List All Connections",
    Run: func(cmd *cobra.Command, args []string) {
        env, _ := cmd.Flags().GetString("env")
        internal.GetConnection(env)
    },
}

func init(){
    rootCmd.AddCommand(connectionCmd)
    connectionCmd.PersistentFlags().String("env", "", "Prefix of environment in config.env")
}
