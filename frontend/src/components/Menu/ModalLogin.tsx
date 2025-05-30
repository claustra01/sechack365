import { postApiV1AuthLogin } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Button, Flex, TextInput, Title } from "@mantine/core";
import { useState } from "react";

export const ModalLogin = () => {
	const [username, setUsername] = useState<string>("");
	const [password, setPassword] = useState<string>("");

	const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setUsername(e.currentTarget.value);
	};

	const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setPassword(e.currentTarget.value);
	};

	const handleSubmit = () => {
		const reqBody = {
			username: username,
			password: password,
		};
		postApiV1AuthLogin(reqBody)
			.then((response) => {
				if (response.status === 204) {
					window.location.reload();
				}
			})
			.catch((error) => {
				console.error(error);
				alert("ログインに失敗しました");
			});
	};

	return (
		<Flex direction="column" gap={12}>
			<Title size="h4" fw={500} c={colors.secondaryColor}>
				Login
			</Title>
			<TextInput
				label="Username"
				placeholder="Username"
				onChange={handleUsernameChange}
			/>
			<TextInput
				label="Password"
				placeholder="Password"
				type="password"
				onChange={handlePasswordChange}
			/>
			<Button
				disabled={username.length === 0 || password.length === 0}
				color={colors.secondaryColor}
				style={{ alignSelf: "flex-end" }}
				onClick={handleSubmit}
			>
				Login
			</Button>
		</Flex>
	);
};
