package telemetryistioio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"telemetryistioio.Telemetry",
		reflect.TypeOf((*Telemetry)(nil)).Elem(),
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
			j := jsiiProxy_Telemetry{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetryProps",
		reflect.TypeOf((*TelemetryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpec",
		reflect.TypeOf((*TelemetrySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecSelector",
		reflect.TypeOf((*TelemetrySpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracing",
		reflect.TypeOf((*TelemetrySpecTracing)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingCustomTags",
		reflect.TypeOf((*TelemetrySpecTracingCustomTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingCustomTagsEnvironment",
		reflect.TypeOf((*TelemetrySpecTracingCustomTagsEnvironment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingCustomTagsHeader",
		reflect.TypeOf((*TelemetrySpecTracingCustomTagsHeader)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingCustomTagsLiteral",
		reflect.TypeOf((*TelemetrySpecTracingCustomTagsLiteral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingProviders",
		reflect.TypeOf((*TelemetrySpecTracingProviders)(nil)).Elem(),
	)
}
