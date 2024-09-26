
build-proto:
	buf generate --clean --config proto/buf.yaml --template proto/buf.gen.yaml
