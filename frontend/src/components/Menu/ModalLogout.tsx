import { postApiV1AuthLogout } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Button, Flex, Text, Title } from "@mantine/core";

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
		<Flex direction="column" gap={12}>
			<Title size="h4" fw={500} c={colors.secondaryColor}>
				Logout
			</Title>
			<Text>Are you sure you want to log out?</Text>
			<Button
				color={colors.secondaryColor}
				style={{ alignSelf: "flex-end" }}
				onClick={handleSubmit}
			>
				Logout
			</Button>
		</Flex>
	);
};
