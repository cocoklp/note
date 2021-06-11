#!/bin/bash
ps -ef|grep haproxy|grep -v grep|grep -v api|cut -c 9-15|xargs kill -9
/app/haproxy -f /app/haproxy.cfg  2>&1
