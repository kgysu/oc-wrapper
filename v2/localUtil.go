package v2

import (
	"encoding/json"
	v1 "github.com/openshift/api/apps/v1"
	v15 "github.com/openshift/api/route/v1"
	"k8s.io/api/apps/v1beta1"
	v12 "k8s.io/api/core/v1"
	v14 "k8s.io/api/rbac/v1"
	"path/filepath"
	"strings"
)

func ListAllLocalItems(namespace string, folderPath string, debug bool) ([]OpenshiftItem, []error) {
	var items []OpenshiftItem
	var errs []error

	// Load all Files
	filePaths, err := filePathWalkDir(filepath.FromSlash(folderPath))
	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}
	// Split
	envFiles, jsonFiles := splitByFileType(filePaths)

	// Load env files
	envs, err := loadEnvsByNamespace(namespace, envFiles, debug)
	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	// Parse json files
	for _, filePath := range jsonFiles {
		if hasSuffixIgnoreCaseAndMinus(filePath, ConfigMapKey+JsonFileSuffix) {
			local := v12.ConfigMap{}
			err = parseConfigMap(filePath, folderPath, namespace, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, ConfigMapKey, marshalltoString(local, &errs)))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, DeploymentConfigKey+JsonFileSuffix) {
			local := v1.DeploymentConfig{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, DeploymentConfigKey, marshalltoString(local, &errs)))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, RoleKey+JsonFileSuffix) {
			local := v14.Role{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, RoleKey, marshalltoString(local, &errs)))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, RoleBindingKey+JsonFileSuffix) {
			local := v14.RoleBinding{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, RoleBindingKey, marshalltoString(local, &errs)))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, RouteKey+JsonFileSuffix) {
			local := v15.Route{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, RouteKey, marshalltoString(local, &errs)))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, SecretKey+JsonFileSuffix) {
			local := v12.Secret{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, SecretKey, marshalltoString(local, &errs)))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, ServiceKey+JsonFileSuffix) {
			local := v12.Service{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, ServiceKey, marshalltoString(local, &errs)))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, ServiceAccountKey+JsonFileSuffix) {
			local := v12.ServiceAccount{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, ServiceAccountKey, marshalltoString(local, &errs)))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, StatefulSetKey+JsonFileSuffix) {
			local := v1beta1.StatefulSet{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, StatefulSetKey, marshalltoString(local, &errs)))
		}
	}
	return items, errs
}

func marshalltoString(local interface{}, errs *[]error) string {
	marshalled, err := json.Marshal(&local)
	if err != nil {
		*errs = append(*errs, err)
		return ""
	}
	return string(marshalled)
}

func checkErrorOrAppend(err error, errs *[]error, items *[]OpenshiftItem, item OpenshiftItem) {
	if err != nil {
		*errs = append(*errs, err)
	} else {
		*items = append(*items, item)
	}
}

func hasSuffixIgnoreCaseAndMinus(value string, suffix string) bool {
	valueNoMinus := strings.ReplaceAll(value, "-", "")
	valueLower := strings.ToLower(valueNoMinus)
	suffixLower := strings.ToLower(suffix)
	return strings.HasSuffix(valueLower, suffixLower)
}
