set -eu
go build .
mv gh-monorepo-stats "../../${GH_REPO}"