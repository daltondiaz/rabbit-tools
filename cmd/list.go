package cmd

import (
	"fmt"

	"github.com/daltondiaz/rabbit-tools/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use: "list",
    Short: "List of Queues",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        queue := args[0]
        result := internal.GetQueues(queue)
        totalItems := 0.0
        for queueName, value := range result {
            fmt.Println("Queue: ", queueName, " Total: ", value)
            totalItems += value
        }
        fmt.Println("Total of Queues: ", len(result))
        fmt.Println("Total of Items on queues: ", totalItems)
    },
}

func init(){
    rootCmd.AddCommand(listCmd)
}
