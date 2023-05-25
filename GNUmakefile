default: testacc

.PHONY: help
help: ## Self documenting help output
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


.PHONY: pull-spec
pull-spec: ## Pull OpenAPI spec and prettify JSON
	curl https://api.supabase.com/api/v1-json | jq . > internal/client/spec-v1.json

.PHONY: gen
gen: ## Generate docs and openapi client code
	go generate ./... -v

.PHONY: testacc
testacc: ## Run acceptance tests
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
