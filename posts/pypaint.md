===
title: Making PyPaint: A simple painting app I made in less than a 150 lines of code
date: 11-09-2022
description: So as a bit of a break while I work on a final version of the Arcana2D framework, I created a simple painting app in Python and Pygame. Here's my breakdown of it.
===
This one took me an afternoon, but was a pretty fun project to flex my Python skills with something a bit more advanced than my Brainf**k project. Once again, it's pretty tiny so there isn't much to break down, but I'll try my best. 

# Using Pygame
I actually started out learning game programming using Pygame, and while I don't think it's the best library for larger indie games because of how much a potato Python can be when it comes to performance, I do think it's a great library for learning how to make games, as well as prototyping. Not only is Python really easy to use and cuts development times down massively, Pygame has a nice API wrapping over the battle-tested SDL2 library.

Anyway, cutting from that tangent, I only needed to use Pygame for this project since it allowed a handy feature to save a surface (which essentially stores and can be used to render pixel data) as an image, which simplifies the process of creating a paint app. So now that I had a library to use, it was time to start hacking!

# The App's Features
Using the information that a surface can be saved as an image, we can create a window and draw a canvas surface to it. The canvas is where the user draws images with a paintbrush so it can be saved to memory. Speaking of the paintbrush, the brush can function by drawing circles to the canvas, and will blit cirlces as long as the left mouse button is being held down. The brush also needs to be resizable to draw bigger circles, and I decided to make that enabled with the mouse scroll for convenience's sake. I needed to also include a color selection panel for users to select colors from to draw with. Finally, I wanted the user to be able to clear the canvas with the C key and save the image with the S key. That's all I wanted to implement for this bite-sized project to keep things simple. So, to summarize:

1. Color selection panel
2. Canvas
3. Paintbrush that draws circles
4. Resizing with mouse scroll
5. Clear canvas with C key
6. Save drawing with S key

# How I implemented each step

1. The color selection panel would own multiple circular color icons, and would blit them all to its own surface. Whenever the left mouse button was clicked, the panel would pass the mouse coordinates to each icon so they can check whether it collided with them. If so, the color panel had a variable that stored the paintbrush's current color which it modified.

2. Simple, create a canvas and fill it with WHITE color. Then, I drew it below the color panel. This did lead to an awkward problem where mouse coordinates wouldn't accurately translate to surface coordinates because the canvas was smaller than the screen, so I just made the paintbrush offset where it drew circles.

3. The paintbrush would read the selected color from the draw panel class and if the mouse button was held down, would draw the circle onto the screen by setting a variable to True or False that specified whether it should draw a circle. This would be set back to False and drawing would be prevent when the left mouse button was released

4. Based on the speed of the mouse scroll, the size of the paintbrush would change by some amount, and that size would equate to the radius of the circle. I did put some limits on the paintbrush size to avoid things getting too ridiculous, or for it to just be shrunken down.

5. Just detect if the C key is pressed and if so, fill the canvas with white color.

6. And finally, if the S key is pressed, invoke the `pygame.image.save` command to save the canvas as a PNG. I decided to set a default PNG filename called `my_drawing.png` so I didn't have to create a weird file dialog or break the flow of the app by printing to a seperate console, prompting input.

It's fun making small projects like these that end up surprisingly complex to implement for their size, but this'll probably be the last one I do until I finish the game engine. Well, I can't wait to get that done with.