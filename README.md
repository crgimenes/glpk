# gothere

This is a small application to help you find the exact path of the packages or project in your GOPATH.

## Install

```
go install github.com/crgimenes/gothere
```

## Example of use

### Displays full path to the project

```
gothere -name packageName
```

### Changes to the project directory automatically

```
cd $(gothere -name packageName)
```
