FROM kinogmt/centos-ssh

ADD misc/testdata/id_rsa /root/.ssh
ADD misc/testdata/id_rsa.pub /root/.ssh

RUN touch /root/.ssh/authorized_keys && \
    chmod 600 /root/.ssh/authorized_keys && \
    cat /root/.ssh/id_rsa.pub > /root/.ssh/authorized_keys
