import { postApiV1Follows } from "@/openapi/api";
import { colors } from "@/styles/colors";
import { Button } from "@mantine/core";

export const FollowButton = (props: { targetId: string }) => {
	const handleFollow = () => {
		const newFollow = {
			target_id: props.targetId,
		};
		postApiV1Follows(newFollow)
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
			onClick={handleFollow}
		>
			Follow
		</Button>
	);
};
