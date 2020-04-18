package converter

import (
	"fmt"
	v14 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/diff"
	"testing"
)

var sampleYaml = []byte(`apiVersion: apps/v1
kind: StatefulSet
metadata:
  annotations:
    app: sample
  creationTimestamp: null
  labels:
    app: sample
  name: sample
spec:
  selector: null
  serviceName: ""
  template:
    metadata:
      annotations:
        app: sample
      creationTimestamp: null
      labels:
        app: sample
      name: sample
    spec:
      containers:
      - env:
        - name: SAMPLE
          value: VALUE
        image: sample:1.0
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /health
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 40
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 5
        name: sample
        ports:
        - containerPort: 8080
          name: basic
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /info
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 40
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 5
        resources:
          limits:
            cpu: "1"
            memory: 1Gi
          requests:
            cpu: 100m
            memory: 100Mi
      dnsPolicy: ClusterFirst
      restartPolicy: Always
  updateStrategy:
    type: RollingUpdate
status:
  replicas: 0
`)

var testObj = v14.StatefulSet{
	TypeMeta: metav1.TypeMeta{
		Kind:       "StatefulSet",
		APIVersion: "apps/v1",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:        "name",
		Labels:      map[string]string{"app": "name"},
		Annotations: map[string]string{"app": "name"},
	},
	Spec: v14.StatefulSetSpec{
		Template: v12.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "name",
				Labels:      map[string]string{"app": "name"},
				Annotations: map[string]string{"app": "name"},
			},
			Spec: v12.PodSpec{
				RestartPolicy: v12.RestartPolicyAlways,
				DNSPolicy:     v12.DNSClusterFirst,
				Containers: []v12.Container{
					{
						Name:  "name",
						Image: "sample:1.0",
						Ports: []v12.ContainerPort{
							{
								Name:          "basic",
								ContainerPort: 8080,
							},
						},
						Env: []v12.EnvVar{
							{
								Name:  "SAMPLE",
								Value: "VALUE",
							},
						},
						Resources: v12.ResourceRequirements{
							Limits: v12.ResourceList{
								"cpu":    resource.MustParse("1"),
								"memory": resource.MustParse("1Gi"),
							},
							Requests: v12.ResourceList{
								"cpu":    resource.MustParse("100m"),
								"memory": resource.MustParse("100Mi"),
							},
						},
						LivenessProbe: &v12.Probe{
							Handler: v12.Handler{
								HTTPGet: &v12.HTTPGetAction{
									Path:   "/health",
									Port:   intstr.FromInt(8080),
									Scheme: v12.URISchemeHTTP,
								},
							},
							InitialDelaySeconds: 40,
							TimeoutSeconds:      5,
							PeriodSeconds:       10,
							SuccessThreshold:    1,
							FailureThreshold:    3,
						},
						ReadinessProbe: &v12.Probe{
							Handler: v12.Handler{
								HTTPGet: &v12.HTTPGetAction{
									Path:   "/info",
									Port:   intstr.FromInt(8080),
									Scheme: v12.URISchemeHTTP,
								},
							},
							InitialDelaySeconds: 40,
							TimeoutSeconds:      5,
							PeriodSeconds:       10,
							SuccessThreshold:    1,
							FailureThreshold:    3,
						},
						ImagePullPolicy: v12.PullIfNotPresent,
					},
				},
			},
		},
		UpdateStrategy: v14.StatefulSetUpdateStrategy{
			Type: v14.RollingUpdateStatefulSetStrategyType,
		},
	},
	Status: v14.StatefulSetStatus{},
}

//func TestConvertObjToYamlString(t *testing.T) {
//	var contentBuilder strings.Builder
//	err := ObjToYaml(&testObj, &contentBuilder, true, false)
//	if err != nil {
//		t.Errorf("No error expected, error: cannot encode to yaml")
//	}
//
//	fmt.Println("----------")
//	fmt.Println(contentBuilder.String())
//	fmt.Println("----------")
//
//	got := contentBuilder.String()
//	want := string(sampleYaml)
//
//	if got != want {
//		t.Errorf("got [%q] \nwant [%q]\n", got, want)
//	}
//}

func TestConvertYamlToObject(t *testing.T) {
	var got v14.StatefulSet
	_, _, err := YamlToObject(sampleYaml, false, &got)
	if err != nil {
		t.Error(err)
	}
	want := testObj

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
