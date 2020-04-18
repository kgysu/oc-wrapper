# openshift/client-go Wrapper

Wrapps up the main functionality used to manage your applications on an Openshift Cluster. 
Setup your app locally and deploy it to remote cluster with only few lines of code. This 
makes it easy to test new poc setups and also to manage large applications with many items.


## Install

You can add oc-wrapper as a go-dependency like this:

```
$ go get github.com/kgysu/oc-wrapper
```

Or directly in code like:

```go
package main
import "github.com/kgysu/oc-wrapper/<submodule>"
```

## How to use it

See [examples](/examples).


## Types

The following Openshift-Item-Types are supported:
 - DeploymentConfig
 - StatefulSet
 - Service
 - Route
 - ConfigMap
 - PersistentVolumeClaim
 - ServiceAccount
 - Role
 - RoleBinding
 - Pod*

(* read only)


## License

Wrapper(oc-wrapper) is licensed under the Apache License 2.0.

