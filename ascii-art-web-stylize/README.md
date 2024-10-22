# ascii-art-web-stylize

## Description

This project is a web server written in Go that allows users to input text and output the text into ascii-art. this version of the project is responsive so it will work with any screen size. 

The webpage makes use of the below banners:

- shadow
- standard
- thinkertoy

## Authors

- [nhussain](https://learn.reboot01.com/git/nhussain)
- [yabuzuha](https://learn.reboot01.com/git/yabuzuha)
- [etarada](https://learn.reboot01.com/git/etarada)

## Usage

Simply run the program like so to start the server:


```
go run .
```

The webpage is available to view at `http://localhost:8080`. On the main page, type your text into the text area. Select the font from the drop down and this will change the style of the ascii-art (Standard by default). 

## Implementation

- The webpage creates ASCII art using the AsciiArt function, which receives input via a form on the webpage. 
- This function searches through text files dedicated to different styles such as standard, shadow, and thinkertoy. 
- For each character in the input, it calculates a modified ASCII value to locate the corresponding ASCII art in the appropriate text file. 
- The results are then presented on a new page using the POST method. 
- Additionally, the website manages 400, 404, and 500 status errors by redirecting users to respective error pages hosted on the server.