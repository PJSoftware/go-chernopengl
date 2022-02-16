# Index Buffers in OpenGL

## Drawing a square

Everything our GPU draws is made up of triangles. so to draw a square we need to actually draw two triangles

The naive method is to specify two triangles by specifying 6 vertices. But two of these are now specified twice. A square has only 4 vertices, so why do we need to specify 50% more vertices than we need?

Index buffers help avoid this. Simply put: we specify each vertex once, then draw pts 1,2,3 and pts 3,4,1.

Using an index (element) buffer, and DrawElements() instead of DrawArrays(), is the preferred approach in 90% of the time.
