import { postApiV1Posts } from "@/openapi/api";
import type { Newpost } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Box, Button, Flex, Modal, Textarea, Title } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import type React from "react";
import { useContext, useState } from "react";
import { CurrentUserContext } from "../Template/PageTemplate";

export const NewPostModal = () => {
	const { user } = useContext(CurrentUserContext);
	const [opened, { open, close }] = useDisclosure(false);
	const [text, setText] = useState<string>("");

	const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
		setText(e.currentTarget.value);
	};

	const handleSubmit = () => {
		if (text.length === 0) {
			return;
		}
		const newPost: Newpost = {
			content: text,
		};
		postApiV1Posts(newPost)
			.then(() => {
				// FIXME: routerを使うようにする
				window.location.reload();
			})
			.catch((error) => {
				alert(error);
			});
	};

	if (!user) {
		return null;
	}

	return (
		<Box p={4}>
			<Modal opened={opened} onClose={close} title="New Post">
				<Flex direction="column" p={24} gap={24}>
					<Textarea
						placeholder="Your content here..."
						autosize
						minRows={4}
						onChange={handleChange}
					/>
					<Button
						disabled={text.length === 0}
						color={colors.secondaryColor}
						style={{ alignSelf: "flex-end" }}
						onClick={handleSubmit}
					>
						New Post
					</Button>
				</Flex>
			</Modal>

			<Button
				p="xs"
				color={colors.primaryColor}
				onClick={open}
				style={{
					borderRadius: "12px",
					boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)",
				}}
			>
				<Title size="h4" fw={500} c={colors.secondaryColor}>
					New Post
				</Title>
			</Button>
		</Box>
	);
};
