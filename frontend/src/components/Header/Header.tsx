import { colors } from "@/styles/colors";
import { ActionIcon, Box, Title } from "@mantine/core";
import Link from "next/link";

export type HeaderProps = {
	title: string;
	icon: JSX.Element;
};

export const Header = (props: HeaderProps) => {
	return (
		<Box
			bg={colors.secondaryColor}
			style={{ display: "flex", alignItems: "center", padding: "24px" }}
		>
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
		</Box>
	);
};
