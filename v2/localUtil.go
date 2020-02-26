package v2

import (
	v1 "github.com/openshift/api/apps/v1"
	v15 "github.com/openshift/api/route/v1"
	"k8s.io/api/apps/v1beta1"
	v12 "k8s.io/api/core/v1"
	v14 "k8s.io/api/rbac/v1"
	"strings"
)

func ListAllLocalItems(namespace string, folderPath string, debug bool) ([]OpenshiftItem, []error) {
	var items []OpenshiftItem
	var errs []error

	// Load all Files
	filePaths, err := filePathWalkDir(folderPath)
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
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, ConfigMapKey, local.String()))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, DeploymentConfigKey+JsonFileSuffix) {
			local := v1.DeploymentConfig{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, DeploymentConfigKey, local.String()))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, RoleKey+JsonFileSuffix) {
			local := v14.Role{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, RoleKey, local.String()))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, RoleBindingKey+JsonFileSuffix) {
			local := v14.RoleBinding{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, RoleBindingKey, local.String()))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, RouteKey+JsonFileSuffix) {
			local := v15.Route{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, RouteKey, local.String()))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, SecretKey+JsonFileSuffix) {
			local := v12.Secret{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, SecretKey, local.String()))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, ServiceKey+JsonFileSuffix) {
			local := v12.Service{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, ServiceKey, local.String()))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, ServiceAccountKey+JsonFileSuffix) {
			local := v12.ServiceAccount{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, ServiceAccountKey, local.String()))
		}
		if hasSuffixIgnoreCaseAndMinus(filePath, StatefulSetKey+JsonFileSuffix) {
			local := v1beta1.StatefulSet{}
			err = parseFileAndReplaceEnvVars(filePath, envs, &local, debug)
			checkErrorOrAppend(err, &errs, &items, NewOpenshiftItem(local.Name, StatefulSetKey, local.String()))
		}
	}
	return items, errs
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
