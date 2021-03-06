package cmd

import (
	"github.com/leopardslab/Dunner/pkg/dunner"
	"github.com/spf13/cobra"
)

type Config struct {
	Image   string    `yaml:"image"`
	Command [] string `yaml:"command"`
}

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do whatever you say",
	Long:  `You can run any task defined on the '.dunner.yaml' with this command`,
	Run:   dunner.Do,
}
