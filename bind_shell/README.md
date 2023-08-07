# Simple Bind Shell in Go

This program is a minimalistic bind shell implemented in Go. It listens on a specified TCP port and spawns a shell (`cmd.exe` on Windows and `/bin/sh` on Unix-like systems) for any incoming connection.

## Usage

1. First, make sure you have Go installed on your machine.
2. Run the code using:
    ```bash
    go run bind_shell.go
    ```

    Alternatively, build the binary using:
    ```bash
    go build bind_shell.go
    ```

3. The program will start listening on port `5000`.
4. You can connect to it using any TCP client, for example, using `netcat`:
    ```bash
    nc localhost 5000
    ```

5. Upon successful connection, you will be presented with a shell from the machine running the bind shell program.