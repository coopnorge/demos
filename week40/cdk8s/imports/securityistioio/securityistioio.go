// securityistioio
package securityistioio

import (
	_init_ "example.com/cdk8s/imports/securityistioio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"example.com/cdk8s/imports/securityistioio/internal"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
)

type AuthorizationPolicy interface {
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

// The jsii proxy struct for AuthorizationPolicy
type jsiiProxy_AuthorizationPolicy struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_AuthorizationPolicy) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthorizationPolicy) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthorizationPolicy) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthorizationPolicy) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthorizationPolicy) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthorizationPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "AuthorizationPolicy" API object.
func NewAuthorizationPolicy(scope constructs.Construct, id *string, props *AuthorizationPolicyProps) AuthorizationPolicy {
	_init_.Initialize()

	j := jsiiProxy_AuthorizationPolicy{}

	_jsii_.Create(
		"securityistioio.AuthorizationPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "AuthorizationPolicy" API object.
func NewAuthorizationPolicy_Override(a AuthorizationPolicy, scope constructs.Construct, id *string, props *AuthorizationPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"securityistioio.AuthorizationPolicy",
		[]interface{}{scope, id, props},
		a,
	)
}

// Renders a Kubernetes manifest for "AuthorizationPolicy".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func AuthorizationPolicy_Manifest(props *AuthorizationPolicyProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"securityistioio.AuthorizationPolicy",
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
func AuthorizationPolicy_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"securityistioio.AuthorizationPolicy",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func AuthorizationPolicy_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"securityistioio.AuthorizationPolicy",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (a *jsiiProxy_AuthorizationPolicy) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (a *jsiiProxy_AuthorizationPolicy) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_AuthorizationPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (a *jsiiProxy_AuthorizationPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_AuthorizationPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (a *jsiiProxy_AuthorizationPolicy) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AuthorizationPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AuthorizationPolicyProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Configuration for access control on workloads.
	//
	// See more details at: https://istio.io/docs/reference/config/security/authorization-policy.html
	Spec *AuthorizationPolicySpec `json:"spec"`
}

// Configuration for access control on workloads.
//
// See more details at: https://istio.io/docs/reference/config/security/authorization-policy.html
type AuthorizationPolicySpec struct {
	// Optional.
	Action AuthorizationPolicySpecAction `json:"action"`
	// Specifies detailed configuration of the CUSTOM action.
	Provider *AuthorizationPolicySpecProvider `json:"provider"`
	// Optional.
	Rules *[]*AuthorizationPolicySpecRules `json:"rules"`
	// Optional.
	Selector *AuthorizationPolicySpecSelector `json:"selector"`
}

// Optional.
type AuthorizationPolicySpecAction string

const (
	AuthorizationPolicySpecAction_ALLOW  AuthorizationPolicySpecAction = "ALLOW"
	AuthorizationPolicySpecAction_DENY   AuthorizationPolicySpecAction = "DENY"
	AuthorizationPolicySpecAction_AUDIT  AuthorizationPolicySpecAction = "AUDIT"
	AuthorizationPolicySpecAction_CUSTOM AuthorizationPolicySpecAction = "CUSTOM"
)

// Specifies detailed configuration of the CUSTOM action.
type AuthorizationPolicySpecProvider struct {
	// Specifies the name of the extension provider.
	Name *string `json:"name"`
}

type AuthorizationPolicySpecRules struct {
	// Optional.
	From *[]*AuthorizationPolicySpecRulesFrom `json:"from"`
	// Optional.
	To *[]*AuthorizationPolicySpecRulesTo `json:"to"`
	// Optional.
	When *[]*AuthorizationPolicySpecRulesWhen `json:"when"`
}

type AuthorizationPolicySpecRulesFrom struct {
	// Source specifies the source of a request.
	Source *AuthorizationPolicySpecRulesFromSource `json:"source"`
}

// Source specifies the source of a request.
type AuthorizationPolicySpecRulesFromSource struct {
	// Optional.
	IpBlocks *[]*string `json:"ipBlocks"`
	// Optional.
	Namespaces *[]*string `json:"namespaces"`
	// Optional.
	NotIpBlocks *[]*string `json:"notIpBlocks"`
	// Optional.
	NotNamespaces *[]*string `json:"notNamespaces"`
	// Optional.
	NotPrincipals *[]*string `json:"notPrincipals"`
	// Optional.
	NotRemoteIpBlocks *[]*string `json:"notRemoteIpBlocks"`
	// Optional.
	NotRequestPrincipals *[]*string `json:"notRequestPrincipals"`
	// Optional.
	Principals *[]*string `json:"principals"`
	// Optional.
	RemoteIpBlocks *[]*string `json:"remoteIpBlocks"`
	// Optional.
	RequestPrincipals *[]*string `json:"requestPrincipals"`
}

type AuthorizationPolicySpecRulesTo struct {
	// Operation specifies the operation of a request.
	Operation *AuthorizationPolicySpecRulesToOperation `json:"operation"`
}

// Operation specifies the operation of a request.
type AuthorizationPolicySpecRulesToOperation struct {
	// Optional.
	Hosts *[]*string `json:"hosts"`
	// Optional.
	Methods *[]*string `json:"methods"`
	// Optional.
	NotHosts *[]*string `json:"notHosts"`
	// Optional.
	NotMethods *[]*string `json:"notMethods"`
	// Optional.
	NotPaths *[]*string `json:"notPaths"`
	// Optional.
	NotPorts *[]*string `json:"notPorts"`
	// Optional.
	Paths *[]*string `json:"paths"`
	// Optional.
	Ports *[]*string `json:"ports"`
}

type AuthorizationPolicySpecRulesWhen struct {
	// The name of an Istio attribute.
	Key *string `json:"key"`
	// Optional.
	NotValues *[]*string `json:"notValues"`
	// Optional.
	Values *[]*string `json:"values"`
}

// Optional.
type AuthorizationPolicySpecSelector struct {
	MatchLabels *map[string]*string `json:"matchLabels"`
}

type PeerAuthentication interface {
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

// The jsii proxy struct for PeerAuthentication
type jsiiProxy_PeerAuthentication struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_PeerAuthentication) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PeerAuthentication) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PeerAuthentication) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PeerAuthentication) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PeerAuthentication) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PeerAuthentication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "PeerAuthentication" API object.
func NewPeerAuthentication(scope constructs.Construct, id *string, props *PeerAuthenticationProps) PeerAuthentication {
	_init_.Initialize()

	j := jsiiProxy_PeerAuthentication{}

	_jsii_.Create(
		"securityistioio.PeerAuthentication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "PeerAuthentication" API object.
func NewPeerAuthentication_Override(p PeerAuthentication, scope constructs.Construct, id *string, props *PeerAuthenticationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"securityistioio.PeerAuthentication",
		[]interface{}{scope, id, props},
		p,
	)
}

// Renders a Kubernetes manifest for "PeerAuthentication".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func PeerAuthentication_Manifest(props *PeerAuthenticationProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"securityistioio.PeerAuthentication",
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
func PeerAuthentication_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"securityistioio.PeerAuthentication",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func PeerAuthentication_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"securityistioio.PeerAuthentication",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (p *jsiiProxy_PeerAuthentication) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (p *jsiiProxy_PeerAuthentication) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PeerAuthentication) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (p *jsiiProxy_PeerAuthentication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PeerAuthentication) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (p *jsiiProxy_PeerAuthentication) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PeerAuthentication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PeerAuthenticationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// PeerAuthentication defines how traffic will be tunneled (or not) to the sidecar.
	Spec *PeerAuthenticationSpec `json:"spec"`
}

