package process

import (
	"context"
	"fmt"

	idxcontext "github.com/choraio/server/idx/context"
)

func Example(ctx context.Context, idxCtx idxcontext.Context) error {
	fmt.Println("example process function")

	return nil
}
