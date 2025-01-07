import type { User } from "@/openapi/schemas";
import { Anchor, Flex, Text } from "@mantine/core";
import { IconArrowBackUp } from "@tabler/icons-react";
import DOMPurify from "dompurify";
import parse from "html-react-parser";
import { Header } from "../Header/Header";
import { UserProfileCard } from "./UserProfileCard";
import { UserProfileCounter } from "./UserProfileCounter";

export const UserProfile = (props: User) => {
	return (
		<>
			<Header title={props.display_name} icon={<IconArrowBackUp />} />
			<Flex direction="column" gap={24} p={24}>
				<UserProfileCard {...props} />
				<Text size="lg" style={{ wordBreak: "break-word" }}>
					{parse(DOMPurify.sanitize(props.profile))}
				</Text>
				<Flex direction="row" align="center" gap={24}>
					<Anchor href="#posts" style={{ textDecoration: "none" }}>
						<UserProfileCounter value={props.post_count} label="Post" />
					</Anchor>
					<Anchor href="#follows" style={{ textDecoration: "none" }}>
						<UserProfileCounter value={props.follow_count} label="Follow" />
					</Anchor>
					<Anchor href="#followers" style={{ textDecoration: "none" }}>
						<UserProfileCounter value={props.follower_count} label="Follower" />
					</Anchor>
				</Flex>
			</Flex>
		</>
	);
};
