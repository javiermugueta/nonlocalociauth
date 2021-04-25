FROM oraclelinux:7-slim
RUN  curl -o /etc/yum.repos.d/public-yum-ol7.repo https://yum.oracle.com/public-yum-ol7.repo && \
     yum-config-manager --enable ol7_developer_golang111 &&\
     yum -y install deltarpm git golang make vi  && \
     rm -rf /var/cache/yum
#
WORKDIR /nonlocalociauth
ADD start.sh .
RUN chmod 700 start.sh
ADD nonlocalociauth.go .
CMD ["/nonlocalociauth/start.sh"]

