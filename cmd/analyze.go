package cmd

import (
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use: "analyze",
	Run: func(cmd *cobra.Command, args []string) {
		//bun := bun()
		//db := mssql.NewDatabase(bun)
		//
		//conf, err := app.ReadConfig()
		//if err != nil {
		//	fmt.Println(err)
		//}
		//
		//db.Connect(conf)
		//db.Analyze()

	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
