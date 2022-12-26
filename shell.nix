let
  pkgs = import <nixpkgs> {};
  projectName = "Unstable CMS";
  goPackagePath = "github.com/bragi-litlausson/UnstablePM";
  pwd = "$PWD";
in pkgs.mkShell rec {
  name = projectName;
  buildInputs = with pkgs; [
      go
      cobra-cli
  ];
  shellHook= ''
    set -v

    export GOPATH="${pwd}/.go"
    export GOCACHE=""
    export GO111MODULE='on'
    export PATH="$PATH:GOPATH";

    go mod init ${goPackagePath}

    set +v
  '';
}
