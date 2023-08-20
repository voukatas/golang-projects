# Golang CLI Reminder

A simple command-line reminder tool written in Go. Set reminders for specific times, and receive desktop notifications when the time arrives.

## Prerequisites

Ensure you have Go installed on your system. Additionally, the program uses the following external libraries:

- `github.com/gen2brain/beeep`
- `github.com/olebedev/when`
- `github.com/olebedev/when/rules/common`
- `github.com/olebedev/when/rules/en`

You can install them using `go get`:

```bash
go get github.com/gen2brain/beeep
go get github.com/olebedev/when
```

# Usage
To set a reminder:
```bash
./program_name <hh:mm> "Your reminder message"

```

For example:
```bash
./program_name 15:30 "This is a reminder message"

```