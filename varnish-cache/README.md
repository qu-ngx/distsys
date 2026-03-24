```
kubectl create configmap varnish-config --from-file=default.vcl
kubectl create -f varnish-deploy.yaml
kubectl create -f varnish-service.yaml
```

