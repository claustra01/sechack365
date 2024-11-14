import type { Post } from "@/openapi/schemas";
import { Box } from "@mantine/core";
import { PostCard } from "../PostCard/PostCard";

export const TimelineBase = ({ posts }: { posts: Post[] }) => {
	if (!posts) {
		return null;
	}

	return (
		<Box>
			{posts.map((post) => {
				return <PostCard key={post.id} {...post} />;
			})}
		</Box>
	);
};
