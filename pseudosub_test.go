// Package simplesub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, simplesub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package simplesub

import (
	"context"
	"testing"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	routed "github.com/libp2p/go-libp2p/p2p/host/routed"
)

/* BEGIN EXPORTED METHODS */

// NewPseudoSub tests the functionality of the NewPseudoSub helper method.
func TestNewPseudoSub(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background()) // Initialize context

	defer cancel() // Cancel

	host, err := libp2p.New(
		ctx,
		libp2p.NATPortMap(),
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/1111",
			"/ip6/::1/tcp/1111",
		),
	) // Initialize host

	if err != nil { // Check for errors
		t.Fatal(err) // Panic
	}

	dht, err := dht.New(ctx, host) // Initialize dht

	if err != nil { // Check for errors
		t.Fatal(err) // Panic
	}

	routedHost := routed.Wrap(host, dht) // Wrap host

	sub, err := NewPseudoSub(routedHost) // Initialize sub

	if err != nil { // Check for errors
		t.Fatal(err) // Panic
	}

	if sub == nil { // Check for errors
		t.Fatal("Nil sub.") // Panic
	}
}

/* END EXPORTED METHODS */
