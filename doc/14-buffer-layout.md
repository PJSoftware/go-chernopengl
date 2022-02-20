# Buffer Layout Abstraction

## _Why_ Abstract Something?

Why exactly do we abstract something?

Simplifying the code, moving it into its own namespace for ease of handling, is a nice end result, but deciding what to abstract away -- and how to organise it -- can be tricky. Sometimes all you end up doing is creating a very thin wrapper, and changing the name of a function.

(**Aside**: given that we cannot do macros in Go, perhaps something like this to wrap each gl.\*\*\* call in error-checking code might actually be a valid approach. Something to keep in the back of my mind...)

One important reason for moving something out into its own class is so that you can identify your actual requirements for that code/object, and hence implement those functions.

### Vertex Array Objects?

For us, a vertex array needs to be able to tie together a vertex buffer with some kind of _buffer layout_.

A vertex buffer by itself is just a bunch of bytes with no inherent concept of its internal structure. The vertex array provides that layout information.
