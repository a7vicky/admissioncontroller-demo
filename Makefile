.PHONY: test
test:
	@echo "\n🛠️  Running unit tests..."
	go test ./...

.PHONY: build
build:
	@echo "\n🔧  Building Go binaries..."
	GOOS=linux GOARCH=amd64 go build -o bin/admission-webhook-demo ./cmd/server/main.go

.PHONY: podman-build
podman-build:
	@echo "\n📦 Building admission-webhook-demo Podman image..."
	podman build --tag  quay.io/aabhishe/webhook-demo:v1 -f Dockerfile

.PHONY: build-push
build-push:
	@echo "\n📦 Building admission-webhook-demo Podman image and push to quay..."
	podman build --tag  quay.io/aabhishe/webhook-demo:v1 -f Dockerfile
	podman push quay.io/aabhishe/webhook-demo:v1