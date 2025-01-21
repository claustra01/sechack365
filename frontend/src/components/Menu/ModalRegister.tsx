import { postApiV1AuthLogin, postApiV1AuthRegister } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Box, Button, TextInput } from "@mantine/core";
import { useState } from "react";

export const ModalRegister = () => {
	const [username, setUsername] = useState<string>("");
	const [password, setPassword] = useState<string>("");
	const [confirmPassword, setConfirmPassword] = useState<string>("");

	const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setUsername(e.currentTarget.value);
	};

	const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setPassword(e.currentTarget.value);
	};

	const handleConfirmPasswordChange = (
		e: React.ChangeEvent<HTMLInputElement>,
	) => {
		setConfirmPassword(e.currentTarget.value);
	};

	const handleSubmit = () => {
		if (password !== confirmPassword) {
			alert("確認用パスワードが一致しません");
			return;
		}
		const reqBody = {
			username: username,
			password: password,
		};
		postApiV1AuthRegister(reqBody)
			.then((response) => {
				if (response.status === 201) {
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
				}
			})
			.catch((error) => {
				console.error(error);
				alert("ユーザー登録に失敗しました");
			});
	};

	return (
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
			<TextInput
				label="ConfirmPassword"
				placeholder="ConfirmPassword"
				type="password"
				onChange={handleConfirmPasswordChange}
			/>
			<Button
				disabled={
					username.length === 0 ||
					password.length === 0 ||
					confirmPassword.length === 0
				}
				color={colors.secondaryColor}
				style={{ alignSelf: "flex-end" }}
				onClick={handleSubmit}
			>
				Register
			</Button>
		</Box>
	);
};
