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

const CurrentUserContext = createContext({} as TCurrentUserContext);

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
						gap={24}
						justify="center"
						pt={24}
					>
						<Box
							bg={colors.white}
							w={720}
							style={{ minHeight: "calc( 100vh - 24px )" }}
						>
							{children}
						</Box>
						<DesktopMenu isAuthenticated={currentUser != null} />
					</Flex>
				</DesktopOnly>
				<MobileOnly>
					<Box bg={colors.white} w={"100%"} style={{ minHeight: "100vh" }}>
						{children}
					</Box>
					<MobileMenu isAuthenticated={currentUser != null} />
				</MobileOnly>
			</CurrentUserContext.Provider>
		</main>
	);
};
