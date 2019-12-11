package aux

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/OnyxPay/OnyxChain/common/config"
	"github.com/OnyxPay/OnyxChain/common/log"
)

// Stopper closure type to stop API handler
type Stopper func() error

const (
	restStopSignal = syscall.Signal(config.DEFAULT_REST_PORT)
	wsStopSignal   = syscall.Signal(config.DEFAULT_WS_PORT)
	rpcStopSignal  = syscall.Signal(config.DEFAULT_RPC_PORT)
)

// HandleStopSignals starts goroutine which will handle UNIX signals and stop API handlers
func HandleStopSignals(restStopper, wsStopper, rpcStopper Stopper) {
	sc := make(chan os.Signal)
	signal.Notify(sc, restStopSignal, wsStopSignal, rpcStopSignal)
	go func() {
		for sig := range sc {
			switch sig {
			case restStopSignal:
				handleSignal("REST", &restStopper)
			case wsStopSignal:
				handleSignal("WS", &wsStopper)
			case rpcStopSignal:
				handleSignal("RPC", &rpcStopper)
			}
		}
	}()
}

func handleSignal(name string, stopper *Stopper) {
	if *stopper == nil {
		return
	}

	err := (*stopper)()
	if err != nil {
		log.Errorf("%s server stopped with an error: %s", name, err)
	} else {
		log.Warnf("%s server stopped.", name)
	}
	*stopper = nil
}
