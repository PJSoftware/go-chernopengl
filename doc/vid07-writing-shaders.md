# Writing a Shader in OpenGL

## Support code

- Compile
- Link
- Delete
  - After the shaders have been linked into the `shader program`, they can be deleted. This deletes the `compiled` version of the shader
  - Technically they should also be `detached` (which removes the shader source code itself!) However, keeping it attached enable more effective debugging if the shader is not working.

Quote of the day: "With an API as fragile as OpenGL, that gives you so little help when things go wrong, you want to master the documentation!"
