#!/usr/bin/bash

# start haproxy in backend

/app/haproxy -f /app/haproxy.cfg  2>&1

# start dataplaneapi in frontend
#  --port 5555 -b /usr/sbin/haproxy -c /etc/haproxy/haproxy.cfg
/app/dataplaneapi --host 0.0.0.0 --port 5555 -b /app/haproxy -c /app/haproxy.cfg -d 5 -u dataplaneapi -r "/app/haproxy -f /app/haproxy.cfg -sf"  -s "/app/restart.sh" --log-to=file --log-file=/app/logs/dataplaneapi.log --log-level=warning
