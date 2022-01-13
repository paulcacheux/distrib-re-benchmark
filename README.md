# Benchmark repo

for PR: https://github.com/distribution/distribution/pull/3566

## run benchmarks

To run the benchmarks, please run:

```
go test -bench=. -benchmem
```

in each of the subdirectories:
- `previous` for the previous state
- `pr` for the state from the PR branch

Runs are also made in this [Github Actions workflow](https://github.com/paulcacheux/distrib-re-benchmark/actions/workflows/benchmark.yaml)
