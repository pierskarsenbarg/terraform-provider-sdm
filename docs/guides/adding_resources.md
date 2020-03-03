# Adding Resources to strongDM

Each resource created adds the credentials for the designated service; these are then used by your strongDM gateways for authentication and leased to users when granted permission.

## Leased Credentials

[Credential Leasing](https://www.strongdm.com/docs/architecture/leasing/)

## Port-overrides

The port-override assignment is the port end-users will use when connecting to the configured service; whereas the port assignment is used for the gateway to establish a connection to the service. All port-override assignments need to be unique across all datasource types.

[Ad Hoc Ports and Port Overrides](https://www.strongdm.com/docs/admin-guide/port-overrides/)

## Examples

### SSH Servers

    resource "sdm_resource" "rack_server" {
      ssh {
        name          = "admin_server"
        hostname      = "10.99.238.2"
        port          = 22
        port_override = 25307
    # where is SSH username?
      }

    output "rack_server_public_key" {
    # just a guess
    	value = sdm_resource.rack_server.public_key
    }

SSH connections are secured using a RSA key pair. The RSA key pair is generated by strongDM when the resource is generated. To complete the authentication, the public key needs to be added to the authorized key file for the SSH username specified.  `~/target_user/.ssh/authorized_keys`

### RDP Server

    resource "sdm_resource" "win_server_2019" {
      rdp {
        name          = "windows server 2019"
        hostname      = "admin.windows.corp.net"
    		username      = "administrator"
    		password      = var.win_admin_pass
        port          = 3389
        port_override = 13390
      }

### Kubernetes

    resource "sdm_resource" "prod_k8s" {
      kubernetes {
        name                           = "k8s1"
        hostname                       = aws_eks_cluster.eks_cluster.endpoint
        port                           = 443
        certificate_authority          = aws_eks_cluster.main_cluster.certificate_authority.0.data
    		certificate_authority_filename =
    # Why are there file names here?
    		client_certificate             =
    		client_certificate_filename    =
    		client_key                     =
    		client_key_filename            =
      }

### Database

    resource "sdm_resource" "mysql" {
      mysql {
        name          = "mysql_prod_db"
        hostname      = "prod.mysql.corp.net"
    		username      = var.mysql-admin
    		password      = var.mysql-pass
        database      = "master"
        port          = 3306
        port_override = 13306
      }