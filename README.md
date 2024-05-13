# gitness-cli
A CLI to work with Gitness

## Build

```
go build -o gitness-cli ./gitness
```

## Run

Set token
```
export GITNESS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5...
```

Example command
```
./gitness-cli pipelines list --repo-ref project/repo
```
