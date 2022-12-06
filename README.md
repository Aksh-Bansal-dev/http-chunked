# http-chunked
A basic server serving images as chunked response using Go.

Suppose a website has many images on a page, then rendering that page would be slow because the browser will try to create separate 
connection for fetching each image. Establishing new connections take time and also most browser only support a limited number of concurrent
connections. 

To improve page load, server can send multiple images in a single request and stream them using `Transfer-Encoding: chunked` so client can render
an image as soon as its data is received instead of waiting for the complete response(all images).

## How it works?
The server is encoding the image in base64 format, and sending it to the client in chunks. The client displays an image as soon as it receives the
data for that image. 

## Reference
[link](https://dropbox.tech/infrastructure/retrieving-thumbnails)
