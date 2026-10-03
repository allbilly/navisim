#!/bin/bash
# Build the sample's legacy code object from the checked-in ReLU object.
set -euo pipefail
cd -- "$(dirname -- "$0")"
ESC="${ESC:-esc}"
command -v "$ESC" >/dev/null
command -v go >/dev/null
artifact_dir="$(mktemp -d)"
trap 'rm -rf -- "$artifact_dir"' EXIT

python3 - "$artifact_dir/kernels.hsaco" <<'PY'
from pathlib import Path
import hashlib
import struct
import sys

source = Path("../relu/kernels.hsaco").read_bytes()
expected_hash = "774ddff0925bb24c8939f7a8bcff89aa2ea200e3e15a6955978f157be15f39dc"
if hashlib.sha256(source).hexdigest() != expected_hash:
    raise SystemExit("ReLU code object changed; inspect its ISA before updating the patch")

data = bytearray(source)
# The ReLU body starts at file offset 0x1100. Replace its two v_max_f32
# instructions with v_add_f32_e32 v2, 1.0, v2 and v_mov_b32_e32 v2, v2.
# Retaining the second max would clamp negative results instead of adding.
struct.pack_into("<II", data, 0x1180, 0x060404F2, 0x7E040302)
Path(sys.argv[1]).write_bytes(data)
PY

(
    cd -- "$artifact_dir"
    "$ESC" -o esc.go -pkg vectoradd -private -modtime 0 kernels.hsaco
)
test -s "$artifact_dir/esc.go"
mv -- "$artifact_dir/kernels.hsaco" kernels.hsaco
mv -- "$artifact_dir/esc.go" esc.go
