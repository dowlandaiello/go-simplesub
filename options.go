// Package pseudosub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, pseudosub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package pseudosub

// Option represents a pseudosub configuration option.
type Option func(*PseudoSub) error

/* BEGIN EXPORTED METHODS */

// WithRoutePrefix defines the WithRoutePrefix option.
// Such an option can used primarily to differentiate between
// different nodes in a network, or partition such networks.
func WithRoutePrefix(prefix string) Option {
	return func(sub *PseudoSub) error {
		sub.RootRoutePath = prefix // Set prefix

		return nil // No error occurred, return nil
	}
}

/* END EXPORTED METHODS */
