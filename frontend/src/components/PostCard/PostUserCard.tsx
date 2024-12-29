import type { SimpleUser } from "@/openapi/schemas";
import { Avatar, Box, Flex, Text, Title } from "@mantine/core";
import Link from "next/link";

export const PostUserCard = (props: SimpleUser) => {
	return (
		<Flex direction="row" align="center" gap={24}>
			<Link href={`/profile/${props.username}`}>
				<Avatar src={props.icon} size="lg" />
			</Link>
			<Flex direction="column">
				<Title size="h4" fw={500}>
					{props.display_name}
				</Title>
				<Box>
					<Text size="sm">{props.username}</Text>
				</Box>
			</Flex>
		</Flex>
	);
};
