"use client";

import { Header } from "@/components/Header/Header";
import { PageTemplate } from "@/components/Template/PageTemplate";
import { postApiV1AuthLogin } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Box, Button, TextInput } from "@mantine/core";
import { IconArrowBackUp } from "@tabler/icons-react";
import { useState } from "react";

export default function LoginPage() {
	const [username, setUsername] = useState<string>("");
	const [password, setPassword] = useState<string>("");

	const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setUsername(e.currentTarget.value);
	};

	const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setPassword(e.currentTarget.value);
	};

	const handleSubmit = () => {
		if (username.length === 0 || password.length === 0) {
			return;
		}
		const reqBody = {
			username: username,
			password: password,
		};
		postApiV1AuthLogin(reqBody)
			.then((response) => {
				if (response.status === 204) {
					// FIXME: routerを使うようにする
					window.location.href = "/";
				}
			})
			.catch((error) => {
				console.error(error);
				alert("ログインに失敗しました");
			});
	};

	return (
		<PageTemplate>
			<Header title="Login" icon={<IconArrowBackUp />} />
			<Box
				style={{
					display: "flex",
					flexDirection: "column",
					padding: "24px",
					gap: "24px",
				}}
			>
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
					color={colors.secondaryColor}
					style={{ alignSelf: "flex-end" }}
					onClick={handleSubmit}
				>
					Login
				</Button>
			</Box>
		</PageTemplate>
	);
}
