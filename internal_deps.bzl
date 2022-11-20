"""Our "development" dependencies

Users should *not* need to install these. If users see a load()
statement from these, that's a bug in our distribution.
"""

load("@bazel_tools//tools/build_defs/repo:http.bzl", _http_archive = "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")

def http_archive(name, **kwargs):
    maybe(_http_archive, name = name, **kwargs)

def rules_mirror_internal_deps():
    "Fetch deps needed for local development"
    http_archive(
        name = "io_bazel_rules_go",
        sha256 = "099a9fb96a376ccbbb7d291ed4ecbdfd42f6bc822ab77ae6f1b5cb9e914e94fa",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.35.0/rules_go-v0.35.0.zip",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.35.0/rules_go-v0.35.0.zip",
        ],
    )

    # gazelle is required by go_image_repositories
    http_archive(
        name = "bazel_gazelle",
        sha256 = "efbbba6ac1a4fd342d5122cbdfdb82aeb2cf2862e35022c752eaddffada7c3f3",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.27.0/bazel-gazelle-v0.27.0.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.27.0/bazel-gazelle-v0.27.0.tar.gz",
        ],
    )

    # Override bazel_skylib distribution to fetch sources instead
    # so that the gazelle extension is included
    # see https://github.com/bazelbuild/bazel-skylib/issues/250
    http_archive(
        name = "bazel_skylib",
        sha256 = "07b4117379dde7ab382345c3b0f5edfc6b7cff6c93756eac63da121e0bbcc5de",
        strip_prefix = "bazel-skylib-1.1.1",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/archive/1.1.1.tar.gz",
            "https://github.com/bazelbuild/bazel-skylib/archive/refs/tags/1.1.1.tar.gz",
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
        sha256 = "1cbbf62315d303c8083d5019a4657623d4f58e925fb51bdc8a41bad4a131f5c9",
        strip_prefix = "bazel-lib-1.8.1",
        url = "https://github.com/aspect-build/bazel-lib/archive/refs/tags/v1.8.1.tar.gz",
    )

    http_archive(
        name = "com_google_protobuf",
        sha256 = "9748c0d90e54ea09e5e75fb7fac16edce15d2028d4356f32211cfa3c0e956564",
        strip_prefix = "protobuf-3.11.4",
        urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.4.zip"],
    )
