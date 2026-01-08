{ pkgs, ... }:

{
  packages = [ pkgs.git ];

  languages.go.enable = true;

  git-hooks.hooks = {
    gofmt.enable = true;
    govet.enable = true;
    golangci-lint.enable = true;
    gotest.enable = true;
    commitizen.enable = true;
  };
}
