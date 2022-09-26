Annotate the Kptfile

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