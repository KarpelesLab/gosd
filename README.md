[![GoDoc](https://godoc.org/github.com/KarpelesLab/gosd?status.svg)](https://godoc.org/github.com/KarpelesLab/gosd)

# gosd

Notify & stuff for go.

## Usage

```go
sd := gosd.New()

// once your daemon is ready:
sd.Ready()

// on shutdown
sd.Notify(gosd.StoppingState)
```
