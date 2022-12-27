let
  pkgs = import <nixpkgs> {};
  projectName = "Unstable CMS";
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
    exec ./bin/${execName}
  '';
in pkgs.mkShell rec {
  name = projectName;
  buildInputs = with pkgs; [
      figlet
      go
      cobra-cli
      build
      run
  ];
  shellHook= ''
    figlet -f doom "Entering goland"

    export GOPATH="${pwd}/.go"
    export GOCACHE=""
    export GO111MODULE='on'
    export PATH="$PATH:GOPATH";

    go mod init ${goPackagePath}
    
    figlet -f doom "Happy hacking!"
  '';
}
