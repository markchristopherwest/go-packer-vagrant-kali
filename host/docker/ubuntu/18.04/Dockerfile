# Download Base Image
FROM ubuntu:18.04
 
# Avoid Interactive Shell
ARG DEBIAN_FRONTEND=noninteractive

# Update Software Repository
RUN apt-get update \
  && apt-get install -y apt-utils aptitude build-essential coreutils g++ git keepalived \
  libcairo2-dev libgirepository1.0-dev libssl-dev libvirt-dev \
  linux-virtual net-tools python3-apt python3-pip python3-dev pkg-config systemd tox vim \
  && cd /usr/local/bin \
  && ln -s /usr/bin/python3 python \
  && pip3 install --upgrade pip \
  && apt-get remove python2.7 -y \
  && apt-get autoclean \
  && apt-get autoremove -y \
  && cd /var/tmp \
  && wget https://www.python.org/ftp/python/3.7.4/Python-3.7.4.tgz \
  && tar -xvf Python-3.7.4.tgz \
  && cd /var/tmp/Python-3.7.4 \
  && ./configure \
  && make -j 2 \
  && make altinstall \
  && cd /usr/local/bin \
  && ln -s /usr/bin/python3 python \
  && curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py \
  && python get-pip.py \
  && pip install ansible docker molecule setuptools wheel \
  && apt-get purge -y \
  && apt-get clean all

# Define Environment Variables
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8