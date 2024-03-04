package main

import (
	"github.com/Itsu8/rest_api/initializers"
	"github.com/Itsu8/rest_api/netEngine"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	net_engine.RunNetEngine()
}
