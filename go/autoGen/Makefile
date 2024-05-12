openapi:
	oapi-codegen -generate "types" -package openapi openapi.yaml > ./backend/openapi/types.gen.go
	oapi-codegen -generate "server" -package openapi openapi.yaml > ./backend/openapi/server.gen.go
	oapi-codegen -generate "spec" -package openapi openapi.yaml > ./backend/openapi/spec.gen.go
