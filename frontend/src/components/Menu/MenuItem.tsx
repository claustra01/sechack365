import { colors } from "@/styles/colors";
import { ActionIcon, Anchor, Flex, Title } from "@mantine/core";

export type MenuItemProps = {
	icon: JSX.Element;
	title: string;
	href: string;
};

export const MenuItem = (props: MenuItemProps) => {
	return (
		<Anchor href={props.href} style={{ textDecoration: "none" }}>
			<Flex direction="row" align="center">
				<ActionIcon variant="subtle" size="xl" color={colors.secondaryColor}>
					{props.icon}
				</ActionIcon>
				<Title size="h3" fw={500} c={colors.secondaryColor}>
					{props.title}
				</Title>
			</Flex>
		</Anchor>
	);
};
