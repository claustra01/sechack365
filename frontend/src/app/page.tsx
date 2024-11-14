"use client";
import { PostCard } from "@/components/Post/PostCard";
import { getApiV1PostsId, getApiV1UsersId } from "@/openapi";
import type { Post, User } from "@/openapi/schemas";
import React from "react";

export default function Home() {
	const [data, setData] = React.useState<Post | null>(null);

	React.useEffect(() => {
		getApiV1PostsId("01932b39-e617-7851-bdc0-2b97a972b48c").then((response) => {
			setData(response.data as unknown as Post);
		});
	}, []);

	if (!data) {
		return <main>Loading...</main>;
	}

	return (
		<main>
			<div>
				<PostCard {...data} />
			</div>
		</main>
	);
}
