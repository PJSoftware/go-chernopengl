# Abstracting Shaders

Much of the code we required had already been written; all that was required was to move it across into its own package(s) and clean up any issues.

## Problem with VSCode

I shall mention this here, because it caught me yet again.

When programming in Go with VSCode, it is strongly recommended to use the official Go extension. It does some clever things to make your life as a Go programmer easier.

One feature is that it will automatically update the `import` list to cover any packages that you reference in the code. This is usually a very handy, time-saving feature. However, in this code we are currently using:

```go
import "github.com/go-gl/gl/v4.1-core/gl"
```

However, if you use one of the `gl` calls in your code and allow the automatic `import` feature to do its thing. it seems to select a version seemingly at random. I have seen all of these, at least:

```go
import "github.com/go-gl/gl/v3.3-core/gl"
import "github.com/go-gl/gl/v4.2-core/gl"
import "github.com/go-gl/gl/all/gl"
```

This will typically compile without issue, but fail in unexpected ways -- `segfault`s are common -- at runtime.

Every time this has happened I've wasted 30 minutes or so trying to determine what I've done wrong.

The added advantage of using a Makefile is that it allows any number of automations to be included in the build process. To that end, I've added a new tool to audit my code for precisely this problem before compiling.
