// Package regen implements the IRI defined within Regen Ledger:
// https://github.com/regen-network/regen-ledger/tree/v5.1.2/x/data
//
// This version modifies the original to support additional configuration
// options including custom prefixes and IRI versioning.
package regen

const (
	DefaultIRIPrefix  = "regen"
	DefaultIRIVersion = IriVersion0
)

// Config is a config struct used for initializing the data module to avoid using globals.
type Config struct {
	// IRIPrefix defines the IRI prefix to use (e.g regen).
	IRIPrefix string
	// IRIVersion defines the IRI version to use (e.g 0).
	IRIVersion byte
}

// DefaultConfig returns the default config for the data module.
func DefaultConfig() Config {
	return Config{
		IRIPrefix:  DefaultIRIPrefix,
		IRIVersion: DefaultIRIVersion,
	}
}
