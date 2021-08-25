# Debug Library

Custom debug loging library

---------------------------------------
  * [Features](#features)
  * [Requirements](#requirements)
  * [Installation](#installation)
  * [Usage](#usage)
  * [License](#license)

---------------------------------------

## Features
  * Print in color based on debug level

## Requirements
  * Supports Go 1.7 or later

---------------------------------------

## Installation
```bash
go get -u github.com/bonagin/debug
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Usage
```go
import (
  "github.com/bonagin/debug"
)

// Initialize the debug library using the application name as the argument
const (
  log = debug.NewDebug("test")
)

// Make sure that "log" is not imported as this will cause a conflict
log.Alert("This is an alert")
log.Success("This is an update")
log.Warning("This is a warning")
log.Info("This is a notification")
log.Dump("This is a dump")
```
Add the config flag when running the exercutable
```bash
$ /path/to/exercutable --debug
```

## License
Go-MySQL-Driver is licensed under the [Mozilla Public License Version 2.0](https://raw.github.com/go-sql-driver/mysql/master/LICENSE)

Mozilla summarizes the license scope as follows:
> MPL: The copyleft applies to any files containing MPLed code.


That means:
  * You can **use** the **unchanged** source code both in private and commercially.
  * When distributing, you **must publish** the source code of any **changed files** licensed under the MPL 2.0 under a) the MPL 2.0 itself or b) a compatible license (e.g. GPL 3.0 or Apache License 2.0).
  * You **needn't publish** the source code of your library as long as the files licensed under the MPL 2.0 are **unchanged**.

Please read the [MPL 2.0 FAQ](https://www.mozilla.org/en-US/MPL/2.0/FAQ/) if you have further questions regarding the license.

You can read the full terms here: [LICENSE](https://github.com/bonagin/config/blob/master/LICENSE).
