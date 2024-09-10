#!/bin/sh

# Apply Namespace Configuration
kubectl apply -f infra/k8s/namespaces/
if [ $? -ne 0 ]; then
    echo "Failed to apply namespace configuration"
    exit 1
fi

# Apply Security Configuration
kubectl apply -f infra/k8s/security.yaml
if [ $? -ne 0 ]; then
    echo "Failed to apply security configuration"
    exit 1
fi

# Apply Metrics Server Configuration
kubectl apply -f infra/k8s/metrics-server.yml
if [ $? -ne 0 ]; then
    echo "Failed to apply Metrics Server configuration"
    exit 1
fi

# Apply ConfigMap app Configuration
kubectl apply -f infra/k8s/config-maps/config-map-app.yml
if [ $? -ne 0 ]; then
    echo "Failed to apply ConfigMap app configuration"
    exit 1
fi

# Apply ConfigMap postgres Configuration
kubectl apply -f infra/k8s/config-maps/config-map-postgres.yml
if [ $? -ne 0 ]; then
    echo "Failed to apply ConfigMap postgres configuration"
    exit 1
fi

# Apply StatefulSet for PostgreSQL
kubectl apply -f infra/k8s/statefulsets/postgres-statefulset.yaml
if [ $? -ne 0 ]; then
    echo "Failed to apply StatefulSet configuration"
    exit 1
fi

# Apply Deployment Configuration
kubectl apply -f infra/k8s/deployments/go-mail-api-deployment.yaml
if [ $? -ne 0 ]; then
    echo "Failed to apply deployment configuration"
    exit 1
fi

# Apply Horizontal Pod Autoscaler Configuration
kubectl apply -f infra/k8s/hpa/go-mail-api-hpa.yml
if [ $? -ne 0 ]; then
    echo "Failed to apply HPA configuration"
    exit 1
fi

# Check the status of the deployment
kubectl rollout status deployment/go-mail-api -n go-mail-api-namespace
if [ $? -ne 0 ]; then
    echo "Deployment rollout failed"
    exit 1
fi

# Forward local port 8080 to the go-mail-api-service's port 80 in the go-mail-api-namespace.
# This allows us to access the service locally by going to http://localhost:8080.
kubectl port-forward service/go-mail-api-service -n go-mail-api-namespace 8080:80

echo "Kubernetes configurations applied successfully"
