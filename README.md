# Utility for setting setting AWS credentials

## Installing

1. Check you have Go installed
2. Open repo and run `go build -ldflags="-s -w"`
3. Move newly created `set-aws` binary with `mv set-aws /usr/local/bin`
4. Set an alias to run the script output in the current shell. This can be achieved by adding the following line to your shell config:

```
alias set-aws='/usr/local/bin/set-aws | source /dev/stdin'
```