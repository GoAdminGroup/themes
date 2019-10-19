# Build your own template

Following the steps to make it work.

## step 1

The files in the ```resource/pages``` is the golang template files of your **Template**.
Now you have to change them to your own custom content. 

## step 2

The files in the ```resource/assets/src``` is the corresponding assets to the template files.
Put them with a right order.

## step 3

See the Makefile, modify it to meet your need.
Finally, run ```make assets```to generate the ```template.go``` and ```assets.go``` and ```assets_list.go```.
