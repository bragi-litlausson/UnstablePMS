let
  pkgs = import <nixpkgs> {};
  projectName = "Unstable PM";
  goPackagePath = "github.com/bragi-litlausson/UnstablePM";
  execName = "upm";
  pwd = "$PWD";
  build = pkgs.writeScriptBin "build" ''
    if [ ! -d "./bin" ]
    then
      mkdir ./bin
    fi

    go build -o ./bin/${execName} main.go
  '';
  run = pkgs.writeScriptBin "run" ''
    build
    exec ./bin/${execName} $*
  '';
  format = pkgs.writeScriptBin "format" ''
    gofmt -l -w -s .
  '';
in pkgs.mkShell rec {
  name = projectName;
  buildInputs = with pkgs; [
      format
      build
      run

      figlet
      go
      cobra-cli
  ];
  shellHook= ''
    figlet -f doom "Entering goland"

    export GOPATH="${pwd}/.go"
    export GOCACHE=""
    export GO111MODULE='on'
    export PATH="$PATH:GOPATH";

    if [ ! -f "./go.mod" ]
    then
      go mod init ${goPackagePath}
    fi
    
    figlet -f doom "Happy hacking!"
  '';
}
