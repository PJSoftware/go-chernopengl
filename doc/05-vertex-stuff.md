# Vertex Attributes and Layouts

Our "vertex" code to date looks like this:

```go
positions := []float32{
  -0.5, 0.5,
  0.0, -0.5,
  0.5, -0.5,
}
```

We have three sets of two values, so each "vertex" represents a position on the screen (or rather, on the `OpenGL` display-space which will map to a position on the screen.) So a "vertex" is a location in space?

No. That is how they are being used here, but to OpenGL a vertex can contain so much more than just coordinates.

It can include:

- position
- texture coordinates
- normals
- colours
- bi-normals
- tangents
- etc

Each of these is an "attribute". Each attribute in the vertex has an index: ie, position is index 0, normal is index 1, etc -- as defined by us, the user.
