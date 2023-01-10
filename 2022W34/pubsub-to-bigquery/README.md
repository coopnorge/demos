# Pub/Sub to BigQuery demo

- [Whiteboard](https://www.figma.com/file/XmnOfhV5h0UeiOsGeMm1P3/PubSub-to-BigQuery-Demo---2022W34)


```bash
# Build the devtools images
docker compose build --progress plain

# Adjust the example files:
cp terraform/terraform.tfvars.example terraform/terraform.tfvars
cp .env.yaml.example .env.yaml

# validate and apply terraform
docker compose run --rm terraform-devtools validate
docker compose run --rm terraform-devtools terraform -chdir=terraform apply

# Clean the state for demo:
docker compose run --rm demo-devtools task demo:clean

# Run schemaless demo:
docker compose run --rm demo-devtools task demo:schemaless

# Run cloudevents schema demo:
docker compose run --rm demo-devtools task demo:cloudevents
```
