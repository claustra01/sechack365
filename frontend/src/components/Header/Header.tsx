import { Box, ActionIcon, Title } from "@mantine/core"
import { IconArrowBackUp } from "@tabler/icons-react"
import Link from "next/link"

export type HeaderProps = {
  title: string
  icon: JSX.Element
}

export const Header = (props: HeaderProps) => {
  return (
    <Box
    bg="blue"
    style={{ display: "flex", alignItems: "center", padding: "24px" }}
  >
    <ActionIcon
      component={Link}
      href="/"
      variant="subtle"
      size="xl"
      c="white"
    >
    {props.icon}
    </ActionIcon>
    <Title size="h3" fw={500} c="white">
      {props.title}
    </Title>
  </Box>

  )
}