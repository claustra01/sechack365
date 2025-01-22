import type { Post } from "@/openapi/schemas";
import { ScrollArea } from "@mantine/core";
import { PostCard } from "../PostCard/PostCard";

export const TimelineBase = ({ posts }: { posts: Post[] }) => {
	if (!posts) {
		return null;
	}

	return (
		<ScrollArea h={"calc( 100vh - 78px )"}>
			{posts.map((post) => {
				return <PostCard key={post.id} {...post} />;
			})}
		</ScrollArea>
	);
};
