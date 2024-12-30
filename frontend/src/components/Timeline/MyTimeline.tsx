import {
	getApiV1UsersIdFollowers,
	getApiV1UsersIdFollows,
	getApiV1UsersIdPosts,
	getApiV1UsersMe,
} from "@/openapi/api";
import type { Post, User } from "@/openapi/schemas";
import { Box } from "@mantine/core";
import { useEffect, useState } from "react";
import { FollowList } from "../FollowList/FollowList";
import { UserProfile } from "../Profile/UserProfile";
import { TimelineBase } from "./TimelineBase";

export const MyTimeline = () => {
	const [hash, setHash] = useState<string>("");
	const [user, setUser] = useState<User | null>(null);
	const [posts, setPosts] = useState<Post[]>([]);
	const [follows, setFollows] = useState<User[]>([]);

	useEffect(() => {
		const handleHashChange = () => {
			setHash(window.location.hash);
		};
		setHash(window.location.hash);
		window.addEventListener("hashchange", handleHashChange);
		return () => {
			window.removeEventListener("hashchange", handleHashChange);
		};
	}, []);

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

	useEffect(() => {
		if (user?.id && hash === "#follows") {
			getApiV1UsersIdFollows(user.id).then((response) => {
				setFollows(response.data as unknown as User[]);
			});
		}
		if (user?.id && hash === "#followers") {
			getApiV1UsersIdFollowers(user.id).then((response) => {
				setFollows(response.data as unknown as User[]);
			});
		}
	}, [user?.id, hash]);

	if (!user) {
		return <Box>Loading...</Box>;
	}

	return (
		<Box>
			<UserProfile {...user} />
			{hash === "#follows" || hash === "#followers" ? (
				<FollowList users={follows} />
			) : (
				<TimelineBase posts={posts} />
			)}
		</Box>
	);
};
