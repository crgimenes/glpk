# glpk

This is a small application to help you find the exact path of the packages or project in your GOPATH.

## Install

```
go install github.com/crgimenes/glpk
```

## Example of use

### Displays full path to the project

```
glpk -name packageName
```

### Changes to the project directory automatically

```
cd $(glpk -name packageName)
```
