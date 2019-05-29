// Package pseudosub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, pseudosub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package pseudosub

import (
	inet "github.com/libp2p/go-libp2p-net"
	routed "github.com/libp2p/go-libp2p/p2p/host/routed"
)

// PseudoSub implements the standard pseudosub pub/sub messaging
// system. The pseudosub type is in many ways analogous
// to the host type in libp2p--it serves as a central hub for
// all pub/sub related operations.
type PseudoSub struct {
	Host *routed.RoutedHost `json:"host"` // Working host

	RootRoutePath string `json:"root_path"` // Root route path

	Handlers map[string]func(inet.Stream) // Message handlers
}

/* BEGIN EXPORTED METHODS */

// NewPseudoSub initializes a new PseudoSub, and sets up all necessary
func NewPseudoSub(host *routed.RoutedHost, opts ...Option) (*PseudoSub, error) {
	sub := &PseudoSub{
		Host: host, // Set host
	} // Initialize pseudosub

	err := sub.applyOptions(opts) // Apply options

	if err != nil { // Check for errors
		return &PseudoSub{}, err // Return found error
	}

	err = sub.setupStreamHandlers() // Setup stream handlers

	if err != nil { // Check for errors
		return &PseudoSub{}, err // Return found error
	}

	return sub, nil // Return initialized sub
}

// Subscribe subscribes to a given topic.
func (sub *PseudoSub) Subscribe(topic string, handler func(inet.Stream)) {
	sub.Handlers[topic] = handler // Set handler
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// applyOptions applies all provided options to the given sub.
func (sub *PseudoSub) applyOptions(opts []Option) error {
	for _, opt := range opts { // Iterate through options
		err := sub.applyOption(opt) // Apply option

		if err != nil { // Check for errors
			return err // Return found error
		}
	}

	return nil // No error occurred, return nil
}

// applyOption applies the provided option to the given sub.
func (sub *PseudoSub) applyOption(opt Option) error {
	return opt(sub) // Apply option
}

/* END INTERNAL METHODS */
