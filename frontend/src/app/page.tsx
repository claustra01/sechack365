"use client";
import { NewPost } from "@/components/NewPost/NewPost";
import { GuestTimeline } from "@/components/Timeline/GuestTimeline";

export default function Home() {
	return (
		<main>
			<div>
				<NewPost />
			</div>
		</main>
	);
}
