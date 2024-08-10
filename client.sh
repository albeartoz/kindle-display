#!/bin/bash

# Dependencies: scrot, imagemagick, curl

while true; do
    # take screenshot
    scrot -F ./image.png -z -o

    # Convert to kindle resolution and b/w, compress
    convert ./image.png -rotate 270 -resize '600x800!' -type GrayScale -depth 8 -colors 256 ./image.jpg

    # POST to kindle server
    # This is stupid
    curl -F "file=@./image.jpg" http://192.168.15.244:4545/upload

    # 1 FPS is all you need
    sleep 1
done
