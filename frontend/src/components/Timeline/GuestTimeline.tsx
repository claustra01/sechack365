import { getApiV1Timeline } from "@/openapi";
import { Post } from "@/openapi/schemas";
import { useState } from "react";
import { Header } from "../Header/Header";
import { IconHome } from "@tabler/icons-react";
import { Box } from "@mantine/core";
import { PostCard } from "../PostCard/PostCard";

export const GuestTimeline = () => {
  const [posts, setPosts] = useState<Post[]>([]);

  getApiV1Timeline().then((response) => {
    setPosts(response.data as unknown as Post[]);
  })

  return(
    <Box w={720}>
      <Header title={"Home"} icon={<IconHome />} />
      {posts.map((post) => {
        return (
          <PostCard key={post.id} {...post} />
        )
      })}
    </Box>
  )
}
