{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    nativeBuildInputs = with pkgs.buildPackages; [ 
        # tools
        watchexec just wget
        # dev
        go_1_22
    ];
    shellHook = ''
        export GOROOT=$(go env GOROOT)
    '';
}