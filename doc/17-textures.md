# Textures in OpenGL

We shall focus on simply adding an image to our square -- but textures can be so much more.

## Using Textures

In the video The Cherno is using a `PNG` file; I've opted for a `JPG` instead. He points out that often a game engine would use its own image format for a variety of reasons.

Reading the image will require a library (unless Go has native support for reading images into a format useable by OpenGL?) and either way, it will likely be different from the `C++` code.

- Use a library (in his case, `stb_image`) to read the image file into a buffer. The function needs to return a "pointer to a buffer of RGBA pixels."
- Upload our pixel array to the GPU as a texture.
- Bind the texture
- Modify our shader to read the texture; the fragment shader can read from the image data and specify the colour of each pixel.
