package upgrade

import "github.com/honorjoey/goctls/internal/cobrax"

// Cmd describes an upgrade command.
var Cmd = cobrax.NewCommand("upgrade", cobrax.WithRunE(upgrade))
