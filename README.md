# Go-Packer-Vagrant-Kali-Vmw

A Vagrant Box for Kali a la VMware Provider

## Foreward

Many people use Kali.  Many people use VMware.  Many people use Vagrant.  Kali creates VMware images in 7zip.  Vagrant Cloud does not contain very many current Kali images.  Wouldn't it be great if everyone on my team could spin up a Vagrant environment that included Kali for testing?

### Preflight Checklist

You should have the following things installed:

```
VMware Fusion (developed & tested using version 11)
Go (you can install using homebrew)
Packer (you can install using homebrew)
Vagrant (you can install using homebrew)
Keka (you can swap this for 7zip & install either using homebrew)
vmrun (comes with VMware so youa re good to if you've got desktop aka fusion)

```

### Installation

"It's always best to start at the beginning..." Glinda the Good Witch

Install Prerequisites

```
brew install vmware-fusion go packer vagrant keka git tree diff ssh-copy-id
```

Clone Repository

```
git clone https://github.com/markchristopherwest/go-packer-vagrant-kali-vmw
```

Go Live

```
go run main.go
```

You can change or uncomment/toggle cmd0 to meet your needs (copy locally vs download, etc.)

## Testing your install

After you run the go file, you'll want to validate that the Vagrant box was created & installed into your local boxes.  To do 

### Vagrant Box List

You might want to check the status of your vagrant boxes by running:

```
➜  Downloads vagrant box list
generic/rhel8      (vmware_desktop, 1.9.16)
generic/ubuntu1604 (vmware_desktop, 1.9.16)
generic/ubuntu1904 (vmware_desktop, 1.9.16)
kali-linux         (vmware_desktop, 0)
```



### Vagrant Up

You'll want to modify your vagrantfile to include the parameters for the new box i.e.:

```
change the default VM in the vagrant file to your new box name or set as default
```

## Deploying

Customized your Vagrant files to create multi-node environments.

## Created With

* [Ansible](https://github.com/ansible/ansible.git) - The Vagrant provisioner (besides shell)
* [Go](http://www.golang.org/) - The higher level language used
* [Packer](https://github.com/hashicorp/packer.git) - Source creation
* [Vagrant](https://github.com/hashicorp/vagrant.git) - Target manifest


## Versioning

We use [VagrantCloud](https://vagrantcloud.com/) for versioning. For the versions available, see the [tags on this repository](https://github.com/markchristopherwest/go-packer-vagrant-kali). 

## Authors

* **Mark Christopher West** - *Initial work* - [MarkChristopherWest](https://github.com/markchristopherwest)

See also the list of [contributors](https://github.com/markchristopherwest/go-packer-vagrant-kali/contributors) who participated in this project.

## License

This project is licensed via MIT License - see our [LICENSE.md](LICENSE.md) for more...

## Acknowledgments

* My Family & their bountiful patience with me
* All my coworkers & friends at Cisco who have been there for 20 years
* All my coworkers & friends at Dell-EMC who taught me the bowels of ESX, VNX & VMAX
* All my coworkers & friends at RedHat who taught me the value of Open Source
