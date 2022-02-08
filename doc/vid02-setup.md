# Getting Started

Typically when creating a Window, you would use OS-specific APIs to generate your on-screen window.

For creating an actual game, it is potentially worth going down that path. (Except since we are going to be using `Go` instead of `C++`, is that even an option?)

For this series, however, we simply want to focus on OpenGL, so we shall use a library to handle our window creation and management for us:

- [glfw](https://www.glfw.org/)
  - It is a nice lightweight library.
  - It is ported to `Go`.

## Adding glfw

The first thing we need to do to setup our program in `Go` is initialise our module:

```sh
go mod init www.github.com/PJSoftware/go-chernopengl
```

Next, to get the `glfw` library:

```sh
go get -u github.com/go-gl/glfw/v3.3/glfw
```

(Cherno is using 3.2, but 3.3 is the current latest version -- and what my other code uses.)

This actually seems rather simpler than achieving the same task in `Visual Studio`, for a `C++` project.

## Legacy OpenGL

We would not ordinarily use Legacy OpenGL in production code, but we use it here to test that everything is working.

First, however, we need to add it to our project:

```sh
go get github.com/go-gl/gl/v4.1-core/gl
```
