import { Box } from "@mantine/core";
import { IconLogin, IconUser } from "@tabler/icons-react";
import { MenuItem, type MenuItemProps } from "./MenuItem";

export const GuestMenu = () => {
	const props: MenuItemProps[] = [
		{
			icon: <IconLogin />,
			title: "Login",
			href: "/login",
		},
		{
			icon: <IconUser />,
			title: "New Account",
			href: "/register",
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
