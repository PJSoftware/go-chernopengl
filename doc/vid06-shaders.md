# How Shaders Work

## What is a Shader?

"A shader is just a program -- a block of code -- that runs on your GPU"

Some things, the GPU does better; some things the CPU does better. You need to strike a balance between what is run on the CPU (via your main program) vs what is run on the GPU (via the shader)

Two main types that we shall be focussed on are:

- Vertex Shaders
- Fragment Shaders (AKA Pixel Shaders)

- Other types include:
  - tesselation shaders
  - geometry shaders
  - compute shaders

Pipeline: Draw call -> Vertex shader -> Fragment shader -> Graphics on screen

This is, of course, simplified, with multiple steps along the line not shown.

**Vertex shader** is called once for each vertex; it tells OpenGL where you want the vertex to be on-screen, in your window. Can also be used to pass data through to the fragment shader.

It may be doing calculations to transform positions depending on camera positions, etc...

**Fragment shader** (AKA pixel shader, although a fragment and a pixel are not quite the same thing) is run once for each pixel which needs to be drawn. Therefore, for a 800x600 window, the fragment shader may be called up to 480000 times. (For our current image of 1 triangle, it will be called once for each pixel in the triangle.) The primary purpose of the fragment shader is to determine what colour that pixel should be.

Any calculation (eg, `5*5`) performed in the Vertex shader, will be performed (in our case) 3 times. The same calculation performed in the Fragment shader will be performed many more times.

Optimisation is an important consideration. Performance should be considered _all the time_. Try to perform calculations as cheaply as possible. **However**, for cases such as lighting (for instance) it is necessary to calculate it per pixel anyway; it cannot be avoided.

Approximately 80-90% of everything you see onscreen in modern games is via these two shaders.

Such shaders can be thousands of lines of code. Game engines may even be generating and compiling shaders on the fly.
