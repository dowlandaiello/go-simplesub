# go-pseudosub

A pub/sub messaging system built on top of the libp2p routed.RoutedHost interface.

## Rationale

```English
Why does this repo exist?
```

Simple: Libp2p's pub/sub implementation simply does not provide enough flexibility
for some use cases (those that utilize a DHT or any sort of routing in particular).

## Installation

```zsh
go get github.com/dowlandaiello/go-pseudosub
```
