package wrapper

import (
	v12 "github.com/openshift/api/project/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProjectList []v12.Project

func ListProjects(ns string, options v1.ListOptions) (ProjectList, error) {
	projectsApi, err := GetProjectApi(ns)
	if err != nil {
		return nil, err
	}
	projects, err := projectsApi.List(options)
	if err != nil {
		return nil, err
	}
	return projects.Items, nil
}

func GetProjectByName(ns string, name string, options v1.GetOptions) (*v12.Project, error) {
	projectsApi, err := GetProjectApi(ns)
	if err != nil {
		return nil, err
	}
	return projectsApi.Get(name, options)
}

func UpdateProject(ns string, project *v12.Project) (*v12.Project, error) {
	projectsApi, err := GetProjectApi(ns)
	if err != nil {
		return nil, err
	}
	return projectsApi.Update(project)
}

func CreateProject(ns string, project *v12.Project) (*v12.Project, error) {
	projectsApi, err := GetProjectApi(ns)
	if err != nil {
		return nil, err
	}
	return projectsApi.Create(project)
}

func DeleteProject(ns string, name string, options v1.DeleteOptions) error {
	projectsApi, err := GetProjectApi(ns)
	if err != nil {
		return err
	}
	return projectsApi.Delete(name, &options)
}

func GetProjectJson(ns string, name string, options v1.GetOptions) (string, error) {
	project, err := GetProjectByName(ns, name, options)
	if err != nil {
		return "", err
	}
	projectData, err := ObjectToJsonString(project)
	if err != nil {
		return "", err
	}
	return string(projectData), nil
}
