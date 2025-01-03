import { getApiV1FollowsFollowingId } from "@/openapi/api";
import type { User } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Avatar, Box, Button, Flex, Text, Title } from "@mantine/core";
import { useContext, useEffect, useState } from "react";
import { bindUsername } from "../../../utils/strings";
import { CurrentUserContext } from "../Template/PageTemplate";
import { FollowButton } from "./FollowButton";
import { UnfollowButton } from "./UnfollowButton";

export const UserProfileCard = (props: User) => {
	const [isFollowed, setIsFollowed] = useState<boolean>(false);
	const { user: currentUser } = useContext(CurrentUserContext);

	useEffect(() => {
		getApiV1FollowsFollowingId(props.id).then((response) => {
			setIsFollowed(response.data.found);
		});
	}, [props.id]);

	return (
		<Flex direction="column" gap={12}>
			<Flex direction="row" align="center" justify="space-between">
				<Avatar src={props.icon} size={80} />
				{currentUser?.id === props.id ? (
					// TODO: edit profile
					<Button color={colors.secondaryColor} size="lg">
						Edit
					</Button>
				) : currentUser && isFollowed ? (
					<UnfollowButton targetId={props.id} />
				) : (
					currentUser && <FollowButton targetId={props.id} />
				)}
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
