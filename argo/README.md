This is a **proof of concept** to integrate kpt with argo. It actually relies on kpt having rendered the 
artifacts completely in git (a typical porch workflow).

It relies on two crazy simple argo plugins (both plugin's are less that 100 lines)

The first on is `kpt-repo`. This is an `app-of-app`'s like argo Application. It looks like this:

```yaml
project: default
source:
  repoURL: 'https://github.com/alexvanboxel/deployments.git'
  path: .
  targetRevision: HEAD
  plugin:
    name: kpt-repo
destination:
  server: 'https://kubernetes.default.svc'
  namespace: argocd
syncPolicy:
  automated:
    prune: true
    selfHeal: true
```

You give it a repository full of deployments, the only thing it looks for though it the `KptArgoRepo` file in the root,
here is an example:

```yaml
apiVersion: argocd.kpt.dev/v1alpha1
kind: KptArgoRepo
metadata:
  name: deployments
  annotations:
    config.kubernetes.io/local-config: "true"
repoUrl: https://github.com/alexvanboxel/deployments.git
namePrefix: deployments-
```

`kpt-repo` uses this fall to recursively traverse all the deployment repo's and look for `Kptfile`'s. Then it looks
in the file and searches for the following annotations:

```yaml
apiVersion: kpt.dev/v1
kind: Kptfile
metadata:
  name: test
  annotations:
    kpt-repo.argocd.kpt.dev/destination-name: "in-cluster"
    kpt-repo.argocd.kpt.dev/project-name: "default"
    kpt-repo.argocd.kpt.dev/sync-policy: "automatic"
```

For each `Kptfile` it will create an Argo Application, using the annotations. That Application targets the second
plugin, `kpt-render`, that simply filters out the KRM files with `config.kubernetes.io/local-config` set to `true`.

Installing

As this is a proof of concept, the install is a bit hacky. A custom Argo build file is included
(see `Dockerfile`). Build the image and replace the image in your deployments of
`argocd-server`.

Add the following too your `argocd-cm` config:

```yaml
  configManagementPlugins: "- name: kpt-repo\n  generate:\n    command: [\"/usr/local/bin/kpt-repo\"]\n
    \ lockRepo: true\n- name: kpt-render\n  generate: \n    command: [\"/usr/local/bin/kpt-render\"]
    \n  lockRepo: true \n"
```
