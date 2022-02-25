# Model, View, Projection

Previously we referred to the projection matrix as the MVP -- but typically "MVP" refers to three separate 4x4 matrices:

- Model matrix (model location in 3D space)
- View Matrix (camera view in 3D space)
- Projection Matrix (projection into 2D space)

These are multiplied together (in that order -- _except that because OpenGL is in column-major order, we actually multiply them in PVM order_) to form the final MVP matrix: our transformation pipeline applied to every vertex.

Translation, rotation, scale
