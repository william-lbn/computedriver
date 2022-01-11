#!/bin/bash
echo "generate kubelet client certs : kubelet-$1.pem"

cd /etc/kubernetes/pki

mkdir nodecerts
cd ./nodecerts

rm -rf kubelet-$1*

openssl genrsa -out kubelet-$1.key 2048

openssl req -new -key kubelet-$1.key -subj "/CN=system:node:$1/O=system:nodes" -out kubelet-$1.csr

openssl x509 -req \
  -CA /etc/kubernetes/pki/ca.crt \
  -CAkey /etc/kubernetes/pki/ca.key \
  -days 3650 \
  -in kubelet-$1.csr \
  -CAcreateserial \
  -extensions v3_req_client \
  -out kubelet-$1.crt

cat kubelet-$1.key kubelet-$1.crt >kubelet-$1.pem

apt-get install -y expect

dir=$(pwd)

passwd='PmL$%*!Fly1314'


/usr/bin/expect<<EOF
set time 50
spawn scp /etc/kubernetes/pki/ca.crt root@$2:/etc/kubernetes/pki/
expect {
"*yes/no" { send "yes\r"; exp_continue }
"*password:" { send "$passwd\r" }
}
expect eof
EOF

/usr/bin/expect<<EOF
set time 50
spawn scp ${dir}/kubelet-$1.pem  root@$2:/var/lib/kubelet/pki/
expect {
"*yes/no" { send "yes\r"; exp_continue }
"*password:" { send "$passwd\r" }
}
expect eof
EOF


