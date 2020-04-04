# Sample implementation of Cluster-Api Bootstrap Provider MicroK8s

Note: This is in super alpha state. Initial implementation 

Using DigitalOcean Infrastructure Provisioner

All samples in [samples](config/samples)

To create a cluster in DIgitalOcean, follow the documentation in https://github.com/kubernetes-sigs/cluster-api-provider-digitalocean/blob/master/docs/getting-started.md

Once the samples are generated apply the `cluster.yaml`

Sample `Microk8sConfig` Control Plane Bootstrapping

```yaml

apiVersion: bootstrap.cluster.x-k8s.io/v1alpha1
kind: Microk8sConfig
metadata:
  name: capdo-microk8s-controlplane-0
spec:
  # Add fields here
  token: abc.123488779
  channel: 1.17/stable

```

* `Channel` : The MicroK8s Channel you want to use to bootstrap the cluster.
* `Token`: The bootstrap token that you can use so that you can join a worker node.

```yaml
apiVersion: bootstrap.cluster.x-k8s.io/v1alpha1
kind: Microk8sConfigTemplate
metadata:
  name: capdo-microk8s-worker-0
spec:
  template:
    spec:
      # Add fields here
      token: abc.123488779
      channel: 1.17/stable
      controlPlaneHost: 128.199.128.28
```

* `controlPlaneHost`: The IP address of the Control Plane.  For now, need to specify.

> **_NOTE:_**  Need to wait for the control plane to be up and running before you can start the provisioning of the worker node.
