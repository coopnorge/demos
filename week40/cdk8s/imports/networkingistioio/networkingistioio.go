// networkingistioio
package networkingistioio

import (
	_init_ "example.com/cdk8s/imports/networkingistioio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"example.com/cdk8s/imports/networkingistioio/internal"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
)

type DestinationRule interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for DestinationRule
type jsiiProxy_DestinationRule struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_DestinationRule) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DestinationRule) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DestinationRule) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DestinationRule) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DestinationRule) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DestinationRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "DestinationRule" API object.
func NewDestinationRule(scope constructs.Construct, id *string, props *DestinationRuleProps) DestinationRule {
	_init_.Initialize()

	j := jsiiProxy_DestinationRule{}

	_jsii_.Create(
		"networkingistioio.DestinationRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "DestinationRule" API object.
func NewDestinationRule_Override(d DestinationRule, scope constructs.Construct, id *string, props *DestinationRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"networkingistioio.DestinationRule",
		[]interface{}{scope, id, props},
		d,
	)
}

// Renders a Kubernetes manifest for "DestinationRule".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func DestinationRule_Manifest(props *DestinationRuleProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"networkingistioio.DestinationRule",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func DestinationRule_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"networkingistioio.DestinationRule",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func DestinationRule_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"networkingistioio.DestinationRule",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (d *jsiiProxy_DestinationRule) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (d *jsiiProxy_DestinationRule) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (d *jsiiProxy_DestinationRule) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (d *jsiiProxy_DestinationRule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DestinationRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (d *jsiiProxy_DestinationRule) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DestinationRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DestinationRuleProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Configuration affecting load balancing, outlier detection, etc.
	//
	// See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html
	Spec *DestinationRuleSpec `json:"spec"`
}

// Configuration affecting load balancing, outlier detection, etc.
//
// See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html
type DestinationRuleSpec struct {
	// A list of namespaces to which this destination rule is exported.
	ExportTo *[]*string `json:"exportTo"`
	// The name of a service from the service registry.
	Host          *string                           `json:"host"`
	Subsets       *[]*DestinationRuleSpecSubsets    `json:"subsets"`
	TrafficPolicy *DestinationRuleSpecTrafficPolicy `json:"trafficPolicy"`
}

type DestinationRuleSpecSubsets struct {
	Labels *map[string]*string `json:"labels"`
	// Name of the subset.
	Name *string `json:"name"`
	// Traffic policies that apply to this subset.
	TrafficPolicy *DestinationRuleSpecSubsetsTrafficPolicy `json:"trafficPolicy"`
}

// Traffic policies that apply to this subset.
type DestinationRuleSpecSubsetsTrafficPolicy struct {
	ConnectionPool *DestinationRuleSpecSubsetsTrafficPolicyConnectionPool `json:"connectionPool"`
	// Settings controlling the load balancer algorithms.
	LoadBalancer     *DestinationRuleSpecSubsetsTrafficPolicyLoadBalancer     `json:"loadBalancer"`
	OutlierDetection *DestinationRuleSpecSubsetsTrafficPolicyOutlierDetection `json:"outlierDetection"`
	// Traffic policies specific to individual ports.
	PortLevelSettings *[]*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettings `json:"portLevelSettings"`
	// TLS related settings for connections to the upstream service.
	Tls *DestinationRuleSpecSubsetsTrafficPolicyTls `json:"tls"`
}

type DestinationRuleSpecSubsetsTrafficPolicyConnectionPool struct {
	// HTTP connection pool settings.
	Http *DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttp `json:"http"`
	// Settings common to both HTTP and TCP upstream connections.
	Tcp *DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcp `json:"tcp"`
}

// HTTP connection pool settings.
type DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttp struct {
	// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
	H2UpgradePolicy DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy `json:"h2UpgradePolicy"`
	// Maximum number of pending HTTP requests to a destination.
	Http1MaxPendingRequests *float64 `json:"http1MaxPendingRequests"`
	// Maximum number of requests to a backend.
	Http2MaxRequests *float64 `json:"http2MaxRequests"`
	// The idle timeout for upstream connection pool connections.
	IdleTimeout *string `json:"idleTimeout"`
	// Maximum number of requests per connection to a backend.
	MaxRequestsPerConnection *float64 `json:"maxRequestsPerConnection"`
	MaxRetries               *float64 `json:"maxRetries"`
	// If set to true, client protocol will be preserved while initiating connection to backend.
	UseClientProtocol *bool `json:"useClientProtocol"`
}

// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
type DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy string

const (
	DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DEFAULT        DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy = "DEFAULT"
	DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy = "DO_NOT_UPGRADE"
	DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_UPGRADE        DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy = "UPGRADE"
)

// Settings common to both HTTP and TCP upstream connections.
type DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcp struct {
	// TCP connection timeout.
	ConnectTimeout *string `json:"connectTimeout"`
	// Maximum number of HTTP1 /TCP connections to a destination host.
	MaxConnections *float64 `json:"maxConnections"`
	// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
	TcpKeepalive *DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcpTcpKeepalive `json:"tcpKeepalive"`
}

// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
type DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcpTcpKeepalive struct {
	// The time duration between keep-alive probes.
	Interval *string  `json:"interval"`
	Probes   *float64 `json:"probes"`
	Time     *string  `json:"time"`
}

// Settings controlling the load balancer algorithms.
type DestinationRuleSpecSubsetsTrafficPolicyLoadBalancer struct {
	ConsistentHash    *DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHash    `json:"consistentHash"`
	LocalityLbSetting *DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSetting `json:"localityLbSetting"`
	Simple            DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple             `json:"simple"`
}

type DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHash struct {
	// Hash based on HTTP cookie.
	HttpCookie *DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHashHttpCookie `json:"httpCookie"`
	// Hash based on a specific HTTP header.
	HttpHeaderName *string `json:"httpHeaderName"`
	// Hash based on a specific HTTP query parameter.
	HttpQueryParameterName *string  `json:"httpQueryParameterName"`
	MinimumRingSize        *float64 `json:"minimumRingSize"`
	// Hash based on the source IP address.
	UseSourceIp *bool `json:"useSourceIp"`
}

// Hash based on HTTP cookie.
type DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHashHttpCookie struct {
	// Name of the cookie.
	Name *string `json:"name"`
	// Path to set for the cookie.
	Path *string `json:"path"`
	// Lifetime of the cookie.
	Ttl *string `json:"ttl"`
}

type DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSetting struct {
	// Optional: only one of distribute or failover can be set.
	Distribute *[]*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingDistribute `json:"distribute"`
	// enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.
	Enabled *bool `json:"enabled"`
	// Optional: only failover or distribute can be set.
	Failover *[]*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingFailover `json:"failover"`
}

type DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingDistribute struct {
	// Originating locality, '/' separated, e.g.
	From *string `json:"from"`
	// Map of upstream localities to traffic distribution weights.
	To *map[string]*float64 `json:"to"`
}

type DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingFailover struct {
	// Originating region.
	From *string `json:"from"`
	To   *string `json:"to"`
}

type DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple string

const (
	DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_ROUND_ROBIN DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple = "ROUND_ROBIN"
	DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_LEAST_CONN  DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple = "LEAST_CONN"
	DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_RANDOM      DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple = "RANDOM"
	DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_PASSTHROUGH DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple = "PASSTHROUGH"
)

type DestinationRuleSpecSubsetsTrafficPolicyOutlierDetection struct {
	// Minimum ejection duration.
	BaseEjectionTime *string `json:"baseEjectionTime"`
	// Number of 5xx errors before a host is ejected from the connection pool.
	Consecutive5XxErrors *float64 `json:"consecutive5XxErrors"`
	ConsecutiveErrors    *float64 `json:"consecutiveErrors"`
	// Number of gateway errors before a host is ejected from the connection pool.
	ConsecutiveGatewayErrors *float64 `json:"consecutiveGatewayErrors"`
	// Time interval between ejection sweep analysis.
	Interval           *string  `json:"interval"`
	MaxEjectionPercent *float64 `json:"maxEjectionPercent"`
	MinHealthPercent   *float64 `json:"minHealthPercent"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettings struct {
	ConnectionPool *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPool `json:"connectionPool"`
	// Settings controlling the load balancer algorithms.
	LoadBalancer     *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancer     `json:"loadBalancer"`
	OutlierDetection *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsOutlierDetection `json:"outlierDetection"`
	Port             *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsPort             `json:"port"`
	// TLS related settings for connections to the upstream service.
	Tls *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTls `json:"tls"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPool struct {
	// HTTP connection pool settings.
	Http *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttp `json:"http"`
	// Settings common to both HTTP and TCP upstream connections.
	Tcp *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcp `json:"tcp"`
}

// HTTP connection pool settings.
type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttp struct {
	// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
	H2UpgradePolicy DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy `json:"h2UpgradePolicy"`
	// Maximum number of pending HTTP requests to a destination.
	Http1MaxPendingRequests *float64 `json:"http1MaxPendingRequests"`
	// Maximum number of requests to a backend.
	Http2MaxRequests *float64 `json:"http2MaxRequests"`
	// The idle timeout for upstream connection pool connections.
	IdleTimeout *string `json:"idleTimeout"`
	// Maximum number of requests per connection to a backend.
	MaxRequestsPerConnection *float64 `json:"maxRequestsPerConnection"`
	MaxRetries               *float64 `json:"maxRetries"`
	// If set to true, client protocol will be preserved while initiating connection to backend.
	UseClientProtocol *bool `json:"useClientProtocol"`
}

// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy string

const (
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DEFAULT        DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy = "DEFAULT"
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy = "DO_NOT_UPGRADE"
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_UPGRADE        DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy = "UPGRADE"
)

// Settings common to both HTTP and TCP upstream connections.
type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcp struct {
	// TCP connection timeout.
	ConnectTimeout *string `json:"connectTimeout"`
	// Maximum number of HTTP1 /TCP connections to a destination host.
	MaxConnections *float64 `json:"maxConnections"`
	// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
	TcpKeepalive *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive `json:"tcpKeepalive"`
}

// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive struct {
	// The time duration between keep-alive probes.
	Interval *string  `json:"interval"`
	Probes   *float64 `json:"probes"`
	Time     *string  `json:"time"`
}

