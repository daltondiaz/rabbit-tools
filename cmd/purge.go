package cmd

import (
	"github.com/daltondiaz/rabbit-tools/internal"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
    Use: "purge",
    Short: "Purge on list",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        queue := args[0]
        internal.PurgeQueue(queue)
    },
}

func init(){
    rootCmd.AddCommand(purgeCmd)
}
