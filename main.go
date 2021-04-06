package main

import (
	"fmt"
	appRoute "github.com/rayonnunes/imersaofsfc-simulator/app/route"
)

func main() {
	route := appRoute.Route{
		ID: "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[0])
}