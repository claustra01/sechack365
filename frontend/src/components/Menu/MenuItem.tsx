import { ActionIcon, Box, Text } from "@mantine/core";
import Link from "next/link";

export type MenuItemProps = {
  icon: JSX.Element;
  title: string;
  href: string;
}

export const MenuItem = (props: MenuItemProps) => {
  return (
    <Box style={{ display: "flex", alignItems: "center" }}>
      <ActionIcon component={Link} href={props.href} variant="subtle" size="xl" color="blue">
        {props.icon}
      </ActionIcon>
      <Text size="xl" fw={500} c="blue">{props.title}</Text>
    </Box>
  )
}
