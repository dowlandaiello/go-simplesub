// Package simplesub implements a pub/sub messaging system through the libp2p
// routed.RoutedHost interface. In contrast with the standard libp2p pub/sub
// package, simplesub has the advantage of letting developers opt for their
// own routing solutions (e.g. kadDHT).
package simplesub

import (
	"bytes"
	"strings"
	"testing"
)

/* BEGIN EXPORTED METHODS */

// TestMessageFromBytes tests the functionality of the MessageFromBytes helper method.
func TestMessageFromBytes(t *testing.T) {
	message := &Message{
		Topic: "test",
		Data:  []byte("test"),
	} // Initialize message

	encodedMessage, err := message.Bytes() // Get message bytes

	if err != nil { // Check for errors
		t.Fatal(err) // Panic
	}

	decodedMessage, err := MessageFromBytes(encodedMessage) // Decode message

	if err != nil { // Check for errors
		t.Fatal(err) // Panic
	}

	if !bytes.Equal(decodedMessage.Data, message.Data) { // Check invalid message
		t.Fatalf("invalid decoded message; got %s, wanted %s", string(decodedMessage.Data), string(message.Data)) // Panic
	}
}

// TestBytesMessage tests the functionality of the Bytes helper method.
func TestBytesMessage(t *testing.T) {
	message := &Message{
		Topic: "test",
		Data:  []byte("test"),
	} // Initialize message

	encodedMessage, err := message.Bytes() // Encode message

	if err != nil { // Check for errors
		t.Fatal(err) // Panic
	}

	if !strings.Contains(string(encodedMessage), "test") { // Check invalid encoded message
		t.Fatal("invalid encoded message") // Panic
	}
}

/* END EXPORTED METHODS */
