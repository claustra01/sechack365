"use client";

import { postApiV1AuthLogout } from "@/openapi/api";
import { useRouter } from "next/navigation";

export default function LogoutPage() {
	const router = useRouter();

	postApiV1AuthLogout()
		.then((response) => {
			if (response.status === 204) {
				// FIXME: routerを使うようにする
				router.push("/");
			}
		})
		.catch((error) => {
			console.error(error);
			alert("ログアウトに失敗しました");
		});

	return <></>;
}
