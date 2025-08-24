package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

// Process the command line arguments and return values set in
// parameters struct.
func processFlags() (Parameters, int) {
	params := NewParameters()
	var p Parameters
	flag.IntVar(&p.SimulationId, "id", params.SimulationId, "Id of simulation")
	flag.IntVar(&p.NumAgents, "agents", params.NumAgents, "Number of agents")
	flag.IntVar(&p.Generations, "generations", params.Generations, "Number of generations to run for")
	flag.Float64Var(&p.GrowthRate, "growth", params.GrowthRate, "Growth rate of population")
	flag.BoolVar(&p.Monogamous, "monog", params.Monogamous, "Agents are monogamous")
	flag.IntVar(&p.MatingK, "matingk", params.MatingK, "Number of agents to search for compatible match")
	flag.BoolVar(&p.Compatible, "compatible", params.Compatible, "Switch off all mating compatibility checks if false")
	flag.BoolVar(&p.MateSelf, "mateself", params.MateSelf, "Agents can mate with themselves")
	flag.BoolVar(&p.MateSibling, "matesibling", params.MateSibling, "Agents can mate with siblings")
	flag.BoolVar(&p.MateCousin, "matecousin", params.MateCousin, "Agents can mate with cousins")
	flag.BoolVar(&p.MateSameSex, "matesamesex", params.MateSameSex, "Agents can mate with same sex")
	flag.IntVar(&p.NumGenes, "genes", params.NumGenes, "Number of genes per agent in initial generation")
	flag.Float64Var(&p.MutationRate, "mutation", params.MutationRate, "Gene mutation rate")
	flag.StringVar(&p.Analysis, "analysis", params.Analysis,
		`N - Number of ancestors
C - Number of common ancestors
D - Generation differences
G - Gene analysis
g - Only do gene analysis on last generation`)
	numSims := 1
	flag.IntVar(&numSims, "numsims", numSims, "Number of simulations to run (will be run in paralllel)")
	flag.Parse()
	return p, numSims
}

func main() {
	parameters, numSims := processFlags()
	var wg sync.WaitGroup
	for i := range numSims {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p := parameters
			p.SimulationId = parameters.SimulationId + i
			simulation := NewSimulation(&p)
			err := simulation.Simulate()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				return
			}
			err = simulation.Analysis()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		}()
	}
	wg.Wait()
}
