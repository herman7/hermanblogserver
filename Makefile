# 默认执行 all 目标
.DEFAULT_GOAL := all

# ==============================================================================
# 定义 Makefile all 伪目标，执行 `make` 时，会默认会执行 all 伪目标
.PHONY: all
all: go.build

# ==============================================================================
# Includes
include scripts/make-rules/common.mk 

## --------------------------------------
## Binaries
## --------------------------------------

##@ build:

.PHONY: build
build:
	@$(MAKE) go.build

.PHONY: go.build
go.build: ## 编译源码.
	@echo "Building Herman Blog..."
	@go build -o _output/hermanblog -v cmd/hermanblog/main.go
