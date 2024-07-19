{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    hardeningDisable = [ "fortify" ];
    nativeBuildInputs = with pkgs.buildPackages; [ 
        # tools
        watchexec just wget delve
        # dev
        go_1_22 golangci-lint
    ];
    shellHook = ''
        export GOROOT=$(go env GOROOT)
        export GOPATH=$(go env GOPATH)
        export PATH=$PATH:$GOPATH/bin
    '';
}