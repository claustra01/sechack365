"use client";
import { Header } from "@/components/Header/Header";
import { PageTemplate } from "@/components/Template/PageTemplate";
import { postApiV1Articles, postApiV1ImagesUpload } from "@/openapi/api";
import type { NewArticle } from "@/openapi/schemas";
import { colors } from "@/styles/colors";
import { Button, Flex, TextInput, Textarea } from "@mantine/core";
import { IconArrowBack } from "@tabler/icons-react";
import { useState } from "react";

export default function NewArticlePage() {
	const [title, setTitle] = useState<string>("");
	const [content, setContent] = useState<string>("");

	const handleChangeTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
		setTitle(e.currentTarget.value);
	};

	const handleChangeContent = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
		setContent(e.currentTarget.value);
	};

	const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const files = e.target.files;
		if (!files) return;
		const file = Array.from(files)[0];
		postApiV1ImagesUpload({ image: file })
			.then(() => {
				alert("Uploaded");
			})
			.catch((error) => {
				alert(error);
			});
	};

	const handleSubmit = () => {
		if (title.length === 0 || content.length === 0) {
			return;
		}
		const newArticle: NewArticle = {
			title: title,
			content: content,
		};
		postApiV1Articles(newArticle)
			.then(() => {
				// FIXME: routerを使うようにする
				window.location.href = "/";
			})
			.catch((error) => {
				alert(error);
			});
	};

	return (
		<PageTemplate>
			<Header title={"New Article"} icon={<IconArrowBack />} />
			<Flex direction="column" p={12} gap={12}>
				<TextInput
					label="Title"
					placeholder="Title"
					onChange={handleChangeTitle}
				/>
				<Textarea
					label="Content"
					placeholder="Your content here..."
					autosize
					minRows={12}
					onChange={handleChangeContent}
				/>
				<Flex justify="space-between">
					<Button variant="outline" color={colors.secondaryColor}>
						<input
							type="file"
							multiple
							accept="image/*"
							onChange={handleFileChange}
						/>
					</Button>
					<Button
						disabled={title.length === 0 || content.length === 0}
						color={colors.secondaryColor}
						style={{ alignSelf: "flex-end" }}
						onClick={handleSubmit}
					>
						Create Article
					</Button>
				</Flex>
			</Flex>
		</PageTemplate>
	);
}