// PeerAuthentication defines how traffic will be tunneled (or not) to the sidecar.
type PeerAuthenticationSpec struct {
	// Mutual TLS settings for workload.
	Mtls *PeerAuthenticationSpecMtls `json:"mtls"`
	// Port specific mutual TLS settings.
	PortLevelMtls *map[string]*PeerAuthenticationSpecPortLevelMtls `json:"portLevelMtls"`
	// The selector determines the workloads to apply the ChannelAuthentication on.
	Selector *PeerAuthenticationSpecSelector `json:"selector"`
}

// Mutual TLS settings for workload.
type PeerAuthenticationSpecMtls struct {
	// Defines the mTLS mode used for peer authentication.
	Mode PeerAuthenticationSpecMtlsMode `json:"mode"`
}

// Defines the mTLS mode used for peer authentication.
type PeerAuthenticationSpecMtlsMode string

const (
	PeerAuthenticationSpecMtlsMode_UNSET      PeerAuthenticationSpecMtlsMode = "UNSET"
	PeerAuthenticationSpecMtlsMode_DISABLE    PeerAuthenticationSpecMtlsMode = "DISABLE"
	PeerAuthenticationSpecMtlsMode_PERMISSIVE PeerAuthenticationSpecMtlsMode = "PERMISSIVE"
	PeerAuthenticationSpecMtlsMode_STRICT     PeerAuthenticationSpecMtlsMode = "STRICT"
)

