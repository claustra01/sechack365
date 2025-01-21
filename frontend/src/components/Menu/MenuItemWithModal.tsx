import { colors } from "@/styles/colors";
import { Box, Button, Flex, Modal, Title } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import React from "react";

export type MenuItemWithModalProps = {
	icon: JSX.Element;
	title: string;
	modalContent: JSX.Element;
};

export const MenuItemWithModal = (props: MenuItemWithModalProps) => {
	const [opened, { open, close }] = useDisclosure(false);

	return (
		<Box p={4}>
			<Modal opened={opened} onClose={close} title={props.title}>
				{props.modalContent}
			</Modal>

			<Button
				variant="subtle"
				px={12}
				onClick={open}
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
