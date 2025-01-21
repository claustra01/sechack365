import { Box } from "@mantine/core";
import { IconLogin, IconQuestionMark, IconUser } from "@tabler/icons-react";
import { MenuItemWithModal } from "./MenuItemWithModal";
import { ModalLogin } from "./ModalLogin";
import { ModalRegister } from "./ModalRegister";
import { MenuItem } from "./MenuItem";

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
			<MenuItem
				icon={<IconQuestionMark />}
				title="How to Demo"
				href="/how_to_demo"
			/>
		</Box>
	);
};
