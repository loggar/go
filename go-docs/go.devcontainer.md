# vscode devcontainer

## issues

###

`go.mod`

packages.Load error: err: exit status 1: stderr: go: could not create module cache: `mkdir /Users`: permission deniedgo list

> This is because of the `./vscode/settings.json` which defines the go env variables for the host. If the volume is mounted for the workspace, then choose one vscode environment (host or dev-container)
