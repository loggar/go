# Go dev env

## local

`direnv`

```
goenv local 1.18.5
```

`.go-version`

```
1.18.5
```

`direnv`

`.envrc`

```
export GOPATH=$HOME/ws-loggar/go
export GOROOT=/Users/charly.lee/.goenv/versions/$(goenv version-name)
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

```
direnv allow
```
