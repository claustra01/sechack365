import { Post } from "@/openapi/schemas"
import { PostUserCard } from "./PostUserCard"
import { Box, Text } from "@mantine/core"

export const PostCard = (props: Post) => {
  console.log(props)
  return (
    <Box style={{display: "flex", flexDirection: "column", padding: "24px", gap: "24px", borderTop: "2px solid #1C7ED6"}}>
      <PostUserCard {...props.user} />
      <Text size="lg">{props.content}</Text>
    </Box>
  )
}
