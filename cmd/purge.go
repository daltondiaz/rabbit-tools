package cmd

import (
	"log"

	"github.com/daltondiaz/rabbit-tools/internal"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
    Use: "purge",
    Short: "Purge on list",
    Run: func(cmd *cobra.Command, args []string) {
        filter, err :=  cmd.Flags().GetString("filter") 
        if err != nil {
            filter = "all"
        }
        env, err := cmd.Flags().GetString("env")
        if err != nil {
            log.Fatal(err)
        }
        internal.PurgeQueue(filter, env)
    },
}

func init(){
    rootCmd.AddCommand(purgeCmd)
    rootCmd.PersistentFlags().String("env","", "Prefix of environment in config.env")
    rootCmd.PersistentFlags().String("filter","", "Filter queue to be purge")
}
