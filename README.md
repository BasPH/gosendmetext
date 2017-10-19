# Gosendmetext

A Golang rewrite of https://github.com/BasPH/sendmetext, a small Bash CLI tool for generating random sentences from a textfile and sending to a given socket.

Main purpose of this rewrite is learning Golang, however the CLI itself can be useful for testing purposes. I sometimes use it for testing data flowing through telnet (streaming APIs often use [telnet examples](https://spark.apache.org/docs/2.2.0/structured-streaming-programming-guide.html#quick-example)).

## Usage
1. In a first window `nc -lk 9999` to listen on tcp port 9999.
2. In a second window `./gosendmetext` to start sending sentences.

Flags:
```
-d/--debug    (optional) print debug logging
-a/--address  (optional) address to send to, default 127.0.0.1
-p/--port     (optional) port to send to, default 9999
-n/--protocol (optional) protocol to use, default tcp
-t/--textfile (optional) file containing words, default /usr/share/dict/words
--minwords    (optional) min number of random words to generate, default 1
--maxwords    (optional) max number of random words to generate, default 50
-s/--sleep    (optional) sleep period in between sending sentences, default 0.7s
```

Example output on tcp port 9999:
```
wifecarl trembling Judy pentrite psychoclinicist empirics cerebrosclerosis blackneb synchronistically 
quadruplicature brandise axillar deepmost freehearted overwind phytomer charge blick liomyofibroma
teretial wanworth bathos
```

## Building from source
`make` to install in your `$GOPATH/bin` directory. Or, `make build` to build in the package source directory.
