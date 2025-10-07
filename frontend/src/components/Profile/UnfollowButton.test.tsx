import { screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import { deleteApiV1Follows } from "@/openapi/api";
import { renderWithMantine } from "@/testutils/renderWithMantine";

import { UnfollowButton } from "./UnfollowButton";

vi.mock("@/openapi/api", () => ({
	deleteApiV1Follows: vi.fn(),
}));

const user = userEvent.setup();
const mockDeleteFollow = vi.mocked(deleteApiV1Follows);
const originalLocation = window.location;
const alertSpy = vi.spyOn(window, "alert").mockImplementation(() => {});
let reloadMock: ReturnType<typeof vi.fn>;

describe("UnfollowButton", () => {
	beforeEach(() => {
		mockDeleteFollow.mockReset();
		reloadMock = vi.fn();
		const newLocation = Object.create(originalLocation);
		Object.defineProperty(newLocation, "reload", {
			configurable: true,
			value: reloadMock,
		});
		Object.defineProperty(window, "location", {
			configurable: true,
			value: newLocation,
		});
		alertSpy.mockClear();
	});

	afterEach(() => {
		Object.defineProperty(window, "location", {
			configurable: true,
			value: originalLocation,
		});
	});

	afterAll(() => {
		alertSpy.mockRestore();
	});

	// 正常系: アンフォローボタンを押すとアンフォローAPIを呼び出してリロードする
	test("正常系: クリック時にアンフォローAPIを呼び出しリロードする", async () => {
		mockDeleteFollow.mockResolvedValue({} as never);

		renderWithMantine(<UnfollowButton targetId="user-4" />);

		await user.click(screen.getByRole("button", { name: "Unfollow" }));

		await waitFor(() =>
			expect(mockDeleteFollow).toHaveBeenCalledWith({ target_id: "user-4" }),
		);
		expect(reloadMock).toHaveBeenCalled();
	});

	// 準異常系: APIエラー時はalertが呼び出されることを確認する
	test("準異常系: APIが失敗した場合はalertを表示する", async () => {
		const error = new Error("network error");
		mockDeleteFollow.mockRejectedValue(error);

		renderWithMantine(<UnfollowButton targetId="user-5" />);

		await user.click(screen.getByRole("button", { name: "Unfollow" }));

		await waitFor(() => expect(mockDeleteFollow).toHaveBeenCalledTimes(1));
		expect(alertSpy).toHaveBeenCalledWith(error);
	});
});
