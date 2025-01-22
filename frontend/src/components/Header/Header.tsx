import { colors } from "@/styles/colors";
import { DesktopOnly, MobileOnly } from "@/styles/devices";
import { ActionIcon, Box, Flex, Title } from "@mantine/core";
import Link from "next/link";
import { NewPostModal } from "./NewPostModal";

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
					justify="space-between"
					w="100%"
					style={{ overflow: "hidden", borderRadius: "12px 12px 0 0" }}
				>
					<Flex p="xs" align="center">
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
					<Box p={5}>
						<NewPostModal />
					</Box>
				</Flex>
			</DesktopOnly>
			<MobileOnly>
				<Flex
					bg={colors.secondaryColor}
					justify="space-between"
					w="100%"
					style={{ overflow: "hidden" }}
				>
					<Flex align="center" p="xs">
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
					<Box p={5} pr={52}>
						<NewPostModal />
					</Box>
				</Flex>
			</MobileOnly>
		</>
	);
};
