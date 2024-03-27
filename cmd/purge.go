package cmd

import (
	"log"
	"strconv"

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
        greaterStr, err := cmd.Flags().GetString("greater")
        greater :=0
        if err != nil {
            log.Fatal(err)
        }  

        if value,err := strconv.Atoi(greaterStr); err != nil {
            log.Fatal(err)
        } else {
            greater = value
        }
        internal.PurgeQueue(filter, env, greater)
    },
}

func init(){
    rootCmd.AddCommand(purgeCmd)
    rootCmd.PersistentFlags().String("env","", "Prefix of environment in config.env")
    rootCmd.PersistentFlags().String("filter","", "Filter queue to be purge")
    rootCmd.PersistentFlags().String("greater","", "Filter queues with equal or greater than N messages")
}
