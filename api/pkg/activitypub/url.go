package activitypub

import "fmt"

func BuildActorUrl(host string, name string) string {
	return fmt.Sprintf("https://%s/api/v1/actor/%s", host, name)
}

func BuildKeyIdUrl(host string, name string) string {
	return BuildActorUrl(host, name) + "#main-key"
}
