"use client";
import { UserTimeline } from "@/components/Timeline/UserTimeline";

export default function UserProfilePage({
	params,
}: { params: { id: string } }) {
	return (
		<main>
			<div>
				<UserTimeline username={params.id} />
			</div>
		</main>
	);
}
