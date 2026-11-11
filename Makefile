
GO ?= go
GORELEASER ?= goreleaser
TEST_FLAGS ?= -count=1

.PHONY: test test-all test-commands test-services test-chirpstack \
        test-gateway test-gateway-service test-clientfactory \
        test-race test-coverage vet release-check snapshot release

# Run all tests.
test: test-all

test-all:
	$(GO) test $(TEST_FLAGS) ./...

# Test all Cobra commands.
test-commands:
	$(GO) test $(TEST_FLAGS) ./internal/commands/...

# Test all domain services.
test-services:
	$(GO) test $(TEST_FLAGS) ./internal/service/...

# Test ChirpStack adapters and client infrastructure.
test-chirpstack:
	$(GO) test $(TEST_FLAGS) ./internal/chirpstack/...

# Test gateway commands.
test-gateway:
	$(GO) test $(TEST_FLAGS) ./internal/commands/resources/gateway/...

# Test the gateway service.
test-gateway-service:
	$(GO) test $(TEST_FLAGS) ./internal/service/gateway/...

# Test the client factory.
test-clientfactory:
	$(GO) test $(TEST_FLAGS) ./internal/chirpstack/clientfactory/...

# Run tests with the race detector.
test-race:
	$(GO) test -race -count=1 ./...

# Generate a coverage report.
test-coverage:
	$(GO) test -count=1 -coverprofile=coverage.out ./...
	$(GO) tool cover -func=coverage.out

# Static analysis.
vet:
	$(GO) vet ./...

# Validate GoReleaser configuration.
release-check:
	$(GORELEASER) check

# Build all release artifacts without publishing.
snapshot:
	$(GORELEASER) release --snapshot --clean

# Build and publish the GitHub release.
release:
	$(GORELEASER) release --clean
