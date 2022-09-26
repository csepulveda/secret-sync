# Secret Sync
This tool copy secrets from AWS SecretsManager to Kubernetes Secrets.
Using this tools its not necesary to use Kubernetes Secrets Store CSI Driver if you only need send the entier secret from AWS to Kubernetes Secret.
The secrets created has the label `created_by=secret-sync` if one of those secrets are removed from de secret-sync configuration, the secret will be removed.

## How to run.
* attach a role to read the aws secrets
* attach a serviceaccount to allow manage secrets
* set INTERVAR env_var

## pod config example:
```
apiVersion: v1
kind: Pod
metadata:
  name: sync
spec:
  serviceAccountName: secrets-admin
  containers:
  - name: sync
    imagePullPolicy: Never
    image: sync:latest
    volumeMounts:
    - name: config-volume
      mountPath: /etc/config
    env:
      - name: AWS_ACCESS_KEY_ID
        value: AKIxxxxx
      - name: AWS_SECRET_ACCESS_KEY
        value: z/MXxxxxx
      - name: AWS_DEFAULT_REGION
        value: us-west-1
      - name: INTERVAL
        value: "120"
  volumes:
    - name: config-volume
      configMap:
        name: secret-sync-config
```

## Example:

### configmap:
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: secret-sync-config
  namespace: default
data:
  config.json: |
    {
        "secrets": [
            {
                "provider":"aws",
                "source":"dev-new-example-secret",
                "dest": "dev-new-example-secret",
                "namespace": "default"
            }
        ]
    }
```

### aws secret:
```
aws secretsmanager get-secret-value --secret-id dev-new-example-secret --region us-west-1 --output json
{
    "ARN": "arn:aws:secretsmanager:us-west-1:xxxxx:secret:dev-new-example-secret-P3SZYf",
    "Name": "dev-new-example-secret",
    "VersionId": "EF6ACFBF-9087-4F63-96ED-F2644F3EF2A0",
    "SecretString": "{\"value1\":\"data1\",\"value2\":\"data2\",\"value3\":\"data3\",\"value4\":\"data4\"}",
    "VersionStages": [
        "AWSCURRENT"
    ],
    "CreatedDate": "2022-07-15T09:32:23.744000-04:00"
}
```

### kubernetes secret generated:
```
kubectl get secrets/dev-new-example-secret -o yaml 
apiVersion: v1
data:
  value1: ZGF0YTE=
  value2: ZGF0YTI=
  value3: ZGF0YTM=
  value4: ZGF0YTQ=
kind: Secret
metadata:
  creationTimestamp: "2022-09-24T03:56:26Z"
  labels:
    created_by: secret-sync
  name: dev-new-example-secret
  namespace: default
  resourceVersion: "19440"
  uid: fa8200a2-ca8f-4f2c-b961-32724f65f57c
type: Opaque
```

## TODO
* Create github actions ci/cd
* Create Test
* Create Helmchart
* evaluate add supports to other kind of secrets (azure, gpc, vault, etc)