# Repro for knative/client-pkg#166

## Steps to reproduce

1. Checkout this code
1. Run test: `go run github.com/google/go-licenses@latest check ./...`
1. You should license check failing, with error:
   ```
   Unknown license type  found for library github.com/mattn/go-localereader
   ```