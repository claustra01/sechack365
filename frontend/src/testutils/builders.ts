import type { SimpleUser } from "@/openapi/schemas";

/**
 * SimpleUser用のテストデータを生成する。
 */
export const buildSimpleUser = (
	overrides: Partial<SimpleUser> = {},
): SimpleUser => {
	return {
		display_name: "Test User",
		icon: "/avatar.png",
		protocol: "atp",
		username: "test.user",
		...overrides,
	};
};
