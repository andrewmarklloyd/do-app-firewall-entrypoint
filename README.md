# do-app-firewall-entrypoint

Digital Ocean does not support configuring a Droplet's firewall to allow inbound addresses or trusted sources from a container running on their [App Platform PAAS](https://docs.digitalocean.com/products/app-platform/).

This utility can be used as an entrypoint of a Docker container running in Digital Ocean. It will add the IP address of the App Platform container to the allowed inbound list.
