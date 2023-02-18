# Stdlog

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