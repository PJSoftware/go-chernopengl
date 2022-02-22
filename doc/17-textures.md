# Textures in OpenGL

We shall focus on simply adding an image to our square -- but textures can be so much more.

## Using Textures

In the video The Cherno is using a `PNG` file; I've opted for a `JPG` instead. He points out that often a game engine would use its own image format for a variety of reasons.

Reading the image will require a library (unless Go has native support for reading images into a format useable by OpenGL?) and either way, it will likely be different from the `C++` code.

- Use a library (in his case, `stb_image`) to read the image file into a buffer. The function needs to return a "pointer to a buffer of RGBA pixels."
- Upload our pixel array to the GPU as a texture.
- Bind the texture
- Modify our shader to read the texture; the fragment shader can read from the image data and specify the colour of each pixel.

## Using Textures in Go

Note that images are typically considered to start at the top-left corner, but `OpenGL` starts at the bottom-left. `stb_image` has a function to flip the buffer to deal with this. If our image comes in upside down, well, _there's yer problem!_

### Texture Slots

Our `Texture.Bind()` has a `slot` parameter. (In the original `C++` code it was specified as an optional parameter with a default of 0; `Go` does not support optional parameters so we'll need to explicitly specify `0`. This is not exactly _bad_!)

Note that a modern PC running `OpenGL` will likely have 32 slots available whereas a smart phone may have 8. However, whatever the number your platform has, it is possible to query `OpenGL` to find out the upper limit.

### Using Texture

To tell our shader to use our texture (and which slot to use) we pass it in as a uniform. Sort of.

It's not really an integer slot, it's a `sampler` slot. This is "a bit hazy, and a bit weird"!

We need to:

- send an integer uniform to our shader
- the integer is the slot we have bound the texture to.
- In the shader code we use that integer to sample from the texture.

## Shader

### Varying

To pass values between the vertex shader (which takes external input) and the fragment shader (which has an output) we use a "varying", typically with a "v\_" prefix. This is passed _out_ from the vertex shader, _in_ to the fragment shader.

## Image

When we finally tried to compile and run our code, we got:

```err
panic: interface conversion: image.Image is *image.YCbCr, not *image.RGBA
```

It seems to me that the simplest approach is to change our `JPG` file to a `PNG`. This now gave me:

```err
2022/02/23 02:06:36 decoding image: image: unknown format
```

Fixing this with:

```go
image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
```

This worked ... with the exception that our image was, as somewhat expected, upside down. The easiest way to change this should be to remap the texture coordinates.

## Final Note

After resolving a few issues, I got my image correctly displaying. The Cherno, however, did not. Turns out this is because his image had transparencies; to make this work we need to enable blending -- which is the topic of the next video!

I attempted to add transparency to my image to test the quick solution he added -- but got this:

```err
panic: interface conversion: image.Image is *image.NRGBA, not *image.RGBA
```

It seems, sooner or later, I'll have to figure out how to solve this issue. But ... not today. For now I'll leave that code in, but commented out.
