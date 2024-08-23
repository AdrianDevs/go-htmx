<h1>GoHTMX</h1>

<h3>A project to explore Go and HTMX</h3>

<h2>Table of Contents</h2>

[TOC]

# Tech

**Frontend**

- HTMX
- surreal.js

**Backend**

- Golang
- Chi
- Templ

**Auth**

- SuperTokens

**Storage**

- PostgreSQL
- Sqlc and pgx
- Goose for migrations
- MinIO

**Hosting**

- Docker-Compose
- Fly.io

**Testing**

- Unit tests with Go
- E2E tests with Playwright



# Run

```sh
$ go mod tidy
$ go install github.com/a-h/templ/cmd/templ@latest
$ go install github.com/air-verse/air@latest
$ templ generate
$ go run home.go
$ air
```



# Setup


## Templ

### Setup

Install Templ

```sh
$ go install github.com/a-h/templ/cmd/templ@latest
```

Make sure Go is added to your PATH. In your `.zshrc` file, add the line.

```sh
export PATH="$PATH:$HOME/go/bin"
```

Add the VS Code extension: https://marketplace.visualstudio.com/items?itemName=a-h.templ.

### Run

Templ files end with a  `.templ` extension. To generate templates from them:

Navigate to the root of your project.

To generate all files, run

```sh
$ templ generate
```

To generate a specific file, run

```sh
$ templ generate <filename>
```

### Watch

Automatically watch for go and templ files and rebuild templates, rerun go and refresh the browser.

```sh
$ templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
```

## WGO

Live reload for Go apps. Watch arbitrary files and respond with arbitrary commands. It supports running multiple invocations in parallel. 

**WGO can be used as an alternative to Templ watch or Air.**

### Setup

```sh
$ go install github.com/bokwoon95/wgo@latest
```

Make sure Go is added to your PATH. In your `.zshrc` file, add the line.

```sh
export PATH="$PATH:$HOME/go/bin"
```

### Run

Automatically watch for go and templ files and rebuild templates, rerun go. Note this will not refresh the browser.

**Using WGO** - this will not refresh the browser.

```sh
$ wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run home.go
```

## Air (preferred)

Live reloading for Go apps with support for automatic browser refreshing. 

**Air can be used as an alternative to Templ watch or WGO.**

Site: https://github.com/air-verse/air

**Reasons for preference**

- Reloads browser for both `.go` and `.temp` file changes
- Is browser agnostic

### Setup

```sh
$ go install github.com/air-verse/air@latest
```

Make sure Go is added to your PATH. In your `.zshrc` file, add the line.

```sh
export PATH="$PATH:$HOME/go/bin"
```

Initialise the `.air.toml` configuration file to the current directory with the default settings running the following command.

```sh
$ air init
```

Add support for automatic browser refreshing by changing the `.air.toml` file contents to the following.

```toml
root = "."
tmp_dir = "tmp"

[build]
  bin = "./tmp/main"
  cmd = "templ generate && go build -o ./tmp/main ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor"]
  exclude_file = []
  exclude_regex = [".*_templ.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "templ", "html"]
  kill_delay = "0s"
  log = "build-errors.log"
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
#  main_only = false
  time = false

[misc]
  clean_on_exit = false

[proxy]
  enabled = true
  proxy_port = 3000
  app_port = 8080
```

### Run

Automatically watch for go and templ file changes, rebuild templates, rerun go and refresh the browser.

```sh
$ air
```

# Code Architecture

Domain - >Handler -> Service -> Database(db_instance)

​                                    -> View


# Miscellaneous

## Go commands

- `$ go install` is used to install a binary, not a package. It will download the binary and put it under GOPATH/bin, which can be executed in the terminal without having to add this to `go.mod`. This is sort of equivalent to global packages in npm. Except in Go, packages cannot be global, and instead, binaries are when placed in  `GOPATH/bin`.
- `$ go get` downloads a package.
- `go mod tidy` is the equivalent of `npm install`. It will scan your project, fetch any dependencies on your code base, and update the `go.mod` file if necessary, such as when there are dependencies in the code base that are not in `go.mod`.
- `$ go mod download` only downloads the dependencies from `go.mod` without modifying the `go.mod` file.

## Creating Diagrams

PlantUML and C4 Diagrams
