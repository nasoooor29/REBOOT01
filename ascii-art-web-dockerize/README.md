# ascii-art-web-dockerize

## Description

This project is a web server written in Go that allows users to input text and output the text into ascii-art. 

The webpage makes use of the below banners:

- shadow
- standard
- thinkertoy

## Authors

- [nhussain](https://learn.reboot01.com/git/nhussain)
- [yabuzuha](https://learn.reboot01.com/git/yabuzuha)
- [etarada](https://learn.reboot01.com/git/etarada)

## Usage

Simply run the program using the supplied bash script `build.sh` like so to start the server:

```
./build.sh

```

Aleternatively refer to the audit questions and run the commands manually 1 by 1.

The webpage is available to view at `http://localhost:8080`. On the main page, type your text into the text area. Select the font from the drop down and this will change the style of the ascii-art (Standard by default). 

## Implementation

- The webpage creates ASCII art using the AsciiArt function, which receives input via a form on the webpage. 
- This function searches through text files dedicated to different styles such as standard, shadow, and thinkertoy. 
- For each character in the input, it calculates a modified ASCII value to locate the corresponding ASCII art in the appropriate text file. 
- The results are then presented on a new page using the POST method. 
- Additionally, the website manages 400, 404, and 500 status errors by redirecting users to respective error pages hosted on the server.
- The Dockerfile uses a multi stage build process to build the app and move the binary and relevant files to a light weight alpine container.
