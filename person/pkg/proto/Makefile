codegen:person.pb.go
codegen:person.pb.gw.go
codegen:person.swagger.json



person.pb.go:person.proto
	protoc -I/usr/local/include -I. \
    -I$$GOPATH/src \
    -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
    person.proto

person.pb.gw.go:person.proto
	protoc -I/usr/local/include -I. \
    -I$$GOPATH/src  -I$$GOPATH/src/github.com/googleapis/googleapis/ \
    --grpc-gateway_out=logtostderr=true:. \
    person.proto


person.swagger.json:person.proto
	protoc -I/usr/local/include -I. \
	-I$$GOPATH/src/grpc-hello-world/proto/google/api \
	--swagger_out=logtostderr=true:. ./hello.proto

clean:
	-rm -rf person.pb.go
	-rm -rf person.pb.gw.go
	-rm -rf person.swagger.json

