#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [ ! -f ./demo.env ] ; then
    echo "You could copy and edit the file:"
    echo "    cp ./hack/compose/demo.env.tmpl ./demo.env"
    exit 1
fi
source ./demo.env
bin/kore-apiserver \
    --kube-api-server http://127.0.0.1:8080 \
    --verbose \
    --admin-pass password \
    --admin-token password \
    --api-public-url http://localhost:10080 \
    --kore-authentication-plugin basicauth \
    --kore-authentication-plugin admintoken \
    --kore-authentication-plugin openid \
    --certificate-authority	hack/ca/ca.pem \
    --certificate-authority-key hack/ca/ca-key.pem \
    --users-db-url 'root:pass@tcp(localhost:3306)/kore?parseTime=true'