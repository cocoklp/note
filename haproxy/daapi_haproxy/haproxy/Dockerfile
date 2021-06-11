FROM centos:7.9.2009


#RUN yum update -y \
#     && yum install vim -y
WORKDIR /app

COPY haproxy /app
COPY haproxy.cfg /app
COPY restart.sh /app
#COPY haproxy.service /usr/lib/systemd/system
COPY start_dpapi_haproxy.sh /app

#RUN systemctl enable haproxy.service
RUN chmod 777 /app/start_dpapi_haproxy.sh \
    && chmod 777 /app/haproxy \
    && chmod 777 /app/restart.sh \
    && mkdir -p /app/logs  \
    && mkdir -p /apps/haproxy
RUN strip /app/haproxy

ENTRYPOINT ["/app/start_dpapi_haproxy.sh"]
