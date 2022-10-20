package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/digitalocean/godo"
)

var staticInboundIP string

func main() {
	staticInboundIP = os.Getenv("STATIC_INBOUND_IP")
	if staticInboundIP == "" {
		log.Fatalln("STATIC_INBOUND_IP env var must be set")
	}

	firewallName := os.Getenv("FIREWALL_NAME")
	if firewallName == "" {
		log.Fatalln("FIREWALL_NAME env var must be set")
	}

	firewallPort := os.Getenv("FIREWALL_PORT")
	if firewallPort == "" {
		log.Fatalln("FIREWALL_PORT env var must be set")
	}

	ownIP, err := getIP()
	if err != nil {
		log.Fatalln(err)
	}

	client := godo.NewFromToken(os.Getenv("DO_ACCESS_TOKEN"))
	ctx := context.Background()
	f, err := getFirewall(ctx, client, firewallName)
	if err != nil {
		log.Fatalln(err)
	}

	newFirewall := updateInboundAddresses(f, firewallPort, ownIP)

	_, _, err = client.Firewalls.Update(ctx, newFirewall.ID, &godo.FirewallRequest{
		Name:          newFirewall.Name,
		InboundRules:  newFirewall.InboundRules,
		OutboundRules: newFirewall.OutboundRules,
		DropletIDs:    newFirewall.DropletIDs,
		Tags:          newFirewall.Tags,
	})
	if err != nil {
		log.Fatalln("updating firewall:", err)
	}
	fmt.Println("successfully updated firewall")
}

func getFirewall(ctx context.Context, c *godo.Client, firewallName string) (godo.Firewall, error) {
	firewalls, _, err := c.Firewalls.List(ctx, &godo.ListOptions{})
	if err != nil {
		return godo.Firewall{}, fmt.Errorf("listing firewalls: %s", err)
	}

	for _, f := range firewalls {
		if f.Name == firewallName {
			return f, nil
		}
	}

	return godo.Firewall{}, fmt.Errorf("no firewalls matched the expected name")
}

func updateInboundAddresses(f godo.Firewall, port string, ownIP string) godo.Firewall {
	newFirewall := f
	for _, r := range newFirewall.InboundRules {
		if r.PortRange == port {
			r.Sources.Addresses = []string{staticInboundIP, ownIP}
		}
	}
	return newFirewall
}

func getIP() (string, error) {
	urlStr := "http://ifconfig.me"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, nil)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("making request to get own ip address: %s", err)
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body from getting own ip address: %s", err)
	}

	return string(resBody), nil
}
