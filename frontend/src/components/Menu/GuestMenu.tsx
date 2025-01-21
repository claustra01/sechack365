import { Box } from "@mantine/core";
import { IconLogin, IconUser } from "@tabler/icons-react";
import { MenuItem } from "./MenuItem";
import { MenuItemWithModal } from "./MenuItemWithModal";
import { ModalLogin } from "./ModalLogin";

export const GuestMenu = () => {
	return (
		<Box>
			<MenuItemWithModal
				icon={<IconLogin />}
				title="Login"
				modalContent={<ModalLogin />}
			/>
			<MenuItem icon={<IconUser />} href="/register" title="Register" />
		</Box>
	);
};
