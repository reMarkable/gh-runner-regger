{ pkgs, ... }:

{
  env.GOPROXY = "direct";


  packages = [ pkgs.git ];

  languages.go.enable = true;

  pre-commit.hooks.gofmt.enable = true;
  pre-commit.hooks.govet.enable = true;
  # pre-commit.hooks.gotest.enable = true;
  pre-commit.hooks.commitizen.enable = true;
}
