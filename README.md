# HTTPHASH: A special HTTP Client

Overview: Takes a list of URLs, outputs each URL and the md5 hash of its response

# Install

install using go

`go install github.com/NH-2/HttpHash@latest`

Add PATH setting your shell config file (.bashrc or .zshrc).

export PATH="$GOPATH/HttpHash:$PATH"

# Build From Source

when Inside the project directory: 

`go build .` creates the binary in the project directory

`go install .` creates the binary in the `$GOPATH`

# Usage

The simplest way to use the CLI is: `HttpHash url1 url2 url3 .... 
example:

```
HttpHash http://example.com messenger.com
```

**output**:
```
http://example.com  84238dfc8092e5d9c0dac8ef93371a07
http://messenger.com  0bbe45b14d908249c34907261a0025ac
```

**Note:** The order of the URLs is not preserved

## Customizing parallel requests
You can set how many parallel requests can be run at a given time by setting the `-parallel` flag.

example: 
```
HttpHash -parallel 2 google.com example.com facebook.com reddit.com http://golang.dev
```

**output**:
```
http://example.com  84238dfc8092e5d9c0dac8ef93371a07
http://google.com  91b2c25f1c2cdf5d4f7224eb288a32f0
http://facebook.com  ffa29b5348a80a64e07323dc49bfa066
http://go.dev  abb74e6791c553562c4d80a3f874385d
http://reddit.com  22e1b1bc00e51f5bba2dc1a24bd3ec37
```

## Invalid Requests and Unreacheable Hosts:

The tool would check that the provided URL is in a valid format. It would display and error otherwise. 
However, this would not block other requests

example:  
```
HttpHash example.com invalid.tn 'hello de/test'
```

**Output**:
```
error: failed to parse url: parse "http://hello de/test": invalid character " " in host name
error: error: perform GET request to the url http://invalid.tn: Get "http://invalid.tn": dial tcp: lookup invalid.tn: no such host
http://example.com  84238dfc8092e5d9c0dac8ef93371a07`
```

# Contributing

PRs are welcome :)

This is a list of possible enhancements I have in mind:
- [ ] Use contexts for making requests. It would enable timeouts and cancel
- [ ] Use a logging library with logging levels.
- [ ] Consider using pointers and passing values by reference for possible performance and memory footprint gains (benchmark results required)
