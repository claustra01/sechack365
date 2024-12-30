import type { SimpleUser } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Avatar, Box, Flex, Text, Title } from "@mantine/core";
import Link from "next/link";

export const SimpleUserCard = (props: SimpleUser) => {
	return (
		<Flex
			direction="row"
			align="center"
			gap={24}
			p={24}
			mx={12}
			mt={12}
			style={{ border: `2px solid ${colors.primaryColor}`, borderRadius: 8 }}
		>
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
