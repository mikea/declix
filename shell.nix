# see https://nixos.wiki/wiki/Go
{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    hardeningDisable = [ "fortify" ];
    nativeBuildInputs = with pkgs.buildPackages; [ 
        # tools
        watchexec just wget delve
        # dev
        go_1_22
    ];
    shellHook = ''
        export GOROOT=$(go env GOROOT)
        export GOPATH=$(go env GOPATH)
    '';
}