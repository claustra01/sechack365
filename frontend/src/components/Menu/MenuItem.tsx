import { colors } from "@/styles/colors";
import { ActionIcon, Box, Title } from "@mantine/core";
import Link from "next/link";

export type MenuItemProps = {
	icon: JSX.Element;
	title: string;
	href: string;
};

export const MenuItem = (props: MenuItemProps) => {
	return (
		<Box style={{ display: "flex", alignItems: "center" }}>
			<ActionIcon
				component={Link}
				href={props.href}
				variant="subtle"
				size="xl"
				color={colors.secondaryColor}
			>
				{props.icon}
			</ActionIcon>
			<Title size="h3" fw={500} c={colors.secondaryColor}>
				{props.title}
			</Title>
		</Box>
	);
};
