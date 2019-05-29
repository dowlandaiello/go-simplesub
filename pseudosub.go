// Package simplesub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, simplesub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package simplesub

import (
	inet "github.com/libp2p/go-libp2p-net"
	routed "github.com/libp2p/go-libp2p/p2p/host/routed"
)

// SimpleSub implements the standard simplesub pub/sub messaging
// system. The simplesub type is in many ways analogous
// to the host type in libp2p--it serves as a central hub for
// all pub/sub related operations.
type SimpleSub struct {
	Host *routed.RoutedHost `json:"host"` // Working host

	RootRoutePath string `json:"root_path"` // Root route path

	Handlers map[string]func(inet.Stream) // Message handlers
}

/* BEGIN EXPORTED METHODS */

// NewSimpleSub initializes a new SimpleSub, and sets up all necessary stream handlers.
func NewSimpleSub(host *routed.RoutedHost, opts ...Option) (*SimpleSub, error) {
	sub := &SimpleSub{
		Host: host, // Set host
	} // Initialize simplesub

	err := sub.applyOptions(opts) // Apply options

	if err != nil { // Check for errors
		return &SimpleSub{}, err // Return found error
	}

	err = sub.setupStreamHandlers() // Setup stream handlers

	if err != nil { // Check for errors
		return &SimpleSub{}, err // Return found error
	}

	return sub, nil // Return initialized sub
}

// Subscribe subscribes to a given topic.
func (sub *SimpleSub) Subscribe(topic string, handler func(inet.Stream)) {
	sub.Handlers[topic] = handler // Set handler
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// applyOptions applies all provided options to the given sub.
func (sub *SimpleSub) applyOptions(opts []Option) error {
	for _, opt := range opts { // Iterate through options
		err := sub.applyOption(opt) // Apply option

		if err != nil { // Check for errors
			return err // Return found error
		}
	}

	return nil // No error occurred, return nil
}

// applyOption applies the provided option to the given sub.
func (sub *SimpleSub) applyOption(opt Option) error {
	return opt(sub) // Apply option
}

/* END INTERNAL METHODS */
