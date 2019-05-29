// Package simplesub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, simplesub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package simplesub

import (
	"bufio"
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	inet "github.com/libp2p/go-libp2p-net"
	protocol "github.com/libp2p/go-libp2p-protocol"
	routed "github.com/libp2p/go-libp2p/p2p/host/routed"
)

// SimpleSub implements the standard simplesub pub/sub messaging
// system. The simplesub type is in many ways analogous
// to the host type in libp2p--it serves as a central hub for
// all pub/sub related operations.
type SimpleSub struct {
	Host *routed.RoutedHost `json:"host"` // Working host

	RootRoutePath string `json:"root_path"` // Root route path

	Handlers map[string]func(inet.Stream, *Message) // Message handlers
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
func (sub *SimpleSub) Subscribe(topic string, handler func(inet.Stream, *Message)) {
	sub.Handlers[topic] = handler // Set handler
}

// Publish publishes to a given topic, to a given subset of peers.
// If no target peers are specified, the message is broadcasted to the entire
// network (i.e. all peers).
func (sub *SimpleSub) Publish(ctx context.Context, topic string, data []byte, peers ...peer.ID) error {
	if len(peers) == 0 { // Check no peers
		return sub.broadcast(ctx, topic, data) // Broadcast
	}

	message := &Message{
		Topic: topic, // Set topic
		Data:  data,  // Set data
	} // Initialize message

	encodedMessage, err := message.Bytes() // Encode message to bytes

	if err != nil { // Check for errors
		return err // Return found error
	}

	for _, peer := range peers { // Iterate through peers
		if sub.Host.ID() == peer { // Check is not self
			continue // Continue
		}

		stream, err := sub.Host.NewStream(ctx, peer, protocol.ID(fmt.Sprintf("%s/sub", sub.RootRoutePath))) // Initialize stream

		if err != nil { // Check for errors
			continue // Continue
		}

		writer := bufio.NewWriter(stream) // Initialize writer

		_, err = writer.Write(append(encodedMessage, '\n')) // Write

		if err != nil { // Check for errors
			continue // Continue
		}
	}

	return nil // No error occurred, return nil
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

// broadcast broadcasts a given message to all available peers.
func (sub *SimpleSub) broadcast(ctx context.Context, topic string, data []byte) error {
	message := &Message{
		Topic: topic, // Set topic
		Data:  data,  // Set data
	} // Initialize message

	encodedMessage, err := message.Bytes() // Encode message to bytes

	if err != nil { // Check for errors
		return err // Return found error
	}

	for _, peer := range sub.Host.Peerstore().Peers() { // Iterate through peers
		if sub.Host.ID() == peer { // Check is not self
			continue // Continue
		}

		stream, err := sub.Host.NewStream(ctx, peer, protocol.ID(fmt.Sprintf("%s/sub", sub.RootRoutePath))) // Initialize stream

		if err != nil { // Check for errors
			continue // Continue
		}

		writer := bufio.NewWriter(stream) // Initialize writer

		_, err = writer.Write(append(encodedMessage, '\n')) // Write

		if err != nil { // Check for errors
			continue // Continue
		}
	}

	return nil // No error occurred, return nil
}

// peerInSLice checks that a given peer ID is in a slice of peer IDs.
func peerInSlice(s []peer.ID, id peer.ID) bool {
	for _, peer := range s { // Iterate through peers
		if peer == id { // Check equivalent
			return true // In slice
		}
	}

	return false // Not in slice
}

/* END INTERNAL METHODS */
