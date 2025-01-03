import { Box } from "@mantine/core";
import { useContext } from "react";
import { CurrentUserContext } from "../Template/PageTemplate";
import { GuestMenu } from "./GuestMenu";
import { UserMenu } from "./UserMenu";

export const DesktopMenu = () => {
	const { user } = useContext(CurrentUserContext);

	return (
		<Box style={{ marginTop: "92px" }}>
			{user ? <UserMenu /> : <GuestMenu />}
		</Box>
	);
};
