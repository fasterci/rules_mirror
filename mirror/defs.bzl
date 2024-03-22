"Public API"

load("@rules_gitops//gitops:provider.bzl", "GitopsPushInfo")
load("@com_adobe_rules_gitops//skylib:push.bzl", "K8sPushInfo")
load("@rules_gitops//skylib:runfile.bzl", "get_runfile_path")

def _mirror_image_impl(ctx):
    digest = ctx.attr.digest
    src_image = ctx.attr.src_image
    v = src_image.split("@", 1)
    s = v[0]
    if len(v) > 1:
        # If the image has a digest, use that.
        if digest and v[1] != digest:
            fail("digest mismatch: %s != %s" % (v[1], digest))
        digest = v[1]
    else:
        # If the image does not have a digest, use the one provided.
        src_image = s + "@" + digest

    if not digest:
        fail("digest must be provided as an attribute to mirror_image or in the src_image")

    dst_without_hash = ""
    if ctx.attr.dst:
        dst = ctx.expand_make_variables("dst", ctx.attr.dst, {})
        dst = dst.split("@", 1)[0]
        v = dst.split(":", 1)
        dst_without_hash = v[0]
        if len(v) > 1:
            fail("dst should not include a tag, only a repository")
    else:
        if not ctx.attr.dst_prefix:
            fail("either dst or dst_prefix must be defined in mirror_image")
        v = s.split(":", 1)
        src_repository = v[0]
        if len(v) > 1:
            src_repository = src_repository + "/" + v[1]
        dst_prefix = ctx.expand_make_variables("dst_prefix", ctx.attr.dst_prefix, {})
        dst_without_hash = dst_prefix.strip("/") + "/" + src_repository

    digest_file = ctx.actions.declare_file(ctx.label.name + ".digest")
    ctx.actions.write(
        output = digest_file,
        content = digest,
    )

    pusher_input = [digest_file]

    ctx.actions.expand_template(
        template = ctx.file._mirror_image_script,
        output = ctx.outputs.executable,
        substitutions = {
            "{mirror_tool}": get_runfile_path(ctx, ctx.executable.mirror_tool),
            "{src_image}": src_image,
            "{digest}": digest,
            "{dst_image}": dst_without_hash,
        },
        is_executable = True,
    )

    dst_registry, dst_repository = dst_without_hash.split("/", 1)

    runfiles = ctx.runfiles(files = pusher_input).merge(ctx.attr.mirror_tool[DefaultInfo].default_runfiles)

    return [
        DefaultInfo(
            runfiles = runfiles,
            executable = ctx.outputs.executable,
        ),
        K8sPushInfo(
            image_label = ctx.label,
            legacy_image_name = ctx.attr.image_name,
            registry = dst_registry,
            repository = dst_repository,
            digestfile = digest_file,
        ),
        GitopsPushInfo(
            image_label = ctx.label,
            repository = dst_without_hash,
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
            doc = "The name that could be referred in manifests. This field is deprecated and used only in legacy com_adobe_rules_gitops.",
        ),
        "digest": attr.string(
            mandatory = False,
            doc = "The digest of the image. If not provided, it will be extracted from the src_image.",
        ),
        "dst_prefix": attr.string(
            doc = "The prefix of the destination image, should include the registry and repository. Either dst_prefix or dst_image must be specified.",
        ),
        "dst": attr.string(
            doc = "The destination image location, should include the registry and repository. Either dst_prefix or dst_image must be specified.",
        ),
        "mirror_tool": attr.label(
            default = Label("//cmd/mirror"),
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