// Settings controlling the load balancer algorithms.
type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancer struct {
	ConsistentHash    *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash    `json:"consistentHash"`
	LocalityLbSetting *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting `json:"localityLbSetting"`
	Simple            DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple             `json:"simple"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash struct {
	// Hash based on HTTP cookie.
	HttpCookie *DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie `json:"httpCookie"`
	// Hash based on a specific HTTP header.
	HttpHeaderName *string `json:"httpHeaderName"`
	// Hash based on a specific HTTP query parameter.
	HttpQueryParameterName *string  `json:"httpQueryParameterName"`
	MinimumRingSize        *float64 `json:"minimumRingSize"`
	// Hash based on the source IP address.
	UseSourceIp *bool `json:"useSourceIp"`
}

// Hash based on HTTP cookie.
type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie struct {
	// Name of the cookie.
	Name *string `json:"name"`
	// Path to set for the cookie.
	Path *string `json:"path"`
	// Lifetime of the cookie.
	Ttl *string `json:"ttl"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting struct {
	// Optional: only one of distribute or failover can be set.
	Distribute *[]*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute `json:"distribute"`
	// enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.
	Enabled *bool `json:"enabled"`
	// Optional: only failover or distribute can be set.
	Failover *[]*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover `json:"failover"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute struct {
	// Originating locality, '/' separated, e.g.
	From *string `json:"from"`
	// Map of upstream localities to traffic distribution weights.
	To *map[string]*float64 `json:"to"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover struct {
	// Originating region.
	From *string `json:"from"`
	To   *string `json:"to"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple string

const (
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_ROUND_ROBIN DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple = "ROUND_ROBIN"
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_CONN  DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple = "LEAST_CONN"
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_RANDOM      DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple = "RANDOM"
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_PASSTHROUGH DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple = "PASSTHROUGH"
)

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsOutlierDetection struct {
	// Minimum ejection duration.
	BaseEjectionTime *string `json:"baseEjectionTime"`
	// Number of 5xx errors before a host is ejected from the connection pool.
	Consecutive5XxErrors *float64 `json:"consecutive5XxErrors"`
	ConsecutiveErrors    *float64 `json:"consecutiveErrors"`
	// Number of gateway errors before a host is ejected from the connection pool.
	ConsecutiveGatewayErrors *float64 `json:"consecutiveGatewayErrors"`
	// Time interval between ejection sweep analysis.
	Interval           *string  `json:"interval"`
	MaxEjectionPercent *float64 `json:"maxEjectionPercent"`
	MinHealthPercent   *float64 `json:"minHealthPercent"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsPort struct {
	Number *float64 `json:"number"`
}

// TLS related settings for connections to the upstream service.
type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTls struct {
	CaCertificates *string `json:"caCertificates"`
	// REQUIRED if mode is `MUTUAL`.
	ClientCertificate *string                                                         `json:"clientCertificate"`
	CredentialName    *string                                                         `json:"credentialName"`
	Mode              DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode `json:"mode"`
	// REQUIRED if mode is `MUTUAL`.
	PrivateKey *string `json:"privateKey"`
	// SNI string to present to the server during TLS handshake.
	Sni             *string    `json:"sni"`
	SubjectAltNames *[]*string `json:"subjectAltNames"`
}

type DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode string

const (
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_DISABLE      DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode = "DISABLE"
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_SIMPLE       DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode = "SIMPLE"
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_MUTUAL       DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode = "MUTUAL"
	DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_ISTIO_MUTUAL DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode = "ISTIO_MUTUAL"
)

// TLS related settings for connections to the upstream service.
type DestinationRuleSpecSubsetsTrafficPolicyTls struct {
	CaCertificates *string `json:"caCertificates"`
	// REQUIRED if mode is `MUTUAL`.
	ClientCertificate *string                                        `json:"clientCertificate"`
	CredentialName    *string                                        `json:"credentialName"`
	Mode              DestinationRuleSpecSubsetsTrafficPolicyTlsMode `json:"mode"`
	// REQUIRED if mode is `MUTUAL`.
	PrivateKey *string `json:"privateKey"`
	// SNI string to present to the server during TLS handshake.
	Sni             *string    `json:"sni"`
	SubjectAltNames *[]*string `json:"subjectAltNames"`
}

type DestinationRuleSpecSubsetsTrafficPolicyTlsMode string

const (
	DestinationRuleSpecSubsetsTrafficPolicyTlsMode_DISABLE      DestinationRuleSpecSubsetsTrafficPolicyTlsMode = "DISABLE"
	DestinationRuleSpecSubsetsTrafficPolicyTlsMode_SIMPLE       DestinationRuleSpecSubsetsTrafficPolicyTlsMode = "SIMPLE"
	DestinationRuleSpecSubsetsTrafficPolicyTlsMode_MUTUAL       DestinationRuleSpecSubsetsTrafficPolicyTlsMode = "MUTUAL"
	DestinationRuleSpecSubsetsTrafficPolicyTlsMode_ISTIO_MUTUAL DestinationRuleSpecSubsetsTrafficPolicyTlsMode = "ISTIO_MUTUAL"
)

type DestinationRuleSpecTrafficPolicy struct {
	ConnectionPool *DestinationRuleSpecTrafficPolicyConnectionPool `json:"connectionPool"`
	// Settings controlling the load balancer algorithms.
	LoadBalancer     *DestinationRuleSpecTrafficPolicyLoadBalancer     `json:"loadBalancer"`
	OutlierDetection *DestinationRuleSpecTrafficPolicyOutlierDetection `json:"outlierDetection"`
	// Traffic policies specific to individual ports.
	PortLevelSettings *[]*DestinationRuleSpecTrafficPolicyPortLevelSettings `json:"portLevelSettings"`
	// TLS related settings for connections to the upstream service.
	Tls *DestinationRuleSpecTrafficPolicyTls `json:"tls"`
}

type DestinationRuleSpecTrafficPolicyConnectionPool struct {
	// HTTP connection pool settings.
	Http *DestinationRuleSpecTrafficPolicyConnectionPoolHttp `json:"http"`
	// Settings common to both HTTP and TCP upstream connections.
	Tcp *DestinationRuleSpecTrafficPolicyConnectionPoolTcp `json:"tcp"`
}

// HTTP connection pool settings.
type DestinationRuleSpecTrafficPolicyConnectionPoolHttp struct {
	// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
	H2UpgradePolicy DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy `json:"h2UpgradePolicy"`
	// Maximum number of pending HTTP requests to a destination.
	Http1MaxPendingRequests *float64 `json:"http1MaxPendingRequests"`
	// Maximum number of requests to a backend.
	Http2MaxRequests *float64 `json:"http2MaxRequests"`
	// The idle timeout for upstream connection pool connections.
	IdleTimeout *string `json:"idleTimeout"`
	// Maximum number of requests per connection to a backend.
	MaxRequestsPerConnection *float64 `json:"maxRequestsPerConnection"`
	MaxRetries               *float64 `json:"maxRetries"`
	// If set to true, client protocol will be preserved while initiating connection to backend.
	UseClientProtocol *bool `json:"useClientProtocol"`
}

// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
type DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy string

const (
	DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DEFAULT        DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy = "DEFAULT"
	DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy = "DO_NOT_UPGRADE"
	DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_UPGRADE        DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy = "UPGRADE"
)

// Settings common to both HTTP and TCP upstream connections.
type DestinationRuleSpecTrafficPolicyConnectionPoolTcp struct {
	// TCP connection timeout.
	ConnectTimeout *string `json:"connectTimeout"`
	// Maximum number of HTTP1 /TCP connections to a destination host.
	MaxConnections *float64 `json:"maxConnections"`
	// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
	TcpKeepalive *DestinationRuleSpecTrafficPolicyConnectionPoolTcpTcpKeepalive `json:"tcpKeepalive"`
}

// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
type DestinationRuleSpecTrafficPolicyConnectionPoolTcpTcpKeepalive struct {
	// The time duration between keep-alive probes.
	Interval *string  `json:"interval"`
	Probes   *float64 `json:"probes"`
	Time     *string  `json:"time"`
}

// Settings controlling the load balancer algorithms.
type DestinationRuleSpecTrafficPolicyLoadBalancer struct {
	ConsistentHash    *DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHash    `json:"consistentHash"`
	LocalityLbSetting *DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSetting `json:"localityLbSetting"`
	Simple            DestinationRuleSpecTrafficPolicyLoadBalancerSimple             `json:"simple"`
}

type DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHash struct {
	// Hash based on HTTP cookie.
	HttpCookie *DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHashHttpCookie `json:"httpCookie"`
	// Hash based on a specific HTTP header.
	HttpHeaderName *string `json:"httpHeaderName"`
	// Hash based on a specific HTTP query parameter.
	HttpQueryParameterName *string  `json:"httpQueryParameterName"`
	MinimumRingSize        *float64 `json:"minimumRingSize"`
	// Hash based on the source IP address.
	UseSourceIp *bool `json:"useSourceIp"`
}

// Hash based on HTTP cookie.
type DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHashHttpCookie struct {
	// Name of the cookie.
	Name *string `json:"name"`
	// Path to set for the cookie.
	Path *string `json:"path"`
	// Lifetime of the cookie.
	Ttl *string `json:"ttl"`
}

type DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSetting struct {
	// Optional: only one of distribute or failover can be set.
	Distribute *[]*DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingDistribute `json:"distribute"`
	// enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.
	Enabled *bool `json:"enabled"`
	// Optional: only failover or distribute can be set.
	Failover *[]*DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingFailover `json:"failover"`
}

type DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingDistribute struct {
	// Originating locality, '/' separated, e.g.
	From *string `json:"from"`
	// Map of upstream localities to traffic distribution weights.
	To *map[string]*float64 `json:"to"`
}

type DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingFailover struct {
	// Originating region.
	From *string `json:"from"`
	To   *string `json:"to"`
}

type DestinationRuleSpecTrafficPolicyLoadBalancerSimple string

const (
	DestinationRuleSpecTrafficPolicyLoadBalancerSimple_ROUND_ROBIN DestinationRuleSpecTrafficPolicyLoadBalancerSimple = "ROUND_ROBIN"
	DestinationRuleSpecTrafficPolicyLoadBalancerSimple_LEAST_CONN  DestinationRuleSpecTrafficPolicyLoadBalancerSimple = "LEAST_CONN"
	DestinationRuleSpecTrafficPolicyLoadBalancerSimple_RANDOM      DestinationRuleSpecTrafficPolicyLoadBalancerSimple = "RANDOM"
	DestinationRuleSpecTrafficPolicyLoadBalancerSimple_PASSTHROUGH DestinationRuleSpecTrafficPolicyLoadBalancerSimple = "PASSTHROUGH"
)

type DestinationRuleSpecTrafficPolicyOutlierDetection struct {
	// Minimum ejection duration.
	BaseEjectionTime *string `json:"baseEjectionTime"`
	// Number of 5xx errors before a host is ejected from the connection pool.
	Consecutive5XxErrors *float64 `json:"consecutive5XxErrors"`
	ConsecutiveErrors    *float64 `json:"consecutiveErrors"`
	// Number of gateway errors before a host is ejected from the connection pool.
	ConsecutiveGatewayErrors *float64 `json:"consecutiveGatewayErrors"`
	// Time interval between ejection sweep analysis.
	Interval           *string  `json:"interval"`
	MaxEjectionPercent *float64 `json:"maxEjectionPercent"`
	MinHealthPercent   *float64 `json:"minHealthPercent"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettings struct {
	ConnectionPool *DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPool `json:"connectionPool"`
	// Settings controlling the load balancer algorithms.
	LoadBalancer     *DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancer     `json:"loadBalancer"`
	OutlierDetection *DestinationRuleSpecTrafficPolicyPortLevelSettingsOutlierDetection `json:"outlierDetection"`
	Port             *DestinationRuleSpecTrafficPolicyPortLevelSettingsPort             `json:"port"`
	// TLS related settings for connections to the upstream service.
	Tls *DestinationRuleSpecTrafficPolicyPortLevelSettingsTls `json:"tls"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPool struct {
	// HTTP connection pool settings.
	Http *DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttp `json:"http"`
	// Settings common to both HTTP and TCP upstream connections.
	Tcp *DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcp `json:"tcp"`
}

// HTTP connection pool settings.
type DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttp struct {
	// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
	H2UpgradePolicy DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy `json:"h2UpgradePolicy"`
	// Maximum number of pending HTTP requests to a destination.
	Http1MaxPendingRequests *float64 `json:"http1MaxPendingRequests"`
	// Maximum number of requests to a backend.
	Http2MaxRequests *float64 `json:"http2MaxRequests"`
	// The idle timeout for upstream connection pool connections.
	IdleTimeout *string `json:"idleTimeout"`
	// Maximum number of requests per connection to a backend.
	MaxRequestsPerConnection *float64 `json:"maxRequestsPerConnection"`
	MaxRetries               *float64 `json:"maxRetries"`
	// If set to true, client protocol will be preserved while initiating connection to backend.
	UseClientProtocol *bool `json:"useClientProtocol"`
}

// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
type DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy string

const (
	DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DEFAULT        DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy = "DEFAULT"
	DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy = "DO_NOT_UPGRADE"
	DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_UPGRADE        DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy = "UPGRADE"
)

// Settings common to both HTTP and TCP upstream connections.
type DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcp struct {
	// TCP connection timeout.
	ConnectTimeout *string `json:"connectTimeout"`
	// Maximum number of HTTP1 /TCP connections to a destination host.
	MaxConnections *float64 `json:"maxConnections"`
	// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
	TcpKeepalive *DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive `json:"tcpKeepalive"`
}

// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
type DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive struct {
	// The time duration between keep-alive probes.
	Interval *string  `json:"interval"`
	Probes   *float64 `json:"probes"`
	Time     *string  `json:"time"`
}

// Settings controlling the load balancer algorithms.
type DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancer struct {
	ConsistentHash    *DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash    `json:"consistentHash"`
	LocalityLbSetting *DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting `json:"localityLbSetting"`
	Simple            DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple             `json:"simple"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash struct {
	// Hash based on HTTP cookie.
	HttpCookie *DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie `json:"httpCookie"`
	// Hash based on a specific HTTP header.
	HttpHeaderName *string `json:"httpHeaderName"`
	// Hash based on a specific HTTP query parameter.
	HttpQueryParameterName *string  `json:"httpQueryParameterName"`
	MinimumRingSize        *float64 `json:"minimumRingSize"`
	// Hash based on the source IP address.
	UseSourceIp *bool `json:"useSourceIp"`
}

// Hash based on HTTP cookie.
type DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie struct {
	// Name of the cookie.
	Name *string `json:"name"`
	// Path to set for the cookie.
	Path *string `json:"path"`
	// Lifetime of the cookie.
	Ttl *string `json:"ttl"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting struct {
	// Optional: only one of distribute or failover can be set.
	Distribute *[]*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute `json:"distribute"`
	// enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.
	Enabled *bool `json:"enabled"`
	// Optional: only failover or distribute can be set.
	Failover *[]*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover `json:"failover"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute struct {
	// Originating locality, '/' separated, e.g.
	From *string `json:"from"`
	// Map of upstream localities to traffic distribution weights.
	To *map[string]*float64 `json:"to"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover struct {
	// Originating region.
	From *string `json:"from"`
	To   *string `json:"to"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple string

const (
	DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_ROUND_ROBIN DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple = "ROUND_ROBIN"
	DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_CONN  DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple = "LEAST_CONN"
	DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_RANDOM      DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple = "RANDOM"
	DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_PASSTHROUGH DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple = "PASSTHROUGH"
)

type DestinationRuleSpecTrafficPolicyPortLevelSettingsOutlierDetection struct {
	// Minimum ejection duration.
	BaseEjectionTime *string `json:"baseEjectionTime"`
	// Number of 5xx errors before a host is ejected from the connection pool.
	Consecutive5XxErrors *float64 `json:"consecutive5XxErrors"`
	ConsecutiveErrors    *float64 `json:"consecutiveErrors"`
	// Number of gateway errors before a host is ejected from the connection pool.
	ConsecutiveGatewayErrors *float64 `json:"consecutiveGatewayErrors"`
	// Time interval between ejection sweep analysis.
	Interval           *string  `json:"interval"`
	MaxEjectionPercent *float64 `json:"maxEjectionPercent"`
	MinHealthPercent   *float64 `json:"minHealthPercent"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettingsPort struct {
	Number *float64 `json:"number"`
}

// TLS related settings for connections to the upstream service.
type DestinationRuleSpecTrafficPolicyPortLevelSettingsTls struct {
	CaCertificates *string `json:"caCertificates"`
	// REQUIRED if mode is `MUTUAL`.
	ClientCertificate *string                                                  `json:"clientCertificate"`
	CredentialName    *string                                                  `json:"credentialName"`
	Mode              DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode `json:"mode"`
	// REQUIRED if mode is `MUTUAL`.
	PrivateKey *string `json:"privateKey"`
	// SNI string to present to the server during TLS handshake.
	Sni             *string    `json:"sni"`
	SubjectAltNames *[]*string `json:"subjectAltNames"`
}

type DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode string

const (
	DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_DISABLE      DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode = "DISABLE"
	DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_SIMPLE       DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode = "SIMPLE"
	DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_MUTUAL       DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode = "MUTUAL"
	DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_ISTIO_MUTUAL DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode = "ISTIO_MUTUAL"
)

// TLS related settings for connections to the upstream service.
type DestinationRuleSpecTrafficPolicyTls struct {
	CaCertificates *string `json:"caCertificates"`
	// REQUIRED if mode is `MUTUAL`.
	ClientCertificate *string                                 `json:"clientCertificate"`
	CredentialName    *string                                 `json:"credentialName"`
	Mode              DestinationRuleSpecTrafficPolicyTlsMode `json:"mode"`
	// REQUIRED if mode is `MUTUAL`.
	PrivateKey *string `json:"privateKey"`
	// SNI string to present to the server during TLS handshake.
	Sni             *string    `json:"sni"`
	SubjectAltNames *[]*string `json:"subjectAltNames"`
}

type DestinationRuleSpecTrafficPolicyTlsMode string

const (
	DestinationRuleSpecTrafficPolicyTlsMode_DISABLE      DestinationRuleSpecTrafficPolicyTlsMode = "DISABLE"
	DestinationRuleSpecTrafficPolicyTlsMode_SIMPLE       DestinationRuleSpecTrafficPolicyTlsMode = "SIMPLE"
	DestinationRuleSpecTrafficPolicyTlsMode_MUTUAL       DestinationRuleSpecTrafficPolicyTlsMode = "MUTUAL"
	DestinationRuleSpecTrafficPolicyTlsMode_ISTIO_MUTUAL DestinationRuleSpecTrafficPolicyTlsMode = "ISTIO_MUTUAL"
)

type EnvoyFilter interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for EnvoyFilter
type jsiiProxy_EnvoyFilter struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_EnvoyFilter) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvoyFilter) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvoyFilter) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvoyFilter) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvoyFilter) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvoyFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "EnvoyFilter" API object.
func NewEnvoyFilter(scope constructs.Construct, id *string, props *EnvoyFilterProps) EnvoyFilter {
	_init_.Initialize()

	j := jsiiProxy_EnvoyFilter{}

	_jsii_.Create(
		"networkingistioio.EnvoyFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "EnvoyFilter" API object.
func NewEnvoyFilter_Override(e EnvoyFilter, scope constructs.Construct, id *string, props *EnvoyFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"networkingistioio.EnvoyFilter",
		[]interface{}{scope, id, props},
		e,
	)
}

// Renders a Kubernetes manifest for "EnvoyFilter".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func EnvoyFilter_Manifest(props *EnvoyFilterProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"networkingistioio.EnvoyFilter",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func EnvoyFilter_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"networkingistioio.EnvoyFilter",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func EnvoyFilter_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"networkingistioio.EnvoyFilter",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (e *jsiiProxy_EnvoyFilter) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (e *jsiiProxy_EnvoyFilter) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (e *jsiiProxy_EnvoyFilter) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (e *jsiiProxy_EnvoyFilter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (e *jsiiProxy_EnvoyFilter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (e *jsiiProxy_EnvoyFilter) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EnvoyFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EnvoyFilterProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Customizing Envoy configuration generated by Istio.
	//
	// See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html
	Spec *EnvoyFilterSpec `json:"spec"`
}

// Customizing Envoy configuration generated by Istio.
//
// See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html
type EnvoyFilterSpec struct {
	// One or more patches with match conditions.
	ConfigPatches    *[]*EnvoyFilterSpecConfigPatches `json:"configPatches"`
	WorkloadSelector *EnvoyFilterSpecWorkloadSelector `json:"workloadSelector"`
}

type EnvoyFilterSpecConfigPatches struct {
	ApplyTo EnvoyFilterSpecConfigPatchesApplyTo `json:"applyTo"`
	// Match on listener/route configuration/cluster.
	Match *EnvoyFilterSpecConfigPatchesMatch `json:"match"`
	// The patch to apply along with the operation.
	Patch *EnvoyFilterSpecConfigPatchesPatch `json:"patch"`
}

type EnvoyFilterSpecConfigPatchesApplyTo string

const (
	EnvoyFilterSpecConfigPatchesApplyTo_INVALID             EnvoyFilterSpecConfigPatchesApplyTo = "INVALID"
	EnvoyFilterSpecConfigPatchesApplyTo_LISTENER            EnvoyFilterSpecConfigPatchesApplyTo = "LISTENER"
	EnvoyFilterSpecConfigPatchesApplyTo_FILTER_CHAIN        EnvoyFilterSpecConfigPatchesApplyTo = "FILTER_CHAIN"
	EnvoyFilterSpecConfigPatchesApplyTo_NETWORK_FILTER      EnvoyFilterSpecConfigPatchesApplyTo = "NETWORK_FILTER"
	EnvoyFilterSpecConfigPatchesApplyTo_HTTP_FILTER         EnvoyFilterSpecConfigPatchesApplyTo = "HTTP_FILTER"
	EnvoyFilterSpecConfigPatchesApplyTo_ROUTE_CONFIGURATION EnvoyFilterSpecConfigPatchesApplyTo = "ROUTE_CONFIGURATION"
	EnvoyFilterSpecConfigPatchesApplyTo_VIRTUAL_HOST        EnvoyFilterSpecConfigPatchesApplyTo = "VIRTUAL_HOST"
	EnvoyFilterSpecConfigPatchesApplyTo_HTTP_ROUTE          EnvoyFilterSpecConfigPatchesApplyTo = "HTTP_ROUTE"
	EnvoyFilterSpecConfigPatchesApplyTo_CLUSTER             EnvoyFilterSpecConfigPatchesApplyTo = "CLUSTER"
	EnvoyFilterSpecConfigPatchesApplyTo_EXTENSION_CONFIG    EnvoyFilterSpecConfigPatchesApplyTo = "EXTENSION_CONFIG"
)

// Match on listener/route configuration/cluster.
type EnvoyFilterSpecConfigPatchesMatch struct {
	// Match on envoy cluster attributes.
	Cluster *EnvoyFilterSpecConfigPatchesMatchCluster `json:"cluster"`
	// The specific config generation context to match on.
	Context EnvoyFilterSpecConfigPatchesMatchContext `json:"context"`
	// Match on envoy listener attributes.
	Listener *EnvoyFilterSpecConfigPatchesMatchListener `json:"listener"`
	// Match on properties associated with a proxy.
	Proxy *EnvoyFilterSpecConfigPatchesMatchProxy `json:"proxy"`
	// Match on envoy HTTP route configuration attributes.
	RouteConfiguration *EnvoyFilterSpecConfigPatchesMatchRouteConfiguration `json:"routeConfiguration"`
}

// Match on envoy cluster attributes.
type EnvoyFilterSpecConfigPatchesMatchCluster struct {
	// The exact name of the cluster to match.
	Name *string `json:"name"`
	// The service port for which this cluster was generated.
	PortNumber *float64 `json:"portNumber"`
	// The fully qualified service name for this cluster.
	Service *string `json:"service"`
	// The subset associated with the service.
	Subset *string `json:"subset"`
}

// The specific config generation context to match on.
type EnvoyFilterSpecConfigPatchesMatchContext string

const (
	EnvoyFilterSpecConfigPatchesMatchContext_ANY              EnvoyFilterSpecConfigPatchesMatchContext = "ANY"
	EnvoyFilterSpecConfigPatchesMatchContext_SIDECAR_INBOUND  EnvoyFilterSpecConfigPatchesMatchContext = "SIDECAR_INBOUND"
	EnvoyFilterSpecConfigPatchesMatchContext_SIDECAR_OUTBOUND EnvoyFilterSpecConfigPatchesMatchContext = "SIDECAR_OUTBOUND"
	EnvoyFilterSpecConfigPatchesMatchContext_GATEWAY          EnvoyFilterSpecConfigPatchesMatchContext = "GATEWAY"
)

// Match on envoy listener attributes.
type EnvoyFilterSpecConfigPatchesMatchListener struct {
	// Match a specific filter chain in a listener.
	FilterChain *EnvoyFilterSpecConfigPatchesMatchListenerFilterChain `json:"filterChain"`
	// Match a specific listener by its name.
	Name       *string  `json:"name"`
	PortName   *string  `json:"portName"`
	PortNumber *float64 `json:"portNumber"`
}

// Match a specific filter chain in a listener.
type EnvoyFilterSpecConfigPatchesMatchListenerFilterChain struct {
	// Applies only to sidecars.
	ApplicationProtocols *string `json:"applicationProtocols"`
	// The destination_port value used by a filter chain's match condition.
	DestinationPort *float64 `json:"destinationPort"`
	// The name of a specific filter to apply the patch to.
	Filter *EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilter `json:"filter"`
	// The name assigned to the filter chain.
	Name *string `json:"name"`
	// The SNI value used by a filter chain's match condition.
	Sni *string `json:"sni"`
	// Applies only to `SIDECAR_INBOUND` context.
	TransportProtocol *string `json:"transportProtocol"`
}

// The name of a specific filter to apply the patch to.
type EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilter struct {
	// The filter name to match on.
	Name      *string                                                              `json:"name"`
	SubFilter *EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilterSubFilter `json:"subFilter"`
}

type EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilterSubFilter struct {
	// The filter name to match on.
	Name *string `json:"name"`
}

// Match on properties associated with a proxy.
type EnvoyFilterSpecConfigPatchesMatchProxy struct {
	Metadata     *map[string]*string `json:"metadata"`
	ProxyVersion *string             `json:"proxyVersion"`
}

// Match on envoy HTTP route configuration attributes.
type EnvoyFilterSpecConfigPatchesMatchRouteConfiguration struct {
	Gateway *string `json:"gateway"`
	// Route configuration name to match on.
	Name *string `json:"name"`
	// Applicable only for GATEWAY context.
	PortName   *string                                                   `json:"portName"`
	PortNumber *float64                                                  `json:"portNumber"`
	Vhost      *EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhost `json:"vhost"`
}

type EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhost struct {
	Name *string `json:"name"`
	// Match a specific route within the virtual host.
	Route *EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRoute `json:"route"`
}

// Match a specific route within the virtual host.
type EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRoute struct {
	// Match a route with specific action type.
	Action EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction `json:"action"`
	Name   *string                                                             `json:"name"`
}

// Match a route with specific action type.
type EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction string

const (
	EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_ANY             EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction = "ANY"
	EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_ROUTE           EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction = "ROUTE"
	EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_REDIRECT        EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction = "REDIRECT"
	EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_DIRECT_RESPONSE EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction = "DIRECT_RESPONSE"
)

// The patch to apply along with the operation.
type EnvoyFilterSpecConfigPatchesPatch struct {
	// Determines the filter insertion order.
	FilterClass EnvoyFilterSpecConfigPatchesPatchFilterClass `json:"filterClass"`
	// Determines how the patch should be applied.
	Operation EnvoyFilterSpecConfigPatchesPatchOperation `json:"operation"`
	// The JSON config of the object being patched.
	Value interface{} `json:"value"`
}

// Determines the filter insertion order.
type EnvoyFilterSpecConfigPatchesPatchFilterClass string

const (
	EnvoyFilterSpecConfigPatchesPatchFilterClass_UNSPECIFIED EnvoyFilterSpecConfigPatchesPatchFilterClass = "UNSPECIFIED"
	EnvoyFilterSpecConfigPatchesPatchFilterClass_AUTHN       EnvoyFilterSpecConfigPatchesPatchFilterClass = "AUTHN"
	EnvoyFilterSpecConfigPatchesPatchFilterClass_AUTHZ       EnvoyFilterSpecConfigPatchesPatchFilterClass = "AUTHZ"
	EnvoyFilterSpecConfigPatchesPatchFilterClass_STATS       EnvoyFilterSpecConfigPatchesPatchFilterClass = "STATS"
)

// Determines how the patch should be applied.
type EnvoyFilterSpecConfigPatchesPatchOperation string

const (
	EnvoyFilterSpecConfigPatchesPatchOperation_INVALID       EnvoyFilterSpecConfigPatchesPatchOperation = "INVALID"
	EnvoyFilterSpecConfigPatchesPatchOperation_MERGE         EnvoyFilterSpecConfigPatchesPatchOperation = "MERGE"
	EnvoyFilterSpecConfigPatchesPatchOperation_ADD           EnvoyFilterSpecConfigPatchesPatchOperation = "ADD"
	EnvoyFilterSpecConfigPatchesPatchOperation_REMOVE        EnvoyFilterSpecConfigPatchesPatchOperation = "REMOVE"
	EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_BEFORE EnvoyFilterSpecConfigPatchesPatchOperation = "INSERT_BEFORE"
	EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_AFTER  EnvoyFilterSpecConfigPatchesPatchOperation = "INSERT_AFTER"
	EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_FIRST  EnvoyFilterSpecConfigPatchesPatchOperation = "INSERT_FIRST"
	EnvoyFilterSpecConfigPatchesPatchOperation_REPLACE       EnvoyFilterSpecConfigPatchesPatchOperation = "REPLACE"
)

type EnvoyFilterSpecWorkloadSelector struct {
	Labels *map[string]*string `json:"labels"`
}

type Gateway interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for Gateway
type jsiiProxy_Gateway struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Gateway) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gateway) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gateway) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gateway) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gateway) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "Gateway" API object.
func NewGateway(scope constructs.Construct, id *string, props *GatewayProps) Gateway {
	_init_.Initialize()

	j := jsiiProxy_Gateway{}

	_jsii_.Create(
		"networkingistioio.Gateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Gateway" API object.
func NewGateway_Override(g Gateway, scope constructs.Construct, id *string, props *GatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"networkingistioio.Gateway",
		[]interface{}{scope, id, props},
		g,
	)
}

// Renders a Kubernetes manifest for "Gateway".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Gateway_Manifest(props *GatewayProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"networkingistioio.Gateway",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func Gateway_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"networkingistioio.Gateway",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Gateway_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"networkingistioio.Gateway",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (g *jsiiProxy_Gateway) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (g *jsiiProxy_Gateway) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (g *jsiiProxy_Gateway) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (g *jsiiProxy_Gateway) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (g *jsiiProxy_Gateway) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (g *jsiiProxy_Gateway) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (g *jsiiProxy_Gateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GatewayProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Configuration affecting edge load balancer.
	//
	// See more details at: https://istio.io/docs/reference/config/networking/gateway.html
	Spec *GatewaySpec `json:"spec"`
}

