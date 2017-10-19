# Gosendmetext

A Golang rewrite of https://github.com/BasPH/sendmetext, a small Bash CLI tool for generating random sentences from a textfile and sending to a given address.

Main purpose of this rewrite is learning Golang, however the CLI itself can be useful for testing purposes. I sometimes use it for testing data flowing through telnet (streaming APIs often use [telnet examples](https://spark.apache.org/docs/2.2.0/structured-streaming-programming-guide.html#quick-example)).

## Usage
```bash
./gosendmetext

Flags:
-d/--debug    (optional) print debug logging
-a/--address  (optional) address to send to, default 127.0.0.1
-p/--port     (optional) port to send to, default 9999
-n/--protocol (optional) protocol to use, default tcp
-t/--textfile (optional) file containing words, default /usr/share/dict/words
--minwords    (optional) min number of random words to generate, default 1
--maxwords    (optional) max number of random words to generate, default 50
-s/--sleep    (optional) sleep period in between sending sentences, default 0.7s

```

## Building from source
