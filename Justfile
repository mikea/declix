GOROOT := `go env GOROOT`
GOPATH := `go env GOPATH`

alias w := watch

watch +WATCH_TARGET='test':
    watchexec -rc -w . -w local/ -- just {{WATCH_TARGET}}

setup: install-pkl install-pkl-gen-go install-cobra-cli

gen: pkl-gen-go

run:
    export PATH=$(pwd)/bin:$PATH && go run main.go apply local/hamd.pkl

[private]
install-pkl:
    mkdir -p bin
    wget -O bin/pkl https://github.com/apple/pkl/releases/download/0.26.1/pkl-alpine-linux-amd64
    chmod +x bin/pkl

[private]
install-pkl-gen-go:
    go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest

[private]
pkl-gen-go:
    export PATH=$(pwd)/bin:$PATH && {{GOPATH}}/bin/pkl-gen-go pkl/System.pkl --base-path github.com/mikea/declix

[private]
install-cobra-cli:
    go install github.com/spf13/cobra-cli@latest