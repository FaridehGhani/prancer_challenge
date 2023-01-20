package delivery

import (
	"fmt"
	"log"
	"math"
	"time"
)

type Agent struct {
	ID       int
	Location Location
	Status   AgentStatus
}

type Location struct {
	X float64
	Y float64
}

type AgentStatus int

const (
	_ AgentStatus = iota
	AgentStatusAvailable
	AgentStatusBusy
)

func (a *Agent) DistanceToPoint(location Location) float64 {
	return math.Sqrt(math.Pow(a.Location.X-location.X, 2) + math.Pow(a.Location.Y-location.Y, 2))
}

func (a *Agent) IsAvailable() bool {
	fmt.Println("agent status ", a.ID, a.Status)
	if a.Status == AgentStatusAvailable {
		return true
	}
	return false
}

func (a *Agent) UpdateStatus(status AgentStatus) {
	a.Status = status
}

func (a *Agent) UpdateLocation(location Location) {
	a.Location = location
}

func (a *Agent) ProcessPoint(point Location) {
	a.UpdateStatus(AgentStatusBusy)

	l := math.Sqrt(math.Pow(point.Y, 2) + math.Pow(point.X, 2))
	for i := l; i > 0; i-- {
		log.Printf("agent %v is moving, %v distance remains to finish \n", a.ID, i)
		time.Sleep(1 * time.Second)
	}
	a.UpdateStatus(AgentStatusAvailable)
	a.UpdateLocation(point)

	log.Printf("Agent %v reached to point %v\n", a.ID, point)
}
