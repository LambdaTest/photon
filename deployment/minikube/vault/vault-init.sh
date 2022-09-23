#!/bin/bash

export VAULT_ADDR=http://0.0.0.0:8200
vault login root

# only for minikube
kubectl config use-context minikube

# Create vault role for photon
vault write auth/kubernetes/role/phoenix-photon \
bound_service_account_names=photon \
bound_service_account_namespaces=phoenix \
policies=microservices-kv-read 

vault kv put microservices/phoenix/photon @../../../.ph.json
