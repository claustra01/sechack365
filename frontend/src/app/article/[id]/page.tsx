"use client";

import { AuthorProfileCard } from "@/components/Article/AuthorProfileCard";
import { Header } from "@/components/Header/Header";
import { PostCard } from "@/components/PostCard/PostCard";
import { PageTemplate } from "@/components/Template/PageTemplate";
import { getApiV1ArticlesId, getApiV1ArticlesIdComments } from "@/openapi/api";
import type { Article, ArticleComment } from "@/openapi/schemas";
import { Box, ScrollArea, Title } from "@mantine/core";
import { IconArrowBack } from "@tabler/icons-react";
import { useEffect, useState } from "react";
import Markdown from "react-markdown";

export default function UserProfilePage({
	params,
}: { params: { id: string } }) {
	const { id } = params;
	const [article, setArticle] = useState<Article | null>(null);
	const [comments, setComments] = useState<ArticleComment[]>([]);

	useEffect(() => {
		getApiV1ArticlesId(id).then((res) => {
			setArticle(res.data);
		});
		getApiV1ArticlesIdComments(id).then((res) => {
			setComments(res.data);
		});
	}, [id]);

	if (!article) {
		return (
			<PageTemplate>
				<Header title="" icon={<IconArrowBack />} />
			</PageTemplate>
		);
	}

	return (
		<PageTemplate>
			<Header title={article.title} icon={<IconArrowBack />} />
			<ScrollArea h={"calc( 100vh - 78px )"}>
				<Box p={24}>
					<Title>{article.title}</Title>
					<Box my={24}>
						<AuthorProfileCard {...article.user} />
					</Box>
					<Markdown>{article.content}</Markdown>
				</Box>
				{comments.map((comment) => {
					return <PostCard key={comment.id} {...comment} />;
				})}
			</ScrollArea>
		</PageTemplate>
	);
}
