import { screen } from "@testing-library/react";

import "@/testutils/mockNextLink";
import { buildSimpleUser } from "@/testutils/builders";
import { renderWithMantine } from "@/testutils/renderWithMantine";

import { AuthorProfileCard } from "./AuthorProfileCard";

describe("AuthorProfileCard", () => {
	// 正常系: 表示名とユーザー名が表示されリンクが正しいことを確認する
	test("正常系: ユーザー情報を表示しプロフィールリンクが正しい", () => {
		renderWithMantine(
			<AuthorProfileCard
				{...buildSimpleUser({
					display_name: "Alice Example",
					username: "alice.example",
				})}
			/>,
		);

		expect(screen.getByText("Alice Example")).toBeInTheDocument();
		expect(screen.getAllByText("alice.example")).toHaveLength(2);
		expect(screen.getByRole("link")).toHaveAttribute(
			"href",
			"/profile/alice.example",
		);
	});

	// 正常系: 長いユーザー名がモバイル表示では省略されることを確認する
	test("正常系: 長いユーザー名はモバイル表示で省略される", () => {
		const longUsername = "averylongusername_that_should_be_cut_off_here";
		renderWithMantine(
			<AuthorProfileCard
				{...buildSimpleUser({
					display_name: "Alice Example",
					username: longUsername,
				})}
			/>,
		);

		expect(
			screen.getByText("averylongusername_that_should_..."),
		).toBeInTheDocument();
		expect(screen.getByText(longUsername)).toBeInTheDocument();
	});

	// 準異常系: アイコンが欠損していても描画が継続する
	test("準異常系: アイコンが欠損していても描画される", () => {
		const { container } = renderWithMantine(
			<AuthorProfileCard
				{...buildSimpleUser({
					display_name: "Alice Example",
					username: "alice.example",
					icon: "",
				})}
			/>,
		);

		expect(container.querySelector("img")).toBeNull();
		expect(screen.getByText("Alice Example")).toBeInTheDocument();
	});
});
