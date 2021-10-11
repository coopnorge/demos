package networkingistioio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"networkingistioio.DestinationRule",
		reflect.TypeOf((*DestinationRule)(nil)).Elem(),
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
			j := jsiiProxy_DestinationRule{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleProps",
		reflect.TypeOf((*DestinationRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpec",
		reflect.TypeOf((*DestinationRuleSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsets",
		reflect.TypeOf((*DestinationRuleSpecSubsets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicy",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPool",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT":        DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE":        DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancer",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"ROUND_ROBIN": DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_CONN":  DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_LEAST_CONN,
			"RANDOM":      DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_PASSTHROUGH,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyOutlierDetection",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettings",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPool",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT":        DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE":        DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancer",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"ROUND_ROBIN": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_CONN":  DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_CONN,
			"RANDOM":      DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_PASSTHROUGH,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsOutlierDetection",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsPort",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTls",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE":      DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_DISABLE,
			"SIMPLE":       DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_SIMPLE,
			"MUTUAL":       DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyTls",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyTlsMode",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE":      DestinationRuleSpecSubsetsTrafficPolicyTlsMode_DISABLE,
			"SIMPLE":       DestinationRuleSpecSubsetsTrafficPolicyTlsMode_SIMPLE,
			"MUTUAL":       DestinationRuleSpecSubsetsTrafficPolicyTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleSpecSubsetsTrafficPolicyTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicy",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPool",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT":        DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE":        DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancer",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"ROUND_ROBIN": DestinationRuleSpecTrafficPolicyLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_CONN":  DestinationRuleSpecTrafficPolicyLoadBalancerSimple_LEAST_CONN,
			"RANDOM":      DestinationRuleSpecTrafficPolicyLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleSpecTrafficPolicyLoadBalancerSimple_PASSTHROUGH,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyOutlierDetection",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettings",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPool",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT":        DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE":        DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancer",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"ROUND_ROBIN": DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_CONN":  DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_CONN,
			"RANDOM":      DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_PASSTHROUGH,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsOutlierDetection",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsPort",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsTls",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE":      DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_DISABLE,
			"SIMPLE":       DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_SIMPLE,
			"MUTUAL":       DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyTls",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyTlsMode",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE":      DestinationRuleSpecTrafficPolicyTlsMode_DISABLE,
			"SIMPLE":       DestinationRuleSpecTrafficPolicyTlsMode_SIMPLE,
			"MUTUAL":       DestinationRuleSpecTrafficPolicyTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleSpecTrafficPolicyTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterClass(
		"networkingistioio.EnvoyFilter",
		reflect.TypeOf((*EnvoyFilter)(nil)).Elem(),
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
			j := jsiiProxy_EnvoyFilter{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterProps",
		reflect.TypeOf((*EnvoyFilterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpec",
		reflect.TypeOf((*EnvoyFilterSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatches",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatches)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesApplyTo",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesApplyTo)(nil)).Elem(),
		map[string]interface{}{
			"INVALID":             EnvoyFilterSpecConfigPatchesApplyTo_INVALID,
			"LISTENER":            EnvoyFilterSpecConfigPatchesApplyTo_LISTENER,
			"FILTER_CHAIN":        EnvoyFilterSpecConfigPatchesApplyTo_FILTER_CHAIN,
			"NETWORK_FILTER":      EnvoyFilterSpecConfigPatchesApplyTo_NETWORK_FILTER,
			"HTTP_FILTER":         EnvoyFilterSpecConfigPatchesApplyTo_HTTP_FILTER,
			"ROUTE_CONFIGURATION": EnvoyFilterSpecConfigPatchesApplyTo_ROUTE_CONFIGURATION,
			"VIRTUAL_HOST":        EnvoyFilterSpecConfigPatchesApplyTo_VIRTUAL_HOST,
			"HTTP_ROUTE":          EnvoyFilterSpecConfigPatchesApplyTo_HTTP_ROUTE,
			"CLUSTER":             EnvoyFilterSpecConfigPatchesApplyTo_CLUSTER,
			"EXTENSION_CONFIG":    EnvoyFilterSpecConfigPatchesApplyTo_EXTENSION_CONFIG,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatch",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchCluster",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchCluster)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchContext",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchContext)(nil)).Elem(),
		map[string]interface{}{
			"ANY":              EnvoyFilterSpecConfigPatchesMatchContext_ANY,
			"SIDECAR_INBOUND":  EnvoyFilterSpecConfigPatchesMatchContext_SIDECAR_INBOUND,
			"SIDECAR_OUTBOUND": EnvoyFilterSpecConfigPatchesMatchContext_SIDECAR_OUTBOUND,
			"GATEWAY":          EnvoyFilterSpecConfigPatchesMatchContext_GATEWAY,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchListener",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchListener)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchListenerFilterChain",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchListenerFilterChain)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilter",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilterSubFilter",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilterSubFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchProxy",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchProxy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchRouteConfiguration",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchRouteConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhost",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhost)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRoute",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRoute)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction)(nil)).Elem(),
		map[string]interface{}{
			"ANY":             EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_ANY,
			"ROUTE":           EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_ROUTE,
			"REDIRECT":        EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_REDIRECT,
			"DIRECT_RESPONSE": EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_DIRECT_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesPatch",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesPatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesPatchFilterClass",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesPatchFilterClass)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": EnvoyFilterSpecConfigPatchesPatchFilterClass_UNSPECIFIED,
			"AUTHN":       EnvoyFilterSpecConfigPatchesPatchFilterClass_AUTHN,
			"AUTHZ":       EnvoyFilterSpecConfigPatchesPatchFilterClass_AUTHZ,
			"STATS":       EnvoyFilterSpecConfigPatchesPatchFilterClass_STATS,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesPatchOperation",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesPatchOperation)(nil)).Elem(),
		map[string]interface{}{
			"INVALID":       EnvoyFilterSpecConfigPatchesPatchOperation_INVALID,
			"MERGE":         EnvoyFilterSpecConfigPatchesPatchOperation_MERGE,
			"ADD":           EnvoyFilterSpecConfigPatchesPatchOperation_ADD,
			"REMOVE":        EnvoyFilterSpecConfigPatchesPatchOperation_REMOVE,
			"INSERT_BEFORE": EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_BEFORE,
			"INSERT_AFTER":  EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_AFTER,
			"INSERT_FIRST":  EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_FIRST,
			"REPLACE":       EnvoyFilterSpecConfigPatchesPatchOperation_REPLACE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecWorkloadSelector",
		reflect.TypeOf((*EnvoyFilterSpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.Gateway",
		reflect.TypeOf((*Gateway)(nil)).Elem(),
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
			j := jsiiProxy_Gateway{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewayProps",
		reflect.TypeOf((*GatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewaySpec",
		reflect.TypeOf((*GatewaySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewaySpecServers",
		reflect.TypeOf((*GatewaySpecServers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewaySpecServersPort",
		reflect.TypeOf((*GatewaySpecServersPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewaySpecServersTls",
		reflect.TypeOf((*GatewaySpecServersTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewaySpecServersTlsMaxProtocolVersion",
		reflect.TypeOf((*GatewaySpecServersTlsMaxProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": GatewaySpecServersTlsMaxProtocolVersion_TLS_AUTO,
			"TLSV1_0":  GatewaySpecServersTlsMaxProtocolVersion_TLSV1_0,
			"TLSV1_1":  GatewaySpecServersTlsMaxProtocolVersion_TLSV1_1,
			"TLSV1_2":  GatewaySpecServersTlsMaxProtocolVersion_TLSV1_2,
			"TLSV1_3":  GatewaySpecServersTlsMaxProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewaySpecServersTlsMinProtocolVersion",
		reflect.TypeOf((*GatewaySpecServersTlsMinProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": GatewaySpecServersTlsMinProtocolVersion_TLS_AUTO,
			"TLSV1_0":  GatewaySpecServersTlsMinProtocolVersion_TLSV1_0,
			"TLSV1_1":  GatewaySpecServersTlsMinProtocolVersion_TLSV1_1,
			"TLSV1_2":  GatewaySpecServersTlsMinProtocolVersion_TLSV1_2,
			"TLSV1_3":  GatewaySpecServersTlsMinProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewaySpecServersTlsMode",
		reflect.TypeOf((*GatewaySpecServersTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"PASSTHROUGH":      GatewaySpecServersTlsMode_PASSTHROUGH,
			"SIMPLE":           GatewaySpecServersTlsMode_SIMPLE,
			"MUTUAL":           GatewaySpecServersTlsMode_MUTUAL,
			"AUTO_PASSTHROUGH": GatewaySpecServersTlsMode_AUTO_PASSTHROUGH,
			"ISTIO_MUTUAL":     GatewaySpecServersTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterClass(
		"networkingistioio.ServiceEntry",
		reflect.TypeOf((*ServiceEntry)(nil)).Elem(),
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
			j := jsiiProxy_ServiceEntry{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntryProps",
		reflect.TypeOf((*ServiceEntryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntrySpec",
		reflect.TypeOf((*ServiceEntrySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntrySpecEndpoints",
		reflect.TypeOf((*ServiceEntrySpecEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.ServiceEntrySpecLocation",
		reflect.TypeOf((*ServiceEntrySpecLocation)(nil)).Elem(),
		map[string]interface{}{
			"MESH_EXTERNAL": ServiceEntrySpecLocation_MESH_EXTERNAL,
			"MESH_INTERNAL": ServiceEntrySpecLocation_MESH_INTERNAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntrySpecPorts",
		reflect.TypeOf((*ServiceEntrySpecPorts)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.ServiceEntrySpecResolution",
		reflect.TypeOf((*ServiceEntrySpecResolution)(nil)).Elem(),
		map[string]interface{}{
			"NONE":   ServiceEntrySpecResolution_NONE,
			"STATIC": ServiceEntrySpecResolution_STATIC,
			"DNS":    ServiceEntrySpecResolution_DNS,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntrySpecWorkloadSelector",
		reflect.TypeOf((*ServiceEntrySpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.Sidecar",
		reflect.TypeOf((*Sidecar)(nil)).Elem(),
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
			j := jsiiProxy_Sidecar{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarProps",
		reflect.TypeOf((*SidecarProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpec",
		reflect.TypeOf((*SidecarSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecEgress",
		reflect.TypeOf((*SidecarSpecEgress)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecEgressCaptureMode",
		reflect.TypeOf((*SidecarSpecEgressCaptureMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT":  SidecarSpecEgressCaptureMode_DEFAULT,
			"IPTABLES": SidecarSpecEgressCaptureMode_IPTABLES,
			"NONE":     SidecarSpecEgressCaptureMode_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecEgressPort",
		reflect.TypeOf((*SidecarSpecEgressPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecIngress",
		reflect.TypeOf((*SidecarSpecIngress)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecIngressCaptureMode",
		reflect.TypeOf((*SidecarSpecIngressCaptureMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT":  SidecarSpecIngressCaptureMode_DEFAULT,
			"IPTABLES": SidecarSpecIngressCaptureMode_IPTABLES,
			"NONE":     SidecarSpecIngressCaptureMode_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecIngressPort",
		reflect.TypeOf((*SidecarSpecIngressPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecOutboundTrafficPolicy",
		reflect.TypeOf((*SidecarSpecOutboundTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecOutboundTrafficPolicyEgressProxy",
		reflect.TypeOf((*SidecarSpecOutboundTrafficPolicyEgressProxy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecOutboundTrafficPolicyEgressProxyPort",
		reflect.TypeOf((*SidecarSpecOutboundTrafficPolicyEgressProxyPort)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecOutboundTrafficPolicyMode",
		reflect.TypeOf((*SidecarSpecOutboundTrafficPolicyMode)(nil)).Elem(),
		map[string]interface{}{
			"REGISTRY_ONLY": SidecarSpecOutboundTrafficPolicyMode_REGISTRY_ONLY,
			"ALLOW_ANY":     SidecarSpecOutboundTrafficPolicyMode_ALLOW_ANY,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecWorkloadSelector",
		reflect.TypeOf((*SidecarSpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.VirtualService",
		reflect.TypeOf((*VirtualService)(nil)).Elem(),
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
			j := jsiiProxy_VirtualService{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceProps",
		reflect.TypeOf((*VirtualServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpec",
		reflect.TypeOf((*VirtualServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttp",
		reflect.TypeOf((*VirtualServiceSpecHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpCorsPolicy",
		reflect.TypeOf((*VirtualServiceSpecHttpCorsPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpCorsPolicyAllowOrigins",
		reflect.TypeOf((*VirtualServiceSpecHttpCorsPolicyAllowOrigins)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpDelegate",
		reflect.TypeOf((*VirtualServiceSpecHttpDelegate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFault",
		reflect.TypeOf((*VirtualServiceSpecHttpFault)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFaultAbort",
		reflect.TypeOf((*VirtualServiceSpecHttpFaultAbort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFaultAbortPercentage",
		reflect.TypeOf((*VirtualServiceSpecHttpFaultAbortPercentage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFaultDelay",
		reflect.TypeOf((*VirtualServiceSpecHttpFaultDelay)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFaultDelayPercentage",
		reflect.TypeOf((*VirtualServiceSpecHttpFaultDelayPercentage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpHeaders",
		reflect.TypeOf((*VirtualServiceSpecHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpHeadersRequest",
		reflect.TypeOf((*VirtualServiceSpecHttpHeadersRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpHeadersResponse",
		reflect.TypeOf((*VirtualServiceSpecHttpHeadersResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatch",
		reflect.TypeOf((*VirtualServiceSpecHttpMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchAuthority",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchAuthority)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchHeaders",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchMethod",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchMethod)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchQueryParams",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchQueryParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchScheme",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchScheme)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchUri",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchUri)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchWithoutHeaders",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchWithoutHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMirror",
		reflect.TypeOf((*VirtualServiceSpecHttpMirror)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMirrorPercentage",
		reflect.TypeOf((*VirtualServiceSpecHttpMirrorPercentage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMirrorPort",
		reflect.TypeOf((*VirtualServiceSpecHttpMirrorPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRedirect",
		reflect.TypeOf((*VirtualServiceSpecHttpRedirect)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRetries",
		reflect.TypeOf((*VirtualServiceSpecHttpRetries)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRewrite",
		reflect.TypeOf((*VirtualServiceSpecHttpRewrite)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRoute",
		reflect.TypeOf((*VirtualServiceSpecHttpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteDestination",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteDestinationPort",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteDestinationPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteHeaders",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteHeadersRequest",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteHeadersRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteHeadersResponse",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteHeadersResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcp",
		reflect.TypeOf((*VirtualServiceSpecTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcpMatch",
		reflect.TypeOf((*VirtualServiceSpecTcpMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcpRoute",
		reflect.TypeOf((*VirtualServiceSpecTcpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcpRouteDestination",
		reflect.TypeOf((*VirtualServiceSpecTcpRouteDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcpRouteDestinationPort",
		reflect.TypeOf((*VirtualServiceSpecTcpRouteDestinationPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTls",
		reflect.TypeOf((*VirtualServiceSpecTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTlsMatch",
		reflect.TypeOf((*VirtualServiceSpecTlsMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTlsRoute",
		reflect.TypeOf((*VirtualServiceSpecTlsRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTlsRouteDestination",
		reflect.TypeOf((*VirtualServiceSpecTlsRouteDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTlsRouteDestinationPort",
		reflect.TypeOf((*VirtualServiceSpecTlsRouteDestinationPort)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.WorkloadEntry",
		reflect.TypeOf((*WorkloadEntry)(nil)).Elem(),
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
			j := jsiiProxy_WorkloadEntry{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadEntryProps",
		reflect.TypeOf((*WorkloadEntryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadEntrySpec",
		reflect.TypeOf((*WorkloadEntrySpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.WorkloadGroup",
		reflect.TypeOf((*WorkloadGroup)(nil)).Elem(),
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
			j := jsiiProxy_WorkloadGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupProps",
		reflect.TypeOf((*WorkloadGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpec",
		reflect.TypeOf((*WorkloadGroupSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecMetadata",
		reflect.TypeOf((*WorkloadGroupSpecMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbe",
		reflect.TypeOf((*WorkloadGroupSpecProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbeExec",
		reflect.TypeOf((*WorkloadGroupSpecProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbeHttpGet",
		reflect.TypeOf((*WorkloadGroupSpecProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkloadGroupSpecProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbeTcpSocket",
		reflect.TypeOf((*WorkloadGroupSpecProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecTemplate",
		reflect.TypeOf((*WorkloadGroupSpecTemplate)(nil)).Elem(),
	)
}
