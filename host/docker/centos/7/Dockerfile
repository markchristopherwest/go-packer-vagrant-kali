FROM centos:7
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
VOLUME [ "/sys/fs/cgroup" ]
CMD ["/usr/sbin/init"]
RUN yum clean all \
    && yum update -y \
    && yum install -y curl gcc git libffi-devel openssl-devel textio vim wget zlib-devel \
    && yum group install -y "Development Tools" \
    && cd /var/tmp \
    && wget https://www.python.org/ftp/python/3.7.4/Python-3.7.4.tgz \
    && tar -xvf Python-3.7.4.tgz \
    && cd /var/tmp/Python-3.7.4 \
    && ./configure \
    && make -j 2 \
    && make altinstall \
    && cd /usr/local/bin \
    && curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py \
    && python get-pip.py \
    && pip install ansible docker setuptools wheel \
    && yum clean all
VOLUME [ "/sys/fs/cgroup" ]
CMD ["/usr/sbin/init"]