package openapi

//go:generate env GOBIN=$PWD/bin GO111MODULE=on go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
//go:generate $PWD/bin/oapi-codegen -generate types -package openapi -o types.generated.go ../../../openapi.yaml
//go:generate $PWD/bin/oapi-codegen -generate server -package openapi -o server.generated.go ../../../openapi.yaml
//go:generate $PWD/bin/oapi-codegen -generate spec -package openapi -o spec.generated.go ../../../openapi.yaml
