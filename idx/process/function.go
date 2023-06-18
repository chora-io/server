package process

import (
	"context"

	"github.com/choraio/server/idx/client"
)

// Function is the function used to advance the process.
type Function func(ctx context.Context, c client.Client) error
