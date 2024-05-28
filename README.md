# gitness-cli
A CLI to work with Gitness


example change

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