// Configuration affecting edge load balancer.
//
// See more details at: https://istio.io/docs/reference/config/networking/gateway.html
type GatewaySpec struct {
	Selector *map[string]*string `json:"selector"`
	// A list of server specifications.
	Servers *[]*GatewaySpecServers `json:"servers"`
}

type GatewaySpecServers struct {
	Bind            *string `json:"bind"`
	DefaultEndpoint *string `json:"defaultEndpoint"`
	// One or more hosts exposed by this gateway.
	Hosts *[]*string `json:"hosts"`
	// An optional name of the server, when set must be unique across all servers.
	Name *string                 `json:"name"`
	Port *GatewaySpecServersPort `json:"port"`
	// Set of TLS related options that govern the server's behavior.
	Tls *GatewaySpecServersTls `json:"tls"`
}

type GatewaySpecServersPort struct {
	// Label assigned to the port.
	Name *string `json:"name"`
	// A valid non-negative integer port number.
	Number *float64 `json:"number"`
	// The protocol exposed on the port.
	Protocol   *string  `json:"protocol"`
	TargetPort *float64 `json:"targetPort"`
}

// Set of TLS related options that govern the server's behavior.
type GatewaySpecServersTls struct {
	// REQUIRED if mode is `MUTUAL`.
	CaCertificates *string `json:"caCertificates"`
	// Optional: If specified, only support the specified cipher list.
	CipherSuites   *[]*string `json:"cipherSuites"`
	CredentialName *string    `json:"credentialName"`
	HttpsRedirect  *bool      `json:"httpsRedirect"`
	// Optional: Maximum TLS protocol version.
	MaxProtocolVersion GatewaySpecServersTlsMaxProtocolVersion `json:"maxProtocolVersion"`
	// Optional: Minimum TLS protocol version.
	MinProtocolVersion GatewaySpecServersTlsMinProtocolVersion `json:"minProtocolVersion"`
	Mode               GatewaySpecServersTlsMode               `json:"mode"`
	// REQUIRED if mode is `SIMPLE` or `MUTUAL`.
	PrivateKey *string `json:"privateKey"`
	// REQUIRED if mode is `SIMPLE` or `MUTUAL`.
	ServerCertificate     *string    `json:"serverCertificate"`
	SubjectAltNames       *[]*string `json:"subjectAltNames"`
	VerifyCertificateHash *[]*string `json:"verifyCertificateHash"`
	VerifyCertificateSpki *[]*string `json:"verifyCertificateSpki"`
}

