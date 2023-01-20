package application

import (
	"github.com/FaridehGhani/prancer_test/deliverpoint/delivery"

	"sync"
)

func DeliverPoint(wg *sync.WaitGroup, agents []*delivery.Agent, location delivery.Location) {
	delivery.AccessAgent(wg, agents, location)
}
