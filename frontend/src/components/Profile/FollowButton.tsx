import { postApiV1FollowsFollow } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Button } from "@mantine/core";

export const FollowButton = (props: { targetId: string }) => {
	const handleFollow = () => {
		const newFollow = {
			target_id: props.targetId,
		};
		postApiV1FollowsFollow(newFollow)
			.then(() => {
				window.location.reload();
			})
			.catch((error) => {
				alert(error);
			});
	};

	return (
		<Button color={colors.secondaryColor} size="lg" onClick={handleFollow}>
			Follow
		</Button>
	);
};
