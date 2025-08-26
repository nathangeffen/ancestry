# Simulation of ancestry and genes

Let's say an animal or plant  population starts off with two individuals. The
parents die and children mate and produce four children. The four children and
produce two children each. After $n$ generations there are $2^n$ individuals.
Many of the individuals will only share the original two individuals as
ancestors. 

In the real world ths quickly becomes unsustainable. Animal populations
fluctuate and cannot grow uncontrollably else they run out of habitat and
food.

This raises some questions:

- How many generations back are the most distantly related individuals in a
population likely to find a common ancestor?

- How many common ancestor do any two individuals have on average?

I wrote this simulation to try to answer these and similar questions.

While doing the initial simulation, further interesting questions arose: what
happens to genetic diversity over time and what mutation rate is needed to
maintain genetic diversity?

## Installation

Download the repo. E.g. *git clone https://github.com/nathangeffen/ancestry*

Run *go build*

The executable is called *ancestry*.

## How it works

The simulation starts off with a specified number of agents in the first
generation. 

It then repeatedly matches pairs of agents to produce agents
(children) for the next generation. A growth rate specifies how many agents
need to be created in the next generation.

The above repeats for a specified number of generations.

There are also parameters specifying genes and a mutation rate.

Additional parameters control whether siblings, cousin and agent sex must be
taken into account when reproducing.

The simulation outputs four different kinds of analysis at the end. 

## Command line options

To get a list of command line options run:

    ./ancestry -h

The main ones are:

- agents: An integer specifying the number of agents to start off withto start
off with.  (default 100)
- generations: Integer indicating the generations to run for (default 4)
- growth: Real number indicating the growth rate of population (default 1.01)
- analysis This tells the simulation what analyses to carry out. There are four
analyses. The letters N, C, D and G represents each one. N - Average ancestors
per agent C - Average common ancestors per agent D - Generation differences G -
Gene analysis g - Only do gene analysis on last generation (default "NCDGg")    
- genes: Integer indicating the number of genes per agent in initial generation
(default 10)
- mutation: Real number indicating the gene mutation rate
- compatible: A boolean indicating whether to do any agent pairing
compatibility checks. For fastest, least complicated results set this to false.
I'm not entirely satisfied yet with the way the simulation handles partner
selection when compatibility is enforced.

There are two matching algorithms. One assumes monogamous partnerships, i.e.
given any agent, it has zero or more children with at most one other agent.
This can be selected with the *-monog=true* command line parameter. This
parameter defaults to false. I'm not entirely satisfied with the implementation
of the monogamous algorithm and recommend that this parameter be left set to false.


## Sample results

TO DO

