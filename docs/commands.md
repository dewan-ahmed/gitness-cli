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
./gitness-cli project ls devdays
```

Output:

```
Project ID: devdays
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