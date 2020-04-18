package items

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/converter"
	"github.com/kgysu/oc-wrapper/fileutils"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
)

var OpPersistentVolumeClaimTypeMeta = v12.TypeMeta{
	Kind:       "PersistentVolumeClaim",
	APIVersion: "v1",
}

type OpPersistentVolumeClaim struct {
	PersistentVolumeClaim *v1.PersistentVolumeClaim
}

func NewOpPersistentVolumeClaim(PersistentVolumeClaim v1.PersistentVolumeClaim) *OpPersistentVolumeClaim {
	if PersistentVolumeClaim.TypeMeta.Kind != OpPersistentVolumeClaimTypeMeta.Kind {
		PersistentVolumeClaim.TypeMeta = OpPersistentVolumeClaimTypeMeta
	}
	return &OpPersistentVolumeClaim{
		PersistentVolumeClaim: &PersistentVolumeClaim,
	}
}

// Methods

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) GetFileName() string {
	return fmt.Sprintf("%s-%s.yaml", oPersistentVolumeClaim.GetName(), oPersistentVolumeClaim.GetKind())
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) WriteToFile(file string) error {
	yamlContent, err := oPersistentVolumeClaim.ToYaml()
	if err != nil {
		return err
	}
	return fileutils.CreateFile(file, []byte(yamlContent))
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) LoadFromFile(file string, envs map[string]string) error {
	tempData, err := fileutils.ReadFile(file)
	if err != nil {
		return err
	}
	data := fileutils.ReplaceEnvs(string(tempData), envs)
	err = oPersistentVolumeClaim.FromData([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) Get(namespace string, restConf *rest.Config, name string) error {
	PersistentVolumeClaimInterface, err := client.GetPersistentVolumeClaimsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	PersistentVolumeClaim, err := PersistentVolumeClaimInterface.Get(name, v12.GetOptions{})
	if err != nil {
		return err
	}
	oPersistentVolumeClaim.PersistentVolumeClaim = PersistentVolumeClaim
	return nil
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) Create(namespace string, restConf *rest.Config) error {
	PersistentVolumeClaimInterface, err := client.GetPersistentVolumeClaimsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	_, err = PersistentVolumeClaimInterface.Create(oPersistentVolumeClaim.PersistentVolumeClaim)
	if err != nil {
		return err
	}
	return nil
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) Delete(namespace string, restConf *rest.Config, options *v12.DeleteOptions) error {
	PersistentVolumeClaimInterface, err := client.GetPersistentVolumeClaimsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	err = PersistentVolumeClaimInterface.Delete(oPersistentVolumeClaim.PersistentVolumeClaim.Name, options)
	if err != nil {
		return err
	}
	return nil
}

func (oPersistentVolumeClaim OpPersistentVolumeClaim) Update(namespace string, restConf *rest.Config) error {
	PersistentVolumeClaimInterface, err := client.GetPersistentVolumeClaimsInterface(namespace, restConf)
	if err != nil {
		return err
	}
	toUpdate, err := PersistentVolumeClaimInterface.Get(oPersistentVolumeClaim.GetName(), v12.GetOptions{})
	if err != nil {
		return err
	}
	toUpdate.Spec = oPersistentVolumeClaim.PersistentVolumeClaim.Spec
	toUpdate.Name = oPersistentVolumeClaim.PersistentVolumeClaim.Name
	toUpdate.Labels = oPersistentVolumeClaim.PersistentVolumeClaim.Labels
	_, err = PersistentVolumeClaimInterface.Update(toUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) UpdateScale(replicas int32, namespace string, restConf *rest.Config) error {
	return nil
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) GetScale() int32 {
	return 0
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) IsScalable() bool {
	return false
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) String() string {
	return fmt.Sprintf("%s %s ", oPersistentVolumeClaim.Info(), oPersistentVolumeClaim.Status())
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) Info() string {
	return fmt.Sprintf("[%s] %s ",
		oPersistentVolumeClaim.GetKind(),
		oPersistentVolumeClaim.GetName())
}

// TODO more infos
func (oPersistentVolumeClaim *OpPersistentVolumeClaim) Status() string {
	accessModes := ""
	for _, accessMode := range oPersistentVolumeClaim.PersistentVolumeClaim.Spec.AccessModes {
		accessModes = accessModes + ":" + string(accessMode)
	}
	return fmt.Sprintf("%s (%v) [%s]",
		oPersistentVolumeClaim.PersistentVolumeClaim.Spec.VolumeName,
		oPersistentVolumeClaim.PersistentVolumeClaim.Spec.StorageClassName,
		accessModes)

}

func (oPersistentVolumeClaim OpPersistentVolumeClaim) InfoStatusHtml() string {
	accessModes := ""
	for _, accessMode := range oPersistentVolumeClaim.PersistentVolumeClaim.Spec.AccessModes {
		accessModes = accessModes + createBadge("secondary", fmt.Sprintf("%s", accessMode))
	}
	return fmt.Sprint(
		createInfo(oPersistentVolumeClaim.GetKind(), oPersistentVolumeClaim.GetName()),
		createLabelBadges(oPersistentVolumeClaim.PersistentVolumeClaim.Labels),
		createStatusButton("secondary", fmt.Sprint("VolumeName ",
			createBadge("light", oPersistentVolumeClaim.PersistentVolumeClaim.Spec.VolumeName))),
		createStatusButton("secondary", fmt.Sprint("StorageClass ",
			createBadge("light", fmt.Sprint(oPersistentVolumeClaim.PersistentVolumeClaim.Spec.StorageClassName)))),
		accessModes,
	)
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) GetName() string {
	return oPersistentVolumeClaim.PersistentVolumeClaim.Name
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) GetKind() string {
	return oPersistentVolumeClaim.PersistentVolumeClaim.Kind
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) ToYaml() (string, error) {
	var contentBuilder strings.Builder
	err := converter.ObjToYaml(oPersistentVolumeClaim.PersistentVolumeClaim, &contentBuilder, true, true)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}

func (oPersistentVolumeClaim *OpPersistentVolumeClaim) FromData(data []byte) error {
	_, _, err := converter.YamlToObject(data, false, oPersistentVolumeClaim.PersistentVolumeClaim)
	if err != nil {
		return err
	}
	return nil
}
