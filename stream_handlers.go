// Package simplesub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, simplesub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package simplesub

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"

	inet "github.com/libp2p/go-libp2p-net"
	protocol "github.com/libp2p/go-libp2p-protocol"
)

/* BEGIN INTERNAL METHODS */

// setupStreamHandlers sets up all the required stream
// handlers for proper simplesub function.
func (sub *SimpleSub) setupStreamHandlers() error {
	sub.Host.SetStreamHandler(protocol.ID(fmt.Sprintf("%s/sub", sub.RootRoutePath)), sub.handleReceiveSub) // Set sub handler

	return nil // No error occurred, return nil
}

// handleReceiveSub handles a received message (i.e. peer published).
func (sub *SimpleSub) handleReceiveSub(stream inet.Stream) {
	reader := bufio.NewReader(stream) // Init reader

	b, err := reader.ReadBytes('\n') // Read up to newline

	if err != nil { // Check for errors
		return // Stop execution
	}

	b = bytes.Trim(b, "\n") // Remove newline

	message, err := MessageFromBytes(b) // Decode message

	if err != nil { // Check for errors
		return // Stop execution
	}

	if !reflect.ValueOf(sub.Handlers[message.Topic]).IsNil() && reflect.ValueOf(sub.Handlers[message.Topic]).IsValid() { // Ensure has handler
		sub.Handlers[message.Topic](stream, message) // Call handler
	}
}

/* END INTERNAL METHODS */
