import { Box } from "@mantine/core";
import {
	IconHome,
	IconLogout,
	IconNote,
	IconSearch,
	IconUser,
} from "@tabler/icons-react";
import { useContext } from "react";
import { CurrentUserContext } from "../Template/PageTemplate";
import { MenuItem } from "./MenuItem";
import { MenuItemWithModal } from "./MenuItemWithModal";
import { ModalLogout } from "./ModalLogout";
import { ModalSearchUser } from "./ModalSearchUser";

export const UserMenu = () => {
	const { user } = useContext(CurrentUserContext);

	return (
		<Box>
			<MenuItem icon={<IconHome />} title="Home" href="/" />
			<MenuItemWithModal
				icon={<IconSearch />}
				title="Search User"
				modalContent={<ModalSearchUser />}
			/>
			<MenuItem
				icon={<IconUser />}
				title="My Profile"
				href={`/profile/@${user?.username}`}
			/>
			<MenuItem icon={<IconNote />} title="New Article" href={"/article/new"} />
			<MenuItemWithModal
				icon={<IconLogout />}
				title="Logout"
				modalContent={<ModalLogout />}
			/>
		</Box>
	);
};
