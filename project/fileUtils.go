package project

import (
	"bufio"
	"encoding/json"
	v1 "github.com/openshift/api/apps/v1"
	v15 "github.com/openshift/api/route/v1"
	"io/ioutil"
	v13 "k8s.io/api/apps/v1"
	"k8s.io/api/apps/v1beta1"
	v12 "k8s.io/api/core/v1"
	v14 "k8s.io/api/rbac/v1"
	"os"
	"path/filepath"
	"strings"
)

const JsonFileSuffix = ".json"
const EnvFileSuffix = ".env"
const ProjectSettingsFileName = "project" + JsonFileSuffix

type ThisProjectMeta struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

func parseLocalItemFiles(namespace string, folderPath string) ([]OpenshiftItem, error) {
	var items []OpenshiftItem

	// Load all Files
	filePaths, err := filePathWalkDir(folderPath)
	if err != nil {
		return nil, err
	}

	var envFiles []string
	var jsonFiles []string
	for _, filePath := range filePaths {
		if strings.HasSuffix(filePath, EnvFileSuffix) {
			envFiles = append(envFiles, filePath)
		}
		if strings.HasSuffix(filePath, JsonFileSuffix) {
			jsonFiles = append(jsonFiles, filePath)
		}
	}

	// Parse env files
	envs, err := loadEnvFile(namespace, envFiles)
	if err != nil {
		return nil, err
	}

	// Parse json files
	for _, filePath := range jsonFiles {
		if strings.HasSuffix(filePath, ConfigMapKey+JsonFileSuffix) {
			localConfigMap := v12.ConfigMap{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &localConfigMap)
			items = append(items, fromConfigMap(localConfigMap))
		}
		if strings.HasSuffix(filePath, DeploymentConfigKey+JsonFileSuffix) {
			localDeploymentConfig := v1.DeploymentConfig{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &localDeploymentConfig)
			items = append(items, fromDeploymentConfig(localDeploymentConfig))
		}
		if strings.HasSuffix(filePath, KubeStatefulSetKey+JsonFileSuffix) {
			localKubeStatefulSet := v13.StatefulSet{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &localKubeStatefulSet)
			items = append(items, fromKubeStatefulSet(localKubeStatefulSet))
		}
		if strings.HasSuffix(filePath, PvKey+JsonFileSuffix) {
			localPv := v12.PersistentVolume{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &localPv)
			items = append(items, fromPersistentVolume(localPv))
		}
		if strings.HasSuffix(filePath, PvClaimKey+JsonFileSuffix) {
			localPvClaim := v12.PersistentVolumeClaim{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &localPvClaim)
			items = append(items, fromPersistentVolumeClaim(localPvClaim))
		}
		if strings.HasSuffix(filePath, RoleKey+JsonFileSuffix) {
			local := v14.Role{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local)
			items = append(items, fromRole(local))
		}
		if strings.HasSuffix(filePath, RoleBindingKey+JsonFileSuffix) {
			local := v14.RoleBinding{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local)
			items = append(items, fromRoleBinding(local))
		}
		if strings.HasSuffix(filePath, RouteKey+JsonFileSuffix) {
			local := v15.Route{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local)
			items = append(items, fromRoute(local))
		}
		if strings.HasSuffix(filePath, SecretKey+JsonFileSuffix) {
			local := v12.Secret{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local)
			items = append(items, fromSecret(local))
		}
		if strings.HasSuffix(filePath, ServiceKey+JsonFileSuffix) {
			local := v12.Service{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local)
			items = append(items, fromService(local))
		}
		if strings.HasSuffix(filePath, ServiceAccountKey+JsonFileSuffix) {
			local := v12.ServiceAccount{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local)
			items = append(items, fromServiceAccount(local))
		}
		if strings.HasSuffix(filePath, StatefulSetKey+JsonFileSuffix) {
			local := v1beta1.StatefulSet{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local)
			items = append(items, fromStatefulSet(local))
		}
	}
	return items, nil
}

func parseProjectFile(filePath string) (*ThisProjectMeta, error) {
	data, err := ioutil.ReadFile(filePath + "/" + ProjectSettingsFileName)
	if err != nil {
		return nil, err
	}
	thisProjectMeta := &ThisProjectMeta{}
	err = json.Unmarshal(data, &thisProjectMeta)
	if err != nil {
		return nil, err
	}
	return thisProjectMeta, nil
}

func parseFileAndReplaceEnvVars(filePath string, envs map[string]string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	dataAsString := string(data)
	replacedString := dataAsString
	for key, value := range envs {
		replacedString = strings.ReplaceAll(replacedString, "$"+key, value)
	}
	err = json.Unmarshal([]byte(replacedString), &v)
	if err != nil {
		return err
	}
	return nil
}

func loadEnvFile(currentNs string, envFiles []string) (map[string]string, error) {
	fileNameSuffixForCurrentNs := currentNs + EnvFileSuffix
	currentEnvs := make(map[string]string)
	for _, envFile := range envFiles {
		if strings.HasSuffix(envFile, fileNameSuffixForCurrentNs) {
			envData, err := os.Open(envFile)
			if err != nil {
				return nil, err
			}
			scanner := bufio.NewScanner(envData)
			for scanner.Scan() {
				text := scanner.Text()
				if !(strings.HasPrefix(text, "#")) {
					splittedEnv := strings.SplitN(text, "=", 2)
					if len(splittedEnv) == 2 {
						currentEnvs[splittedEnv[0]] = splittedEnv[1]
					}
				}
			}
			if err := scanner.Err(); err != nil {
				return nil, err
			}
		}
	}
	return currentEnvs, nil
}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
