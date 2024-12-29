import { Box } from "@mantine/core";
import { GuestMenu } from "./GuestMenu";
import { UserMenu } from "./UserMenu";

export const DesktopMenu = (props: { isAuthenticated: boolean }) => {
	return (
		<Box style={{ marginTop: "92px" }}>
			{props.isAuthenticated ? <UserMenu /> : <GuestMenu />}
		</Box>
	);
};
