import type { ReactNode } from "react";

import { MantineProvider } from "@mantine/core";
import { render, screen } from "@testing-library/react";

import type { SimpleUser } from "@/openapi/schemas";

import { AuthorProfileCard } from "./AuthorProfileCard";

vi.mock("next/link", () => ({
	default: ({ children, href }: { children: ReactNode; href: string }) => (
		<a href={href}>{children}</a>
	),
}));

const createUser = (overrides: Partial<SimpleUser> = {}): SimpleUser => ({
	display_name: "Alice Example",
	icon: "/avatar.png",
	protocol: "atp",
	username: "alice.example",
	...overrides,
});

describe("AuthorProfileCard", () => {
	// 正常系: 表示名とユーザー名が表示されリンクが正しいことを確認する
	test("renders display name, username, and profile link", () => {
		render(
			<MantineProvider>
				<AuthorProfileCard {...createUser()} />
			</MantineProvider>,
		);

		expect(screen.getByText("Alice Example")).toBeInTheDocument();
		expect(screen.getAllByText("alice.example")).toHaveLength(2);
		expect(screen.getByRole("link")).toHaveAttribute(
			"href",
			"/profile/alice.example",
		);
	});

	// 正常系: 長いユーザー名がモバイル表示では省略されることを確認する
	test("truncates long usernames for mobile view", () => {
		const longUsername = "averylongusername_that_should_be_cut_off_here";
		render(
			<MantineProvider>
				<AuthorProfileCard {...createUser({ username: longUsername })} />
			</MantineProvider>,
		);

		expect(
			screen.getByText("averylongusername_that_should_..."),
		).toBeInTheDocument();
		expect(screen.getByText(longUsername)).toBeInTheDocument();
	});

	// 準異常系: ユーザー名が空の場合でもコンポーネントが描画されることを確認する
	test("renders even when username is empty", () => {
		render(
			<MantineProvider>
				<AuthorProfileCard {...createUser({ username: "" })} />
			</MantineProvider>,
		);

		expect(screen.getByRole("link")).toHaveAttribute("href", "/profile/");
	});
});
