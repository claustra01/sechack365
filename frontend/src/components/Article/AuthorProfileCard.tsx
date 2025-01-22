import type { SimpleUser } from "@/openapi/schemas";
import { DesktopOnly, MobileOnly } from "@/styles/devices";
import { Avatar, Box, Flex, Text, Title } from "@mantine/core";
import Link from "next/link";

export const AuthorProfileCard = (props: SimpleUser) => {
	return (
		<Flex direction="row" align="center" gap="sm">
			<Link href={`/profile/${props.username}`}>
				<Avatar src={props.icon} size="md" />
			</Link>
			<Flex direction="column">
				<Title size="h5" fw={500}>
					{props.display_name}
				</Title>
				<Box>
					<DesktopOnly>
						<Text size="xs">{props.username}</Text>
					</DesktopOnly>
					<MobileOnly>
						<Text size="xs">
							{props.username.length > 30
								? `${props.username.slice(0, 30)}...`
								: props.username}
						</Text>
					</MobileOnly>
				</Box>
			</Flex>
		</Flex>
	);
};
