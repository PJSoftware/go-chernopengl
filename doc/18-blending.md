# Blending

## How do we control blending?

- Three Ways:

  - `gl.Enable(gl.BLEND)` -- `gl.Disable(gl.BLEND)`
  - `gl.BlendFunc(src, dest)`
    - src = how the src RGBA factor is calculated (default: `gl.ONE`)
    - dest = how the dest RGBA factor is calculated (default: `gl.ZERO`)
  - `gl.BlendEquation(mode)`
    - mode = how we combine the src and dest colors (default: `gl.FUNC_ADD`)
  - The above defaults => "use the source value"

- In our code we used:
  - src = `gl.SRC_ALPHA`
  - dest = `gl.ONE_MINUS_SRC_ALPHA`
  - This => "if transparent, use the destination colour"
