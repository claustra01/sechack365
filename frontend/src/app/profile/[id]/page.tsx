"use client";
import { PageTemplate } from "@/components/Template/PageTemplate";
import { UserTimeline } from "@/components/Timeline/UserTimeline";

export default function UserProfilePage({
	params,
}: { params: { id: string } }) {
	return (
		<PageTemplate>
			<UserTimeline username={params.id} />
		</PageTemplate>
	);
}
