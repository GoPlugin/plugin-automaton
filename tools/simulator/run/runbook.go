package run

import (
	"os"

	"github.com/goplugin/plugin-automaton/tools/simulator/config"
)

func LoadSimulationPlan(path string) (config.SimulationPlan, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return config.SimulationPlan{}, err
	}

	return config.DecodeSimulationPlan(data)
}
