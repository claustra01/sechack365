"use client";
import { NewPost } from "@/components/NewPost/NewPost";
import { GuestTimeline } from "@/components/Timeline/GuestTimeline";
import { UserTimeline } from "@/components/Timeline/UserTimeline";

export default function Home() {
	return (
		<main>
			<div>
				<UserTimeline />
			</div>
		</main>
	);
}
