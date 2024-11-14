"use client";

import { PageTemplate } from "@/components/Template/PageTemplate";
import { HomeTimeline } from "@/components/Timeline/HomeTimeline";

export default function Home() {
	return (
		<main>
			<div>
				<PageTemplate>
					<HomeTimeline />
				</PageTemplate>
			</div>
		</main>
	);
}
