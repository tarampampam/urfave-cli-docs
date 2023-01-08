# urfave CLI docs

![go_version][badge_go_version]
[![tests][badge_tests]][actions]
[![coverage][badge_coverage]][coverage]
[![docs][badge_docs]][docs]

<!-- Documentation inside this block generated automatically; DO NOT EDIT -->

CLI interface - myapp
---------------------

does awesome things.

> myapp does awesome things.

Usage:

```bash
$ ./app [GLOBAL FLAGS] command [SUBCOMMAND] [COMMAND FLAGS] [ARGUMENTS...]
```

Global flags:

| Name                        | Description        | Default value | Environment variables |
|-----------------------------|--------------------|:-------------:|-----------------------|
| `--app-string-flag` (`-a`\) | app string flag    |   `default`   | `STRING_FLAG_ENV_VAR` |
| `--app-int-flag`            | app int flag       |     `123`     | `INT_FLAG_ENV_VAR`    |
| `--l`                       | short app int flag |     `321`     |                       |

### `cmd` command

does awesome things.

The following flags are supported:

| Name            | Description | Default value | Environment variables |
|-----------------|-------------|:-------------:|-----------------------|
| `--string-flag` | string flag |   `default`   | `STRING_FLAG_ENV_VAR` |
| `--int-flag`    | int flag    |     `123`     | `INT_FLAG_ENV_VAR`    |

#### **`cmd subcmd`** sub-command

does awesome things.

<!-- End of automatically generated block -->

[badge_tests]:https://img.shields.io/github/actions/workflow/status/tarampampam/urfave-cli-docs/tests.yml?branch=master
[badge_coverage]:https://img.shields.io/codecov/c/github/tarampampam/urfave-cli-docs/master.svg?maxAge=30
[badge_docs]:https://pkg.go.dev/badge/mod/github.com/tarampampam/urfave-cli-docs
[badge_go_version]:https://img.shields.io/badge/go%20version-%3E=1.16-61CFDD.svg
[actions]:https://github.com/tarampampam/urfave-cli-docs/actions
[coverage]:https://codecov.io/gh/tarampampam/urfave-cli-docs
[docs]:https://pkg.go.dev/github.com/tarampampam/urfave-cli-docs
