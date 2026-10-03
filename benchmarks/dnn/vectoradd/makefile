.PHONY: all regenerate

ESC ?= esc

all: esc.go

esc.go: kernels.hsaco
	ESC="$(ESC)" ./patch_from_relu.sh

kernels.hsaco: ../relu/kernels.hsaco patch_from_relu.sh
	ESC="$(ESC)" ./patch_from_relu.sh

# Recreate both artifacts without requiring a GPU compiler.
regenerate:
	ESC="$(ESC)" ./patch_from_relu.sh