// Optional: Maximum TLS protocol version.
type GatewaySpecServersTlsMaxProtocolVersion string

const (
	GatewaySpecServersTlsMaxProtocolVersion_TLS_AUTO GatewaySpecServersTlsMaxProtocolVersion = "TLS_AUTO"
	GatewaySpecServersTlsMaxProtocolVersion_TLSV1_0  GatewaySpecServersTlsMaxProtocolVersion = "TLSV1_0"
	GatewaySpecServersTlsMaxProtocolVersion_TLSV1_1  GatewaySpecServersTlsMaxProtocolVersion = "TLSV1_1"
	GatewaySpecServersTlsMaxProtocolVersion_TLSV1_2  GatewaySpecServersTlsMaxProtocolVersion = "TLSV1_2"
	GatewaySpecServersTlsMaxProtocolVersion_TLSV1_3  GatewaySpecServersTlsMaxProtocolVersion = "TLSV1_3"
)

// Optional: Minimum TLS protocol version.
type GatewaySpecServersTlsMinProtocolVersion string

const (
	GatewaySpecServersTlsMinProtocolVersion_TLS_AUTO GatewaySpecServersTlsMinProtocolVersion = "TLS_AUTO"
	GatewaySpecServersTlsMinProtocolVersion_TLSV1_0  GatewaySpecServersTlsMinProtocolVersion = "TLSV1_0"
	GatewaySpecServersTlsMinProtocolVersion_TLSV1_1  GatewaySpecServersTlsMinProtocolVersion = "TLSV1_1"
	GatewaySpecServersTlsMinProtocolVersion_TLSV1_2  GatewaySpecServersTlsMinProtocolVersion = "TLSV1_2"
	GatewaySpecServersTlsMinProtocolVersion_TLSV1_3  GatewaySpecServersTlsMinProtocolVersion = "TLSV1_3"
)

