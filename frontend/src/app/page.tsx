"use client";
import { UserProfile } from "@/components/Profile/UserProfile";
import { getApiV1UsersId } from "@/openapi";
import type { User } from "@/openapi/schemas";
import React from "react";

export default function Home() {
	const [data, setData] = React.useState<User | null>(null);

	React.useEffect(() => {
		getApiV1UsersId("019324b7-ab40-7c7c-ada1-702fe243847f").then((response) => {
			setData(response.data as unknown as User);
		});
	}, []);

	if (!data) {
		return <main>Loading...</main>;
	}

	return (
		<main>
			<div>
				<UserProfile {...data} />
			</div>
		</main>
	);
}
