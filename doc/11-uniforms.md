# Uniforms in OpenGL

## What is a Uniform?

Essentially a way of passing data from our CPU (ie, our C++ or Golang code) to the GPU (the shader).

Passing data in this way can be done via `uniforms`, or via a `vertex buffer`.

Uniforms are per-draw, so you can set a uniform up before each call to draw.

Note that when using `GetUniformLocation()`, a value of -1 means the specified uniform could not be found.

This may indicate an error. However, if the uniform is specified in the shader code but not actually used, the shader compiler will probably drop it.

Note: For our current code, this does not allow us to draw different triangles in different colours. It is per-**draw**. To have different triangles be different colours, we would need to do that via a vertex attribute.