type GatewaySpecServersTlsMode string

const (
	GatewaySpecServersTlsMode_PASSTHROUGH      GatewaySpecServersTlsMode = "PASSTHROUGH"
	GatewaySpecServersTlsMode_SIMPLE           GatewaySpecServersTlsMode = "SIMPLE"
	GatewaySpecServersTlsMode_MUTUAL           GatewaySpecServersTlsMode = "MUTUAL"
	GatewaySpecServersTlsMode_AUTO_PASSTHROUGH GatewaySpecServersTlsMode = "AUTO_PASSTHROUGH"
	GatewaySpecServersTlsMode_ISTIO_MUTUAL     GatewaySpecServersTlsMode = "ISTIO_MUTUAL"
)

type ServiceEntry interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ServiceEntry
type jsiiProxy_ServiceEntry struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ServiceEntry) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEntry) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEntry) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEntry) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEntry) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEntry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "ServiceEntry" API object.
func NewServiceEntry(scope constructs.Construct, id *string, props *ServiceEntryProps) ServiceEntry {
	_init_.Initialize()

	j := jsiiProxy_ServiceEntry{}

	_jsii_.Create(
		"networkingistioio.ServiceEntry",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ServiceEntry" API object.
func NewServiceEntry_Override(s ServiceEntry, scope constructs.Construct, id *string, props *ServiceEntryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"networkingistioio.ServiceEntry",
		[]interface{}{scope, id, props},
		s,
	)
}

// Renders a Kubernetes manifest for "ServiceEntry".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ServiceEntry_Manifest(props *ServiceEntryProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"networkingistioio.ServiceEntry",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ServiceEntry_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"networkingistioio.ServiceEntry",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ServiceEntry_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"networkingistioio.ServiceEntry",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (s *jsiiProxy_ServiceEntry) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (s *jsiiProxy_ServiceEntry) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (s *jsiiProxy_ServiceEntry) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (s *jsiiProxy_ServiceEntry) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (s *jsiiProxy_ServiceEntry) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (s *jsiiProxy_ServiceEntry) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServiceEntry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceEntryProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Configuration affecting service registry.
	//
	// See more details at: https://istio.io/docs/reference/config/networking/service-entry.html
	Spec *ServiceEntrySpec `json:"spec"`
}

// Configuration affecting service registry.
//
// See more details at: https://istio.io/docs/reference/config/networking/service-entry.html
type ServiceEntrySpec struct {
	// The virtual IP addresses associated with the service.
	Addresses *[]*string `json:"addresses"`
	// One or more endpoints associated with the service.
	Endpoints *[]*ServiceEntrySpecEndpoints `json:"endpoints"`
	// A list of namespaces to which this service is exported.
	ExportTo *[]*string `json:"exportTo"`
	// The hosts associated with the ServiceEntry.
	Hosts    *[]*string               `json:"hosts"`
	Location ServiceEntrySpecLocation `json:"location"`
	// The ports associated with the external service.
	Ports *[]*ServiceEntrySpecPorts `json:"ports"`
	// Service discovery mode for the hosts.
	Resolution      ServiceEntrySpecResolution `json:"resolution"`
	SubjectAltNames *[]*string                 `json:"subjectAltNames"`
	// Applicable only for MESH_INTERNAL services.
	WorkloadSelector *ServiceEntrySpecWorkloadSelector `json:"workloadSelector"`
}

type ServiceEntrySpecEndpoints struct {
	Address *string `json:"address"`
	// One or more labels associated with the endpoint.
	Labels *map[string]*string `json:"labels"`
	// The locality associated with the endpoint.
	Locality *string `json:"locality"`
	Network  *string `json:"network"`
	// Set of ports associated with the endpoint.
	Ports          *map[string]*float64 `json:"ports"`
	ServiceAccount *string              `json:"serviceAccount"`
	// The load balancing weight associated with the endpoint.
	Weight *float64 `json:"weight"`
}

type ServiceEntrySpecLocation string

