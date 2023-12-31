package cmd

import (
	"log"

	"github.com/daltondiaz/rabbit-tools/internal"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of Queues",
	Run: func(cmd *cobra.Command, args []string) {
		filter, err:= cmd.Flags().GetString("filter")
        if  err != nil {
            filter = "all"
        }

        env, err := cmd.Flags().GetString("env")
        if err != nil {
            log.Fatal(err)
        }  
            
		result := internal.GetQueues(filter, env)
        var prettyResult internal.PrettyModelResult
        prettyResult.Title = "List of Queues"
        prettyResult.Header = table.Row{"Queue", "Messages"} 
        totalItems := 0.0
		for queueName, value := range result {
			row := table.Row{queueName, value}
			totalItems += value
			prettyResult.Contents = append(prettyResult.Contents, row)
		}
        prettyResult.Footers = append(prettyResult.Footers, table.Row{"Total Queues", len(result)})
        prettyResult.Footers = append(prettyResult.Footers, table.Row{"Total Messages", totalItems})
        internal.PrettyResult(prettyResult)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
    listCmd.PersistentFlags().String("env", "", "Prefix of environment in config.env")
    listCmd.PersistentFlags().String("filter", "", "Filter queues")
}
