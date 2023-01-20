package delivery

import (
	"fmt"
	"log"
	"math"
	"sync"
)

type Center struct {
	Agents []Agent
}

func InitiateAgents(count int) []*Agent {
	var initiatedAgents []*Agent

	for i := 0; i < count; i++ {
		agent := &Agent{
			ID: i,
			Location: Location{
				X: 0,
				Y: 0,
			},
			Status: AgentStatusAvailable,
		}
		initiatedAgents = append(initiatedAgents, agent)
	}

	return initiatedAgents
}

func AccessAgent(wg *sync.WaitGroup, agents []*Agent, point Location) {
	var availableAgents []Agent

	for _, agent := range agents {
		wg.Add(1)
		go func(a Agent) {
			defer wg.Done()
			if available := a.IsAvailable(); available {
				availableAgents = append(availableAgents, a)
			}
		}(*agent)
	}
	wg.Wait()
	log.Printf("available agents %v\n", availableAgents)

	if len(availableAgents) > 0 {
		minDistanceAgentToPoint := math.MaxFloat64
		var minDistanceAgent Agent

		for _, agent := range availableAgents {
			wg.Add(1)
			go func(a Agent) {
				defer wg.Done()
				distance := a.DistanceToPoint(point)
				fmt.Printf("distance %v\n", distance)
				if minDistanceAgentToPoint > distance {
					minDistanceAgentToPoint = distance
					minDistanceAgent = a
				}
			}(agent)
		}
		wg.Wait()
		fmt.Printf("selected agent %v\n", minDistanceAgent)

		minDistanceAgent.ProcessPoint(point)
	} else {
		// TODO: implement queue for not processed points
	}
}
