import type { SimpleUser } from "@/openapi/schemas";
import { Avatar, Box, Text, Title } from "@mantine/core";
import Link from "next/link";

export const PostUserCard = (props: SimpleUser) => {
	return (
		<Box style={{ display: "flex", alignItems: "center", gap: "24px" }}>
			<Link href={`/profile/${props.username}`}>
				<Avatar src={props.icon} size="lg" />
			</Link>
			<Box style={{ display: "flex", flexDirection: "column" }}>
				<Title size="h4" fw={500}>
					{props.display_name}
				</Title>
				<Box>
					<Text size="sm">{props.username}</Text>
				</Box>
			</Box>
		</Box>
	);
};
