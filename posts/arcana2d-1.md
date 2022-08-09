===
title: Designing my game framework (Pt 1)
date: 09-08-2021
description: Designing a game framework isn't easy, and I can tell from experience. I'm going to talk about my old and now discontinued game framework Glimmer2D, why I redesigned it, and how my new game framework works.
===
# The life and death of Glimmer2D
So I'm currently working on a game engine kinda similar to Love2D, but using the Wren scripting language to expose the API instead. To make the design easy, I split the engine into a frontend and a backend. The backend implement the features of a game engine like windowing, rendering and input handling, whereas the frontend would help it interface with the Wren API. I initially designed the two together, but then I realised I could design the backend seperately as a seperate game framework and wrap it with the Wren frontend in the game engine.

And that's how Glimmer2D was born! Since I'd mostly worked around with the Raylib games library( a super-underrated game library that you should totally check [out](https://www.raylib.com/index.html)), I used it as my inspiration for the API. 

Raylib's API, in my opinion, is amazing for beginner developers with its usage of the pleasant-to-read Pascal case and its comfortable naming conventions that simply tell you what a function does or the purpose of a given struct is, especially in comparision to other game libraries written in the C programming language. Something like `Font` or `GetFPS()` is a lot more readable than something like `sapp_sgcontext()`, which, while is from the wonderful Sokol game library, still looks a lot more daunting to an amateur programmer. 

Anyway, coming back to my game framework, I not only based my API on Raylib's, but also ended up learning a lot from its code when it came to game engine design. However, I found myself aping too much from Raylib's API, to the point it just felt kinda like a C++ binding for Raylib, so it ended up feeling like a bit of a pointless creation. That wasn't what discouraged me from finishing the framework though, it was the endless bugs I ran into when designing the rendering, stuff that I wouldn't solve until I swapped to working on my new framework, Arcana2D, where I decided to start on working on the renderer instead, since that's the beefiest and most difficult part of the engine. 

# Arcana2D - how it works
Arcana2D is my newer version of the game framework with an entirely different API, and I ended up liking the API a bit more. I took a bit of inspiration from game engines like Sokol and Bevy where instead of running the game in the `main()` function, it instead passes control to an `App` object that the user configures, sets an `init()`, `update()`, `render()` and `finish()` function so the app can automatically run them, passing down arguments for the user to interface with the engine's API while maintaining privacy. Each function passes down different data, and serves a different purpose.

1. `init()` passes down an `AppConfig` so that the user can set details such as window size, background color, while also initialising/setting whatever data they want. A `UserData` object is used to store a void pointer to an object that stores of all of the game's relevant data that has to be shared by multiple functions, hence removing the need for global variables. All functions pass the `UserData` argument.

2. `update()` passes down a `GameContext` object that stores relevant information such as inputs buffered and delta time between frames. By controlling what data is passed down to the `update()` function, I can prevent draw commands from happening during the update phase of the game engine, so that the renderer and window can be prepared. The `update()` function just serves to update the game logic based on tabulated inputs and such, so that's why I wanted to seperate rendering from it.

3. `render()` passes down a `RenderContext` that simply allows the user to draw objects to the screen. How the renderer works and how I came to the decision of how I should implement it is something I look forward to discussing in a later blogpost.

4. `finish()` is just called when the game is closed, and only passes down  `UserData`. All of the game engine's data is automatically removed by the `App`'s destructor, so it just exists to do stuff like saving data and freeing user-allocated objects.

The structure of the framework is a lot more different from my initial version of it, but I am fond of it's design. It feels very contained and organised, and makes for a bit less of a hassle to code certain features with it, such as calculating delta time and capping frame rates. Well, hopefully I can finish up with this quickly so I can proceed to getting back to designing the Wren interface for my **ACTUAL** game engine. Until then, see you next time!