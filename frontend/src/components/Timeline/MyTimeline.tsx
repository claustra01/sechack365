import { getApiV1LookupUsername, getApiV1UsersIdPosts, getApiV1UsersMe } from "@/openapi";
import type { Post, User } from "@/openapi/schemas";
import { Box } from "@mantine/core";
import { useEffect, useState } from "react";
import { UserProfile } from "../Profile/UserProfile";
import { TimelineBase } from "./TimelineBase";

export const MyTimeline = () => {
	const [user, setUser] = useState<User | null>(null);
	const [posts, setPosts] = useState<Post[]>([]);

	useEffect(() => {
		getApiV1UsersMe().then((response) => {
			setUser(response.data as unknown as User);
		});
	}, []);

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
      <TimelineBase posts={posts} />
		</Box>
	);
};
