package main

import "github.com/taction/wit-dapr/dapr/state"

func init() {

}

func main() {
	//state.DaprStateStateInterfaceSet("state", state.DaprStateStateTypesSetRequest{Key: "z", Value: []byte("value")})
	state.DaprStateStateInterfaceGet("state", state.DaprStateStateTypesGetRequest{Key: "z"})
}
