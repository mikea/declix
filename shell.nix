{ pkgs ? import <nixpkgs> {} }:
let
    currentDir = builtins.toString ./.;
in
  pkgs.mkShell {
    hardeningDisable = [ "fortify" ];
    nativeBuildInputs = with pkgs.buildPackages; [ 
        # tools
        watchexec just wget delve
        # dev
        go_1_22 golangci-lint 
        # podman sdk needs theses
        linuxHeaders btrfs-progs pkg-config gpgme
    ];
    shellHook = ''
        export GOROOT=$(go env GOROOT)
        export GOPATH=$(go env GOPATH)
        export PATH=$PATH:$GOPATH/bin:${currentDir}/bin/
    '';
}