const (
	ServiceEntrySpecLocation_MESH_EXTERNAL ServiceEntrySpecLocation = "MESH_EXTERNAL"
	ServiceEntrySpecLocation_MESH_INTERNAL ServiceEntrySpecLocation = "MESH_INTERNAL"
)

type ServiceEntrySpecPorts struct {
	// Label assigned to the port.
	Name *string `json:"name"`
	// A valid non-negative integer port number.
	Number *float64 `json:"number"`
	// The protocol exposed on the port.
	Protocol   *string  `json:"protocol"`
	TargetPort *float64 `json:"targetPort"`
}

// Service discovery mode for the hosts.
type ServiceEntrySpecResolution string

const (
	ServiceEntrySpecResolution_NONE   ServiceEntrySpecResolution = "NONE"
	ServiceEntrySpecResolution_STATIC ServiceEntrySpecResolution = "STATIC"
	ServiceEntrySpecResolution_DNS    ServiceEntrySpecResolution = "DNS"
)

// Applicable only for MESH_INTERNAL services.
type ServiceEntrySpecWorkloadSelector struct {
	Labels *map[string]*string `json:"labels"`
}

type Sidecar interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for Sidecar
type jsiiProxy_Sidecar struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Sidecar) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sidecar) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sidecar) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sidecar) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sidecar) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sidecar) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "Sidecar" API object.
func NewSidecar(scope constructs.Construct, id *string, props *SidecarProps) Sidecar {
	_init_.Initialize()

	j := jsiiProxy_Sidecar{}

	_jsii_.Create(
		"networkingistioio.Sidecar",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Sidecar" API object.
func NewSidecar_Override(s Sidecar, scope constructs.Construct, id *string, props *SidecarProps) {
	_init_.Initialize()

	_jsii_.Create(
		"networkingistioio.Sidecar",
		[]interface{}{scope, id, props},
		s,
	)
}

// Renders a Kubernetes manifest for "Sidecar".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Sidecar_Manifest(props *SidecarProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"networkingistioio.Sidecar",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func Sidecar_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"networkingistioio.Sidecar",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Sidecar_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"networkingistioio.Sidecar",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (s *jsiiProxy_Sidecar) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (s *jsiiProxy_Sidecar) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (s *jsiiProxy_Sidecar) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (s *jsiiProxy_Sidecar) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (s *jsiiProxy_Sidecar) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (s *jsiiProxy_Sidecar) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_Sidecar) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SidecarProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Configuration affecting network reachability of a sidecar.
	//
	// See more details at: https://istio.io/docs/reference/config/networking/sidecar.html
	Spec *SidecarSpec `json:"spec"`
}

// Configuration affecting network reachability of a sidecar.
//
// See more details at: https://istio.io/docs/reference/config/networking/sidecar.html
type SidecarSpec struct {
	Egress  *[]*SidecarSpecEgress  `json:"egress"`
	Ingress *[]*SidecarSpecIngress `json:"ingress"`
	// Configuration for the outbound traffic policy.
	OutboundTrafficPolicy *SidecarSpecOutboundTrafficPolicy `json:"outboundTrafficPolicy"`
	WorkloadSelector      *SidecarSpecWorkloadSelector      `json:"workloadSelector"`
}

type SidecarSpecEgress struct {
	Bind        *string                      `json:"bind"`
	CaptureMode SidecarSpecEgressCaptureMode `json:"captureMode"`
	Hosts       *[]*string                   `json:"hosts"`
	// The port associated with the listener.
	Port *SidecarSpecEgressPort `json:"port"`
}

type SidecarSpecEgressCaptureMode string

const (
	SidecarSpecEgressCaptureMode_DEFAULT  SidecarSpecEgressCaptureMode = "DEFAULT"
	SidecarSpecEgressCaptureMode_IPTABLES SidecarSpecEgressCaptureMode = "IPTABLES"
	SidecarSpecEgressCaptureMode_NONE     SidecarSpecEgressCaptureMode = "NONE"
)

// The port associated with the listener.
type SidecarSpecEgressPort struct {
	// Label assigned to the port.
	Name *string `json:"name"`
	// A valid non-negative integer port number.
	Number *float64 `json:"number"`
	// The protocol exposed on the port.
	Protocol   *string  `json:"protocol"`
	TargetPort *float64 `json:"targetPort"`
}

type SidecarSpecIngress struct {
	// The IP to which the listener should be bound.
	Bind            *string                       `json:"bind"`
	CaptureMode     SidecarSpecIngressCaptureMode `json:"captureMode"`
	DefaultEndpoint *string                       `json:"defaultEndpoint"`
	// The port associated with the listener.
	Port *SidecarSpecIngressPort `json:"port"`
}

type SidecarSpecIngressCaptureMode string

const (
	SidecarSpecIngressCaptureMode_DEFAULT  SidecarSpecIngressCaptureMode = "DEFAULT"
	SidecarSpecIngressCaptureMode_IPTABLES SidecarSpecIngressCaptureMode = "IPTABLES"
	SidecarSpecIngressCaptureMode_NONE     SidecarSpecIngressCaptureMode = "NONE"
)

// The port associated with the listener.
type SidecarSpecIngressPort struct {
	// Label assigned to the port.
	Name *string `json:"name"`
	// A valid non-negative integer port number.
	Number *float64 `json:"number"`
	// The protocol exposed on the port.
	Protocol   *string  `json:"protocol"`
	TargetPort *float64 `json:"targetPort"`
}

// Configuration for the outbound traffic policy.
type SidecarSpecOutboundTrafficPolicy struct {
	EgressProxy *SidecarSpecOutboundTrafficPolicyEgressProxy `json:"egressProxy"`
	Mode        SidecarSpecOutboundTrafficPolicyMode         `json:"mode"`
}

type SidecarSpecOutboundTrafficPolicyEgressProxy struct {
	// The name of a service from the service registry.
	Host *string `json:"host"`
	// Specifies the port on the host that is being addressed.
	Port *SidecarSpecOutboundTrafficPolicyEgressProxyPort `json:"port"`
	// The name of a subset within the service.
	Subset *string `json:"subset"`
}

// Specifies the port on the host that is being addressed.
type SidecarSpecOutboundTrafficPolicyEgressProxyPort struct {
	Number *float64 `json:"number"`
}

type SidecarSpecOutboundTrafficPolicyMode string

const (
	SidecarSpecOutboundTrafficPolicyMode_REGISTRY_ONLY SidecarSpecOutboundTrafficPolicyMode = "REGISTRY_ONLY"
	SidecarSpecOutboundTrafficPolicyMode_ALLOW_ANY     SidecarSpecOutboundTrafficPolicyMode = "ALLOW_ANY"
)

type SidecarSpecWorkloadSelector struct {
	Labels *map[string]*string `json:"labels"`
}

type VirtualService interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for VirtualService
type jsiiProxy_VirtualService struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_VirtualService) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "VirtualService" API object.
func NewVirtualService(scope constructs.Construct, id *string, props *VirtualServiceProps) VirtualService {
	_init_.Initialize()

	j := jsiiProxy_VirtualService{}

	_jsii_.Create(
		"networkingistioio.VirtualService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "VirtualService" API object.
func NewVirtualService_Override(v VirtualService, scope constructs.Construct, id *string, props *VirtualServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"networkingistioio.VirtualService",
		[]interface{}{scope, id, props},
		v,
	)
}

// Renders a Kubernetes manifest for "VirtualService".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func VirtualService_Manifest(props *VirtualServiceProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"networkingistioio.VirtualService",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func VirtualService_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"networkingistioio.VirtualService",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func VirtualService_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"networkingistioio.VirtualService",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (v *jsiiProxy_VirtualService) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (v *jsiiProxy_VirtualService) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (v *jsiiProxy_VirtualService) OnPrepare() {
	_jsii_.InvokeVoid(
		v,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (v *jsiiProxy_VirtualService) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		v,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (v *jsiiProxy_VirtualService) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (v *jsiiProxy_VirtualService) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (v *jsiiProxy_VirtualService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualServiceProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Configuration affecting label/content routing, sni routing, etc.
	//
	// See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html
	Spec *VirtualServiceSpec `json:"spec"`
}

// Configuration affecting label/content routing, sni routing, etc.
//
// See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html
type VirtualServiceSpec struct {
	// A list of namespaces to which this virtual service is exported.
	ExportTo *[]*string `json:"exportTo"`
	// The names of gateways and sidecars that should apply these routes.
	Gateways *[]*string `json:"gateways"`
	// The destination hosts to which traffic is being sent.
	Hosts *[]*string `json:"hosts"`
	// An ordered list of route rules for HTTP traffic.
	Http *[]*VirtualServiceSpecHttp `json:"http"`
	// An ordered list of route rules for opaque TCP traffic.
	Tcp *[]*VirtualServiceSpecTcp `json:"tcp"`
	Tls *[]*VirtualServiceSpecTls `json:"tls"`
}

type VirtualServiceSpecHttp struct {
	// Cross-Origin Resource Sharing policy (CORS).
	CorsPolicy *VirtualServiceSpecHttpCorsPolicy `json:"corsPolicy"`
	Delegate   *VirtualServiceSpecHttpDelegate   `json:"delegate"`
	// Fault injection policy to apply on HTTP traffic at the client side.
	Fault   *VirtualServiceSpecHttpFault    `json:"fault"`
	Headers *VirtualServiceSpecHttpHeaders  `json:"headers"`
	Match   *[]*VirtualServiceSpecHttpMatch `json:"match"`
	Mirror  *VirtualServiceSpecHttpMirror   `json:"mirror"`
	// Percentage of the traffic to be mirrored by the `mirror` field.
	MirrorPercent *float64 `json:"mirrorPercent"`
	// Percentage of the traffic to be mirrored by the `mirror` field.
	MirrorPercentage *VirtualServiceSpecHttpMirrorPercentage `json:"mirrorPercentage"`
	// The name assigned to the route for debugging purposes.
	Name *string `json:"name"`
	// A HTTP rule can either redirect or forward (default) traffic.
	Redirect *VirtualServiceSpecHttpRedirect `json:"redirect"`
	// Retry policy for HTTP requests.
	Retries *VirtualServiceSpecHttpRetries `json:"retries"`
	// Rewrite HTTP URIs and Authority headers.
	Rewrite *VirtualServiceSpecHttpRewrite `json:"rewrite"`
	// A HTTP rule can either redirect or forward (default) traffic.
	Route *[]*VirtualServiceSpecHttpRoute `json:"route"`
	// Timeout for HTTP requests, default is disabled.
	Timeout *string `json:"timeout"`
}

// Cross-Origin Resource Sharing policy (CORS).
type VirtualServiceSpecHttpCorsPolicy struct {
	AllowCredentials *bool      `json:"allowCredentials"`
	AllowHeaders     *[]*string `json:"allowHeaders"`
	// List of HTTP methods allowed to access the resource.
	AllowMethods *[]*string `json:"allowMethods"`
	// The list of origins that are allowed to perform CORS requests.
	AllowOrigin *[]*string `json:"allowOrigin"`
	// String patterns that match allowed origins.
	AllowOrigins  *[]*VirtualServiceSpecHttpCorsPolicyAllowOrigins `json:"allowOrigins"`
	ExposeHeaders *[]*string                                       `json:"exposeHeaders"`
	MaxAge        *string                                          `json:"maxAge"`
}

type VirtualServiceSpecHttpCorsPolicyAllowOrigins struct {
	Exact  *string `json:"exact"`
	Prefix *string `json:"prefix"`
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex *string `json:"regex"`
}

type VirtualServiceSpecHttpDelegate struct {
	// Name specifies the name of the delegate VirtualService.
	Name *string `json:"name"`
	// Namespace specifies the namespace where the delegate VirtualService resides.
	Namespace *string `json:"namespace"`
}

// Fault injection policy to apply on HTTP traffic at the client side.
type VirtualServiceSpecHttpFault struct {
	Abort *VirtualServiceSpecHttpFaultAbort `json:"abort"`
	Delay *VirtualServiceSpecHttpFaultDelay `json:"delay"`
}

type VirtualServiceSpecHttpFaultAbort struct {
	GrpcStatus *string `json:"grpcStatus"`
	Http2Error *string `json:"http2Error"`
	// HTTP status code to use to abort the Http request.
	HttpStatus *float64 `json:"httpStatus"`
	// Percentage of requests to be aborted with the error code provided.
	Percentage *VirtualServiceSpecHttpFaultAbortPercentage `json:"percentage"`
}

// Percentage of requests to be aborted with the error code provided.
type VirtualServiceSpecHttpFaultAbortPercentage struct {
	Value *float64 `json:"value"`
}

type VirtualServiceSpecHttpFaultDelay struct {
	ExponentialDelay *string `json:"exponentialDelay"`
	// Add a fixed delay before forwarding the request.
	FixedDelay *string `json:"fixedDelay"`
	// Percentage of requests on which the delay will be injected (0-100).
	Percent *float64 `json:"percent"`
	// Percentage of requests on which the delay will be injected.
	Percentage *VirtualServiceSpecHttpFaultDelayPercentage `json:"percentage"`
}

// Percentage of requests on which the delay will be injected.
type VirtualServiceSpecHttpFaultDelayPercentage struct {
	Value *float64 `json:"value"`
}

type VirtualServiceSpecHttpHeaders struct {
	Request  *VirtualServiceSpecHttpHeadersRequest  `json:"request"`
	Response *VirtualServiceSpecHttpHeadersResponse `json:"response"`
}

type VirtualServiceSpecHttpHeadersRequest struct {
	Add    *map[string]*string `json:"add"`
	Remove *[]*string          `json:"remove"`
	Set    *map[string]*string `json:"set"`
}

type VirtualServiceSpecHttpHeadersResponse struct {
	Add    *map[string]*string `json:"add"`
	Remove *[]*string          `json:"remove"`
	Set    *map[string]*string `json:"set"`
}

type VirtualServiceSpecHttpMatch struct {
	Authority *VirtualServiceSpecHttpMatchAuthority `json:"authority"`
	// Names of gateways where the rule should be applied.
	Gateways *[]*string                                      `json:"gateways"`
	Headers  *map[string]*VirtualServiceSpecHttpMatchHeaders `json:"headers"`
	// Flag to specify whether the URI matching should be case-insensitive.
	IgnoreUriCase *bool                              `json:"ignoreUriCase"`
	Method        *VirtualServiceSpecHttpMatchMethod `json:"method"`
	// The name assigned to a match.
	Name *string `json:"name"`
	// Specifies the ports on the host that is being addressed.
	Port *float64 `json:"port"`
	// Query parameters for matching.
	QueryParams  *map[string]*VirtualServiceSpecHttpMatchQueryParams `json:"queryParams"`
	Scheme       *VirtualServiceSpecHttpMatchScheme                  `json:"scheme"`
	SourceLabels *map[string]*string                                 `json:"sourceLabels"`
	// Source namespace constraining the applicability of a rule to workloads in that namespace.
	SourceNamespace *string                         `json:"sourceNamespace"`
	Uri             *VirtualServiceSpecHttpMatchUri `json:"uri"`
	// withoutHeader has the same syntax with the header, but has opposite meaning.
	WithoutHeaders *map[string]*VirtualServiceSpecHttpMatchWithoutHeaders `json:"withoutHeaders"`
}

type VirtualServiceSpecHttpMatchAuthority struct {
	Exact  *string `json:"exact"`
	Prefix *string `json:"prefix"`
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex *string `json:"regex"`
}

type VirtualServiceSpecHttpMatchHeaders struct {
	Exact  *string `json:"exact"`
	Prefix *string `json:"prefix"`
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex *string `json:"regex"`
}

type VirtualServiceSpecHttpMatchMethod struct {
	Exact  *string `json:"exact"`
	Prefix *string `json:"prefix"`
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex *string `json:"regex"`
}

type VirtualServiceSpecHttpMatchQueryParams struct {
	Exact  *string `json:"exact"`
	Prefix *string `json:"prefix"`
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex *string `json:"regex"`
}

type VirtualServiceSpecHttpMatchScheme struct {
	Exact  *string `json:"exact"`
	Prefix *string `json:"prefix"`
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex *string `json:"regex"`
}

type VirtualServiceSpecHttpMatchUri struct {
	Exact  *string `json:"exact"`
	Prefix *string `json:"prefix"`
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex *string `json:"regex"`
}

type VirtualServiceSpecHttpMatchWithoutHeaders struct {
	Exact  *string `json:"exact"`
	Prefix *string `json:"prefix"`
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex *string `json:"regex"`
}

type VirtualServiceSpecHttpMirror struct {
	// The name of a service from the service registry.
	Host *string `json:"host"`
	// Specifies the port on the host that is being addressed.
	Port *VirtualServiceSpecHttpMirrorPort `json:"port"`
	// The name of a subset within the service.
	Subset *string `json:"subset"`
}

// Percentage of the traffic to be mirrored by the `mirror` field.
type VirtualServiceSpecHttpMirrorPercentage struct {
	Value *float64 `json:"value"`
}

// Specifies the port on the host that is being addressed.
type VirtualServiceSpecHttpMirrorPort struct {
	Number *float64 `json:"number"`
}

// A HTTP rule can either redirect or forward (default) traffic.
type VirtualServiceSpecHttpRedirect struct {
	Authority    *string  `json:"authority"`
	RedirectCode *float64 `json:"redirectCode"`
	Uri          *string  `json:"uri"`
}

// Retry policy for HTTP requests.
type VirtualServiceSpecHttpRetries struct {
	// Number of retries to be allowed for a given request.
	Attempts *float64 `json:"attempts"`
	// Timeout per attempt for a given request, including the initial call and any retries.
	PerTryTimeout *string `json:"perTryTimeout"`
	// Specifies the conditions under which retry takes place.
	RetryOn *string `json:"retryOn"`
	// Flag to specify whether the retries should retry to other localities.
	RetryRemoteLocalities *bool `json:"retryRemoteLocalities"`
}

// Rewrite HTTP URIs and Authority headers.
type VirtualServiceSpecHttpRewrite struct {
	// rewrite the Authority/Host header with this value.
	Authority *string `json:"authority"`
	Uri       *string `json:"uri"`
}

type VirtualServiceSpecHttpRoute struct {
	Destination *VirtualServiceSpecHttpRouteDestination `json:"destination"`
	Headers     *VirtualServiceSpecHttpRouteHeaders     `json:"headers"`
	Weight      *float64                                `json:"weight"`
}

type VirtualServiceSpecHttpRouteDestination struct {
	// The name of a service from the service registry.
	Host *string `json:"host"`
	// Specifies the port on the host that is being addressed.
	Port *VirtualServiceSpecHttpRouteDestinationPort `json:"port"`
	// The name of a subset within the service.
	Subset *string `json:"subset"`
}

// Specifies the port on the host that is being addressed.
type VirtualServiceSpecHttpRouteDestinationPort struct {
	Number *float64 `json:"number"`
}

type VirtualServiceSpecHttpRouteHeaders struct {
	Request  *VirtualServiceSpecHttpRouteHeadersRequest  `json:"request"`
	Response *VirtualServiceSpecHttpRouteHeadersResponse `json:"response"`
}

type VirtualServiceSpecHttpRouteHeadersRequest struct {
	Add    *map[string]*string `json:"add"`
	Remove *[]*string          `json:"remove"`
	Set    *map[string]*string `json:"set"`
}

type VirtualServiceSpecHttpRouteHeadersResponse struct {
	Add    *map[string]*string `json:"add"`
	Remove *[]*string          `json:"remove"`
	Set    *map[string]*string `json:"set"`
}

type VirtualServiceSpecTcp struct {
	Match *[]*VirtualServiceSpecTcpMatch `json:"match"`
	// The destination to which the connection should be forwarded to.
	Route *[]*VirtualServiceSpecTcpRoute `json:"route"`
}

type VirtualServiceSpecTcpMatch struct {
	// IPv4 or IPv6 ip addresses of destination with optional subnet.
	DestinationSubnets *[]*string `json:"destinationSubnets"`
	// Names of gateways where the rule should be applied.
	Gateways *[]*string `json:"gateways"`
	// Specifies the port on the host that is being addressed.
	Port         *float64            `json:"port"`
	SourceLabels *map[string]*string `json:"sourceLabels"`
	// Source namespace constraining the applicability of a rule to workloads in that namespace.
	SourceNamespace *string `json:"sourceNamespace"`
	// IPv4 or IPv6 ip address of source with optional subnet.
	SourceSubnet *string `json:"sourceSubnet"`
}

type VirtualServiceSpecTcpRoute struct {
	Destination *VirtualServiceSpecTcpRouteDestination `json:"destination"`
	Weight      *float64                               `json:"weight"`
}

type VirtualServiceSpecTcpRouteDestination struct {
	// The name of a service from the service registry.
	Host *string `json:"host"`
	// Specifies the port on the host that is being addressed.
	Port *VirtualServiceSpecTcpRouteDestinationPort `json:"port"`
	// The name of a subset within the service.
	Subset *string `json:"subset"`
}

// Specifies the port on the host that is being addressed.
type VirtualServiceSpecTcpRouteDestinationPort struct {
	Number *float64 `json:"number"`
}

type VirtualServiceSpecTls struct {
	Match *[]*VirtualServiceSpecTlsMatch `json:"match"`
	// The destination to which the connection should be forwarded to.
	Route *[]*VirtualServiceSpecTlsRoute `json:"route"`
}

type VirtualServiceSpecTlsMatch struct {
	// IPv4 or IPv6 ip addresses of destination with optional subnet.
	DestinationSubnets *[]*string `json:"destinationSubnets"`
	// Names of gateways where the rule should be applied.
	Gateways *[]*string `json:"gateways"`
	// Specifies the port on the host that is being addressed.
	Port *float64 `json:"port"`
	// SNI (server name indicator) to match on.
	SniHosts     *[]*string          `json:"sniHosts"`
	SourceLabels *map[string]*string `json:"sourceLabels"`
	// Source namespace constraining the applicability of a rule to workloads in that namespace.
	SourceNamespace *string `json:"sourceNamespace"`
}

type VirtualServiceSpecTlsRoute struct {
	Destination *VirtualServiceSpecTlsRouteDestination `json:"destination"`
	Weight      *float64                               `json:"weight"`
}

type VirtualServiceSpecTlsRouteDestination struct {
	// The name of a service from the service registry.
	Host *string `json:"host"`
	// Specifies the port on the host that is being addressed.
	Port *VirtualServiceSpecTlsRouteDestinationPort `json:"port"`
	// The name of a subset within the service.
	Subset *string `json:"subset"`
}

// Specifies the port on the host that is being addressed.
type VirtualServiceSpecTlsRouteDestinationPort struct {
	Number *float64 `json:"number"`
}

type WorkloadEntry interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for WorkloadEntry
type jsiiProxy_WorkloadEntry struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_WorkloadEntry) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadEntry) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadEntry) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadEntry) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadEntry) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadEntry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "WorkloadEntry" API object.
func NewWorkloadEntry(scope constructs.Construct, id *string, props *WorkloadEntryProps) WorkloadEntry {
	_init_.Initialize()

	j := jsiiProxy_WorkloadEntry{}

	_jsii_.Create(
		"networkingistioio.WorkloadEntry",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "WorkloadEntry" API object.
func NewWorkloadEntry_Override(w WorkloadEntry, scope constructs.Construct, id *string, props *WorkloadEntryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"networkingistioio.WorkloadEntry",
		[]interface{}{scope, id, props},
		w,
	)
}

// Renders a Kubernetes manifest for "WorkloadEntry".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func WorkloadEntry_Manifest(props *WorkloadEntryProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"networkingistioio.WorkloadEntry",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func WorkloadEntry_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"networkingistioio.WorkloadEntry",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func WorkloadEntry_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"networkingistioio.WorkloadEntry",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (w *jsiiProxy_WorkloadEntry) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (w *jsiiProxy_WorkloadEntry) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (w *jsiiProxy_WorkloadEntry) OnPrepare() {
	_jsii_.InvokeVoid(
		w,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (w *jsiiProxy_WorkloadEntry) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		w,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (w *jsiiProxy_WorkloadEntry) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (w *jsiiProxy_WorkloadEntry) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WorkloadEntry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WorkloadEntryProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Configuration affecting VMs onboarded into the mesh.
	//
	// See more details at: https://istio.io/docs/reference/config/networking/workload-entry.html
	Spec *WorkloadEntrySpec `json:"spec"`
}

// Configuration affecting VMs onboarded into the mesh.
//
// See more details at: https://istio.io/docs/reference/config/networking/workload-entry.html
type WorkloadEntrySpec struct {
	Address *string `json:"address"`
	// One or more labels associated with the endpoint.
	Labels *map[string]*string `json:"labels"`
	// The locality associated with the endpoint.
	Locality *string `json:"locality"`
	Network  *string `json:"network"`
	// Set of ports associated with the endpoint.
	Ports          *map[string]*float64 `json:"ports"`
	ServiceAccount *string              `json:"serviceAccount"`
	// The load balancing weight associated with the endpoint.
	Weight *float64 `json:"weight"`
}

type WorkloadGroup interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for WorkloadGroup
type jsiiProxy_WorkloadGroup struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_WorkloadGroup) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadGroup) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadGroup) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadGroup) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadGroup) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "WorkloadGroup" API object.
func NewWorkloadGroup(scope constructs.Construct, id *string, props *WorkloadGroupProps) WorkloadGroup {
	_init_.Initialize()

	j := jsiiProxy_WorkloadGroup{}

	_jsii_.Create(
		"networkingistioio.WorkloadGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "WorkloadGroup" API object.
func NewWorkloadGroup_Override(w WorkloadGroup, scope constructs.Construct, id *string, props *WorkloadGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"networkingistioio.WorkloadGroup",
		[]interface{}{scope, id, props},
		w,
	)
}

// Renders a Kubernetes manifest for "WorkloadGroup".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func WorkloadGroup_Manifest(props *WorkloadGroupProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"networkingistioio.WorkloadGroup",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func WorkloadGroup_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"networkingistioio.WorkloadGroup",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func WorkloadGroup_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"networkingistioio.WorkloadGroup",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (w *jsiiProxy_WorkloadGroup) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (w *jsiiProxy_WorkloadGroup) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (w *jsiiProxy_WorkloadGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		w,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (w *jsiiProxy_WorkloadGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		w,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (w *jsiiProxy_WorkloadGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (w *jsiiProxy_WorkloadGroup) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WorkloadGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WorkloadGroupProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Describes a collection of workload instances.
	//
	// See more details at: https://istio.io/docs/reference/config/networking/workload-group.html
	Spec *WorkloadGroupSpec `json:"spec"`
}

