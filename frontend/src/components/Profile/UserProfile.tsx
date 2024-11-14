import type { User } from "@/openapi/schemas";
import { ActionIcon, Avatar, Box, Image, Text, Title } from "@mantine/core";
import { IconArrowBackUp } from "@tabler/icons-react";
import Link from "next/link";
import { UserProfileCard } from "./UserProfileCard";
import { UserProfileCounter } from "./UserProfileCounter";
import { Header } from "../Header/Header";

export const UserProfile = (props: User) => {
	return (
		<Box>
			<Header title={props.display_name} icon={<IconArrowBackUp />} />
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
