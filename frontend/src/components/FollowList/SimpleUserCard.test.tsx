import type { ReactNode } from "react";

import { MantineProvider } from "@mantine/core";
import { render, screen } from "@testing-library/react";

import type { SimpleUser } from "@/openapi/schemas";

import { SimpleUserCard } from "./SimpleUserCard";

vi.mock("next/link", () => ({
	default: ({ children, href }: { children: ReactNode; href: string }) => (
		<a href={href}>{children}</a>
	),
}));

const renderWithMantine = (ui: ReactNode) => {
	return render(<MantineProvider>{ui}</MantineProvider>);
};

const createUser = (overrides: Partial<SimpleUser> = {}): SimpleUser => ({
	display_name: "Carol Example",
	icon: "/avatar.png",
	protocol: "atp",
	username: "carol.example",
	...overrides,
});

describe("SimpleUserCard", () => {
	// 正常系: 表示名とユーザー名が表示され正しいプロフィールリンクを指すことを確認する
	test("正常系: ユーザー情報を表示しプロフィールリンクが正しい", () => {
		renderWithMantine(<SimpleUserCard {...createUser()} />);

		expect(screen.getByText("Carol Example")).toBeInTheDocument();
		expect(screen.getAllByText("carol.example")).toHaveLength(2);
		expect(screen.getByRole("link")).toHaveAttribute(
			"href",
			"/profile/carol.example",
		);
	});

	// 正常系: 長いユーザー名がモバイル表示では省略されることを確認する
	test("正常系: 長いユーザー名はモバイル表示で省略される", () => {
		const longUsername = "averylongusername_that_should_be_cut_off_here";
		renderWithMantine(
			<SimpleUserCard {...createUser({ username: longUsername })} />,
		);

		expect(
			screen.getByText("averylongusername_that_should_..."),
		).toBeInTheDocument();
		expect(screen.getByText(longUsername)).toBeInTheDocument();
	});

	// 準異常系: アイコンが欠損したデータでも描画が継続する
	test("準異常系: アイコンが欠損していても描画される", () => {
		const { container } = renderWithMantine(
			<SimpleUserCard {...createUser({ icon: "" })} />,
		);

		expect(container.querySelector("img")).toBeNull();
		expect(screen.getByText("Carol Example")).toBeInTheDocument();
	});
});
