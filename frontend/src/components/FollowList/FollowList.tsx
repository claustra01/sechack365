import type { SimpleUser } from "@/openapi/schemas";
import { Box } from "@mantine/core";
import { SimpleUserCard } from "./SimpleUserCard";

export const FollowList = ({ users }: { users: SimpleUser[] }) => {
	if (!users) {
		return null;
	}

	return (
		<Box style={{ overflowY: "auto", height: "calc( 100vh - 301px )" }}>
			{users.map((user) => {
				return <SimpleUserCard key={user.username} {...user} />;
			})}
		</Box>
	);
};
