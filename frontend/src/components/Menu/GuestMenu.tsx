import { Box } from "@mantine/core";
import { IconLogin, IconUser } from "@tabler/icons-react";
import { MenuItemWithModal } from "./MenuItemWithModal";
import { ModalLogin } from "./ModalLogin";
import { ModalRegister } from "./ModalRegister";

export const GuestMenu = () => {
	return (
		<Box>
			<MenuItemWithModal
				icon={<IconLogin />}
				title="Login"
				modalContent={<ModalLogin />}
			/>
			<MenuItemWithModal
				icon={<IconUser />}
				title="Register"
				modalContent={<ModalRegister />}
			/>
		</Box>
	);
};
