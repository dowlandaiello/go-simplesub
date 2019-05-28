// Package pseudosub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, pseudosub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package pseudosub

import "testing"

/* BEGIN EXPORTED METHODS */

// NewPseudoSub tests the functionality of the NewPseudoSub helper method.
func TestNewPseudoSub(t *testing.T) {
	// host, err := libp2p.New()
	// sub, err := NewPseudoSub(host *routedhost.RoutedHost, opts ...Option)
}

/* END EXPORTED METHODS */