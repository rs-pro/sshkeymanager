FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y openssh-server sudo

RUN mkdir /var/run/sshd

RUN adduser --disabled-password --gecos "" test && \
    echo 'test ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

RUN mkdir -p /root/.ssh /home/test/.ssh && \
    chmod 700 /root/.ssh /home/test/.ssh

COPY ./testdata/id_rsa.pub /root/.ssh/authorized_keys
COPY ./testdata/id_rsa.pub /home/test/.ssh/authorized_keys

RUN chown -R root:root /root/.ssh && \
    chown -R test:test /home/test/.ssh && \
    chmod 600 /root/.ssh/authorized_keys /home/test/.ssh/authorized_keys

RUN sed -i 's/^#*Port 22$/Port 2222/' /etc/ssh/sshd_config && \
    sed -i 's/^#*PermitRootLogin.*/PermitRootLogin yes/' /etc/ssh/sshd_config && \
    sed -i 's/^#*PubkeyAuthentication.*/PubkeyAuthentication yes/' /etc/ssh/sshd_config && \
    sed -i 's/^#*PasswordAuthentication.*/PasswordAuthentication no/' /etc/ssh/sshd_config

RUN sed -i 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' /etc/pam.d/sshd

EXPOSE 2222

CMD ["/usr/sbin/sshd", "-D", "-e"]
