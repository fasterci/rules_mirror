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

#Fetch deps needed for local development"
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_oci",
    sha256 = "4a276e9566c03491649eef63f27c2816cc222f41ccdebd97d2c5159e84917c3b",
    strip_prefix = "rules_oci-1.7.4",
    url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.7.4/rules_oci-v1.7.4.tar.gz",
)

load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")

rules_oci_dependencies()

load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "oci_register_toolchains")

oci_register_toolchains(
    name = "oci",
    crane_version = LATEST_CRANE_VERSION,
    # Uncommenting the zot toolchain will cause it to be used instead of crane for some tasks.
    # Note that it does not support docker-format images.
    # zot_version = LATEST_ZOT_VERSION,
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
    ],
)

# gazelle is required by go_image_repositories
http_archive(
    name = "bazel_gazelle",
    sha256 = "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
    ],
)

# Override bazel_skylib distribution to fetch sources instead
# so that the gazelle extension is included
# see https://github.com/bazelbuild/bazel-skylib/issues/250
http_archive(
    name = "bazel_skylib",
    sha256 = "74d544d96f4a5bb630d465ca8bbcfe231e3594e5aae57e1edbf17a6eb3ca2506",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_stardoc",
    sha256 = "05fb57bb4ad68a360470420a3b6f5317e4f722839abc5b17ec4ef8ed465aaa47",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/stardoc/releases/download/0.5.2/stardoc-0.5.2.tar.gz",
        "https://github.com/bazelbuild/stardoc/releases/download/0.5.2/stardoc-0.5.2.tar.gz",
    ],
)

http_archive(
    name = "aspect_bazel_lib",
    sha256 = "44f4f6d1ea1fc5a79ed6ca83f875038fee0a0c47db4f9c9beed097e56f8fad03",
    strip_prefix = "bazel-lib-1.34.0",
    url = "https://github.com/aspect-build/bazel-lib/releases/download/v1.34.0/bazel-lib-v1.34.0.tar.gz",
)

load("//mirror:repositories.bzl", "rules_mirror_dependencies")

# Fetch dependencies which users need as well
rules_mirror_dependencies()

# For running our own unit tests
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

load("@aspect_bazel_lib//lib:repositories.bzl", "aspect_bazel_lib_dependencies")

aspect_bazel_lib_dependencies()

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
# Setup for com_adobe_rules_gitops
load("@com_adobe_rules_gitops//gitops:deps.bzl", "rules_gitops_dependencies")

rules_gitops_dependencies()

load("@com_adobe_rules_gitops//gitops:repositories.bzl", "rules_gitops_repositories")

rules_gitops_repositories()

load("@com_adobe_rules_gitops//skylib:k8s.bzl", "kubeconfig")

kubeconfig(
    name = "k8s_test",
    cluster = "kind",
    use_host_config = True,
)

############################################
# Setup for rules_gitops
load("@rules_gitops//gitops:deps.bzl", rg_dependencies = "rules_gitops_dependencies")

rg_dependencies()

load("@rules_gitops//gitops:repositories.bzl", rg_repositories = "rules_gitops_repositories")

rg_repositories()

load("@rules_gitops//skylib:k8s.bzl", rg_kubeconfig = "kubeconfig")

rg_kubeconfig(
    name = "rg_remote_test",
    cluster = "kind",
    use_host_config = True,
)
