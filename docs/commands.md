# Commands

## Global Options

These options can be used with any command:

- `--token value`: Gitness personal access token (can also be set using `$GITNESS_TOKEN`)
- `--url value`: Gitness server URL (default: "http://localhost:3000/") (can also be set using `$GITNESS_URL`)
- `--help, -h`: Show help message and exit

## start

Starts a Gitness Docker container running on port 3000.

```bash
./gitness-cli start
```

**Description**

This command starts a Docker container for Gitness, running it on port 3000. The container is set to always restart to ensure high availability, and it uses persistent storage to keep your data safe across restarts.

## project

Manage projects.

### ls

List details of a specific project.

```bash
./gitness-cli project ls <projectName>
```

**Example**

```bash
./gitness-cli project ls myProject
```

Output:

```
Project ID: myProject
Project Description: Gitness is awesome
Project Visibility: private
Project Created: 2024-05-22 15:12:38
```

### create

Create a new Gitness project.

```bash
./gitness-cli project create <projectName>
```

**Example**

```bash
./gitness-cli project create myProject
```

### delete

Delete an existing Gitness project including all the resources (repositories, pipelines, secrets, etc.) within it.

```bash
./gitness-cli project delete <projectName>
```

**Example**

```bash
./gitness-cli project delete myProject
```

## repo

### import 

Import a repository from another SCM to a Gitness project.

```bash
./gitness-cli repo import --project-id <projectName> --uid <repoName> <sourceRepo>
```

**Example**

```bash
./gitness-cli repo import --project-id myProject --uid podinfo harness-community/podinfo
```

**Output:**

```bash
imported podinfo
```

## pipeline

### list

List all pipelines under a specific project/repo.

```bash
./gitness-cli pipeline list --repo-ref <projectName>/<repoName>
```

**Example**

```bash
./gitness-cli pipeline list --repo-ref myProject/myRepo
```

**Output:**

```
.harness/build-deploy-pipeline.yaml
.harness/hello-pipeline.yaml
.harness/volume-pipeline.yaml
.harness/webhook-pipeline.yaml
```

### create

Create a new Gitness pipeline under a specific project/repo.

```bash
./gitness-cli pipeline create --repo-ref <projectName>/<repoName> <pipelineName>
```

**Example**

```bash
./gitness-cli pipeline create --repo-ref myProject/myRepo   --config-path .harness/test.yml --default-branch master test
```

The above command creates a new Gitness pipeline under the project `myProject` and repository `myRepo` with the name `test` and the configuration file `.harness/test.yml` and sets the default branch to `master`.

### delete

Delete an existing Gitness pipeline under a specific project/repo.

```bash
./gitness-cli pipeline delete --repo-ref <projectName>/<repoName> --pipeline-id <pipelineName>
```

**Example**

```bash
./gitness-cli pipeline delete --repo-ref myProject/myRepo --pipeline-id myPipeline
```

### executions

List all the executions of a specific pipeline under a specific project/repo.

```bash
./gitness-cli pipeline executions --repo-ref <projectName>/<repoName> --pipeline-id <pipelineName>
```

**Example**

```bash
./gitness-cli pipeline executions --repo-ref myProject/myRepo --pipeline-id myPipeline
```

Output:

```bash
3
2
1
```

### trigger

List all the triggers of a specific pipeline under a specific project/repo.

```bash
./gitness-cli pipeline trigger --repo-ref <projectName>/<repoName> --pipeline-id <pipelineName>
```

**Example**

```bash
./gitness-cli pipeline trigger --repo-ref myProject/myRepo --pipeline-id myPipeline
```

Output:

```bash
default
trigger_on_pr
```

## help, h

Shows a list of commands or help for one command.

```bash
./gitness-cli help
```

**Example**

To get help for a specific command:

```bash
./gitness-cli help start
```