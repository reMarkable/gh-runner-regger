{ pkgs, ... }:

{
  packages = [ pkgs.git ];

  languages.go.enable = true;

  pre-commit.hooks.golangci-lint.enable = true;
  pre-commit.hooks.commitzen.enable = true;
}
