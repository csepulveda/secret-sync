# Secret Sync

## Parameters

### Common Parameters

| Name                         | Description                                                                                                         | Value                     |
| ---------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------- |
| `replicaCount`               | Number of secret-sync replicas                                                                                      | `1`                       |
| `image.repository`           | The repository to use for the secret-sync image.                                                                    | `csepulvedab/secret-sync` |
| `image.pullPolicy`           | the pull policy to use for the secret-sync image                                                                    | `IfNotPresent`            |
| `image.tag`                  | The secret-sync image tag. Defaults to the chart's AppVersion                                                       | `""`                      |
| `serviceAccount.annotations` | Annotations for service account. Evaluated as a template. Only used if `create` is `true`.                          | `{}`                      |
| `serviceAccount.create`      | Specifies whether a ServiceAccount should be created                                                                | `true`                    |
| `serviceAccount.name`        | Name of the service account to use. If not set and create is true, a name is generated using the fullname template. | `secret-sync`             |
| `role.annotations`           | Annotations for the tole. Evaluated as a template. Only used if `create` is `true`.                                 | `{}`                      |
| `role.create`                | Specifies whether a role should be created                                                                          | `true`                    |
| `role.name`                  | Name of the role to use. If not set and create is true, a name is generated using the fullname template.            | `secret-sync`             |
| `roleBinding.annotations`    | Annotations for role binding. Evaluated as a template. Only used if `create` is `true`.                             | `{}`                      |
| `roleBinding.create`         | Specifies whether a roleBinding should be created                                                                   | `true`                    |
| `roleBinding.name`           | Name of the role binding to use. If not set and create is true, a name is generated using the fullname template.    | `secret-sync`             |
| `podAnnotations`             | Add extra annotations to the secret-sync pod(s)                                                                     | `{}`                      |
| `podSecurityContext`         | Add extra podSecurityContext to the secret-sync pod(s)                                                              | `{}`                      |
| `securityContext`            | Add extra securityContext to the secret-sync pod(s)                                                                 | `{}`                      |
| `resources.limits`           | The resources limits for the secret-sync container                                                                  | `{}`                      |
| `resources.requests`         | The requested resources for the secret-sync container                                                               | `{}`                      |
| `env`                        | Additional env vars for secret-sync pod(s).                                                                         | `{}`                      |
| `nodeSelector`               | Node labels for pod assignment                                                                                      | `{}`                      |
| `tolerations`                | Tolerations for secret-sync pod assignment.                                                                         | `[]`                      |
| `affinity`                   | Node labels for pod assignment                                                                                      | `{}`                      |


### Secrets Definitions

| Name      | Description          | Value |
| --------- | -------------------- | ----- |
| `secrets` | secret to be synced. | `[]`  |


