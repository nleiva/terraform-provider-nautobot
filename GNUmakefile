TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=github.com
NAMESPACE=nleiva
NAME=nautobot
BINARY=terraform-provider-${NAME}
VERSION=0.3.2
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)


default: install

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p $(HOME)/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} $(HOME)/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

get-api:
	cd client; wget https://demo.nautobot.com/api/swagger.yaml\?api_version\=1.3 -O swagger.yaml

generate: get-api	
	cd client; oapi-codegen -generate client -o nautobot.go -package nautobot swagger.yaml && \
	oapi-codegen -generate types -o types.go -package nautobot swagger.yaml && \
	go mod tidy

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4   

gpg-key:
	gpg --armor --export-secret-key $(EMAIL) -w0 | xclip -selection clipboard -i

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

local: install
	sed -i "s-/home/nleiva-${HOME}-" test/.terraform/plugin_path
	sed -i 's-version =.*-version = "${VERSION}"-' test/main.tf
	cd test; terraform init -upgrade && \
	terraform apply -auto-approve; cd ..
	
tag: local
	git add .
	git commit -m "Bump to version ${VERSION}"
	git tag -a -m "Bump to version ${VERSION}" v${VERSION}
	git push --follow-tag