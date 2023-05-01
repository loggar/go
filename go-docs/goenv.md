# goenv

Install `goenv`

```
# mac
brew install goenv
```

```
export PATH="$HOME/.goenv/bin:$PATH"
eval "$(goenv init -)"
```

Install go versions

```
# list versions
goenv install --list

# install version
goenv install <version>
```

Use go version

```
$ goenv versions
  system
  1.16.15
* 1.17.8 (set by /Users/<user>/.goenv/version)
  1.18.5

# Set the Go version for your current shell session:
$ goenv shell 1.18.5

# Set the installed Go version for a specific project, navigate to the project directory and use:
$ goenv local 1.18.5
```
