a tool for automating the search for gtfobins on the attacked machine 
all information about binaries is taken from here -> https://github.com/GTFOBins/GTFOBins.github.io

example usage: SUID
```sh
find / -uid 0 -perm -4000 -type f 2>/dev/null | ./gtfo_search
```
