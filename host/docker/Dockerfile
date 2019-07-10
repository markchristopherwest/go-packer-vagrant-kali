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
  && apt-get autoclean \
  && apt-get autoremove -y \
  && pip install -U ansible docker molecule setuptools wheel

# Define Environment Variables
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8