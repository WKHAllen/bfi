# BFI (Brainfuck interpreter)

An interpreter for the Brainfuck language, written for the web using Go.

## Design Choices

The backend for this project is written in Go, using [Gin](https://github.com/gin-gonic/gin) to run the HTTP server. The frontend logic is written in JavaScript, using [jQuery](https://jquery.com/) for AJAX calls.

The brainfuck code is sent to the server once the button is pressed. The server will create a new interpreter object and begin evaluating the code. When a '.' or ',' is reached, the server will stop and send the request to either print a character or get keyboard input from the user. The client will then tell the server this has been done, and the server will continue evaluating. This ends when the server reaches the end of the code, at which point the client is notified.

I know, it's not the ideal way to design such a piece of software. It could all be done on the frontend, using JavaScript. It would be much faster and simpler that way. But doing it in such a way would be far to straightforward and boring. I designed the program the way I did because I wanted a challenge.
