# Dealing With Shaders

- Quote of the day: "DirectX is just a better API"

## Reading Shaders from a file

I have already done this -- but The Cherno prefers to have both shaders in a single file, so let's do that.

Rather than slurping in the entire shader file, we need to read it line by line, looking for section breaks. My chosen break is:

```glsl
// shader: <shadertype>
```
