package main

import (
	"fmt"
	"github.com/taction/wit-dapr/dapr/state"
)

func main() {
	state.DaprStateStateInterfaceSet("state", state.DaprStateStateTypesSetRequest{Key: "z", Value: []byte("value")})
	res := state.DaprStateStateInterfaceGet("state", state.DaprStateStateTypesGetRequest{Key: "z"})
	if res.IsOk() {
		fmt.Println(string(res.Val.Data))
	} else {
		fmt.Println(res.Err)
	}
}
