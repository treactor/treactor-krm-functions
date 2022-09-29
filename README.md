# treactor-kpt-functions

This repository contains the supporting KRM functions to create the 
Kubernetes Resources to deploy "Telemetry Reactor". Telemetry Reactor
is a playground project, mainly used to test and demonstrate several
concept in the "Cloud Native" space. Originally it focused on Telemetry,
hence it's name, but it's starting to grow beyond the original scope.

KRM functions are a nice fit to create the Kubernetes Resources as TReactor
has a lot of permutations possible for deployments (from a few micro services
to 100+ services).

The KRM functions are tested with `kpt` (https://kpt.dev), but in theory 
should be independently usable, as long as you have a compliant KRM function
operator, see the [KRM Function Spec](https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/docs/api-conventions/functions-spec.md). 
for more details.

## Treactor Functions

* **create-atoms**: create the deployments, services and paths
* **ensure-sa**: make sure each pod has a k8s service account
* **ensure-telemetry**: add telemetry environment variables
* **gateway-api**: understands the gateway-api resources

### create-atoms

create the deployments, services and paths

```yaml
apiVersion: fn.treactor.io/v1alpha1
kind: CreateAtoms
metadata:
name: create-atoms
annotations:
  config.kubernetes.io/local-config: "true"
maxBond: 0
maxNumber: 0
frameworks:
- name: "treactorGo"
  image: "gcr.io/treactor/treactor-go"
  imageTag: "latest"
- name: "treactorJava"
  image: "gcr.io/treactor/treactor-java"
  imageTag: "latest"
- name: "treactorNode"
  image: "gcr.io/treactor/treactor-node"
  imageTag: "latest"
numbers:
  - number: 2
    framework: "treactorNode"
  - number: 3
    framework: "treactorJava"
groups:
  - group: 1
    framework: "treactorGo"
```

#### Configuration main parameters

#### Configuring Frameworks

* treactorGo
* treactorNode
* treactorJava


#### Implementation nodes

### ensure-sa

Ensures that each deployment has it's own unique service account. It will loop 
over each deployment and create a kubernetes service account if it didn't exist
yet. If a deployment disappears it will remove the service account if it was 
managed by this function.

#### Implementation nodes

This is a fairly simple KRM function. It doesn't have a configuration and will
loop over each deployment, add the service account reference in the pod template
spec, and create the actual service account reference.

The created resources have the `config.kubernetes.io/managed-by: fn:gen:gcr.io/treactor/kpt-fn/ensure-sa`
annotation, so it will only delete the service account that the function created
in a previous execution.

### ensure-telemetry

Adds telemetry environment variables.

### gateway-api

understands the gateway-api resources

* **noop** : Disable any processing 
* **envoy** : Create an Istio proxy
* **istio** : Create Istio virtual services
* **gateway-api** : Keep/Restore the GatewayApi resources

```yaml
apiVersion: fn.treactor.io/v1alpha1
kind: GatewayAPI
metadata:
  name: gateway-api
  annotations:
    config.kubernetes.io/local-config: "true"
output: envoy
```

#### Configuring Envoy mode


#### Configuring Istio mode


#### Configure Ingress mode


## History

Telemetry Reactor is a configurable playground project with lots of 
micro-services. Creating Kubernetes resources for so may deployments
was very impractical. In the very beginning a go program was created
do generate the Kubernetes Resource Model (KRM) resources. Although it
worked, it doesn't reflect what happens in the "real" world.

A second attempt was done to create a "industry" standard helm chart.
Although it works, a lot of mind bends were necessary to create the 
chart. Although "industry" standard I have never been happy with it, and
it was not a nice experience.

With the introduction of `kpt` and the use of KRM functions it seems
that we finally have a model that I'm happy with. It's powerful, with
the functions, but still simple enough to be understandable. For a lot
of kubernetes deployments this seems like a good alternative for an 
operator (as long as you don't need lifecycle management and runtime
information).

