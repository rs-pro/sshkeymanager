# DO NOT use linuxserver/openssh-server - it starts ssh as a non-root user
# https://docs.docker.com/engine/examples/running_ssh_service/
FROM debian:buster

RUN apt-get update && apt-get install -y openssh-server apt-utils sudo

RUN mkdir /var/run/sshd

RUN adduser --disabled-password --gecos "" test

COPY ./testdata/id_rsa.pub /root/.ssh/authorized_keys
COPY ./testdata/id_rsa.pub /home/test/.ssh/authorized_keys

RUN sed -i 's/#*Port 22/Port 2222/g' /etc/ssh/sshd_config
#RUN sed -i 's/#*PermitRootLogin prohibit-password/PermitRootLogin yes/g' /etc/ssh/sshd_config

# SSH login fix. Otherwise user is kicked off after login
RUN sed -i 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' /etc/pam.d/sshd

RUN echo 'test ALL=(ALL) NOPASSWD:ALL' > /etc/sudoers

RUN mkdir /data

EXPOSE 2222
CMD ["/usr/sbin/sshd", "-D", "-e"]
