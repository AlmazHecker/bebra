## About Bebra

This is a tool to automate reverse engineering of APK files. It can:

1. Decompile APK
2. Compile
3. Extract APK from XAPK
4. Build XAPK from dir with given APKs
5. Create keystore.jks
6. Multi sign of APK files

### Built With

Go and Cobra

### Prerequisites

You need to install APKTool and Android SDK Tools

## Installation Guide

There are 2 ways to install bebra.

#### Simple installation

1. Download from bin directory build that supports your PC
2. Then you can run it in CMD or Bash:

```sh
./bebra.
```

#### Advanced installation

You have to install Golang.

1. Clone the repo
2. Compile to binaries
   ```sh
   ./scripts/build.sh
   ```
3. Choose binary that's right for your PC & run it:

```bash
 ./bebra
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
