import { getApiV1Timeline } from "@/openapi";
import type { Post } from "@/openapi/schemas";
import { Box } from "@mantine/core";
import { IconHome } from "@tabler/icons-react";
import { useState } from "react";
import { Header } from "../Header/Header";
import { TimelineBase } from "./TimelineBase";

export const GuestTimeline = () => {
	const [posts, setPosts] = useState<Post[]>([]);

	getApiV1Timeline().then((response) => {
		setPosts(response.data as unknown as Post[]);
	});

	return (
		<Box>
			<Header title={"Home"} icon={<IconHome />} />
			<TimelineBase posts={posts} />
		</Box>
	);
};
