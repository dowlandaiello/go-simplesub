# go-simplesub

A minimalistic, yet powerful pubsub messaging system built on top of libp2p.

## Rationale

```English
Why does this repo exist?
```

Simple: Libp2p's pub/sub implementation simply does not provide enough flexibility
for some use cases (those that utilize a DHT or any sort of routing in particular).

## Installation

```zsh
go get github.com/dowlandaiello/go-simplesub
```

## Usage

```Go
package main

import (
    "context"

    "github.com/libp2p/go-libp2p"
    dht "github.com/libp2p/go-libp2p-kad-dht"
    routed "github.com/libp2p/go-libp2p/p2p/host/routed"
    "github.com/dowlandaiello/go-simplesub"
)

func main() {
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
        panic(err) // Panic
    }

    dht, err := dht.New(ctx, host) // Initialize dht

    if err != nil { // Check for errors
        panic(err) // Panic
    }

    routedHost := routed.Wrap(host, dht) // Wrap host

    sub, err := simplesub.NewSimpleSub(routedHost) // Initialize sub

    if err != nil { // Check for errors
        panic(err) // Panic
    }

    sub.Subscribe("test_topic") // Subscribe

    err = sub.Publish(ctx, "test_topic", []byte("test")) // Publish to topic

    if err != nil { // Check for errors
        panic(err) // Panic
    }
}
```
