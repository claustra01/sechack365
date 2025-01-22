import { getApiV1LookupUsername } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Button, Flex, TextInput, Title } from "@mantine/core";
import { useState } from "react";

export const ModalSearchUser = () => {
	const [username, setUsername] = useState<string>("");

	const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setUsername(e.currentTarget.value);
	};

	const handleSubmit = () => {
		getApiV1LookupUsername(username)
			.then((response) => {
				if (response.status === 200) {
					window.location.href = `/profile/${username}`;
				}
			})
			.catch((error) => {
				console.error(error);
				alert("ユーザーが見つかりませんでした");
			});
	};

	return (
		<Flex direction="column" gap={12}>
			<Title size="h4" fw={500} c={colors.secondaryColor}>
				Search User
			</Title>
			<TextInput
				label="Username"
				placeholder="@user@example.com / npub1xxxxxx...."
				onChange={handleUsernameChange}
			/>
			<Button
				disabled={username.length === 0}
				color={colors.secondaryColor}
				style={{ alignSelf: "flex-end" }}
				onClick={handleSubmit}
			>
				Search
			</Button>
		</Flex>
	);
};
