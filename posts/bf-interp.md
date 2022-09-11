===
title: Writing yet another Brainf**k interpreter in Python
date: 11-09-2021
description: For a bit of fun, I decided to make yet another Brainf**k interpret in Python. Here's a simple breakdown of the language and how the code works
===
Here's a little fun project I whipped up in an hour. The implementation is just under a 100 lines of code, and with stripping comments and spaces out, it probably comes out to be around 60 lines. Let's jump right into it!

# What is Brainf**k anyway?
If you already know how the language works, you can skip this section. If you don't know, you can either read this section, or check out this amazing [video](https://www.youtube.com/watch?v=hdHjjBS4cs8&ab_channel=Fireship) by Fireship.io that summarises the whole in a 100-something seconds.

So basically, Brainf**k is an esoteric programming language famous for being incredibly minimalist, with only 8 commands. While it's Turing complete, meaning the programming language can theoretically solve any computational program, it's really impractical. It's still a fun thing to hack around with. So here's how the language works.

First, it creates a one-dimensional array of at least 30,000 bytes with all values initialised to 0, and a pointer is set to the first element in the array. The programming language essentially works by manipulating this array. Second, everything that isn't a command gets ignored and treated as a comment. So, what are the commands that the language offers? Well, Brainf**k offers the following 8 single-character commands:

1. `>` to increment the data pointer and point to the element to the right
2. `<` to decrement the data pointer and point to the element to the left
3. `+` to increment the value stored in the current array cell the pointer is pointing towards
4. `-` to decrement the current value
5. `.` outputs the value stored at the pointer as a character
6. `,` take an input and store it in the pointed cell
7. `[` initiates a loop. If the byte at the data pointer is zero however, it breaks the loop and jumps to the next `]`
8. `]` if the byte at the data pointer is not zero, it jumps the program back to the last seen `[`, closing the loop

So that's a simple rundown of Brainf**k. The source code for a simple Hello World program can be found [here](https://en.wikipedia.org/wiki/Brainfuck#Hello_World!) (credits to Wikipedia). Looks pretty confusing doesn't it? But it's surprisingly easy to implement.

# How my interpreter works
1. First, the program takes a .bf file as input, and strips all new lines from it, before sending that data to an `interpret()` function.
2. The interpret creates an array of 30,000 integers(since bytes aren't a Python datatype as far as I know), and initialises it.
3. The program then loops through the source code, one character at a time. If it finds a command, it processes it. Else, it ignores it.
4. Note: Since everything in the array is an integer, the program has to cap output to numbers ranging from 0-255 cause otherwise it wouldn't be a valid ASCII character.
5. If the program runs into an error, it prints out an error message, the instruction which caused it and where it happened, and exits.

This works pretty well for linearly executed code. However, the interpreter also needs to implement jump commands, right? How does it do that?

# The Loop Stack
The idea's real simple: when a loop is initiated by a `[` command, it gets pushed onto a loop stack that memorises it's current position. When the program encounters a `]`, it checks the value stored at the pointer, as well as pops the last item of the stack being the command that initiated that particular loop. If the value at the pointer is 0, the program just moves on as usual. Otherwise, it jumps to the byte before the location of the `[` that created the loop, starting it once again. Stacks work well for this model because it allows for nested loops.

You can find the source code [here](https://github.com/stelleron/bf-interpreter/blob/master/src/main.py) I've done my best to comment everything in the project, but even then it should be easy to understand because the code is pretty small. Brainf**k interpreters are pretty popular projects to do for beginners, and I can see why. They're simple, require some knowledge of your programming language and hence test it well, don't need any external libraries, and can be really fun to make, like this one! 