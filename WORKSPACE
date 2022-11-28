# Declare the local Bazel workspace.
workspace(
    name = "com_fasterci_rules_mirror",
)
# gazelle:repo bazel_gazelle

# optionally override the default version of rules_gitops with a local copy
# local_repository(
#     name = "com_adobe_rules_gitops",
#     path = "../rules_gitops",
# )

load(":internal_deps.bzl", "rules_mirror_internal_deps")

# Fetch deps needed only locally for development
rules_mirror_internal_deps()

load("//mirror:repositories.bzl", "rules_mirror_dependencies")

# Fetch dependencies which users need as well
rules_mirror_dependencies()

# For running our own unit tests
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

############################################
# Gazelle, for generating bzl_library targets
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:go_repositories.bzl", "go_dependencies")

# gazelle:repository_macro go_repositories.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.19.3")

gazelle_dependencies()

############################################
# Setup for rules_gitops
load("@com_adobe_rules_gitops//gitops:deps.bzl", "rules_gitops_dependencies")

rules_gitops_dependencies()

load("@com_adobe_rules_gitops//gitops:repositories.bzl", "rules_gitops_repositories")

rules_gitops_repositories()

load("@com_adobe_rules_gitops//skylib:k8s.bzl", "kubeconfig")

kubeconfig(
    name = "k8s_test",
    cluster = "testcluster",
    use_host_config = True,
)
