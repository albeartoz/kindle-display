# Kindle-Display

A super janky way to mirror your monitor to your Kindle  

### Requirements
- A jailbroken Kindle (see MobileReads Kindle developer corner) with the USBNetwork package installed
- A Linux machine with scrot, imagemagick, and curl installed

### Getting started
- Compile server.go targetting ARMv5
- Connect to your kindle via USBNet
- scp the compiled file to your Kindle
- Run the server on the kindle
- Run the client on your Linux machine

### Limitations
This thing tops out at 1fps on my Kindle Basic 2. It might perform better on more modern Kindles but you'll have to modify the sleep in the bash script to get it to sent more screenshots.  
I'm sure sending screenshots over http is not the best way to do this but any other option is much more complicated. Maybe send a video stream over netcat and just pipe it directly to the framebuffer somehow? IDK  

### TODO
- make the screenshot path and timestamp variables command line options
- port X server to Kindle
