# Golang Mock Server
This utility can be used to mock apis and respond with xml/json response.

## Setup
This service runs on go.

- Install go
  - On OSX run `brew install go`.
  - Follow instructions on https://golang.org/doc/install for other OSes.
- Setup go
  - Make sure that the executable `go` is in your shell's path.
  - Add the following in your .zshrc or .bashrc: (where `<workspace_dir>` is the directory in
    which you'll checkout your code)

```
GOPATH=<workspace_dir>
export GOPATH
PATH="${PATH}:${GOPATH}/bin"
export PATH
```

- Checkout the code and build the project:
git clone https://github.com/bhambri94/mockit.git

cd mockit/

docker build -t mockit:latest .

docker images

docker run -it --rm -p 8010:8010 -v $PWD/src:/go/src/mockit mockit