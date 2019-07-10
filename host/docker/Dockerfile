#Download base image ubuntu 16.04
FROM ubuntu:18.04
 
#Avoid interaction
ARG DEBIAN_FRONTEND=noninteractive

# Update Software repository
RUN apt-get update \
  && apt-get install -y apt-utils aptitude build-essential coreutils curl g++ git haproxy iperf3 keepalived \
  libcairo2-dev libgirepository1.0-dev libssl-dev libvirt-dev \
  linux-virtual net-tools python3-apt python3-pip python3-dev pkg-config systemd tox wget vim \
  && cd /usr/local/bin \
  && ln -s /usr/bin/python3 python \
  && pip3 install --upgrade pip \
  && apt-get autoclean \
  && apt-get autoremove -y \
  && pip install -U ansible docker molecule setuptools wheel

#Define the ENV variable
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8