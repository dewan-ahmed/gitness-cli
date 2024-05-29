# gitness-cli
A CLI to work with Gitness

## Install on Mac

Intel CPU

```sh
curl -L https://github.com/dewan-ahmed/gitness-cli/releases/latest/download/gitness-cli-darwin-amd64.tar.gz | tar zx
```

Apple silicon CPU

```sh
curl -L https://github.com/dewan-ahmed/gitness-cli/releases/latest/download/gitness-cli-darwin-arm64.tar.gz | tar zx
```

Verify

```sh
./gitness-cli --help
```

## Build

```
go build -o gitness-cli ./gitness
```

## Run

Set token
```
export GITNESS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5...
```

Example commands
```
./gitness-cli pipelines list --repo-ref project/repo
./gitness-cli pipeline list --repo-ref project/repo --pipeline-id example
./gitness-cli pipeline executions --repo-ref project/repo --pipeline-id example --format "Number: {{ .Number }} Event: {{ .Event }}"
```
