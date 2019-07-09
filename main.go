package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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

	// // VMX change encoding to UTF-8
	cmd07 := commandClass{Command: "/bin/sh helper.sh", Path: dir, Message: "VMX sed .encoding"}

	// VMW vmrun start
	cmd08 := commandClass{Command: "vmrun start /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp/", Message: "VMW vmrun start"}

	// VMW vmrun copyFileFromHostToGuest
	cmd09 := commandClax	ss{Command: "vmrun -gu root -gp toor copyFileFromHostToGuest /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx bootstrap.sh /var/tmp/bootstrap.sh", Path: dir, Message: "VMW vmrun copyFileFromHostToGuest"}

	// VMW vmrun runScriptInGuest
	cmd10 := commandClass{Command: "vmrun -gu root -gp toor runScriptInGuest /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx /bin/bash /var/tmp/bootstrap.sh", Path: "/var/tmp", Message: "VMW vmrun runScriptInGuest"}

	// VMW vmrun shutdown
	cmd11 := commandClass{Command: "vmrun stop /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx soft", Path: "/var/tmp", Message: "VMW vmrun stop soft"}

	// Packer build vagrant
	cmd12 := commandClass{Command: "/usr/local/bin/packer build packer.json", Path: dir, Message: "Packer build Vagrant box & vagrant box add local repo"}

	// VMW delete guest
	cmd13 := commandClass{Command: "vmrun deleteVM /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp", Message: "VMW vmrun deleteVM"}

	// Purge source files
	cmd14 := commandClass{Command: "rm -rf /var/tmp/kali/", Path: dir, Message: "CleanUp Project Workspace"}

	// Initialize Kali for Vagrant
	cmd15 := commandClass{Command: "vagrant init --output /var/tmp/kali/Vagrantfile --box-version kali-linux2019.2", Path: "/var/tmp", Message: "Create Vagrant File"}

	// Initialize Kali for Vagrant
	cmd16 := commandClass{Command: "vagrant up", Path: "/var/tmp/kali", Message: "Launch Vagrant Up"}

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

}
