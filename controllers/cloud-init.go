package controllers

const (
	controllerInit = `
#cloud-config
package_update: true
snap:
  commands:
    00: snap install microk8s --channel=1.18/stable --classic
runcmd:
  - 'until /snap/bin/microk8s.status --wait-ready; do sleep 1; echo "waiting for status.."; done'
  - 'touch /var/snap/microk8s/current/credentials/cluster-tokens.txt'
  - 'echo "abc.99998888833346" >> /var/snap/microk8s/current/credentials/cluster-tokens.txt'
`

	workerInit = `
#cloud-config
package_update: true
snap:
  commands:
    00: snap install microk8s --channel=1.18/stable --classic
runcmd:
  - 'until /snap/bin/microk8s.status --wait-ready; do sleep 1; echo "waiting for status.."; done'
  - '/snap/bin/microk8s.join ${digitalocean_droplet.microk8s-controller.ipv4_address_private}:25000/${var.cluster-token}-${count.index}'
`
)
