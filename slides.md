---
# vim: set ft=markdown:
author: Christoph Lipautz
title: Rubyists, Lets Go
subtitle: Viennarb 2023
theme: Copenhagen
colortheme: beaver
date: \today
fontsize: 14pt
keywords: ['ruby', 'golang']
---

# Slides & Resources

https://github.com/unused/rubyists-lets-go/

# Go Programming Language

\vspace{2.5cm}
\center
![](figures/goopher.png)

# What is it good for?

`script < command < application`

# Are we on the same page?

- Statically Typed
<!-- nono, not like rbs or sorbet -->

- Compiled to Cross-Platform

- Nice Community, Much Packages

- Testing and Documentation is a Good Practice

# Learn Go

::: columns

:::: column

- Go by Example

- Tour of Go

- Effective Go

::::

:::: column

![](figures/go-by-example.png)

::::

:::

# Starter

Much improved `go-mod`, but no `bundler` yet.

```sh
# creative project names
# including "go" highly welcome
$ go mod init github.com/unused/gorgonzola
```

# In the Wild

Hugo, Kubernetes, Docker, ngrok, influxDB, fzf, esbuild, anycable

<!-- TODO: add logos -->

# Defer Go

```go
func write(filename, body string) {
    f, _ := os.Create(filename)
    f.WriteString(body)
    defer f.Close()
}
```

# Defer Go and Ruby

```go
func write(filename, body string) {
    f, _ := os.Create(filename)
    defer f.Close()
    f.WriteString(body)
}
```

\vspace{1cm}

```ruby
def write(filename, body)
  file = File.open(filename, 'w') do |file|
  file.write body
  file.close
end
```

# Defer Ruby Block

```ruby
def write(filename, body)
  file = File.open(filename, 'w') do |file|
  file.write body
  file.close
end
```

\vspace{1cm}

```ruby
def write(filename, body)
  File.open(filename, 'w') do |file|
    file.write body
  end
end
```

# Defer Go in Good Style

```go
func write(filename, body string) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    return f.WriteString(body)
}
```

# Errors

```go
err := json.Unmarshal(byt, &dat)
if err != nil {
    return nil, err
}
return dat, nil

// or


if err := json.Unmarshal(byt, &dat); err != nil {
    return nil, err
}
// continue
```

<!--
Exceptions, something like that exists, but is discouraged. Use safe error
handling instead.
-->

# Errors (for this example)

```go
// for this example...
func fromJson(res []byte, &dat Response) (*Response, error) {
    return json.Unmarshal(res, &dat)
}
```

# Conventions in Language Design

`myFunc` vs `MyFunc`

People write things very similar (like with Ruby)

# Rubocop Ships with Language Tooling

```sh
$ gofmt main.go
```

# Embrace Spaghetti

Ruby is Developer Happiness

Go is happy Developer

# Summary

\small https://github.com/unused/rubyists-lets-go

\spacer

\small Learn: https://gobyexample.com/ https://go.dev/tour/
[https://go.dev/doc/effective\_go](https://go.dev/doc/effective_go)

