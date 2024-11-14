import type { User } from "@/openapi/schemas";
import { ActionIcon, Avatar, Box, Image, Text, Title } from "@mantine/core";
import { IconArrowBackUp } from "@tabler/icons-react";
import Link from "next/link";
import { UserProfileCard } from "./UserProfileCard";
import { UserProfileCounter } from "./UserProfileCounter";

export const UserProfile = (props: User) => {
	return (
		<Box>
			<Box
				bg="blue"
				style={{ display: "flex", alignItems: "center", padding: "24px" }}
			>
				<ActionIcon
					component={Link}
					href="/"
					variant="subtle"
					size="xl"
					c="white"
				>
					<IconArrowBackUp />
				</ActionIcon>
				<Title size="h3" fw={500} c="white">
					{props.display_name}
				</Title>
			</Box>
			{/* TODO: implement profile header */}
			<Image src="https://placehold.jp/540x180.png" />
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
