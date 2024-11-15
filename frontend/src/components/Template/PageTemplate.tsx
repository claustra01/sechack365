import { getApiV1UsersMe } from "@/openapi";
import { Box } from "@mantine/core";
import { useEffect, useState } from "react";
import { GuestMenu } from "../Menu/GuestMenu";
import { UserMenu } from "../Menu/UserMenu";
import { colors } from "@/styles/colors";

export const PageTemplate = ({ children }: { children: React.ReactNode }) => {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);

	useEffect(() => {
		getApiV1UsersMe().then((response) => {
			setIsAuthenticated(true);
		});
	}, []);

	return (
		<main>
			<Box
				bg={colors.primaryColor}
				style={{
					display: "flex",
					flexDirection: "column",
					alignItems: "center",
					minHeight: "100vh",
				}}
			>
				<Box
					style={{
						display: "flex",
						gap: "24px",
						marginTop: "24px",
					}}
				>
					<Box bg={colors.white} w={720} style={{ minHeight: "100vh" }}>
						{children}
					</Box>
					<Box style={{ marginTop: "92px" }}>
						{isAuthenticated ? <UserMenu /> : <GuestMenu />}
					</Box>
				</Box>
			</Box>
		</main>
	);
};
