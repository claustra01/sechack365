import type { ReactNode } from "react";

import { MantineProvider } from "@mantine/core";
import { render, screen } from "@testing-library/react";

import { UserProfileCounter } from "./UserProfileCounter";

const renderWithMantine = (ui: ReactNode) => {
	return render(<MantineProvider>{ui}</MantineProvider>);
};

describe("UserProfileCounter", () => {
	// 正常系: 値とラベルが表示されることを確認する
	test("正常系: 値とラベルを表示する", () => {
		renderWithMantine(<UserProfileCounter value={42} label="Posts" />);

		expect(screen.getByText("42")).toBeInTheDocument();
		expect(screen.getByText("Posts")).toBeInTheDocument();
	});

	// 準異常系: 負の値でも描画が継続することを確認する
	test("準異常系: 負の値でもそのまま表示する", () => {
		renderWithMantine(<UserProfileCounter value={-1} label="Followers" />);

		expect(screen.getByText("-1")).toBeInTheDocument();
		expect(screen.getByText("Followers")).toBeInTheDocument();
	});
});
