# Download Base Image
FROM ubuntu:17.10
 
# Avoid Interactive Shell
ARG DEBIAN_FRONTEND=noninteractive

# Update Software Repository
RUN cp /etc/apt/sources.list /etc/apt/sources.bkup \
  && sed -i -r 's/([a-z]{2}.)?archive.ubuntu.com/old-releases.ubuntu.com/g' /etc/apt/sources.list \
  && sed -i -r 's/security.ubuntu.com/old-releases.ubuntu.com/g' /etc/apt/sources.list \
  && apt-get update \
  && apt-get clean all \
  && apt-get install -y curl g++ git build-essential coreutils g++ git  \
  libcairo2-dev libc-dev libffi-dev libgirepository1.0-dev libssl-dev libvirt-dev \
  libssl-dev openssl vim wget zlib1g \
  && apt-get remove python3 -y \
  && apt-get purge -y \
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