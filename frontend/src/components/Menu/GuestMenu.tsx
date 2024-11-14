import { IconHome, IconLogin, IconLogout, IconUser } from "@tabler/icons-react"
import { MenuItem, MenuItemProps } from "./MenuItem"
import { Box } from "@mantine/core"

export const GuestMenu = () => {
  const props: MenuItemProps[] = [
    {
      icon: <IconLogin />,
      title: "Login",
      href: "/login",
    },
    {
      icon: <IconUser />,
      title: "New Account",
      href: "/register",
    },
  ]
  return (
    <Box>
      {props.map((item, index) => (
        <MenuItem key={index} {...item} />
      ))}
    </Box>
  )
}