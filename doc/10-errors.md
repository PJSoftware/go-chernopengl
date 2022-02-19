# Dealing with Errors in OpenGL

For this episode we'll focus on `gl.GetError()`, but from version OpenGL 4.3 they have added `gl.DebugMessageCallback()` which gives much more useful information, and is generally nicer to use.

Ultimately this video shows how to use C++ macros to wrap every call to a gl-library function in a call which clears errors before the required function call, and checks for errors afterwards. Go does not appear to make this possible. Macros are not available, and passing a function to another function requires that it be specified exactly -- including parameters and return type.

I shall leave the existing glClearError() and glPanicOnError() in place for now -- the latter demonstrates how best to use a for loop in place of the while loop used in the video (Go does not have `while`, and my first attempt did it wrong) but I'll need to do some more research into `go-gl` error handling suggestions.
