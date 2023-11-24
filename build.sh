docker build -t josepht96/scout:latest .
docker push josepht96/scout:latest
kubectl delete -f deployment.yaml
kubectl apply -f deployment.yaml