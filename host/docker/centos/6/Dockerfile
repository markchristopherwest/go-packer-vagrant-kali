FROM centos:6
ENV container docker
VOLUME [ "/sys/fs/cgroup" ]
CMD ["/usr/sbin/init"]
RUN yum update -y \
    && yum groupinstall -y "Development Tools" \
    && yum install -y yum-utils coreutils curl g++ git keepalived make \
    cairo cairo-devel mesa-libgl openssl-devel libvirt-dev \
    qemu-kvm qemu-img virt-manager libvirt libvirt-client libffi-dev virt-install virt-viewer bridge-utils \
    linux-virtual net-tools python3-dev python3-pip systemd vim wget \
    && curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
    && cd /var/tmp \
    && wget https://www.python.org/ftp/python/3.7.4/Python-3.7.4.tgz \
    && tar -xvf Python-3.7.4.tgz \
    && cd /var/tmp/Python-3.7.4 \
    && ./configure --enable-optimizations \
    && make -j 4 \
    && make altinstall \
    && cd /usr/local/bin \
    && ln -s /usr/bin/python3 python \
    && curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py \
    && python get-pip.py \
    && pip install --upgrade pip \
    && pip install ansible wheel
VOLUME [ "/sys/fs/cgroup" ]
CMD ["/usr/sbin/init"]


# Define Environment Variables
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8