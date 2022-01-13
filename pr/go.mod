module example.com/pr

go 1.17

require github.com/distribution/distribution/v3 v3.0.0-20220112110024-bb1fb6144562

require github.com/opencontainers/go-digest v1.0.0 // indirect

replace github.com/distribution/distribution/v3 => github.com/paulcacheux/distribution/v3 v3.0.0-20220113210606-89622d99a13e
