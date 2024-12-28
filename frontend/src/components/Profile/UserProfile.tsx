import type { User } from "@/openapi/schemas";
import { Flex, Text } from "@mantine/core";
import { IconArrowBackUp } from "@tabler/icons-react";
import { Header } from "../Header/Header";
import { UserProfileCard } from "./UserProfileCard";
import { UserProfileCounter } from "./UserProfileCounter";

export const UserProfile = (props: User) => {
	return (
		<>
			<Header title={props.display_name} icon={<IconArrowBackUp />} />
			<Flex direction="column" gap={24} p={24}>
				<UserProfileCard {...props} />
				<Text size="lg">{props.profile}</Text>
				<Flex direction="row" align="center" gap={24}>
					<UserProfileCounter value={props.post_count} label="Post" />
					<UserProfileCounter value={props.follow_count} label="Follow" />
					<UserProfileCounter value={props.follower_count} label="Follower" />
				</Flex>
			</Flex>
		</>
	);
};
