#!/bin/bash
etcd --log-level 'info' --name infra1 --data-dir infra1 \
--listen-client-urls https://0.0.0.0:2379 \
--advertise-client-urls https://0.0.0.0:2379 \
--listen-peer-urls https://0.0.0.0:12380 \
--initial-advertise-peer-urls https://127.0.0.1:12380 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster-state new \
--cert-file=/root/cfssl/server.pem \
--key-file=/root/cfssl/server-key.pem \
--peer-client-cert-auth \
--peer-trusted-ca-file=/root/cfssl/ca.pem \
--peer-cert-file=/root/cfssl/member1.pem \
--peer-key-file=/root/cfssl/member1-key.pem \
--discovery https://discovery.etcd.io/d49e6153e1a73157e3251e904adb8a72