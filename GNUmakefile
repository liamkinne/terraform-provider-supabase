default: testacc

# Pull OpenAPI spec and prettify JSON
.PHONY: pull-spec
pull-spec:
	curl https://api.supabase.com/api/v1-json | jq . > internal/client/spec-v1.json

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
