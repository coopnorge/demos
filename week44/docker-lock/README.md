# Demo of docker-lock

[docker-lock](https://github.com/safe-waters/docker-lock) is a tool to help manage Docker image digests.

Key features:

- Add missing digests to Dockerfiles, docker-compose files, and K8S manifests
- Update existing digests

## Demo


```bash
make clean
make setup
```

### Example 1: Adding digests

```bash
cd var/demo/000-nodigest

# generate a docker-lock.json file with digests
docker-lock lock generate
# rewrite Docker image references to include digests
docker-lock lock rewrite

# Docker image references in Dockerfile now has digests
diff --color --new-file -u -r ../../../share/states/000-nodigest/ ./

cd ../../../
```

### Example 2: Updating outdated digests

```bash
cd var/demo/010-olddigest

# generate and rewrite makes no change
docker-lock lock generate
docker-lock lock rewrite
diff --color -u -r ../../../share/states/010-olddigest/ ./

# generate with --update-existing-digests queries registries for new digests
docker-lock lock generate --update-existing-digests
# rewrite Docker image references to include new digests
docker-lock lock rewrite

# Docker image references in Dockerfile are now updated
diff --color --new-file -u -r ../../../share/states/010-olddigest/ ./

cd ../../../
```
