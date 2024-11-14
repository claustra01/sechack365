import { MenuItem } from "@/components/Menu/MenuItem";
import { IconHeart } from "@tabler/icons-react";

export default function Home() {
	return (
		<main>
			<div>
				<MenuItem icon={<IconHeart />} title="Menu Item" href="/" />
			</div>
		</main>
	);
}
