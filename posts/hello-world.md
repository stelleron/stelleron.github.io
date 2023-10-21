===
title: Hello World!
date: 01-08-2022
description: My first ever post. I look at my previous blog and why I took it down, and how I came to create this blog and the static site generator that powers it in Python
===
# Hello World!
As the title and description says, this is my first ever post on this blog! However, this isn't the first time I've ever created a blog. Earlier this year, I created a blog using Jekyll Now during March to document my progress on my game engine and explore my design decisions. However, I found myself with a few problems during the making of the blog:
<br>
<br>

- **Updating way too inconsistently**: I only had three posts on the blog, and I made each post after a month's time. Since I was bogged down with schoolwork and exams and all of the pains and suffering of being a high schooler, I couldn't really make much time for this blog. Besides, when I did have some spare time, I would devote it to improving the game engine, segueing into my next problem:
- **Game Engine development progressed faster than I could blog about it**: Even though I had only three posts about my game engine's progress, I'd made many more leaps and bounds in the engine's development. In my second post I detailed how I setting up the basic framework for my engine. At that exact same point in time, I was working on extension loading and creating a Shader class for the engine.
I wanted to be able to track my development step by step, but I was running way too fast for that. I ended up losing track of what concept I should blog about next since I would have finished implementing it long ago and have forgotten about it soon after.
- **I didn't like using Jekyll Now**: This is more of a personal preference issue, since I just didn't link how Jekyll Now looked. Granted, it's an incredible tool and it made making blogs super convenient, but I just disliked how the font looked. Furthermore, I was struggling to get Jekyll installed on my Mac since the system's version of Ruby sucks and I kept getting errors saying that I didn't have permission access to download it. This meant that if I wanted to preview Jekyll's changes, I had to preview it after pushing my changes to my GitHub repo, and those changes could take many minutes to load.

<br>
So, sometime during May, I took my blog down and flushed all evidence of it's existence down the toilet. But then come July and with a lot more free time now that I wasn't being anchored down by exams, I decided to reinstate my blog. However, I decided not to use Jekyll Now this time. I ended up looking into other options.

## Creating my second blog
First, I looked into Zola, since I liked the Hyde theme ported from Jekyll, and since I was struggling with getting Jekyll installed, I went for the next option. But I kept struggling with using Zola and loading in the Hyde theme, so I eventually ditched it out of impatience. I eventually figured out how to install Jekyll via using another version of Ruby instead of the system version, but I came to realise how outdated Hyde was, and that it was incompatiable with recent versions of Jekyll. I eventually gave up with using external static site generators, and created one of my own.

## How I made Nullsite
I first tried to create Nullsite with Rust since most of what resources I looked into to learn how to create my own static site generator were in Rust, but I eventually swapped to Python for convenience. It had been over a whole year since I'd touched Python, and I must say, coming back to Python after using C++ feels like a breath of fresh air. The ease of usage and zero compile times, not needing to type in `#includes` everywhere and setup a project with Premake, shell files and all of my libraries made the developer experience so much more enjoyable. Oh, how I've missed you Python.

The idea is simple - the Python file simply just looks at a provided config.toml file to find some important details, then loads all of the required Markdown files, converts it into HTML and injects it into a given template and saves it to create a blogpost. The information from all of the blogposts is saved and used to generate the index, which also sorts all of the blogposts by date.

Also note that each Markdown blogpost comes with a YAML Frontmatter, which is essentially the stuff at the beginning of each Markdown file between the two lines made up of `=` symbols. This just contains all of the metadata, which includes:

- Title of the blogpost
- Date of publishing
- A basic description

The end result of Nullsite is slightly janky since I made it in a day, and it does use tons of global variables and repeated code blocks, but it works, so I don't really mind. It'll probably be a pain to add anything new to it though.

I'm still happy with the end result, which you can check in the [Projects](/site/projects.html) tab.

## Signing off
Overall, I'm looking forward to start blogging again properly. My dev logs aren't necessarily gonna be tracking any of my projects, but rather just exploring stuff I've learned or stuff I've done in my projects, and I'll try to keep new blogposts with two weeks of each other at the very most. Now, time for me to get back to work!