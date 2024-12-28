import type { User } from "@/openapi/schemas";
import { Box, Text } from "@mantine/core";
import { IconArrowBackUp } from "@tabler/icons-react";
import { Header } from "../Header/Header";
import { UserProfileCard } from "./UserProfileCard";
import { UserProfileCounter } from "./UserProfileCounter";

export const UserProfile = (props: User) => {
	return (
		<Box>
			<Header title={props.display_name} icon={<IconArrowBackUp />} />
			<Box
				style={{
					display: "flex",
					flexDirection: "column",
					padding: "24px",
					gap: "24px",
				}}
			>
				<UserProfileCard {...props} />
				<Text size="lg">{props.profile}</Text>
				<Box style={{ display: "flex", alignItems: "center", gap: "24px" }}>
					{/* TODO: add user profile count */}
					<UserProfileCounter value={0} label="Follow" />
					<UserProfileCounter value={0} label="Follower" />
				</Box>
			</Box>
		</Box>
	);
};
