package aux

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/OnyxPay/OnyxChain/common/log"
)

// Stopper closure type to stop API handler
type Stopper func() error

const (
	restStopSignal = syscall.Signal(35)     // SIGRTMIN on Alpine
	wsStopSignal   = syscall.Signal(35 + 1) // SIGRTMIN+1
	rpcStopSignal  = syscall.Signal(35 + 2) // SIGRTMIN+2
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
