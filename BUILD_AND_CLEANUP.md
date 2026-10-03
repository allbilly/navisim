# Build and cleanup notes

## What is using disk space

The largest generated files are root-level `isa_*.debug` files. The current set has 256 files totaling about 3.91 GiB: 64 non-empty traces for GPU 1 and 192 empty files for GPUs 2–4. NaviSim opens these per-CU files when ISA debugging is enabled (`-debug-isa`). They are execution traces, not a build cache. Removing them frees space but discards the captured register-state history.

There are also sample executables under `samples/` (about 147 MiB total). These are compiled Go binaries and can be rebuilt from their sample packages. `samples/spmv/spmv` was untracked in Git but identified as an ELF executable; it is treated as a generated binary here.

`localmem/cache/` is source code from a local Akita memory checkout, not a cache directory to clean. The optional `localmem/` directory is ignored by Git; leave it intact. The root `go.mod` uses `gitlab.com/akita/mem v1.12.0` without a local replacement, so normal builds use the downloaded module. Go's compiler cache is stored outside this repository and is not responsible for the large repository size. The small `metrics.csv` files are run results and are not included in the cleanup commands below.

## Build

Compile all Go packages (this checks/builds packages but does not leave a separate executable for each sample):

```sh
go build ./...
```

Build one sample executable:

```sh
(cd samples/atax && go build -o atax .)
```

Rebuild the sample executables removed by the cleanup below:

```sh
for sample in atax bicg fir kmeans matrixmultiplication matrixtranspose maxpooling relu simpleconvolution spmv stencil2d; do
  (cd "samples/$sample" && go build -o "$sample" .)
done
```

To recreate ISA traces, rerun the desired workload with `-debug-isa`; outputs are written to the process's current working directory. For the documented small vector-add run:

```sh
(cd samples/vectoradd && go run . -timing -verify -debug-isa -report-all -length 4096)
```

This regenerates traces for that run, not necessarily the same large root-level traces. The command and workload that produced the original 3.91 GiB trace set are not recorded, so reproducing those exact contents requires recovering the original workload and flags.

## Cleanup

From the repository root, inspect the exact root trace files first:

```sh
find . -maxdepth 1 -type f -name 'isa_*.debug' -printf '%s %p\n'
```

Remove only those root-level traces with:

```sh
find . -maxdepth 1 -type f -name 'isa_*.debug' -delete
```

Remove the named sample executables with:

```sh
for sample in atax bicg fir kmeans matrixmultiplication matrixtranspose maxpooling relu simpleconvolution spmv stencil2d; do
  find "samples/$sample" -maxdepth 1 -type f -name "$sample" -delete
done
```

These commands do not delete source files, kernel code objects, sample `metrics.csv` results, or trace files inside sample directories. Rebuild executables with the loop above. Regenerate traces by rerunning the chosen workload with `-debug-isa`.
