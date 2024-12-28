import type { User } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Avatar, Box, Flex, Text, Title } from "@mantine/core";
import { bindUsername } from "../../../utils/strings";

export const UserProfileCard = (props: User) => {
	return (
		<Flex align="center" gap={20}>
			<Avatar src={props.icon} size={80} />
			<Flex direction="column" gap={4}>
				<Title size="h2" fw={500}>
					{props.display_name}
				</Title>
				<Box style={{ maxWidth: "calc( 100vw - 144px )", overflowX: "auto" }}>
					<Text size="sm" c={colors.black}>
						{bindUsername(props)}
					</Text>
				</Box>
			</Flex>
		</Flex>
	);
};
