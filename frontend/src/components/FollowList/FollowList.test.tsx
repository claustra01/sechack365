import { screen } from "@testing-library/react";

import type { SimpleUser } from "@/openapi/schemas";

import { buildSimpleUser } from "@/testutils/builders";
import { renderWithMantine } from "@/testutils/renderWithMantine";

import { FollowList } from "./FollowList";

vi.mock("./SimpleUserCard", () => ({
	SimpleUserCard: ({ display_name }: { display_name: string }) => (
		<div data-testid="simple-user-card">{display_name}</div>
	),
}));

describe("FollowList", () => {
	// 正常系: ユーザー配列を受け取り各ユーザーが描画されることを確認する
	test("正常系: ユーザー一覧を描画する", () => {
		renderWithMantine(
			<FollowList
				users={[
					buildSimpleUser({
						display_name: "Dave Example",
						username: "dave.example",
					}),
					buildSimpleUser({
						display_name: "Eve Example",
						username: "eve",
					}),
				]}
			/>,
		);

		expect(screen.getAllByTestId("simple-user-card")).toHaveLength(2);
		expect(screen.getByText("Dave Example")).toBeInTheDocument();
		expect(screen.getByText("Eve Example")).toBeInTheDocument();
	});

	// 準異常系: ユーザー配列が未定義の場合は描画されないことを確認する
	test("準異常系: usersがundefinedの場合はnullを返す", () => {
		const { container } = renderWithMantine(
			<FollowList users={undefined as unknown as SimpleUser[]} />,
		);

		expect(container.querySelector("div")).toBeNull();
	});

	// 異常系: 無効なエントリを含む場合はエラーになることを確認する
	test("異常系: 無効なユーザーデータを含むとエラーになる", () => {
		const consoleErrorSpy = vi
			.spyOn(console, "error")
			.mockImplementation(() => {});

		try {
			expect(() =>
				renderWithMantine(
					<FollowList
						users={[
							buildSimpleUser({
								display_name: "Dave Example",
								username: "dave.example",
							}),
							null as unknown as SimpleUser,
						]}
					/>,
				),
			).toThrow();
		} finally {
			consoleErrorSpy.mockRestore();
		}
	});
});