// Describes a collection of workload instances.
//
// See more details at: https://istio.io/docs/reference/config/networking/workload-group.html
type WorkloadGroupSpec struct {
	// Metadata that will be used for all corresponding `WorkloadEntries`.
	Metadata *WorkloadGroupSpecMetadata `json:"metadata"`
	// `ReadinessProbe` describes the configuration the user must provide for healthchecking on their workload.
	Probe *WorkloadGroupSpecProbe `json:"probe"`
	// Template to be used for the generation of `WorkloadEntry` resources that belong to this `WorkloadGroup`.
	Template *WorkloadGroupSpecTemplate `json:"template"`
}

// Metadata that will be used for all corresponding `WorkloadEntries`.
type WorkloadGroupSpecMetadata struct {
	Annotations *map[string]*string `json:"annotations"`
	Labels      *map[string]*string `json:"labels"`
}

// `ReadinessProbe` describes the configuration the user must provide for healthchecking on their workload.
type WorkloadGroupSpecProbe struct {
	// Health is determined by how the command that is executed exited.
	Exec *WorkloadGroupSpecProbeExec `json:"exec"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	FailureThreshold *float64                       `json:"failureThreshold"`
	HttpGet          *WorkloadGroupSpecProbeHttpGet `json:"httpGet"`
	// Number of seconds after the container has started before readiness probes are initiated.
	InitialDelaySeconds *float64 `json:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	PeriodSeconds *float64 `json:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	SuccessThreshold *float64 `json:"successThreshold"`
	// Health is determined by if the proxy is able to connect.
	TcpSocket *WorkloadGroupSpecProbeTcpSocket `json:"tcpSocket"`
	// Number of seconds after which the probe times out.
	TimeoutSeconds *float64 `json:"timeoutSeconds"`
}

// Health is determined by how the command that is executed exited.
type WorkloadGroupSpecProbeExec struct {
	// Command to run.
	Command *[]*string `json:"command"`
}

type WorkloadGroupSpecProbeHttpGet struct {
	// Host name to connect to, defaults to the pod IP.
	Host *string `json:"host"`
	// Headers the proxy will pass on to make the request.
	HttpHeaders *[]*WorkloadGroupSpecProbeHttpGetHttpHeaders `json:"httpHeaders"`
	// Path to access on the HTTP server.
	Path *string `json:"path"`
	// Port on which the endpoint lives.
	Port   *float64 `json:"port"`
	Scheme *string  `json:"scheme"`
}

type WorkloadGroupSpecProbeHttpGetHttpHeaders struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

// Health is determined by if the proxy is able to connect.
type WorkloadGroupSpecProbeTcpSocket struct {
	Host *string  `json:"host"`
	Port *float64 `json:"port"`
}

// Template to be used for the generation of `WorkloadEntry` resources that belong to this `WorkloadGroup`.
type WorkloadGroupSpecTemplate struct {
	Address *string `json:"address"`
	// One or more labels associated with the endpoint.
	Labels *map[string]*string `json:"labels"`
	// The locality associated with the endpoint.
	Locality *string `json:"locality"`
	Network  *string `json:"network"`
	// Set of ports associated with the endpoint.
	Ports          *map[string]*float64 `json:"ports"`
	ServiceAccount *string              `json:"serviceAccount"`
	// The load balancing weight associated with the endpoint.
	Weight *float64 `json:"weight"`
}
