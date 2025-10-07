import type { ReactNode } from "react";

import { MantineProvider } from "@mantine/core";
import { render, screen, waitFor } from "@testing-library/react";

import { getApiV1FollowsFollowingId } from "@/openapi/api";
import type { User } from "@/openapi/schemas";

import { CurrentUserContext } from "../Template/PageTemplate";
import { UserProfileCard } from "./UserProfileCard";

vi.mock("@/openapi/api", () => ({
	getApiV1FollowsFollowingId: vi.fn(),
}));

vi.mock("./FollowButton", () => ({
	FollowButton: ({ targetId }: { targetId: string }) => (
		<button type="button" data-testid="follow-button">
			{`Follow ${targetId}`}
		</button>
	),
}));

vi.mock("./UnfollowButton", () => ({
	UnfollowButton: ({ targetId }: { targetId: string }) => (
		<button type="button" data-testid="unfollow-button">
			{`Unfollow ${targetId}`}
		</button>
	),
}));

const mockGetFollow = vi.mocked(getApiV1FollowsFollowingId);

const baseUser: User = {
	id: "user-1",
	username: "example.user",
	display_name: "Example User",
	icon: "/avatar.png",
	profile: "Hello",
	post_count: 10,
	follow_count: 20,
	follower_count: 30,
	protocol: "activitypub",
	created_at: "2024-01-01T00:00:00Z",
	updated_at: "2024-01-02T00:00:00Z",
	identifiers: {
		activitypub: {
			host: "example.com",
			local_username: "example.user",
			public_key: "pubkey",
		},
		nostr: {
			npub: "npub1example",
		},
	},
};

const renderWithProviders = (
	ui: ReactNode,
	currentUser: User | null = baseUser,
) => {
	return render(
		<MantineProvider>
			<CurrentUserContext.Provider
				value={{ user: currentUser, setUser: vi.fn() }}
			>
				{ui}
			</CurrentUserContext.Provider>
		</MantineProvider>,
	);
};

describe("UserProfileCard", () => {
	beforeEach(() => {
		mockGetFollow.mockResolvedValue({ data: { found: false } } as never);
	});

	afterEach(() => {
		vi.clearAllMocks();
	});

	// 正常系: 自分自身のプロフィールを表示する際にEditボタンが表示される
	test("正常系: 現在のユーザー本人ならEditボタンを表示する", async () => {
		renderWithProviders(<UserProfileCard {...baseUser} />);

		await waitFor(() =>
			expect(mockGetFollow).toHaveBeenCalledWith(baseUser.id),
		);

		expect(screen.getByRole("button", { name: "Edit" })).toBeInTheDocument();
	});

	// 準異常系: ログインしていない場合でも描画が継続しフォローボタンが表示されない
	test("準異常系: 現在のユーザーがnullならフォロー関連ボタンを表示しない", async () => {
		renderWithProviders(<UserProfileCard {...baseUser} />, null);

		await waitFor(() =>
			expect(mockGetFollow).toHaveBeenCalledWith(baseUser.id),
		);

		expect(screen.queryByTestId("follow-button")).toBeNull();
		expect(screen.queryByTestId("unfollow-button")).toBeNull();
		expect(screen.queryByRole("button", { name: "Edit" })).toBeNull();
	});

	// 異常系: 想定しないユーザーデータ（identifiers欠損）が渡されるとエラーになる
	test("異常系: identifiersが欠損したユーザーを渡すとエラーになる", () => {
		const corruptedUser = {
			...baseUser,
			identifiers: undefined,
		} as unknown as User;

		expect(() =>
			renderWithProviders(<UserProfileCard {...corruptedUser} />),
		).toThrow();
	});
});
