import { postApiV1AuthLogout } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Box, Button, Text } from "@mantine/core";

export const ModalLogout = () => {
	const handleSubmit = () => {
		postApiV1AuthLogout()
			.then((response) => {
				if (response.status === 204) {
					window.location.reload();
				}
			})
			.catch((error) => {
				console.error(error);
				alert("ログアウトに失敗しました");
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
			<Text>Are you sure you want to log out?</Text>
			<Button
				color={colors.secondaryColor}
				style={{ alignSelf: "flex-end" }}
				onClick={handleSubmit}
			>
				Logout
			</Button>
		</Box>
	);
};
