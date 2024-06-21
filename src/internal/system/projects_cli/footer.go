package projects_cli

import (
	"context"
	"fmt"
	terminal_table "github.com/aquasecurity/table"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"sm-box/internal/common/entities"
	error_list "sm-box/internal/common/errors"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	"strconv"
	"strings"
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
			var projectVersion, projectDescription string

			var cmd = &cobra.Command{
				Use:   "create [project name]",
				Short: "Creating a new project in the system. ",
				Long: `
Creating a new project in the system without an owner, 
with initial configuration. 
`,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {
					var name, version, description string

					// Сбор аргументов
					{
						name = args[0]
						version = projectVersion
						description = projectDescription
					}

					if cErr := cli_.controllers.Projects.Create(context.Background(), name, description, version); cErr != nil {
						cli_.components.Logger.Error().
							Format("An error occurred while executing the command: '%s'. ", cErr).Write()

						fmt.Println(cErr.Message())
						return
					}
				},
			}

			cmd.Flags().StringVarP(&projectVersion, "version", "v", "", "specifying the project version")
			cmd.Flags().StringVarP(&projectDescription, "description", "d", "", "specifying the project description")

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
					var table = terminal_table.New(os.Stdout)

					table.SetDividers(terminal_table.UnicodeRoundedDividers)
					table.SetLineStyle(terminal_table.StyleBrightBlack)

					table.SetHeaders("id", "uuid", "name", "version", "owner", "description")
					table.AddHeaders("id", "uuid", "name", "version", "id", "name", "description")

					table.SetHeaderColSpans(0, 1, 1, 1, 1, 2, 1)
					table.SetAutoMergeHeaders(true)

					var projects []*entities.Project

					// Получение проектов
					{
						var cErr c_errors.Error

						if projects, cErr = cli_.controllers.Projects.GetAll(context.Background()); cErr != nil {
							cli_.components.Logger.Error().
								Format("An error occurred while executing the command: '%s'. ", cErr).Write()

							fmt.Println(cErr.Message())
							return
						}
					}

					for _, project := range projects {
						var items = make([]string, 0, 7)

						items = append(items, strconv.Itoa(int(project.ID)))
						items = append(items, project.UUID.String())
						items = append(items, project.Name)
						items = append(items, project.Version)

						items = append(items, strconv.Itoa(int(project.Owner.ID)))
						items = append(items, project.Owner.Username)

						items = append(items, project.Description)

						table.AddRow(items...)
					}

					table.Render()
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

		// env
		{
			var setEnv string

			var cmd = &cobra.Command{
				Use:   "env [project id or uuid]",
				Short: "Project Environment Management. ",
				Long: `
Project Environment Management.
`,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {
					var id string

					// Сбор аргументов
					{
						id = args[0]
					}

					// Установить значение (если такой флаг указан)
					{
						if setEnv != "" {
							if matched, _ := regexp.MatchString(`^([0-9a-zA-Z_]+=[\s\S]+)$`, setEnv); matched {
								var (
									splitFlag = strings.Split(setEnv, "=")
									key       = splitFlag[0]
									value     = splitFlag[1]
								)

								if cErr := cli_.controllers.Projects.SetEnv(context.Background(), id, key, value); cErr != nil {
									cli_.components.Logger.Error().
										Format("An error occurred while executing the command: '%s'. ", cErr).Write()

									fmt.Println(cErr.Message())
									return
								}
							} else {
								var cErr = error_list.InvalidFlag()

								cli_.components.Logger.Error().
									Format("An error occurred while executing the command: '%s'. ", cErr).Write()

								fmt.Println(cErr.Message())
								return
							}

							return
						}
					}

					// Отображение окружения
					{
						var (
							env  entities.ProjectEnv
							cErr c_errors.Error
						)

						// Получение данных
						{
							if env, cErr = cli_.controllers.Projects.GetEnv(context.Background(), id); cErr != nil {
								cli_.components.Logger.Error().
									Format("An error occurred while executing the command: '%s'. ", cErr).Write()

								fmt.Println(cErr.Message())
								return
							}
						}

						// Отображение
						{
							var table = terminal_table.New(os.Stdout)

							table.SetDividers(terminal_table.UnicodeRoundedDividers)
							table.SetLineStyle(terminal_table.StyleBrightBlack)

							table.SetHeaders("ENVIRONMENT")
							table.AddHeaders("KEY", "VALUE")

							table.SetHeaderColSpans(0, 2)

							for _, v := range env {
								table.AddRow(v.Key, v.Value)
							}

							table.Render()
						}
					}
				},
			}

			cmd.Flags().StringVarP(&setEnv, "set", "s", "", "set the value of the environment variable")

			root.AddCommand(cmd)
		}

		// remove/rm
		{
			var cmd1 = &cobra.Command{
				Use:   "remove [project id or uuid]",
				Short: "Deleting a project. ",
				Long: `
Deleting a project.
`,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {
					var id string

					// Сбор аргументов
					{
						id = args[0]
					}

					if cErr := cli_.controllers.Projects.Remove(context.Background(), id); cErr != nil {
						cli_.components.Logger.Error().
							Format("An error occurred while executing the command: '%s'. ", cErr).Write()

						fmt.Println(cErr.Message())
						return
					}
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
