#!/bin/sh

# mod vmx config (because escaping string literals for sed in Go's commandExec is not fun)
/usr/local/bin/gsed -i 's/.encoding.*/.encoding = \"UTF-8\"/g' /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx
/usr/local/bin/gsed -i 's/uuid.bios.*/msg.autoAnswer = \"TRUE\"/g' /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx
/usr/local/bin/gsed -i 's/uuid.location.*/uuid.action = \"CREATE\"/g' /var/tmp/kali/Kali-Linux-2019.2-vmware-amd64/Kali-Linux-2019.2-vmware-amd64.vmx