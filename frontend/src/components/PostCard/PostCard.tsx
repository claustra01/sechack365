import type { Post } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Flex, Text } from "@mantine/core";
import { PostUserCard } from "./PostUserCard";

export const PostCard = (props: Post) => {
	console.log(props);
	return (
		<Flex
			direction="column"
			gap={12}
			p={24}
			mx={12}
			mt={12}
			style={{ border: `2px solid ${colors.primaryColor}`, borderRadius: 8 }}
		>
			<PostUserCard {...props.user} />
			<Text size="lg">{props.content}</Text>
		</Flex>
	);
};
