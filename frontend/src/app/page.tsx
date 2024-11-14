"use client";

import { UserTimeline } from "@/components/Timeline/UserTimeline";

export default function Home() {
	return (
		<main>
			<div>
				<UserTimeline username={"mock"} />
			</div>
		</main>
	);
}
