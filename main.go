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

	// YMMV
	// cmd00 := commandClass{Command: "curl https://images.offensive-security.com/virtual-images/kali-linux-2019.2-vmware-amd64.7z", Path: "/var/tmp/", Message: "Downloading VHD..."}

	// Copy expected checksum
	cmd01 := commandClass{Command: "cp -rf /Users/mark/Downloads/kali-linux-2019.2-vmware-amd64.7z /var/tmp", Path: "/var/tmp", Message: "Downloading VHD..."}

	// Validate the file
	cmd02 := commandClass{Command: "sha256sum -c kali-linux-2019.2-vmware-amd64.7z.txt.sha256sum", Path: "/var/tmp", Message: "Validating Checksum..."}

	// Keka Extract 7zip
	cmd03 := commandClass{Command: "/Applications/Keka.app/Contents/Resources/keka7z x /var/tmp/kali-linux-2019.2-vmware-amd64.7z", Path: "/var/tmp", Message: "Extracting 7zip..."}

	// VMW vmrun silence (vm package lock)
	cmd04 := commandClass{Command: "rm -f /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx.lck", Path: "/var/tmp", Message: "VMW remove guest package lock"}

	// VMW vmrun silence (vmx hush dialog)
	cmd05 := commandClass{Command: "echo 'msg.autoanswer = \"true\"'$'\r' >> 'Kali-Linux-2019.2-vmware-amd64.vmx'", Path: "/var/tmp/kali-linux-2019-2-vmware-amd64-7z/Kali-Linux-2019.2-vmware-amd64", Message: "VMW append vmx options"}

	// VMW vmrun upgrade
	cmd06 := commandClass{Command: "vmrun upgradevm /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp", Message: "VMW append vmx options"}

	// VMW vmrun start
	cmd07 := commandClass{Command: "vmrun start /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp/", Message: "VMW vmrun start guest"}

	// VMW vmrun copyFileFromHostToGuest
	cmd08 := commandClass{Command: "vmrun -gu root -gp toor copyFileFromHostToGuest /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx bootstrap.sh /var/tmp/bootstrap.sh", Path: dir, Message: "VMW vmrun copyFileFromHostToGuest"}

	// VMW vmrun runScriptInGuest
	cmd09 := commandClass{Command: "vmrun -gu root -gp toor runScriptInGuest /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx /bin/bash /var/tmp/bootstrap.sh", Path: "/var/tmp", Message: "VMW vmrun runScriptInGuest"}

	// VMW vmrun shutdown
	cmd10 := commandClass{Command: "vmrun stop /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx soft", Path: "/var/tmp", Message: "VMW vmrun shutdown"}

	// VMKFSTOOLS defrag disk (to reduce clone time)
	cmd11 := commandClass{Command: "vmware-vdiskmanager -d /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmdk", Path: "/Applications/VMware Fusion.app/Contents/Library", Message: "VMKFSTOOLS defrag disk (to reduce clone time)"}

	// VMKFSTOOLS shrink disk (to reduce disk size)
	cmd12 := commandClass{Command: "vmware-vdiskmanager -k /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmdk", Path: "/Applications/VMware Fusion.app/Contents/Library", Message: "VMKFSTOOLS shrink disk (to reduce disk size)"}

	// Execute packer build
	cmd13 := commandClass{Command: "packer build packer.json", Path: dir, Message: "Execute packer build"}

	// Cleanup source directory
	cmd14 := commandClass{Command: "rm -rf /var/tmp/kali-linux-2019.2-vmware-amd64.7z", Path: "/var/tmp", Message: "Cleanup 7zip cache"}

	// Cleanup source directory
	cmd15 := commandClass{Command: "vmrun deleteVM /var/tmp/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx", Path: "/var/tmp", Message: "Cleanup VMW cache"}

	// Delete source box
	cmd16 := commandClass{Command: "rm -rf packer_cache output-vmware-vmx /var/tmp/Kali-Linux-2019.2-vmware-amd64", Path: dir, Message: "Cleanup Packer & VMW cache"}

	// executor(cmd00)
	executor(cmd01)
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
