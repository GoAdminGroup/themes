all:
	make -C ./adminlte
	make -C ./sword
	make fmt

fmt:
	GO111MODULE=off go fmt ./...
	GO111MODULE=off goimports -l -w .