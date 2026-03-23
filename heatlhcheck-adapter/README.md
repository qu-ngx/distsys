# Instructions:

# STEP 1: Run this command to create the secret in your cluster

```
 kubectl create secret generic mysql-creds \
  --from-literal=username=root \
  --from-literal=password=your-secure-password-here
```

# STEP 2:

# 1. Apply the manifest
```
 kubectl apply -f mysql-adapter.yaml
```

# 2. Watch the status
```
kubectl get pods -w
```

# 3. Check the Adapter's logs to see it connecting to MySQL
```
kubectl logs -f deployment/mysql-monitored -c health-adapter
```
