import { colors } from "@/styles/colors";
import { Box, Button, Flex, Title } from "@mantine/core";
import { useRouter } from "next/navigation";
import React from "react";

export type MenuItemProps = {
	icon: JSX.Element;
	title: string;
	href: string;
};

export const MenuItem = (props: MenuItemProps) => {
	const router = useRouter();

	return (
		<Box p={4}>
			<Button
				variant="subtle"
				px={12}
				onClick={() => router.push(props.href)}
				style={{ borderRadius: "16px" }}
			>
				<Flex direction="row" align="center">
					{React.cloneElement(props.icon, { color: colors.secondaryColor })}
					<Title size="h3" fw={500} c={colors.secondaryColor} ml="md">
						{props.title}
					</Title>
				</Flex>
			</Button>
		</Box>
	);
};
