===
title: Designing my game framework (Pt 4)
date: 10-06-2022
description: Well, I've done it! I've released a stable version of the Arcana2D framework, as well as of the MAGE Engine, and I wanted to share my thoughts about designing both of them.
===
These past few months have been exhausting, but I finally produced a release version of both the Arcana2D framework **AND** the MAGE Engine! It's been a rough ride, especially with the amount of code I've had to write, which totalled up to 5000+ lines, more than any other project I've ever made before. For this final blogpost, I'll just be talking a bit about my grievances with rendering TTF fonts and playing audio in C++, and my overall thoughts on the Arcana2D framework.

# Font Rendering
I still have nightmares of trying and failing to render TrueTypes, with the framework either throwing a segmentation fault or generating an incomprehensible bitmap. The main difficulty with trying to get TTFs to work for me was trying to generate the bitmap. TTF font rendering is a tricky process that involves a couple of steps to follow:

1. First, you have to get the data for each glyph in the TTF and store it in an image. 
2. Then, you have to fuse all of those images into one large bitmap. This is done so that you don't have to bind a texture for each glyph, which can get out of hand and slow down text rendering as the engine constantly switches between textures for rendering different characters. The target rectangles for each glyph should also be stored, so that texture coordinates can be determined by the renderer to identify seperate characters and which ones should be rendered.
3. Since TTF image data is stored in grayscale, it has to first be converted to a gray alpha format so that all of the blank area within each glyph is made transparent. From there, the bitmap can be converted into a texture.
4. From there, the texture can be passed to the renderer to render text. Each of the aforementioned target rectangles has to be mapped to a character so that the renderer can find where each character is in the texture.

So font rendering is a pretty complex task. My first issue was trying to find a library for loading TTFs. I initially tried to use the Freetype library, but constantly ran into problems trying to set it up on my system. I ended up deciding to use `stb_truetype.h` since it was the best alternative and easy to set up (I just had to download and include the header file).

However, I had no idea whatsoever as to how I could implement bitmap generation, so I just shamelessly took from the Raylib code again. I would like to mention here quickly that the Raylib source code's definitely something to look into when designing your game engine for the first time. It's well documented, and implements a ton of features that you might want to add in your game engine.

Anyway, back to TTF rendering. While I took Raylib's bitmap generation algorithm, I kept running into issues with it throwing segmentation faults or generating buggy bitmaps, so I had to spend a few days tweaking the algorithm until it functioned properly. Afterwards, it was pretty easy to make the renderer draw text to the screen. It just had to iterate through each character in the line of text to be printed, draw each character to the screen from a certain specified position, and then add the character's rectangular dimensions as an offset so that the next character can be print to the right of the current one.

# Playing Audio
I never realised just how complex audio management was until I tried to add it to Arcana2D. I initially tried to work with Miniaudio for audio management, but I found myself wasting too much time trying to figure out how the whole thing worked rather than implementing anything in the framework. So, for the time being, I settled on using irrKlang for a really simple and barebones audio player. In the future, I'll probably slowly phase it out for Minaudio and provide a lot more features, but until I've tested it out for myself and figured out how to make audio management work, the framework will be relying on irrKlang to play audio.

# Overall Thoughts
Overall, I'm pretty satisfied with how the framework turned out, considering that it was my first time ever making something like this. I do wish I could've fleshed out the audio API a lot more, and while future updates will definitely expand it, it is probably the weakest part of the framework right now. But otherwise, I think Arcana2D offers a lot of the base features needed for developing games, and it's API feels nice to use. 

For the MAGE Engine, I feel that Arcana2D's API is a bit weirder, but I still think it works. Creating the Wren API required a load of effort on my part, but finishing it was satisfying, knowing I had finally finished a big project. I'm looking forward to further improving it, but for the meantime I'll be taking a break from it cause I certainly need one.