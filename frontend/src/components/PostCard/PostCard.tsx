import type { Post } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Flex, Text } from "@mantine/core";
import { parseHtml } from "../../../utils/html";
import { PostUserCard } from "./PostUserCard";

export const PostCard = (props: Post) => {
	return (
		<Flex
			direction="column"
			gap={12}
			p={24}
			mx={12}
			mt={12}
			style={{
				border: `2px solid ${colors.primaryColor}`,
				borderRadius: 8,
				boxShadow: "5px 5px 5px rgba(0, 0, 0, 0.1)",
			}}
		>
			<PostUserCard {...props.user} />
			<Text size="lg" style={{ wordBreak: "break-word" }}>
				{parseHtml(props.content)}
			</Text>
		</Flex>
	);
};
