package activitypub

import "fmt"

func BuildActorUrl(host, id string) string {
	return fmt.Sprintf("https://%s/api/v1/users/%s", host, id)
}

func BuildKeyIdUrl(host string, name string) string {
	return BuildActorUrl(host, name) + "#main-key"
}
