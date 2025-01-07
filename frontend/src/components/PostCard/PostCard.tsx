import type { Post } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Flex, Text } from "@mantine/core";
import DOMPurify from "dompurify";
import parse from "html-react-parser";
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
			<Text size="lg" style={{ wordBreak: "break-word" }}>
				{parse(DOMPurify.sanitize(props.content))}
			</Text>
		</Flex>
	);
};
