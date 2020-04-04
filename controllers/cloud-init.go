package controllers

const (
	ControlPlaneInitTemplate = `
#cloud-config
package_update: true
snap:
  commands:
    00: snap install microk8s --channel={{ .Channel }} --classic
runcmd:
  - 'until /snap/bin/microk8s.status --wait-ready; do sleep 1; echo "waiting for status.."; done'
  - 'touch /var/snap/microk8s/current/credentials/cluster-tokens.txt'
  - 'echo "{{ .Token }}" >> /var/snap/microk8s/current/credentials/cluster-tokens.txt'
`

	WorkerInitTemplate = `
#cloud-config
package_update: true
snap:
  commands:
    00: snap install microk8s --channel={{ .Channel }} --classic
runcmd:
  - 'until /snap/bin/microk8s.status --wait-ready; do sleep 1; echo "waiting for status.."; done'
  - '/snap/bin/microk8s.join {{ .ControlPlaneHost }}:25000/{{ .Token }}'
`
)

// UserData is a data structure used to templatized the cloud init user data
type UserData struct {
	Channel          string
	Token            string
	ControlPlaneHost string
}
