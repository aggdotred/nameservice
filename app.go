package app

import (
	"log"
	"os"
)

const appName = "nameservice"

var (
	// DefaultCLIHome sets dirs for the app CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.nscli")

	// DefaultNodeHome sets the folder where the app data and config will be stored
	DefaultNodeHome = os.ExpandEnv("$home/.nsd")

	// ModuleBasics is in charge of setting up basic module elements
	ModuleBasics = module.NewBasicManager()
)

type nameServiceApp struct {
	*bam.BaseApp
}

// NewNameServiceApp is the app constructor
func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {
	// define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nameServiceApp{
		BaseApp: bApp,
		cdc:     cdc,
	}
}
