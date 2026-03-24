# STEP1: The first step is to upload your certificate as a secret to Kubernetes:
```
kubectl create secret tls ssl --cert=server.crt --key=server.key
```

# STEP2: As with Varnish, you need to transform this into a ConfigMap object
```
kubectl create configmap nginx-conf --from-file=nginx.conf
```


# STEP3:To create the replicated nginx servers, you use: 
```
kubectl create -f nginx-deploy.yaml
```

# To create this load-balancing service run:
```
kubectl create -f nginx-service.yaml
```

