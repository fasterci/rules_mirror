"""Declare runtime dependencies

These are needed for local dev, and users must install them as well.
See https://docs.bazel.build/versions/main/skylark/deploying.html#dependencies
"""

load("@bazel_tools//tools/build_defs/repo:http.bzl", _http_archive = "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")

def http_archive(name, **kwargs):
    maybe(_http_archive, name = name, **kwargs)

def rules_mirror_dependencies():
    # The minimal version of bazel_skylib we require
    http_archive(
        name = "bazel_skylib",
        sha256 = "f7be3474d42aae265405a592bb7da8e171919d74c16f082a5457840f06054728",
        urls = [
            "https://github.com/bazelbuild/bazel-skylib/releases/download/1.2.1/bazel-skylib-1.2.1.tar.gz",
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.2.1/bazel-skylib-1.2.1.tar.gz",
        ],
    )

    http_archive(
        name = "rules_oci",
        sha256 = "176e601d21d1151efd88b6b027a24e782493c5d623d8c6211c7767f306d655c8",
        strip_prefix = "rules_oci-1.2.0",
        url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.2.0/rules_oci-v1.2.0.tar.gz",
    )

    http_archive(
        name = "rules_gitops",
        sha256 = "37be3b22375a91b0839385364882938a972564c5510b6c6ca92f6bfc0221fe79",
        strip_prefix = "rules_gitops-0.31.6",
        urls = ["https://github.com/fasterci/rules_gitops/releases/download/v0.31.6/rules_gitops-v0.31.6.tar.gz"],
    )
