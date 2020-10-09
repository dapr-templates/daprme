# daprme

[![Go Report Card](https://goreportcard.com/badge/github.com/dapr-templates/daprme)](https://goreportcard.com/report/github.com/dapr-templates/daprme)
[![Build Status](https://github.com/dapr-templates/daprme/workflows/dapr_cli/badge.svg)](https://github.com/dapr-templates/daprme/actions?workflow=dapr_cli)
[![codecov](https://codecov.io/gh/dapr-templates/daprme/branch/master/graph/badge.svg)](https://codecov.io/gh/dapr-templates/daprme)

This new app wizard CLI guides you through the creation of a new Dapr project. Including the Dapr component selection and scaffolding of a new application for either command-line ClI, or HTTP and gPRC service application types.

## Installing Dapr CLI

**MacOS**

Install the latest darwin CLI to `/usr/local/bin`

```bash
curl -fsSL https://raw.githubusercontent.com/dapr-templates/daprme/master/install/install.sh | /bin/bash
```

**Linux**

Install the latest linux CLI to `/usr/local/bin`

```bash
wget -q https://raw.githubusercontent.com/dapr-templates/daprme/master/install/install.sh -O - | /bin/bash
```

**Windows**

Install the latest windows CLI to `c:\daprme` and add this directory to User PATH environment variable.

```powershell
powershell -Command "iwr -useb https://raw.githubusercontent.com/dapr-templates/daprme/master/install/install.ps1 | iex"
```





#### From the Binary Releases

Each release of CLI includes various OSes and architectures.

1. Download the [Dapr CLI](https://github.com/dapr/cli/releases)
2. Unpack it (e.g. dapr_linux_amd64.tar.gz, dapr_windows_amd64.zip)
3. Move it to your desired location.
   * For Linux/MacOS - `/usr/local/bin`
   * For Windows, create a directory and add this to your System PATH. For example create a directory called `c:\dapr` and add this directory to your path, by editing your system environment variable.
