<div align="center">
  <h1>docker-tui</h1>

### NOTE: docker-tui is not done yet and isn't functional for use.
  
docker-tui is a terminal interface for easily managing docker containers running on your machine.

![License](https://img.shields.io/github/license/marcelpkg/docker-tui)
![GitHub branch check runs](https://img.shields.io/github/check-runs/marcelpkg/docker-tui/main?label=tests)
![GitHub Release](https://img.shields.io/github/v/release/marcelpkg/docker-tui)

</div>

## Installation

### Requirements

- [Docker Engine](https://docs.docker.com/engine/install/)
- [Git](https://git-scm.com/downloads/)
- [Go](https://go.dev/doc/install)

### Linux

1. Clone the repository

```bash
$ git clone https://github.com/marcelpkg/docker-tui.git
$ cd docker-tui
```

2. Build the binary & install

```bash
$ sudo GOBIN=/usr/local/bin/ go install
```
