# echo "Installing gogo/protobuf..."
# rm -rf $GOGOPROTO_ROOT
# go get -v github.com/gogo/protobuf/{proto,protoc-gen-gogo,gogoproto,protoc-gen-gofast}
# go get -v golang.org/x/tools/cmd/goimports
# pushd "${GOGOPROTO_ROOT}"
#   git reset --hard HEAD
#   make install
# popd

protoc --gofast_out=plugins=grpc:. \
  --proto_path=$GOPATH/src:$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1:. \
  --go_out=./ \
  *.proto;