import { getApiV1UsersMe } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { DesktopOnly, MobileOnly } from "@/styles/devices";
import { Box, Flex } from "@mantine/core";
import { useEffect, useState } from "react";
import { DesktopMenu } from "../Menu/DesktopMenu";
import { MobileMenu } from "../Menu/MobileMenu";

export const PageTemplate = ({ children }: { children: React.ReactNode }) => {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);

	useEffect(() => {
		getApiV1UsersMe().then(() => {
			setIsAuthenticated(true);
		});
	}, []);

	return (
		<main>
			<DesktopOnly>
				<Flex
					bg={colors.primaryColor}
					direction={{ base: "row" }}
					gap={{ base: "24px" }}
					justify={{ base: "center" }}
					pt={24}
				>
					<Box
						bg={colors.white}
						w={720}
						style={{ minHeight: "calc( 100vh - 24px )" }}
					>
						{children}
					</Box>
					<DesktopMenu isAuthenticated={isAuthenticated} />
				</Flex>
			</DesktopOnly>
			<MobileOnly>
				<Box bg={colors.white} w={"100%"} style={{ minHeight: "100vh" }}>
					{children}
				</Box>
				<MobileMenu isAuthenticated={isAuthenticated} />
			</MobileOnly>
		</main>
	);
};
