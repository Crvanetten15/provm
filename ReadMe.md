# ProVM

_created and mantained by Connor Van Etten_

> Progress ABL Version Manager

<br />

#### Motive

Within my own company we utelize several different DLC's of OpenEdge Progress ABL. Unfortunately after searching I was unable to find a easy way to manage each of these versions from within my terminal. Each version had it's own bin with its own commands, that would only work for that version. Now this is completely good if you are always staying within one distribution of Progress ABL, but as my job requires much more variability between versions I decided to create my own version manager.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Conclusion](#conclusion)
4. [Thanks](#thanks)

## Installation

This CLI is built using GoLang, therefore any users of this application will need to have it installed on there machine before they will have access to _ProVM_. Navigate to the [GoLang's Official Installation](https://go.dev/doc/install) to make sure you have the required installs before proceeding.

If you believe you already have Go installed on your machine you can run the command `go version` to check that installation has bee successful.

Now to install _ProVM_ please fork my GitHub Repository and open it locally on your machine. Navigate to the project folder and run the commands below.

```bash
go build ; # creating the go exe
go install # Installing your GO CLI to your Go Bin Folder
```

By default your Go folder should be added to your Environment Path by default, but if for some reason this doesnt work, add this go folder to your path.

To test if all has worked properly run the command :

```bash
provm -v
```

## Usage

ProVM is mean to just be a version controller, and will have some limited function due to this. It will have the ability to set a constant global version of Progress ABL that is able top be target simply by running `provm`. THis will allow a developer to skip using multiple ProEnv's to run Progress commands.

## Conclusion

## Thanks
