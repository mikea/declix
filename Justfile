TARGET := "local/hamd-target.pkl"
RESOURCES := "local/hamd.pkl"

GOROOT := `go env GOROOT`
GOPATH := `go env GOPATH`

alias w := watch

watch +WATCH_TARGET='test':
    watchexec -rc -w . --ignore *.pkl.go -- just {{WATCH_TARGET}}

setup: install-pkl install-pkl-gen-go install-cobra-cli

gen: pkl-gen-go

run: test state actions apply

state:
    export PATH=$(pwd)/bin:$PATH && go run main.go state -t {{TARGET}} -r {{RESOURCES}}

actions:
    export PATH=$(pwd)/bin:$PATH && go run main.go actions -t {{TARGET}} -r {{RESOURCES}}

apply:
    export PATH=$(pwd)/bin:$PATH && go run main.go apply -t {{TARGET}} -r {{RESOURCES}}

test:
    go test ./...

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
    find . -name "*.pkl.go" -exec rm -f {} +
    export PATH=$(pwd)/bin:$PATH && {{GOPATH}}/bin/pkl-gen-go resources/apt/Apt.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{GOPATH}}/bin/pkl-gen-go resources/dpkg/Dpkg.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{GOPATH}}/bin/pkl-gen-go resources/filesystem/FileSystem.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{GOPATH}}/bin/pkl-gen-go content/Content.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{GOPATH}}/bin/pkl-gen-go resources/Resources.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{GOPATH}}/bin/pkl-gen-go target/Target.pkl --base-path mikea/declix

[private]
install-cobra-cli:
    go install github.com/spf13/cobra-cli@latest