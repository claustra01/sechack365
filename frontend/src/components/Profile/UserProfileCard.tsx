import { getApiV1FollowsFollowingId, getApiV1UsersMe } from "@/openapi/api";
import type { User } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Avatar, Box, Button, Flex, Text, Title } from "@mantine/core";
import { useEffect, useState } from "react";
import { bindUsername } from "../../../utils/strings";
import { FollowButton } from "./FollowButton";

const followButton = (props: User) => {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
	const [isFollowed, setIsFollowed] = useState<boolean>(false);
	const [currentUser, setCurrentUser] = useState<User | null>(null);

	useEffect(() => {
		getApiV1UsersMe().then((response) => {
			setIsAuthenticated(true);
			setCurrentUser(response.data as unknown as User);
		});
		getApiV1FollowsFollowingId(props.id).then((response) => {
			setIsFollowed(response.data.found);
		});
	}, [props.id]);

	if (!isAuthenticated) {
		return null;
	}

	// TODO: edit profile
	if (currentUser?.id === props.id) {
		return (
			<Button color={colors.secondaryColor} size="lg">
				Edit
			</Button>
		);
	}

	// TODO: unfollow
	if (isFollowed) {
		return (
			<Button color={colors.secondaryColor} size="lg">
				Unfollow
			</Button>
		);
	}

	return <FollowButton targetId={props.id} />;
};

export const UserProfileCard = (props: User) => {
	return (
		<Flex direction="column" gap={12}>
			<Flex direction="row" align="center" justify="space-between">
				<Avatar src={props.icon} size={80} />
				{followButton(props)}
			</Flex>
			<Flex direction="column" gap={4}>
				<Title size="h3" fw={500}>
					{props.display_name}
				</Title>
				<Box style={{ maxWidth: "calc( 100vw - 48px )", overflowX: "auto" }}>
					<Text size="sm" c={colors.black}>
						{bindUsername(props)}
					</Text>
				</Box>
			</Flex>
		</Flex>
	);
};
