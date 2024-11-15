import type { Post } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Box, Text } from "@mantine/core";
import { PostUserCard } from "./PostUserCard";

export const PostCard = (props: Post) => {
	console.log(props);
	return (
		<Box
			style={{
				display: "flex",
				flexDirection: "column",
				padding: "24px",
				gap: "24px",
				borderTop: `2px solid ${colors.primaryColor}`,
			}}
		>
			<PostUserCard {...props.user} />
			<Text size="lg">{props.content}</Text>
		</Box>
	);
};
