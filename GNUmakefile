HOSTNAME=cappyzawa.com
NAMESPACE=cappyzawa
NAME=christmas-tree
VERSION=0.0.1
OS_ARCH=darwin_amd64
BINARY=terraform-provider-${NAME}

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test -cover ./... -v $(TESTARGS) -timeout 120m

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
