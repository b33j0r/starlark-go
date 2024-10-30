print("hello")
"world" |> print
1 |> print

f = lambda x, y: x + y
g = partial(f, 1)
g(2) |> print

# example of arity 3
f2 = lambda x, y, z: x + y + z
h = partial(f2, 1, 2)
h(3) |> print
