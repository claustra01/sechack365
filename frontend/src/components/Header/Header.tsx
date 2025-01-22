import { colors } from "@/styles/colors";
import { DesktopOnly, MobileOnly } from "@/styles/devices";
import { ActionIcon, Flex, Title } from "@mantine/core";
import Link from "next/link";

export type HeaderProps = {
	title: string;
	icon: JSX.Element;
};

export const Header = (props: HeaderProps) => {
	return (
		<>
			<DesktopOnly>
				<Flex
					bg={colors.secondaryColor}
					align="center"
					w="100%"
					p="xs"
					style={{ overflow: "hidden", borderRadius: "12px 12px 0 0" }}
				>
					<ActionIcon
						component={Link}
						href="/"
						variant="subtle"
						size="lg"
						c={colors.primaryColor}
					>
						{props.icon}
					</ActionIcon>
					<Title size="h4" fw={500} c={colors.primaryColor}>
						{props.title}
					</Title>
				</Flex>
			</DesktopOnly>
			<MobileOnly>
				<Flex
					bg={colors.secondaryColor}
					align="center"
					w="100%"
					p="xs"
					style={{ overflow: "hidden" }}
				>
					<ActionIcon
						component={Link}
						href="/"
						variant="subtle"
						size="lg"
						c={colors.primaryColor}
					>
						{props.icon}
					</ActionIcon>
					<Title size="h4" fw={500} c={colors.primaryColor}>
						{props.title}
					</Title>
				</Flex>
			</MobileOnly>
		</>
	);
};
