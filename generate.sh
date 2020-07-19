export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin
protoc -I src/ --go_out=src/ src/simple/simple.proto
protoc -I src/ --go_out=src/ src/enum_example/enum_example.proto