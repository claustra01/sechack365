package activitypub

import "fmt"

func BuildActorUrl(host string, name string) string {
	return fmt.Sprintf("https://%s/api/v1/actor/%s", host, name)
}
