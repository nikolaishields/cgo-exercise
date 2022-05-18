{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    bashInteractive
    gnumake
    go_1_17
    shellcheck
    shfmt
  ];
}
