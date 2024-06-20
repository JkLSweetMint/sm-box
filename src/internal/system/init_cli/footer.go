package init_cli

import (
	"context"
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
			Use: "init",
		}

		// Настройка root команды
		{
			var cmd = &cobra.Command{
				Use:    "completion",
				Hidden: true,
			}

			root.AddCommand(cmd)
		}

		// initial
		{
			var cmd = &cobra.Command{
				Use:   "initial",
				Short: "Initial initialization of the system. ",
				Long: `
Initial initialization of the system. 
`,
				Args: cobra.NoArgs,
				Run: func(cmd *cobra.Command, args []string) {
					// Очистка (если требуется)
					{
						if ok, err := cmd.Flags().GetBool("clear"); err != nil {
							cli_.components.Logger.Error().
								Format("An error occurred while executing the command: '%s'. ", err).Write()
							return
						} else if ok {
							if cErr := cli_.controllers.Initialization.Clear(ctx); cErr != nil {
								cli_.components.Logger.Error().
									Format("An error occurred while executing the command: '%s'. ", cErr).Write()
								return
							}
						}
					}

					if cErr := cli_.controllers.Initialization.Initialize(ctx); cErr != nil {
						cli_.components.Logger.Error().
							Format("An error occurred while executing the command: '%s'. ", cErr).Write()
						return
					}
				},
			}

			cmd.Flags().Bool("clear", false, "cleaning the system before initialization")

			root.AddCommand(cmd)
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
