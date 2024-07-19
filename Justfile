TARGET := "local/hamd-target.pkl"
RESOURCES := "local/hamd.pkl"

GOROOT := `go env GOROOT`
GOPATH := `go env GOPATH`

alias w := watch

watch +WATCH_TARGET='test':
    watchexec -rc -w . --ignore *.pkl.go --ignore main --print-events -- just {{WATCH_TARGET}}

setup: install
install: install-pkl install-pkl-gen-go install-cobra-cli install-go-releaser

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

build:
    go build main.go
    mv main bin/declix

release version: clean dist (build-release version) (gen-pkl version)

dist:
    mkdir dist

clean:
    rm -rf dist

cut-release VERSION:
    # check clean tree
    git diff --exit-code
    git diff --cached --exit-code

    git tag {{VERSION}}
    git push origin
    git push origin {{VERSION}}


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

[private]
install-go-releaser:
    go install github.com/goreleaser/goreleaser/v2@latest

[private]
build-release version:
    GOOS="linux" GOARCH="amd64"  go build -o dist/declix
    zip dist/declix-{{version}}-linux-amd64.zip dist/declix

[private]
gen-pkl version:
    #!/usr/bin/env bash
    set -euxo pipefail
    find . -type f -name "*.pkl" -print0 | zip dist/pkl@{{version}}.zip -@
    cp pkl.json.tpl dist/pkl@{{version}}.json
    sed -i "s/VERSION/{{version}}/g" dist/pkl@{{version}}.json
    read -r SHA256 _ < <(sha256sum dist/pkl@{{version}}.zip) 
    sed -i "s/SHA256/$SHA256/g" dist/pkl@{{version}}.json
