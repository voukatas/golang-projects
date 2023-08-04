# Go Port Scanner

This repository contains a simple port scanner written in Go. It scans a given range of ports for a specified host to check if they are open.

## Versions

There are two versions of the scanner:

1. `scanner_procedural.go`: A procedural version where all logic is contained within independent functions. This version is straightforward and easy to follow.

2. `scanner_struct.go`: An object-oriented version that encapsulates the state and behavior of a scanner into a `Scanner` struct. This version is easier to reason about, test, and reuse.

## Usage

Both versions of the scanner are command-line utilities that take three arguments:

1. `hostname`: The address of the target host to be scanned.
2. `start-port`: The first port number in the range to be scanned.
3. `end-port`: The last port number in the range to be scanned.

For example, you can run the scanner like this:

```bash
go run scanner_procedural.go localhost 1 1024
go run scanner_struct.go localhost 1 1024

```
These commands will scan ports 1 through 1024 on localhost.