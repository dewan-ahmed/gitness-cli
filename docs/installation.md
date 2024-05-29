# Installation

## Prerequisites

Ensure you have the following installed:

- Docker
- Go (Golang)

## Install Gitness CLI

* Clone the repository:

```bash
git clone https://github.com/dewan-ahmed/gitness-cli.git
cd gitness-cli
```

* Build the binary:

```bash
go build -o gitness-cli ./gitness
```

## Configure Gitness Token

* From `http://localhost:3000/profile`, click **+ New Token** to create a Gitness token and export the following environment variable:

```bash
export GITNESS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5... # your token
```