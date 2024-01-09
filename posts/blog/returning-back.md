===
title: Returning back!
date: 01-08-2024
description: I have returned! This blog post will mainly be covering some updates - basic summary of what's changed and my new projects since - as well as the new and improved Nullsite, and how and why I re-implemented it in Go.
===
# I'm back! 
It's been about an year since I made my last blog post, and a lot has changed since then. I finally graduated from high school, and joined the University of Massachusetts Amherst as an Electrical Engineering and Computer Science major. I've also just been practicing my coding skills, reworking my old game framework and ended up creating a new one (which I'll discuss more in a bit!), and I've begun creating a roadmap for my coding career and for the future of this blog. I've also revamped Nullsite, the static site generator that produces the pages for this blog! I have a lot to cover in this blogpost, so I'll start moving down the list.

## The new game framework - sorta
So I created a new-ish game framework called Xenon, which is basically just my old Arcana2D framework but with an SDL2 backend instead of GLFW. When I was initially working on a game prototype in my old Arcana2D framework, I noticed that when I tried to create a fullscreen window, I wasn't able to get a cool looking native fullscreen window on my MacBook. When I did some research on why I couldn't get native fullscreen working, I learnt that GLFW hadn't actually implemented native fullscreen on MacOS.
For Windows users, MacOS allows you to open and run apps in a new window (as shown below by the image), which I believe is different from traditional maximize/fullscreen displays on a Windows computer. I initially wanted to work around this by having users of the game open the window in native fullscreen themselves, but GLFW couldn't detect native fullscreen, so I eventually decided to ditch GLFW and rewrite the framework with SDL, since SDL had native fullscreen for MacOS.

![Example of Fullscreen Window on MacOS](/images/returning-back/window.png)

Hence, Xenon was born. Initially starting out as a rewrite of my game engine with an SDL2 backend, I would eventually keep updating the Xenon framework while deprecating the Arcana2D framework. Now I've added expanded audio features, improved font rendering, and even cleaned out a few bugs that were in the original framework. Moving forward, everything I do will be in the Xenon framework, as it has now distanced itself a quite bit from the original framework it originated from, and I'm looking forward to building new games with it.

## Nullsite Go - The new and improved Nullsite
When I first created this blog, I built it with a custom static site generator written in Python which took a bunch of Markdown files and converted them to HTML, to produce an end result that looked like this: 

![Old Nullsite](/images/returning-back/old-nullsite.png)

However, I found that the old Nullsite was a bit antiquated and suffered from a couple of issues, being:

- The website wasn't responsive and the top bar would become misshapen on smaller screens or width
- The Python markdown-to-HTML converter I was using didn't interpret code highlighting statements, so I wasn't able to use libraries like [highlight.js](https://highlightjs.org/) to implement syntax highlighting, which is pretty crucial for a coding blog.
-  I wanted to be able to add other pages for displaying my resume and individual project pages, as well as rework the visuals for my website to look more appealing to recruiters and just visitors of the site in general.

So, I decided to rewrite my project, this time using Go instead of Python since I wanted to do a project to learn Go and because I was able to find a markdown-to-HTML converter that could implement syntax highlighting. The code I would implement would be pretty similar to my original Nullsite implementation, but in a different language, and parsing a different file structure used to arrange the website like so:

- Homepage 
    - [Projects] 
    - About 
    - Resume 
    - Blog 
        - [Blogposts] 

I would also add more parameters to the ```config.toml``` file to allow for more customization, and add a few animations to display each of the blogposts and projects. While Nullsite Go is still under construction and needs a few more touches, I hope to complete it within the week.

## Looking to the future
Aside from finishing up the blog, there's a lot of projects I'm looking to try and get done this year and blog about, and I wish to commemorate this blog by starting out with a either a blogpost or even a series detailing my exploration into full-stack web development. I'm currently working on a project with React and Django, and if all goes well, I'll be able to complete it within January, and have a blogpost ready then. Until then, this is me signing out for what I hope doesn't end up being another silent year! 