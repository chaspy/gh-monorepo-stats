# gh-monorepo-stats

gh extension to output language-specific statistics for services in monorepo

## Installation

```sh
gh extension install chaspy/gh-monorepo-stats
```

To upgrade,

```sh
gh extension upgrade chaspy/gh-monorepo-stats
```

## Usage

```sh
gh monorepo-stats
```

Output example:

```sh
backend1, Gemfile, Ruby, 3121
backend2, go.mod, Go, 2189
frontend, yarn.lock, TypeScript, 10876
```

## Environment Variables

| Name          | Description                                                                                                                           |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `GH_REPO`     | The repository to query. Defaults to the current repository.                                                                          |
| `IGNORE_PATH` | List of path for ignoring to output LOC. it should be comma separated like IGNORE_PATH="app1/generated/path.go,app2/generated.path.go |
