![Build and test](https://github.com/marcelovicentegc/kontrolio-cli/workflows/Go/badge.svg)

<p align="center">
  <img alt="kontrolio logo" src="./assets/logo.png" height="300" />
  <h3 align="center">kontrolio-cli</h3>
  <p align="center">Kontrolio's CLI time clock, clock card machine, punch clock, or time recorder.</p>
</p>

## Installation

- You can install it with your favorite package manager:
  - `yarn global add @kontrolio/cli`
  - `npm i -g @kontrolio/cli`
- [Or, you can download the binaries for Linux (64 and 32 bit) and macOS here](https://github.com/marcelovicentegc/kontrolio-cli/releases/latest)

## Usage

```bash
$ kontrolio
```

```plain
NAME:
   kontrolio - your cli time clock, clock card machine, punch clock or time recorder

USAGE:
   kontrolio [global options] command [command options] [arguments...]

VERSION:
   0.x

COMMANDS:
   logs, l    Navigate through all your records
   punch, p   Punch your clock
   status, s  Check how many hours have you worked today
   sync       Sync offline and online records
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Troubleshooting

In case you receive an `EACESS` error while trying to execute `kontrolio`, change the file permissions with `chmod 755 <path_to_binaries>`.

## About

### ✈️ Offline mode

Kontrolio works offline by default. If you want to save your data on Kontrolio's platform, check the [online mode](#-online-mode)

### 🌐 Online mode

In order to register your data remotely on Kontrolio's database, you need to create an account on [kontrolio.com](https://kontrolio.com) and set the generated API Key on the [configuration file](#-configuration).

### 🧰 Configuration

| Functionality      | Enabled by default |
| ------------------ | ------------------ |
| Saves data offline | ✔️                 |
| Saves data online  | opt-in             |

Kontrolio has a [configuration file](../.kontrolio.example.yaml) that allows you to configure it. This is optional, you don't need to create this file unless you want to customize some default behavior.

The file must be named `.kontrolio.yaml`.

Kontrolio looks for it in your home directory (`/home/marcelo` on Linux, `/Users/Marcelo` on macOs, `C:\Users\Marcelo` on Windows).

This is how `.kontrolio.yaml` should look like:

```yaml
# Required if you want to save your data on Kontrolio's database,
# thus have created an account on https://kontrolio.com.
api_key: "YOUR_API_KEY"
```

## Development

Make sure you have a `.kontrolio.yaml` file under your home directory with `dev` set to `true`, like this:

```yaml
dev: true
```
