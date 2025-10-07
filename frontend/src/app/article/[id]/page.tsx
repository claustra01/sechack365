import { notFound } from "next/navigation";
import { ArticlePageClient } from "./ArticlePageClient";

type PageProps = {
	params?: Promise<{ id: string | string[] | undefined }>;
};

export default async function ArticlePage({ params }: PageProps) {
	const resolvedParams = params ? await params : undefined;
	const rawId = resolvedParams?.id;
	const id = Array.isArray(rawId) ? rawId[0] : rawId;

	if (!id) {
		notFound();
	}

	return <ArticlePageClient id={id} />;
}
