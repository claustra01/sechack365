import { screen } from "@testing-library/react";

import "@/testutils/mockNextLink";
import { buildSimpleUser } from "@/testutils/builders";
import { renderWithMantine } from "@/testutils/renderWithMantine";

import { PostUserCard } from "./PostUserCard";

describe("PostUserCard", () => {
	// 正常系: 表示名とユーザー名が表示され正しいプロフィールリンクを指すことを確認する
	test("正常系: ユーザー情報を表示しプロフィールリンクが正しい", () => {
		renderWithMantine(
			<PostUserCard
				{...buildSimpleUser({
					display_name: "Bob Example",
					username: "bob.example",
				})}
			/>,
		);

		expect(screen.getByText("Bob Example")).toBeInTheDocument();
		expect(screen.getAllByText("bob.example")).toHaveLength(2);
		expect(screen.getByRole("link")).toHaveAttribute(
			"href",
			"/profile/bob.example",
		);
	});

	// 正常系: 長いユーザー名がモバイル表示では省略されることを確認する
	test("正常系: 長いユーザー名はモバイル表示で省略される", () => {
		const longUsername = "averylongusername_that_should_be_cut_off_here";
		renderWithMantine(
			<PostUserCard
				{...buildSimpleUser({
					display_name: "Bob Example",
					username: longUsername,
				})}
			/>,
		);

		expect(
			screen.getByText("averylongusername_that_should_..."),
		).toBeInTheDocument();
		expect(screen.getByText(longUsername)).toBeInTheDocument();
	});

	// 準異常系: アイコンが欠損したデータでも描画が継続する
	test("準異常系: アイコンが欠損していても描画される", () => {
		const { container } = renderWithMantine(
			<PostUserCard
				{...buildSimpleUser({
					display_name: "Bob Example",
					username: "bob.example",
					icon: "",
				})}
			/>,
		);

		expect(container.querySelector("img")).toBeNull();
		expect(screen.getByText("Bob Example")).toBeInTheDocument();
	});
});
