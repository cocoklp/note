#!/bin/sh
	cd haproxy
	make
	
	cd ../
	
	cd dataplaneapi
	make -f Makefile.docker
