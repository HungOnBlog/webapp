# Bootstrap the Canary Deployment with Linked and Flagger

This guide will show you how to bootstrap a canary deployment for your Kubernetes cluster using Flagger and Linkerd.

## Prerequisites

Flagger requires a Kubernetes cluster **v1.11** or newer and Linkerd **v2.4** or newer.

If you don't have a Kubernetes cluster, I recommend using [KinD](https://kind.sigs.k8s.io/) to create one.

Step 1: Install Linkerd on your cluster:

```bash
./install_linkerd.sh
```

Step 2: Install Flagger and the Prometheus add-on:

```bash
./install_flagger.sh
```

Step 3: Create the the canary testing namespace:

```bash
cd kubectl apply -f k8s/webapp_namespace.yaml
```

Step 4: Deploy the webapp service:

```bash
kubectl apply -f k8s/webapp
```

**NOTE**: The webapp service must be meshed with Linkerd using the `linkerd.io/inject: enabled` annotation.

Step 5: Deploy Locust:

```bash
kubectl apply -f k8s/locust
kubectl port-forward -n webapp service/locust-master-ui 30627:8089
```

Step 5.1: Create traffic to the webapp
Open the Locust UI at <http://localhost:30627> and start a 5 minute test with 2 users.

And choose 100 concurrent users and 100 spawn rate.

Then, click the `Start swarming` button.

Open new `terminal`

Step 6: Create `Canary` custom resource:

```yaml
apiVersion: flagger.app/v1beta1
kind: Canary
metadata:
  name: webapp
  namespace: webapp
spec:
  # deployment reference
  targetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: webapp
  # HPA reference (optional)
  autoscalerRef:
    apiVersion: autoscaling/v2beta2
    kind: HorizontalPodAutoscaler
    name: webapp
  # the maximum time in seconds for the canary deployment
  # to make progress before it is rollback (default 600s)
  progressDeadlineSeconds: 60
  service:
    # ClusterIP port number
    port: 3000
    # container port number or name (optional)
    targetPort: 3000
  analysis:
    # schedule interval (default 60s)
    interval: 30s
    # max number of failed metric checks before rollback
    threshold: 5
    # max traffic percentage routed to canary
    # percentage (0-100)
    maxWeight: 50
    # canary increment step
    # percentage (0-100)
    stepWeight: 5
    # Linkerd Prometheus checks
    metrics:
    - name: request-success-rate
      # minimum req success rate (non 5xx responses)
      # percentage (0-100)
      thresholdRange:
        min: 99
      interval: 1m
    - name: request-duration
      # maximum req duration P99
      # milliseconds
      thresholdRange:
        max: 500
      interval: 30s
```

Step 7: Deploy the canary:

```bash
kubectl apply -f k8s/canary.yaml
```

After a few seconds Flagger will create the primary and canary deployments, service and HPA then it will start the canary analysis.

Step 8: Trigger a canary deployment by updating the container image:

```bash
kubectl -n webapp set image deployment/webapp webapp=hungtpplay/webapp:1.0.6
```

Step 9: Watch the canary deployment progress:

```bash
watch kubectl get canaries --all-namespaces
```
