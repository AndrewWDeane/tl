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


Examples
-------

Tail install and system log files, alerting with various colours for specific events.  

	tl -f=install.log,system.log -r="install,failed,network is,Unrecognized leaf certificate" -c=y,r,b,r

Tail install and system log files, alerting on network activity and loss of network, having switched delimiter  

	tl -f=install.log,system.log -d="|" -r="not reachable, netbiosd|network" -c="r|m"  

Tail system log highlighting service numbers  

	tl -f=system.log -r=service\[[0-9]*\]

Tail system log highlighting date time  

	tl -f=system.log -r="[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}"
