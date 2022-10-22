# do-app-firewall-entrypoint

Digital Ocean does not support configuring a Droplet's firewall to allow inbound addresses or trusted sources from a container running on their [App Platform PAAS](https://docs.digitalocean.com/products/app-platform/).

This utility can be used in the entrypoint of a Docker container running on the Digital Ocean App Platform. It will add the IP address of the App Platform container to the allowed inbound list.

## Example

Use an entrypoint script similar to below:

```shell
#!/bin/sh

/app/do-app-firewall-entrypoint

unset DO_TOKEN
unset DO_FIREWALL_ID

<path-to-app-binary>
```

Then include the `do-app-firewall-entrypoint` tool:
```
FROM golang:1.19-alpine as builder
RUN go install github.com/andrewmarklloyd/do-app-firewall-entrypoint@latest
...
COPY entrypoint.sh /entrypoint.sh
...
ENTRYPOINT ["/entrypoint.sh"]
```
