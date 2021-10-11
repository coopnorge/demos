# cdk8s

https://cdk8s.io/

### define kubernetes templates using
* golang
* python
* typescript
* java

### cdk

cdk = cloud development toolkit 

created by AWS https://aws.amazon.com/cdk/

current implementations
* cdk8s (kubernetes)
* cdktf (terraform)

### Why

* you don't need to learn yaml
* can test with language specific testing tools
* IDE's generally perform better with 'real' programming languages
* Integrate chart with code resources

### What about

**kustomize**: kustomize is more of a patching engine not a templating engine.

**helm**: similar, but yaml, dsl has limitations

**pulumi**: looks great! (but not free/opensource)


### Demo

For demo code, check `main.go`

To start a new project
```console
cdk8s init go-app
```

To render templates

```
cdk8s sync # (-p) or just go run main.go 
```

To import CDRs

```
crd cdk8s import istio/istio_crds.yaml -l go
```
