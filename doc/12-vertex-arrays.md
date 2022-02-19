# Vertex Arrays in OpenGL

What is a `vertex array` and how does it differ from a `vertex buffer`?

They _are_ similar -- but vertex arrays are unique to OpenGL; none of the other graphics APIs (DirectX, etc) have this concept.

They are "a way to bind _vertex buffers_ with a specification for the layout of that vertex buffer."

Currently, if we had multiple meshes, etc, we would need to:

- bind our vertex buffer,
- bind our index buffer,
- draw our geometry.

After we bind the vertex buffer we still need to specify our layout.

Let's look at what we would need to do if we had multiple vertex buffers with different attribute layouts. Essentially we need to unbind everything we have bound by calling bind with 0, and then rebind before the draw call.
