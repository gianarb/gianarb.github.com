{ pkgs ? import <nixpkgs> { } }:
pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    glibc
    jekyll
    ruby

    libyaml
    libffi
    zlib
    pkg-config
    readline
  ];
}
