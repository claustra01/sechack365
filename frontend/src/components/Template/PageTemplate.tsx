import { getApiV1UsersMe } from "@/openapi/api";
import type { User } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { DesktopOnly, MobileOnly } from "@/styles/devices";
import { Box, Flex } from "@mantine/core";
import { createContext, useEffect, useState } from "react";
import { DesktopMenu } from "../Menu/DesktopMenu";
import { MobileMenu } from "../Menu/MobileMenu";

type TCurrentUserContext = {
	user: User | null;
	setUser: React.Dispatch<React.SetStateAction<User | null>>;
};

export const CurrentUserContext = createContext({} as TCurrentUserContext);

export const PageTemplate = ({ children }: { children: React.ReactNode }) => {
	const [currentUser, setCurrentUser] = useState<User | null>(null);

	useEffect(() => {
		getApiV1UsersMe().then((r) => {
			setCurrentUser(r.data as unknown as User);
		});
	}, []);

	return (
		<main>
			<CurrentUserContext.Provider
				value={{ user: currentUser, setUser: setCurrentUser }}
			>
				<DesktopOnly>
					<Flex
						bg={colors.primaryColor}
						direction="row"
						py="sm"
						justify="center"
					>
						<Box
							bg={colors.white}
							w={540}
							style={{
								minHeight: "calc( 100vh - 24px )",
								borderRadius: "12px",
								boxShadow: "5px 5px 10px rgba(0, 0, 0, 0.1)",
							}}
						>
							{children}
						</Box>
						<DesktopMenu />
					</Flex>
				</DesktopOnly>
				<MobileOnly>
					<Box bg={colors.white} w={"100%"} style={{ minHeight: "100vh" }}>
						{children}
					</Box>
					<MobileMenu />
				</MobileOnly>
			</CurrentUserContext.Provider>
		</main>
	);
};
