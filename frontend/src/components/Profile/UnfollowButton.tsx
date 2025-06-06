import { deleteApiV1Follows } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Button } from "@mantine/core";

export const UnfollowButton = (props: { targetId: string }) => {
	const handleUnfollow = () => {
		const newFollow = {
			target_id: props.targetId,
		};
		deleteApiV1Follows(newFollow)
			.then(() => {
				window.location.reload();
			})
			.catch((error) => {
				alert(error);
			});
	};

	return (
		<Button
			color={colors.secondaryColor}
			size="md"
			style={{ borderRadius: "12px" }}
			onClick={handleUnfollow}
		>
			Unfollow
		</Button>
	);
};
