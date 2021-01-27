function build() {
    go build -ldflags="-s -w" -o "out/$1/$2/prettyJson" .
}

command -v go
go version
env

build windows 386
build darwin amd64
build linux amd64