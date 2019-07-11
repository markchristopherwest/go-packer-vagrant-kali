FROM rhel:7
ENV container docker
RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == \
systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;
RUN yum update -y \
    && yum upgrade -y \
    && yum groupinstall -y "Development Tools" \
    && yum install -y yum-utils coreutils curl g++ git keepalived make \
    cairo cairo-devel mesa-libgl openssl-devel libvirt-dev \
    qemu-kvm qemu-img virt-manager libvirt libvirt-client libffi-dev virt-install virt-viewer bridge-utils \
    linux-virtual net-tools python3-dev python3-pip systemd vim  \
    && curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
    && cd /usr/local/bin \
    && ln -s /usr/bin/python3 python \
    && pip3 install --upgrade pip \
    && pip install ansible wheel

VOLUME [ "/sys/fs/cgroup" ]
CMD ["/usr/sbin/init"]


# Define Environment Variables
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8