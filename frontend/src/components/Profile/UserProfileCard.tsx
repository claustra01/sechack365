import type { User } from "@/openapi/schemas";
import { Avatar, Box, Text, Title } from "@mantine/core";

const bindUsername = (props: User) => {
	switch (props.protocol) {
		case "local":
			// TODO: nostr support
			return `${props.username}@${props.host}`;
		case "activitypub":
			return `${props.username}@${props.host}`;
		case "nostr":
			// TODO: nostr support
			return "";
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
					<Text size="lg">{bindUsername(props)}</Text>
				</Box>
			</Box>
		</Box>
	);
};
