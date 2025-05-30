import { colors } from "@/styles/colors";
import { Flex, Text } from "@mantine/core";

export type UserProfileCounterProps = {
	value: number;
	label: string;
};

export const UserProfileCounter = (props: UserProfileCounterProps) => {
	return (
		<Flex direction="row" align="center" gap={4}>
			<Text size="md" fw={500} c={colors.black}>
				{props.value}
			</Text>
			<Text size="md" c={colors.black}>
				{props.label}
			</Text>
		</Flex>
	);
};
