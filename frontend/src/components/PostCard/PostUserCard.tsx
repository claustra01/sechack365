import type { SimpleUser } from "@/openapi/schemas";
import { Avatar, Box, Flex, Text } from "@mantine/core";
import Link from "next/link";

export const PostUserCard = (props: SimpleUser) => {
	return (
		<Flex direction="row" align="center" gap="md">
			<Link href={`/profile/${props.username}`}>
				<Avatar src={props.icon} size="md" />
			</Link>
			<Flex direction="column">
				<Text size="md" fw={500}>
					{props.display_name}
				</Text>
				<Box>
					<Text size="xs">
						{props.username.length > 30
							? `${props.username.slice(0, 30)}...`
							: props.username}
					</Text>
				</Box>
			</Flex>
		</Flex>
	);
};
