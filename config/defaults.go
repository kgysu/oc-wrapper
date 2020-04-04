package config

const defaultNamespaceEnvName = "NAMESPACE"

var namespaceEnvName = defaultNamespaceEnvName

func GetNamespaceEnvNameOrDefault() string {
	return namespaceEnvName
}
func SetNamespaceEnvName(name string) {
	namespaceEnvName = name
}

const defaultRootFolder = "/projects"

var rootFolder = defaultRootFolder

func GetRootFolderOrDefault() string {
	return rootFolder
}

func SetRootFolder(folder string) {
	rootFolder = folder
}
