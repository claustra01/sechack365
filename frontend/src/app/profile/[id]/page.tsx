import { notFound } from "next/navigation";
import { UserProfilePageClient } from "./UserProfilePageClient";

type PageProps = {
	params?: Promise<{ id: string | string[] | undefined }>;
};

export default async function UserProfilePage({ params }: PageProps) {
	const resolvedParams = params ? await params : undefined;
	const rawId = resolvedParams?.id;
	const id = Array.isArray(rawId) ? rawId[0] : rawId;

	if (!id) {
		notFound();
	}

	return <UserProfilePageClient id={id} />;
}
