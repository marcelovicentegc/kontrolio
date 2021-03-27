![Build and test](https://github.com/marcelovicentegc/kontrolio-cli/workflows/Go/badge.svg)

<p align="center">
  <img alt="kontrolio logo" src="./assets/logo.png" height="300" />
  <h3 align="center">kontrolio-cli</h3>
  <p align="center">Kontrolio's CLI time clock, clock card machine, punch clock, or time recorder.</p>
</p>

## Installation

- You can:
  - Install with homebrew if you're a Mac user :beer::
    - `brew install ktrlio/tools/kontrolio`
  - Install it with your favorite package manager:
    - `yarn global add @kontrolio/cli`
    - `npm i -g @kontrolio/cli`
  - [Download the binaries for Linux (64 and 32 bit), macOS or Windows here](https://github.com/marcelovicentegc/kontrolio-cli/releases/latest)

See [Troubleshooting](#troubleshooting) if you have any issues on installation.

## Usage

```bash
$ kontrolio
```

```plain
NAME:
   kontrolio - Your cli time clock, clock card machine, punch clock or time recorder

USAGE:
   kontrolio [global options] command [command options] [arguments...]

VERSION:
   0.x

COMMANDS:
   auth, a                Authenticate on Kontrolio
   config, c              Configure Kontrolio
   logs, l                Navigate through all your records
   punch, p               Punch your clock
   status, s              Check how many hours have you worked today
   sync                   Sync offline and online records
   help, h                Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Troubleshooting

In case you receive an `EACESS` error while trying to execute `kontrolio` after installing it with `npm` or `yarn`, change the file permissions with `chmod 755 <path_to_binaries>`.

## About

### Offline mode

Kontrolio works offline by default. If you want to save your data on Kontrolio's platform, check the [online mode](#online-mode)

### Online mode

In order to register your data remotely on Kontrolio's database, you need to create an account at [kontrolio.com](https://kontrolio.com) and authenticate with:

```bash
$ kontrolio auth
```

## Development

Make sure you have a `.kontrolio.yaml` file under your home directory (`/home/marcelo` on Linux, `/Users/Marcelo` on macOs, `C:\Users\Marcelo` on Windows) with `dev` set to `true`, like this:

```yaml
dev: true
```
