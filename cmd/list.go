package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/daltondiaz/rabbit-tools/internal"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of Queues",
	Run: func(cmd *cobra.Command, args []string) {
		filter, err := cmd.Flags().GetString("filter")
		if err != nil {
			filter = "all"
		}

		env, err := cmd.Flags().GetString("env")
		if err != nil {
			log.Fatal(err)
		}

		greaterStr, err := cmd.Flags().GetString("greater")
		if err != nil {
			log.Fatal(err)
		}

		if greaterStr == "" {
			greaterStr = "0"
		}

		if _, err := strconv.Atoi(greaterStr); err != nil {
			log.Fatal(fmt.Errorf("error to convert greater flag to int, error: %w", err))
		}

		greater, _ := strconv.Atoi(greaterStr)
		result := internal.GetQueues(filter, env, greater)
		var prettyResult internal.PrettyModelResult
		prettyResult.Title = "List of Queues"
		prettyResult.Header = table.Row{"Queue", "Messages"}
		totalItems := 0
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
	listCmd.PersistentFlags().String("greater", "", "Filter queues with equal or greater than N messages")
}
