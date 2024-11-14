import { Box, Text } from "@mantine/core";

export type UserProfileCounterProps = {
	value: number;
	label: string;
};

export const UserProfileCounter = (props: UserProfileCounterProps) => {
	return (
		<Box style={{ display: "flex", alignItems: "center", gap: "4px" }}>
			<Text size="lg" fw={500}>
				{props.value}
			</Text>
			<Text size="lg">{props.label}</Text>
		</Box>
	);
};
