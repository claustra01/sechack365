import { colors } from "@/styles/colors";
import { Box, Burger, Drawer } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { useContext } from "react";
import { CurrentUserContext } from "../Template/PageTemplate";
import { GuestMenu } from "./GuestMenu";
import { UserMenu } from "./UserMenu";

export const MobileMenu = () => {
	const { user } = useContext(CurrentUserContext);
	const [opened, { toggle }] = useDisclosure();

	return (
		<Box style={{ position: "absolute", top: "12px", right: "12px" }}>
			<Burger
				lineSize={2}
				size="lg"
				color={colors.white}
				opened={opened}
				onClick={toggle}
			/>
			<Drawer opened={opened} onClose={toggle}>
				{user ? <UserMenu /> : <GuestMenu />}
			</Drawer>
		</Box>
	);
};
