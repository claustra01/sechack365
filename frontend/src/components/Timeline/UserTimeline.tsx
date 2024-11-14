import { getApiV1LookupUsername, getApiV1UsersIdPosts } from "@/openapi";
import type { Post, User } from "@/openapi/schemas";
import { Box } from "@mantine/core";
import { IconHome } from "@tabler/icons-react";
import { useEffect, useState } from "react";
import { Header } from "../Header/Header";
import { PostCard } from "../PostCard/PostCard";
import { UserProfile } from "../Profile/UserProfile";

export type UserTimelineProps = {
	username: string;
};

export const UserTimeline = (props: UserTimelineProps) => {
	const [user, setUser] = useState<User | null>(null);
	const [posts, setPosts] = useState<Post[]>([]);

	useEffect(() => {
		getApiV1LookupUsername(props.username).then((response) => {
			setUser(response.data as unknown as User);
		});
	}, [props.username]);

	useEffect(() => {
		if (user?.id) {
			getApiV1UsersIdPosts(user.id).then((response) => {
				setPosts(response.data as unknown as Post[]);
			});
		}
	}, [user?.id]);

	if (!user) {
		return <Box>Loading...</Box>;
	}

	return (
		<Box w={720}>
			<UserProfile {...user} />
			{posts.map((post) => {
				return <PostCard key={post.id} {...post} />;
			})}
		</Box>
	);
};
