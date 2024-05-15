package sodagar

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/openjogd/sodagar/cmd"
	"github.com/spf13/cobra"
)

var Version = "(untracked)"

type SodaGar struct {

	// RootCmd is the main console command
	RootCmd *cobra.Command
}

type Config struct{}

func New() *SodaGar {
	return NewWithConfig(Config{})
}
func NewWithConfig(config Config) *SodaGar {
	sg := &SodaGar{
		RootCmd: &cobra.Command{
			Use:     filepath.Base(os.Args[0]),
			Short:   "Sodagar CLI",
			Version: Version,

			// no need to provide the default cobra completion command
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		},
	}

	return sg
}

func (sg *SodaGar) Start() error {
	// register system commands
	sg.RootCmd.AddCommand(cmd.NewAdminCommand(sg))
	sg.RootCmd.AddCommand(cmd.NewServeCommand(sg))

	return sg.Execute()
}

func (sg *SodaGar) Execute() error {

	if err := sg.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return fmt.Errorf("ERROR IN EXECUTION")
	// note: leave to the commands to decide whether to print their error
}
