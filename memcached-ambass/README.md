```
kubectl create -f memecached-shards.yaml
kubectl create -f memcached-service.yaml
kubectl create configmap --from-file=nutcracker.yaml twem-config
kubectl create -f memcached-ambassador-pod.yaml
```
