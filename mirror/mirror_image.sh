#!/bin/sh
set -eu
{mirror_tool} -from {src_image} -digest {digest} -to {dst_image}
