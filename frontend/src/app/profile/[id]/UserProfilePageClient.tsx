"use client";

import { PageTemplate } from "@/components/Template/PageTemplate";
import { UserTimeline } from "@/components/Timeline/UserTimeline";

type UserProfilePageClientProps = {
	id: string;
};

export function UserProfilePageClient({ id }: UserProfilePageClientProps) {
	return (
		<PageTemplate>
			<UserTimeline username={id} />
		</PageTemplate>
	);
}
