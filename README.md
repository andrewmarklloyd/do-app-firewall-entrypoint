# do-app-firewall-entrypoint

Digital Ocean does not support configuring a Droplet's firewall to allow inbound addresses or trusted sources from a container running on their [App Platform PAAS](https://docs.digitalocean.com/products/app-platform/).

This utility can be used in the entrypoint of a Docker container running on the Digital Ocean App Platform. It will add the IP address of the App Platform container to the allowed inbound list.

## Required Environment Variables

```
STATIC_INBOUND_IP: ip address to keep in the inbound rules

FIREWALL_NAME: name of the firewall in Digital Ocean to update

FIREWALL_PORT: port of the firewall to match against

DO_ACCESS_TOKEN: access token to update the Digital Ocean firewall
```

## Example

Use an entrypoint script similar to below:

```shell
#!/bin/sh

/app/do-app-firewall-entrypoint

# best to unset these variables before running the app
unset STATIC_INBOUND_IP
unset FIREWALL_NAME
unset FIREWALL_PORT
unset DO_ACCESS_TOKEN

<path-to-app-binary>
```

Then include the `do-app-firewall-entrypoint` tool:
```
FROM golang:1.22-alpine as builder
RUN go install github.com/andrewmarklloyd/do-app-firewall-entrypoint@latest
...
COPY entrypoint.sh /entrypoint.sh
...
ENTRYPOINT ["/entrypoint.sh"]
```
