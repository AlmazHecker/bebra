## About Bebra

This is a tool to automate reverse engineering of APK files. It can:

1. Decompile APK
2. Compile
3. Extract APKs from XAPK
4. Create keystore.jks
5. Multi sign of APK files
6. Install multiple APK files

### Requirements

You need to install APKTool and ADB & APKSigner from Android SDK

## Installation Guide

> **NOTE:** Windows detects this file as a virus, but it's not!

1. Download bebra [here](https://github.com/AlmazHecker/bebra/releases)
2. Unzip the archive
3. Then you can run it in CMD or Bash:

```sh
# For Linux
./bebra --help

# For Windows
./bebra.exe --help
```

## After installation

run:

```
bebra init
```

This command will create _bebra.config.json_ file. You will need to set some variables.

Note: you can add bebra to PATH to make it accessible everywhere:

```bash
# For Linux
sudo mv bebra /usr/local/bin
```

Otherwise you need full path of bebra location to run it.

## Usage

Run

```bash
bebra --help
```

## HELP NEEDED!

Bebra currently depends on tools like ADB, apksigner and APKTool.
It would be great if we could implement tools like ADB and apksigner natively!
Because in order to install them we need to install whole Android SDK ):<

I also need help with the documentation for this project. XD

## Why I choose Golang ?

1. It's cross-platform
2. Compiles into native binaries(no dependencies needed from user)
3. if err != nil XD
