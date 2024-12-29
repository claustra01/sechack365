import { colors } from "@/styles/colors";
import { ActionIcon, Flex, Title } from "@mantine/core";
import Link from "next/link";

export type HeaderProps = {
	title: string;
	icon: JSX.Element;
};

export const Header = (props: HeaderProps) => {
	return (
		<Flex bg={colors.secondaryColor} align="center" w="100%" p={12} style={{ overflow: "hidden" }}>
			<ActionIcon
				component={Link}
				href="/"
				variant="subtle"
				size="xl"
				c={colors.primaryColor}
			>
				{props.icon}
			</ActionIcon>
			<Title size="h3" fw={500} c={colors.primaryColor}>
				{props.title}
			</Title>
		</Flex>
	);
};
