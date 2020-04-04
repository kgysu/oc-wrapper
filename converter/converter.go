package converter

import (
	jsonEncoding "encoding/json"
	"github.com/openshift/api"
	"github.com/openshift/api/apps"
	v1 "github.com/openshift/api/apps/v1"
	"github.com/openshift/api/route"
	templatev1 "github.com/openshift/api/template/v1"
	"gopkg.in/yaml.v2"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	// The Kubernetes Go client (nested within the OpenShift Go client)
	// automatically registers its types in scheme.Scheme, however the
	// additional OpenShift types must be registered manually.  AddToScheme
	// registers the API group types (e.g. route.openshift.io/v1, Route) only.
	//appsv1.AddToScheme(scheme.Scheme)
	//authorizationv1.AddToScheme(scheme.Scheme)
	//buildv1.AddToScheme(scheme.Scheme)
	//imagev1.AddToScheme(scheme.Scheme)
	//networkv1.AddToScheme(scheme.Scheme)
	//oauthv1.AddToScheme(scheme.Scheme)
	//projectv1.AddToScheme(scheme.Scheme)
	//quotav1.AddToScheme(scheme.Scheme)
	//routev1.AddToScheme(scheme.Scheme)
	//securityv1.AddToScheme(scheme.Scheme)
	//templatev1.AddToScheme(scheme.Scheme)
	//userv1.AddToScheme(scheme.Scheme)

	err := api.Install(scheme.Scheme)
	if err != nil {
		panic(err)
	}

	err = api.InstallKube(scheme.Scheme)
	if err != nil {
		panic(err)
	}

	err = apps.Install(scheme.Scheme)
	if err != nil {
		panic(err)
	}

	err = v1.Install(scheme.Scheme)
	if err != nil {
		panic(err)
	}

	err = templatev1.Install(scheme.Scheme)
	if err != nil {
		panic(err)
	}

	err = route.Install(scheme.Scheme)
	if err != nil {
		panic(err)
	}
}

func ObjToYaml(obj runtime.Object, w io.Writer, prettyPrint bool, strict bool) error {
	serializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme,
		scheme.Scheme, json.SerializerOptions{
			Yaml:   true,
			Pretty: prettyPrint,
			Strict: strict,
		})

	err := serializer.Encode(obj, w)
	if err != nil {
		return err
	}
	return nil
}

func YamlToObject(data []byte, strict bool, obj runtime.Object) (runtime.Object, *schema.GroupVersionKind, error) {
	serializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme,
		scheme.Scheme, json.SerializerOptions{
			Yaml:   true,
			Strict: strict,
		})

	if obj == nil {
		resultObj, gvk, err := serializer.Decode(data, nil, nil)
		if err != nil {
			return nil, nil, err
		}
		return resultObj, gvk, nil
	} else {
		origGvk := obj.GetObjectKind().GroupVersionKind()
		resultObj, gvk, err := serializer.Decode(data, &origGvk, obj)
		if err != nil {
			return nil, nil, err
		}
		return resultObj, gvk, nil
	}
}

// TODO: check if needed
func RawToObject(rawExtension runtime.RawExtension, prettyPrint bool, strict bool) (runtime.Object, *schema.GroupVersionKind, error) {
	serializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme,
		scheme.Scheme, json.SerializerOptions{
			Yaml:   true,
			Pretty: prettyPrint,
			Strict: strict,
		})

	resultObj, gvk, err := serializer.Decode(rawExtension.Raw, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	return resultObj, gvk, nil
}

func RawToRealObject(rawExtension runtime.RawExtension, prettyPrint bool, strict bool, obj runtime.Object) (runtime.Object, *schema.GroupVersionKind, error) {
	serializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme,
		scheme.Scheme, json.SerializerOptions{
			Yaml:   true,
			Pretty: prettyPrint,
			Strict: strict,
		})

	origGvk := obj.GetObjectKind().GroupVersionKind()
	resultObj, gvk, err := serializer.Decode(rawExtension.Raw, &origGvk, obj)
	if err != nil {
		return nil, nil, err
	}

	return resultObj, gvk, nil
}

func InterfaceToJson(i interface{}, pretty bool) (string, error) {
	content, err := InterfaceToJsonBytes(i, pretty)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func InterfaceToJsonBytes(i interface{}, pretty bool) ([]byte, error) {
	if pretty {
		return jsonEncoding.MarshalIndent(i, "", "  ")
	}
	return jsonEncoding.Marshal(i)
}

func InterfaceToYaml(i interface{}) (string, error) {
	content, err := InterfaceToYamlBytes(i)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func InterfaceToYamlBytes(i interface{}) ([]byte, error) {
	return yaml.Marshal(i)
}

func YamlToInterface(input string, i interface{}) error {
	return YamlBytesToInterface([]byte(input), i)
}

// Does not work for Openshift Items
func YamlBytesToInterface(input []byte, i interface{}) error {
	return yaml.Unmarshal(input, i)
}

func JsonToInterface(input string, i interface{}) error {
	return JsonBytesToInterface([]byte(input), i)
}

func JsonBytesToInterface(input []byte, i interface{}) error {
	return jsonEncoding.Unmarshal(input, i)
}
