import type { User } from "@/openapi/schemas";
import { Avatar, Box, Text, Title } from "@mantine/core";

const bindUsername = (props: User) => {
	switch (props.protocol) {
		case "local":
			// TODO: nostr support
			return `@${props.username}@${props.identifiers.activitypub?.host}\n${props.identifiers.nostr?.public_key}`;
		case "activitypub":
			return `@${props.identifiers.activitypub?.local_username}@${props.identifiers.activitypub?.host}`;
		case "nostr":
			return props.identifiers.nostr?.public_key;
		default:
			return "";
	}
};

export const UserProfileCard = (props: User) => {
	return (
		<Box style={{ display: "flex", alignItems: "center", gap: "24px" }}>
			<Avatar src={props.icon} size={84} />
			<Box style={{ display: "flex", flexDirection: "column" }}>
				<Title size="h2" fw={500}>
					{props.display_name}
				</Title>
				<Box>
					<Text size="md">{bindUsername(props)}</Text>
				</Box>
			</Box>
		</Box>
	);
};
