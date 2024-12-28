import { Flex, Text } from "@mantine/core";

export type UserProfileCounterProps = {
	value: number;
	label: string;
};

export const UserProfileCounter = (props: UserProfileCounterProps) => {
	return (
		<Flex direction="row" align="center" gap={4}>
			<Text size="lg" fw={500}>
				{props.value}
			</Text>
			<Text size="lg">{props.label}</Text>
		</Flex>
	);
};
