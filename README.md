<center><h1>gosplitter</h1></center>
<center><a href="https://travis-ci.org/NIPE-SYSTEMS/gosplitter"><img src="https://api.travis-ci.org/NIPE-SYSTEMS/gosplitter.svg?branch=master" alt="Travis CI" /></a></center>
<center><img src="https://github.com/NIPE-SYSTEMS/gosplitter/raw/master/screenshot.png" alt="Screenshot" /></center>

## Installation

    go get github.com/NIPE-SYSTEMS/gosplitter

## Cretae a new splitter

NewSplitter() creates a new splitter from a given input channel. It returns an add function that may be used for adding more
outputs to the splitter. The add function returns a new output channel with the given capacity and a remove
function. To prevent memory leaks the remove function must be called when the output channel is no longer needed.

Create splitter:

    input := make(chan interface{})
    add := NewGenericSplitter(input, capacity)

Add output channel:

    output, remove := add()
    defer remove() // ensure that output gets removed

## License

MIT
