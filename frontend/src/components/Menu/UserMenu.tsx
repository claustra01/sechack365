import { Box } from "@mantine/core";
import { IconHome, IconLogout, IconUser } from "@tabler/icons-react";
import { MenuItem, type MenuItemProps } from "./MenuItem";

export const UserMenu = () => {
	const props: MenuItemProps[] = [
		{
			icon: <IconHome />,
			title: "Home",
			href: "/",
		},
		{
			icon: <IconUser />,
			title: "My Profile",
			href: "/profile",
		},
		{
			icon: <IconLogout />,
			title: "Logout",
			href: "/logout",
		},
	];
	return (
		<Box>
			{props.map((item, _) => (
				<MenuItem key={item.title} {...item} />
			))}
		</Box>
	);
};
