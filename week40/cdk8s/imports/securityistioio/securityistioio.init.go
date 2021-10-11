package securityistioio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"securityistioio.AuthorizationPolicy",
		reflect.TypeOf((*AuthorizationPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AuthorizationPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicyProps",
		reflect.TypeOf((*AuthorizationPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpec",
		reflect.TypeOf((*AuthorizationPolicySpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"securityistioio.AuthorizationPolicySpecAction",
		reflect.TypeOf((*AuthorizationPolicySpecAction)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW":  AuthorizationPolicySpecAction_ALLOW,
			"DENY":   AuthorizationPolicySpecAction_DENY,
			"AUDIT":  AuthorizationPolicySpecAction_AUDIT,
			"CUSTOM": AuthorizationPolicySpecAction_CUSTOM,
		},
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpecProvider",
		reflect.TypeOf((*AuthorizationPolicySpecProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpecRules",
		reflect.TypeOf((*AuthorizationPolicySpecRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpecRulesFrom",
		reflect.TypeOf((*AuthorizationPolicySpecRulesFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpecRulesFromSource",
		reflect.TypeOf((*AuthorizationPolicySpecRulesFromSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpecRulesTo",
		reflect.TypeOf((*AuthorizationPolicySpecRulesTo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpecRulesToOperation",
		reflect.TypeOf((*AuthorizationPolicySpecRulesToOperation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpecRulesWhen",
		reflect.TypeOf((*AuthorizationPolicySpecRulesWhen)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.AuthorizationPolicySpecSelector",
		reflect.TypeOf((*AuthorizationPolicySpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"securityistioio.PeerAuthentication",
		reflect.TypeOf((*PeerAuthentication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PeerAuthentication{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"securityistioio.PeerAuthenticationProps",
		reflect.TypeOf((*PeerAuthenticationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.PeerAuthenticationSpec",
		reflect.TypeOf((*PeerAuthenticationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.PeerAuthenticationSpecMtls",
		reflect.TypeOf((*PeerAuthenticationSpecMtls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"securityistioio.PeerAuthenticationSpecMtlsMode",
		reflect.TypeOf((*PeerAuthenticationSpecMtlsMode)(nil)).Elem(),
		map[string]interface{}{
			"UNSET":      PeerAuthenticationSpecMtlsMode_UNSET,
			"DISABLE":    PeerAuthenticationSpecMtlsMode_DISABLE,
			"PERMISSIVE": PeerAuthenticationSpecMtlsMode_PERMISSIVE,
			"STRICT":     PeerAuthenticationSpecMtlsMode_STRICT,
		},
	)
	_jsii_.RegisterStruct(
		"securityistioio.PeerAuthenticationSpecPortLevelMtls",
		reflect.TypeOf((*PeerAuthenticationSpecPortLevelMtls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"securityistioio.PeerAuthenticationSpecPortLevelMtlsMode",
		reflect.TypeOf((*PeerAuthenticationSpecPortLevelMtlsMode)(nil)).Elem(),
		map[string]interface{}{
			"UNSET":      PeerAuthenticationSpecPortLevelMtlsMode_UNSET,
			"DISABLE":    PeerAuthenticationSpecPortLevelMtlsMode_DISABLE,
			"PERMISSIVE": PeerAuthenticationSpecPortLevelMtlsMode_PERMISSIVE,
			"STRICT":     PeerAuthenticationSpecPortLevelMtlsMode_STRICT,
		},
	)
	_jsii_.RegisterStruct(
		"securityistioio.PeerAuthenticationSpecSelector",
		reflect.TypeOf((*PeerAuthenticationSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"securityistioio.RequestAuthentication",
		reflect.TypeOf((*RequestAuthentication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RequestAuthentication{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"securityistioio.RequestAuthenticationProps",
		reflect.TypeOf((*RequestAuthenticationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.RequestAuthenticationSpec",
		reflect.TypeOf((*RequestAuthenticationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.RequestAuthenticationSpecJwtRules",
		reflect.TypeOf((*RequestAuthenticationSpecJwtRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.RequestAuthenticationSpecJwtRulesFromHeaders",
		reflect.TypeOf((*RequestAuthenticationSpecJwtRulesFromHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"securityistioio.RequestAuthenticationSpecSelector",
		reflect.TypeOf((*RequestAuthenticationSpecSelector)(nil)).Elem(),
	)
}
