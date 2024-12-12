{ pkgs, ... }:

{
  packages = [ pkgs.git ];

  languages.go.enable = true;

  pre-commit.hooks.gofmt.enable = true;
  pre-commit.hooks.govet.enable = true;
  pre-commit.hooks.golangci-lint.enable = true;
  pre-commit.hooks.gotest.enable = true;
  pre-commit.hooks.commitizen.enable = true;
}
