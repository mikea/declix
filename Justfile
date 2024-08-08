TARGET := "local/hamd-target.pkl"
RESOURCES := "local/hamd.pkl"

GOROOT := `go env GOROOT`
GOPATH := `go env GOPATH`

PKL_GEN_GO := GOPATH + "/bin/pkl-gen-go"

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

test: build
    go test ./...

build:
    go build -o bin/declix

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
    export PATH=$(pwd)/bin:$PATH && {{PKL_GEN_GO}} resources/apt/Apt.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{PKL_GEN_GO}} resources/dpkg/Dpkg.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{PKL_GEN_GO}} resources/users/Users.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{PKL_GEN_GO}} systemd/systemd.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{PKL_GEN_GO}} resources/filesystem/FileSystem.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{PKL_GEN_GO}} content/Content.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{PKL_GEN_GO}} resources/Resources.pkl --base-path mikea/declix
    export PATH=$(pwd)/bin:$PATH && {{PKL_GEN_GO}} target/Target.pkl --base-path mikea/declix

[private]
install-cobra-cli:
    go install github.com/spf13/cobra-cli@latest

[private]
install-go-releaser:
    go install github.com/goreleaser/goreleaser/v2@latest

[private]
build-release VERSION: (build-archive VERSION "linux" "amd64") (build-archive VERSION "linux" "arm64")

[private]
build-archive VERSION OS ARCH:
    GOOS={{OS}} GOARCH={{ARCH}} go build -o dist/declix-{{VERSION}}-{{OS}}-{{ARCH}}
    tar \
        -cvzf dist/declix-{{VERSION}}-{{OS}}-{{ARCH}}.tgz \
        --transform "s/declix-{{VERSION}}-{{OS}}-{{ARCH}}/declix/" \
        -C dist declix-{{VERSION}}-{{OS}}-{{ARCH}}
    rm dist/declix-{{VERSION}}-{{OS}}-{{ARCH}}

[private]
gen-pkl version:
    #!/usr/bin/env bash
    set -euxo pipefail
    find . -type f -name "*.pkl" -not -path "*/tests/*" | zip dist/pkl@{{version}}.zip -@
    cp pkl.json.tpl dist/pkl@{{version}}.json
    sed -i "s/VERSION/{{version}}/g" dist/pkl@{{version}}.json
    read -r SHA256 _ < <(sha256sum dist/pkl@{{version}}.zip) 
    sed -i "s/SHA256/$SHA256/g" dist/pkl@{{version}}.json
