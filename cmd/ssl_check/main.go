package main

# Add CA cert bundle to list of trusted CA certs.  For example, you would do this with a golang TLS config by parsing the certificate chain and adding the parsed certificates to the RootCAs field in the tls.Config struct.  https://kubernetes.io/docs/tasks/tls/managing-tls-in-a-cluster/#trusting-tls-in-a-cluster

# You can distribute the CA certificate as a ConfigMap that your pods have access to use.
