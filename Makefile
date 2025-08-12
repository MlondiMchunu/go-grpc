protoc:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	./**/*.proto
