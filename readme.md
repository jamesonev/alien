### Getting started

Here is a really simple controller function you can use to simulate the attack. Wherever this is being invoked from needs to have the input file.

```
package main

import "github.com/jamesonev/alien"

func main() {
	alien.Attack()
}
```

You can run this with `go run main.go 5`, with however many aliens you want. You also have the option to specify an input file for the program as the final argument in the command line. If no file is specified, the program will look for a file named `input.txt` in the current directory.

At the end, Attack prints the current atlas including any roads between cities. It also does a check of all remaining cities for aliens and prints out the location and number of any found, which is helpful for checking the correctness of the implementation

### Development process

This is my first Go project, and so it had quite a steep learning curve. Conceptually, I'm used to separating my code into .c/.h files and linking them in a makefile, so transitioning into thinking in terms of packages was challenging. You can see that I tried to use good cohesion/coupling practices, like how `atlas.go` has no knowledge of the battle. However, figuring out the most elegant way to do that in a package was challenging.
