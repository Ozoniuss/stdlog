# Stdlog


> Note: I've written this some time ago, and it's time this package gets archived. Even though I still haven't encountered any issues with it, I moved away from not really caring about it to having a preference for loggers passed via dependency injection. This logger, even though it's mostly just a formatter of the standard logger, still acts similar to a global variable, making it very hard to disable in cases such as testing a function that logs something, where you would not necessarily want to see the same logs. I've also linked below some articles discussing this topic. Btw, log/slog will be out in go 1.21, so can't wait to test that out!

Reference to mentioned articles:

- [Dave Cheney, a GoD](https://dave.cheney.net/2017/01/23/the-package-level-logger-anti-pattern)
- [Go for Industrial Programming](https://peter.bourgon.org/go-for-industrial-programming/#logging)
- [People like to argue about programming :)](https://www.reddit.com/r/golang/comments/p7kzti/understanding_global_scope_and_how_it_impacts/)

Introduction
------------

This package defines methods for logging information to standard output. It provides an info logging message, a warning logging message and an error logging message. Since that is what I often use in my code, I found this to be the simplest approach for me personally to handle logging.


Installation and usage
----------------------

To download the package, run 

```
go get "github.com/Ozoniuss/stdlog"
```

Sample usage:

```go
package main

import (
	"fmt"
	"os"

	log "github.com/Ozoniuss/stdlog"
)

// readFile reads the content of a file to a buffer.
func readFile(buffer []byte, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	f.Read(buffer)
	return nil
}

func main() {
	buffer := make([]byte, 32)
	log.Infoln("Initialized buffer.")

	err := readFile(buffer, "a.txt")
	if err != nil {
		log.Errf("Could not read file: %s", err.Error())
		os.Exit(1)
	}

	log.Infoln("File read and stored in buffer.")
	fmt.Printf("%s", buffer)
}
```

Generated logs:

```
[info] 2023/02/18 15:09:11 Initialized buffer.
[info] 2023/02/18 15:09:11 File read and stored in buffer.
Hello world
```

```
[info] 2023/02/18 15:07:29 Initialized buffer.
[error] 2023/02/18 15:07:29 Could not read file: could not open file: open b.txt: The system cannot find the file specified.
```

Thread-safety
-------------

It is safe to call the logging methods of these package from different goroutines. Printing to the standard output in Go is serialized correctly even from different loggers, and messages do not get interlaced. Check out [this example](https://github.com/Ozoniuss/what-does-this-code-do/blob/main/golang/03-multithreaded-prints/main.go) for reference.


Logging approach
----------------

This package is of course much more limited than the standard log package. In most cases it's not ideal for production setups, where a configurable logger such as the default one or even dedicated third-party logging packages such as [go-hclog](https://github.com/hashicorp/go-hclog) or [zap](https://github.com/uber-go/zap) are better suited. For my own projects though, I usually couldn't care less about logging so I just want to have something easy to use that formats output nicely.

There are three approaches that I've generally considered for logging:

- Global logger as global variable;
- Passing a logger explicit via dependency injection;
- New logger for every log message.

I would have normally gone with passing the logger as an explicit dependency to all functions that write logs, but that extra parameter is kind of ugly, and I wanted to avoid that for no other reason than I like my functions to have fewer parameters. I could have also gone with global loggers from the standard log package, there's nothing particularly wrong about that, but global variables are generally frowned upon, so I do usually prefer avoiding them.

Thus, I went with creating a new logger for each log message in this package. This never caused me any problems, so for now I had no reason to do something more complex. Additionally, the overhead added by creating a new logger every time is insignificant, the performance seems to be almost the same with using the fmt package: https://github.com/Ozoniuss/misc/blob/main/test-stdlog