# Demo of github.com/c2fo/vfs

This directory contains a demo batch processor that reads files from a location
and then creaters marker files containing file checksums for each file that was
found.

The batch processor uses c2fo/vfs and can thus work with a variety of backends
including OS, Google Cloud Storage, Azure File Storage, S3 and SFTP.

This makes it easier to write logic that is not tied to a specific storage
backend and provides more options for testing applications that will use cloud
storage in production.

```bash
# Build the latest devtools environment
docker-compose build

# Check available targets
docker-compose run --rm devtools make help

# Build and test
docker-compose run --rm devtools make test gotest_args=-v

# Run on a OS location (file-uri)
ls -l test/data/
go run ./cmd/govfsdemo/ run --url "file://$(pwd)/test/data/" --mark "$(date +%Y%m%d%H%M)"
\rm -v test/data/*marker*

# Run on a GCP storage bucket
gsutil ls gs://coopnorge-iwana-dev-pharnavaz/
go run ./cmd/govfsdemo/ run --url "gs://coopnorge-iwana-dev-pharnavaz/" --mark "$(date +%Y%m%d%H%M)"
gsutil rm -v gs://coopnorge-iwana-dev-pharnavaz/*marker*
```

