# Writing a Basic Renderer

This was fairly straightforward; moved the couple of gl calls from our loop in `main()` to `renderer.Draw()` and `renderer.Clear()`.

Because I've handled my code slightly different from The Cherno -- putting `shaderUniform` in its own class -- I've ended up with this in my loop:

```go
renderer.Clear()

uniform.SetUniform4f(r, 0.1, 0.3, 1.0)
renderer.Draw(va, ib, shader)
```

The `uniform` is linked (or related) to the `shader` by the earlier call:

```go
uniform := shaderUniform.New(shader, "u_Colour")
```

It would perhaps be useful to make this relationship a little more explicit in the code -- somehow. (Of course, at this point we only have one shader, so even the repeated `.Bind()` calls are overkill. But that will no doubt change...)

Of course, for a more advanced application we would probably be using a `material` rather than a `shader`... A `material` is essentially a `shader` plus all its `uniforms` -- exactly what I was talking about above without realising it.
