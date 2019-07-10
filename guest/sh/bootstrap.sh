#!/bin/sh

# create vagrant user
useradd vagrant -p vagrant
usermod -aG sudo vagrant
mkhomedir_helper vagrant

# install vagrant keys
# https://www.vagrantup.com/docs/boxes/base.html
cd /var/tmp
git clone https://github.com/hashicorp/vagrant.git
mkdir -p /home/vagrant/.ssh
cp -rf /var/tmp/guest/ansible /home/vagrant/ansible
touch -f /home/vagrant/.ssh/authorized_keys
cat /var/tmp/vagrant/keys/vagrant.pub >> /home/vagrant/.ssh/authorized_keys
rm -rf /var/tmp/vagrant
chown vagrant.vagrant /home/vagrant/.ssh
echo "vagrant ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers
sed -i "s/.*PubkeyAuthentication.*/PubkeyAuthentication yes/g" /etc/ssh/sshd_config
sed -i "s/.*PasswordAuthentication.*/PasswordAuthentication yes/g" /etc/ssh/sshd_config

# enable root login for packer
sed -i "s/.*PermitRootLogin.*/PermitRootLogin  yes/g" /etc/ssh/sshd_config

# disable UseDNS for vagrant
sed -i "s/.*UseDNS.*/UseDNS no/g" /etc/ssh/sshd_config

# enable & restart (if exists)
systemctl enable ssh.service
systemctl restart ssh.service

# install preferred provisioner(s)
apt install -y ansible cloud-init

# clean up after
rm -rf /var/tmp/vagrant