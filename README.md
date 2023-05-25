# Supabase Terraform Provider

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.18

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To see which other commands are available, run `make help` to see the possible make targets with descriptions.

## References

- [Supabase CLI](https://github.com/supabase/cli) (provides the HTTP client)
- [Management API Reference](https://supabase.com/docs/reference/api/introduction)
