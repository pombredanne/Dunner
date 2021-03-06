package dunner

import (
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/leopardslab/Dunner/internal/logger"
	"github.com/leopardslab/Dunner/pkg/config"
	"github.com/leopardslab/Dunner/pkg/docker"
	"github.com/spf13/cobra"
	"os"
)

var log = logger.Log

func Do(_ *cobra.Command, args []string) {

	// TODO Should get the name of the Dunner file from a constant or an environment variable or config file
	var dunnerFile = ".dunner.yaml"

	configs, err := config.GetConfigs(dunnerFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, stepDefinition := range configs.Tasks[args[0]] {
		step := docker.Step{
			Task:    args[0],
			Name:    stepDefinition.Name,
			Image:   stepDefinition.Image,
			Command: stepDefinition.Command,
		}
		pout, err := step.Do()
		if err != nil {
			log.Fatal(err)
		}

		_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, *pout)
		if err != nil {
			log.Fatal(err)
		}
	}

}
