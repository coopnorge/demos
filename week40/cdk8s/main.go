package main

import (
	"example.com/cdk8s/imports/networkingistioio"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus20"
)

// commands
// init: cdk8s init go-app
// render: cdk8s sync (-p) or just go run main.go
// import: crd cdk8s import istio/istio_crds.yaml -l go
func main() {
	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir:         jsii.String("templates"),
		YamlOutputType: cdk8s.YamlOutputType_FILE_PER_RESOURCE,
	})

	chart := cdk8s.NewChart(app, jsii.String("cdk8s-demo"), &cdk8s.ChartProps{
		Labels: &map[string]*string{
			"app": jsii.String("cdk8s-demo"),
		},
	})

	//deploy :=
	deploy := cdk8splus20.NewDeployment(chart, jsii.String("deploy"), &cdk8splus20.DeploymentProps{
		Replicas: jsii.Number(3),
		//	DefaultSelector: jsii.Bool(false),
		Containers: &[]*cdk8splus20.ContainerProps{{
			Image: jsii.String("ubuntu"),
			Name:  jsii.String("ubuntu"),
			Liveness: cdk8splus20.Probe_FromHttpGet(jsii.String("/health"),
				&cdk8splus20.HttpGetProbeOptions{
					InitialDelaySeconds: cdk8s.Duration_Seconds(jsii.Number(30)),
				},
			),
			Port: jsii.Number(8080),
			Env: &map[string]cdk8splus20.EnvValue{
				"env1": cdk8splus20.EnvValue_FromValue(jsii.String("valone")),
				"env2": cdk8splus20.EnvValue_FromValue(jsii.String("valtwo")),
			},
		}},
	})

	//deploy.SelectByLabel(jsii.String("klaas"), jsii.String("jan"))
	//deploy.SelectByLabel(jsii.String("piet"), jsii.String("karel"))

	svc := deploy.Expose(jsii.Number(883), &cdk8splus20.ExposeOptions{})

	ing := cdk8splus20.NewIngressV1Beta1(chart, jsii.String("ing"), &cdk8splus20.IngressV1Beta1Props{})
	ing.AddHostDefaultBackend(
		jsii.String("host"),
		cdk8splus20.IngressV1Beta1Backend_FromService(
			svc,
			&cdk8splus20.ServiceIngressV1BetaBackendOptions{},
		))

	networkingistioio.NewVirtualService(chart, jsii.String("vs"), &networkingistioio.VirtualServiceProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("test-vs"),
		},
		Spec: &networkingistioio.VirtualServiceSpec{
			Gateways: jsii.Strings("test-gateway"),
			Hosts:    jsii.Strings("bla.com", "boe.com"),
			Http: &[]*networkingistioio.VirtualServiceSpecHttp{
				{
					Match: &[]*networkingistioio.VirtualServiceSpecHttpMatch{
						{
							Uri: &networkingistioio.VirtualServiceSpecHttpMatchUri{
								Prefix: jsii.String("/test"),
							},
						},
					},
					Route: &[]*networkingistioio.VirtualServiceSpecHttpRoute{
						{
							Destination: &networkingistioio.VirtualServiceSpecHttpRouteDestination{
								Host: svc.Name(),
								Port: &networkingistioio.VirtualServiceSpecHttpRouteDestinationPort{
									(*svc.Ports())[0].Port,
								},
							},
						},
					},
				},
			},
		},
	})

	app.Synth()
}
