import { IconHome, IconLogout, IconUser } from "@tabler/icons-react"
import { MenuItem, MenuItemProps } from "./MenuItem"
import { Box } from "@mantine/core"

export const UserMenu = () => {
  const props: MenuItemProps[] = [
    {
      icon: <IconHome />,
      title: "Home",
      href: "/",
    },
    {
      icon: <IconUser />,
      title: "My Profile",
      href: "/profile",
    },
    {
      icon: <IconLogout />,
      title: "Logout",
      href: "/logout",
    }
  ]
  return (
    <Box>
      {props.map((item, index) => (
        <MenuItem key={index} {...item} />
      ))}
    </Box>
  )
}