# OpenGL with The Cherno

On YouTube, [The Cherno](https://www.youtube.com/channel/UCQ-W1KE9EYfdxhL6S4twUNw) has an excellent [OpenGL playlist](https://www.youtube.com/watch?v=W3gAzLwfIP0&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2), which provides a guide to learning OpenGL, using `C++`.

This is my attempt to follow along, while translating on the fly: I want to use `Go` instead.

Additionally, check out:

- [the Cherno's discord channel](https://thecherno.com/discord)
- the [OpenGL API Documentation](https://docs.gl)

## Videos

1. Welcome to OpenGL ([doc](doc/01-welcome.md)) ([video](https://www.youtube.com/watch?v=W3gAzLwfIP0&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=1))
2. Setting up OpenGL and Creating a Window in ~~C++~~ Go ([doc](doc/02-setup.md)) ([video](https://www.youtube.com/watch?v=OR4fNpBjmq8&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=2))
3. Using Modern OpenGL in ~~C++~~ Go ([doc](doc/03-modern-opengl.md)) ([video](https://www.youtube.com/watch?v=H2E3yO0J7TM&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=3))
4. Vertex Buffers & Drawing a Triangle ([doc](doc/04-vbuf-triangle.md)) ([video](https://www.youtube.com/watch?v=0p9VxImr7Y0&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=4))
5. Vertex Attributes and Layouts ([doc](doc/05-vertex-stuff.md)) ([video](https://www.youtube.com/watch?v=x0H--CL2tUI&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=5))
6. How Shaders Work in OpenGL ([doc](doc/06-shaders.md)) ([video](https://www.youtube.com/watch?v=5W7JLgFCkwI&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=6))
7. Writing a Shader ([doc](doc/07-writing-shaders.md)) ([video](https://www.youtube.com/watch?v=71BLZwRGUJE&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=7))
8. Dealing with Shaders ([doc](doc/08-dealing-with-shaders.md)) ([video](https://www.youtube.com/watch?v=2pv0Fbo-7ms&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=8))
9. Index Buffers in OpenGL ([doc](doc/09-index-buffers.md)) ([video](https://www.youtube.com/watch?v=MXNMC1YAxVQ&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=9))
10. Dealing with Errors in OpenGL ([doc](doc/10-errors.md)) ([video](https://www.youtube.com/watch?v=FBbPWSOQ0-w&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=10))
11. Uniforms in OpenGL ([doc](doc/11-uniforms.md)) ([video](https://www.youtube.com/watch?v=DE6Xlx_kbo0&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=11))
12. Vertex Arrays in OpenGL ([doc](doc/12-vertex-arrays.md)) ([video](https://www.youtube.com/watch?v=Bcs56Mm-FJY&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=12))
13. Abstracting OpenGL into Classes ([doc](doc/13-abstracting-opengl.md)) ([video](https://www.youtube.com/watch?v=bTHqmzjm2UI&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=13))
14. Buffer Layout Abstraction in OpenGL ([doc](doc/14-buffer-layout.md)) ([video](https://www.youtube.com/watch?v=oD1dvfbyf6A&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=14))
15. Shader Abstraction in OpenGL ([doc](doc/15-abstracting-shaders.md)) ([video](https://www.youtube.com/watch?v=gDtHL6hy9R8&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=15))
16. Writing a Basic Renderer in OpenGL ([doc](doc/16-renderer.md)) ([video](https://www.youtube.com/watch?v=jjaTTRFXRAk&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=16))
17. Textures in OpenGL ([doc](doc/17-textures.md)) ([video](https://www.youtube.com/watch?v=n4k7ANAFsIQ&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=17))
18. Blending in OpenGL ([doc](doc/18-blending.md)) ([video](https://www.youtube.com/watch?v=o1_yJ60UIxs&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=18))
19. Maths in OpenGL ([doc](doc/19-maths-in-opengl.md)) ([video](https://www.youtube.com/watch?v=VuYnjsDOx60&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=19))

### Still Pending

- Projection Matrices in OpenGL<!--  ([doc](doc/20)) ([video]()) -->
- Model View Projection Matrices in OpenGL<!--  ([doc](doc/21)) ([video]()) -->
- ImGui in OpenGL<!--  ([doc](doc/22)) ([video]()) -->
- Rendering Multiple Objects in OpenGL<!--  ([doc](doc/23)) ([video]()) -->
- Setting up a Test Framework for OpenGL<!--  ([doc](doc/24)) ([video]()) -->
- Creating Tests in OpenGL<!--  ([doc](doc/25)) ([video]()) -->
- Creating a Texture Test in OpenGL<!--  ([doc](doc/26)) ([video]()) -->
- How to make your UNIFORMS FASTER in OpenGL<!--  ([doc](doc/27)) ([video]()) -->
- Batch Rendering - An Introduction<!--  ([doc](doc/28)) ([video]()) -->
- Batch Rendering - Colors<!--  ([doc](doc/29)) ([video]()) -->
- Batch Rendering - Textures<!--  ([doc](doc/30)) ([video]()) -->
- Batch Rendering - Dynamic Geometry<!--  ([doc](doc/31)) ([video]()) -->

### Other Topics

- Materials(?)