type PeerAuthenticationSpecPortLevelMtls struct {
	// Defines the mTLS mode used for peer authentication.
	Mode PeerAuthenticationSpecPortLevelMtlsMode `json:"mode"`
}

// Defines the mTLS mode used for peer authentication.
type PeerAuthenticationSpecPortLevelMtlsMode string

const (
	PeerAuthenticationSpecPortLevelMtlsMode_UNSET      PeerAuthenticationSpecPortLevelMtlsMode = "UNSET"
	PeerAuthenticationSpecPortLevelMtlsMode_DISABLE    PeerAuthenticationSpecPortLevelMtlsMode = "DISABLE"
	PeerAuthenticationSpecPortLevelMtlsMode_PERMISSIVE PeerAuthenticationSpecPortLevelMtlsMode = "PERMISSIVE"
	PeerAuthenticationSpecPortLevelMtlsMode_STRICT     PeerAuthenticationSpecPortLevelMtlsMode = "STRICT"
)

// The selector determines the workloads to apply the ChannelAuthentication on.
type PeerAuthenticationSpecSelector struct {
	MatchLabels *map[string]*string `json:"matchLabels"`
}

type RequestAuthentication interface {
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

// The jsii proxy struct for RequestAuthentication
type jsiiProxy_RequestAuthentication struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_RequestAuthentication) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthentication) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthentication) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthentication) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthentication) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthentication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

// Defines a "RequestAuthentication" API object.
func NewRequestAuthentication(scope constructs.Construct, id *string, props *RequestAuthenticationProps) RequestAuthentication {
	_init_.Initialize()

	j := jsiiProxy_RequestAuthentication{}

	_jsii_.Create(
		"securityistioio.RequestAuthentication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "RequestAuthentication" API object.
func NewRequestAuthentication_Override(r RequestAuthentication, scope constructs.Construct, id *string, props *RequestAuthenticationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"securityistioio.RequestAuthentication",
		[]interface{}{scope, id, props},
		r,
	)
}

// Renders a Kubernetes manifest for "RequestAuthentication".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func RequestAuthentication_Manifest(props *RequestAuthenticationProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"securityistioio.RequestAuthentication",
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
func RequestAuthentication_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"securityistioio.RequestAuthentication",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func RequestAuthentication_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"securityistioio.RequestAuthentication",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (r *jsiiProxy_RequestAuthentication) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (r *jsiiProxy_RequestAuthentication) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_RequestAuthentication) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (r *jsiiProxy_RequestAuthentication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_RequestAuthentication) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (r *jsiiProxy_RequestAuthentication) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_RequestAuthentication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RequestAuthenticationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// RequestAuthentication defines what request authentication methods are supported by a workload.
	Spec *RequestAuthenticationSpec `json:"spec"`
}

// RequestAuthentication defines what request authentication methods are supported by a workload.
type RequestAuthenticationSpec struct {
	// Define the list of JWTs that can be validated at the selected workloads' proxy.
	JwtRules *[]*RequestAuthenticationSpecJwtRules `json:"jwtRules"`
	// The selector determines the workloads to apply the RequestAuthentication on.
	Selector *RequestAuthenticationSpecSelector `json:"selector"`
}

type RequestAuthenticationSpecJwtRules struct {
	Audiences *[]*string `json:"audiences"`
	// If set to true, the orginal token will be kept for the ustream request.
	ForwardOriginalToken *bool `json:"forwardOriginalToken"`
	// List of header locations from which JWT is expected.
	FromHeaders *[]*RequestAuthenticationSpecJwtRulesFromHeaders `json:"fromHeaders"`
	// List of query parameters from which JWT is expected.
	FromParams *[]*string `json:"fromParams"`
	// Identifies the issuer that issued the JWT.
	Issuer *string `json:"issuer"`
	// JSON Web Key Set of public keys to validate signature of the JWT.
	Jwks                  *string `json:"jwks"`
	JwksUri               *string `json:"jwksUri"`
	OutputPayloadToHeader *string `json:"outputPayloadToHeader"`
}

type RequestAuthenticationSpecJwtRulesFromHeaders struct {
	// The HTTP header name.
	Name *string `json:"name"`
	// The prefix that should be stripped before decoding the token.
	Prefix *string `json:"prefix"`
}

// The selector determines the workloads to apply the RequestAuthentication on.
type RequestAuthenticationSpecSelector struct {
	MatchLabels *map[string]*string `json:"matchLabels"`
}
