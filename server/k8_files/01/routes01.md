# Version 01

Very basic setup with only a few routes

| route    | http method | Request | Response                            | Other info                                  |
| -------- | ----------- | ------- | ----------------------------------- | ------------------------------------------- |
| /        | GET         |         | _simple hello response_             |                                             |
| /pod     | GET         |         | _name of the k8 pod_                |                                             |
| /secrets | GET         |         | _the data in the secrets yaml file_ | Prints out all the env variables to the log |
