package project_cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// exec - внутренний метод для запуска CLI.
func (cli_ *cli) exec(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	// Завершение работы
	{
		defer func() {
			var failed = err != nil

			if err = cli_.core.Shutdown(); err != nil {
				cli_.components.Logger.Error().
					Format("An error occurred when starting maintenance of the '%s': '%s'. ",
						env.Vars.SystemName,
						err).Write()
				return
			}

			if failed {
				os.Exit(1)
			}
		}()
	}

	cli_.components.Logger.Info().
		Format("Starting '%s' CLI... ", env.Vars.SystemName).Write()

	// Команды
	{
		var root = &cobra.Command{
			Use: "project",
		}

		// Настройка root команды
		{
			var cmd = &cobra.Command{
				Use:    "completion",
				Hidden: true,
			}

			root.AddCommand(cmd)
		}

		// create
		{
			var projectVersion string

			var cmd = &cobra.Command{
				Use:   "create [project name] [owner login]",
				Short: "Creating a new project in the system. ",
				Long: `
Creating a new project in the system without an owner, 
with initial configuration. 
`,
				Args: cobra.MinimumNArgs(2),
				Run: func(cmd *cobra.Command, args []string) {
					var (
						projectName = args[0]
						ownerLogin  = args[1]
					)

					fmt.Printf("Project name: %s\n", projectName)
					fmt.Printf("Project version: %s\n", projectVersion)
					fmt.Printf("Owner login: %s\n", ownerLogin)
				},
			}

			cmd.Flags().StringVarP(&projectVersion, "version", "v", "", "specifying the project version")
			root.AddCommand(cmd)
		}

		// list/ls
		{
			var cmd1 = &cobra.Command{
				Use:   "list",
				Short: "Getting a list of projects. ",
				Long: `
Getting a list of projects.
`,
				Args: cobra.NoArgs,
				Run: func(cmd *cobra.Command, args []string) {

				},
			}

			var cmd2 = &cobra.Command{
				Use:   "ls",
				Short: cmd1.Short,
				Long:  cmd1.Long,
				Args:  cmd1.Args,
				Run:   cmd1.Run,
			}

			root.AddCommand(cmd1)
			root.AddCommand(cmd2)
		}

		// remove/rm
		{
			var cmd1 = &cobra.Command{
				Use:   "remove [project id]",
				Short: "Deleting a project. ",
				Long: `
Deleting a project.
`,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {
					var projectID = args[0]

					fmt.Printf("Project id: %s\n", projectID)
				},
			}

			var cmd2 = &cobra.Command{
				Use:   "rm",
				Short: cmd1.Short,
				Long:  cmd1.Long,
				Args:  cmd1.Args,
				Run:   cmd1.Run,
			}

			root.AddCommand(cmd1)
			root.AddCommand(cmd2)
		}

		if err = root.Execute(); err != nil {
			cli_.components.Logger.Error().
				Format("The command could not be executed: '%s'. ",
					err).Write()
			return
		}
	}

	cli_.components.Logger.Info().
		Format("The '%s' CLI has been successfully finished. ", env.Vars.SystemName).Write()

	return
}
