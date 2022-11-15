# Secret Sync
This tool copy secrets from AWS SecretsManager to Kubernetes Secrets.
Using this tools its not necesary to use Kubernetes Secrets Store CSI Driver if you only need send the entier secret from AWS to Kubernetes Secret.
The secrets created has the label `created_by=secret-sync` if one of those secrets are removed from de secret-sync configuration, the secret will be removed.

## How to run.
* attach a role to read the aws secrets
* attach a serviceaccount to allow manage secrets
* set INTERVAR env_var

### run example:
local values file:
```
cat ~/values.yaml
env:
  "AWS_ACCESS_KEY_ID": "xxxxxxx"
  "AWS_SECRET_ACCESS_KEY": "xxxxxxx"
  "AWS_DEFAULT_REGION": "us-west-2"

secrets:
  - provider: aws
    source: dev-new-example-secret
    dest: dev-new-example-secret
    namespace: default
```

add repo
```
(⎈ |N/A:N/A)➜  secret_sync git:(dev) helm repo add secret-sync https://csepulveda.github.io/secret-sync/  
(⎈ |N/A:N/A)➜  ~ helm repo add secret-sync https://csepulveda.github.io/secret-sync/
"secret-sync" has been added to your repositories
(⎈ |N/A:N/A)➜  ~ helm repo update secret-sync                                       
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "secret-sync" chart repository
Update Complete. ⎈Happy Helming!⎈
```
Start minukube cluster
```
(⎈ |N/A:N/A)➜  ~ minikube start       
😄  minikube v1.27.0 on Darwin 12.6 (arm64)
    ▪ MINIKUBE_ACTIVE_DOCKERD=minikube
❗  Kubernetes 1.25.0 has a known issue with resolv.conf. minikube is using a workaround that should work for most use cases.
❗  For more information, see: https://github.com/kubernetes/kubernetes/issues/112135
✨  Automatically selected the docker driver. Other choices: parallels, ssh, qemu2 (experimental)
📌  Using Docker Desktop driver with root privileges
👍  Starting control plane node minikube in cluster minikube
🚜  Pulling base image ...
🔥  Creating docker container (CPUs=2, Memory=5891MB) ...
🐳  Preparing Kubernetes v1.25.0 on Docker 20.10.17 ...
    ▪ Generating certificates and keys ...
    ▪ Booting up control plane ...
    ▪ Configuring RBAC rules ...
🔎  Verifying Kubernetes components...
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
🌟  Enabled addons: storage-provisioner, default-storageclass
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default
```
Install helm chart
```                                             
(⎈ |minikube:default)➜  ~ helm upgrade -i secret-sync secret-sync/secret-sync  -f ~/values.yaml
Release "secret-sync" does not exist. Installing it now.
NAME: secret-sync
LAST DEPLOYED: Mon Sep 26 11:46:42 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

check logs and secret
```
(⎈ |minikube:default)➜  ~ kubectl logs deployment/secret-sync  
2022/09/26 14:46:50 running every 120 seconds
2022/09/26 14:46:50 Sync 1 of 1 secrets
(⎈ |minikube:default)➜  ~ kubectl describe secrets/dev-new-example-secret                      
Name:         dev-new-example-secret
Namespace:    default
Labels:       created_by=secret-sync
Annotations:  <none>

Type:  Opaque

Data
====
value1:  5 bytes
value2:  5 bytes
value3:  5 bytes
value4:  5 bytes
```


## TODO
* Create Tests
* Evaluate add supports to other kind of secrets (azure, gpc, vault, etc)