package converter

import (
	"fmt"
	v1 "github.com/openshift/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/diff"
	"reflect"
	"testing"
)

var sampleJson = []byte(`{
  "metadata": {
    "name": "Test",
    "creationTimestamp": null
  },
  "spec": {
    "strategy": {
      "resources": {}
    },
    "triggers": null,
    "replicas": 0,
    "test": false
  },
  "status": {
    "latestVersion": 0,
    "observedGeneration": 0,
    "replicas": 0,
    "updatedReplicas": 0,
    "availableReplicas": 0,
    "unavailableReplicas": 0
  }
}`)

var sampleYaml = []byte(`typemeta:
  kind: ""
  apiversion: ""
objectmeta:
  name: Test
  generatename: ""
  namespace: ""
  selflink: ""
  uid: ""
  resourceversion: ""
  generation: 0
  creationtimestamp: "0001-01-01T00:00:00Z"
  deletiontimestamp: null
  deletiongraceperiodseconds: null
  labels: {}
  annotations: {}
  ownerreferences: []
  finalizers: []
  clustername: ""
  managedfields: []
spec:
  strategy:
    type: ""
    customparams: null
    recreateparams: null
    rollingparams: null
    resources:
      limits: {}
      requests: {}
    labels: {}
    annotations: {}
    activedeadlineseconds: null
  minreadyseconds: 0
  triggers: []
  replicas: 0
  revisionhistorylimit: null
  test: false
  paused: false
  selector: {}
  template: null
status:
  latestversion: 0
  observedgeneration: 0
  replicas: 0
  updatedreplicas: 0
  availablereplicas: 0
  unavailablereplicas: 0
  details: null
  conditions: []
  readyreplicas: 0
`)

func TestConvertObjToYamlString(t *testing.T) {
	testObj := v1.DeploymentConfig{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name: "Test",
		},
		Spec:   v1.DeploymentConfigSpec{},
		Status: v1.DeploymentConfigStatus{},
	}

	yamlString, err := InterfaceToYaml(&testObj)
	if err != nil {
		t.Errorf("No error expected, error: cannot encode to yaml")
	}

	got := yamlString
	want := string(sampleYaml)

	if got != want {
		t.Errorf("got [%q] \nwant [%q]\n", got, want)
	}
}

func TestConvertYamlToObject(t *testing.T) {
	var got v1.DeploymentConfig
	err := YamlBytesToInterface(sampleYaml, &got)
	if err != nil {
		t.Errorf("No error expected, error: cannot decode yaml")
	}

	want := v1.DeploymentConfig{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name: "Test",
		},
		Spec:   v1.DeploymentConfigSpec{},
		Status: v1.DeploymentConfigStatus{},
	}

	//if !reflect.DeepEqual(want, got) {
	if want.Name != got.Name {
		fmt.Println("_____")
		fmt.Printf("%v", want)
		fmt.Println("_____")
		fmt.Printf("%v", got)
		fmt.Println("_____")
		fmt.Print(diff.ObjectDiff(want, got))
		fmt.Println("_____")
		t.Errorf("Objects not equal!")
	}
}

func TestConvertObjectToJson(t *testing.T) {
	testObj := v1.DeploymentConfig{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name: "Test",
		},
		Spec:   v1.DeploymentConfigSpec{},
		Status: v1.DeploymentConfigStatus{},
	}

	jsonString, err := InterfaceToJson(&testObj, true)
	if err != nil {
		t.Errorf("No error expected, error: cannot encode to yaml")
	}

	got := jsonString
	want := string(sampleJson)

	if got != want {
		t.Errorf("got=[%q] \n want=[%q]\n", got, want)
	}
}

func TestConvertJsonToObject(t *testing.T) {
	var got v1.DeploymentConfig
	err := JsonBytesToInterface(sampleJson, &got)
	if err != nil {
		t.Errorf("No error expected, error: cannot decode yaml")
	}

	want := v1.DeploymentConfig{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name: "Test",
		},
		Spec:   v1.DeploymentConfigSpec{},
		Status: v1.DeploymentConfigStatus{},
	}

	if !reflect.DeepEqual(want, got) {
		//if want.Name != got.Name {
		fmt.Println("_____")
		fmt.Printf("%v", want)
		fmt.Println("_____")
		fmt.Printf("%v", got)
		fmt.Println("_____")
		fmt.Print(diff.ObjectDiff(want, got))
		fmt.Println("_____")
		t.Errorf("Objects not equal!")
	}
}
