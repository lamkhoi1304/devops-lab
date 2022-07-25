# Getting Started with Vagrant

## Vagrant Architecture
![Vagrant Architecture](vagrant.png?raw=true "vagrant.png")

## Prerequisites
* A public key in .ssh/id_rsa.pub in your home directory
* Directly download [Virtualbox](https://www.virtualbox.org/) and install
    ```
    sudo apt install virtualbox
    ```
* Vagrant
    ```
    wget https://releases.hashicorp.com/vagrant/2.2.19/vagrant_2.2.19_x86_64.deb
    sudo apt install ./vagrant_2.2.19_x86_64.deb
    ```

## Provision VMs
```
git clone https://github.com/lamkhoi1304/devops-lab.git

cd devops-lab/vagrant/singleVM
vagrant up
ssh ci@192.168.56.11

cd devops-lab/vagrant/multiVM
vagrant up
ssh ci@192.168.56.11
ssh ci@192.168.56.12
```

## Vagrant Cheat Sheet
- `vagrant up` -- starts vagrant
- `vagrant ssh <boxname>` -- connects to machine via SSH
- `vagrant halt` -- stops the vagrant machine
- `vagrant destroy` -- stops and deletes all traces of the vagrant machine
- `vagrant -v` -- get the vagrant version
- `vagrant status` -- outputs status of the vagrant machine
- `vagrant global-status` -- outputs status of all vagrant machines
