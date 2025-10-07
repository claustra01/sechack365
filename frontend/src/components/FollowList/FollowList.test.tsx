import type { ReactNode } from "react";

import { MantineProvider } from "@mantine/core";
import { render, screen } from "@testing-library/react";

import type { SimpleUser } from "@/openapi/schemas";

import { FollowList } from "./FollowList";

vi.mock("./SimpleUserCard", () => ({
	SimpleUserCard: ({ display_name }: { display_name: string }) => (
		<div data-testid="simple-user-card">{display_name}</div>
	),
}));

const renderWithMantine = (ui: ReactNode) => {
	return render(<MantineProvider>{ui}</MantineProvider>);
};

const createUser = (overrides: Partial<SimpleUser> = {}): SimpleUser => ({
	display_name: "Dave Example",
	icon: "/avatar.png",
	protocol: "atp",
	username: "dave.example",
	...overrides,
});

describe("FollowList", () => {
	// 正常系: ユーザー配列を受け取り各ユーザーが描画されることを確認する
	test("正常系: ユーザー一覧を描画する", () => {
		renderWithMantine(
			<FollowList
				users={[
					createUser(),
					createUser({ username: "eve", display_name: "Eve Example" }),
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
					<FollowList users={[createUser(), null as unknown as SimpleUser]} />,
				),
			).toThrow();
		} finally {
			consoleErrorSpy.mockRestore();
		}
	});
});
