# Vanity

> [!WARNING]
> This was mostly made to suit my needs, some things may or may not work for you and I do not spend a lot of time maintaining it, use at your own risk.

A vanity URL service for Go packages (modules, projects and executables)

## Installation

You can install Vanity with Go:

```sh
go install go.trulyao.dev/vanity@latest
```

## Usage

Run the following command to create a config file:
```sh
vanity --init
```

Optionally, you can use `--config=[path]` flag to customize your config file location and name. You can now run Vanity by using the following command:

```sh
vanity --config=path/to/config.json
```

## Docker

Docker is the recommended way to use Vanity but I haven't taken time to setup the workflow yet (there aren't even versions yet at this point) but you can use the unofficially supported image or build from source.

```sh
docker pull trulyao/vanity:latest
```
