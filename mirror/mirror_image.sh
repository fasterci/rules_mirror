#!/bin/sh
set -e
{mirror_tool} -from {src_image} -digest {digest} -to_prefix {dst_prefix}
