import { postApiV1Posts } from "@/openapi/api";
import type { Newpost } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Button, Flex, Textarea } from "@mantine/core";
import { useState } from "react";

export const NewPost = () => {
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
				window.location.href = "/";
			})
			.catch((error) => {
				alert(error);
			});
	};

	return (
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
	);
};
