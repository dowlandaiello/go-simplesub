// Package simplesub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, simplesub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package simplesub

import (
	"encoding/json"
)

// Message outlines a simplesub message.
type Message struct {
	Topic string `json:"topic"` // Message topic

	Data []byte `json:"data"` // Message data
}

/* BEGIN EXPORTED METHODS */

// MessageFromBytes attempts to decode a message from a given byte slice.
func MessageFromBytes(b []byte) (*Message, error) {
	buffer := new(Message) // Init message buffer

	if err := json.Unmarshal(b, buffer); err != nil { // Check for errors
		return &Message{}, err // Return found error
	}

	return buffer, nil // Return decoded message
}

// Bytes serializes a given message to a byte slice.
func (message *Message) Bytes() ([]byte, error) {
	return json.Marshal(*message) // Return encoded
}

/* END EXPORTED METHODS */
