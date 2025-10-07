import { screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import { postApiV1Follows } from "@/openapi/api";
import { renderWithMantine } from "@/testutils/renderWithMantine";

import { FollowButton } from "./FollowButton";

vi.mock("@/openapi/api", () => ({
	postApiV1Follows: vi.fn(),
}));

const user = userEvent.setup();
const mockPostFollow = vi.mocked(postApiV1Follows);
const originalLocation = window.location;
const alertSpy = vi.spyOn(window, "alert").mockImplementation(() => {});
let reloadMock: ReturnType<typeof vi.fn>;

describe("FollowButton", () => {
	beforeEach(() => {
		mockPostFollow.mockReset();
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

	// 正常系: フォローボタンを押すとフォローAPIを呼び出してリロードする
	test("正常系: クリック時にフォローAPIを呼び出しリロードする", async () => {
		mockPostFollow.mockResolvedValue({} as never);

		renderWithMantine(<FollowButton targetId="user-2" />);

		await user.click(screen.getByRole("button", { name: "Follow" }));

		await waitFor(() =>
			expect(mockPostFollow).toHaveBeenCalledWith({ target_id: "user-2" }),
		);
		expect(reloadMock).toHaveBeenCalled();
	});

	// 準異常系: APIエラー時はalertが呼び出されることを確認する
	test("準異常系: APIが失敗した場合はalertを表示する", async () => {
		const error = new Error("network error");
		mockPostFollow.mockRejectedValue(error);

		renderWithMantine(<FollowButton targetId="user-3" />);

		await user.click(screen.getByRole("button", { name: "Follow" }));

		await waitFor(() => expect(mockPostFollow).toHaveBeenCalledTimes(1));
		expect(alertSpy).toHaveBeenCalledWith(error);
	});
});
