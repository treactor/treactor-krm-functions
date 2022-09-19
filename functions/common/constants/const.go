package constants

const (
	FnConfigGroup      = "fn.treactor.io"
	FnConfigVersion    = "v1alpha1"
	FnConfigAPIVersion = FnConfigGroup + "/" + FnConfigVersion
	FnUriPrefix        = "fn:gen:gcr.io/treactor/kpt-fn/"
	//legacyFnConfigKind = "SetAnnotationConfig"
	//fnConfigKind       = "SetAnnotations"
	//fnUri              = "fn:gen:gcr.io/treactor/kpt-fn/create-atoms"
	K8sAnnLocalConfig = "config.kubernetes.io/local-config"
)
