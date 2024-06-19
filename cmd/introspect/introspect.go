package introspect

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/comsma/knead/app"
	"github.com/comsma/knead/pkg/db/mssql"
	"github.com/comsma/knead/pkg/orm/bun"
	"github.com/spf13/cobra"
	"os"
	"time"
)

const (
	databaseFlag = "database"
	ormFlag      = "orm"
	configFlag   = "config"
)

func NewCmd() *cobra.Command {
	var databaseType string
	var ormType string
	var configPath string

	var cmd = &cobra.Command{
		Use:   "introspect",
		Short: "Introspect database and generate models",

		Run: func(cmd *cobra.Command, args []string) {
			db := mssql.NewDatabase()

			conf, err := app.ReadConfig(configPath)
			if err != nil {
				fmt.Println(err)
			}

			db.Connect(conf)

			os.MkdirAll("gen-test", 0777)

			spinnerSet := []string{"\u2588", "\u2580", "\u2584", "\u2588"}
			s := spinner.New(spinnerSet, 100*time.Millisecond)
			s.Start()
			defer s.Stop()

			for _, t := range conf.Tables {
				func(t string) {
					fi, err := os.Create(fmt.Sprintf("gen-test/%s.go", t))
					defer fi.Close()
					table, err := db.GetTableInfo(t)

					generator := bun.Generator{}

					err = generator.WriteFile(fi, table)
					if err != nil {
						fmt.Println(err)
					}
				}(t)

			}

		},
	}
	cmd.Flags().StringVarP(&databaseType, databaseFlag, "d", "", "Database to introspect")
	cmd.Flags().StringVarP(&ormType, ormFlag, "o", "", "ORM to generate models for")
	cmd.Flags().StringVarP(&configPath, configFlag, "c", "config.yml", "Config file to use")

	cmd.RegisterFlagCompletionFunc(databaseFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"mssql"}, cobra.ShellCompDirectiveDefault
	})
	cmd.RegisterFlagCompletionFunc(ormFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"bun"}, cobra.ShellCompDirectiveDefault
	})

	return cmd
}
