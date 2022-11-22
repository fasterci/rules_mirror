"Public API re-exports"

load("@io_bazel_rules_docker//container:providers.bzl", "PushInfo")
load("@com_adobe_rules_gitops//skylib:push.bzl", "K8sPushInfo")

def _mirror_image_impl(ctx):
    digest_file = ctx.actions.declare_file(ctx.label.name + ".digest")
    ctx.actions.write(
        output = digest_file,
        content = ctx.attr.digest,
    )

    ctx.actions.expand_template(
        template = ctx.file._mirror_image_script,
        output = ctx.outputs.executable,
        substitutions = {
            "{mirror_tool}": ctx.executable.mirror_tool.short_path,
            "{src_image}": ctx.attr.src_image,
            "{digest}": ctx.attr.digest,
            "{dst_prefix}": ctx.attr.dst_prefix,
        },
        is_executable = True,
    )

    # keep this in sync with image mirror tool implementation
    # _, s = ctx.attr.src_image.split("/", 1)
    s = ctx.attr.src_image.split("@", 1)[0]
    src_repository = s.split(":", 1)[0]

    dst_without_hash = ctx.attr.dst_prefix + "/" + src_repository
    dst_registry, dst_repository = dst_without_hash.split("/", 1)

    return [
        DefaultInfo(
            runfiles = ctx.runfiles(files = [ctx.executable.mirror_tool]),
            files = depset([digest_file]),
            executable = ctx.outputs.executable,
        ),
        PushInfo(
            registry = dst_registry,
            repository = dst_repository,
            digest = digest_file,
        ),
        K8sPushInfo(
            image_label = ctx.label,
            legacy_image_name = ctx.attr.image_name,
            registry = dst_registry,
            repository = dst_repository,
            digestfile = digest_file,
        ),
    ]

mirror_image = rule(
    implementation = _mirror_image_impl,
    attrs = {
        "src_image": attr.string(
            mandatory = True,
            doc = "The image to mirror",
        ),
        "image_name": attr.string(
            doc = "The name that could be referred in manifests. This field is optional since the image can always be referred by its target label.",
        ),
        "digest": attr.string(
            mandatory = True,
            doc = "The digest of the image",
        ),
        "dst_prefix": attr.string(
            mandatory = True,
            doc = "The prefix of the destination image, should include the registry and repository",
        ),
        "mirror_tool": attr.label(
            default = Label("@com_fasterci_rules_mirror//cmd/mirror"),
            executable = True,
            cfg = "exec",
        ),
        "_mirror_image_script": attr.label(
            default = ":mirror_image.sh",
            allow_single_file = True,
        ),
    },
    executable = True,
    doc = """Mirror an image to a local registry. 
Implements the K8sPushInfo provider so the returned image can be injected into manifests by rules_gitops
""",
)
