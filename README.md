# Gola

[![Build Status][1]][2]
[![Go Report][3]][4]

Gola is a Golang tool for automated scripting purpose

## How To Install

You can install `gola` using below command.

```bash
# TODO: Add install command
```

## Example

### Configuration

```yaml
commands:
  - name: mysql
    cmd:
      path: "mysql"
      args:
        - "-u"
        - "rashad"
        - "-h"
        - "127.0.0.1"
      env:
        - "MYSQL_PWD=pass"
```

### Usage

```bash
gola run mysql
```

[1]: https://img.shields.io/drone/build/RashadAnsari/gola.svg?style=flat-square&logo=drone
[2]: https://cloud.drone.io/RashadAnsari/gola
[3]: https://goreportcard.com/badge/github.com/RashadAnsari/gola?style=flat-square
[4]: https://goreportcard.com/report/github.com/RashadAnsari/gola
