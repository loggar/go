# asdf

## go

Install Go Plugin for ASDF

```
asdf plugin-add go https://github.com/kennyp/asdf-golang.git
```

Install Go

```
asdf install go 1.20
```

Set Go Version in a local project

```
asdf local go 1.20
```

Alternatively, you can set the global version like this:

```
asdf global go 1.20
```

## local file

`.go-version`

```
1.20
```

or `.tool-versions`

```
go 1.20
```

## environment variables

```
export GOPATH=$HOME/ws-loggar/go
export GOROOT=$(asdf where golang)/go
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

## vscode

`settings.json`

```
{
    "go.goroot": "/Users/charly.lee/.asdf/installs/golang/1.20/go",
    "go.toolsGopath": "/Users/charly.lee/ws-loggar/go"
}
```
