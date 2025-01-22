import { Box } from "@mantine/core";
import { useContext } from "react";
import { CurrentUserContext } from "../Template/PageTemplate";
import { GuestMenu } from "./GuestMenu";
import { UserMenu } from "./UserMenu";

export const DesktopMenu = () => {
	const { user } = useContext(CurrentUserContext);

	return (
		<Box mt={60} ml={12}>
			{user ? <UserMenu /> : <GuestMenu />}
		</Box>
	);
};
