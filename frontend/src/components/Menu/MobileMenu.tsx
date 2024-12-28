import { Box, Burger, Drawer } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { GuestMenu } from "./GuestMenu";
import { UserMenu } from "./UserMenu";

export const MobileMenu = (props: { isAuthenticated: boolean }) => {
	const [opened, { toggle }] = useDisclosure();
	return (
		<Box style={{ position: "absolute", top: "12px", right: "12px" }}>
			<Burger lineSize={2} size="lg" opened={opened} onClick={toggle} />
			<Drawer opened={opened} onClose={toggle}>
				{props.isAuthenticated ? <UserMenu /> : <GuestMenu />}
			</Drawer>
		</Box>
	);
};
