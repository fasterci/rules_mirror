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
        name = "com_adobe_rules_gitops",
        sha256 = "83124a8056b1e0f555c97adeef77ec6dff387eb3f5bc58f212b376ba06d070dd",
        strip_prefix = "rules_gitops-0.17.2",
        urls = ["https://github.com/adobe/rules_gitops/archive/refs/tags/v0.17.2.tar.gz"],
    )

    http_archive(
        name = "io_bazel_rules_docker",
        sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
        urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
    )

    http_archive(
        name = "rules_oci",
        sha256 = "647f4c6fd092dc7a86a7f79892d4b1b7f1de288bdb4829ca38f74fd430fcd2fe",
        strip_prefix = "rules_oci-1.7.6",
        url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.7.6/rules_oci-v1.7.6.tar.gz",
    )

    RULES_GITOPS_COMMIT = "aa6320b36aec9beacd0af9c3ad3bb479e41f3bce"

    http_archive(
        name = "rules_gitops",
        sha256 = "694cad4247a0d4919fa5d41d17770d9aab0d9858ef03ca5658787de70557889a",
        strip_prefix = "rules_gitops-%s" % RULES_GITOPS_COMMIT,
        urls = ["https://github.com/fasterci/rules_gitops/archive/%s.tar.gz" % RULES_GITOPS_COMMIT],
    )
