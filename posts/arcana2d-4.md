===
title: Designing my game framework (Pt 4)
date: 06-10-2021
description: Well, I've done it! I've released a stable version of the Arcana2D framework, as well as of the MAGE Engine, and I wanted to share my thoughts about designing both of them.
===
These past few months have been exhausting, but I finally produced a release version of both the Arcana2D framework **AND** the MAGE Engine! It's been a rough ride, especially with the amount of code I've had to write, which totalled up to 5000+ lines, more than any other project I've ever made before. For this final blogpost, I'll just be talking a bit about my grievances with rendering TTF fonts and playing audio in C++, and my overall thoughts on the Arcana2D framework.

# Font Rendering
I still have nightmares of trying and failing to render TrueTypes, with the framework either throwing a segmentation fault or generating an incomprehensible bitmap. The main difficulty with trying to get TTFs to work for me was trying to generate the bitmap. TTF font rendering is a tricky process that involves a couple of steps to follow:

1. First, you have to get the data for each glyph in the TTF and store it in an image. 
2. Then, you have to fuse all of those images into one large bitmap. This is done so that you don't have to bind a texture for each glyph, which can get out of hand and slow down text rendering as the engine constantly switches between textures for rendering different characters. The target rectangles for each glyph should also be stored, so that texture coordinates can be determined by the renderer to identify seperate characters and which ones should be rendered.
3. Since TTF image data is stored in grayscale, it has to first be converted to a gray alpha format so that all of the blank area within each glyph is made transparent. From there, the bitmap can be converted into a texture.
4. From there, the texture can be passed to the renderer to render text. Each of the aforementioned target rectangles has to be mapped to a character so that the renderer can find where each character is in the texture.

So font rendering is a pretty complex task. My first issue was trying to find a library for loading TTFs. I initially tried to use the Freetype library, but constantly ran into problems trying to set it up on my system. I ended up deciding to use STB Freetype since it was the best alternative and easy to set up (I just had to download and include the header file). 