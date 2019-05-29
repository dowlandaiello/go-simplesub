// Package simplesub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, simplesub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package simplesub

import "strings"

// Option represents a pseudosub configuration option.
type Option func(*PseudoSub) error

/* BEGIN EXPORTED METHODS */

// WithRoutePrefix defines the WithRoutePrefix option.
// Such an option can used primarily to differentiate between
// different nodes in a network, or partition such networks.
func WithRoutePrefix(prefix string) Option {
	return func(sub *PseudoSub) error {
		prefixCpy := prefix // Copy prefix

		if !strings.Contains(prefix, "/") { // Check has no slash
			prefixCpy = "/" + prefixCpy // Prepend slash
		}

		sub.RootRoutePath = prefixCpy // Set prefix

		return nil // No error occurred, return nil
	}
}

/* END EXPORTED METHODS */
