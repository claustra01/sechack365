import type { User } from "@/openapi/schemas";
import { Avatar, Box, Text, Title } from "@mantine/core";
import Link from "next/link";

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

export const PostUserCard = (props: User) => {
	return (
		<Box style={{ display: "flex", alignItems: "center", gap: "24px" }}>
			<Link href={`/profile/${bindUsername(props)}`}>
				<Avatar src={props.icon} size="lg" />
			</Link>
			<Box style={{ display: "flex", flexDirection: "column" }}>
				<Title size="h4" fw={500}>
					{props.display_name}
				</Title>
				<Box>
					<Text size="sm">{bindUsername(props)}</Text>
				</Box>
			</Box>
		</Box>
	);
};
