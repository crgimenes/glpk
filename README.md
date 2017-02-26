# lpk

This is a small application to help you find the exact path of the packages or project in your GOPATH.

## Install

```
go install github.com/crgimenes/lpk
```

## Example of use

### Displays full path to the project

```
lpk packageName
```

### Changes to the project directory automatically

```
cd $(lpk packageName)
```

### Tips

Set up an alias for the cd and jump straight to your project directory.

```
alias gofn="cd $(lpk gofn)"
```

