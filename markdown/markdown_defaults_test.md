<!-- Documentation inside this block generated by gh.tarampamp.am/urfave-cli-docs/markdown; DO NOT EDIT -->

## CLI interface - docker

A self-sufficient runtime for containers.

Docker usage goes here.

> myapp does awesome things.

Usage:

```bash
$ ./app [GLOBAL FLAGS] [COMMAND] [COMMAND FLAGS] <any-arguments>
```

Global flags:

| Name                     | Description                                         | Default value | Environment variables |
|--------------------------|-----------------------------------------------------|:-------------:|:---------------------:|
| `--context="…"` (`-c`)   | Name of the context to use to connect to the daemon |               |   `DOCKER_CONTEXT`    |
| `--debug` (`-D`)         | Enable debug mode                                   |    `false`    |        *none*         |
| `--log-level="…"` (`-l`) | Set the logging level                               |      `2`      |        *none*         |

### `images` command

Manage images.

Usage:

```bash
$ ./app [GLOBAL FLAGS] images [COMMAND FLAGS] <any-args>
```

The following flags are supported:

| Name                | Description | Default value | Environment variables |
|---------------------|-------------|:-------------:|:---------------------:|
| `--string-flag="…"` | string flag |   `default`   | `STRING_FLAG_ENV_VAR` |

### `images ls` subcommand

List images.

Usage:

```bash
$ ./app [GLOBAL FLAGS] images ls [COMMAND FLAGS] [ARGUMENTS...]
```

The following flags are supported:

| Name                  | Description                                         | Default value | Environment variables |
|-----------------------|-----------------------------------------------------|:-------------:|:---------------------:|
| `--all` (`-a`)        | Show all images (default hides intermediate images) |    `true`     |        *none*         |
| `--filter="…"` (`-f`) | Filter output based on conditions provided          |               |    `DOCKER_FILTER`    |

### `images rm` subcommand

Remove one or more images.

Usage:

```bash
$ ./app [GLOBAL FLAGS] images rm [COMMAND FLAGS] IMAGE [IMAGE...]
```

The following flags are supported:

| Name             | Description                    | Default value | Environment variables |
|------------------|--------------------------------|:-------------:|:---------------------:|
| `--force` (`-f`) | Force removal of the image     |    `false`    |        *none*         |
| `--no-prune`     | Do not delete untagged parents |    `false`    |        *none*         |

### `images rm test` subcommand (aliases: `t`, `tst`)

Just for a test.

Usage:

```bash
$ ./app [GLOBAL FLAGS] images rm test [ARGUMENTS...]
```

<!-- End of automatically generated block -->
