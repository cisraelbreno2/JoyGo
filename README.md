
# JoyGo
Initial project of a bridge between a joystick with a keyboard and mouse

## Objective
The goal of this project is to create a bridge between a joystick and a keyboard and mouse. The idea is to allow the user to use the joystick as a keyboard and mouse, to control the computer in order to be able to place it in an environment such as a living room without having to keep the keyboard and mouse with them.

## Functionality
- The joystick can control the mouse and keyboard.
- When entering a game, simply press the right analog stick and "select" to disable keyboard and mouse input.
- So far, the joystick has only been tested on **macOS**.

## ⚠️ Compatibility Notice
Currently, the project **only works on macOS**.  
You must have the **SDL2 library installed** on your system.

To install it using Homebrew (MacOs):

```bash
brew install sdl2
```

### How to run the project

Clone the repository

Install the dependencies with
```bash
go get
```
Run the project with
```bash
go build main.go
```
Enjoy!

## Observations
- The project is still in the initial phase, so there are many things to be done.
- The project is being developed in Go, so it is necessary to have Go installed on your machine.
- The project does not yet have unit tests.