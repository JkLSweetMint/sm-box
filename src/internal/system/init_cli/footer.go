package init_cli

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"os"
	"path"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

var (
	initFile = path.Join(env.Paths.SystemLocation, env.Paths.System.Path, ".init")
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

					cli_.components.Logger.Info().
						Format("Starting initialization '%s' CLI... ", env.Vars.SystemName).Write()

					// Проверка что уже инициализировано (существует init файл)
					{
						if _, err = os.Stat(initFile); errors.Is(err, os.ErrNotExist) {
							err = nil
						} else {
							if err == nil {
								cli_.components.Logger.Info().
									Text("The system is initialized, no reinitialization is required. ").Write()
							} else {
								cli_.components.Logger.Error().
									Format("Failed to initialize '%s' CLI: '%s'. ", env.Vars.SystemName, err).Write()
							}

							return
						}
					}

					// Логика
					{
						if err = cli_.initSystemDB(ctx); err != nil {
							cli_.components.Logger.Error().
								Format("Failed to initialize the system database: '%s'. ", err).Write()
							return
						}
					}

					// Создание init файла
					{
						if err = os.WriteFile(initFile, []byte{}, 0666); err != nil {
							cli_.components.Logger.Error().
								Format("Failed to initialize system: '%s'. ", err).Write()
							return
						}
					}

					cli_.components.Logger.Info().
						Format("The initialization '%s' CLI has been successfully finished. ", env.Vars.SystemName).Write()
				},
			}

			root.AddCommand(cmd)
		}

		// clean
		{
			var cmd = &cobra.Command{
				Use:   "clean",
				Short: "Clean initialization of the system. ",
				Long: `
Clean initialization of the system, 
the system will be returned to its original form.
`,
				Args: cobra.NoArgs,
				Run: func(cmd *cobra.Command, args []string) {
					if err = cli_.initSystemDB(ctx); err != nil {
						cli_.components.Logger.Error().
							Format("Failed to initialize the system database: '%s'. ", err).Write()
						return
					}
				},
			}

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
