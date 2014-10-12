tl
========

Tail multiple files with coloured pattern matching.

Usage of tl:
		-f="": 		CSV list of files to tail  
		-p=false: 	Prefix output with file name  
		-r="": 		List of regexs to search for  
		-c="": 		List of regex colours  
		-d=",":		Regex list delimiter  


Available colours:
* r = red
* g = green
* y = yellow
* b = blue
* m = magenta
* c = cyan
* w = white
* k = black  

default = r


Example
-------

Tail install and system log files, alerting with various colours for specific events.  

tl -f=install.log,system.log -r="install,failed,network is,Unrecognized leaf certificate" -c=y,r,b,r
