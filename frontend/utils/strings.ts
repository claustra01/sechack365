import { User } from "@/openapi/schemas";

export const bindUsername = (props: User) => {
	switch (props.protocol) {
		case "local":
			return `@${props.username}@${props.identifiers.activitypub?.host}\n${props.identifiers.nostr?.npub}`;
		case "activitypub":
			return `@${props.identifiers.activitypub?.local_username}@${props.identifiers.activitypub?.host}`;
		case "nostr":
			return props.identifiers.nostr?.npub;
		default:
			return "";
	}
};
