
API_SCHEMA_FILE = schema/apidata.js


.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


.PHONY: tools
tools: ## Build tools
	@mkdir -p bin
	@go build tools/gen-api -o bin/gen-api

.PHONY: gen-api
gen-api: ## Generate api by schema file
	@go run tools/gen-api/main.go --schema-file $(API_SCHEMA_FILE) --output-dir pve


.PHONY: download-schema
download-schema: ## Download api schema file from source
	@mkdir -p $(shell dirname $(API_SCHEMA_FILE))
	@curl -L https://github.com/proxmox/pve-docs/raw/master/api-viewer/apidata.js -o $(API_SCHEMA_FILE)

