package host_state

import (
	"context"
	"github.com/tetratelabs/wazero"
)

const ModuleName = "dapr:state/state-interface"

func Instantiate(ctx context.Context, rt wazero.Runtime) error {
	s := State
	_, err := rt.NewHostModuleBuilder(ModuleName).
		NewFunctionBuilder().WithFunc(s.Get).Export("get").
		NewFunctionBuilder().WithFunc(s.Set).Export("set").
		NewFunctionBuilder().WithFunc(s.Delete).Export("delete").
		Instantiate(ctx)
	return err
}
