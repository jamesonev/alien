### Here's where I break down what I will build

I'll begin by making an 'atlas' package (called atlas as map is reserved). The atlas is a mapping to city structs. This will make it easy to give the parsed input consistency (if I'm told Foo is south of Bar, then I can easily set Bar's north to Foo)

It seems best to make each city be a struct with a Name and 4 pointers, n s e and w.
