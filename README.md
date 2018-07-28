# gosplitter

[![Travis CI](https://api.travis-ci.org/NIPE-SYSTEMS/gosplitter.svg?branch=master)](https://travis-ci.org/NIPE-SYSTEMS/gosplitter) [![GoDoc](https://godoc.org/github.com/NIPE-SYSTEMS/gosplitter?status.svg)](https://godoc.org/github.com/NIPE-SYSTEMS/gosplitter)

![Screenshot](screenshot.png)

This repository contains a channel splitter which broadcasts messages received from an input channel to a dynamic amount of output channels.

## Installation

    go get github.com/NIPE-SYSTEMS/gosplitter

## Create a new splitter

NewSplitter() creates a new splitter from a given input channel. It returns an add function that may be used for adding more outputs to the splitter. The add function returns a new output channel with the given capacity and a remove function. To prevent memory leaks the remove function must be called when the output channel is no longer needed.

Create splitter:

```go
input := make(chan interface{})
add := NewSplitter(input, capacity)
```

Add output channel:

```go
output, remove := add()
defer remove() // ensure that output gets removed
```

Close the input channel when done:

```go
close(input) // will close output channel if it has not been removed yet
```

## License

MIT
