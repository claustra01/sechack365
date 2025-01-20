import type { Post } from "@/openapi/schemas";
import { Box } from "@mantine/core";
import { PostCard } from "../PostCard/PostCard";

export const TimelineBase = ({ posts }: { posts: Post[] }) => {
	if (!posts) {
		return null;
	}

	return (
		<Box style={{ overflowY: "auto", height: "calc( 100vh - 325px )" }}>
			{posts.map((post) => {
				return <PostCard key={post.id} {...post} />;
			})}
		</Box>
	);
};
