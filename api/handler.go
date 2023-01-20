package api

import (
	"github.com/FaridehGhani/prancer_test/deliverpoint/application"
	"github.com/FaridehGhani/prancer_test/deliverpoint/delivery"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// Agents FIXME: global variable, bad practice
var Agents []*delivery.Agent

type apiHandler struct{}

func (api apiHandler) DeliverPoint(ctx *gin.Context) {
	log.Printf("initiated agents %v\n", Agents)

	requestBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Printf("error in reading request body %v\n", err)
	}

	var point Point
	if err = json.Unmarshal(requestBody, &point); err != nil {
		log.Printf("error in unmarshalling request body %v\n", err)
	}
	log.Printf("requested point %v\n", point)

	wg := new(sync.WaitGroup)
	go application.DeliverPoint(wg, Agents, PointToLocation(point))

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "point is sent for processing",
	})
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func PointToLocation(p Point) delivery.Location {
	return delivery.Location{
		X: p.X,
		Y: p.Y,
	}
}
