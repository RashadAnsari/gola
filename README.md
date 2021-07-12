# Gola

[![Build Status][1]][2]

Gola is a Golang tool for automated scripting purpose

## How To Install

You can install `gola` using below command.

```bash
OS=darwin ARCH=amd64 VERSION=v0.2.2; wget https://github.com/RashadAnsari/gola/releases/download/$VERSION/gola-$VERSION-$OS-$ARCH && chmod +x gola-$VERSION-$OS-$ARCH && mv gola-$VERSION-$OS-$ARCH gola && mv gola /usr/local/bin
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
