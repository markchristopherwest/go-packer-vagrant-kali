package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type commandClass struct {
	Command string
	Path    string
	Message string
}

func (class commandClass) StructMethod() {
	fmt.Println(class.Command)

}

func executor(runner commandClass) {

	fmt.Println("\nRunning in PWD: ", runner.Path)
	fmt.Println("\nExecute in CLI: ", runner.Command)
	fmt.Println("\nMessage in LOG: ", runner.Message)
	// Split up the command line arguments using space as a delimiter
	parts := strings.Split(runner.Command, " ")

	// The first part is the command, the rest are the args:
	head := parts[0]
	args := parts[1:len(parts)]

	// Format the command
	cmd := exec.Command(head, args...)
	cmd.Dir = runner.Path

	// Handle the Output
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("\ncmd.Run() our with cmd: ", cmd)
		log.Fatalln("\ncmd.Run() our with arg: ", args)
		log.Fatalln("\ncmd.Run() our with out: ", string(out))
		log.Fatalln("\ncmd.Run() err with err: ", err)
	}
	fmt.Println("\nSuccess @ cmd.Run() command: ", runner.Command)
	fmt.Println("\nSuccess @ cmd.Run() outputs: ", string(out))
	log.Println("Success @ cmd.Run() message: ", runner.Message)

}

func main() {

	// PWD
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	for _, p := range os.Args[1:] {
		pid, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		process, err := os.FindProcess(int(pid))
		if err != nil {
			fmt.Printf("Failed to find process: %s\n", err)
		} else {
			err := process.Signal(syscall.Signal(0))
			fmt.Printf("process.Signal on pid %d returned: %v\n", pid, err)
		}

	}

	// Curl expected file
	// cmd01 := commandClass{Command: "curl https://images.offensive-security.com/virtual-images/kali-linux-2019.2-vmware-amd64.7z", Path: "/users/Mark/Downloads/kali", Message: "Downloading VHD..."}

	// Copy expecteed file
	cmd02 := commandClass{Command: "/usr/bin/rsync -avh /Users/mark/Downloads/kali /var/tmp", Path: "/var/tmp", Message: "Kali rsync to /var/tmp"}

	// Validate file checkso
	cmd03 := commandClass{Command: "/usr/local/bin/sha256sum -c kali-linux-2019.2-vmware-amd64.7z.txt.sha256sum", Path: "/var/tmp/kali", Message: "Validating Checksum in /var/tmp"}

	// Keka Extract 7zip
	cmd04 := commandClass{Command: "/Applications/Keka.app/Contents/Resources/keka7z x /var/tmp/kali/kali-linux-2019.2-vmware-amd64.7z", Path: "/var/tmp/kali", Message: "Extracting 7z via Keka"}

	// recode VMX encoding
	cmd05 := commandClass{Command: "/usr/local/bin/recode cp1251..utf8 /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp/kali/Kali-Linux-2019.2-vmware-amd64", Message: "VMX change encoding"}

	// VMW vmrun upgrade
	cmd06 := commandClass{Command: "vmrun upgradevm /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp/", Message: "VMW vmrun upgradevm"}

	// VMX change encoding to UTF-8
	cmd07 := commandClass{Command: "/bin/sh host/sh/helper.sh", Path: dir, Message: "Launch helper shell"}

	// VMW vmrun start
	cmd08 := commandClass{Command: "vmrun start /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp/", Message: "VMW vmrun start"}

	// VMW vmrun copyFileFromHostToGuest
	cmd09 := commandClass{Command: "vmrun -gu root -gp toor copyFileFromHostToGuest /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx guest /var/tmp/guest", Path: dir, Message: "VMW vmrun copyFileFromHostToGuest"}

	// VMW vmrun runScriptInGuest
	cmd10 := commandClass{Command: "vmrun -gu root -gp toor runScriptInGuest /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx /bin/bash /var/tmp/guest/sh/bootstrap.sh", Path: "/var/tmp", Message: "VMW vmrun runScriptInGuest"}

	// VMW vmrun shutdown
	cmd11 := commandClass{Command: "vmrun stop /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx soft", Path: "/var/tmp", Message: "VMW vmrun stop soft"}

	// Packer build vagrant
	cmd12 := commandClass{Command: "/usr/local/bin/packer build host/packer/kali.json", Path: dir, Message: "Packer build Vagrant box & vagrant box add local repo"}

	// VMW delete guest
	cmd13 := commandClass{Command: "vmrun deleteVM /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp", Message: "VMW vmrun deleteVM"}

	// Purge source files
	cmd14 := commandClass{Command: "rm -rf output-vmware-vmx packer-cache /var/tmp/kali", Path: dir, Message: "CleanUp Project Workspace"}

	// Initialize Kali for Vagrant
	cmd15 := commandClass{Command: "cp Vagrantfile /var/tmp", Path: dir, Message: "Place Vagrant File"}

	// Initialize Kali for Vagrant
	cmd16 := commandClass{Command: "vagrant up", Path: "/var/tmp", Message: "Launch Vagrant Up"}

	// Initialize Kali for Vagrant
	cmd17 := commandClass{Command: "vagrant ssh kali", Path: "/var/tmp", Message: "Test Vagrant Connection"}

	// Initialize Kali for Vagrant
	cmd18 := commandClass{Command: "vagrant global-status", Path: "/var/tmp", Message: "Vagrant Global Status"}

	// Initialize Kali for Vagrant
	cmd19 := commandClass{Command: "ansible-playbook host/ansible/site.yml", Path: dir, Message: "Run iperf3 test & diff netstat before & after"}

	// Initialize Kali for Vagrant
	cmd20 := commandClass{Command: "sleep 5", Path: "/var/tmp", Message: "Everybody Take Five"}

	// Initialize Kali for Vagrant
	cmd21 := commandClass{Command: "sleep 5", Path: "/var/tmp", Message: "Everybody Take Five"}

	// Initialize Kali for Vagrant
	cmd22 := commandClass{Command: "sleep 5", Path: "/var/tmp", Message: "Everybody Take Five"}

	// Initialize Kali for Vagrant
	cmd23 := commandClass{Command: "sleep 5", Path: "/var/tmp", Message: "Everybody Take Five"}

	// Initialize Kali for Vagrant
	cmd24 := commandClass{Command: "sleep 5", Path: "/var/tmp", Message: "Everybody Take Five"}
	// executor(cmd01)
	executor(cmd02)
	executor(cmd03)
	executor(cmd04)
	executor(cmd05)
	executor(cmd06)
	executor(cmd07)
	executor(cmd08)
	executor(cmd09)
	executor(cmd10)
	executor(cmd11)
	executor(cmd12)
	executor(cmd13)
	executor(cmd14)
	executor(cmd15)
	executor(cmd16)
	executor(cmd17)
	executor(cmd18)
	executor(cmd19)
	executor(cmd20)
	executor(cmd21)
	executor(cmd22)
	executor(cmd23)
	executor(cmd24)

}
