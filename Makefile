INSTALL_PREFIX ?= $(shell pwd)/build

BIN := \
	strava-heatmap-auth \
	strava-heatmap-proxy \
	strava-auth

.PHONY: all
all: $(BIN)

%: ## Build all commands 
	CGO_ENABLED=0 go build -o $(INSTALL_PREFIX)/$@ cmd/$@/main.go