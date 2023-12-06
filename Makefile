# Columbo makefile

build: # Build Columbo
	@echo "Building Columbo..."
	@go build -o bin/schema_gen cmd/schema_gen.go ingest/main.go

# Generate structs from schemas, depends on build target above
generate: build
	./bin/schema_gen
