import { Box } from "@mantine/core";
import { IconHome, IconLogout, IconUser } from "@tabler/icons-react";
import { useContext } from "react";
import { CurrentUserContext } from "../Template/PageTemplate";
import { MenuItem } from "./MenuItem";
import { MenuItemWithModal } from "./MenuItemWithModal";
import { ModalLogout } from "./ModalLogout";

export const UserMenu = () => {
	const { user } = useContext(CurrentUserContext);

	return (
		<Box>
			<MenuItem icon={<IconHome />} title="Home" href="/" />
			<MenuItem
				icon={<IconUser />}
				title="My Profile"
				href={`/profile/@${user?.username}`}
			/>
			<MenuItemWithModal
				icon={<IconLogout />}
				title="Logout"
				modalContent={<ModalLogout />}
			/>
		</Box>
	);
};
