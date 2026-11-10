
# Go test configuration.
GO ?= go
TEST_FLAGS ?= -count=1

.PHONY: test test-all test-commands test-services test-chirpstack \
        test-gateway test-gateway-service test-clientfactory \
        test-race test-coverage

# Default: run every test in the Go module.
test: test-all

test-all:
	$(GO) test $(TEST_FLAGS) ./...

# Test all Cobra commands and their subpackages.
test-commands:
	$(GO) test $(TEST_FLAGS) ./internal/commands/...

# Test all domain service packages.
test-services:
	$(GO) test $(TEST_FLAGS) ./internal/service/...

# Test ChirpStack adapters and client infrastructure.
test-chirpstack:
	$(GO) test $(TEST_FLAGS) ./internal/chirpstack/...

# Test gateway commands only.
test-gateway:
	$(GO) test $(TEST_FLAGS) ./internal/commands/resources/gateway/...

# Test the gateway service package.
test-gateway-service:
	$(GO) test $(TEST_FLAGS) ./internal/service/gateway/...

# Test the client factory package.
test-clientfactory:
	$(GO) test $(TEST_FLAGS) ./internal/chirpstack/clientfactory/...

# Check for data races across the project.
test-race:
	$(GO) test -race -count=1 ./...

# Generate an overall coverage profile.
test-coverage:
	$(GO) test -count=1 -coverprofile=coverage.out ./...
	$(GO) tool cover -func=coverage.out
